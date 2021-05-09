package indicators

//#include "../tulipindicators/indicators/mfi.c"
import "C"
import "fmt"

// MFI function wraps `mfi' function that provides "Money Flow Index"
//
// Reference: https://tulipindicators.org/mfi
func MFI(high, low, close, volume []float64, period int) (mfi []float64, err error) {
	input_length := len(high)
	options := []float64{float64(period)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_mfi_start(option_input)
	if err != nil {
		return
	}

	all_input_data := newIndicatorData(input_length, 4)
	all_input_data.Set([][]float64{high, low, close, volume})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_mfi(
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
	mfi = outputs[0]
	return
}
