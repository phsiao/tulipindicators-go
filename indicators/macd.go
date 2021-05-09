package indicators

//#include "../tulipindicators/indicators/macd.c"
import "C"
import "fmt"

// MACD function wraps `macd' function that provides "Moving Average Convergence/Divergence"
//
// Reference: https://tulipindicators.org/macd
func MACD(real []float64, short_period, long_period, signal_period int) (macd, macd_signal, macd_histogram []float64, err error) {
	input_length := len(real)
	options := []float64{float64(short_period), float64(long_period), float64(signal_period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_macd_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{real})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 3)
	defer all_output_data.Destroy()
	ret, err := C.ti_macd(
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
	macd = outputs[0]
	macd_signal = outputs[1]
	macd_histogram = outputs[2]
	return
}
