package main

import (
	"github.com/DevolvingSpud/rogue-core/internal/entities"
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()

	mapLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})
	mapLevel.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
	game.Screen().SetLevel(mapLevel)

	player := entities.Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		Level:  mapLevel,
	}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	mapLevel.AddEntity(&player)

	game.Start()
}
