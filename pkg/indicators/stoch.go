// stoch
// Stochastic Oscillator
package indicators

//#include "../../tulipindicators/indicators/stoch.c"
import "C"
import "fmt"

func STOCH(input1, input2, input3 []float64, options1, options2, options3 int) (output1, output2 []float64, err error) {
	input_length := len(input1)
	options := []float64{float64(options1), float64(options2), float64(options3)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_stoch_start(option_input)
	if err != nil {
		return
	}

	all_input_data := NewIndicatorData(input_length, 3)
	all_input_data.Set([][]float64{input1, input2, input3})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := NewIndicatorData(output_length, 2)
	defer all_output_data.Destroy()
	ret, err := C.ti_stoch(
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
	return
}
