package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/usvc/config"
)

var conf = config.Map{
	"str1": &config.String{
		Default: "default",
	},
	"str2": &config.String{
		Default: "default",
	},
	"uint1": &config.Uint{
		Default: 1234567,
	},
	"uint2": &config.Uint{
		Default: 1234567,
	},
	"int1": &config.Int{
		Default: 1234567,
	},
	"int2": &config.Int{
		Default: 1234567,
	},
	"bool": &config.Bool{},
}

func main() {
	rootCommand := cobra.Command{
		Use: "config",
		PreRun: func(command *cobra.Command, args []string) {
			conf.GetFromEnvironment()
		},
		Run: func(command *cobra.Command, args []string) {
			fmt.Println(conf["str1"].GetValue())
			fmt.Println(conf["uint1"].GetValue())
			fmt.Println(conf["bool"].GetValue())
		},
	}
	conf.ApplyToCobra(&rootCommand)
	rootCommand.Execute()
}
