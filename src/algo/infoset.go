package algo

import (
	"github.com/aidoraide/cfr/src/util"
)

type Infoset struct {
	regretSums   []float32
	strategySums []float32
	actionSpace  []Action
}

var epsilon = float32(0.0001)

func NewInfoset(actionSpace []Action) *Infoset {
	return &Infoset{
		regretSums:   util.NVals(0, len(actionSpace)),
		strategySums: util.NVals(0, len(actionSpace)),
		actionSpace:  actionSpace,
	}
}

func (is *Infoset) GetStrategy() []float32 {
	strategy := util.Max(is.regretSums, 0)
	normalizer := util.Sum(strategy)
	if normalizer == 0 {
		return util.NVals(1/float32(len(strategy)), len(strategy))
	}

	util.DivideBy(strategy, normalizer)

	// Remove small values and renormalize
	util.Clamp(strategy, func(x float32) bool { return x < epsilon }, 0)
	util.DivideBy(strategy, util.Sum(strategy))
	return strategy
}

func (is *Infoset) GetAverageStrategy() []float32 {
	normalizer := util.Sum(is.strategySums)
	var strategy []float32
	if normalizer == 0 {
		strategy = util.NVals(1/float32(len(is.strategySums)), len(is.strategySums))
	} else {
		strategy = util.Copy(is.strategySums)
		util.DivideBy(strategy, normalizer)
	}

	// Remove small values and renormalize
	util.Clamp(strategy, func(x float32) bool { return x < epsilon }, 0)
	util.DivideBy(strategy, util.Sum(strategy))
	return strategy
}

func (is *Infoset) GetActionSet() []Action {
	return is.actionSpace
}

func (is *Infoset) Accumulate(regret []float32, strategy []float32) {
	util.AddVectorTo(regret, is.regretSums)
	util.AddVectorTo(strategy, is.strategySums)
}
