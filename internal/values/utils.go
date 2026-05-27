package values

import (
	"flag"
)

// IsBool checks if a given value is a bool value, i.e. implements the BoolValued interface
func IsBool(v flag.Value) bool { _ = "STUB: not implemented"; return false }

// SetFromEnv fills a value from a list of env vars
func SetFromEnv(into flag.Value, envVars string) bool { _ = "STUB: not implemented"; return false }

func setMultivalued(into MultiValued, values []string) error { _ = "STUB: not implemented"; return nil }

func DefaultValue(v flag.Value) string { _ = "STUB: not implemented"; return "" }
