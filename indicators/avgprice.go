package indicators

//#include "../tulipindicators/indicators/avgprice.c"
import "C"
import "fmt"

// AVGPRICE function wraps `avgprice' function that provides "Average Price"
//
// Reference: https://tulipindicators.org/avgprice
func AVGPRICE(open, high, low, close []float64) (avgprice []float64, err error) {
	input_length := len(open)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 4)
	all_input_data.Set([][]float64{open, high, low, close})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_avgprice(
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
	avgprice = outputs[0]
	return
}
