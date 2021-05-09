package indicators

//#include "../tulipindicators/indicators/bbands.c"
import "C"
import "fmt"

// BBANDS function wraps `bbands' function that provides "Bollinger Bands"
//
// Reference: https://tulipindicators.org/bbands
func BBANDS(input1 []float64, option1, option2 int) (output1, output2, output3 []float64, err error) {
	input_length := len(input1)
	options := []float64{float64(option1), float64(option2)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_bbands_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 1)
	all_input_data.Set([][]float64{input1})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 3)
	defer all_output_data.Destroy()
	ret, err := C.ti_bbands(
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
	output2 = outputs[1]
	output3 = outputs[2]
	return
}
