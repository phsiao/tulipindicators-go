package indicators

//#include "../tulipindicators/indicators/psar.c"
import "C"
import "fmt"

// PSAR function wraps `psar' function that provides "Parabolic SAR"
//
// Reference: https://tulipindicators.org/psar
func PSAR(high, low []float64, acceleration_factor_step, acceleration_factor_maximum int) (psar []float64, err error) {
	input_length := len(high)
	options := []float64{float64(acceleration_factor_step), float64(acceleration_factor_maximum)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_psar_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{high, low})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_psar(
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
	psar = outputs[0]
	return
}
