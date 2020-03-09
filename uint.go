package config

import "github.com/spf13/pflag"

// Uint represents a configuration which should be interpreted
// as an unsigned integer
type Uint struct {
	Shorthand string
	Usage     string
	Default   uint
	Value     uint
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
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *Uint) SetValue(value interface{}) {
	s.Value = value.(uint)
}
