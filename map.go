package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Map map[string]Config

func (m *Map) ApplyToCobra(command *cobra.Command) {
	flags := command.Flags()
	for name, conf := range *m {
		conf.ApplyToFlagSet(name, flags)
	}
}

func (m *Map) GetFromEnvironment() {
	env := viper.New()
	for key, conf := range *m {
		if conf.GetDefault() != nil {
			fmt.Println("setting default for", key)
			env.SetDefault(key, conf.GetDefault())
		}
	}
	env.AutomaticEnv()
	for key, conf := range *m {
		defaultValue := conf.GetDefault()
		switch conf.(type) {
		case *String:
			envValue := env.GetString(key)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *StringSlice:
			envValue := env.GetStringSlice(key)
			if envValue != nil {
				conf.SetValue(envValue)
			}
		case *Int:
			envValue := env.GetInt(key)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *Uint:
			envValue := env.GetUint(key)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *Float:
			envValue := env.GetFloat64(key)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		case *Bool:
			envValue := env.GetBool(key)
			if envValue != defaultValue {
				conf.SetValue(envValue)
			}
		}
	}
}
