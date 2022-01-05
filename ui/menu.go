package ui

import "github.com/hajimehoshi/ebiten/v2"

// This should handle drawing the player menu from which commands can be selected

const ATTACK = "Attack" //0
const DEFEND = "Defend" //1
const ITEMS = "Items"   //2
const RUN = "Run"       //3

var menu_options = make([]string, 4)

func init() {
	menu_options[0] = ATTACK
	menu_options[1] = DEFEND
	menu_options[2] = ITEMS
	menu_options[3] = RUN
}

// Draws the menu
func drawMenu(op []string, x, y int, r *ebiten.Image) {
	spacing := 20
	for i, s := range op {
		drawText(r, s, x, y+(i*spacing))
	}
}

func DrawMainMenu(x, y int, r *ebiten.Image) {
	drawMenu(menu_options, x, y, r)
}
