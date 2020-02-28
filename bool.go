package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

type Bool struct {
	Shorthand string
	Usage     string
	Default   bool
	Value     bool
}

func (s *Bool) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		value     = s.GetDefault().(bool)
		pointer   = s.GetValuePointer().(*bool)
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if isZeroValue(shorthand) {
		fmt.Print("reached 2")
		flags.BoolVar(pointer, name, value, usage)
	} else {
		fmt.Print("reached 3")
		flags.BoolVarP(pointer, name, shorthand, value, usage)
	}
}

func (s *Bool) GetValuePointer() interface{} {
	return &s.Value
}

func (s *Bool) GetDefault() interface{} {
	return s.Default
}

func (s *Bool) GetShorthand() string {
	return s.Shorthand
}

func (s *Bool) GetUsage() string {
	return s.Usage
}

func (s *Bool) GetValue() interface{} {
	return s.Value
}

func (s *Bool) SetValue(value interface{}) {
	s.Value = value.(bool)
}
