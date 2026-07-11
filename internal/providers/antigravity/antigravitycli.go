package antigravity

import (
	"context"
	"deepn/internal/providers"
	"deepn/internal/utils"
	"encoding/json"
	"fmt"
)

type AntigravityCLI struct {
	subproc        *utils.Subproc
	conversationId string
}

func NewAntigravityCLI(subproc *utils.Subproc) (*AntigravityCLI, error) {

	startupStatus := map[string]any{
		"ready":          false,
		"conversationId": nil,
	}

	response := subproc.Run("agy", "-p", "ONLY respond with a JSON object {\"ready\":boolean, \"conversationId\":string} true/false if you are ready and this conversation id. NEVER include any other text")
	if response.ExitCode != 0 {
		return nil, fmt.Errorf("failed to start Antigravity CLI: %s", string(response.Stderr))
	}

	if err := json.Unmarshal(response.Stdout, &startupStatus); err != nil {
		return nil, fmt.Errorf("failed to parse Antigravity CLI startup status: %s", err)
	}

	if ready, ok := startupStatus["ready"].(bool); !ok || !ready {
		return nil, fmt.Errorf("Antigravity CLI is not ready or returned invalid status object")
	}

	if convId, ok := startupStatus["conversationId"].(string); ok {
		return &AntigravityCLI{
			subproc:        subproc,
			conversationId: convId,
		}, nil
	} else {
		return nil, fmt.Errorf("failed to get conversationId from Antigravity CLI startup status")
	}
}

func (a *AntigravityCLI) Chat(ctx context.Context, req providers.Request) (providers.Response, error) {

	response := a.subproc.Run("agy", "-p", req.Message, "--continue", a.conversationId)
	if response.ExitCode != 0 {
		return providers.Response{}, fmt.Errorf("Antigravity CLI chat failed: %s", string(response.Stderr))
	}

	return providers.Response{Message: string(response.Stdout)}, nil
}
