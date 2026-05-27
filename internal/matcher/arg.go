package matcher

import (
	"github.com/jawher/mow.cli/internal/container"
)

// NewArg creates an (positional) argument matcher
func NewArg(a *container.Container) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

type arg struct {
	arg *container.Container
}

func (arg *arg) Match(args []string, c *ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func (*arg) Priority() int { _ = "STUB: not implemented"; return 0 }

func (arg *arg) String() string { _ = "STUB: not implemented"; return "" }
