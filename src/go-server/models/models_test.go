package models

import (
	"strings"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	const runs = 10000

	tests := []struct {
		name  string
		check func(t *testing.T, code string)
	}{
		{
			name: "length is eight",
			check: func(t *testing.T, code string) {
				if len(code) != codeLength {
					t.Errorf("len(%q) = %d, want %d", code, len(code), codeLength)
				}
			},
		},
		{
			name: "only alphanumeric characters",
			check: func(t *testing.T, code string) {
				for _, c := range code {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("code %q contains %q, which is not in the alphabet", code, c)
					}
				}
			},
		},
	}

	seen := make(map[string]bool, runs)
	codes := make([]string, 0, runs)
	for i := 0; i < runs; i++ {
		code, err := GenerateShortURL()
		if err != nil {
			t.Fatalf("GenerateShortURL() error: %v", err)
		}
		codes = append(codes, code)
		if seen[code] {
			t.Fatalf("duplicate code %q after %d runs", code, i)
		}
		seen[code] = true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, code := range codes {
				tt.check(t, code)
			}
		})
	}
}
