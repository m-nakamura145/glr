package main

import (
	"github.com/mitchellh/cli"
	"os"
	"fmt"
)

var Version string = "0.1.0"

func main() {
	c := cli.NewCLI("glr", Version)
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"sync": func() (cli.Command, error) {
			return &Sync{}, nil
		},
	}

	exitStatus, err := c.Run()

	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}
