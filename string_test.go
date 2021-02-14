package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type StringTests struct {
	suite.Suite
}

func TestString(t *testing.T) {
	suite.Run(t, new(StringTests))
}

func (s *StringTests) TestInterface() {
	var conf Config
	b := String{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *StringTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	stringConf := &String{}
	stringConf.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetString("no-shorthand")
	s.Nil(err)
	s.Equal("", val)

	stringConf = &String{Default: "default"}
	stringConf.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetString("with-default")
	s.Nil(err)
	s.Equal("default", val)

	stringConf = &String{Shorthand: "s"}
	stringConf.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetString("with-shorthand")
	s.Nil(err)
	s.Equal("", val)
}

func (s *StringTests) Test_IsSetExplicitlyByFlag() {
	flags := &pflag.FlagSet{}
	conf := &String{}
	conf.ApplyToFlagSet("test", flags)
	s.False(conf.IsSetExplicitlyByFlag())
	flags.Set("test", "test")
	s.True(conf.IsSetExplicitlyByFlag())
}

func (s *StringTests) Test_IsSet() {
	conf := &String{}
	s.False(conf.IsSet())
	s.Nil(conf.SetValue("test"))
	s.True(conf.IsSet())
}

func (s *StringTests) Test_GettersSetters() {
	conf := &String{
		Default:   "default",
		Value:     "value",
		Shorthand: "t",
		Usage:     "usage",
	}
	s.Equal("default", conf.GetDefault())
	s.Equal("value", conf.GetValue())
	s.Equal("t", conf.GetShorthand())
	s.Equal("usage", conf.GetUsage())
	valuePointer, ok := conf.GetValuePointer().(*string)
	s.True(ok)
	s.Equal(conf.GetValue(), *valuePointer)
	s.Nil(conf.SetValue("new value"))
	s.Equal("new value", conf.GetValue())
}
