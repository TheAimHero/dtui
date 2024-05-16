package main

import (
	"fmt"

	"github.com/TheAimHero/dtui/cmd/tui/tabs"
)

func main() {
	err := tabs.NewTui()
	if err != nil {
		fmt.Println(err.Error())
	}
}
