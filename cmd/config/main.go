package main

var (
	Version   string
	Commit    string
	Timestamp string
)

func main() {
	rootCommand := GetCommand()
	conf.LoadFromEnvironment()
	conf.ApplyToCobra(rootCommand)
	rootCommand.Execute()
}
