package vips

// #cgo pkg-config: vips
// #include "bridge.h"
import "C"

import (
	"errors"
	"fmt"
	dbg "runtime/debug"
)

var (
	// ErrUnsupportedImageFormat when image type is unsupported
	ErrUnsupportedImageFormat = errors.New("unsupported image format")
)

func handleVipsError(out *C.VipsImage) error {
	if out != nil {
		unrefImage(out)
	}

	s := C.GoString(C.vips_error_buffer())
	C.vips_error_clear()

	stack := string(dbg.Stack())

	return fmt.Errorf("%v\nStack:\n%s", s, stack)
}
