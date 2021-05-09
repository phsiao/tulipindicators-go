package indicators

//#include "../tulipindicators/indicators/sinh.c"
import "C"
import "fmt"

// SINH function wraps `sinh' function that provides "Vector Hyperbolic Sine"
//
// Reference: https://tulipindicators.org/sinh
func SINH(real []float64) (sinh []float64, err error) {
	input_length := len(real)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{real})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_sinh(
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
	sinh = outputs[0]
	return
}
