package indicators

//#include "../tulipindicators/indicators/medprice.c"
import "C"
import "fmt"

// MEDPRICE function wraps `medprice' function that provides "Median Price"
//
// Reference: https://tulipindicators.org/medprice
func MEDPRICE(high, low []float64) (medprice []float64, err error) {
	input_length := len(high)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{high, low})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_medprice(
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
	medprice = outputs[0]
	return
}
