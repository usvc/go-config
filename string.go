package config

import (
	"github.com/spf13/pflag"
)

// String represents a configuration which should be interpreted
// as a string-typed value
type String struct {
	Shorthand string
	Usage     string
	Default   string
	Value     string
}

// ApplyToFlagSet applies the configuration to a provided flag set
func (s *String) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     string
		pointer   *string
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(string); !ok {
		value = *new(string)
	}
	pointer = s.GetValuePointer().(*string)
	if isZeroValue(shorthand) {
		flags.StringVar(pointer, name, value, usage)
	} else {
		flags.StringVarP(pointer, name, shorthand, value, usage)
	}
}

// GetDefault retrieves the default value of this configuration
func (s *String) GetDefault() interface{} {
	return s.Default
}

// GetShorthand retrieves the short form of the flag if available
func (s *String) GetShorthand() string {
	return s.Shorthand
}

// GetUsage retrieves the usage string for this configuration
func (s *String) GetUsage() string {
	return s.Usage
}

// GetValuePointer returns a pointer that points to the instance of the configured value
func (s *String) GetValuePointer() interface{} {
	return &s.Value
}

// GetValue returns the value of this configuration
func (s *String) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *String) SetValue(value interface{}) {
	s.Value = value.(string)
}
