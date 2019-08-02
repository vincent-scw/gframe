package simulator

import (
	"strings"
	"testing"
)

func TestInjector(t *testing.T) {

}

func TestGetRandPlayerName(t *testing.T) {
	name := getRandPlayerName()
	slices := strings.Split(name, " ")
	if len(slices) != 2 {
		t.Error("Wrong name formatting")
	}
}
