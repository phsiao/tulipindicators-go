package indicators

//#include "../tulipindicators/indicators/div.c"
import "C"
import "fmt"

// DIV function wraps `div' function that provides "Vector Division"
//
// Reference: https://tulipindicators.org/div
func DIV(real1, real2 []float64) (div []float64, err error) {
	input_length := len(real1)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{real1, real2})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_div(
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
	div = outputs[0]
	return
}
