package indicators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	input1 := []float64{81.59, 81.06, 82.87}
	input2 := []float64{81.59, 81.06, 82.87}

	output, err := ADD(input1, input2)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%.2f", output[0]), "163.18")
	assert.Equal(t, fmt.Sprintf("%.2f", output[1]), "162.12")
	assert.Equal(t, fmt.Sprintf("%.2f", output[2]), "165.74")
}
