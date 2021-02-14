package main

import (
	"github.com/usvc/go-config"
)

const (
	DefaultBoolean   = false
	DefaultFloat     = 1.23456
	DefaultInt       = -1
	DefaultString    = "default"
	DefaultUint      = 1
	ParamBoolean     = "bool"
	ParamFloat       = "float"
	ParamInt         = "int"
	ParamIntSlice    = "int slice"
	ParamString      = "string"
	ParamStringSlice = "string slice"
	ParamUint        = "uint"
)

var (
	DefaultIntSlice    = []int{-2, -3, -4}
	DefaultStringSlice = []string{"hello", "world"}
)

func NewConfiguration() config.Map {
	return config.Map{
		// with env : BOOL=true
		// with flag: --bool | -b
		ParamBoolean: &config.Bool{
			Default:   DefaultBoolean,
			Shorthand: "b",
			Usage:     "specifies a boolean value",
		},
		// with env : FLOAT=-123
		// with flag: --float -123 | -f -123
		ParamFloat: &config.Float{
			Default:   DefaultFloat,
			Shorthand: "f",
			Usage:     "specifies a floating point value",
		},
		// with env : INT=-123
		// with flag: --int -123 | -i -123
		ParamInt: &config.Int{
			Default:   DefaultInt,
			Shorthand: "i",
			Usage:     "specifies a signed integer value",
		},
		// with env : INTS="-123 -456"
		// with flag: --int-slice -123,-456 | -I -123,-456
		ParamIntSlice: &config.IntSlice{
			Default:   DefaultIntSlice,
			Shorthand: "I",
			Usage:     "specifies a slice of signed integers value",
		},
		// with env : STRING=value
		// with flag: --string value | -s value
		ParamString: &config.String{
			Default:   DefaultString,
			Shorthand: "s",
			Usage:     "specifies a string value",
		},
		// with env : STRING_SLICE="value1 value2"
		// with flag: --string-slice value1,value2 | -S value1,value2
		ParamStringSlice: &config.StringSlice{
			Default:   DefaultStringSlice,
			Shorthand: "S",
			Usage:     "specifies a slice of strings value",
		},
		// with env : UINT=123
		// with flag: --uint 123 | -u 123
		ParamUint: &config.Uint{
			Default:   DefaultUint,
			Shorthand: "u",
			Usage:     "specifies an unsigned integer value",
		},
	}
}
