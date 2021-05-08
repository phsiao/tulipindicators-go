package indicators

//#include "../../tulipindicators/indicators.h"
//#include "../../tulipindicators/indicators/adx.c"
import "C"
import (
	"fmt"
)

func ADX(high, low, close []float64, period int) ([]float64, int, error) {

	input_length := len(high)

	options := []float64{float64(period)}

	option_input := (*C.double)(&options[0])
	start, err := C.ti_adx_start(option_input)
	if err != nil {
		return nil, int(start), err
	}

	all_input_data := NewIndicatorData(input_length, 3)
	all_input_data.Set([][]float64{high, low, close})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := NewIndicatorData(output_length, 1)
	defer all_output_data.Destroy()

	ret, err := C.ti_adx(
		(C.int)(input_length),
		(**C.double)(all_input_data.buffer),
		(*C.double)(&options[0]),
		(**C.double)(all_output_data.buffer))
	if err != nil {
		return nil, int(start), err
	}
	if ret != C.TI_OKAY {
		return nil, int(start), fmt.Errorf("ret = %d", ret)
	}

	return all_output_data.Get()[0], int(start), nil
}
