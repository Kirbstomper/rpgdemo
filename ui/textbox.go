package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//Meant to display a simple text box at the bottom of the screen to display battle status
//I think leveraging the menu here could actually work where a textbox feeds lines until it is at len
// Then sets battle state back to plater

type TextBox struct {
	x, y, current int
	lines         []string
}

var BattleLog *TextBox

func init() {
	BattleLog = &TextBox{lines: []string{"BAttleLog!"}, current: 0, x: 20, y: 800}
}

//Draws textbox to the screen
func (t *TextBox) Draw(r *ebiten.Image) {

	background := ebiten.NewImage(1000, 200)
	background.Fill(color.RGBA{0, 0, 255, 255})
	drawTextRed(background, t.lines[t.current], 100, 100)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x), float64(t.y))
	r.DrawImage(background, op)

}
