package service

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Account-aware service calls must go through DoWithTLS. A nil profile is the
// deliberate standard-Go-TLS fallback; calling Do directly silently bypasses
// an account's explicit TLS profile.
func TestServiceHTTPUpstreamCallsUseTLSHelper(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate static TLS audit test")
	}
	entries, err := os.ReadDir(filepath.Dir(thisFile))
	if err != nil {
		t.Fatalf("read service directory: %v", err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		body, err := os.ReadFile(filepath.Join(filepath.Dir(thisFile), name))
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		if strings.Contains(string(body), "httpUpstream.Do(") {
			t.Errorf("%s bypasses DoWithTLS; use a nil profile for standard Go TLS", name)
		}
	}
}
