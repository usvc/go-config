package config

import "github.com/spf13/pflag"

type Bool struct {
	Default bool
	Value   bool
	Base
}

func (s *Bool) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     bool
		pointer   *bool
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(bool); !ok {
		value = *new(bool)
	}
	if pointer, ok = s.GetValuePointer().(*bool); ok {
		if isZeroValue(shorthand) {
			flags.BoolVar(pointer, name, value, usage)
		} else {
			flags.BoolVarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.BoolP(name, shorthand, value, usage)
		} else {
			flags.Bool(name, value, usage)
		}
	}
}

func (s *Bool) GetValuePointer() interface{} {
	return &s.Value
}

func (s *Bool) GetDefault() interface{} {
	return s.Default
}

func (s *Bool) GetValue() interface{} {
	return s.Value
}

func (s *Bool) SetValue(value interface{}) {
	s.Value = value.(bool)
}
