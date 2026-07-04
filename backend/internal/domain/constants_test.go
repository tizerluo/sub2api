package domain

import "testing"

func TestDefaultAntigravityModelMapping_LegacyCompatibilityAliases(t *testing.T) {
	t.Parallel()

	// 旧模型名映射到实际支持的模型
	cases := map[string]string{
		"gemini-2.5-flash":      "gemini-3-flash",
		"gemini-2.5-pro":        AntigravityGemini31ProAgentModel,
		"gemini-3-pro-high":     AntigravityGemini31ProAgentModel,
		"claude-opus-4-8":       "claude-opus-4-6-thinking",
		"claude-fable-5":        "claude-opus-4-6-thinking",
		"claude-haiku-4-5":      "claude-sonnet-4-6-thinking",
		"claude-sonnet-4-5-20250929": "claude-sonnet-4-6-thinking",
	}

	for from, want := range cases {
		got, ok := DefaultAntigravityModelMapping[from]
		if !ok {
			t.Fatalf("expected mapping for %q to exist", from)
		}
		if got != want {
			t.Fatalf("unexpected mapping for %q: got %q want %q", from, got, want)
		}
	}
}

func TestDefaultAntigravityModelMapping_ActualSupportedModels(t *testing.T) {
	t.Parallel()

	// 实测 2026-07-04：实际支持的 6 个模型必须存在且映射到自身（或上游路由名）
	cases := map[string]string{
		"gemini-3.1-pro-high":       AntigravityGemini31ProAgentModel,
		"gemini-3.1-pro-low":        "gemini-3.1-pro-low",
		"gemini-3-flash":            "gemini-3-flash",
		"claude-sonnet-4-6-thinking": "claude-sonnet-4-6-thinking",
		"claude-opus-4-6-thinking":   "claude-opus-4-6-thinking",
		"gpt-oss-120b-medium":       "gpt-oss-120b-medium",
	}
	for from, want := range cases {
		got, ok := DefaultAntigravityModelMapping[from]
		if !ok {
			t.Fatalf("expected mapping for %q to exist", from)
		}
		if got != want {
			t.Fatalf("unexpected mapping for %q: got %q want %q", from, got, want)
		}
	}
}

func TestDefaultAntigravityModelMapping_Gemini31ProAliases(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		AntigravityGemini31ProAgentModel: AntigravityGemini31ProAgentModel,
		"gemini-3.1-pro":                 AntigravityGemini31ProAgentModel,
		"gemini-3.1-pro-high":            AntigravityGemini31ProAgentModel,
		"gemini-3.1-pro-preview":         AntigravityGemini31ProAgentModel,
		"gemini-3.1-pro-low":             "gemini-3.1-pro-low",
	}

	for from, want := range cases {
		got, ok := DefaultAntigravityModelMapping[from]
		if !ok {
			t.Fatalf("expected mapping for %q to exist", from)
		}
		if got != want {
			t.Fatalf("unexpected mapping for %q: got %q want %q", from, got, want)
		}
	}
}

func TestDefaultBedrockModelMapping_ContainsNewClaudeModels(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"claude-fable-5":  "anthropic.claude-fable-5",
		"claude-opus-4-8": "us.anthropic.claude-opus-4-8-v1",
	}
	for from, want := range cases {
		got, ok := DefaultBedrockModelMapping[from]
		if !ok {
			t.Fatalf("expected Bedrock mapping for %q to exist", from)
		}
		if got != want {
			t.Fatalf("unexpected Bedrock mapping for %q: got %q want %q", from, got, want)
		}
	}
}
