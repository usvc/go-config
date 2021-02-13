package main

var (
	Version   string
	Commit    string
	Timestamp string
)

var conf = NewConfiguration()

func main() {
	rootCommand := GetCommand()
	conf.ApplyToCobra(rootCommand)
	rootCommand.Execute()
}
