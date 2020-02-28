package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type UintTests struct {
	suite.Suite
}

func TestUint(t *testing.T) {
	suite.Run(t, new(UintTests))
}

func (s *UintTests) TestInterface() {
	var conf Config
	b := Uint{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *UintTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	uintConf := &Uint{}
	uintConf.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetUint("no-shorthand")
	s.Nil(err)
	s.Equal(uint(0), val)

	uintConf = &Uint{Default: 1}
	uintConf.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetUint("with-default")
	s.Nil(err)
	s.Equal(uint(1), val)

	uintConf = &Uint{Shorthand: "s"}
	uintConf.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetUint("with-shorthand")
	s.Nil(err)
	s.Equal(uint(0), val)
}
