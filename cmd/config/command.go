package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type logger struct {
	writer io.Writer
}

func (l logger) Printf(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, format, args...)
}

func (l logger) Println(args ...interface{}) {
	fmt.Fprintln(l.writer, args...)
}

func GetCommand() *cobra.Command {
	command := cobra.Command{
		Use:     "config",
		Version: fmt.Sprintf("%s-%s %s", Version, Commit, Timestamp),
		Run: func(command *cobra.Command, args []string) {
			log := logger{command.OutOrStdout()}
			for key, conf := range conf {
				log.Printf("%s: %v", key, conf.GetValue())
				length := -1
				switch v := conf.GetValue().(type) {
				case []string:
					length = len(v)
				case []int:
					length = len(v)
				}
				if length > -1 {
					log.Printf(" (length: %v)", length)
				}
				log.Println()
			}
		},
	}
	return &command
}
