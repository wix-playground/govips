package vips

// #cgo pkg-config: vips
// #include <vips/vips.h>
import "C"

import (
	"errors"
	"fmt"
	dbg "runtime/debug"
	"unsafe"
)

var (
	// ErrUnsupportedImageFormat when image type is unsupported
	ErrUnsupportedImageFormat = errors.New("unsupported image format")
)

func handleImageError(out *C.VipsImage) error {
	if out != nil {
		unrefImage(out)
	}

	return handleVipsError()
}

func handleSaveBufferError(out unsafe.Pointer) error {
	if out != nil {
		gFreePointer(out)
	}

	return handleVipsError()
}

func handleVipsError() error {
	s := C.GoString(C.vips_error_buffer())
	C.vips_error_clear()

	stack := string(dbg.Stack())

	return fmt.Errorf("%v\nStack:\n%s", s, stack)
}
