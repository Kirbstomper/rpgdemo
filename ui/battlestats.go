package ui

import (
	"strconv"

	"github.com/Kirbstomper/rpgdemo/battle"
	"github.com/hajimehoshi/ebiten/v2"
)

//Draws player information to screen
func DrawPlayerInformation(r *ebiten.Image) {

	plyimg := ebiten.NewImage(200, 100)

	drawTextRed(plyimg, "Health: "+strconv.Itoa(battle.Hero.Health), 20, 20)
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(20, 400)
	r.DrawImage(plyimg, op)
}

//Draws enemy information (Health, enemies remaining to screen)
func DrawEnemyInformation(r *ebiten.Image) {
	plyimg := ebiten.NewImage(200, 100)

	for i, e := range battle.Foes {
		drawTextWhite(plyimg, e.Name+" : "+strconv.Itoa(e.Health), 20, 20+(i*20))

	}
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(400, 400)
	r.DrawImage(plyimg, op)
}
