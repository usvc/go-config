package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type IntTests struct {
	suite.Suite
}

func TestInt(t *testing.T) {
	suite.Run(t, new(IntTests))
}

func (s *IntTests) TestInterface() {
	var conf Config
	b := Int{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *IntTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	intConf := &Int{}
	intConf.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetInt("no-shorthand")
	s.Nil(err)
	s.Equal(int(0), val)

	intConf = &Int{Default: -1}
	intConf.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetInt("with-default")
	s.Nil(err)
	s.Equal(-1, val)

	intConf = &Int{Shorthand: "s"}
	intConf.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetInt("with-shorthand")
	s.Nil(err)
	s.Equal(int(0), val)
}

func (s *IntTests) Test_IsSetExplicitlyByFlag() {
	flags := &pflag.FlagSet{}
	conf := &Int{}
	conf.ApplyToFlagSet("test", flags)
	s.False(conf.IsSetExplicitlyByFlag())
	flags.Set("test", "1")
	s.True(conf.IsSetExplicitlyByFlag())
}

func (s *IntTests) Test_IsSet() {
	conf := &Int{}
	s.False(conf.IsSet())
	s.Nil(conf.SetValue(1))
	s.True(conf.IsSet())
}

func (s *IntTests) Test_GettersSetters() {
	conf := &Int{
		Default:   1,
		Value:     2,
		Shorthand: "t",
		Usage:     "usage",
	}
	s.Equal(1, conf.GetDefault())
	s.Equal(2, conf.GetValue())
	s.Equal("t", conf.GetShorthand())
	s.Equal("usage", conf.GetUsage())
	valuePointer, ok := conf.GetValuePointer().(*int)
	s.True(ok)
	s.Equal(conf.GetValue(), *valuePointer)
	s.Nil(conf.SetValue(3))
	s.Equal(3, conf.GetValue())
}
