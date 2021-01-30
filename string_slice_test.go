package config

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

type StringSliceTests struct {
	suite.Suite
}

func TestStringSlice(t *testing.T) {
	suite.Run(t, new(StringSliceTests))
}

func (s *StringSliceTests) TestInterface() {
	var conf Config
	b := StringSlice{}
	conf = &b
	s.Equal(conf.GetValue(), b.Value)
	s.Equal(conf.GetValuePointer(), &b.Value)
	s.Equal(conf.GetDefault(), b.Default)
	s.Equal(conf.GetUsage(), b.Usage)
	s.Equal(conf.GetShorthand(), b.Shorthand)
}

func (s *StringSliceTests) TestApplyToFlagSet() {
	flagSet := &pflag.FlagSet{}
	stringSlice := &StringSlice{}
	stringSlice.ApplyToFlagSet("no-shorthand", flagSet)
	val, err := flagSet.GetStringSlice("no-shorthand")
	s.Nil(err)
	s.Equal([]string{}, val)

	stringSlice = &StringSlice{Default: []string{"a", "b"}}
	stringSlice.ApplyToFlagSet("with-default", flagSet)
	val, err = flagSet.GetStringSlice("with-default")
	s.Nil(err)
	s.Equal([]string{"a", "b"}, val)

	stringSlice = &StringSlice{Shorthand: "s"}
	stringSlice.ApplyToFlagSet("with-shorthand", flagSet)
	val, err = flagSet.GetStringSlice("with-shorthand")
	s.Nil(err)
	s.Equal([]string{}, val)
}

func (s *StringSliceTests) TestLoadFromEnvironment() {
	os.Setenv("TEST_STRING_SLICE", "a,b,c")
	testMap := Map{
		"test-string-slice": &StringSlice{
			Default: []string{"d", "e", "f"},
		},
	}
	testMap.LoadFromEnvironment()
	stringSlice := testMap.GetStringSlice("test-string-slice")
	s.Equal([]string{"a", "b", "c"}, stringSlice)
}
