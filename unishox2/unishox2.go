package unishox2

/*
#cgo CFLAGS: -I.
#include "unishox2.h"
#include <stdlib.h>

int compressDefault(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_DFLT, prev_lines_ptr, skip_bits);
}
int decompressDefault(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_DFLT, prev_lines_ptr, skip_bits);
}

int compressAlphaOnly(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_ONLY, prev_lines_ptr, skip_bits);
}
int decompressAlphaOnly(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_ONLY, prev_lines_ptr, skip_bits);
}

int compressAlphaNumOnly(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_ONLY, prev_lines_ptr, skip_bits);
}
int decompressAlphaNumOnly(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_ONLY, prev_lines_ptr, skip_bits);
}

int compressAlphaNumSymOnly(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY, prev_lines_ptr, skip_bits);
}
int decompressAlphaNumSymOnly(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY, prev_lines_ptr, skip_bits);
}

int compressAlphaNumSymOnlyText(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY_TXT, prev_lines_ptr, skip_bits);
}
int decompressAlphaNumSymOnlyText(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY_TXT, prev_lines_ptr, skip_bits);
}

int compressFavorAlpha(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_ALPHA, prev_lines_ptr, skip_bits);
}
int decompressFavorAlpha(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_ALPHA, prev_lines_ptr, skip_bits);
}

int compressFavorDict(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_DICT, prev_lines_ptr, skip_bits);
}
int decompressFavorDict(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_DICT, prev_lines_ptr, skip_bits);
}

int compressFavorSym(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_SYM, prev_lines_ptr, skip_bits);
}
int decompressFavorSym(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_SYM, prev_lines_ptr, skip_bits);
}

int compressFavorUmlaut(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_UMLAUT, prev_lines_ptr, skip_bits);
}
int decompressFavorUmlaut(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_UMLAUT, prev_lines_ptr, skip_bits);
}

int compressNoDict(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_NO_DICT, prev_lines_ptr, skip_bits);
}
int decompressNoDict(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_NO_DICT, prev_lines_ptr, skip_bits);
}

int compressNoUni(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_NO_UNI, prev_lines_ptr, skip_bits);
}
int decompressNoUni(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_NO_UNI, prev_lines_ptr, skip_bits);
}

int compressNoUniFavorText(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_NO_UNI_FAVOR_TEXT, prev_lines_ptr, skip_bits);
}
int decompressNoUniFavorText(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_NO_UNI_FAVOR_TEXT, prev_lines_ptr, skip_bits);
}

int compressURL(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_URL, prev_lines_ptr, skip_bits);
}
int decompressURL(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_URL, prev_lines_ptr, skip_bits);
}

int compressJSON(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_JSON, prev_lines_ptr, skip_bits);
}
int decompressJSON(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_JSON, prev_lines_ptr, skip_bits);
}

int compressJSONNoUni(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_JSON_NO_UNI, prev_lines_ptr, skip_bits);
}
int decompressJSONNoUni(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_JSON_NO_UNI, prev_lines_ptr, skip_bits);
}

int compressXML(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_XML, prev_lines_ptr, skip_bits);
}
int decompressXML(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_XML, prev_lines_ptr, skip_bits);
}

int compressHTML(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_HTML, prev_lines_ptr, skip_bits);
}
int decompressHTML(const char *in, int len, char *out, int olen, char *previous_line, uint8_t skip_bits) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_HTML, prev_lines_ptr, skip_bits);
}
*/
import "C"
import (
	"unsafe"
)

func compress(input, output []byte, delta string, skipBits uint8, compressFunc func(*C.char, C.int, *C.char, C.int, *C.char, C.uint8_t) C.int) int {
	cIn := (*C.char)(unsafe.Pointer(unsafe.SliceData(input)))
	cLen := C.int(len(input))

	cOut := (*C.char)(unsafe.Pointer(unsafe.SliceData(output)))
	cOLen := C.int(len(output))

	var cDelta *C.char
	if len(delta) > 0 {
		cDelta = C.CString(delta)
		defer C.free(unsafe.Pointer(cDelta))
	}

	return int(compressFunc(cIn, cLen, cOut, cOLen, cDelta, C.uint8_t(skipBits)))
}
func decompress(input, output []byte, delta string, skipBits uint8, decompressFunc func(*C.char, C.int, *C.char, C.int, *C.char, C.uint8_t) C.int) int {
	cIn := (*C.char)(unsafe.Pointer(unsafe.SliceData(input)))
	cLen := C.int(len(input))

	cOut := (*C.char)(unsafe.Pointer(unsafe.SliceData(output)))
	cOLen := C.int(len(output))

	var cDelta *C.char
	if len(delta) > 0 {
		cDelta = C.CString(delta)
		defer C.free(unsafe.Pointer(cDelta))
	}

	return int(decompressFunc(cIn, cLen, cOut, cOLen, cDelta, C.uint8_t(skipBits)))
}

func CompressDefault(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressDefault(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressDefault(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressDefault(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressAlphaOnly(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressAlphaOnly(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressAlphaOnly(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressAlphaOnly(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressAlphaNumOnly(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressAlphaNumOnly(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressAlphaNumOnly(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressAlphaNumOnly(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressAlphaNumSymOnly(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressAlphaNumSymOnly(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressAlphaNumSymOnly(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressAlphaNumSymOnly(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressAlphaNumSymOnlyText(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressAlphaNumSymOnlyText(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressAlphaNumSymOnlyText(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressAlphaNumSymOnlyText(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressFavorAlpha(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressFavorAlpha(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressFavorAlpha(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressFavorAlpha(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressFavorDict(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressFavorDict(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressFavorDict(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressFavorDict(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressFavorSym(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressFavorSym(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressFavorSym(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressFavorSym(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressFavorUmlaut(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressFavorUmlaut(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressFavorUmlaut(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressFavorUmlaut(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressNoDict(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressNoDict(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressNoDict(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressNoDict(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressNoUni(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressNoUni(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressNoUni(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressNoUni(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressNoUniFavorText(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressNoUniFavorText(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressNoUniFavorText(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressNoUniFavorText(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressURL(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressURL(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressURL(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressURL(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressJSON(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressJSON(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressJSON(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressJSON(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressJSONNoUni(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressJSONNoUni(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressJSONNoUni(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressJSONNoUni(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressXML(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressXML(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressXML(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressXML(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}

func CompressHTML(input, output []byte, delta string, skipBits uint8) int {
	return compress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.compressHTML(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
func DecompressHTML(input, output []byte, delta string, skipBits uint8) int {
	return decompress(input, output, delta, skipBits, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char, skipBits C.uint8_t) C.int {
		return C.decompressHTML(inPtr, inLen, outPtr, outLen, deltaPtr, skipBits)
	})
}
