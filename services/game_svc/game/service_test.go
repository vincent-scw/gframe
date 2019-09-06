package game

import (
	"testing"

	c "github.com/vincent-scw/gframe/contracts"
)

func TestJudge(t *testing.T) {
	judgeUtil(c.Shape_Rock, c.Shape_Paper, 1, t)
	judgeUtil(c.Shape_Rock, c.Shape_Scissors, 0, t)
	judgeUtil(c.Shape_Rock, c.Shape_Rock, -1, t)
	judgeUtil(c.Shape_Paper, c.Shape_Scissors, 1, t)
}

func judgeUtil(s1, s2 c.Shape, winner int32, t *testing.T) {
	p1 := c.Playing{Shape: s1}
	p2 := c.Playing{Shape: s2}
	plays := []*c.Playing{&p1, &p2}

	result := judge(plays)
	if result.Winner != winner {
		t.Errorf("judege error %s: %s", s1, s2)
	}
}