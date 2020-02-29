package config

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type MapTests struct {
	suite.Suite
}

func TestMap(t *testing.T) {
	suite.Run(t, new(MapTests))
}

func (s *MapTests) TestApplyToCobra() {
	conf := Map{
		"testing string": &String{Default: "default"},
	}
	cmd := &cobra.Command{}
	conf.ApplyToCobra(cmd)
	str, err := cmd.Flags().GetString("testing-string")
	s.Nil(err)
	s.Equal("default", str)
}

func (s *MapTests) TestGetBool() {
	expectedName := "bool"
	expectedValue := true
	expectedConfig := Bool{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetBool(expectedName))
}

func (s *MapTests) TestGetFloat() {
	expectedName := "float"
	expectedValue := 3.142
	expectedConfig := Float{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetFloat(expectedName))
}

func (s *MapTests) TestGetInt() {
	expectedName := "int"
	expectedValue := -12345
	expectedConfig := Int{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetInt(expectedName))
}

func (s *MapTests) TestGetIntSlice() {
	expectedName := "int-slice"
	expectedValue := []int{-1, -2, -3}
	expectedConfig := IntSlice{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetIntSlice(expectedName))
}

func (s *MapTests) TestGetString() {
	expectedName := "string"
	expectedValue := "hello world"
	expectedConfig := String{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetString(expectedName))
}

func (s *MapTests) TestGetStringSlice() {
	expectedName := "string-slice"
	expectedValue := []string{"hello", "world"}
	expectedConfig := StringSlice{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetStringSlice(expectedName))
}

func (s *MapTests) TestGetUint() {
	expectedName := "uint"
	expectedValue := uint(1234567890)
	expectedConfig := Uint{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetUint(expectedName))
}

func (s *MapTests) TestGetUintSlice() {
	expectedName := "uint-slice"
	expectedValue := []uint{12345, 67890}
	expectedConfig := UintSlice{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetUintSlice(expectedName))
}
