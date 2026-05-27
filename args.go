package cli

import (
	"flag"

	"github.com/jawher/mow.cli/internal/container"
)

// BoolArg describes a boolean argument
type BoolArg struct {
	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's initial value
	Value bool
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a BoolArg) value(into *bool) (flag.Value, *bool) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// StringArg describes a string argument
type StringArg struct {
	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's initial value
	Value string
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a StringArg) value(into *string) (flag.Value, *string) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// IntArg describes an int argument
type IntArg struct {
	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's initial value
	Value int
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a IntArg) value(into *int) (flag.Value, *int) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// Float64Arg describes an float64 argument
type Float64Arg struct {
	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's initial value
	Value float64
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a Float64Arg) value(into *float64) (flag.Value, *float64) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// StringsArg describes a string slice argument
type StringsArg struct {
	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The argument's initial value
	Value []string
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a StringsArg) value(into *[]string) (flag.Value, *[]string) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// IntsArg describes an int slice argument
type IntsArg struct {
	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The argument's initial value
	Value []int
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a IntsArg) value(into *[]int) (flag.Value, *[]int) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// Floats64Arg describes an int slice argument
type Floats64Arg struct {
	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The argument's initial value
	Value []float64
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a Floats64Arg) value(into *[]float64) (flag.Value, *[]float64) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// VarArg describes an argument where the type and format of the value is controlled by the developer
type VarArg struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// A value implementing the flag.Value type (will hold the final value)
	Value flag.Value
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this arg was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (a VarArg) value() flag.Value {
	_ = "STUB: not implemented"

	/*
	   BoolArg defines a boolean argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

	   The result should be stored in a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
	*/return *new(flag.Value)
}

func (c *Cmd) BoolArg(name string, value bool, desc string) *bool {
	_ = "STUB: not implemented"
	return nil
}

/*
BoolArgPtr defines a boolean argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The into parameter points to a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) BoolArgPtr(into *bool, name string, value bool, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
StringArg defines a string argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringArg(name string, value string, desc string) *string {
	_ = "STUB: not implemented"
	return nil
}

/*
StringArgPtr defines a string argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The into parameter points to a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringArgPtr(into *string, name string, value string, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
IntArg defines an int argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntArg(name string, value int, desc string) *int {
	_ = "STUB: not implemented"
	return nil
}

/*
IntArgPtr defines an int argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The into parameter points to a variable (a pointer to a int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntArgPtr(into *int, name string, value int, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
Float64Arg defines an float64 argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an float64) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Float64Arg(name string, value float64, desc string) *float64 {
	_ = "STUB: not implemented"
	return nil
}

/*
Float64ArgPtr defines an float64 argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The into parameter points to a variable (a pointer to a float64) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Float64ArgPtr(into *float64, name string, value float64, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
StringsArg defines a string slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsArg(name string, value []string, desc string) *[]string {
	_ = "STUB: not implemented"
	return nil
}

/*
StringsArgPtr defines a string slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The into parameter points to a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsArgPtr(into *[]string, name string, value []string, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
IntsArg defines an int slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsArg(name string, value []int, desc string) *[]int {
	_ = "STUB: not implemented"
	return nil
}

/*
IntsArgPtr defines a int slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The into parameter points to a variable (a pointer to a int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsArgPtr(into *[]int, name string, value []int, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
Floats64Arg defines an float64 slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an float64 slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Floats64Arg(name string, value []float64, desc string) *[]float64 {
	_ = "STUB: not implemented"
	return nil
}

/*
Floats64ArgPtr defines a float64 slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The into parameter points to a variable (a pointer to a float64 slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Floats64ArgPtr(into *[]float64, name string, value []float64, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
VarArg defines an argument where the type and format is controlled by the developer on the command c named `name` and a description of `desc` which will be used in help messages.

The result will be stored in the value parameter (a value implementing the flag.Value interface) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) VarArg(name string, value flag.Value, desc string) { _ = "STUB: not implemented"; return }

func (c *Cmd) mkArg(arg container.Container) { _ = "STUB: not implemented"; return }

func validArgName(n string) bool { _ = "STUB: not implemented"; return false }
