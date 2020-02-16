package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aidoraide/cfr/src/algo"
	"github.com/aidoraide/cfr/src/games/dudo"
)

var PrintIter = 100000

func getIterationsFromArgs() int {
	nIter := 100000
	if len(os.Args) > 1 {
		arg, err := strconv.ParseInt(os.Args[1], 0, 64)
		nIter = int(arg)
		if err != nil {
			panic(err)
		}
	}
	return nIter
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	nIters := getIterationsFromArgs()
	fmt.Println("Running with nIters = ", nIters)
	model := algo.NewCFRModel(dudo.NewDudo(2, 1))

	for i := 0; i < nIters; i += PrintIter {
		trainIters := nIters - i
		if trainIters > PrintIter {
			trainIters = PrintIter
		}
		utility := model.Train(trainIters)
		p1DifferenceFromOptimalUtility := math.Abs(dudo.P1OptimalUtility - float64(utility[0]))
		percentProgress := 100.0 * float64(i+trainIters) / float64(nIters)
		fmt.Printf("[%6.2f%%] l1loss=%.5f Utility=%v \n", percentProgress, p1DifferenceFromOptimalUtility, utility)
		// games.PrintKuhnStrategy(model)
	}
}
