package indicators

//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include "../tulipindicators/utils/buffer.c"
import "C"
import (
	"unsafe"
)

const (
	// https://dlintw.github.io/gobyexample/public/memory-and-sizeof.html
	ptrSize    = (32 << uintptr(^uintptr(0)>>63)) / 8
	doubleSize = C.sizeof_double
)

type indicatorData struct {
	length int
	rows   int
	buffer unsafe.Pointer
}

func newIndicatorData(output_length int, rows int) indicatorData {
	rval := indicatorData{
		length: output_length,
		rows:   rows,
	}

	buffer := C.malloc((C.ulong)(ptrSize * rval.rows))
	for i := 0; i < rval.rows; i++ {
		offset := unsafe.Pointer(uintptr(unsafe.Pointer(buffer)) + uintptr(ptrSize*i))
		*(**C.double)(offset) = (*C.double)(C.malloc((C.ulong)(doubleSize * rval.length)))
	}

	rval.buffer = buffer

	return rval
}

func (io indicatorData) Set(input [][]float64) {
	if len(input) != io.rows {
		panic("Set() with incorect shape")
	}
	for _, elmt := range input {
		if len(elmt) != io.length {
			panic("Set() with incorect shape")
		}
	}

	for i := 0; i < io.rows; i++ {
		offset := *(**C.double)(unsafe.Pointer(uintptr(io.buffer) + uintptr(ptrSize*i)))
		for j := 0; j < io.length; j++ {
			elmt_addr := unsafe.Pointer(uintptr(unsafe.Pointer(offset)) + uintptr(j*doubleSize))
			(*(*C.double)(elmt_addr)) = (C.double)(input[i][j])
		}
	}
}

func (io indicatorData) Get() [][]float64 {
	rval := make([][]float64, io.rows)
	for i := 0; i < io.rows; i++ {
		rval[i] = make([]float64, io.length)
		offset := *(**C.double)(unsafe.Pointer(uintptr(io.buffer) + uintptr(ptrSize*i)))
		for j := 0; j < io.length; j++ {
			rval[i][j] = (*(*float64)(
				unsafe.Pointer(
					(uintptr(unsafe.Pointer(offset)) + uintptr(j*doubleSize)),
				)))
		}
	}
	return rval
}

func (io indicatorData) Destroy() {
	for i := 0; i < io.rows; i++ {
		offset := *(**C.double)(unsafe.Pointer(uintptr(unsafe.Pointer(io.buffer)) + uintptr(ptrSize*i)))
		C.free(unsafe.Pointer(offset))
	}
	C.free(unsafe.Pointer(io.buffer))
}
