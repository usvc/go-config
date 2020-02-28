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
