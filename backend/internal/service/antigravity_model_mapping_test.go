//go:build unit

package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAntigravityGatewayServiceGetMappedModel(t *testing.T) {
	svc := &AntigravityGatewayService{}

	tests := []struct {
		name           string
		requestedModel string
		mapping        map[string]any
		expected       string
	}{
		{
			name:           "explicit mapping takes precedence",
			requestedModel: "claude-sonnet-4-6",
			mapping:        map[string]any{"claude-sonnet-4-6": "custom-upstream-model"},
			expected:       "custom-upstream-model",
		},
		{
			name:           "explicit wildcard is allowed",
			requestedModel: "claude-sonnet-4-6",
			mapping:        map[string]any{"claude-*": "custom-upstream-model"},
			expected:       "custom-upstream-model",
		},
		{
			name:           "REST fallback model",
			requestedModel: "gemini-2.5-flash",
			expected:       "gemini-2.5-flash",
		},
		{
			name:           "gRPC-only Claude is absent from fallback",
			requestedModel: "claude-sonnet-4-6",
			expected:       "",
		},
		{
			name:           "unknown model is absent from fallback",
			requestedModel: "gemini-unknown",
			expected:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &Account{Platform: PlatformAntigravity}
			if tt.mapping != nil {
				account.Credentials = map[string]any{"model_mapping": tt.mapping}
			}
			require.Equal(t, tt.expected, svc.getMappedModel(account, tt.requestedModel))
		})
	}
}

func TestAntigravityGatewayServiceIsModelSupportedUsesRESTFallback(t *testing.T) {
	svc := &AntigravityGatewayService{}
	for _, model := range []string{
		"gemini-2.5-pro",
		"gemini-2.5-flash",
		"gemini-2.5-flash-lite",
		"gemini-3.1-flash-lite",
	} {
		require.True(t, svc.IsModelSupported(model), model)
	}
	for _, model := range []string{"claude-sonnet-4-6", "gpt-oss-120b-medium", "gemini-3-flash", ""} {
		require.False(t, svc.IsModelSupported(model), model)
	}
}
