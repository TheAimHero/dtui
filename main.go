package main

import "github.com/TheAimHero/dtui/cmd/tui"

func main() {
	err := tui.Init()
	if err != nil {
		panic(err)
	}
}
