package ui

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// This will handle actually drawing text to a screen to abstract it away

var mplusNormalFont font.Face

//Initialize font used
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

// Draws text on given image
func drawTextWhite(r *ebiten.Image, t string, x, y int) {
	text.Draw(r, t, mplusNormalFont, x, y, color.White)
}

func drawTextRed(r *ebiten.Image, t string, x, y int) {
	text.Draw(r, t, mplusNormalFont, x, y, color.RGBA{255, 0, 0, 255})

}
