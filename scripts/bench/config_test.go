package main

import (
	"path/filepath"
	"testing"
)

func TestParseConfigResultDirDefaultJSONPath(t *testing.T) {
	dir := t.TempDir()
	cfg, err := parseConfigFrom([]string{"-result-dir", dir})
	if err != nil {
		t.Fatalf("parse config: %v", err)
	}

	want := filepath.Join(dir, "result.json")
	if cfg.jsonOut != want {
		t.Fatalf("jsonOut = %q, want %q", cfg.jsonOut, want)
	}
}
