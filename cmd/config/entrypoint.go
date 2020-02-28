package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCommand := cobra.Command{
		Use: "config",
		PreRun: func(command *cobra.Command, args []string) {
			conf.GetFromEnvironment()
		},
		Run: func(command *cobra.Command, args []string) {
			for key, conf := range conf {
				fmt.Printf("%s: %v", key, conf.GetValue())
				length := -1
				switch v := conf.GetValue().(type) {
				case []string:
					length = len(v)
				case []uint:
					length = len(v)
				case []int:
					length = len(v)
				}
				if length > -1 {
					fmt.Printf(" (length: %v)", length)
				}
				fmt.Println()
			}
		},
	}
	conf.ApplyToCobra(&rootCommand)
	rootCommand.Execute()
}
