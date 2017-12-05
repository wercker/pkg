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

import (
	"strings"
)

// MatchesWildcard returns true if the target match the wildcard where value
// can be like "*pattern", "pattern*" or "pattern".
func MatchesWildcard(target, wildcard string) bool {

	if target == "" || wildcard == "" {
		return false
	}

	if wildcard == "*" || wildcard == target {
		return true
	}

	wildcardIndex := strings.Index(wildcard, "*")

	// example: */features
	if wildcardIndex == 0 {
		suffix := wildcard[1:]
		return strings.HasSuffix(target, suffix)
	}

	// example: features/*
	if wildcardIndex == len(wildcard)-1 {
		prefix := wildcard[:wildcardIndex]
		return strings.HasPrefix(target, prefix)
	}

	return false
}

// MatchesAnyAsWildcard tests whether the target string matches any of the wildcards elements
func MatchesAnyAsWildcard(target string, wildcards []string) bool {
	for _, wildcard := range wildcards {
		if MatchesWildcard(target, wildcard) {
			return true
		}
	}
	return false
}
