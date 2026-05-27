package values

import (
	"flag"
)

// BoolValued is an interface values can implement to indicate that they are a bool option, i.e. can be set without providing a value with just -f for example
type BoolValued interface {
	flag.Value
	// IsBoolFlag should return true to indicate that this value is a bool value
	IsBoolFlag() bool
}

// MultiValued is an interface ti indicate that a value can hold multiple values
type MultiValued interface {
	flag.Value
	// Clear should clear the list of values
	Clear()
}

// DefaultValued in an interface to determine if the value stored is the default value, and thus does not need be shown in the help message
type DefaultValued interface {
	// IsDefault should return true if the value stored is the default value, and thus does not need be shown in the help message
	IsDefault() bool
}

/******************************************************************************/
/* BOOL                                                                        */
/******************************************************************************/

// BoolValue is a flag.Value type holding boolean values
type BoolValue bool

var (
	_ flag.Value    = NewBool(new(bool), false)
	_ BoolValued    = NewBool(new(bool), false)
	_ DefaultValued = NewBool(new(bool), false)
)

// NewBool creates a new bool value
func NewBool(into *bool, v bool) *BoolValue { _ = "STUB: not implemented"; return nil }

// Set sets the value from a provided string
func (bo *BoolValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

// IsBoolFlag returns true
func (bo *BoolValue) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }

func (bo *BoolValue) String() string { _ = "STUB: not implemented"; return "" }

// IsDefault return true if the bool value is false
func (bo *BoolValue) IsDefault() bool {
	_ = "STUB: not implemented"

	/******************************************************************************/
	/* STRING                                                                        */
	/******************************************************************************/
	return false
}

// StringValue is a flag.Value type holding string values
type StringValue string

var (
	_ flag.Value    = NewString(new(string), "")
	_ DefaultValued = NewString(new(string), "")
)

// NewString creates a new string value
func NewString(into *string, v string) *StringValue { _ = "STUB: not implemented"; return nil }

// Set sets the value from a provided string
func (sa *StringValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (sa *StringValue) String() string { _ = "STUB: not implemented"; return "" }

// IsDefault return true if the string value is empty
func (sa *StringValue) IsDefault() bool { _ = "STUB: not implemented"; return false }

/******************************************************************************/
/* INT                                                                        */
/******************************************************************************/

// IntValue is a flag.Value type holding int values
type IntValue int

var (
	_ flag.Value = NewInt(new(int), 0)
)

// NewInt creates a new int value
func NewInt(into *int, v int) *IntValue { _ = "STUB: not implemented"; return nil }

// Set sets the value from a provided string
func (ia *IntValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (ia *IntValue) String() string { _ = "STUB: not implemented"; return "" }

/******************************************************************************/
/* Float64                                                                        */
/******************************************************************************/

// Float64Value is a flag.Value type holding int values
type Float64Value float64

var (
	_ flag.Value = NewFloat64(new(float64), 0)
)

// NewFloat64 creates a new int value
func NewFloat64(into *float64, v float64) *Float64Value { _ = "STUB: not implemented"; return nil }

// Set sets the value from a provided string
func (ia *Float64Value) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (ia *Float64Value) String() string { _ = "STUB: not implemented"; return "" }

/******************************************************************************/
/* STRINGS                                                                    */
/******************************************************************************/

// StringsValue is a flag.Value type holding string slices values
type StringsValue []string

var (
	_ flag.Value    = NewStrings(new([]string), nil)
	_ MultiValued   = NewStrings(new([]string), nil)
	_ DefaultValued = NewStrings(new([]string), nil)
)

// NewStrings creates a new multi-string value
func NewStrings(into *[]string, v []string) *StringsValue { _ = "STUB: not implemented"; return nil }

// Set sets the value from a provided string
func (sa *StringsValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (sa *StringsValue) String() string { _ = "STUB: not implemented"; return "" }

// Clear clears the slice
func (sa *StringsValue) Clear() {
	_ = "STUB: not implemented"

	// IsDefault return true if the string slice is empty
	return
}

func (sa *StringsValue) IsDefault() bool { _ = "STUB: not implemented"; return false }

/******************************************************************************/
/* INTS                                                                       */
/******************************************************************************/

// IntsValue is a flag.Value type holding int values
type IntsValue []int

var (
	_ flag.Value    = NewInts(new([]int), nil)
	_ MultiValued   = NewInts(new([]int), nil)
	_ DefaultValued = NewInts(new([]int), nil)
)

// NewInts creates a new multi-int value
func NewInts(into *[]int, v []int) *IntsValue { _ = "STUB: not implemented"; return nil }

// Set sets the value from a provided string
func (ia *IntsValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (ia *IntsValue) String() string { _ = "STUB: not implemented"; return "" }

// Clear clears the slice
func (ia *IntsValue) Clear() {
	_ = "STUB: not implemented"

	// IsDefault return true if the int slice is empty
	return
}

func (ia *IntsValue) IsDefault() bool { _ = "STUB: not implemented"; return false }

/******************************************************************************/
/* FLOATs64                                                                       */
/******************************************************************************/

// Floats64Value is a flag.Value type holding int values
type Floats64Value []float64

var (
	_ flag.Value    = NewFloats64(new([]float64), nil)
	_ MultiValued   = NewFloats64(new([]float64), nil)
	_ DefaultValued = NewFloats64(new([]float64), nil)
)

// NewFloats64 creates a new multi-int value
func NewFloats64(into *[]float64, v []float64) *Floats64Value {
	_ = "STUB: not implemented"
	return nil
}

// Set sets the value from a provided string
func (ia *Floats64Value) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (ia *Floats64Value) String() string { _ = "STUB: not implemented"; return "" }

// Clear clears the slice
func (ia *Floats64Value) Clear() {
	_ = "STUB: not implemented"

	// IsDefault return true if the int slice is empty
	return
}

func (ia *Floats64Value) IsDefault() bool { _ = "STUB: not implemented"; return false }
