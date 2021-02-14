package config

import (
	"github.com/spf13/pflag"
)

// Float stores the configuration details for a floating
// point value
type Float struct {
	Shorthand    string
	Usage        string
	Default      float64
	Value        float64
	controller   *pflag.FlagSet
	internalName string
	isSet        bool
}

// ApplyToFlagSet applies the configuration to a provided flag set
func (s *Float) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     float64
		pointer   *float64
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(float64); !ok {
		value = *new(float64)
	}
	pointer = s.GetValuePointer().(*float64)
	if isZeroValue(shorthand) {
		flags.Float64Var(pointer, name, value, usage)
	} else {
		flags.Float64VarP(pointer, name, shorthand, value, usage)
	}
	s.controller = flags
	s.internalName = name
}

// IsSetExplicitlyByFlag returns true if the value was set by the user even if it equals the default value
func (s Float) IsSetExplicitlyByFlag() bool {
	if s.controller == nil {
		return false
	}
	return s.controller.Changed(s.internalName)
}

// IsSet returns ture if the value was set by the .SetValue method of this instance
func (s Float) IsSet() bool {
	return s.isSet
}

// GetDefault retrieves the default value of this configuration
func (s *Float) GetDefault() interface{} {
	return s.Default
}

// GetShorthand retrieves the short form of the flag if available
func (s *Float) GetShorthand() string {
	return s.Shorthand
}

// GetUsage retrieves the usage string for this configuration
func (s *Float) GetUsage() string {
	return s.Usage
}

// GetValuePointer returns a pointer that points to the instance of the configured value
func (s *Float) GetValuePointer() interface{} {
	return &s.Value
}

// GetValue returns the value of this configuration
func (s *Float) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *Float) SetValue(value interface{}) {
	s.Value = value.(float64)
	s.isSet = true
}
