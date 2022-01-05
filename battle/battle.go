package battle

// Class to handle battles and turns

//Since there will only ever be one battle at a time, having it managed here should be fine right?
var (
	Hero    *Player
	Enemies []enemy
	State   int
)

const (
	PLAYER_TURN = 1
	ENEMY_TURN  = 2
)

func init() {
	Enemies = append(Enemies, enemy{Name: "Elf 1", Health: 10}, enemy{Name: "Elf 2", Health: 10})
	Hero = &Player{Health: 20, baseDefense: 3, baseAttack: 4}
}

//This loop should continue until Either player health is 0 or no more enemies exist
func TurnLoop() {
	switch State {
	case PLAYER_TURN:
		{
			println("It's the Hero turn!")
			Hero.ResetStats()
		}
	//Enemies act
	case ENEMY_TURN:
		{
			println("Its the enemy turn!")

		}
	}
}
