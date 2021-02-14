package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type CommandTests struct {
	suite.Suite
	command *cobra.Command
	output  *bytes.Buffer
}

func TestCommand(t *testing.T) {
	suite.Run(t, &CommandTests{})
}

func (s *CommandTests) SetupTest() {
	s.output = bytes.NewBuffer([]byte{})
	command := GetCommand()
	conf = NewConfiguration()
	conf.ApplyToCobra(command)
	command.SetOut(s.output)
	s.command = command
}

func (s CommandTests) Test_env() {
	os.Setenv("BOOL", "true")
	os.Setenv("FLOAT", "3.142")
	os.Setenv("INT", "-2")
	os.Setenv("INT_SLICE", "3,-4,5,-6,7")
	os.Setenv("STRING", "not default")
	os.Setenv("STRING_SLICE", "a,b,c,d,e")
	os.Setenv("UINT", "2")
	os.Setenv("UINT_SLICE", "3,4,5,6,7")
	s.command.Execute()
	s.Contains(s.output.String(), "bool: true")
	s.Contains(s.output.String(), "float: 3.142")
	s.Contains(s.output.String(), "int: -2")
	s.Contains(s.output.String(), "int slice: [3 -4 5 -6 7]")
	s.Contains(s.output.String(), "string: not default")
	s.Contains(s.output.String(), "string slice: [a b c d e]")
	s.Contains(s.output.String(), "uint: 2")
}

func (s CommandTests) Test_flags() {
	s.command.SetArgs([]string{
		"--bool", "true",
		"--float", "3.142",
		"--int", "-2",
		"--int-slice", "3,-4,5,-6,7",
		"--string", "a",
		"--string-slice", "a,b,c,d,e",
		"--uint", "2",
	})
	s.command.Execute()
	s.Contains(s.output.String(), "bool: true")
	s.Contains(s.output.String(), "float: 3.142")
	s.Contains(s.output.String(), "int: -2")
	s.Contains(s.output.String(), "int slice: [3 -4 5 -6 7]")
	s.Contains(s.output.String(), "string: a")
	s.Contains(s.output.String(), "string slice: [a b c d e]")
	s.Contains(s.output.String(), "uint: 2")
}

func (s CommandTests) Test_priority_flag_and_env() {
	s.command.SetArgs([]string{
		"--bool", "true",
		"--float", "3.142",
		"--int", "1",
		"--int-slice", "1,-2,3,-4,5",
		"--string", "a",
		"--string-slice", "a,b,c,d,e",
		"--uint", "1",
	})
	os.Setenv("BOOL", "false")
	os.Setenv("FLOAT", "1.618")
	os.Setenv("INT", "2")
	os.Setenv("INT_SLICE", "2,-3,4,-5,6")
	os.Setenv("STRING", "b")
	os.Setenv("STRING_SLICE", "b,c,d,e,f")
	os.Setenv("UINT", "2")
	os.Setenv("UINT_SLICE", "2,3,4,5,6")
	s.command.Execute()
	s.Contains(s.output.String(), "bool: true")
	s.Contains(s.output.String(), "float: 3.142")
	s.Contains(s.output.String(), "int: 1")
	s.Contains(s.output.String(), "int slice: [1 -2 3 -4 5]")
	s.Contains(s.output.String(), "string: a")
	s.Contains(s.output.String(), "string slice: [a b c d e]")
	s.Contains(s.output.String(), "uint: 1")
}
