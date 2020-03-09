package config

import "github.com/spf13/pflag"

// UintSlice represents a configuration which should be interpreted as
// a slice of unsigned integers
type UintSlice struct {
	Shorthand string
	Usage     string
	Default   []uint
	Value     []uint
}

// ApplyToFlagSet applies the configuration to a provided flag set
func (s *UintSlice) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     []uint
		pointer   *[]uint
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().([]uint); !ok {
		value = *new([]uint)
	}
	pointer = s.GetValuePointer().(*[]uint)
	if isZeroValue(shorthand) {
		flags.UintSliceVar(pointer, name, value, usage)
	} else {
		flags.UintSliceVarP(pointer, name, shorthand, value, usage)
	}
}

// GetDefault retrieves the default value of this configuration
func (s *UintSlice) GetDefault() interface{} {
	return s.Default
}

// GetShorthand retrieves the short form of the flag if available
func (s *UintSlice) GetShorthand() string {
	return s.Shorthand
}

// GetUsage retrieves the usage string for this configuration
func (s *UintSlice) GetUsage() string {
	return s.Usage
}

// GetValuePointer returns a pointer that points to the instance of the configured value
func (s *UintSlice) GetValuePointer() interface{} {
	return &s.Value
}

// GetValue returns the value of this configuration
func (s *UintSlice) GetValue() interface{} {
	return s.Value
}

// SetValue sets the value of this configuration
func (s *UintSlice) SetValue(value interface{}) {
	s.Value = value.([]uint)
}
