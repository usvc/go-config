package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type UintSliceTests struct {
	suite.Suite
}

func TestUintSlice(t *testing.T) {
	suite.Run(t, new(UintSliceTests))
}

func (s *UintSliceTests) TestInterface() {
	var conf Config
	b := UintSlice{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *UintSliceTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	uintSlice := &UintSlice{}
	uintSlice.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetUintSlice("no-shorthand")
	s.Nil(err)
	s.Equal(int(0), val)

	uintSlice = &UintSlice{Default: []uint{1, 2}}
	uintSlice.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetUintSlice("with-default")
	s.Nil(err)
	s.Equal([]int{1, 2}, val)

	uintSlice = &UintSlice{Shorthand: "s"}
	uintSlice.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetUintSlice("with-shorthand")
	s.Nil(err)
	s.Equal(int(0), val)
}
