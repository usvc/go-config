package config

import "github.com/spf13/pflag"

// Int represents a configuration which should be interpreted
// as a signed integer
type Int struct {
	Shorthand string
	Usage     string
	Default   int
	Value     int
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
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

// SetValue sets the value of this configuration
func (s *Int) SetValue(value interface{}) {
	s.Value = value.(int)
}
