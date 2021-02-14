package config

import (
	"github.com/spf13/pflag"
)

// Bool represents a configuration which should be interpreted
// as a boolean-typed value
type Bool struct {
	Shorthand    string
	Usage        string
	Default      bool
	Value        bool
	controller   *pflag.FlagSet
	internalName string
	isSet        bool
}

// ApplyToFlagSet applies the configuration to a provided flag set
func (s *Bool) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		value     = s.GetDefault().(bool)
		pointer   = s.GetValuePointer().(*bool)
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if isZeroValue(shorthand) {
		flags.BoolVar(pointer, name, value, usage)
	} else {
		flags.BoolVarP(pointer, name, shorthand, value, usage)
	}
	s.controller = flags
	s.internalName = name
}

// IsSetExplicitlyByFlag returns true if the value was set by the user even if it equals the default value
func (s Bool) IsSetExplicitlyByFlag() bool {
	if s.controller == nil {
		return false
	}
	return s.controller.Changed(s.internalName)
}

// IsSet returns ture if the value was set by the .SetValue method of this instance
func (s Bool) IsSet() bool {
	return s.isSet
}

// GetDefault retrieves the default value of this configuration
func (s *Bool) GetDefault() interface{} {
	return s.Default
}

// GetShorthand retrieves the short form of the flag if available
func (s *Bool) GetShorthand() string {
	return s.Shorthand
}

// GetUsage retrieves the usage string for this configuration
func (s *Bool) GetUsage() string {
	return s.Usage
}

// GetValuePointer returns a pointer that points to the instance of the configured value
func (s *Bool) GetValuePointer() interface{} {
	return &s.Value
}

// GetValue returns the value of this configuration
func (s *Bool) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *Bool) SetValue(value interface{}) {
	s.Value = value.(bool)
	s.isSet = true
}
