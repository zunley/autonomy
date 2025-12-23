# Agent 设计

## 输入输出

1. 以配置文件启动
2. 创建定时任务
3. 任务执行完成后输出任务结果


## 详细设计

### 配置文件

```
agent_id: "macbook-pro"

control_node:
    url: "http://127.0.0.1:8080"
    token: ""

log_level: "info"
```

### 工作流文件

```
name: "sync-notes"

description: ""

schedule: "0 */6 * * *"

steps:
    - name: "sync notes"
      working_dir: ""
      shell: |
        git add .
        git commit -m "autonomy: sync $(date -Iseconds)"
        git push origin main
```

### 任务结果输出

```json
{
    "agent_id": "macbook-pro",
    "workflow_name": "sync-notes",
    "run_id": "abc123",
    "started_at": 11111111,
    "completed_at": 222222222,
    "status": "success", // "success" | "failed" | "skipped"
}
```

