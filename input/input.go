package input

//Handles reading input from the user to interact with the game

import (
	"github.com/Kirbstomper/rpgdemo/battle"
	"github.com/Kirbstomper/rpgdemo/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//Reads user input to determine what should be controled based on current state
func ReadInput() {
	if battle.IsPlayersTurn() {
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

	if !battle.IsPlayersTurn() {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if !battle.IsLogEmpty() {
				println(battle.ReadLogLine())
			}
		}
	}

}
