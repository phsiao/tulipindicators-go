package indicators

//#include "../tulipindicators/indicators/vwma.c"
import "C"
import "fmt"

// VWMA function wraps `vwma' function that provides "Volume Weighted Moving Average"
//
// Reference: https://tulipindicators.org/vwma
func VWMA(close, volume []float64, period int) (vwma []float64, err error) {
	input_length := len(close)
	options := []float64{float64(period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_vwma_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{close, volume})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_vwma(
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
	vwma = outputs[0]
	return
}
