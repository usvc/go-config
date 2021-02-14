package config

import (
	"github.com/spf13/pflag"
)

// Config defines an interface that all configuration keys
// should implement
type Config interface {
	ApplyToFlagSet(name string, flagset *pflag.FlagSet)
	GetDefault() interface{}
	GetShorthand() string
	GetUsage() string
	GetValue() interface{}
	GetValuePointer() interface{}
	IsSetExplicitlyByFlag() bool
	SetValue(value interface{}) (err error)
}
