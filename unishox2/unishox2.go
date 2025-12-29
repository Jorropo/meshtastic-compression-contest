package unishox2

/*
#cgo CFLAGS: -I.
#include "unishox2.h"
#include <stdlib.h>

int compressDefault(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_DFLT, prev_lines_ptr);
}

int decompressDefault(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_DFLT, prev_lines_ptr);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func CompressDefault(input []byte, delta string) ([]byte, error) {
	cIn := (*C.char)(unsafe.Pointer(unsafe.SliceData(input)))
	cLen := C.int(len(input))

	var outBuf [256]byte
	cOut := (*C.char)(unsafe.Pointer(&outBuf[0]))
	cOLen := C.int(len(outBuf))

	var cDelta *C.char
	if len(delta) > 0 {
		cDelta = C.CString(delta)
	}

	retLen := C.compressDefault(cIn, cLen, cOut, cOLen, cDelta)

	if retLen < 0 {
		return nil, errors.New("unishox2 compression failed")
	}

	return outBuf[:retLen], nil
}

func DecompressDefault(input []byte, delta string) ([]byte, error) {
	cIn := (*C.char)(unsafe.Pointer(unsafe.SliceData(input)))
	cLen := C.int(len(input))

	var outBuf [256]byte
	cOut := (*C.char)(unsafe.Pointer(&outBuf[0]))
	cOLen := C.int(len(outBuf))

	var cDelta *C.char
	if len(delta) > 0 {
		cDelta = C.CString(delta)
	}

	retLen := C.decompressDefault(cIn, cLen, cOut, cOLen, cDelta)

	if retLen < 0 {
		return nil, errors.New("unishox2 decompression failed")
	}

	return outBuf[:retLen], nil
}
