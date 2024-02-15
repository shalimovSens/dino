package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})
	game.Screen().SetLevel(level)
	game.Start()
}
