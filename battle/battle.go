package battle

// Holds state associated with the battle

// I need to traverse through turns a little more elegantly

// I like the idea of battle state for controls, but then enemy turn needs to be going through enemy list
// ANd working in tandem with UI to advance things along

// Control -> UI -> (reads data from) -> Battle

var (
	Hero *player
	Foes []enemy
)

const ()

type enemy struct {
	Name   string
	Health int
}

func init() {
	Foes = []enemy{{Name: "Goblin 1", Health: 20}, {Name: "Goblin 2", Health: 25}}
	Hero = &player{Health: 100}
}
