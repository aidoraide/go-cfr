package games

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/aidoraide/cfr/src/algo"
)

type kuhnCard = int32

type KuhnPokerHistory struct {
	p1Card     kuhnCard
	p2Card     kuhnCard
	betHistory string
}

type KuhnPoker struct{}

var P1 = int32(0)
var P2 = int32(1)
var players = []int32{P1, P2}

var Pass = int32('p')
var Bet = int32('b')
var PassIdx = 0
var BetIdx = 1
var ActionSpace = []int32{Pass, Bet}

var KuhnPokerCards = []int32{1, 2, 3}

func dealCards() (int32, int32) {
	d := []int32{1, 2, 3}
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return d[0], d[1]
}

func higherCardRegret(p1Card, p2Card kuhnCard, reward float32) (float32, float32) {
	if p1Card > p2Card {
		return -reward, reward // p1 gets negative regret for winning (regret = -reward)
	}
	return reward, -reward
}

func kuhnRegret(p1Card, p2Card kuhnCard, betHistory string) (float32, float32, error) {
	if len(betHistory) > 3 {
		panic(betHistory)
	}
	r1, r2 := float32(0), float32(0)
	if betHistory == "pp" {
		r1, r2 = higherCardRegret(p1Card, p2Card, 1)
	} else if betHistory == "pbp" {
		r1, r2 = 1, -1
	} else if betHistory == "pbb" {
		r1, r2 = higherCardRegret(p1Card, p2Card, 2)
	} else if betHistory == "bp" {
		r1, r2 = -1, 1
	} else if betHistory == "bb" {
		r1, r2 = higherCardRegret(p1Card, p2Card, 2)
	} else {
		return r1, r2, fmt.Errorf("Non terminal state")
	}
	return r1, r2, nil
}

// Fulfill Game interface requirements

func (kn *KuhnPoker) NActions() int32 {
	return 2
}

func (kn *KuhnPoker) NewGame() algo.History {
	p1Card, p2Card := dealCards()
	return &KuhnPokerHistory{
		p1Card:     p1Card,
		p2Card:     p2Card,
		betHistory: "",
	}
}

func (kn *KuhnPoker) Value(h algo.History) []float32 {
	kh, ok := h.(*KuhnPokerHistory)
	if !ok {
		panic("KuhnPokerGame was given a History instance for another game")
	}
	p1Regret, p2Regret, _ := kuhnRegret(kh.p1Card, kh.p2Card, kh.betHistory)
	value := make([]float32, 2)
	value[0] = -p1Regret
	value[1] = -p2Regret
	return value
}

func (kn *KuhnPoker) PlayerSet() []int32 {
	return players
}

// Fulfill History interface requirements

func (kh *KuhnPokerHistory) TurnToAct() int32 {
	return int32(len(kh.betHistory) % 2)
}

func (kh *KuhnPokerHistory) Infoset() *algo.Infoset {
	return algo.NewInfoset(ActionSpace)
}

func (kh *KuhnPokerHistory) InfosetKey() string {
	if kh.TurnToAct() == P1 {
		return fmt.Sprintf("%d%s", kh.p1Card, kh.betHistory)
	}
	return fmt.Sprintf("%d%s", kh.p2Card, kh.betHistory)
}

func (kh *KuhnPokerHistory) TakeAction(action int32) algo.History {
	return &KuhnPokerHistory{
		p1Card:     kh.p1Card,
		p2Card:     kh.p2Card,
		betHistory: fmt.Sprintf("%s%c", kh.betHistory, action),
	}
}

func (kh *KuhnPokerHistory) IsTerminal() bool {
	_, _, err := kuhnRegret(kh.p1Card, kh.p2Card, kh.betHistory)
	return err == nil // regret is defined only for terminal states, so if getting regret returns no error then we are terminal
}

func (kh *KuhnPokerHistory) String() string {
	return fmt.Sprintf("%d%d%s", kh.p1Card, kh.p2Card, kh.betHistory)
}

// Debugging

/*
From https://en.wikipedia.org/wiki/Kuhn_poker#Optimal_strategy

Optimal strategy is defined as

s1 = strategy(1x)
s1[Pass] = 1-a
s1[Bet] = a

s2 = strategy(2xpb)
s2[Pass] = 2/3 - a
s2[Bet] = 1/3 + a

s3 = strategy(3x)
s3[Pass] = 1-3a
s3[Bet] = 3a

s2b = strategy(x2b)
s2b[Pass] = 2/3
s2b[Bet] = 1/3

s1p = strategy(x1p)
s1p[Pass] = 2/3
s1p[Bet] = 1/3
*/

func PrintKuhnStrategy(model algo.Model) {
	for _, c1 := range KuhnPokerCards {
		for _, c2 := range KuhnPokerCards {
			if c1 == c2 {
				continue
			}
			h := &KuhnPokerHistory{
				p1Card:     c1,
				p2Card:     c2,
				betHistory: "",
			}
			stack := []*KuhnPokerHistory{h}
			for len(stack) > 0 {
				hp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if hp.IsTerminal() {
					continue
				}
				strategy := model.GetStrategy(hp)
				fmt.Printf("%s%s %.3f%s%f\n", strings.Repeat("    ", len(hp.betHistory)), hp, strategy[BetIdx], strings.Repeat("     ", 5-len(hp.betHistory)), strategy[BetIdx])
				for _, action := range ActionSpace {
					hpa := hp.TakeAction(action)
					kuhnHpa, _ := hpa.(*KuhnPokerHistory)
					stack = append(stack, kuhnHpa)
				}
			}
		}
	}
}
