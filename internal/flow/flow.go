package flow

/*
ExitCode is a value used in a call to panic to signify that code execution should be stopped,
before/after listeners executed and finally that the app whould exit with the provided exit code
*/
type ExitCode int

/*
Step is the building block of execution flow.
It has a code block to run, a success step to go to if the former succeeds, or go to an error step otherwise
*/
type Step struct {
	Do      func()
	Success *Step
	Error   *Step
	Desc    string
	Exiter  func(code int)
}

/*
Run call the code block of the step, moves to the success step if the call went ok, opr the the error step otherwise
*/
func (s *Step) Run(p interface{}) { _ = "STUB: not implemented"; return }

func (s *Step) callDo(p interface{}) { _ = "STUB: not implemented"; return }
