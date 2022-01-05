package ui

// This should handle drawing the player menu from which commands can be selected

const ATTACK = "Attack" //0
const DEFEND = "Defend" //1
const ITEMS = "Items"   //2
const RUN = "Run"       //3

var menu_options = make([]string, 4)

func init() {
	menu_options[0] = ATTACK
	menu_options[1] = DEFEND
	menu_options[2] = ITEMS
	menu_options[3] = RUN
}
