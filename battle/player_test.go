package battle

import "testing"

func Test_Attack(t *testing.T) {
	p := player{}
	e := enemy{Health: 10}

	p.Attack(&e)

	if e.Health != 5 {
		t.Fail()
	}
}
