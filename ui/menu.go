package ui

import "github.com/hajimehoshi/ebiten/v2"

// This should handle drawing the player menu from which commands can be selected

const ATTACK = "Attack" //0
const DEFEND = "Defend" //1
const ITEMS = "Items"   //2
const RUN = "Run"       //3

var MAIN_MENU = &Menu{[]string{ATTACK, DEFEND, ITEMS, RUN}, 0}

type Menu struct {
	options  []string
	selected int
}

// Draw the menu at given position on image
func (m Menu) Draw(x, y int, r *ebiten.Image) {
	spacing := 20
	for i, s := range m.options {
		if m.selected == i {
			drawTextRed(r, s, x, y+(i*spacing))
		} else {
			drawTextWhite(r, s, x, y+(i*spacing))
		}
	}
}

//Updates item selected for a menu. If new selection is greater or less than menu optiosn
//will wrap options around
func (m Menu) Select(i int) {

}
