package vips

// #cgo pkg-config: vips
// #include "bridge.h"
// #include "foreign.h"
// #include "color.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type vipsLabelOptions struct {
	Text      *C.char
	Font      *C.char
	Width     C.int
	Height    C.int
	OffsetX   C.int
	OffsetY   C.int
	Alignment C.VipsAlign
	DPI       C.int
	Margin    C.int
	Opacity   C.float
	Color     [3]C.double
}

func vipsCall(name string, options []*Option) error {
	operation := vipsOperationNew(name)

	return vipsCallOperation(operation, options)
}

func vipsOperationNew(name string) *C.VipsOperation {
	cName := C.CString(name)
	defer freeCString(cName)

	return C.vips_operation_new(cName)
}

func vipsCallOperation(operation *C.VipsOperation, options []*Option) error {
	// todo: replace with https://jcupitt.github.io/libvips/API/current/VipsOperation.html#vips-cache-operation-build

	for _, option := range options {
		if option.Output() {
			continue
		}
		defer option.Close()

		cName := C.CString(option.Name)
		defer freeCString(cName)

		C.gobject_set_property((*C.VipsObject)(unsafe.Pointer(operation)), cName, option.GValue())
	}

	if ret := C.vips_cache_operation_buildp(&operation); ret != 0 {
		return handleVipsError(nil)
	}

	defer unrefPointer(unsafe.Pointer(operation))

	for _, option := range options {
		if !option.Output() {
			continue
		}
		defer option.Close()

		cName := C.CString(option.Name)
		defer freeCString(cName)

		C.g_object_get_property((*C.GObject)(unsafe.Pointer(operation)), (*C.gchar)(cName), option.GValue())
	}

	return nil
}

func vipsExportBuffer(image *C.VipsImage, params *ExportParams) ([]byte, ImageType, error) {
	tmpImage, err := vipsPrepareForExport(image, params)
	if err != nil {
		return nil, ImageTypeUnknown, err
	}

	// If these are equal, then we don't want to deref the original image as
	// the original will be returned if the target colorspace is not supported
	if tmpImage != image {
		defer unrefImage(tmpImage)
	}

	cLen := C.size_t(0)
	var cErr C.int
	interlaced := C.int(boolToInt(params.Interlaced))
	quality := C.int(params.Quality)
	lossless := C.int(boolToInt(params.Lossless))
	stripMetadata := C.int(boolToInt(params.StripMetadata))
	format := params.Format

	if format != ImageTypeUnknown && !IsTypeSupported(format) {
		return nil, ImageTypeUnknown, fmt.Errorf("cannot save to %#v", ImageTypes[format])
	}

	if params.BackgroundColor != nil && vipsHasAlpha(tmpImage) {
		tmpImage, err = vipsFlatten(tmpImage, params.BackgroundColor)
		if err != nil {
			return nil, ImageTypeUnknown, err
		}
	}

	var ptr unsafe.Pointer

	switch format {
	case ImageTypeWEBP:
		incOpCounter("save_webp_buffer")
		cErr = C.save_webp_buffer(tmpImage, &ptr, &cLen, stripMetadata, quality, lossless)
	case ImageTypePNG:
		incOpCounter("save_png_buffer")
		cErr = C.save_png_buffer(tmpImage, &ptr, &cLen, stripMetadata, C.int(params.Compression), quality, interlaced)
	case ImageTypeTIFF:
		incOpCounter("save_tiff_buffer")
		cErr = C.save_tiff_buffer(tmpImage, &ptr, &cLen)
	case ImageTypeHEIF:
		incOpCounter("save_heif_buffer")
		cErr = C.save_heif_buffer(tmpImage, &ptr, &cLen, quality, lossless)
	default:
		incOpCounter("save_jpeg_buffer")
		format = ImageTypeJPEG
		cErr = C.save_jpeg_buffer(tmpImage, &ptr, &cLen, stripMetadata, quality, interlaced)
	}

	if int(cErr) != 0 {
		return nil, ImageTypeUnknown, handleVipsError(nil)
	}

	buf := C.GoBytes(ptr, C.int(cLen))
	gFreePointer(ptr)

	return buf, format, nil
}

func vipsPrepareForExport(in *C.VipsImage, params *ExportParams) (*C.VipsImage, error) {
	if params.StripProfile && vipsHasProfile(in) {
		C.remove_icc_profile(in)
	}

	if params.Quality == 0 {
		params.Quality = defaultQuality
	}

	if params.Compression == 0 {
		params.Compression = defaultCompression
	}

	// Use a default interpretation and cast it to C type
	if params.Interpretation == 0 {
		params.Interpretation = Interpretation(in.Type)
	}

	interpretation := C.VipsInterpretation(params.Interpretation)

	// Apply the proper colour space
	if int(C.is_colorspace_supported(in)) == 1 && interpretation != in.Type {
		var out *C.VipsImage

		err := C.to_colorspace(in, &out, interpretation)
		if int(err) != 0 {
			return nil, handleVipsError(out)
		}

		return out, nil
	}

	return in, nil
}

func vipsDetermineImageType(buf []byte) ImageType {
	if len(buf) < 12 {
		return ImageTypeUnknown
	}
	if buf[0] == 0xFF && buf[1] == 0xD8 && buf[2] == 0xFF {
		return ImageTypeJPEG
	}
	if IsTypeSupported(ImageTypeGIF) && buf[0] == 0x47 && buf[1] == 0x49 && buf[2] == 0x46 {
		return ImageTypeGIF
	}
	if buf[0] == 0x89 && buf[1] == 0x50 && buf[2] == 0x4E && buf[3] == 0x47 {
		return ImageTypePNG
	}
	if IsTypeSupported(ImageTypeTIFF) &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) ||
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) {
		return ImageTypeTIFF
	}
	if IsTypeSupported(ImageTypePDF) && buf[0] == 0x25 && buf[1] == 0x50 && buf[2] == 0x44 && buf[3] == 0x46 {
		return ImageTypePDF
	}
	if IsTypeSupported(ImageTypeWEBP) && buf[8] == 0x57 && buf[9] == 0x45 && buf[10] == 0x42 && buf[11] == 0x50 {
		return ImageTypeWEBP
	}
	if IsTypeSupported(ImageTypeSVG) && buf[0] == 0x3C && buf[1] == 0x3F && buf[2] == 0x78 && buf[3] == 0x6D {
		return ImageTypeSVG
	}
	// https://github.com/strukturag/libheif/blob/master/libheif/heif.cc
	if IsTypeSupported(ImageTypeHEIF) && (buf[4] == 'f' && buf[5] == 't' && buf[6] == 'y' && buf[7] == 'p') &&
		(buf[8] == 'h' && buf[9] == 'e' && buf[10] == 'i' && buf[11] == 'c') {
		return ImageTypeHEIF
	}
	return ImageTypeUnknown
}

func vipsLabel(in *C.VipsImage, params *LabelParams) (*C.VipsImage, error) {
	incOpCounter("label")
	var out *C.VipsImage

	text := C.CString(params.Text)
	defer freeCString(text)

	font := C.CString(params.Font)
	defer freeCString(font)

	color := [3]C.double{C.double(params.Color.R), C.double(params.Color.G), C.double(params.Color.B)}
	w := params.Width.GetRounded(int(in.Xsize))
	h := params.Height.GetRounded(int(in.Ysize))
	offsetX := params.OffsetX.GetRounded(int(in.Xsize))
	offsetY := params.OffsetY.GetRounded(int(in.Ysize))

	opts := vipsLabelOptions{
		Text:      text,
		Font:      font,
		Width:     C.int(w),
		Height:    C.int(h),
		OffsetX:   C.int(offsetX),
		OffsetY:   C.int(offsetY),
		Alignment: C.VipsAlign(params.Alignment),
		Opacity:   C.float(params.Opacity),
		Color:     color,
	}

	err := C.label(in, &out, (*C.LabelOptions)(unsafe.Pointer(&opts)))
	if err != 0 {
		return nil, handleVipsError(out)
	}

	return out, nil
}
