package indicators

//#include "../tulipindicators/indicators/obv.c"
import "C"
import "fmt"

// OBV function wraps `obv' function that provides "On Balance Volume"
//
// Reference: https://tulipindicators.org/obv
func OBV(close, volume []float64) (obv []float64, err error) {
	input_length := len(close)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{close, volume})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_obv(
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
	obv = outputs[0]
	return
}
