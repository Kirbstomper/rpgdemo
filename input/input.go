package input

import (
	"github.com/Kirbstomper/rpgdemo/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func ReadInput() {

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
