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
