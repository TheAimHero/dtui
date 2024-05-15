package main

import "github.com/TheAimHero/dtui/cmd/tui/tabs"

func main() {
	err := tui.NewTui()
	if err != nil {
		panic(err)
	}
}
