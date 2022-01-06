package ui

import (
	"github.com/Kirbstomper/rpgdemo/battle"
	"github.com/hajimehoshi/ebiten/v2"
)

//Draws the lines from the log
func DrawBattleInformation(r *ebiten.Image) {

	background := ebiten.NewImage(600, 100)

	drawTextWhite(background, battle.Latestlog, 20, 20)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 300)
	r.DrawImage(background, op)
}
