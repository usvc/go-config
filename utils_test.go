package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, new(UtilsTests))
}

func (s *UtilsTests) Test_normalizeName_cases() {
	testCases := map[string]string{
		"abcd": "abcd",
		"ABCD": "abcd",
		"1234": "1234",
		"1bc4": "1bc4",
		"1BC4": "1bc4",
		"1bcD": "1bcd",
		"Abc4": "abc4",
	}
	for input, expected := range testCases {
		s.Equal(expected, normalizeName(input, '_'))
	}
}

func (s *UtilsTests) Test_normalizeName_specials() {
	testCases := map[string]string{
		"`~!@":  "____", // doing an exhaustive test here
		"#$%^":  "____",
		"&*()":  "____",
		"-_=+":  "____",
		"[{]}":  "____",
		"\\|;:": "____",
		"'\",<": "____",
		".>/?":  "____",
		"a cd":  "a_cd",
		"ab!d":  "ab_d", // middle of string
		"abc!":  "abc_", // end of string
		"!bcd":  "_bcd", // start of string
	}
	for input, expected := range testCases {
		s.Equal(expected, normalizeName(input, '_'))
	}
}
