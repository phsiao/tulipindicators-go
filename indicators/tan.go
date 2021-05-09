package indicators

//#include "../tulipindicators/indicators/tan.c"
import "C"
import "fmt"

// TAN function wraps `tan' function that provides "Vector Tangent"
//
// Reference: https://tulipindicators.org/tan
func TAN(real []float64) (tan []float64, err error) {
	input_length := len(real)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{real})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_tan(
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
	tan = outputs[0]
	return
}
