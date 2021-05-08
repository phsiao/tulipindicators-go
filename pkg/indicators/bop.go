// bop
// Balance of Power
package indicators

// #cgo LDFLAGS: -lm -L../../tulipindicators -lindicators
//#include "../../tulipindicators/indicators.h"
//#include "../../tulipindicators/utils/buffer.h"
//#include "../../tulipindicators/indicators/bop.c"
import "C"
import "fmt"

func BOP(input1, input2, input3, input4 []float64) (output1 []float64, err error) {
	input_length := len(input1)
	options := []float64{0}
	start := 0

	all_input_data := NewIndicatorData(input_length, 4)
	all_input_data.Set([][]float64{input1, input2, input3, input4})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := NewIndicatorData(output_length, 1)
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
	output1 = outputs[0]
	return
}
