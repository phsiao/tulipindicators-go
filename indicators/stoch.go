package indicators

//#include "../tulipindicators/indicators/stoch.c"
import "C"
import "fmt"

// STOCH function wraps `stoch' function that provides "Stochastic Oscillator"
//
// Reference: https://tulipindicators.org/stoch
func STOCH(high, low, close []float64, pctk_period, pctk_slowing_period, pctd_period int) (stoch_k, stoch_d []float64, err error) {
	input_length := len(high)
	options := []float64{float64(pctk_period), float64(pctk_slowing_period), float64(pctd_period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_stoch_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 3)
	all_input_data.Set([][]float64{high, low, close})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 2)
	defer all_output_data.Destroy()
	ret, err := C.ti_stoch(
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
	stoch_k = outputs[0]
	stoch_d = outputs[1]
	return
}
