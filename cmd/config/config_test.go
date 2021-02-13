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

type e2eTests struct {
	suite.Suite
	command                          *cobra.Command
	expectedStringSliceOutput        string
	expectedDefaultStringSliceOutput string
	inputStringSlice                 []string
	inputStringSliceAlternate        []string
	output                           *bytes.Buffer
}

func Test_e2e(t *testing.T) {
	suite.Run(t, &e2eTests{})
}

func (s *e2eTests) BeforeTest(suite, test string) {
	s.inputStringSlice = []string{"a", "ab", "abc"}
	s.inputStringSliceAlternate = []string{"z", "zy", "zyx"}
	s.expectedStringSliceOutput = fmt.Sprintf("%s: [a ab abc] (length: 3)", ParamStringSlice)
	s.expectedDefaultStringSliceOutput = fmt.Sprintf("%s: [%s] (length: 2)", ParamStringSlice, "hello world")
}

func (s *e2eTests) SetupTest() {
	fmt.Println("SetupTest")
	s.output = bytes.NewBuffer([]byte{})
	s.command = GetCommand()
	conf = NewConfiguration()
	conf.ApplyToCobra(s.command)
	s.command.SetOut(s.output)
}

func (s e2eTests) Test_StringSlice_env_comma_delimited() {
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSlice, ","))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should correctly parse comma delimited values in the envionment")
}

func (s e2eTests) Test_StringSlice_env_space_delimited() {
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSlice, " "))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should correctly parse space delimited values in the envionment")
}

func (s e2eTests) Test_StringSlice_flag_flag_delimited() {
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

func (s e2eTests) Test_StringSlice_flag_comma_delimited() {
	s.command.SetArgs([]string{"--string-slice", strings.Join(s.inputStringSlice, ",")})
	defer s.command.SetArgs([]string{})
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should correctly parse comma delimited values in flags")
}

func (s e2eTests) Test_StringSlice_priority_defaults() {
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedDefaultStringSliceOutput,
		"it should consume the default value when neither environment nor flag is set")
}

func (s e2eTests) Test_StringSlice_priority_only_flag() {
	s.command.SetArgs([]string{"--string-slice", strings.Join(s.inputStringSlice, ",")})
	defer s.command.SetArgs([]string{})
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should consume the flag's value when only the flag value is set")
}

func (s e2eTests) Test_StringSlice_priority_only_env() {
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSlice, ","))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should consume the environment value when only environment value is set")
}

func (s e2eTests) Test_StringSlice_priority_flag_and_env() {
	s.command.SetArgs([]string{"--string-slice", strings.Join(s.inputStringSlice, ",")})
	defer s.command.SetArgs([]string{})
	os.Setenv("STRING_SLICE", strings.Join(s.inputStringSliceAlternate, ","))
	defer os.Setenv("STRING_SLICE", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedStringSliceOutput,
		"it should consume the flag's value when both environment and flag is set")
}
