package config

import (
	"github.com/spf13/pflag"
)

type Config interface {
	ApplyToFlagSet(name string, flagset *pflag.FlagSet)
	GetDefault() interface{}
	GetShorthand() string
	GetUsage() string
	GetValue() interface{}
	GetValuePointer() interface{}
	SetValue(value interface{})
}
