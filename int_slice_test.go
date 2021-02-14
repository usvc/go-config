package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type IntSliceTests struct {
	suite.Suite
}

func TestIntSlice(t *testing.T) {
	suite.Run(t, new(IntSliceTests))
}

func (s *IntSliceTests) TestInterface() {
	var conf Config
	b := IntSlice{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *IntSliceTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	intSlice := &IntSlice{}
	intSlice.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetIntSlice("no-shorthand")
	s.Nil(err)
	s.Equal([]int{}, val)

	intSlice = &IntSlice{Default: []int{-1, -2}}
	intSlice.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetIntSlice("with-default")
	s.Nil(err)
	s.Equal([]int{-1, -2}, val)

	intSlice = &IntSlice{Shorthand: "s"}
	intSlice.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetIntSlice("with-shorthand")
	s.Nil(err)
	s.Equal([]int{}, val)
}

func (s *IntSliceTests) Test_IsSetExplicitlyByFlag() {
	flags := &pflag.FlagSet{}
	conf := &IntSlice{}
	conf.ApplyToFlagSet("test", flags)
	s.False(conf.IsSetExplicitlyByFlag())
	flags.Set("test", "1,2,3,4")
	s.True(conf.IsSetExplicitlyByFlag())
}

func (s *IntSliceTests) Test_IsSet() {
	conf := &IntSlice{}
	s.False(conf.IsSet())
	s.Nil(conf.SetValue([]int{1, 2, 3, 4}))
	s.True(conf.IsSet())
}

func (s *IntSliceTests) Test_GettersSetters() {
	conf := &IntSlice{
		Default:   []int{1, 2, 3, 4},
		Value:     []int{-1, -2, -3, -4},
		Shorthand: "t",
		Usage:     "usage",
	}
	s.Equal([]int{1, 2, 3, 4}, conf.GetDefault())
	s.Equal([]int{-1, -2, -3, -4}, conf.GetValue())
	s.Equal("t", conf.GetShorthand())
	s.Equal("usage", conf.GetUsage())
	valuePointer, ok := conf.GetValuePointer().(*[]int)
	s.True(ok)
	s.Equal(conf.GetValue(), *valuePointer)
	s.Nil(conf.SetValue([]int{-1, 2, -3, 4}))
	s.Equal([]int{-1, 2, -3, 4}, conf.GetValue())

	s.Contains(conf.SetValue("").Error(), "interface {} is string, not []int")
}
