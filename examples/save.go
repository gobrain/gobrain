package examples

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/goml/gobrain"
	"github.com/goml/gobrain/persist"
)

func Save(filename string) {
	// set the random seed to 0
	rand.Seed(0)

	// create the XOR representation patter to train the network
	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	// instantiate the Feed Forward
	ff := &gobrain.FeedForward{}

	// initialize the Neural Network;
	// the networks structure will contain:
	// 2 inputs, 2 hidden nodes and 1 output.
	ff.Init(2, 2, 1)

	// train the network using the XOR patterns
	// the training will run for 1000 epochs
	// the learning rate is set to 0.6 and the momentum factor to 0.4
	// use true in the last parameter to receive reports about the learning error
	ff.Train(patterns, 1000, 0.6, 0.4, false)

	// saves neural network to file
	err := persist.Save(filename, ff)
	if err != nil {
		log.Println("impossible to save network on file: ", err.Error())
	}

	// sends inputs to the neural network
	inputs := []float64{1, 1}

	// saves the result
	result := ff.Update(inputs)

	// prints the result
	fmt.Println(result)
}
