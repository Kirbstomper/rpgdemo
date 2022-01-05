package manager

import (
	"github.com/Kirbstomper/rpgdemo/battle"
	"github.com/Kirbstomper/rpgdemo/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func BattleLoop() {
	if battle.State == battle.ENEMY_TURN {
		for _, e := range battle.Enemies {
			ui.SetBattlelog("Enemy Attacks!" + e.Name)
		}
		battle.State = battle.BETWEEN_TURN
		battle.Hero.ResetStats()
	}
}

func Handleinput() {
	switch battle.State {
	case battle.PLAYER_TURN:
		{
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
		}
	default:
		{
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				if ui.HasLines() {
					ui.ReadLog()
				} else {
					battle.State = battle.PLAYER_TURN
					ui.SetCurrentlog("Its the heroes turn!")
				}
			}
		}
	}
}
