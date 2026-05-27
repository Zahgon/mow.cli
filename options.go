package cli

import (
	"flag"

	"github.com/jawher/mow.cli/internal/container"
)

// BoolOpt describes a boolean option
type BoolOpt struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// The option's initial value
	Value bool
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o BoolOpt) value(into *bool) (flag.Value, *bool) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// StringOpt describes a string option
type StringOpt struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// The option's initial value
	Value string
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o StringOpt) value(into *string) (flag.Value, *string) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// IntOpt describes an int option
type IntOpt struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// The option's initial value
	Value int
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o IntOpt) value(into *int) (flag.Value, *int) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// Float64Opt describes an float64 option
type Float64Opt struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// The option's initial value
	Value float64
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o Float64Opt) value(into *float64) (flag.Value, *float64) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// StringsOpt describes a string slice option
type StringsOpt struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The option's initial value
	Value []string
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o StringsOpt) value(into *[]string) (flag.Value, *[]string) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// IntsOpt describes an int slice option
type IntsOpt struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The option's initial value
	Value []int
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o IntsOpt) value(into *[]int) (flag.Value, *[]int) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// Floats64Opt describes an int slice option
type Floats64Opt struct {
	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The option's initial value
	Value []float64
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o Floats64Opt) value(into *[]float64) (flag.Value, *[]float64) {
	_ = "STUB: not implemented"
	return *new(flag.Value), nil
}

// VarOpt describes an option where the type and format of the value is controlled by the developer
type VarOpt struct {
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
	// Set to true if this option was set by the user (as opposed to being set from env or not set at all)
	SetByUser *bool
}

func (o VarOpt) value() flag.Value {
	_ = "STUB: not implemented"

	/*
	   BoolOpt defines a boolean option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

	   The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	   The one letter names will then be called with a single dash (short option), the others with two (long options).

	   The result should be stored in a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
	*/return *new(flag.Value)
}

func (c *Cmd) BoolOpt(name string, value bool, desc string) *bool {
	_ = "STUB: not implemented"
	return nil
}

/*
BoolOptPtr defines a bool option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The into parameter points to a variable (a pointer to a int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) BoolOptPtr(into *bool, name string, value bool, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
StringOpt defines a string option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The result should be stored in a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringOpt(name string, value string, desc string) *string {
	_ = "STUB: not implemented"
	return nil
}

/*
StringOptPtr defines a string option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The into parameter points to a variable (a pointer to a int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringOptPtr(into *string, name string, value string, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
IntOpt defines an int option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The result should be stored in a variable (a pointer to an int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntOpt(name string, value int, desc string) *int {
	_ = "STUB: not implemented"
	return nil
}

/*
IntOptPtr defines a int option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The into parameter points to a variable (a pointer to an int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntOptPtr(into *int, name string, value int, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
Float64Opt defines an float64 option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The result should be stored in a variable (a pointer to an float64) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Float64Opt(name string, value float64, desc string) *float64 {
	_ = "STUB: not implemented"
	return nil
}

/*
Float64OptPtr defines a float64 option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The into parameter points to a variable (a pointer to a float64) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Float64OptPtr(into *float64, name string, value float64, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
StringsOpt defines a string slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The result should be stored in a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsOpt(name string, value []string, desc string) *[]string {
	_ = "STUB: not implemented"
	return nil
}

/*
StringsOptPtr defines a string slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The into parameter points to a variable (a pointer to a int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsOptPtr(into *[]string, name string, value []string, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
IntsOpt defines an int slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The result should be stored in a variable (a pointer to an int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsOpt(name string, value []int, desc string) *[]int {
	_ = "STUB: not implemented"
	return nil
}

/*
IntsOptPtr defines a int slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The into parameter points to a variable (a pointer to a int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsOptPtr(into *[]int, name string, value []int, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
Floats64Opt defines an float64 slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The result should be stored in a variable (a pointer to an float64 slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Floats64Opt(name string, value []float64, desc string) *[]float64 {
	_ = "STUB: not implemented"
	return nil
}

/*
Floats64OptPtr defines a int slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The into parameter points to a variable (a pointer to a int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) Floats64OptPtr(into *[]float64, name string, value []float64, desc string) {
	_ = "STUB: not implemented"
	return
}

/*
VarOpt defines an option where the type and format is controlled by the developer.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).

The result will be stored in the value parameter (a value implementing the flag.Value interface) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) VarOpt(name string, value flag.Value, desc string) { _ = "STUB: not implemented"; return }

func mkOptStrs(optName string) []string { _ = "STUB: not implemented"; return nil }

func (c *Cmd) mkOpt(opt container.Container) { _ = "STUB: not implemented"; return }
