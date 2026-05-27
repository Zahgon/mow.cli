package matcher

import (
	"github.com/jawher/mow.cli/internal/container"
)

// NewOptions create an Options matcher which can parse a group of options
func NewOptions(opts []*container.Container, index map[string]*container.Container) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

type options struct {
	options []*container.Container
	index   map[string]*container.Container
}

func (*options) Priority() int { _ = "STUB: not implemented"; return 0 }

func (om *options) Match(args []string, c *ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func (om *options) try(args []string, c *ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func (om *options) String() string { _ = "STUB: not implemented"; return "" }
