package ui

import "github.com/hajimehoshi/ebiten/v2"

// This should handle drawing the player menu from which commands can be selected

const ATTACK = "Attack" //0
const DEFEND = "Defend" //1
const ITEMS = "Items"   //2
const RUN = "Run"       //3

var menu_options = make([]string, 4)

var MAIN_MENU = &Menu{menu_options, 0}

type Menu struct {
	options  []string
	selected int
}

func init() {
	menu_options[0] = ATTACK
	menu_options[1] = DEFEND
	menu_options[2] = ITEMS
	menu_options[3] = RUN
}

// Draws the menu
func (m Menu) DrawMenu(x, y int, r *ebiten.Image) {
	spacing := 20
	for i, s := range m.options {
		if m.selected == i {
			drawTextRed(r, s, x, y+(i*spacing))
		} else {
			drawTextWhite(r, s, x, y+(i*spacing))
		}
	}
}
