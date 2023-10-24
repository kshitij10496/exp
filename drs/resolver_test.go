package main

import "testing"

func TestResolver(t *testing.T) {
	t.Run("success on resolving example.com via Cloudflare DNS", func(t *testing.T) {
		site, dnsServer := "example.com", "1.1.1.1"
		got, err := resolve(site, dnsServer)
		if err != nil {
			t.Fatalf("expected no error, got err=%v", err)
		}
		if got == "" {
			t.Fatalf("expected some resolution, got=%s", got)
		}
	})

	t.Run("success on resolving example.com via Google DNS", func(t *testing.T) {
		site, dnsServer := "example.com", "8.8.8.8"
		got, err := resolve(site, dnsServer)
		if err != nil {
			t.Fatalf("expected no error, got err=%v", err)
		}
		if got == "" {
			t.Fatalf("expected some resolution, got=%s", got)
		}
	})

	t.Run("return error for invalid DNS server", func(t *testing.T) {
		site, dnsServer := "example.com", "foobar"
		_, err := resolve(site, dnsServer)
		if err == nil {
			t.Fatalf("expected error, got err=%v", err)
		}
	})
}
