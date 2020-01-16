package vips

// #cgo pkg-config: vips
// #include "header.h"
import "C"
import "unsafe"

func vipsHasICCProfile(in *C.VipsImage) bool {
	return int(C.has_icc_profile(in)) != 0
}

func vipsRemoveICCProfile(in *C.VipsImage) bool {
	return fromGboolean(C.remove_icc_profile(in))
}

func vipsHasIPTC(in *C.VipsImage) bool {
	return int(C.has_iptc(in)) != 0
}

func vipsRemoveMetadata(in *C.VipsImage) {
	C.remove_metadata(in)
}

func vipsGetMetaOrientation(in *C.VipsImage) int {
	return int(C.get_meta_orientation(in))
}

func vipsRemoveMetaOrientation(in *C.VipsImage) {
	C.remove_meta_orientation(in)
}

func vipsSetMetaOrientation(in *C.VipsImage, orientation int) {
	C.set_meta_orientation(in, C.int(orientation))
}

func vipsGetPagesNumber(in *C.VipsImage) int {
	return int(C.get_pages_number(in))
}

func vipsGetPagesDelays(in *C.VipsImage) []int {
	var out *C.int
	var outLength C.int

	err := C.get_pages_delays(in, &out, &outLength)
	if err != 0 {
		info("failed to get pages delays")
		return nil
	}

	cdelays := (*[1 << 28]C.int)(unsafe.Pointer(out))[:outLength:outLength]

	delays := make([]int, int(outLength))
	for i := range cdelays {
		delays[i] = int(cdelays[i])
	}

	return delays
}
