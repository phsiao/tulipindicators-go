package indicators

//#include "../tulipindicators/indicators/aroon.c"
import "C"
import "fmt"

// AROON function wraps `aroon' function that provides "Aroon"
//
// Reference: https://tulipindicators.org/aroon
func AROON(high, low []float64, period int) (aroon_down, aroon_up []float64, err error) {
	input_length := len(high)
	options := []float64{float64(period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_aroon_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{high, low})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 2)
	defer all_output_data.Destroy()
	ret, err := C.ti_aroon(
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
	aroon_down = outputs[0]
	aroon_up = outputs[1]
	return
}
