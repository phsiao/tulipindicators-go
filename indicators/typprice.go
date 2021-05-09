package indicators

//#include "../tulipindicators/indicators/typprice.c"
import "C"
import "fmt"

// TYPPRICE function wraps `typprice' function that provides "Typical Price"
//
// Reference: https://tulipindicators.org/typprice
func TYPPRICE(high, low, close []float64) (typprice []float64, err error) {
	input_length := len(high)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 3)
	all_input_data.Set([][]float64{high, low, close})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_typprice(
		(C.int)(input_length),
		(**C.double)(all_input_data.buffer),
		(*C.double)(&options[0]),
		(**C.double)(all_output_data.buffer),
	)

	if err != nil {
		return
	}
	if ret != C.TI_OKAY {
		err = fmt.Errorf("ret = %d", ret)
		return
	}
	outputs := all_output_data.Get()
	typprice = outputs[0]
	return
}
