package matcher

// NewOptsEnd returns the special matcher that matches the -- operator
func NewOptsEnd() Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

const (
	theOptsEnd = optsEnd(true)
)

type optsEnd bool

func (optsEnd) Match(args []string, c *ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func (optsEnd) Priority() int { _ = "STUB: not implemented"; return 0 }

func (optsEnd) String() string { _ = "STUB: not implemented"; return "" }
