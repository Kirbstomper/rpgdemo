package ui

//This is used for all main menu logic

const ATTACK = "Attack" //0
const DEFEND = "Defend" //1
const ITEMS = "Items"   //2
const RUN = "Run"       //3

var MAIN_MENU Menu

func init() {
	MAIN_MENU = Menu{
		parent:   nil,
		options:  []string{ATTACK, DEFEND, ITEMS, RUN},
		selected: 0,
		actions:  []func(){attackAction, defendAction, itemsAction, runAction},

		x: 20,
		y: 20,
	}
}

func attackAction() {
	updateAttackOptions()
	CurrentMenu = &ATTACK_MENU
}
func defendAction() {
	println("Defend!")
}
func itemsAction() {
	println("Items!")
}
func runAction() {
	println("Run!")
}
