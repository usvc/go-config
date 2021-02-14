package config

import (
	"os"
	"strconv"
	"strings"
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

func (s *MapTests) TestApplyToCobraPersistent() {
	conf := Map{
		"testing string": &String{Default: "default"},
	}
	cmd := &cobra.Command{}
	conf.ApplyToCobraPersistent(cmd)
	str, err := cmd.PersistentFlags().GetString("testing-string")
	s.Nil(err)
	s.Equal("default", str)
}

func (s *MapTests) TestLoadFromEnvironment_bool() {
	conf := Map{
		"lfe_bool": &Bool{Default: false},
	}
	originalLFEBool := os.Getenv("LFE_BOOL")
	defer os.Setenv("LFE_BOOL", originalLFEBool)
	expectedLFEBool := true
	os.Setenv("LFE_BOOL", "1")
	conf.LoadFromEnvironment()
	s.Equal(expectedLFEBool, conf.GetBool("lfe_bool"))
}

func (s *MapTests) TestLoadFromEnvironment_float() {
	conf := Map{
		"lfe_float": &Float{Default: 3.142},
	}
	originalLFEFloat := os.Getenv("LFE_FLOAT")
	defer os.Setenv("LFE_FLOAT", originalLFEFloat)
	expectedLFEFloat := 1.618
	os.Setenv("LFE_FLOAT", strconv.FormatFloat(expectedLFEFloat, 'f', 3, 64))
	conf.LoadFromEnvironment()
	s.Equal(expectedLFEFloat, conf.GetFloat("lfe_float"))
}

func (s *MapTests) TestLoadFromEnvironment_int() {
	conf := Map{
		"lfe_int": &Int{Default: 42},
	}
	originalLFEInt := os.Getenv("LFE_INT")
	defer os.Setenv("LFE_INT", originalLFEInt)
	expectedLFEInt := 43
	os.Setenv("LFE_INT", strconv.Itoa(expectedLFEInt))
	conf.LoadFromEnvironment()
	s.Equal(expectedLFEInt, conf.GetInt("lfe_int"))
}

func (s *MapTests) TestLoadFromEnvironment_string() {
	conf := Map{
		"lfe_string": &String{Default: "default"},
	}
	originalLFEString := os.Getenv("LFE_STRING")
	defer os.Setenv("LFE_STRING", originalLFEString)
	expectedLFEString := "expected"
	os.Setenv("LFE_STRING", expectedLFEString)
	conf.LoadFromEnvironment()
	s.Equal(expectedLFEString, conf.GetString("lfe_string"))
}

func (s *MapTests) TestLoadFromEnvironment_stringSlice() {
	conf := Map{
		"lfe_string_slice": &StringSlice{Default: []string{"default", "string"}},
	}
	originalLFEStringSlice := os.Getenv("LFE_STRING_SLICE")
	defer func() {
		os.Setenv("LFE_STRING_SLICE", originalLFEStringSlice)
	}()
	expectedLFEStringSlice := "expected text"
	os.Setenv("LFE_STRING_SLICE", expectedLFEStringSlice)
	conf.LoadFromEnvironment()
	s.EqualValues(strings.Split(expectedLFEStringSlice, " "), conf.GetStringSlice("lfe_string_slice"))
}

func (s *MapTests) TestLoadFromEnvironment_uint() {
	conf := Map{
		"lfe_uint": &Uint{Default: 748},
	}
	originalLFEUint := os.Getenv("LFE_UINT")
	defer func() {
		os.Setenv("LFE_UINT", originalLFEUint)
	}()
	expectedLFEUint := uint(749)
	os.Setenv("LFE_UINT", strconv.FormatUint(uint64(expectedLFEUint), 10))
	conf.LoadFromEnvironment()
	s.Equal(expectedLFEUint, conf.GetUint("lfe_uint"))
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
