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
