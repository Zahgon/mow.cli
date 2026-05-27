package lexer

// TokenType is a type representing the different kinds of tokens
type TokenType string

const (
	// TTArg is an arg token, e.g. ARG, SRC. DST
	TTArg TokenType = "Arg"
	// TTOpenPar (
	TTOpenPar TokenType = "OpenPar"
	// TTClosePar )
	TTClosePar TokenType = "ClosePar"
	// TTOpenSq [
	TTOpenSq TokenType = "OpenSq"
	// TTCloseSq ]
	TTCloseSq TokenType = "CloseSq"
	// TTChoice |
	TTChoice TokenType = "Choice"
	// TTOptions is the special OPTIONS keyword
	TTOptions TokenType = "Options"
	// TTRep ...
	TTRep TokenType = "Rep"
	// TTShortOpt -a, -f, ...
	TTShortOpt TokenType = "ShortOpt"
	// TTLongOpt --force, --retry, ...
	TTLongOpt TokenType = "LongOpt"
	// TTOptSeq a folded option sequence, -rm
	TTOptSeq TokenType = "OptSeq"
	// TTOptValue is the special =<example> syntax token
	TTOptValue TokenType = "OptValue"
	// TTDoubleDash --
	TTDoubleDash TokenType = "DblDash"
)

// Token has a type, a value and a position in the input
type Token struct {
	// Type is the token type
	Typ TokenType
	// Val the textual content
	Val string
	// Pos is the token position in the input
	Pos int
}

func (t *Token) String() string { _ = "STUB: not implemented"; return "" }

// ParseError represents a parsing error
type ParseError struct {
	// Input is the text to parse
	Input string
	// Msg s the error message
	Msg string
	// Post is where in the input the error occurred
	Pos int
}

func (t *ParseError) ident() string { _ = "STUB: not implemented"; return "" }

func (t *ParseError) Error() string { _ = "STUB: not implemented"; return "" }

// Tokenize transforms the provided input into a slice of tokens or returns a ParseError
func Tokenize(usage string) ([]*Token, error) { _ = "STUB: not implemented"; return nil, nil }

func isLowercase(c uint8) bool { _ = "STUB: not implemented"; return false }

func isUppercase(c uint8) bool { _ = "STUB: not implemented"; return false }

func isOkInArg(c uint8) bool { _ = "STUB: not implemented"; return false }

func isLetter(c uint8) bool { _ = "STUB: not implemented"; return false }

func isDigit(c uint8) bool { _ = "STUB: not implemented"; return false }

func isOkLongOpt(c uint8, first bool) bool { _ = "STUB: not implemented"; return false }
