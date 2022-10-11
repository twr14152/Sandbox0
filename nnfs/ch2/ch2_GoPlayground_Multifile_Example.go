//This example uses multifile approach to emulate numpy example in the book
//https://go.dev/play/p/TdmXQvaD2iN


package main

import (
	"fmt"
	"play.ground/foo"
)

func main() {
	inputs := []float64{1, 2, 3, 2.5}
	weights := [][]float64{
		{0.2, 0.8, -0.5, 1},
		{0.5, -0.91, 0.26, -0.5},
		{-0.26, -0.27, 0.17, 0.87},
	}
	biases := []float64{2, 3, 0.5}
	results := foo.DotProd_Fn(inputs, weights, biases)
	fmt.Println(results)
}
-- go.mod --
module play.ground
-- foo/foo.go --
package foo

func DotProd_Fn(inputs []float64, weights [][]float64, biases []float64) []float64 {
	layers_output := make([]float64, 3)

	for i := range weights {
		neuron_weights := weights[i]
		neuron_bias := biases[i]
		var neuron_output float64 = 0
		for j := range inputs {
			neuron_output = inputs[j] * neuron_weights[j]

		}
		neuron_output += neuron_bias
		layers_output[i] = neuron_output
	}

	return (layers_output)
}
/* Results
[4.5 1.75 2.675]

Program exited.
*/
