package config

import (
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

func isZeroValue(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v == *new(bool)
	case int:
		return v == *new(int)
	case string:
		return v == *new(string)
	case uint:
		return v == *new(uint)
	case float64:
		return v == *new(float64)
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
