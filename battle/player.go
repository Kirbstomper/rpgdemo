package battle

//Handles all player related things, such as seting up defense, end of turn actions, etc

type Player struct {
	Health         int
	currentDefense int
	currentAttack  int
	baseAttack     int
	baseDefense    int
}

//Buff defense for player by 2x
func (p *Player) DefendPlayer() {
	p.currentDefense = p.baseDefense * 2
}

func (p *Player) AttackEnemy(e *enemy) {

}

//Reset defense back to regular
func (p *Player) EndOfTurn() {
	p.currentDefense = p.baseDefense
}
