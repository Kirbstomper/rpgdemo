package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// This should handle drawing menus from which commands can be selected

var CurrentMenu *Menu

type Menu struct {
	parent   *Menu
	options  []string
	selected int
	actions  []func()
	x, y     int
}

// Draw the menu at given position on image
func (m *Menu) Draw(r *ebiten.Image) {
	var basex, basey int

	if m.parent != nil {
		basex, basey = m.parent.x, m.parent.y
		m.parent.Draw(r)
	}
	basex += 20
	basey += 20
	spacing := 20
	background := ebiten.NewImage(150, 100)
	background.Fill(color.RGBA{0, 0, 255, 255})
	for i, s := range m.options {
		if m.selected == i {
			drawTextRed(background, s, basex, basey+(i*spacing))
		} else {
			drawTextWhite(background, s, basex, basey+(i*spacing))
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(basex), float64(basey))
	r.DrawImage(background, op)
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

// Goes to the parent menu item if non-nil
func (m *Menu) GoToParent() {
	if m.parent != nil {
		CurrentMenu = m.parent
	}
}
