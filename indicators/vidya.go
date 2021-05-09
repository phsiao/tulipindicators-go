package indicators

//#include "../tulipindicators/indicators/vidya.c"
import "C"
import "fmt"

// VIDYA function wraps `vidya' function that provides "Variable Index Dynamic Average"
//
// Reference: https://tulipindicators.org/vidya
func VIDYA(real []float64, short_period, long_period, alpha int) (vidya []float64, err error) {
	input_length := len(real)
	options := []float64{float64(short_period), float64(long_period), float64(alpha)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_vidya_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{real})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_vidya(
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
	vidya = outputs[0]
	return
}
