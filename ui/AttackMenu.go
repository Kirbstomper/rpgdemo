package ui

import (
	"github.com/Kirbstomper/rpgdemo/battle"
)

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

// Updates attack menu with information from battle about enemies
func updateAttackMenu() {
	nOpts := []string{}
	nActions := []func(){}
	for i, _ := range battle.Foes {
		var index = i
		foe := &battle.Foes[index]
		nOpts = append(nOpts, foe.Name)
		nActions = append(nActions, func() {
			battle.Hero.Attack(foe)
			println(foe.Name)
		})
	}
	ATTACK_MENU.options = nOpts
	ATTACK_MENU.actions = nActions
}
