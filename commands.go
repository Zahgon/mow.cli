package cli

import (
	"flag"
	"io"

	"github.com/jawher/mow.cli/internal/container"
	"github.com/jawher/mow.cli/internal/flow"
	"github.com/jawher/mow.cli/internal/fsm"
)

/*
Cmd represents a command (or sub command) in a CLI application. It should be constructed
by calling Command() on an app to create a top level command or by calling Command() on another
command to create a sub command
*/
type Cmd struct {
	// The code to execute when this command is matched
	Action func()
	// The code to execute before this command or any of its children is matched
	Before func()
	// The code to execute after this command or any of its children is matched
	After func()
	// The command options and arguments
	Spec string
	// The command long description to be shown when help is requested
	LongDesc string
	// Hide this command in the help messages
	Hidden bool
	// The command error handling strategy
	ErrorHandling flag.ErrorHandling

	init    CmdInitializer
	name    string
	aliases []string
	desc    string

	commands   []*Cmd
	options    []*container.Container
	optionsIdx map[string]*container.Container
	args       []*container.Container
	argsIdx    map[string]*container.Container

	parents []string

	fsm *fsm.State
}

/*
BoolParam represents a Bool option or argument
*/
type BoolParam interface {
	value(into *bool) (flag.Value, *bool)
}

/*
StringParam represents a String option or argument
*/
type StringParam interface {
	value(into *string) (flag.Value, *string)
}

/*
IntParam represents an Int option or argument
*/
type IntParam interface {
	value(into *int) (flag.Value, *int)
}

/*
Float64Param represents an Float64 option or argument
*/
type Float64Param interface {
	value(into *float64) (flag.Value, *float64)
}

/*
StringsParam represents a string slice option or argument
*/
type StringsParam interface {
	value(into *[]string) (flag.Value, *[]string)
}

/*
IntsParam represents an float64 slice option or argument
*/
type IntsParam interface {
	value(into *[]int) (flag.Value, *[]int)
}

/*
Floats64Param represents an float64 slice option or argument
*/
type Floats64Param interface {
	value(into *[]float64) (flag.Value, *[]float64)
}

/*
VarParam represents an custom option or argument where the type and format are controlled by the developer
*/
type VarParam interface {
	value() flag.Value
}

/*
CmdInitializer is a function that configures a command by adding options, arguments, a spec, sub commands and the code
to execute when the command is called
*/
type CmdInitializer func(*Cmd)

/*
Command adds a new (sub) command to c where name is the command name (what you type in the console),
description is what would be shown in the help messages, e.g.:

	Usage: git [OPTIONS] COMMAND [arg...]

	Commands:
	  $name	$desc

the last argument, init, is a function that will be called by mow.cli to further configure the created
(sub) command, e.g. to add options, arguments and the code to execute
*/
func (c *Cmd) Command(name, desc string, init CmdInitializer) { _ = "STUB: not implemented"; return }

/*
Bool can be used to add a bool option or argument to a command.
It accepts either a BoolOpt or a BoolArg struct.

The result should be stored in a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Bool(p BoolParam) *bool { _ = "STUB: not implemented"; return nil }

/*
BoolPtr can be used to add a bool option or argument to a command.
It accepts either a pointer to a bool var and a BoolOpt or a BoolArg struct.

The into parameter points to a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) BoolPtr(into *bool, p BoolParam) { _ = "STUB: not implemented"; return }

/*
String can be used to add a string option or argument to a command.
It accepts either a StringOpt or a StringArg struct.

The result should be stored in a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) String(p StringParam) *string { _ = "STUB: not implemented"; return nil }

/*
StringPtr can be used to add a string option or argument to a command.
It accepts either a pointer to a string var and a StringOpt or a StringArg struct.

The into parameter points to a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringPtr(into *string, p StringParam) { _ = "STUB: not implemented"; return }

/*
Int can be used to add an int option or argument to a command.
It accepts either a IntOpt or a IntArg struct.

The result should be stored in a variable (a pointer to an int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Int(p IntParam) *int { _ = "STUB: not implemented"; return nil }

/*
IntPtr can be used to add a int option or argument to a command.
It accepts either a pointer to a int var and a IntOpt or a IntArg struct.

The into parameter points to a variable (a pointer to a int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntPtr(into *int, p IntParam) { _ = "STUB: not implemented"; return }

/*
Float64 can be used to add a float64 option or argument to a command.
It accepts either a Float64Opt or a Float64Arg struct.

The result should be stored in a variable (a pointer to a float64) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Float64(p Float64Param) *float64 { _ = "STUB: not implemented"; return nil }

/*
Float64Ptr can be used to add a float64 option or argument to a command.
It accepts either a pointer to a float64 var and a Float64Opt or a Float64Arg struct.

The into parameter points to a variable (a pointer to a float64) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Float64Ptr(into *float64, p Float64Param) { _ = "STUB: not implemented"; return }

/*
Strings can be used to add a string slice option or argument to a command.
It accepts either a StringsOpt or a StringsArg struct.

The result should be stored in a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Strings(p StringsParam) *[]string { _ = "STUB: not implemented"; return nil }

/*
StringsPtr can be used to add a string slice option or argument to a command.
It accepts either a pointer to a string slice var and a StringsOpt or a StringsArg struct.

The into parameter points to a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsPtr(into *[]string, p StringsParam) { _ = "STUB: not implemented"; return }

/*
Ints can be used to add an int slice option or argument to a command.
It accepts either a IntsOpt or a IntsArg struct.

The result should be stored in a variable (a pointer to an int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Ints(p IntsParam) *[]int { _ = "STUB: not implemented"; return nil }

/*
IntsPtr can be used to add a int slice option or argument to a command.
It accepts either a pointer to a int slice var and a IntsOpt or a IntsArg struct.

The into parameter points to a variable (a pointer to a int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsPtr(into *[]int, p IntsParam) { _ = "STUB: not implemented"; return }

/*
Floats64 can be used to add an float64 slice option or argument to a command.
It accepts either a Floats64Opt or a Floats64Arg struct.

The result should be stored in a variable (a pointer to an float64 slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Floats64(p Floats64Param) *[]float64 { _ = "STUB: not implemented"; return nil }

/*
Floats64Ptr can be used to add a float64 slice option or argument to a command.
It accepts either a pointer to a float64 slice var and a Floats64Opt or a Floats64Arg struct.

The into parameter points to a variable (a pointer to a float64 slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Floats64Ptr(into *[]float64, p Floats64Param) { _ = "STUB: not implemented"; return }

/*
Var can be used to add a custom option or argument to a command.
It accepts either a VarOpt or a VarArg struct.

As opposed to the other built-in types, this function does not return a pointer the the value.
Instead, the VarOpt or VarOptArg structs hold the said value.
*/
func (c *Cmd) Var(p VarParam) { _ = "STUB: not implemented"; return }

func (c *Cmd) doInit() error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) onError(err error) { _ = "STUB: not implemented"; return }

/*
PrintHelp prints the command's help message.
In most cases the library users won't need to call this method, unless
a more complex validation is needed
*/
func (c *Cmd) PrintHelp() {
	_ = "STUB: not implemented"

	/*
	   PrintLongHelp prints the command's help message using the command long description if specified.
	   In most cases the library users won't need to call this method, unless
	   a more complex validation is needed
	*/return
}

func (c *Cmd) PrintLongHelp() { _ = "STUB: not implemented"; return }

func (c *Cmd) printHelp(longDesc bool) { _ = "STUB: not implemented"; return }

func formatOptNamesForHelp(o *container.Container) string { _ = "STUB: not implemented"; return "" }

// 2 spaces instead of the short option (-x), one space for the comma (,) and one space for the after comma blank

func formatValueForHelp(hide bool, v string) string { _ = "STUB: not implemented"; return "" }

func formatEnvVarsForHelp(envVars string) string { _ = "STUB: not implemented"; return "" }

func (c *Cmd) parse(args []string, entry, inFlow, outFlow *flow.Step) error {
	_ = "STUB: not implemented"
	return nil
}

// help was requested, but not for this command, skip the validation

// impossible case

func (c *Cmd) helpIndex(args []string) int { _ = "STUB: not implemented"; return 0 }

func (c *Cmd) isFirstItemAmong(args []string, searchSet []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Cmd) getOptsAndArgs(args []string) int { _ = "STUB: not implemented"; return 0 }

func (c *Cmd) isAlias(arg string) bool { _ = "STUB: not implemented"; return false }

func joinStrings(parts ...string) string { _ = "STUB: not implemented"; return "" }

func printTabbedRow(w io.Writer, s1 string, s2 string) { _ = "STUB: not implemented"; return }
