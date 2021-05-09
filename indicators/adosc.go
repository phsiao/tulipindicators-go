package indicators

//#include "../tulipindicators/indicators/adosc.c"
import "C"
import "fmt"

// ADOSC function wraps `adosc' function that provides "Accumulation/Distribution Oscillator"
//
// Reference: https://tulipindicators.org/adosc
func ADOSC(high, low, close, volume []float64, short_period, long_period int) (adosc []float64, err error) {
	input_length := len(high)
	options := []float64{float64(short_period), float64(long_period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_adosc_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 4)
	all_input_data.Set([][]float64{high, low, close, volume})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_adosc(
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
	adosc = outputs[0]
	return
}
