package indicators

//#cgo LDFLAGS:
//#include <stdio.h>
//#include <stdlib.h>
import "C"
import "unsafe"

const (
	// https://dlintw.github.io/gobyexample/public/memory-and-sizeof.html
	PtrSize    = 32 << uintptr(^uintptr(0)>>63)
	DoubleSize = C.sizeof_double
)

type IndicatorData struct {
	length int
	rows   int
	buffer unsafe.Pointer
}

func NewIndicatorData(output_length int, rows int) IndicatorData {
	rval := IndicatorData{
		length: output_length,
		rows:   rows,
	}

	buffer := C.malloc((C.ulong)(PtrSize * rval.rows))
	for i := 0; i < rval.rows; i++ {
		offset := unsafe.Pointer(uintptr(unsafe.Pointer(buffer)) + uintptr(PtrSize*i))
		*(**C.double)(offset) = (*C.double)(C.malloc((C.ulong)(DoubleSize * rval.length)))
	}

	rval.buffer = buffer

	return rval
}

func (io IndicatorData) Set(input [][]float64) {
	if len(input) != io.rows {
		panic("Set() with incorect shape")
	}
	if len(input[0]) != io.length {
		panic("Set() with incorect shape")
	}

	for i := 0; i < io.rows; i++ {
		offset := *(**C.double)(unsafe.Pointer(uintptr(io.buffer) + uintptr(PtrSize*i)))
		for j := 0; j < io.length; j++ {
			elmt_addr := unsafe.Pointer(uintptr(unsafe.Pointer(offset)) + uintptr(j*DoubleSize))
			(*(*C.double)(elmt_addr)) = (C.double)(input[i][j])
		}
	}
}

func (io IndicatorData) Get() [][]float64 {
	rval := make([][]float64, io.rows)
	for i := 0; i < io.rows; i++ {
		rval[i] = make([]float64, io.length)
		offset := *(**C.double)(unsafe.Pointer(uintptr(io.buffer) + uintptr(PtrSize*i)))
		for j := 0; j < io.length; j++ {
			rval[i][j] = (*(*float64)(
				unsafe.Pointer(
					(uintptr(unsafe.Pointer(offset)) + uintptr(j*DoubleSize)),
				)))
		}
	}
	return rval
}

func (io IndicatorData) Destroy() {
	for i := 0; i < io.rows; i++ {
		offset := (**C.double)(unsafe.Pointer(uintptr(unsafe.Pointer(io.buffer)) + uintptr(PtrSize*i)))
		C.free(unsafe.Pointer(*offset))
	}
	C.free(unsafe.Pointer(io.buffer))
}
