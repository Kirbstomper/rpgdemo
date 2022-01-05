package ui

//This is used for all main menu logic

const ATTACK = "Attack" //0
const DEFEND = "Defend" //1
const ITEMS = "Items"   //2
const RUN = "Run"       //3

var MAIN_MENU = &Menu{[]string{ATTACK, DEFEND, ITEMS, RUN}, 0, []func(){attackAction, defendAction, itemsAction, runAction}}

func attackAction() {
	println("Attack!")
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
