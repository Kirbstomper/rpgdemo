package battle

// The hero of our short example battle

type player struct {
	Health int
}

//Attacks an enemy and deals a set number of damage
func (player) Attack(e *enemy) {
	e.Health = e.Health - 5
}
