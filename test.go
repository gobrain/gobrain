package main

import (
	"fmt"
	"math/rand"

	"github.com/jonhkr/gobrain/neuralnetwork"
)

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	rand.Seed(0)
	// patterns := [][][]float64{
	// 	{{0, 0}, {0}},
	// 	{{0, 1}, {1}},
	// 	{{1, 0}, {1}},
	// 	{{1, 1}, {0}},
	// }

	// patterns := [][][]float64{
	//  {{.0}, {.0 * .0}},
	//  {{.1}, {.1 * .1}},
	//  {{.2}, {.2 * .2}},
	//  {{.3}, {.3 * .3}},
	//  {{.4}, {.4 * .4}},
	// }

	patterns := [][][]float64{
		{{0.01}, {0.01}},
		{{0.02}, {0.01}},
		{{0.03}, {0.02}},
		{{0.04}, {0.03}},
		{{0.05}, {0.05}},
		{{0.06}, {0.08}},
		{{0.07}, {0.13}},
		{{0.08}, {0.21}},
		{{0.09}, {0.34}},
		{{0.10}, {0.55}},
	}

	testPatterns := [][][]float64{
		{{0.06}, {0.08}},
		{{0.07}, {0.13}},
		{{0.08}, {0.21}},
		{{0.09}, {0.34}},
		{{0.10}, {0.55}},
		{{0.11}, {0.89}},
		{{1}, {0}},
	}

	nn := neuralnetwork.New(1, 8, 1, true)

	fmt.Println(nn)

	nn.Train(patterns, 100000, 0.2, 0.05)

	nn.Test(testPatterns)

}