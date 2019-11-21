package main

import "github.com/techdecaf/tasks/cmd"

// VERSION - This is converted to the git tag at compile time using the make build command.
var VERSION string

func main() {
	cmd.VERSION = VERSION

	cmd.Execute()
}