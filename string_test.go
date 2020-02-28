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
