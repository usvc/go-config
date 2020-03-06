package config

import "github.com/spf13/pflag"

type Uint struct {
	Shorthand string
	Usage     string
	Default   uint
	Value     uint
}

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

func (s *Uint) GetDefault() interface{} {
	return s.Default
}

func (s *Uint) GetShorthand() string {
	return s.Shorthand
}

func (s *Uint) GetUsage() string {
	return s.Usage
}

func (s *Uint) GetValuePointer() interface{} {
	return &s.Value
}

func (s *Uint) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

func (s *Uint) SetValue(value interface{}) {
	s.Value = value.(uint)
}
