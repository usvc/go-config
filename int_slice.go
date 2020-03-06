package config

import "github.com/spf13/pflag"

type IntSlice struct {
	Shorthand string
	Usage     string
	Default   []int
	Value     []int
}

func (s *IntSlice) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     []int
		pointer   *[]int
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().([]int); !ok {
		value = *new([]int)
	}
	pointer = s.GetValuePointer().(*[]int)
	if isZeroValue(shorthand) {
		flags.IntSliceVar(pointer, name, value, usage)
	} else {
		flags.IntSliceVarP(pointer, name, shorthand, value, usage)
	}
}

func (s *IntSlice) GetDefault() interface{} {
	return s.Default
}

func (s *IntSlice) GetShorthand() string {
	return s.Shorthand
}

func (s *IntSlice) GetUsage() string {
	return s.Usage
}

func (s *IntSlice) GetValuePointer() interface{} {
	return &s.Value
}

func (s *IntSlice) GetValue() interface{} {
	if isZeroValue(s.Value) {
		return s.Default
	}
	return s.Value
}

func (s *IntSlice) SetValue(value interface{}) {
	s.Value = value.([]int)
}
