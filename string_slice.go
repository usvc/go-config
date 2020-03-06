package config

import (
	"github.com/spf13/pflag"
)

type StringSlice struct {
	Shorthand string
	Usage     string
	Default   []string
	Value     []string
}

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

func (s *StringSlice) GetDefault() interface{} {
	return s.Default
}

func (s *StringSlice) GetShorthand() string {
	return s.Shorthand
}

func (s *StringSlice) GetUsage() string {
	return s.Usage
}

func (s *StringSlice) GetValuePointer() interface{} {
	return &s.Value
}

func (s *StringSlice) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

func (s *StringSlice) SetValue(value interface{}) {
	s.Value = value.([]string)
}
