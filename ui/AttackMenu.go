package ui

import "github.com/Kirbstomper/rpgdemo/battle"

// This will contain the logic related to the attack menu

var ATTACK_MENU Menu

func init() {
	ATTACK_MENU = Menu{
		parent:   &MAIN_MENU,
		options:  []string{""},
		actions:  []func(){func() { println("Hit Target 1") }, func() { println("Hit Target 2") }},
		selected: 0,
	}
}

func updateAttackOptions() {
	var newOptions = make([]string, len(battle.Enemies))
	var newActions = make([]func(), len(battle.Enemies))
	for i, e := range battle.Enemies {
		newOptions[i] = e.Name
		newActions[i] = func() { println("Attacking ", e.Name) } //Lol is this even safe?
	}
	ATTACK_MENU.actions = newActions
	ATTACK_MENU.options = newOptions
}
