package indicators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
date        input   sma
2005-11-01	81.59
2005-11-02	81.06
2005-11-03	82.87
2005-11-04	83.00
2005-11-07	83.61	82.43
2005-11-08	83.15	82.74
2005-11-09	82.84	83.09
2005-11-10	83.99	83.32
2005-11-11	84.55	83.63
2005-11-14	84.36	83.78
2005-11-15	85.53	84.25
2005-11-16	86.54	84.99
2005-11-17	86.89	85.57
2005-11-18	87.77	86.22
*/

func TestSMA(t *testing.T) {
	input := []float64{81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.82,
		83.99, 84.55, 84.36, 85.53, 86.54, 86.89, 87.77}

	output, err := SMA(input, 5)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%.2f", output[0]), "82.43")
	assert.Equal(t, fmt.Sprintf("%.2f", output[1]), "82.74")
	assert.Equal(t, fmt.Sprintf("%.2f", output[2]), "83.09")
}
