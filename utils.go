package config

import (
	"fmt"
	"strings"
)

const (
	byteA         rune = 'a'
	byteZ         rune = 'z'
	byte0         rune = '0'
	byte9         rune = '9'
	separatorEnv  rune = '_'
	separatorFlag rune = '-'
)

func areEqualFloatSlice(sliceA, sliceB []float64) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, value := range sliceA {
		if value != sliceB[index] {
			return false
		}
	}
	return true
}

func areEqualIntSlice(sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, value := range sliceA {
		if value != sliceB[index] {
			return false
		}
	}
	return true
}

func areEqualStringSlice(sliceA, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, value := range sliceA {
		if value != sliceB[index] {
			return false
		}
	}
	return true
}

func assertIDExists(configMap Map, id string) {
	if configMap[id] == nil {
		var environmentKeys []string
		for key := range configMap {
			environmentKeys = append(environmentKeys, fmt.Sprintf("'%s'", key))
		}
		panic(fmt.Errorf("provided id '%s' could not be found (available keys: %v)", id, environmentKeys))
	}
}

func isZeroValue(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v == *new(bool)
	case float64:
		return v == *new(float64)
	case int:
		return v == *new(int)
	case []int:
		return areEqualIntSlice(v, []int{})
	case string:
		return v == *new(string)
	case []string:
		return areEqualStringSlice(v, []string{})
	case uint:
		return v == *new(uint)
	default:
		return v == nil
	}
}

func normalizeName(name string, separator ...rune) string {
	selectedSeparator := *new(rune)
	if len(separator) > 0 {
		selectedSeparator = separator[0]
	}
	normalizedName := []rune(strings.ToLower(name))
	for n := 0; n < len(normalizedName); n++ {
		char := normalizedName[n]
		isAlphabetic := (char >= byteA && char <= byteZ)
		isNumeric := (char >= byte0 && char <= byte9)
		if !(isAlphabetic || isNumeric) {
			normalizedName[n] = selectedSeparator
		}
	}
	returnedName := string(normalizedName)
	return returnedName
}
