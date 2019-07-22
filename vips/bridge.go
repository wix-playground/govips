package vips

// #cgo pkg-config: vips
// #include "bridge.h"
import "C"
import (
	"fmt"
)

// ExportParams are options when exporting an image to file or buffer
type ExportParams struct {
	Format          ImageType
	Quality         int
	Compression     int
	Interlaced      bool
	Lossless        bool
	StripProfile    bool
	StripMetadata   bool
	Interpretation  Interpretation
	BackgroundColor *Color
}

func vipsExportBuffer(image *C.VipsImage, params *ExportParams) ([]byte, ImageType, error) {
	var buf []byte
	var err error

	tmpImage, err := vipsPrepareForExport(image, params)
	if err != nil {
		return nil, ImageTypeUnknown, err
	}

	// If these are equal, then we don't want to deref the original image as
	// the original will be returned if the target colorspace is not supported
	if tmpImage != image {
		defer unrefImage(tmpImage)
	}

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

	switch format {
	case ImageTypeWEBP:
		buf, err = vipsSaveWebPToBuffer(tmpImage, params.StripMetadata, params.Quality, params.Lossless)
	case ImageTypePNG:
		buf, err = vipsSavePNGToBuffer(tmpImage, params.StripMetadata, params.Compression, params.Quality, params.Interlaced)
	case ImageTypeTIFF:
		buf, err = vipsSaveTIFFToBuffer(tmpImage)
	case ImageTypeHEIF:
		buf, err = vipsSaveHEIFToBuffer(tmpImage, params.Quality, params.Lossless)
	default:
		format = ImageTypeJPEG
		buf, err = vipsSaveJPEGToBuffer(tmpImage, params.Quality, params.StripMetadata, params.Interlaced)
	}

	if err != nil {
		return nil, ImageTypeUnknown, err
	}

	return buf, format, nil
}

func vipsPrepareForExport(in *C.VipsImage, params *ExportParams) (*C.VipsImage, error) {
	if params.StripProfile {
		vipsRemoveICCProfile(in)
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
	if vipsIsColorSpaceSupported(in) && interpretation != in.Type {
		return vipsToColorSpace(in, params.Interpretation)
	}

	return in, nil
}
