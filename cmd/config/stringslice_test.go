package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type StringSliceTests struct {
	suite.Suite
	command                          *cobra.Command
	expectedStringSliceOutput        string
	expectedDefaultStringSliceOutput string
	inputStringSlice                 []string
	inputStringSliceAlternate        []string
	output                           *bytes.Buffer
}

func TestSTringSlice(t *testing.T) {
	suite.Run(t, &StringSliceTests{})
}

func (s *StringSliceTests) BeforeTest(suite, test string) {
	s.inputStringSlice = []string{"a", "ab", "abc"}
	s.inputStringSliceAlternate = []string{"z", "zy", "zyx"}
	s.expectedStringSliceOutput = fmt.Sprintf("%s: [a ab abc] (length: 3)", ParamStringSlice)
	defaultStringSlice := strings.Join(DefaultStringSlice, " ")
	s.expectedDefaultStringSliceOutput = fmt.Sprintf("%s: [%s] (length: %v)", ParamStringSlice, defaultStringSlice, len(DefaultStringSlice))
}

func (s *StringSliceTests) SetupTest() {
	s.output = bytes.NewBuffer([]byte{})
	command := GetCommand()
	conf = NewConfiguration()
	conf.ApplyToCobra(command)
	command.SetOut(s.output)
	s.command = command
}

func (s StringSliceTests) Test_env_comma_delimited() {
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSlice, ","))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should correctly parse comma delimited values in the envionment")
}

func (s StringSliceTests) Test_env_space_delimited() {
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSlice, " "))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should correctly parse space delimited values in the envionment")
}

func (s StringSliceTests) Test_flag_flag_delimited() {
	args := []string{}
	for i := 0; i < len(s.inputStringSlice); i++ {
		args = append(args, "--string-slice", s.inputStringSlice[i])
	}
	s.command.SetArgs(args)
	defer s.command.SetArgs([]string{})
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should correctly parse flag delimited values in flags")
}

func (s StringSliceTests) Test_flag_comma_delimited() {
	s.command.SetArgs([]string{"--string-slice", strings.Join(s.inputStringSlice, ",")})
	defer s.command.SetArgs([]string{})
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should correctly parse comma delimited values in flags")
}

func (s StringSliceTests) Test_priority_defaults() {
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedDefaultStringSliceOutput,
		"it should consume the default value when neither environment nor flag is set")
}

func (s StringSliceTests) Test_priority_only_flag() {
	s.command.SetArgs([]string{"--string-slice", strings.Join(s.inputStringSlice, ",")})
	defer s.command.SetArgs([]string{})
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should consume the flag's value when only the flag value is set")
}

func (s StringSliceTests) Test_priority_only_env() {
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSlice, ","))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should consume the environment value when only environment value is set")
}

func (s StringSliceTests) Test_priority_flag_and_env() {
	s.command.SetArgs([]string{"--string-slice", strings.Join(s.inputStringSlice, ",")})
	defer s.command.SetArgs([]string{})
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSliceAlternate, ","))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should consume the string slice flag's value when both environment and flag is set")
}

func (s StringSliceTests) Test_String_priority_flag_and_env() {
	s.command.SetArgs([]string{"--string", "hello world"})
	defer s.command.SetArgs([]string{})
	os.Setenv("STRING", "hola mundo")
	defer os.Setenv("STRING", "")
	s.command.Execute()
	s.Contains(s.output.String(), "string: hello world",
		"it should consume the string flag's value when both environment and flag is set")
}
