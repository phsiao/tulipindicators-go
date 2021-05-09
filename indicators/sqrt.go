package indicators

//#include "../tulipindicators/indicators/sqrt.c"
import "C"
import "fmt"

// SQRT function wraps `sqrt' function that provides "Vector Square Root"
//
// Reference: https://tulipindicators.org/sqrt
func SQRT(input1 []float64) (output1 []float64, err error) {
	input_length := len(input1)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{input1})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_sqrt(
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
	output1 = outputs[0]
	return
}
