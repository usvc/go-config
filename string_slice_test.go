package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type StringSliceTests struct {
	suite.Suite
}

func TestStringSlice(t *testing.T) {
	suite.Run(t, new(StringSliceTests))
}

func (s *StringSliceTests) TestInterface() {
	var conf Config
	b := StringSlice{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *StringSliceTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	stringSlice := &StringSlice{}
	stringSlice.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetStringSlice("no-shorthand")
	s.Nil(err)
	s.Equal([]string{}, val)

	stringSlice = &StringSlice{Default: []string{"a", "b"}}
	stringSlice.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetStringSlice("with-default")
	s.Nil(err)
	s.Equal([]string{"a", "b"}, val)

	stringSlice = &StringSlice{Shorthand: "s"}
	stringSlice.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetStringSlice("with-shorthand")
	s.Nil(err)
	s.Equal([]string{}, val)
}

func (s *StringSliceTests) Test_IsSetExplicitlyByFlag() {
	flags := &pflag.FlagSet{}
	conf := &StringSlice{}
	conf.ApplyToFlagSet("test", flags)
	s.False(conf.IsSetExplicitlyByFlag())
	flags.Set("test", "hello,world")
	s.True(conf.IsSetExplicitlyByFlag())
}

func (s *StringSliceTests) Test_IsSet() {
	conf := &StringSlice{}
	s.False(conf.IsSet())
	s.Nil(conf.SetValue([]string{"hello", "world"}))
	s.True(conf.IsSet())
}

func (s *StringSliceTests) Test_GettersSetters() {
	conf := &StringSlice{
		Default:   []string{"hello", "world"},
		Value:     []string{"hola", "mundo"},
		Shorthand: "t",
		Usage:     "usage",
	}
	s.Equal([]string{"hello", "world"}, conf.GetDefault())
	s.Equal([]string{"hola", "mundo"}, conf.GetValue())
	s.Equal("t", conf.GetShorthand())
	s.Equal("usage", conf.GetUsage())
	valuePointer, ok := conf.GetValuePointer().(*[]string)
	s.True(ok)
	s.Equal(conf.GetValue(), *valuePointer)
	s.Nil(conf.SetValue([]string{"halo", "dunia"}))
	s.Equal([]string{"halo", "dunia"}, conf.GetValue())

	s.Contains(conf.SetValue("").Error(), "interface {} is string, not []string")
}
