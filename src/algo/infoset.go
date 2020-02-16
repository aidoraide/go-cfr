package algo

import (
	"github.com/aidoraide/cfr/src/util"
)

type Infoset struct {
	regretSums   []float64
	strategySums []float64
	actionSpace  []Action
}

var epsilon = float64(0.00001)

func NewInfoset(actionSpace []Action) *Infoset {
	return &Infoset{
		regretSums:   util.NVals(0, len(actionSpace)),
		strategySums: util.NVals(0, len(actionSpace)),
		actionSpace:  actionSpace,
	}
}

func (is *Infoset) GetStrategy() []float64 {
	strategy := util.Max(is.regretSums, 0)
	normalizer := util.Sum(strategy)
	if normalizer == 0 {
		return util.NVals(1/float64(len(strategy)), len(strategy))
	}

	util.DivideBy(strategy, normalizer)

	// Remove small values and renormalize
	util.Clamp(strategy, func(x float64) bool { return x < epsilon }, 0)
	util.DivideBy(strategy, util.Sum(strategy))
	return strategy
}

func (is *Infoset) GetAverageStrategy() []float64 {
	normalizer := util.Sum(is.strategySums)
	var strategy []float64
	if normalizer == 0 {
		strategy = util.NVals(1/float64(len(is.strategySums)), len(is.strategySums))
	} else {
		strategy = util.Copy(is.strategySums)
		util.DivideBy(strategy, normalizer)
	}

	// Remove small values and renormalize
	util.Clamp(strategy, func(x float64) bool { return x < epsilon }, 0)
	util.DivideBy(strategy, util.Sum(strategy))
	return strategy
}

func (is *Infoset) GetActionSet() []Action {
	return is.actionSpace
}

func (is *Infoset) Accumulate(regret []float64, strategy []float64) {
	util.AddVectorTo(regret, is.regretSums)
	util.AddVectorTo(strategy, is.strategySums)
}

func (is *Infoset) ResetAverageStrategy() {
	is.strategySums = util.NVals(0, len(is.actionSpace))
}
