package indicators

//#include "../tulipindicators/indicators/bbands.c"
import "C"
import "fmt"

// BBANDS function wraps `bbands' function that provides "Bollinger Bands"
//
// Reference: https://tulipindicators.org/bbands
func BBANDS(real []float64, period, stddev int) (bbands_lower, bbands_middle, bbands_upper []float64, err error) {
	input_length := len(real)
	options := []float64{float64(period), float64(stddev)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_bbands_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{real})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 3)
	defer all_output_data.Destroy()
	ret, err := C.ti_bbands(
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
	bbands_lower = outputs[0]
	bbands_middle = outputs[1]
	bbands_upper = outputs[2]
	return
}
