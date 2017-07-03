package strings

import "testing"

func TestMatchWildcard(t *testing.T) {
	tests := []struct {
		target   string
		wildcard string
		exp      bool
	}{
		{"pattern", "*", true},
		{"😀", "*", true},
		{"pattern", "*pattern", true},
		{"hey/pattern", "*pattern", true},
		{"pattern/hey", "*pattern", false},
		{"word", "*pattern", false},
		{"pattern", "pattern*", true},
		{"hey/pattern", "pattern*", false},
		{"pattern/hey", "pattern*", true},
		{"word", "pattern*", false},
		{"hey/pattern/hey", "*pattern*", false},
		{"pattern", "pattern", true},
		{"pat|tern", "pat*tern", false},
	}

	for _, tt := range tests {
		out := MatchesWildcard(tt.target, tt.wildcard)
		if out != tt.exp {
			t.Errorf("expect \"%v\" got \"%v\" for \"%v\" and \"%v\"", out, tt.exp, tt.target, tt.wildcard)
		}
	}
}
