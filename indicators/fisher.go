package indicators

//#include "../tulipindicators/indicators/fisher.c"
import "C"
import "fmt"

// FISHER function wraps `fisher' function that provides "Fisher Transform"
//
// Reference: https://tulipindicators.org/fisher
func FISHER(high, low []float64, period int) (fisher, fisher_signal []float64, err error) {
	input_length := len(high)
	options := []float64{float64(period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_fisher_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{high, low})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 2)
	defer all_output_data.Destroy()
	ret, err := C.ti_fisher(
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
	fisher = outputs[0]
	fisher_signal = outputs[1]
	return
}
