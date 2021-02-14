package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type BoolTests struct {
	suite.Suite
}

func TestBool(t *testing.T) {
	suite.Run(t, new(BoolTests))
}

func (s *BoolTests) TestInterface() {
	var conf Config
	b := Bool{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *BoolTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	boolConf := &Bool{}
	boolConf.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetBool("no-shorthand")
	s.Nil(err)
	s.Equal(false, val)

	boolConf = &Bool{Default: true}
	boolConf.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetBool("with-default")
	s.Nil(err)
	s.Equal(true, val)

	boolConf = &Bool{Shorthand: "n"}
	boolConf.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetBool("with-shorthand")
	s.Nil(err)
	s.Equal(false, val)
}

func (s *BoolTests) Test_IsSetExplicitlyByFlag() {
	flags := &pflag.FlagSet{}
	conf := &Bool{}
	conf.ApplyToFlagSet("test", flags)
	s.False(conf.IsSetExplicitlyByFlag())
	flags.Set("test", "true")
	s.True(conf.IsSetExplicitlyByFlag())
}

func (s *BoolTests) Test_IsSet() {
	conf := &Bool{}
	s.False(conf.IsSet())
	s.Nil(conf.SetValue(true))
	s.True(conf.IsSet())
}

func (s *BoolTests) Test_GettersSetters() {
	conf := &Bool{
		Default:   true,
		Value:     false,
		Shorthand: "t",
		Usage:     "usage",
	}
	s.Equal(true, conf.GetDefault())
	s.Equal(true, conf.GetValue())
	s.Equal("t", conf.GetShorthand())
	s.Equal("usage", conf.GetUsage())
	valuePointer, ok := conf.GetValuePointer().(*bool)
	s.True(ok)
	s.NotEqual(conf.GetValue(), *valuePointer,
		"the value of .Value and .GetValue() shoudld be different in this edge case")
	s.Nil(conf.SetValue(false))
	s.Equal(false, conf.GetValue())

	s.Contains(conf.SetValue("").Error(), "interface {} is string, not bool")
}
