package indicators

//#include "../tulipindicators/indicators/di.c"
import "C"
import "fmt"

// DI function wraps `di' function that provides "Directional Indicator"
//
// Reference: https://tulipindicators.org/di
func DI(high, low, close []float64, period int) (plus_di, minus_di []float64, err error) {
	input_length := len(high)
	options := []float64{float64(period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_di_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 3)
	all_input_data.Set([][]float64{high, low, close})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 2)
	defer all_output_data.Destroy()
	ret, err := C.ti_di(
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
	plus_di = outputs[0]
	minus_di = outputs[1]
	return
}
