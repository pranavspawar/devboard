// Starter test for env() so `go test ./...` does real work in CI. No DB needed.
package main

import (
	"os"
	"testing"
)

func TestEnvReturnsFallbackWhenUnset(t *testing.T) {
	if err := os.Unsetenv("DEVBOARD_TEST_KEY"); err != nil {
		t.Fatalf("Unsetenv failed: %v", err)
	}
	if got := env("DEVBOARD_TEST_KEY", "fallback"); got != "fallback" {
		t.Errorf("env() = %q, want %q", got, "fallback")
	}
}

func TestEnvReturnsValueWhenSet(t *testing.T) {
	if err := os.Setenv("DEVBOARD_TEST_KEY", "real"); err != nil {
		t.Fatalf("Setenv failed: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("DEVBOARD_TEST_KEY"); err != nil {
			t.Logf("Unsetenv failed: %v", err)
		}
	}()
	if got := env("DEVBOARD_TEST_KEY", "fallback"); got != "real" {
		t.Errorf("env() = %q, want %q", got, "real")
	}
}
