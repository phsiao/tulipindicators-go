package main

//#include "../../tulipindicators/indicators.h"
//#include "../../tulipindicators/indicators/sma.c"
import "C"

import (
	"fmt"
	"unsafe"
)

func print_array(p []float64, size int) {
	for i := 0; i < size; i++ {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%.1f", p[i])
	}
	fmt.Printf("\n")
}
func main() {
	data_in := []float64{5, 8, 12, 11, 9, 8, 7, 10, 11, 13}
	input_length := len(data_in)

	fmt.Printf("We have %d bars of input data.\n", input_length)
	print_array(data_in, input_length)

	options := []float64{3}
	fmt.Printf("Our option array is: ")
	print_array(options, len(options))

	input := (*C.double)(&options[0])
	start, err := C.ti_sma_start(input)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The start amount is: %d\n", start)

	output_length := input_length - int(start)
	data_out := [][]float64{
		make([]float64, output_length),
	}
	fmt.Printf("The output length is: %d\n", output_length)

	all_inputs := (**C.double)(C.malloc(8))
	*all_inputs = (*C.double)(&data_in[0])

	all_outputs := (**C.double)(C.malloc(8))
	*all_outputs = (*C.double)(C.malloc((C.ulong)(8 * output_length)))

	ret, err := C.ti_sma((C.int)(input_length),
		all_inputs,
		(*C.double)(&options[0]),
		all_outputs)
	if err != nil {
		panic(err.Error())
	}
	if ret != C.TI_OKAY {
		panic(fmt.Sprintf("ret = %d", ret))
	}

	for i := 0; i < output_length; i++ {
		data_out[0][i] = (float64)(*(*float64)(unsafe.Pointer(
			(uintptr(unsafe.Pointer(*all_outputs)) + uintptr(i*8)),
		)))
	}

	fmt.Printf("The output data is: ")
	print_array(data_out[0], output_length)
}
