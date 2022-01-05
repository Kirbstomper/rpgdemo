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
		ui.MAIN_MENU.SelectPrevious()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		ui.MAIN_MENU.SelectNext()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		ui.MAIN_MENU.Enter()
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
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.MAIN_MENU.Draw(10, 20, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
