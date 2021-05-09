package indicators

//#include "../tulipindicators/indicators/bop.c"
import "C"
import "fmt"

// BOP function wraps `bop' function that provides "Balance of Power"
//
// Reference: https://tulipindicators.org/bop
func BOP(open, high, low, close []float64) (bop []float64, err error) {
	input_length := len(open)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 4)
	all_input_data.Set([][]float64{open, high, low, close})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_bop(
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
	bop = outputs[0]
	return
}
