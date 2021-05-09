package indicators

//#include "../tulipindicators/indicators/cci.c"
import "C"
import "fmt"

// CCI function wraps `cci' function that provides "Commodity Channel Index"
//
// Reference: https://tulipindicators.org/cci
func CCI(high, low, close []float64, period int) (cci []float64, err error) {
	input_length := len(high)
	options := []float64{float64(period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_cci_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 3)
	all_input_data.Set([][]float64{high, low, close})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_cci(
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
	cci = outputs[0]
	return
}
