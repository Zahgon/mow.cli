package parser

import (
	"github.com/jawher/mow.cli/internal/container"
	"github.com/jawher/mow.cli/internal/fsm"
	"github.com/jawher/mow.cli/internal/lexer"
)

// Params are used to cofigure the parser
type Params struct {
	Spec       string
	Options    []*container.Container
	OptionsIdx map[string]*container.Container
	Args       []*container.Container
	ArgsIdx    map[string]*container.Container
}

// Parse transforms a slice of tokens into an FSM or returns an ParseError
func Parse(tokens []*lexer.Token, params Params) (*fsm.State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type parser struct {
	spec       string
	options    []*container.Container
	optionsIdx map[string]*container.Container
	args       []*container.Container
	argsIdx    map[string]*container.Container

	tokens []*lexer.Token

	tkpos int

	matchedToken *lexer.Token

	rejectOptions bool
}

func (p *parser) parse() (s *fsm.State, err error) { _ = "STUB: not implemented"; return nil, nil }

func (p *parser) seq(required bool) (*fsm.State, *fsm.State) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) choice() (*fsm.State, *fsm.State) { _ = "STUB: not implemented"; return nil, nil }

func (p *parser) atom() (*fsm.State, *fsm.State) { _ = "STUB: not implemented"; return nil, nil }

func (p *parser) canAtom() bool { _ = "STUB: not implemented"; return false }

func (p *parser) found(t lexer.TokenType) bool { _ = "STUB: not implemented"; return false }

func (p *parser) is(t lexer.TokenType) bool { _ = "STUB: not implemented"; return false }

func (p *parser) expect(t lexer.TokenType) { _ = "STUB: not implemented"; return }

func (p *parser) back() { _ = "STUB: not implemented"; return }

func (p *parser) eof() bool { _ = "STUB: not implemented"; return false }

func (p *parser) token() *lexer.Token { _ = "STUB: not implemented"; return nil }
