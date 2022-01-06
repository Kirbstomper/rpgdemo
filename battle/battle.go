package battle

// Holds state associated with the battle

// I need to traverse through turns a little more elegantly

// I like the idea of battle state for controls, but then enemy turn needs to be going through enemy list
// ANd working in tandem with UI to advance things along

// Control -> UI -> (reads data from) -> Battle

var (
	Hero      *player
	Foes      []enemy
	foeTurn   int
	battlelog []string
	Latestlog string

	State int
)

const (
	PLAYER_TURN = iota + 1
	ENEMY_TURN  = iota + 1
)

type enemy struct {
	Name   string
	Health int
}

func init() {
	Foes = []enemy{{Name: "Goblin 1", Health: 20}, {Name: "Goblin 2", Health: 25}}
	Hero = &player{Health: 100}
	State = PLAYER_TURN
	foeTurn = 0

}

func BattleLoop() {
	if IsLogEmpty() {
		if State == ENEMY_TURN {
			if foeTurn == len(Foes) {
				foeTurn = 0
				State = PLAYER_TURN
			} else {
				if !IsLogEmpty() {
					ReadLogLine()
				}
				battlelog = append(battlelog, "Player Hit by "+Foes[foeTurn].Name)
				Latestlog = "Player Hit by " + Foes[foeTurn].Name
				Hero.Health = Hero.Health - 5
				foeTurn += 1
			}
		}
	}

}

//Reads a entry in the log and pops it off
func ReadLogLine() {
	if len(battlelog) > 1 {
		battlelog = battlelog[1:]
	}
	if len(battlelog) == 1 {
		battlelog = []string{}
	}
}

//Returns if log is empty
func IsLogEmpty() bool {
	return len(battlelog) == 0
}

func IsPlayersTurn() bool {
	return State == PLAYER_TURN
}

