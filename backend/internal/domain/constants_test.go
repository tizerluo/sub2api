package domain

import "testing"

func TestDefaultAntigravityModelMapping_ContainsOnlyRESTFallbacks(t *testing.T) {
	t.Parallel()

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
	if len(DefaultAntigravityModelMapping) != len(cases) {
		t.Fatalf("unexpected REST fallback count: got %d want %d", len(DefaultAntigravityModelMapping), len(cases))
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
