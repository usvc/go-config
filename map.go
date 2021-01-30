package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Map stores a map of configurations identifed by a string key
type Map map[string]Config

// ApplyToCobra applies the configuration stored in the instance
// of Map to a cobra.Command instance as flags (cannot be accessed
// by child commands)
func (m *Map) ApplyToCobra(command *cobra.Command) {
	m.ApplyToFlagSet(command.Flags())
}

// ApplyToCobraPersistent applies the configuration stored in the instance
// of Map to a cobra.Command instance as persistent flags (can be
// accessed by child commands)
func (m *Map) ApplyToCobraPersistent(command *cobra.Command) {
	m.ApplyToFlagSet(command.PersistentFlags())
}

// ApplyToFlagSet applies the configuration stored in the instance
// of Map to a provided set of flags
func (m *Map) ApplyToFlagSet(flags *pflag.FlagSet) {
	for rawFlagString, conf := range *m {
		flagString := normalizeName(rawFlagString, separatorFlag)
		conf.ApplyToFlagSet(flagString, flags)
	}
}

// LoadFromEnvironment loads the configuration from pre-defined
// environment variables
func (m *Map) LoadFromEnvironment() {
	env := viper.New()
	for rawEnvKey, conf := range *m {
		envKey := normalizeName(rawEnvKey, separatorEnv)
		if conf.GetDefault() != nil {
			env.SetDefault(envKey, conf.GetDefault())
		}
	}
	env.AutomaticEnv()
	for rawEnvKey, conf := range *m {
		envKey := normalizeName(rawEnvKey, separatorEnv)
		defaultValue := conf.GetDefault()
		switch conf.(type) {
		case *String:
			envValue := env.GetString(envKey)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *StringSlice:
			// this is a hack to work around a viper/pflags integration
			// issue, see the following isues for details:
			// - https://github.com/spf13/viper/issues/200
			// -  https://github.com/spf13/viper/issues/380
			env.Set(envKey, strings.ReplaceAll(env.GetString(envKey), ",", " "))
			envValue := env.GetStringSlice(envKey)
			fmt.Println(envValue)
			if envValue != nil && !areEqualStringSlice(envValue, defaultValue.([]string)) {
				conf.SetValue(envValue)
			}
		case *Int:
			envValue := env.GetInt(envKey)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *Uint:
			envValue := env.GetUint(envKey)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *Float:
			envValue := env.GetFloat64(envKey)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *Bool:
			envValue := env.GetBool(envKey)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		}
	}
}

// Get retrieves the value of the configuration identified
// by the key :id as an `interface{}` type
func (m Map) Get(id string) interface{} {
	assertIDExists(m, id)
	return m[id].GetValue()
}

// GetBool retrieves the value of the configuration identified
// by the key :id as a boolean type
func (m Map) GetBool(id string) bool {
	return m.Get(id).(bool)
}

// GetFloat retrieves the value of the configuration identified
// by the key :id as a floating point type
func (m Map) GetFloat(id string) float64 {
	return m.Get(id).(float64)
}

// GetInt retrieves the value of the configuration identified
// by the key :id as an integer type
func (m Map) GetInt(id string) int {
	return m.Get(id).(int)
}

// GetIntSlice retrieves the value of the configuration identified
// by the key :id as an integer slice type
func (m Map) GetIntSlice(id string) []int {
	return m.Get(id).([]int)
}

// GetString retrieves the value of the configuration identified
// by the key :id as a string type
func (m Map) GetString(id string) string {
	return m.Get(id).(string)
}

// GetStringSlice retrieves the value of the configuration identified
// by the key :id as a string slice type
func (m Map) GetStringSlice(id string) []string {
	return m.Get(id).([]string)
}

// GetUint retrieves the value of the configuration identified
// by the key :id as an unsigned integer type
func (m Map) GetUint(id string) uint {
	return m.Get(id).(uint)
}

// GetUintSlice retrieves the value of the configuration identified
// by the key :id as an unsigned intger slice type
func (m Map) GetUintSlice(id string) []uint {
	return m.Get(id).([]uint)
}
