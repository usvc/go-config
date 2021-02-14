package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

// StringSlice represents a configuration which should be represented
// as a slice of strings
type StringSlice struct {
	Shorthand    string
	Usage        string
	Default      []string
	Value        []string
	controller   *pflag.FlagSet
	internalName string
	isSet        bool
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
	s.controller = flags
	s.internalName = name
}

// IsSetExplicitlyByFlag returns true if the value was set by the user even if it equals the default value
func (s StringSlice) IsSetExplicitlyByFlag() bool {
	if s.controller == nil {
		return false
	}
	return s.controller.Changed(s.internalName)
}

// IsSet returns ture if the value was set by the .SetValue method of this instance
func (s StringSlice) IsSet() bool {
	return s.isSet
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
	if !s.IsSet() && !s.IsSetExplicitlyByFlag() && isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *StringSlice) SetValue(value interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	s.Value = value.([]string)
	s.isSet = true
	return nil
}
