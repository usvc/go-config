package config

import "github.com/spf13/pflag"

type Int struct {
	Shorthand string
	Usage     string
	Default   int
	Value     int
}

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

func (s *Int) GetDefault() interface{} {
	return s.Default
}

func (s *Int) GetShorthand() string {
	return s.Shorthand
}

func (s *Int) GetUsage() string {
	return s.Usage
}

func (s *Int) GetValuePointer() interface{} {
	return &s.Value
}

func (s *Int) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

func (s *Int) SetValue(value interface{}) {
	s.Value = value.(int)
}
