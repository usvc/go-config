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
	s.Equal(int(0), val)

	intSlice = &IntSlice{Default: []int{-1, -2}}
	intSlice.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetIntSlice("with-default")
	s.Nil(err)
	s.Equal([]int{-1, -2}, val)

	intSlice = &IntSlice{Shorthand: "s"}
	intSlice.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetIntSlice("with-shorthand")
	s.Nil(err)
	s.Equal(int(0), val)
}
