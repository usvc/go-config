package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type e2eTests struct {
	suite.Suite
}

func Test_e2e(t *testing.T) {
	suite.Run(t, &e2eTests{})
}

func (s e2eTests) Test_StringSlice_EnvFormat() {
	var out bytes.Buffer
	command := GetCommand()
	conf.ApplyToCobra(command)
	command.SetOut(&out)

	os.Setenv("STRING_SLICE", "g,h,i")
	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [g h i] (length: 3)", ParamStringSlice),
		"it should correctly parse comma delimited values in the envionment")
	out.Reset()
	conf.Reset()

	os.Setenv("STRING_SLICE", "j k l")
	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [j k l] (length: 3)", ParamStringSlice),
		"it should correctly parse space delimited values in the envionment")
	out.Reset()
	conf.Reset()
	os.Setenv("STRING_SLICE", "")
}

func (s e2eTests) Test_StringSlice_FlagFormat() {
	var out bytes.Buffer
	command := GetCommand()
	conf.ApplyToCobra(command)
	command.SetOut(&out)

	command.SetArgs([]string{"--string-slice", "a", "--string-slice", "b", "--string-slice", "c"})
	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [a b c] (length: 3)", ParamStringSlice),
		"it should correctly parse flag delimited values in flags")
	out.Reset()
	conf.Reset()

	command.SetArgs([]string{"--string-slice", "d,e,f"})
	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [d e f] (length: 3)", ParamStringSlice),
		"it should correctly parse comma delimited values in flags")
	out.Reset()
	conf.Reset()
}

func (s e2eTests) Test_StringSlice_Input() {
	var out bytes.Buffer
	command := GetCommand()
	conf.ApplyToCobra(command)
	command.SetOut(&out)

	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [hello world] (length: 2)", ParamStringSlice),
		"it should consume the default value when neither environment nor flag is set")
	out.Reset()
	conf.Reset()

	command.SetArgs([]string{"--string-slice", "a,b,c"})
	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [a b c] (length: 3)", ParamStringSlice),
		"it should consume the flag's value when only the flag value is set")
	out.Reset()
	conf.Reset()
	command.SetArgs([]string{})

	os.Setenv("STRING_SLICE", "d,e,f")
	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [d e f] (length: 3)", ParamStringSlice),
		"it should consume the environment value when only environment value is set")
	out.Reset()
	conf.Reset()

	command.SetArgs([]string{"--string-slice", "g,h,i"})
	command.Execute()
	s.Contains(out.String(), fmt.Sprintf("%s: [g h i] (length: 3)", ParamStringSlice),
		"it should consume the flag's value when both environment and flag is set")
	out.Reset()
	conf.Reset()
}
