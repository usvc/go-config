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
	confTypes := []Config{
		&Bool{},
		&Float{},
		&Int{},
		&IntSlice{},
		&String{},
		&StringSlice{},
		&Uint{},
	}
	for _, confType := range confTypes {
		s.Zero(confType.GetValue())
	}
}
