package main

import "fmt"

func main() {
	inputs := []float64{1, 2, 3, 2.5}
	weights := [][]float64{
		{0.2, 0.8, -0.5, 1},
		{0.5, -0.91, 0.26, -0.5},
		{-0.26, -0.27, 0.17, 0.87},
	}
	biases := []float64{2, 3, 0.5}
	layersOutput := make([]float64, 3)

	for i := range weights {
		neuronWeights := weights[i]
		neuronBias := biases[i]
		var neuronOutput float64 = 0
		for j := range inputs {
			neuronOutput = inputs[j] * neuronWeights[j]
		}
		neuronOutput += neuronBias
		layersOutput[i] = neuronOutput
	}

	fmt.Println(layersOutput)
}
/*
results:
[4.5 1.75 2.675]

Program exited.
*/
