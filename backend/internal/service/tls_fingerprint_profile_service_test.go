package service

import (
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/model"
)

func TestResolveTLSProfilePlatformDefaults(t *testing.T) {
	svc := &TLSFingerprintProfileService{localCache: map[int64]*model.TLSFingerprintProfile{
		7: {ID: 7, Name: "Measured custom profile"},
	}}

	if got := svc.ResolveTLSProfile(&Account{Platform: PlatformGrok}); got != nil {
		t.Fatalf("disabled account resolved profile %q", got.Name)
	}
	if got := svc.ResolveTLSProfile(&Account{Platform: PlatformGrok, Extra: map[string]any{"enable_tls_fingerprint": true}}); got != nil {
		t.Fatalf("Grok without an explicit profile resolved %q", got.Name)
	}
	if got := svc.ResolveTLSProfile(&Account{Platform: PlatformAnthropic, Type: AccountTypeOAuth, Extra: map[string]any{"enable_tls_fingerprint": true}}); got == nil || got.Name != "Built-in Default (Node.js 24.x)" {
		t.Fatalf("Anthropic default profile = %#v", got)
	}
	if got := svc.ResolveTLSProfile(&Account{Platform: PlatformAntigravity, Extra: map[string]any{"enable_tls_fingerprint": true, "tls_fingerprint_profile_id": 7}}); got == nil || got.Name != "Measured custom profile" {
		t.Fatalf("explicit Antigravity profile = %#v", got)
	}
}
