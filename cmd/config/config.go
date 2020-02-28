package main

import (
	"github.com/usvc/config"
)

var conf = config.Map{
	// with env : BOOL=true
	// with flag: --bool | -b
	"bool": &config.Bool{
		Default:   false,
		Shorthand: "b",
		Usage:     "specifies a boolean value",
	},
	// with env : FLOAT=-123
	// with flag: --float -123 | -f -123
	"float": &config.Float{
		Default:   1.6180339887498948482045868343,
		Shorthand: "f",
		Usage:     "specifies a floating point value",
	},
	// with env : INT=-123
	// with flag: --int -123 | -i -123
	"int": &config.Int{
		Default:   -1,
		Shorthand: "i",
		Usage:     "specifies a signed integer value",
	},
	// with env : INTS="-123 -456"
	// with flag: --int-slice -123,-456 | -I -123,-456
	"int slice": &config.IntSlice{
		Default:   []int{-2, -3},
		Shorthand: "I",
		Usage:     "specifies a slice of signed integers value",
	},
	// with env : STRING=value
	// with flag: --string value | -s value
	"string": &config.String{
		Default:   "default",
		Shorthand: "s",
		Usage:     "specifies a string value",
	},
	// with env : STRING_SLICE="value1 value2"
	// with flag: --string-slice value1,value2 | -S value1,value2
	"string slice": &config.StringSlice{
		Default:   []string{"hello", "world"},
		Shorthand: "S",
		Usage:     "specifies a slice of strings value",
	},
	// with env : UINT=123
	// with flag: --uint 123 | -u 123
	"uint": &config.Uint{
		Default:   1,
		Shorthand: "u",
		Usage:     "specifies an unsigned integer value",
	},
	// with env : UINT_SLICE="123 456"
	// with flag: --uint-slice 123,456 | -U 123,456
	"uint slice": &config.UintSlice{
		Default:   []uint{2, 3},
		Shorthand: "U",
		Usage:     "specifies a slice of unsigned integers value",
	},
}
