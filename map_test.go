package config

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type MapTests struct {
	suite.Suite
	ConfigMap Map
}

func TestMap(t *testing.T) {
	mapTests := new(MapTests)
	mapTests.ConfigMap = Map{
		"test apply to cobra": &String{Default: "default"},
	}
	suite.Run(t, mapTests)
}

func (s *MapTests) Test_ApplyToCobra_PreRun() {
	var output bytes.Buffer
	cmd := &cobra.Command{
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
		PreRun: func(command *cobra.Command, args []string) {},
	}
	cmd.SetOut(&output)
	s.ConfigMap.ApplyToCobra(cmd)
	s.Nil(cmd.Execute())
	str, err := cmd.Flags().GetString("test-apply-to-cobra")
	s.Nil(err)
	s.Equal("default", str)
	s.Contains(output.String(), "--test-apply-to-cobra")
}

func (s *MapTests) Test_ApplyToCobra_PreRunE() {
	var output bytes.Buffer
	cmd := &cobra.Command{
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
		PreRunE: func(command *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.SetOut(&output)
	s.ConfigMap.ApplyToCobra(cmd)
	s.Nil(cmd.Execute())
	str, err := cmd.Flags().GetString("test-apply-to-cobra")
	s.Nil(err)
	s.Equal("default", str)
	s.Contains(output.String(), "--test-apply-to-cobra")
}

func (s *MapTests) Test_ApplyToCobraPersistent_PreRun() {
	var output bytes.Buffer
	cmd := &cobra.Command{
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
		PersistentPreRun: func(command *cobra.Command, args []string) {
			output.Write([]byte("PersistentPreRun ran"))
		},
	}
	cmd.SetOut(&output)
	s.ConfigMap.ApplyToCobraPersistent(cmd)
	s.Nil(cmd.Execute())
	str, err := cmd.PersistentFlags().GetString("test-apply-to-cobra")
	s.Nil(err)
	s.Equal("default", str)
	s.Contains(output.String(), "--test-apply-to-cobra")
	s.Contains(output.String(), "PersistentPreRun ran")
}

func (s *MapTests) Test_ApplyToCobraPersistent_PreRunE() {
	var output bytes.Buffer
	cmd := &cobra.Command{
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
		PersistentPreRunE: func(command *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.SetOut(&output)
	s.ConfigMap.ApplyToCobraPersistent(cmd)
	s.Nil(cmd.Execute())
	str, err := cmd.PersistentFlags().GetString("test-apply-to-cobra")
	s.Nil(err)
	s.Equal("default", str)
	s.Contains(output.String(), "--test-apply-to-cobra")
}

func (s *MapTests) Test_LoadFromEnvironment_bool() {
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

func (s *MapTests) Test_LoadFromEnvironment_float() {
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

func (s *MapTests) Test_LoadFromEnvironment_int() {
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

func (s *MapTests) Test_LoadFromEnvironment_intSlice() {
	defaultIntSlice := []int{2021, 2, 15, 0, 32, 38, 800}
	conf := Map{
		"lfe_int_slice": &IntSlice{Default: defaultIntSlice},
	}
	originalLFEIntSlice := os.Getenv("LFE_INT_SLICE")
	defer os.Setenv("LFE_INT_SLICE", originalLFEIntSlice)
	expectedLFEIntSlice := []int{1, 2, 3, 4, 5}
	lfeIntAsString := "1,2,3,4,5"
	os.Setenv("LFE_INT_SLICE", lfeIntAsString)
	s.Equal(defaultIntSlice, conf.GetIntSlice("lfe_int_slice"))
	conf.LoadFromEnvironment()
	s.Equal(expectedLFEIntSlice, conf.GetIntSlice("lfe_int_slice"))
}

func (s *MapTests) Test_LoadFromEnvironment_string() {
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

func (s *MapTests) Test_LoadFromEnvironment_stringSlice() {
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

func (s *MapTests) Test_LoadFromEnvironment_uint() {
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

func (s *MapTests) Test_GetBool() {
	expectedName := "bool"
	expectedValue := true
	expectedConfig := Bool{Default: false}
	expectedConfig.SetValue(true)
	conf := Map{
		expectedName: &expectedConfig,
	}

	expectedValue = false
	expectedConfig = Bool{Default: true}
	expectedConfig.SetValue(false)
	conf = Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetBool(expectedName))
}

func (s *MapTests) Test_GetFloat() {
	expectedName := "float"
	expectedValue := 3.142
	expectedConfig := Float{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetFloat(expectedName))
}

func (s *MapTests) Test_GetInt() {
	expectedName := "int"
	expectedValue := -12345
	expectedConfig := Int{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetInt(expectedName))
}

func (s *MapTests) Test_GetIntSlice() {
	expectedName := "int-slice"
	expectedValue := []int{-1, -2, -3}
	expectedConfig := IntSlice{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetIntSlice(expectedName))
}

func (s *MapTests) Test_GetString() {
	expectedName := "string"
	expectedValue := "hello world"
	expectedConfig := String{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetString(expectedName))
}

func (s *MapTests) Test_GetStringSlice() {
	expectedName := "string-slice"
	expectedValue := []string{"hello", "world"}
	expectedConfig := StringSlice{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetStringSlice(expectedName))
}

func (s *MapTests) Test_GetUint() {
	expectedName := "uint"
	expectedValue := uint(1234567890)
	expectedConfig := Uint{Value: expectedValue}
	conf := Map{
		expectedName: &expectedConfig,
	}
	s.Equal(expectedValue, conf.GetUint(expectedName))
}

func (s *MapTests) Test_Reset() {
	conf := Map{
		"bool":         &Bool{Value: true},
		"float":        &Float{Value: 3.142},
		"int":          &Int{Value: -1},
		"int-slice":    &IntSlice{Value: []int{-1}},
		"string":       &String{Value: "a"},
		"string-slice": &StringSlice{Value: []string{"a"}},
		"uint":         &Uint{Value: 1},
	}
	s.Nil(conf.Reset())
	s.Equal(*new(bool), conf.GetBool("bool"))
	s.Equal(*new(float64), conf.GetFloat("float"))
	s.Equal(*new(int), conf.GetInt("int"))
	s.Equal(*new([]int), conf.GetIntSlice("int-slice"))
	s.Equal(*new(string), conf.GetString("string"))
	s.Equal(*new([]string), conf.GetStringSlice("string-slice"))
	s.Equal(*new(uint), conf.GetUint("uint"))
}
