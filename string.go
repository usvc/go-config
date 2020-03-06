package config

import (
	"github.com/spf13/pflag"
)

type String struct {
	Shorthand string
	Usage     string
	Default   string
	Value     string
}

func (s *String) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     string
		pointer   *string
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(string); !ok {
		value = *new(string)
	}
	pointer = s.GetValuePointer().(*string)
	if isZeroValue(shorthand) {
		flags.StringVar(pointer, name, value, usage)
	} else {
		flags.StringVarP(pointer, name, shorthand, value, usage)
	}
}

func (s *String) GetDefault() interface{} {
	return s.Default
}

func (s *String) GetShorthand() string {
	return s.Shorthand
}

func (s *String) GetUsage() string {
	return s.Usage
}

func (s *String) GetValuePointer() interface{} {
	return &s.Value
}

func (s *String) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

func (s *String) SetValue(value interface{}) {
	s.Value = value.(string)
}
