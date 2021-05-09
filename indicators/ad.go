package indicators

//#include "../tulipindicators/indicators/ad.c"
import "C"
import "fmt"

// AD function wraps `ad' function that provides "Accumulation/Distribution Line"
//
// Reference: https://tulipindicators.org/ad
func AD(high, low, close, volume []float64) (ad []float64, err error) {
	input_length := len(high)
	options := []float64{0}
	start := 0

	all_input_data := newIndicatorData(input_length, 4)
	all_input_data.Set([][]float64{high, low, close, volume})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := newIndicatorData(output_length, 1)
	defer all_output_data.Destroy()
	ret, err := C.ti_ad(
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
	ad = outputs[0]
	return
}
