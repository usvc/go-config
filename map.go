package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Map map[string]Config

func (m *Map) ApplyToCobra(command *cobra.Command) {
	flags := command.Flags()
	for rawFlagString, conf := range *m {
		flagString := normalizeName(rawFlagString, separatorFlag)
		conf.ApplyToFlagSet(flagString, flags)
	}
}

func (m *Map) GetFromEnvironment() {
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
			envValue := env.GetStringSlice(envKey)
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
