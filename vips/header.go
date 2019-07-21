package vips

// #cgo pkg-config: vips
// #include "header.h"
import "C"

func vipsHasProfile(in *C.VipsImage) bool {
	return int(C.has_profile_embed(in)) > 0
}
