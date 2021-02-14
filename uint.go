package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Uint represents a configuration which should be interpreted
// as an unsigned integer
type Uint struct {
	Shorthand    string
	Usage        string
	Default      uint
	Value        uint
	controller   *pflag.FlagSet
	internalName string
	isSet        bool
}

// ApplyToFlagSet applies the configuration to a provided flag set
func (s *Uint) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     uint
		pointer   *uint
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(uint); !ok {
		value = *new(uint)
	}
	pointer = s.GetValuePointer().(*uint)
	if isZeroValue(shorthand) {
		flags.UintVar(pointer, name, value, usage)
	} else {
		flags.UintVarP(pointer, name, shorthand, value, usage)
	}
	s.controller = flags
	s.internalName = name
}

// IsSetExplicitlyByFlag returns true if the value was set by the user even if it equals the default value
func (s Uint) IsSetExplicitlyByFlag() bool {
	if s.controller == nil {
		return false
	}
	return s.controller.Changed(s.internalName)
}

// IsSet returns ture if the value was set by the .SetValue method of this instance
func (s Uint) IsSet() bool {
	return s.isSet
}

// GetDefault retrieves the default value of this configuration
func (s *Uint) GetDefault() interface{} {
	return s.Default
}

// GetShorthand retrieves the short form of the flag if available
func (s *Uint) GetShorthand() string {
	return s.Shorthand
}

// GetUsage retrieves the usage string for this configuration
func (s *Uint) GetUsage() string {
	return s.Usage
}

// GetValuePointer returns a pointer that points to the instance of the configured value
func (s *Uint) GetValuePointer() interface{} {
	return &s.Value
}

// GetValue returns the value of this configuration
func (s *Uint) GetValue() interface{} {
	if !s.IsSet() && !s.IsSetExplicitlyByFlag() && isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *Uint) SetValue(value interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	s.Value = value.(uint)
	s.isSet = true
	return nil
}
