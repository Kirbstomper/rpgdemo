package main

import (
	"log"

	"github.com/Kirbstomper/rpgdemo/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Game struct{}

var menu_pos = 0

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		ui.CurrentMenu.SelectPrevious()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		ui.CurrentMenu.SelectNext()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		ui.CurrentMenu.Enter()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		ui.CurrentMenu.GoToParent()
	}

	return nil
}

var mplusNormalFont font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	ui.CurrentMenu = &ui.MAIN_MENU
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.CurrentMenu.Draw(screen)
	ui.BattleLog.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 1000
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
