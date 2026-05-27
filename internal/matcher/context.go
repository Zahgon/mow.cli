package matcher

import "github.com/jawher/mow.cli/internal/container"

// ParseContext holds the state of the arguments parsing, i.e. the encountered options and arguments values, etc.
type ParseContext struct {
	Args          map[*container.Container][]string
	Opts          map[*container.Container][]string
	ExcludedOpts  map[*container.Container]struct{}
	RejectOptions bool
}

// NewParseContext create a new ParseContext
func NewParseContext() ParseContext { _ = "STUB: not implemented"; return *new(ParseContext) }

// Merge adds the values in the provided context in the current context
func (pc ParseContext) Merge(o ParseContext) { _ = "STUB: not implemented"; return }
