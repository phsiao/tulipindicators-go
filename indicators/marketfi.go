package indicators

//#include "../tulipindicators/indicators/marketfi.c"
import "C"
import "fmt"

// MARKETFI function wraps `marketfi' function that provides "Market Facilitation Index"
//
// Reference: https://tulipindicators.org/marketfi
func MARKETFI(high, low, volume []float64) (marketfi []float64, err error) {
	input_length := len(high)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 3)
	all_input_data.Set([][]float64{high, low, volume})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_marketfi(
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
	marketfi = outputs[0]
	return
}
