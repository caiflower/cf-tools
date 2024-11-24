package main

import "github.com/caiflower/cf-tools/command"

func main() {
	cfCommand := command.NewCfCommand()
	cfCommand.Execute()
}
