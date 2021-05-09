package indicators

//#include "../tulipindicators/indicators/nvi.c"
import "C"
import "fmt"

// NVI function wraps `nvi' function that provides "Negative Volume Index"
//
// Reference: https://tulipindicators.org/nvi
func NVI(close, volume []float64) (nvi []float64, err error) {
	input_length := len(close)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{close, volume})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_nvi(
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
	nvi = outputs[0]
	return
}
