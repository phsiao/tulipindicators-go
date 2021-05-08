package indicators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	input1 := []float64{81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.82,
		83.99, 84.55, 84.36, 85.53, 86.54, 86.89, 87.77}
	input2 := []float64{81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.82,
		83.99, 84.55, 84.36, 85.53, 86.54, 86.89, 87.77}

	output, err := ADD(input1, input2)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%.2f", output[0]), "163.18")
	assert.Equal(t, fmt.Sprintf("%.2f", output[1]), "162.12")
	assert.Equal(t, fmt.Sprintf("%.2f", output[2]), "165.74")
}
