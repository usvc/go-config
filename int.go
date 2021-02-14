package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Int represents a configuration which should be interpreted
// as a signed integer
type Int struct {
	Shorthand    string
	Usage        string
	Default      int
	Value        int
	controller   *pflag.FlagSet
	internalName string
	isSet        bool
}

// ApplyToFlagSet applies the configuration to a provided flag set
func (s *Int) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     int
		pointer   *int
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(int); !ok {
		value = *new(int)
	}
	pointer = s.GetValuePointer().(*int)
	if isZeroValue(shorthand) {
		flags.IntVar(pointer, name, value, usage)
	} else {
		flags.IntVarP(pointer, name, shorthand, value, usage)
	}
	s.controller = flags
	s.internalName = name
}

// IsSetExplicitlyByFlag returns true if the value was set by the user even if it equals the default value
func (s Int) IsSetExplicitlyByFlag() bool {
	if s.controller == nil {
		return false
	}
	return s.controller.Changed(s.internalName)
}

// IsSet returns ture if the value was set by the .SetValue method of this instance
func (s Int) IsSet() bool {
	return s.isSet
}

// GetDefault retrieves the default value of this configuration
func (s *Int) GetDefault() interface{} {
	return s.Default
}

// GetShorthand retrieves the short form of the flag if available
func (s *Int) GetShorthand() string {
	return s.Shorthand
}

// GetUsage retrieves the usage string for this configuration
func (s *Int) GetUsage() string {
	return s.Usage
}

// GetValuePointer returns a pointer that points to the instance of the configured value
func (s *Int) GetValuePointer() interface{} {
	return &s.Value
}

// GetValue returns the value of this configuration
func (s *Int) GetValue() interface{} {
	if !s.IsSet() && !s.IsSetExplicitlyByFlag() && isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *Int) SetValue(value interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	s.Value = value.(int)
	s.isSet = true
	return nil
}
