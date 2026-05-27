package fsmdot

import (
	"github.com/jawher/mow.cli/internal/fsm"
)

// Dot generates a graphviz dot representation of an FSM
func Dot(s *fsm.State) string { _ = "STUB: not implemented"; return "" }

func dot(s *fsm.State, sn *stateNames, visited map[*fsm.State]struct{}) []string {
	_ = "STUB: not implemented"
	return nil
}

func mkStateNames() *stateNames { _ = "STUB: not implemented"; return nil }

type stateNames struct {
	counter int
	ids     map[*fsm.State]int
}

func (sn *stateNames) id(s *fsm.State) int { _ = "STUB: not implemented"; return 0 }
