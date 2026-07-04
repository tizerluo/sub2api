package antigravity

import (
	"testing"
)

func TestDefaultModels_ContainsActualSupportedModels(t *testing.T) {
	t.Parallel()

	models := DefaultModels()
	byID := make(map[string]ClaudeModel, len(models))
	for _, m := range models {
		byID[m.ID] = m
	}

	// 实测 2026-07-04：Google AI Pro 订阅实际支持的 6 个模型
	requiredIDs := []string{
		"gemini-3.1-pro-high",
		"gemini-3.1-pro-low",
		"gemini-3-flash",
		"claude-sonnet-4-6-thinking",
		"claude-opus-4-6-thinking",
		"gpt-oss-120b-medium",
	}

	for _, id := range requiredIDs {
		if _, ok := byID[id]; !ok {
			t.Fatalf("expected model %q to be exposed in DefaultModels", id)
		}
	}

	// 总数应该是 6 个（2 Claude + 3 Gemini + 1 GPT-OSS）
	if len(models) != 6 {
		t.Fatalf("expected 6 models, got %d: %v", len(models), byID)
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
