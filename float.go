package config

import "github.com/spf13/pflag"

type Float struct {
	Shorthand string
	Usage     string
	Default   float64
	Value     float64
	Base
}

func (s *Float) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     float64
		pointer   *float64
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(float64); !ok {
		value = *new(float64)
	}
	if pointer, ok = s.GetValuePointer().(*float64); ok {
		if isZeroValue(shorthand) {
			flags.Float64Var(pointer, name, value, usage)
		} else {
			flags.Float64VarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.Float64P(name, shorthand, value, usage)
		} else {
			flags.Float64(name, value, usage)
		}
	}
}

func (f *Float) GetDefault() interface{} {
	return f.Default
}

func (f *Float) GetValuePointer() interface{} {
	return &f.Value
}

func (f *Float) GetValue() interface{} {
	return f.Value
}

func (f *Float) SetValue(value interface{}) {
	f.Value = value.(float64)
}
