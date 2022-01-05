package ui

import "testing"

var action_string = ""
var m = Menu{parent: nil, options: []string{"TEST", "TEST2", "TEST3", "TEST4"}, actions: []func(){func() { action_string = "HIT" }, nil, nil, nil}}

func Test_SelectBasic(t *testing.T) {
	m.selected = 0

	m.Select(1)
	if m.selected != 1 {
		t.Fail()
	}
}

func Test_SelectOver(t *testing.T) {
	//Select something over length of options
	m.selected = len(m.options)
	m.Select(len(m.options) + 1)
	if m.selected != 0 {
		t.Fail()
	}
}

func Test_SelectUnder(t *testing.T) {
	//Select something under length of options
	m.Select(-1)
	if m.selected != len(m.options)-1 {
		t.Fail()
	}
}

func Test_SelectNext(t *testing.T) {
	m.selected = 0
	m.SelectNext()

	if m.selected != 1 {
		t.Fail()
	}
}

func Test_SelectPrevious(t *testing.T) {
	m.selected = 1
	m.SelectPrevious()
	if m.selected != 0 {
		t.Fail()
	}
}

func Test_EnterNonNilAction(t *testing.T) {
	m.selected = 0
	action_string = ""
	m.Enter()
	if action_string != "HIT" {
		t.Fail()
	}
}

func Test_EnterNilAction(t *testing.T) {
	m.selected = 1
	action_string = ""
	m.Enter()
	if action_string != "" {
		t.Fail()
	}
}

func Test_GotoParent(t *testing.T) {
	CurrentMenu = nil
	m.GoToParent()
	if CurrentMenu != nil {
		t.Fail()
	}
	menuWithParent := Menu{parent: &m}
	menuWithParent.GoToParent()
	if CurrentMenu != &m {
		t.Fail()
	}

}
