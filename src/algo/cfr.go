package algo

import (
	"fmt"

	"github.com/aidoraide/cfr/src/util"
)

type Model interface {
	Train(nIter int)
	GetStrategy(h History) []float32
}

type cfrModel struct {
	game     Game
	cfr_iter func(h History, player int32, pi []float32) float32
	h2I      func(h History) *Infoset
}

func cfr(h History, player int32, pi []float32, h2I func(h History) *Infoset, game Game) float32 {
	if h.IsTerminal() {
		value := game.Value(h)
		return value[player]
	}

	infoset := h2I(h)
	actionSet := infoset.GetActionSet()
	value := float32(0)
	actionValues := util.NVals(0, len(actionSet))
	strategy := infoset.GetStrategy()
	for actionIdx, a := range actionSet {
		ha := h.TakeAction(a)
		// fmt.Println(pi)
		pia := util.Copy(pi)
		pia[h.TurnToAct()] *= strategy[actionIdx]
		actionValues[actionIdx] = cfr(ha, player, pia, h2I, game)
		value += strategy[actionIdx] * actionValues[actionIdx]
	}

	if h.TurnToAct() == player {
		util.AddTo(-value, actionValues)
		util.MultBy(pi[1-player], actionValues)
		regret := actionValues
		// actionValues now holds regret values

		util.MultBy(pi[player], strategy)
		infoset.Accumulate(regret, strategy)
	}

	return value
}

func gameWrapCoreFunctions(game Game) (func(h History, player int32, pi []float32) float32, func(h History) *Infoset) {
	h2IMap := map[string]*Infoset{}
	h2I := func(h History) *Infoset {
		key := h.InfosetKey()
		infoset, ok := h2IMap[key]
		if !ok {
			infoset = h.Infoset()
			h2IMap[key] = infoset
		}
		return infoset
	}

	cfr_clean := func(h History, player int32, pi []float32) float32 {
		return cfr(h, player, pi, h2I, game)
	}

	return cfr_clean, h2I
}

func NewCFRModel(game Game) Model {
	cfr_iter, h2I := gameWrapCoreFunctions(game)
	return &cfrModel{
		game:     game,
		cfr_iter: cfr_iter,
		h2I:      h2I,
	}
}

func (m *cfrModel) Train(nIter int) {
	utility := float32(0.0)
	for i := 0; i < nIter; i++ {
		for _, player := range m.game.PlayerSet() {
			h := m.game.NewGame()
			actionSet := m.h2I(h).GetActionSet()
			utility += m.cfr_iter(h, player, util.NVals(1.0, len(actionSet)))
		}
		if i%10000 == 0 {
			fmt.Printf("t: %5.1f%s\r", 100*float32(i)/float32(nIter), "%")
		}
	}
	fmt.Println("                                                 ")
}

func (m *cfrModel) GetStrategy(h History) []float32 {
	return m.h2I(h).GetAverageStrategy()
}
