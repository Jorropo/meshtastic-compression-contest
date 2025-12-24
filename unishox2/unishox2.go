// FIXME: This package is unsafe, it lacks bound check on output compression buffers.
package unishox2

/*
#cgo CFLAGS: -I.
#include "unishox2.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// CompressSimple wraps unishox2_compress_simple.
func CompressSimple(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}

	cIn := (*C.char)(unsafe.Pointer(unsafe.SliceData(input)))
	cLen := C.int(len(input))

	maxOutLen := 256 * 4
	outBuf := make([]byte, maxOutLen)
	cOut := (*C.char)(unsafe.Pointer(&outBuf[0]))

	retLen := C.unishox2_compress_simple(cIn, cLen, cOut)

	if retLen < 0 {
		return nil, errors.New("unishox2 compression failed")
	}

	return outBuf[:retLen], nil
}

// DecompressSimple wraps unishox2_decompress_simple.
func DecompressSimple(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}

	cIn := (*C.char)(unsafe.Pointer(&input[0]))
	cLen := C.int(len(input))

	maxOutLen := 256 * 4
	outBuf := make([]byte, maxOutLen)
	cOut := (*C.char)(unsafe.Pointer(&outBuf[0]))

	retLen := C.unishox2_decompress_simple(cIn, cLen, cOut)

	if retLen < 0 {
		return nil, errors.New("unishox2 decompression failed (buffer too small or corrupt data)")
	}

	return outBuf[:retLen], nil
}
