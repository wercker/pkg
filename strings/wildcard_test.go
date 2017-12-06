//-----------------------------------------------------------------------------
// Copyright (c) 2017 Oracle and/or its affiliates.  All rights reserved.
// This program is free software: you can modify it and/or redistribute it
// under the terms of:
//
// (i)  the Universal Permissive License v 1.0 or at your option, any
//      later version (http://oss.oracle.com/licenses/upl); and/or
//
// (ii) the Apache License v 2.0. (http://www.apache.org/licenses/LICENSE-2.0)
//-----------------------------------------------------------------------------

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
