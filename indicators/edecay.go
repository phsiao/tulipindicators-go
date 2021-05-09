package indicators

//#include "../tulipindicators/indicators/edecay.c"
import "C"
import "fmt"

// EDECAY function wraps `edecay' function that provides "Exponential Decay"
//
// Reference: https://tulipindicators.org/edecay
func EDECAY(real []float64, period int) (edecay []float64, err error) {
	input_length := len(real)
	options := []float64{float64(period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_edecay_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{real})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_edecay(
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
	edecay = outputs[0]
	return
}
