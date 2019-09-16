package util

import (
	"testing"
)

func TestNextRandom(t *testing.T) {
	nxt := NextRandom()
	if len(nxt) != 7 {
		t.Error("Random error")
	}
}