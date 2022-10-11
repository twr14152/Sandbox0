package main

import (
	"fmt"
)

func emulateDotProd(inputs []float64, weights [][]float64, biases []float64) []float64 {
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

	return layers_output
}

func main() {
	inputs := []float64{1, 2, 3, 2.5}
	weights := [][]float64{
		{0.2, 0.8, -0.5, 1},
		{0.5, -0.91, 0.26, -0.5},
		{-0.26, -0.27, 0.17, 0.87},
	}
	biases := []float64{2, 3, 0.5}
	results := emulateDotProd(inputs, weights, biases)
	fmt.Println(results)
}
pi@RaspPi4:~/Coding/Go_folder/nnfs/ch2 $ go run ch2_numpy_emulator.go 
[4.5 1.75 2.675]
pi@RaspPi4:~/Coding/Go_folder/nnfs/ch2 $ 
