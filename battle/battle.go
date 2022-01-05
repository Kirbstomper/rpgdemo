package battle

// Class to handle battles and turns

//Since there will only ever be one battle at a time, having it managed here should be fine right?
var (
	Enemies []enemy
)

func init() {
	Enemies = append(Enemies, enemy{Name: "Elf 1", Health: 10}, enemy{Name: "Elf 2", Health: 10})
}
