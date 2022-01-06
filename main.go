package main

import (
	"log"

	"github.com/Kirbstomper/rpgdemo/input"
	"github.com/Kirbstomper/rpgdemo/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var menu_pos = 0

func (g *Game) Update() error {
	input.ReadInput() // Read User input

	return nil
}

func init() {
	ui.CurrentMenu = &ui.MAIN_MENU
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.CurrentMenu.Draw(screen)
	ui.DrawPlayerInformation(screen)
	ui.DrawEnemyInformation(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
