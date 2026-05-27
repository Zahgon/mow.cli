package matcher

type shortcut bool

const (
	theShortcut = shortcut(true)
)

// NewShortcut create a special matcher that always matches and doesn't consume any input
func NewShortcut() Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

func (shortcut) Match(args []string, c *ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func (shortcut) Priority() int { _ = "STUB: not implemented"; return 0 }

func (shortcut) String() string { _ = "STUB: not implemented"; return "" }
