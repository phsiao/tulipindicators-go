// fisher
// Fisher Transform
package indicators

// #cgo LDFLAGS: -lm -L../../tulipindicators -lindicators
//#include "../../tulipindicators/indicators.h"
//#include "../../tulipindicators/utils/buffer.h"
//#include "../../tulipindicators/indicators/fisher.c"
import "C"
import "fmt"

func FISHER(input1, input2 []float64, options1 int) (output1, output2 []float64, err error) {
	input_length := len(input1)
	options := []float64{float64(options1)}
	option_input := (*C.double)(&options[0])
	start, err := C.ti_fisher_start(option_input)
	if err != nil {
		return
	}

	all_input_data := NewIndicatorData(input_length, 2)
	all_input_data.Set([][]float64{input1, input2})
	defer all_input_data.Destroy()

	output_length := input_length - int(start)
	all_output_data := NewIndicatorData(output_length, 2)
	defer all_output_data.Destroy()
	ret, err := C.ti_fisher(
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
