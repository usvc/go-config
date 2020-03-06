package config

import "github.com/spf13/pflag"

type Float struct {
	Shorthand string
	Usage     string
	Default   float64
	Value     float64
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
	pointer = s.GetValuePointer().(*float64)
	if isZeroValue(shorthand) {
		flags.Float64Var(pointer, name, value, usage)
	} else {
		flags.Float64VarP(pointer, name, shorthand, value, usage)
	}
}

func (s *Float) GetDefault() interface{} {
	return s.Default
}

func (s *Float) GetShorthand() string {
	return s.Shorthand
}

func (s *Float) GetUsage() string {
	return s.Usage
}

func (s *Float) GetValuePointer() interface{} {
	return &s.Value
}

func (s *Float) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

func (s *Float) SetValue(value interface{}) {
	s.Value = value.(float64)
}
