package config

import (
	"github.com/spf13/pflag"
)

// StringSlice represents a configuration which should be represented
// as a slice of strings
type StringSlice struct {
	Shorthand string
	Usage     string
	Default   []string
	Value     []string
}

// // ApplyToFlagSet applies the configuration to a provided flag set
func (s *StringSlice) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     []string
		pointer   *[]string
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().([]string); !ok {
		value = *new([]string)
	}
	pointer = s.GetValuePointer().(*[]string)
	if isZeroValue(shorthand) {
		flags.StringSliceVar(pointer, name, value, usage)
	} else {
		flags.StringSliceVarP(pointer, name, shorthand, value, usage)
	}
}

// GetDefault retrieves the default value of this configuration
func (s *StringSlice) GetDefault() interface{} {
	return s.Default
}

// GetShorthand retrieves the short form of the flag if available
func (s *StringSlice) GetShorthand() string {
	return s.Shorthand
}

// GetUsage retrieves the usage string for this configuration
func (s *StringSlice) GetUsage() string {
	return s.Usage
}

// GetValuePointer returns a pointer that points to the instance of the configured value
func (s *StringSlice) GetValuePointer() interface{} {
	return &s.Value
}

// GetValue returns the value of this configuration
func (s *StringSlice) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *StringSlice) SetValue(value interface{}) {
	s.Value = value.([]string)
}
