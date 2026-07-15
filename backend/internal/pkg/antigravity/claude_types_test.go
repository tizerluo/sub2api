package antigravity

import "testing"

func TestDefaultModels_ContainsOnlyRESTFallbacks(t *testing.T) {
	t.Parallel()

	models := DefaultModels()
	byID := make(map[string]ClaudeModel, len(models))
	for _, m := range models {
		byID[m.ID] = m
	}

	requiredIDs := []string{
		"gemini-2.5-pro",
		"gemini-2.5-flash",
		"gemini-2.5-flash-lite",
		"gemini-3.1-flash-lite",
	}

	for _, id := range requiredIDs {
		if _, ok := byID[id]; !ok {
			t.Fatalf("expected model %q to be exposed in DefaultModels", id)
		}
	}
	if len(models) != len(requiredIDs) {
		t.Fatalf("expected %d REST models, got %d", len(requiredIDs), len(models))
	}
}
