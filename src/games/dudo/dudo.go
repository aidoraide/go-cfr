package dudo

import (
	"math/rand"

	"github.com/aidoraide/go-cfr/src/algo"
	"github.com/aidoraide/go-cfr/src/util"
)

var P1OptimalUtility = -7.0 / 258.0

type die uint8 // die E [1, 6]

var dieSides = 6

func idx2Die(idx int) die {
	// idx E [0, 5]
	// 5 -> 1
	// n -> n+2, n != 5
	return die((idx+1)%dieSides + 1)
}

func die2Idx(d die) int {
	return (int(d) + dieSides - 2) % dieSides
}

func randomDie() die {
	return idx2Die(rand.Intn(63) % dieSides)
}

type dudo struct {
	nPlayers    int
	nDie        int
	playerSet   []int
	actionSpace []algo.Action
	dudoAction  algo.Action
}

func nActions(nPlayers, nDie int) int {
	return nPlayers*nDie*dieSides + 1
}

func NewDudo(nPlayers, nDie int) algo.Game {
	playerSet := util.Range(0, nPlayers, 1)
	nActions := nActions(nPlayers, nDie)
	actionSpace := make([]algo.Action, nActions)
	for i := range actionSpace {
		actionSpace[i] = algo.Action(i)
	}
	return &dudo{
		nPlayers:    nPlayers,
		nDie:        nDie,
		playerSet:   playerSet,
		actionSpace: actionSpace,
		dudoAction:  algo.Action(nActions - 1),
	}
}

func (d *dudo) NewGame() algo.History {
	return NewDudoHistory(d)
}

func (d *dudo) PlayerSet() []int {
	return d.playerSet
}

func breakdownClaim(claim algo.Action) (die, int) {
	claimInt := int(claim)
	d := idx2Die(claimInt % dieSides)
	count := claimInt/6 + 1
	return d, count
}
