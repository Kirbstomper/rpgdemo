package ui

import "testing"

var m = Menu{options: []string{"TEST", "TEST2", "TEST3", "TEST4"}}

func Test_SelectBasic(t *testing.T) {
	m.selected = 0

	m.Select(1)
	if m.selected != 1 {
		t.Fail()
	}
}

func Test_SelectOver(t *testing.T) {
	//Select something over length of options
	m.Select(len(m.options) + 1)
	if m.selected != 0 {
		t.Fail()
	}
}

func Test_SelectUnder(t *testing.T) {
	//Select something over length of options
	m.Select(-1)
	if m.selected != len(m.options)-1 {
		t.Fail()
	}
}
