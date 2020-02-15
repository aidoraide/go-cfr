package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aidoraide/cfr/src/algo"
	"github.com/aidoraide/cfr/src/games"
)

var PrintIter = 1000000

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
	model := algo.NewCFRModel(&games.KuhnPoker{})

	for i := 0; i < nIters; i += PrintIter {
		trainIters := nIters - i
		if trainIters > PrintIter {
			trainIters = PrintIter
		}
		fmt.Printf("[%d, %d]\n", i, i+trainIters)
		model.Train(trainIters)
		games.PrintKuhnStrategy(model)
	}
}
