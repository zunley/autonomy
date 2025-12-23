default: binaries

binaries: autonomy-agentd autonomy-core

autonomy-agentd:

autonomy-agent:

autonomy-core:

lint:
	find -name \*.go | xargs go fmt