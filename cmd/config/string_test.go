package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type StringTests struct {
	suite.Suite
	command               *cobra.Command
	expectedOutput        string
	expectedDefaultOutput string
	input                 string
	inputAlternate        string
	output                *bytes.Buffer
}

func TestString(t *testing.T) {
	suite.Run(t, &StringTests{})
}

func (s *StringTests) BeforeTest(suite, test string) {
	s.input = "hello world"
	s.inputAlternate = "hola mundo"
	s.expectedOutput = fmt.Sprintf("%s: %s", ParamString, s.input)
	s.expectedDefaultOutput = fmt.Sprintf("%s: %s", ParamString, DefaultString)
}

func (s *StringTests) SetupTest() {
	s.output = bytes.NewBuffer([]byte{})
	command := GetCommand()
	conf = NewConfiguration()
	conf.ApplyToCobra(command)
	command.SetOut(s.output)
	s.command = command
}

func (s StringTests) Test_env() {
	os.Setenv("STRING", s.input)
	defer os.Setenv("STRING", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedOutput,
		"it should correctly parse string values in the envionment")
}

func (s StringTests) Test_flag() {
	s.command.SetArgs([]string{"--string", s.input})
	defer s.command.SetArgs([]string{})
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedOutput,
		"it should correctly parse string values in flags")
}

func (s StringTests) Test_priority_defaults() {
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedDefaultOutput,
		"it should consume the default value when neither environment nor flag is set")
}

func (s StringTests) Test_priority_only_flag() {
	s.command.SetArgs([]string{"--string", s.input})
	defer s.command.SetArgs([]string{})
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedOutput,
		"it should consume the flag's value when only the flag value is set")
}

func (s StringTests) Test_priority_only_env() {
	os.Setenv("STRING", s.input)
	defer os.Setenv("STRING", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedOutput,
		"it should consume the environment value when only environment value is set")
}

func (s StringTests) Test_priority_flag_and_env() {
	s.command.SetArgs([]string{"--string", s.input})
	defer s.command.SetArgs([]string{})
	os.Setenv("STRING", s.inputAlternate)
	defer os.Setenv("STRING", "")
	s.command.Execute()
	s.Contains(s.output.String(), s.expectedOutput,
		"it should consume the string slice flag's value when both environment and flag is set")
}
