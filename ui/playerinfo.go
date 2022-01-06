package ui

import (
	"strconv"

	"github.com/Kirbstomper/rpgdemo/battle"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawPlayerInformation(r *ebiten.Image) {

	plyimg := ebiten.NewImage(200, 100)

	drawTextRed(plyimg, "Health: "+strconv.Itoa(battle.Hero.Health), 20, 20)
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(20, 200)
	r.DrawImage(plyimg, op)
}
