package matcher

import (
	"github.com/jawher/mow.cli/internal/container"
)

// NewOpt create an option matcher that can consume short and long options
func NewOpt(o *container.Container, index map[string]*container.Container) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

type opt struct {
	theOne *container.Container
	index  map[string]*container.Container
}

func (*opt) Priority() int { _ = "STUB: not implemented"; return 0 }

func (o *opt) String() string { _ = "STUB: not implemented"; return "" }

func (o *opt) Match(args []string, c *ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func (o *opt) matchLongOpt(args []string, idx int, c *ParseContext) (bool, int, []string) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (o *opt) matchShortOpt(args []string, idx int, c *ParseContext) (bool, int, []string) {
	_ = "STUB: not implemented"
	return false, 0, nil
}
