package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTests struct {
	suite.Suite
}

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTests))
}

func (s *ConfigTests) TestTypes() {
	var conf Config
	conf = &Bool{}
	conf = &Float{}
	conf = &Int{}
	conf = &IntSlice{}
	conf = &String{}
	conf = &StringSlice{}
	conf = &Uint{}
	conf = &UintSlice{}
	s.Nil(conf.GetValue())
}
