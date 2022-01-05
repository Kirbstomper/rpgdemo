package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//Meant to display a simple text box at the bottom of the screen to display battle status
//I think leveraging the menu here could actually work where a textbox feeds lines until it is at len
// Then sets battle state back to plater

type TextBox struct {
	x, y  int
	lines []string
	cur   string
}

var BattleLog *TextBox

func init() {
	BattleLog = &TextBox{cur: "", x: 20, y: 800}
}

//Draws textbox to the screen
func (t *TextBox) Draw(r *ebiten.Image) {

	background := ebiten.NewImage(1000, 200)
	background.Fill(color.RGBA{0, 0, 255, 255})
	drawTextRed(background, t.cur, 100, 100)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x), float64(t.y))
	r.DrawImage(background, op)

}
func SetCurrentlog(s string) {
	BattleLog.cur = s
}
func SetBattlelog(s string) {
	BattleLog.lines = append(BattleLog.lines, s)
}
func ReadLog() {
	BattleLog.cur = BattleLog.lines[0]
	BattleLog.lines = BattleLog.lines[1:]
}
func HasLines() bool {
	return len(BattleLog.lines) > 0
}
