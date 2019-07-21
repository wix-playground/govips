package vips

// #cgo pkg-config: vips
// #include "foreign.h"
import "C"
import (
	"log"
	"runtime"
	"unsafe"
)

func vipsLoadFromBuffer(buf []byte) (*C.VipsImage, ImageType, error) {
	// Reference buf here so it's not garbage collected during image initialization.
	defer runtime.KeepAlive(buf)

	var out *C.VipsImage

	imageType := vipsDetermineImageType(buf)
	if imageType == ImageTypeUnknown {
		if len(buf) > 2 {
			log.Printf("Failed to understand image format size=%d %x %x %x", len(buf), buf[0], buf[1], buf[2])
		} else {
			log.Printf("Failed to understand image format size=%d", len(buf))
		}
		return nil, ImageTypeUnknown, ErrUnsupportedImageFormat
	}

	bufLength := C.size_t(len(buf))
	imageBuf := unsafe.Pointer(&buf[0])

	err := C.init_image(imageBuf, bufLength, C.int(imageType), &out)
	if err != 0 {
		return nil, ImageTypeUnknown, handleVipsError(out)
	}

	return out, imageType, nil
}
