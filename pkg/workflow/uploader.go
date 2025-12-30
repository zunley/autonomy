package workflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zunley/autonomy/pkg/types"
	"net/http"
	"time"
)

type Uploader interface {
	Upload(result *types.RunResult) error
}
type HTTPUploader struct {
	url    string // e.g., "https://core.example.com/api/v1"
	token  string
	client *http.Client
}

func NewHTTPUploader(config *types.ControlNode) *HTTPUploader {
	return &HTTPUploader{
		url:    config.URL,
		token:  config.Token,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (u *HTTPUploader) Upload(result *types.RunResult) error {
	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("marshal result: %w", err)
	}

	req, err := http.NewRequest("POST", u.url+"/runs", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.token)

	resp, err := u.client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("upload failed: HTTP %d", resp.StatusCode)
	}

	return nil
}
