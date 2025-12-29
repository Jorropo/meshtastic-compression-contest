package unishox2

/*
#cgo CFLAGS: -I.
#include "unishox2.h"
#include <stdlib.h>

/// Default preset parameter set. When composition of text is know beforehand, the other parameter sets in this section can be
/// used to achieve more compression.
#define USX_PSET_DFLT USX_HCODES_DFLT, USX_HCODE_LENS_DFLT, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set for English Alphabet only content
#define USX_PSET_ALPHA_ONLY USX_HCODES_ALPHA_ONLY, USX_HCODE_LENS_ALPHA_ONLY, USX_FREQ_SEQ_TXT, USX_TEMPLATES
/// Preset parameter set for Alpha numeric content
#define USX_PSET_ALPHA_NUM_ONLY USX_HCODES_ALPHA_NUM_ONLY, USX_HCODE_LENS_ALPHA_NUM_ONLY, USX_FREQ_SEQ_TXT, USX_TEMPLATES
/// Preset parameter set for Alpha numeric and symbol content
#define USX_PSET_ALPHA_NUM_SYM_ONLY                                                                                              \
    USX_HCODES_ALPHA_NUM_SYM_ONLY, USX_HCODE_LENS_ALPHA_NUM_SYM_ONLY, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set for Alpha numeric symbol content having predominantly text
#define USX_PSET_ALPHA_NUM_SYM_ONLY_TXT                                                                                          \
    USX_HCODES_ALPHA_NUM_SYM_ONLY, USX_HCODE_LENS_ALPHA_NUM_SYM_ONLY, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set favouring Alphabet content
#define USX_PSET_FAVOR_ALPHA USX_HCODES_FAVOR_ALPHA, USX_HCODE_LENS_FAVOR_ALPHA, USX_FREQ_SEQ_TXT, USX_TEMPLATES
/// Preset parameter set favouring repeating sequences
#define USX_PSET_FAVOR_DICT USX_HCODES_FAVOR_DICT, USX_HCODE_LENS_FAVOR_DICT, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set favouring symbols
#define USX_PSET_FAVOR_SYM USX_HCODES_FAVOR_SYM, USX_HCODE_LENS_FAVOR_SYM, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set favouring unlaut letters
#define USX_PSET_FAVOR_UMLAUT USX_HCODES_FAVOR_UMLAUT, USX_HCODE_LENS_FAVOR_UMLAUT, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set for when there are no repeating sequences
#define USX_PSET_NO_DICT USX_HCODES_NO_DICT, USX_HCODE_LENS_NO_DICT, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set for when there are no unicode symbols
#define USX_PSET_NO_UNI USX_HCODES_NO_UNI, USX_HCODE_LENS_NO_UNI, USX_FREQ_SEQ_DFLT, USX_TEMPLATES
/// Preset parameter set for when there are no unicode symbols favouring text
#define USX_PSET_NO_UNI_FAVOR_TEXT USX_HCODES_NO_UNI, USX_HCODE_LENS_NO_UNI, USX_FREQ_SEQ_TXT, USX_TEMPLATES
/// Preset parameter set favouring URL content
#define USX_PSET_URL USX_HCODES_DFLT, USX_HCODE_LENS_DFLT, USX_FREQ_SEQ_URL, USX_TEMPLATES
/// Preset parameter set favouring JSON content
#define USX_PSET_JSON USX_HCODES_DFLT, USX_HCODE_LENS_DFLT, USX_FREQ_SEQ_JSON, USX_TEMPLATES
/// Preset parameter set favouring JSON content having no Unicode symbols
#define USX_PSET_JSON_NO_UNI USX_HCODES_NO_UNI, USX_HCODE_LENS_NO_UNI, USX_FREQ_SEQ_JSON, USX_TEMPLATES
/// Preset parameter set favouring XML content
#define USX_PSET_XML USX_HCODES_DFLT, USX_HCODE_LENS_DFLT, USX_FREQ_SEQ_XML, USX_TEMPLATES
/// Preset parameter set favouring HTML content
#define USX_PSET_HTML USX_HCODES_DFLT, USX_HCODE_LENS_DFLT, USX_FREQ_SEQ_HTML, USX_TEMPLATES

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

int compressAlphaOnly(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_ONLY, prev_lines_ptr);
}
int decompressAlphaOnly(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_ONLY, prev_lines_ptr);
}

int compressAlphaNumOnly(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_ONLY, prev_lines_ptr);
}
int decompressAlphaNumOnly(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_ONLY, prev_lines_ptr);
}

int compressAlphaNumSymOnly(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY, prev_lines_ptr);
}
int decompressAlphaNumSymOnly(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY, prev_lines_ptr);
}

int compressAlphaNumSymOnlyText(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY_TXT, prev_lines_ptr);
}
int decompressAlphaNumSymOnlyText(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_ALPHA_NUM_SYM_ONLY_TXT, prev_lines_ptr);
}

int compressFavorAlpha(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_ALPHA, prev_lines_ptr);
}
int decompressFavorAlpha(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_ALPHA, prev_lines_ptr);
}

int compressFavorDict(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_DICT, prev_lines_ptr);
}
int decompressFavorDict(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_DICT, prev_lines_ptr);
}

int compressFavorSym(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_SYM, prev_lines_ptr);
}
int decompressFavorSym(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_SYM, prev_lines_ptr);
}

int compressFavorUmlaut(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_FAVOR_UMLAUT, prev_lines_ptr);
}
int decompressFavorUmlaut(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_FAVOR_UMLAUT, prev_lines_ptr);
}

int compressNoDict(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_NO_DICT, prev_lines_ptr);
}
int decompressNoDict(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_NO_DICT, prev_lines_ptr);
}

int compressNoUni(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_NO_UNI, prev_lines_ptr);
}
int decompressNoUni(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_NO_UNI, prev_lines_ptr);
}

int compressNoUniFavorText(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_NO_UNI_FAVOR_TEXT, prev_lines_ptr);
}
int decompressNoUniFavorText(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_NO_UNI_FAVOR_TEXT, prev_lines_ptr);
}

int compressURL(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_URL, prev_lines_ptr);
}
int decompressURL(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_URL, prev_lines_ptr);
}

int compressJSON(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_JSON, prev_lines_ptr);
}
int decompressJSON(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_JSON, prev_lines_ptr);
}

int compressJSONNoUni(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_JSON_NO_UNI, prev_lines_ptr);
}
int decompressJSONNoUni(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_JSON_NO_UNI, prev_lines_ptr);
}

int compressXML(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_XML, prev_lines_ptr);
}
int decompressXML(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_XML, prev_lines_ptr);
}

int compressHTML(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_compress_lines(in, len, out, olen, USX_PSET_HTML, prev_lines_ptr);
}
int decompressHTML(const char *in, int len, char *out, int olen, char *previous_line) {
	struct us_lnk_lst prev_lines = {0};
	struct us_lnk_lst *prev_lines_ptr = NULL;
	if (previous_line != NULL) {
		prev_lines.data = previous_line;
		prev_lines_ptr = &prev_lines;
	}
	return unishox2_decompress_lines(in, len, out, olen, USX_PSET_HTML, prev_lines_ptr);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func compress(input []byte, delta string, compressFunc func(*C.char, C.int, *C.char, C.int, *C.char) C.int) ([]byte, error) {
	cIn := (*C.char)(unsafe.Pointer(unsafe.SliceData(input)))
	cLen := C.int(len(input))

	var outBuf [256]byte
	cOut := (*C.char)(unsafe.Pointer(&outBuf[0]))
	cOLen := C.int(len(outBuf))

	var cDelta *C.char
	if len(delta) > 0 {
		cDelta = C.CString(delta)
		defer C.free(unsafe.Pointer(cDelta))
	}

	retLen := compressFunc(cIn, cLen, cOut, cOLen, cDelta)
	if retLen < 0 {
		return nil, errors.New("unishox2 compression failed")
	}

	return outBuf[:retLen], nil
}
func decompress(input []byte, delta string, decompressFunc func(*C.char, C.int, *C.char, C.int, *C.char) C.int) ([]byte, error) {
	cIn := (*C.char)(unsafe.Pointer(unsafe.SliceData(input)))
	cLen := C.int(len(input))

	var outBuf [256]byte
	cOut := (*C.char)(unsafe.Pointer(&outBuf[0]))
	cOLen := C.int(len(outBuf))

	var cDelta *C.char
	if len(delta) > 0 {
		cDelta = C.CString(delta)
		defer C.free(unsafe.Pointer(cDelta))
	}

	retLen := decompressFunc(cIn, cLen, cOut, cOLen, cDelta)
	if retLen < 0 {
		return nil, errors.New("unishox2 decompression failed")
	}

	return outBuf[:retLen], nil
}

func CompressDefault(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressDefault(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressDefault(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressDefault(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressAlphaOnly(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressAlphaOnly(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressAlphaOnly(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressAlphaOnly(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressAlphaNumOnly(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressAlphaNumOnly(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressAlphaNumOnly(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressAlphaNumOnly(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressAlphaNumSymOnly(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressAlphaNumSymOnly(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressAlphaNumSymOnly(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressAlphaNumSymOnly(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressAlphaNumSymOnlyText(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressAlphaNumSymOnlyText(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressAlphaNumSymOnlyText(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressAlphaNumSymOnlyText(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressFavorAlpha(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressFavorAlpha(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressFavorAlpha(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressFavorAlpha(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressFavorDict(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressFavorDict(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressFavorDict(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressFavorDict(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressFavorSym(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressFavorSym(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressFavorSym(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressFavorSym(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressFavorUmlaut(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressFavorUmlaut(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressFavorUmlaut(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressFavorUmlaut(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressNoDict(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressNoDict(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressNoDict(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressNoDict(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressNoUni(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressNoUni(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressNoUni(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressNoUni(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressNoUniFavorText(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressNoUniFavorText(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressNoUniFavorText(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressNoUniFavorText(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressURL(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressURL(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressURL(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressURL(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressJSON(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressJSON(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressJSON(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressJSON(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressJSONNoUni(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressJSONNoUni(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressJSONNoUni(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressJSONNoUni(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressXML(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressXML(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressXML(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressXML(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}

func CompressHTML(input []byte, delta string) ([]byte, error) {
	return compress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.compressHTML(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
func DecompressHTML(input []byte, delta string) ([]byte, error) {
	return decompress(input, delta, func(inPtr *C.char, inLen C.int, outPtr *C.char, outLen C.int, deltaPtr *C.char) C.int {
		return C.decompressHTML(inPtr, inLen, outPtr, outLen, deltaPtr)
	})
}
