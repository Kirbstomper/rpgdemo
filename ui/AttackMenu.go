package ui

// This will contain the logic related to the attack menu

var ATTACK_MENU Menu

func init() {
	ATTACK_MENU = Menu{
		parent:   &MAIN_MENU,
		options:  []string{"Target 1", "Target 2"},
		actions:  []func(){func() { println("Hit Target 1") }, func() { println("Hit Target 2") }},
		selected: 0,
	}
}
