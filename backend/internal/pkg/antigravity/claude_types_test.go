package antigravity

import (
	"testing"
)

func TestDefaultModels_ContainsNewAndLegacyImageModels(t *testing.T) {
	t.Parallel()

	models := DefaultModels()
	byID := make(map[string]ClaudeModel, len(models))
	for _, m := range models {
		byID[m.ID] = m
	}

	requiredIDs := []string{
		"claude-fable-5",
		"claude-opus-4-8",
		"claude-opus-4-6-thinking",
		"gemini-2.5-flash-image",
		"gemini-2.5-flash-image-preview",
		"gemini-3.1-flash-image",
		"gemini-3.1-flash-image-preview",
		"gemini-3-pro-image", // legacy compatibility
	}

	for _, id := range requiredIDs {
		if _, ok := byID[id]; !ok {
			t.Fatalf("expected model %q to be exposed in DefaultModels", id)
		}
	}
}

func TestDefaultModels_EnvOverride(t *testing.T) {
	// t.Setenv 会在测试结束后自动恢复环境变量

	tests := []struct {
		name    string
		envVal  string
		wantLen int
		wantIDs []string
	}{
		{
			name:    "正常覆盖",
			envVal:  "claude-sonnet-4-6,gemini-3-flash,claude-opus-4-8",
			wantLen: 3,
			wantIDs: []string{"claude-sonnet-4-6", "gemini-3-flash", "claude-opus-4-8"},
		},
		{
			name:    "含空白项被跳过",
			envVal:  "claude-sonnet-4-6, , gemini-3-flash,",
			wantLen: 2,
			wantIDs: []string{"claude-sonnet-4-6", "gemini-3-flash"},
		},
		{
			name:    "单个模型",
			envVal:  "gemini-3-pro-high",
			wantLen: 1,
			wantIDs: []string{"gemini-3-pro-high"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(AntigravityModelsEnv, tt.envVal)
			models := DefaultModels()
			if len(models) != tt.wantLen {
				t.Fatalf("got %d models, want %d", len(models), tt.wantLen)
			}
			byID := make(map[string]bool, len(models))
			for _, m := range models {
				byID[m.ID] = true
			}
			for _, wantID := range tt.wantIDs {
				if !byID[wantID] {
					t.Errorf("expected model %q not found in result", wantID)
				}
			}
		})
	}
}

func TestDefaultModels_EnvEmptyFallback(t *testing.T) {
	// 空字符串应 fallback 到硬编码默认列表
	t.Setenv(AntigravityModelsEnv, "")
	models := DefaultModels()
	if len(models) == 0 {
		t.Fatal("空 env 应 fallback 到默认列表，但返回了空列表")
	}
	// 纯空白也应 fallback
	t.Setenv(AntigravityModelsEnv, "   ")
	models = DefaultModels()
	if len(models) == 0 {
		t.Fatal("纯空白 env 应 fallback 到默认列表，但返回了空列表")
	}
}
