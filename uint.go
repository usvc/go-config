package config

import "github.com/spf13/pflag"

type Uint struct {
	Shorthand string
	Usage     string
	Default   uint
	Value     uint
	Base
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
	if pointer, ok = s.GetValuePointer().(*uint); ok {
		if isZeroValue(shorthand) {
			flags.UintVar(pointer, name, value, usage)
		} else {
			flags.UintVarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.UintP(name, shorthand, value, usage)
		} else {
			flags.Uint(name, value, usage)
		}
	}
}

func (u *Uint) GetDefault() interface{} {
	return u.Default
}

func (u *Uint) GetValuePointer() interface{} {
	return &u.Value
}

func (u *Uint) GetValue() interface{} {
	return u.Value
}

func (s *Uint) SetValue(value interface{}) {
	s.Value = value.(uint)
}
