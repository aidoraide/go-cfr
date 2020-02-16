package dudo

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	dice := []int{2, 3, 4, 5, 6, 1}
	for i, d := range dice {
		t.Run(fmt.Sprintf("idx %d -> die %d", i, d), func(t *testing.T) {
			if idx2Die(i) != die(d) {
				t.Errorf("idx2Die(%d) != die(%d)", i, d)
			}
		})
	}

	for i, d := range dice {
		t.Run(fmt.Sprintf("die %d -> idx %d", d, i), func(t *testing.T) {
			if die2Idx(die(d)) != i {
				t.Errorf("die2Idx(%d) != idx(%d)", d, i)
			}
		})
	}

}
