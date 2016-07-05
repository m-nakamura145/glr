package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

var Version string = "0.1.0"

func main() {
	c := cli.NewCLI("glr", Version)
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"sync": func() (cli.Command, error) {
			return &Sync{}, nil
		},
		"status start": func() (cli.Command, error) {
			return &StatusStart{}, nil
		},
		"status stop": func() (cli.Command, error) {
			return &StatusStop{}, nil
		},
	}

	exitStatus, err := c.Run()

	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}
