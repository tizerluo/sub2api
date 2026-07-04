//go:build unit

package service

import (
	"testing"
)

func TestApplyThinkingModelSuffix(t *testing.T) {
	tests := []struct {
		name            string
		mappedModel     string
		thinkingEnabled bool
		expected        string
	}{
		// thinking 关闭：透传
		{"thinking disabled - passthrough", "gemini-2.5-pro", false, "gemini-2.5-pro"},
		// thinking 开启：仍透传（当前 Antigravity REST 路径无 thinking 后缀转换逻辑）
		{"thinking enabled - passthrough", "gemini-2.5-flash", true, "gemini-2.5-flash"},
		{"thinking enabled - gemini-3.1-flash-lite passthrough", "gemini-3.1-flash-lite", true, "gemini-3.1-flash-lite"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applyThinkingModelSuffix(tt.mappedModel, tt.thinkingEnabled)
			if result != tt.expected {
				t.Errorf("applyThinkingModelSuffix(%q, %v) = %q, want %q",
					tt.mappedModel, tt.thinkingEnabled, result, tt.expected)
			}
		})
	}
}
