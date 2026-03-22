package main

import (
	"fmt"
	"os"

	"github.com/TheAimHero/dtui/cmd/tui/tabs"
)

func main() {
	err := tabs.NewTui()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
