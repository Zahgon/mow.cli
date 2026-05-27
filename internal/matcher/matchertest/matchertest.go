package matchertest

import (
	"github.com/jawher/mow.cli/internal/matcher"
)

// NewArg creates a positional argument matcher given its name, e.g. SRC
func NewArg(name string) matcher.Matcher { _ = "STUB: not implemented"; return *new(matcher.Matcher) }

// NewOpt creates a short and long option matcher given its (space separated) names, e.g. -f --force
func NewOpt(name string) matcher.Matcher { _ = "STUB: not implemented"; return *new(matcher.Matcher) }

// NewOptions create an options matcher given their names, e.g. -abc
func NewOptions(names string) matcher.Matcher {
	_ = "STUB: not implemented"
	return *new(matcher.Matcher)
}
