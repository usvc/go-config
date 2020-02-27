package config

import (
	"github.com/spf13/pflag"
)

type String struct {
	Default string
	Value   string
	Base
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
	if pointer, ok = s.GetValuePointer().(*string); ok {
		if isZeroValue(shorthand) {
			flags.StringVar(pointer, name, value, usage)
		} else {
			flags.StringVarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.StringP(name, shorthand, value, usage)
		} else {
			flags.String(name, value, usage)
		}
	}
}

func (s *String) GetValuePointer() interface{} {
	return &s.Value
}

func (s *String) GetDefault() interface{} {
	return s.Default
}

func (s *String) GetValue() interface{} {
	return s.Value
}

func (s *String) SetValue(value interface{}) {
	s.Value = value.(string)
}
