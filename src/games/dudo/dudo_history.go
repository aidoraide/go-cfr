package dudo

import (
	"bytes"
	"fmt"

	"github.com/aidoraide/cfr/src/algo"
)

type dudoHistory struct {
	game          *dudo
	playerDice    [][]die
	actionHistory []algo.Action
}

func nRandomDice(n int) []die {
	dice := make([]die, n)
	for i := range dice {
		dice[i] = randomDie()
	}
	return dice
}

func NewDudoHistory(d *dudo) *dudoHistory {
	playerDice := make([][]die, d.nPlayers)
	for i := range playerDice {
		rolled := nRandomDice(d.nDie)
		playerDice[i] = rolled
	}
	return &dudoHistory{
		game:          d,
		playerDice:    playerDice,
		actionHistory: make([]algo.Action, 0),
	}
}

func (dh *dudoHistory) TurnToAct() int {
	return len(dh.actionHistory) % dh.game.nPlayers
}

func (dh *dudoHistory) Infoset() *algo.Infoset {
	if len(dh.actionHistory) == 0 {
		return algo.NewInfoset(dh.game.actionSpace[:len(dh.game.actionSpace)-1]) // Remove DUDO action (last action)
	} else {
		lastAction := dh.actionHistory[len(dh.actionHistory)-1]
		return algo.NewInfoset(dh.game.actionSpace[lastAction+1:])
	}
}

func writeDiceString(buf *bytes.Buffer, dice []die) {
	for _, d := range dice {
		fmt.Fprintf(buf, "%d", d)
	}
}

func writeActionHistory(buf *bytes.Buffer, actionHistory []algo.Action, dudoAction algo.Action) {
	for _, a := range actionHistory {
		if a == dudoAction {
			buf.WriteString("DUDO,")
		} else {
			fmt.Fprintf(buf, "%dx%d,", int(a)/dieSides+1, int(a)%dieSides+1)
		}
	}
	buf.Truncate(buf.Len() - 1) // Remove trailing ","
}

func (dh *dudoHistory) InfosetKey() string {
	var buf bytes.Buffer
	writeDiceString(&buf, dh.playerDice[dh.TurnToAct()])
	buf.WriteString(":")
	writeActionHistory(&buf, dh.actionHistory, dh.game.dudoAction)
	return buf.String()
}

func (dh *dudoHistory) TakeAction(action algo.Action) algo.History {
	actionHistory := make([]algo.Action, len(dh.actionHistory)+1)
	for i, a := range dh.actionHistory {
		actionHistory[i] = a
	}
	actionHistory[len(dh.actionHistory)] = action
	return &dudoHistory{
		game:          dh.game,
		playerDice:    dh.playerDice,
		actionHistory: actionHistory,
	}
}

func (dh *dudoHistory) IsTerminal() bool {
	return len(dh.actionHistory) > 0 && dh.actionHistory[len(dh.actionHistory)-1] == dh.game.dudoAction
}

func (dh *dudoHistory) Value() []float32 {
	if len(dh.actionHistory) == 0 || dh.actionHistory[len(dh.actionHistory)-1] != dh.game.dudoAction {
		panic(fmt.Errorf("Tried to get value for non terminal state: %s", dh))
	}

	if len(dh.playerDice) > 2 {
		panic("Not yet supported for playing with more than 2 players")
	}
	for _, dice := range dh.playerDice {
		if len(dice) > 1 {
			panic("Not yet supported for playing with more than 1 die")
		}
	}

	lastClaim := dh.actionHistory[len(dh.actionHistory)-2]
	lastClaimant := (len(dh.actionHistory) - 2) % len(dh.playerDice)
	dudoCaller := (len(dh.actionHistory) - 1) % len(dh.playerDice)
	claimDie, claimCount := breakdownClaim(lastClaim)

	actualCount := 0
	for _, dice := range dh.playerDice {
		for _, d := range dice {
			if d == die(1) || d == claimDie {
				actualCount += 1
			}
		}
	}

	value := make([]float32, 2)
	if claimCount == actualCount {
		value[dudoCaller] = -1
		value[lastClaimant] = 1
	} else if actualCount < claimCount {
		value[dudoCaller] = 1
		value[lastClaimant] = -1
	} else {
		value[dudoCaller] = -1
		value[lastClaimant] = 1
	}
	return value
}

func (dh *dudoHistory) String() string {
	var buf bytes.Buffer
	for _, dice := range dh.playerDice {
		writeDiceString(&buf, dice)
		buf.WriteString(":")
	}
	writeActionHistory(&buf, dh.actionHistory, dh.game.dudoAction)
	return buf.String()
}
