package domain

import "testing"

func TestDefaultAntigravityModelMapping_AgyUIToRestAliases(t *testing.T) {
	t.Parallel()

	// agy UI 标签 → REST 上游模型 ID（实测 2026-07-04）
	cases := map[string]string{
		"gemini-3.1-pro-high":            "gemini-2.5-pro",
		"gemini-3.1-pro-low":             "gemini-2.5-pro",
		"gemini-3-flash":                 "gemini-2.5-flash",
		AntigravityGemini31ProAgentModel: "gemini-2.5-pro",
		"gemini-3.1-pro":                 "gemini-2.5-pro",
		"gemini-3.1-pro-preview":         "gemini-2.5-pro",
		"gemini-3-pro-high":              "gemini-2.5-pro",
		"gemini-3-pro-low":               "gemini-2.5-pro",
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

func TestDefaultAntigravityModelMapping_RestAccessibleModels(t *testing.T) {
	t.Parallel()

	// 实测 2026-07-04：通过 REST streamGenerateContent 可用的 4 个 Gemini 模型
	// （retrieveUserQuota 权威确认），映射到自身
	cases := map[string]string{
		"gemini-2.5-pro":        "gemini-2.5-pro",
		"gemini-2.5-flash":      "gemini-2.5-flash",
		"gemini-2.5-flash-lite": "gemini-2.5-flash-lite",
		"gemini-3.1-flash-lite": "gemini-3.1-flash-lite",
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
