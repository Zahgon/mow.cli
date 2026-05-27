package fsmtest

import (
	"github.com/jawher/mow.cli/internal/fsm"
	"github.com/jawher/mow.cli/internal/matcher"
)

// NopeMatcher is a matcher that always fails
type NopeMatcher struct{}

// Match always returns false without consuming any args
func (NopeMatcher) Match(args []string, c *matcher.ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false,

		// Priority returns the magic value 666
		nil
}

func (NopeMatcher) Priority() int { _ = "STUB: not implemented"; return 0 }

func (NopeMatcher) String() string {
	_ = "STUB: not implemented"

	// YepMatcher is a matcher that always succeeds without consuming any args
	return ""
}

type YepMatcher struct{}

// Match always returns true without consuming any args
func (YepMatcher) Match(args []string, c *matcher.ParseContext) (bool, []string) {
	_ = "STUB: not implemented"

	// Priority returns the magic value 666
	return false, nil
}

func (YepMatcher) Priority() int { _ = "STUB: not implemented"; return 0 }

func (YepMatcher) String() string {
	_ = "STUB: not implemented"

	// TestMatcher is a matcher with a configurable match function and priority
	return ""
}

type TestMatcher struct {
	MatchFunc    func(args []string, c *matcher.ParseContext) (bool, []string)
	TestPriority int
}

// Match executes the provided match func
func (t TestMatcher) Match(args []string, c *matcher.ParseContext) (bool, []string) {
	_ = "STUB: not implemented"
	return false,

		// Priority returns the provided priority
		nil
}

func (t TestMatcher) Priority() int { _ = "STUB: not implemented"; return 0 }

/*
NewFsm constructs an FSM from the provided string spec and a list of defined matchers
The spec syntax is:

S1 t1 S2
S2 t2 (S3)
<source state name> <transition name> <target state name>

states between parenthesis are final states
*/
func NewFsm(spec string, matchers map[string]matcher.Matcher) *fsm.State {
	_ = "STUB: not implemented"
	return nil
}

// TransitionStrs returns a string slice with the transitions names
func TransitionStrs(trs fsm.StateTransitions) []string { _ = "STUB: not implemented"; return nil }

func stateNameTerm(name string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func mkStateNames() *stateNames { _ = "STUB: not implemented"; return nil }

type stateNames struct {
	counter int
	ids     map[*fsm.State]int
}

func (sn *stateNames) id(s *fsm.State) int { _ = "STUB: not implemented"; return 0 }

func stateName(s *fsm.State, sn *stateNames) string { _ = "STUB: not implemented"; return "" }

// FsmStr generates a string representation of the provided FSM
func FsmStr(s *fsm.State) string { _ = "STUB: not implemented"; return "" }

func fsmStrVis(s *fsm.State, sn *stateNames, visited map[*fsm.State]struct{}) fsmStrings {
	_ = "STUB: not implemented"
	return *new(fsmStrings)
}

type fsmStrings []string

func (t fsmStrings) Len() int           { _ = "STUB: not implemented"; return 0 }
func (t fsmStrings) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (t fsmStrings) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func isParen(r rune) bool { _ = "STUB: not implemented"; return false }
