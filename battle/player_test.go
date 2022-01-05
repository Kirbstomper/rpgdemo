package battle

import "testing"

func Test_DefendPlayer(t *testing.T) {
	p := Player{currentDefense: 3, baseDefense: 3}

	p.DefendPlayer()
	if p.currentDefense != (p.baseDefense * 2) {
		t.Fail()
	}

}

func Test_ResetStats(t *testing.T) {

	p := Player{currentDefense: 99, baseDefense: 3}

	p.ResetStats()
	if p.currentDefense != p.baseDefense {
		t.Fail()
	}
}
