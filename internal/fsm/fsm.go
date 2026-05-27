package fsm

import (
	"github.com/jawher/mow.cli/internal/container"
	"github.com/jawher/mow.cli/internal/matcher"
)

/*
State is the basic building block in the FSM.
A State can be final or not, and has transitions to other states
*/
type State struct {
	Terminal    bool
	Transitions StateTransitions
}

/*
Transition links 2 states.
If a transition's matcher matches, the next state can be reached
*/
type Transition struct {
	Matcher matcher.Matcher
	Next    *State
}

// StateTransitions is a sortable slice of transitions according to their priorities
type StateTransitions []*Transition

func (t StateTransitions) Len() int           { _ = "STUB: not implemented"; return 0 }
func (t StateTransitions) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (t StateTransitions) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// NewState create a new state
func NewState() *State { _ = "STUB: not implemented"; return nil }

// T creates a transition between 2 states
func (s *State) T(matcher matcher.Matcher, next *State) *State {
	_ = "STUB: not implemented"
	return nil
}

// Prepare simplifies the FSM and sorts the transitions according to their priorities
func (s *State) Prepare() { _ = "STUB: not implemented"; return }

func sortTransitions(s *State, visited map[*State]bool) { _ = "STUB: not implemented"; return }

func simplify(start, s *State, visited map[*State]bool) { _ = "STUB: not implemented"; return }

func (s *State) simplifySelf(start *State) bool { _ = "STUB: not implemented"; return false }

func removeTransitionAt(idx int, arr StateTransitions) StateTransitions {
	_ = "STUB: not implemented"
	return *new(StateTransitions)
}

func (s *State) has(tr *Transition) bool { _ = "STUB: not implemented"; return false }

// Parse tries to navigate into the FSM according to the provided args
func (s *State) Parse(args []string) error { _ = "STUB: not implemented"; return nil }

func fillContainers(containers map[*container.Container][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) apply(args []string, pc matcher.ParseContext) bool {
	_ = "STUB: not implemented"
	return false
}
