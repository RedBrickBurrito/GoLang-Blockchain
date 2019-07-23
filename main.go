package main

import (
	"os"

	"github.com/RedBrickBurrito/GoLang-Blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
