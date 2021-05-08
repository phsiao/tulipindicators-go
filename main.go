package main

import (
	"fmt"

	"github.com/phsiao/tulipindicators-go/pkg/indicators"
)

func main() {

	input := [][]float64{
		{82.15, 81.89, 83.03, 83.30, 83.85, 83.90, 83.33, 84.30,
			84.84, 85.00, 85.90, 86.58, 86.98, 88.00, 87.87},
		{81.29, 80.64, 81.31, 82.65, 83.07, 83.11, 82.49, 82.30,
			84.15, 84.11, 84.03, 85.39, 85.76, 87.17, 87.01},
		{81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.82, 83.99,
			84.55, 84.36, 85.53, 86.65, 86.89, 87.77, 87.29},
	}

	output, _, err := indicators.ADD(input[0], input[1])
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v\n", output)
}
