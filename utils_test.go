package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, new(UtilsTests))
}

func (s *UtilsTests) Test_areEqualFloatSlice_equal() {
	var equals = [][][]float64{
		[][]float64{[]float64{1.1, 2.2}, []float64{1.1, 2.2}},
		[][]float64{[]float64{-1.1, -2.2}, []float64{-1.1, -2.2}},
	}
	for _, cases := range equals {
		s.True(areEqualFloatSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualFloatSlice_unequalLength() {
	var unequals = [][][]float64{
		[][]float64{[]float64{1.1}, []float64{1.1, 2.2}},
		[][]float64{[]float64{-1.1}, []float64{-1.1, 2.2}},
	}
	for _, cases := range unequals {
		s.False(areEqualFloatSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualFloatSlice_unequalValue() {
	var unequals = [][][]float64{
		[][]float64{[]float64{1.1, 2.2}, []float64{1.1, -2.2}},
		[][]float64{[]float64{-1.1, -2.2}, []float64{-1.1, 2.2}},
	}
	for _, cases := range unequals {
		s.False(areEqualFloatSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualIntSlice_equal() {
	var equals = [][][]int{
		[][]int{[]int{1, 2}, []int{1, 2}},
		[][]int{[]int{-1, -2}, []int{-1, -2}},
	}
	for _, cases := range equals {
		s.True(areEqualIntSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualIntSlice_unequalLength() {
	var unequals = [][][]int{
		[][]int{[]int{1}, []int{1, 2}},
		[][]int{[]int{-1}, []int{-1, 2}},
	}
	for _, cases := range unequals {
		s.False(areEqualIntSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualIntSlice_unequalValue() {
	var unequals = [][][]int{
		[][]int{[]int{1, 2}, []int{1, -2}},
		[][]int{[]int{-1, -2}, []int{-1, 2}},
	}
	for _, cases := range unequals {
		s.False(areEqualIntSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualStringSlice_equal() {
	var equalStringSlices = [][][]string{
		[][]string{[]string{"a", "b"}, []string{"a", "b"}},
		[][]string{[]string{"1", "2"}, []string{"1", "2"}},
	}
	for _, cases := range equalStringSlices {
		s.True(areEqualStringSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualStringSlice_unequalLength() {
	var unequalStringSlices = [][][]string{
		[][]string{[]string{"a"}, []string{"a", "b"}},
		[][]string{[]string{"1"}, []string{"1", "2"}},
	}
	for _, cases := range unequalStringSlices {
		s.False(areEqualStringSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_areEqualStringSlice_unequalValue() {
	var unequalStringSlices = [][][]string{
		[][]string{[]string{"a", "2"}, []string{"a", "b"}},
		[][]string{[]string{"1", "b"}, []string{"1", "2"}},
	}
	for _, cases := range unequalStringSlices {
		s.False(areEqualStringSlice(cases[0], cases[1]))
	}
}

func (s *UtilsTests) Test_assertIDExists() {
	expectedKey := "__not_found"
	defer func() {
		if r := recover(); r != nil {
			s.Contains(fmt.Sprintf("%s", r), fmt.Sprintf("'%s' could not be found", expectedKey))
		} else {
			s.False(true)
		}
	}()
	conf := Map{}
	conf.GetString(expectedKey)
}

func (s *UtilsTests) Test_isZeroValue() {
	var boolValue bool
	s.True(isZeroValue(boolValue))
	var intValue int
	s.True(isZeroValue(intValue))
	var stringValue string
	s.True(isZeroValue(stringValue))
	var uintValue uint
	s.True(isZeroValue(uintValue))
	var float64Value float64
	s.True(isZeroValue(float64Value))
	s.True(isZeroValue(nil))
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

func (s *UtilsTests) Test_shouldEnvironmentVariableBeSet() {
	envValue := "a"
	conf := &String{
		Default: "a",
		Value:   "a",
	}
	s.False(shouldEnvironmentVariableBeSet(envValue, conf))
}
