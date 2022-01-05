package ui

import "github.com/hajimehoshi/ebiten/v2"

// This should handle drawing menus from which commands can be selected

type Menu struct {
	options  []string
	selected int
	actions  []func()
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
func (m *Menu) Select(i int) {
	switch {
	case i < 0:
		m.selected = len(m.options) - 1
	case i > len(m.options)-1:
		m.selected = 0
	default:
		m.selected = i
	}
}

//Selects next menu item
func (m *Menu) SelectNext() {
	m.Select(m.selected + 1)
}

func (m *Menu) SelectPrevious() {
	m.Select(m.selected - 1)
}

// Executes function at selected
func (m *Menu) Enter() {
	if m.actions[m.selected] != nil {
		m.actions[m.selected]() //Call function
	}
}
