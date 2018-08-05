package main

import (
	"fmt"
	"os"

	"github.com/jpk/cmd"
)

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	must(cmd.RootCmd.Execute())
}
