package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type FloatTests struct {
	suite.Suite
}

func TestFloat(t *testing.T) {
	suite.Run(t, new(FloatTests))
}

func (s *FloatTests) TestInterface() {
	var conf Config
	b := Float{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *FloatTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	floatConf := &Float{}
	floatConf.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetFloat64("no-shorthand")
	s.Nil(err)
	s.Equal(float64(0), val)

	floatConf = &Float{Default: 3.142}
	floatConf.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetFloat64("with-default")
	s.Nil(err)
	s.Equal(3.142, val)

	floatConf = &Float{Shorthand: "s"}
	floatConf.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetFloat64("with-shorthand")
	s.Nil(err)
	s.Equal(float64(0), val)
}
