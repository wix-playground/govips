package vips

// #cgo pkg-config: vips
// #include "bridge.h"
import "C"

// todo: move to bridge
// todo: replace with direct C.<method> invocation
// todo: make all private, callable from ImageRef

// Abs executes the 'abs' operation
func Abs(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("abs")
	err = vipsCall("abs", options)
	return out, err
}

// Add executes the 'add' operation
func Add(left *C.VipsImage, right *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("left", left),
		InputImage("right", right),
		OutputImage("out", &out),
	)
	incOpCounter("add")
	err = vipsCall("add", options)
	return out, err
}

// Avg executes the 'avg' operation
func Avg(in *C.VipsImage, options ...*Option) (float64, error) {
	var out float64
	var err error
	options = append(options,
		InputImage("in", in),
		OutputDouble("out", &out),
	)
	incOpCounter("avg")
	err = vipsCall("avg", options)
	return out, err
}

// Bandbool executes the 'bandbool' operation
func Bandbool(in *C.VipsImage, boolean OperationBoolean, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("boolean", int(boolean)),
		OutputImage("out", &out),
	)
	incOpCounter("bandbool")
	err = vipsCall("bandbool", options)
	return out, err
}

// Bandfold executes the 'bandfold' operation
func Bandfold(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("bandfold")
	err = vipsCall("bandfold", options)
	return out, err
}

// Bandmean executes the 'bandmean' operation
func Bandmean(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("bandmean")
	err = vipsCall("bandmean", options)
	return out, err
}

// Bandunfold executes the 'bandunfold' operation
func Bandunfold(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("bandunfold")
	err = vipsCall("bandunfold", options)
	return out, err
}

// Buildlut executes the 'buildlut' operation
func Buildlut(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("buildlut")
	err = vipsCall("buildlut", options)
	return out, err
}

// Byteswap executes the 'byteswap' operation
func Byteswap(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("byteswap")
	err = vipsCall("byteswap", options)
	return out, err
}

// Cache executes the 'cache' operation
func Cache(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("cache")
	err = vipsCall("cache", options)
	return out, err
}

// Cast executes the 'cast' operation
func Cast(in *C.VipsImage, format BandFormat, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("format", int(format)),
		OutputImage("out", &out),
	)
	incOpCounter("cast")
	err = vipsCall("cast", options)
	return out, err
}

// Cmc2Lch executes the 'CMC2LCh' operation
func Cmc2Lch(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("CMC2LCh")
	err = vipsCall("CMC2LCh", options)
	return out, err
}

// Colourspace executes the 'colourspace' operation
func Colourspace(in *C.VipsImage, space Interpretation, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("space", int(space)),
		OutputImage("out", &out),
	)
	incOpCounter("colourspace")
	err = vipsCall("colourspace", options)
	return out, err
}

// Complex executes the 'complex' operation
func Complex(in *C.VipsImage, cmplx OperationComplex, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("cmplx", int(cmplx)),
		OutputImage("out", &out),
	)
	incOpCounter("complex")
	err = vipsCall("complex", options)
	return out, err
}

// Complexget executes the 'complexget' operation
func Complexget(in *C.VipsImage, get OperationComplexGet, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("get", int(get)),
		OutputImage("out", &out),
	)
	incOpCounter("complexget")
	err = vipsCall("complexget", options)
	return out, err
}

// Falsecolour executes the 'falsecolour' operation
func Falsecolour(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("falsecolour")
	err = vipsCall("falsecolour", options)
	return out, err
}

// FillNearest executes the 'fill_nearest' operation
func FillNearest(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("fill_nearest")
	err = vipsCall("fill_nearest", options)
	return out, err
}

// Float2Rad executes the 'float2rad' operation
func Float2Rad(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("float2rad")
	err = vipsCall("float2rad", options)
	return out, err
}

// Fwfft executes the 'fwfft' operation
func Fwfft(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("fwfft")
	err = vipsCall("fwfft", options)
	return out, err
}

// Gamma executes the 'gamma' operation
func Gamma(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("gamma")
	err = vipsCall("gamma", options)
	return out, err
}

// Globalbalance executes the 'globalbalance' operation
func Globalbalance(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("globalbalance")
	err = vipsCall("globalbalance", options)
	return out, err
}

// Grid executes the 'grid' operation
func Grid(in *C.VipsImage, tileHeight int, across int, down int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("tile-height", tileHeight),
		InputInt("across", across),
		InputInt("down", down),
		OutputImage("out", &out),
	)
	incOpCounter("grid")
	err = vipsCall("grid", options)
	return out, err
}

// HistCum executes the 'hist_cum' operation
func HistCum(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hist_cum")
	err = vipsCall("hist_cum", options)
	return out, err
}

// HistEqual executes the 'hist_equal' operation
func HistEqual(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hist_equal")
	err = vipsCall("hist_equal", options)
	return out, err
}

// HistFind executes the 'hist_find' operation
func HistFind(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hist_find")
	err = vipsCall("hist_find", options)
	return out, err
}

// HistFindNdim executes the 'hist_find_ndim' operation
func HistFindNdim(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hist_find_ndim")
	err = vipsCall("hist_find_ndim", options)
	return out, err
}

// HistLocal executes the 'hist_local' operation
func HistLocal(in *C.VipsImage, width int, height int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("width", width),
		InputInt("height", height),
		OutputImage("out", &out),
	)
	incOpCounter("hist_local")
	err = vipsCall("hist_local", options)
	return out, err
}

// HistNorm executes the 'hist_norm' operation
func HistNorm(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hist_norm")
	err = vipsCall("hist_norm", options)
	return out, err
}

// HistPlot executes the 'hist_plot' operation
func HistPlot(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hist_plot")
	err = vipsCall("hist_plot", options)
	return out, err
}

// HoughCircle executes the 'hough_circle' operation
func HoughCircle(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hough_circle")
	err = vipsCall("hough_circle", options)
	return out, err
}

// HoughLine executes the 'hough_line' operation
func HoughLine(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("hough_line")
	err = vipsCall("hough_line", options)
	return out, err
}

// Hsv2Srgb executes the 'HSV2sRGB' operation
func Hsv2Srgb(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("HSV2sRGB")
	err = vipsCall("HSV2sRGB", options)
	return out, err
}

// IccExport executes the 'icc_export' operation
func IccExport(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("icc_export")
	err = vipsCall("icc_export", options)
	return out, err
}

// IccImport executes the 'icc_import' operation
func IccImport(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("icc_import")
	err = vipsCall("icc_import", options)
	return out, err
}

// IccTransform executes the 'icc_transform' operation
func IccTransform(in *C.VipsImage, outputProfile string, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputString("output-profile", outputProfile),
		OutputImage("out", &out),
	)
	incOpCounter("icc_transform")
	err = vipsCall("icc_transform", options)
	return out, err
}

// Invertlut executes the 'invertlut' operation
func Invertlut(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("invertlut")
	err = vipsCall("invertlut", options)
	return out, err
}

// Invfft executes the 'invfft' operation
func Invfft(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("invfft")
	err = vipsCall("invfft", options)
	return out, err
}

// Lab2Labq executes the 'Lab2LabQ' operation
func Lab2Labq(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("Lab2LabQ")
	err = vipsCall("Lab2LabQ", options)
	return out, err
}

// Lab2Labs executes the 'Lab2LabS' operation
func Lab2Labs(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("Lab2LabS")
	err = vipsCall("Lab2LabS", options)
	return out, err
}

// Lab2Lch executes the 'Lab2LCh' operation
func Lab2Lch(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("Lab2LCh")
	err = vipsCall("Lab2LCh", options)
	return out, err
}

// Lab2Xyz executes the 'Lab2XYZ' operation
func Lab2Xyz(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("Lab2XYZ")
	err = vipsCall("Lab2XYZ", options)
	return out, err
}

// Labelregions executes the 'labelregions' operation
func Labelregions(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var mask *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("mask", &mask),
	)
	incOpCounter("labelregions")
	err = vipsCall("labelregions", options)
	return mask, err
}

// Labq2Lab executes the 'LabQ2Lab' operation
func Labq2Lab(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("LabQ2Lab")
	err = vipsCall("LabQ2Lab", options)
	return out, err
}

// Labq2Labs executes the 'LabQ2LabS' operation
func Labq2Labs(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("LabQ2LabS")
	err = vipsCall("LabQ2LabS", options)
	return out, err
}

// Labq2Srgb executes the 'LabQ2sRGB' operation
func Labq2Srgb(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("LabQ2sRGB")
	err = vipsCall("LabQ2sRGB", options)
	return out, err
}

// Labs2Lab executes the 'LabS2Lab' operation
func Labs2Lab(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("LabS2Lab")
	err = vipsCall("LabS2Lab", options)
	return out, err
}

// Labs2Labq executes the 'LabS2LabQ' operation
func Labs2Labq(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("LabS2LabQ")
	err = vipsCall("LabS2LabQ", options)
	return out, err
}

// Lch2Cmc executes the 'LCh2CMC' operation
func Lch2Cmc(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("LCh2CMC")
	err = vipsCall("LCh2CMC", options)
	return out, err
}

// Lch2Lab executes the 'LCh2Lab' operation
func Lch2Lab(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("LCh2Lab")
	err = vipsCall("LCh2Lab", options)
	return out, err
}

// Linecache executes the 'linecache' operation
func Linecache(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("linecache")
	err = vipsCall("linecache", options)
	return out, err
}

// Math executes the 'math' operation
func Math(in *C.VipsImage, math OperationMath, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("math", int(math)),
		OutputImage("out", &out),
	)
	incOpCounter("math")
	err = vipsCall("math", options)
	return out, err
}

// Measure executes the 'measure' operation
func Measure(in *C.VipsImage, h int, v int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("h", h),
		InputInt("v", v),
		OutputImage("out", &out),
	)
	incOpCounter("measure")
	err = vipsCall("measure", options)
	return out, err
}

// Msb executes the 'msb' operation
func Msb(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("msb")
	err = vipsCall("msb", options)
	return out, err
}

// Premultiply executes the 'premultiply' operation
func Premultiply(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("premultiply")
	err = vipsCall("premultiply", options)
	return out, err
}

// Rad2Float executes the 'rad2float' operation
func Rad2Float(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("rad2float")
	err = vipsCall("rad2float", options)
	return out, err
}

// Rank executes the 'rank' operation
func Rank(in *C.VipsImage, width int, height int, index int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("width", width),
		InputInt("height", height),
		InputInt("index", index),
		OutputImage("out", &out),
	)
	incOpCounter("rank")
	err = vipsCall("rank", options)
	return out, err
}

// Reduce executes the 'reduce' operation
func Reduce(in *C.VipsImage, hshrink float64, vshrink float64, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputDouble("hshrink", hshrink),
		InputDouble("vshrink", vshrink),
		OutputImage("out", &out),
	)
	incOpCounter("reduce")
	err = vipsCall("reduce", options)
	return out, err
}

// Reduceh executes the 'reduceh' operation
func Reduceh(in *C.VipsImage, hshrink float64, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputDouble("hshrink", hshrink),
		OutputImage("out", &out),
	)
	incOpCounter("reduceh")
	err = vipsCall("reduceh", options)
	return out, err
}

// Reducev executes the 'reducev' operation
func Reducev(in *C.VipsImage, vshrink float64, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputDouble("vshrink", vshrink),
		OutputImage("out", &out),
	)
	incOpCounter("reducev")
	err = vipsCall("reducev", options)
	return out, err
}

// Replicate executes the 'replicate' operation
func Replicate(in *C.VipsImage, across int, down int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("across", across),
		InputInt("down", down),
		OutputImage("out", &out),
	)
	incOpCounter("replicate")
	err = vipsCall("replicate", options)
	return out, err
}

// Rotate45 executes the 'rot45' operation
func Rot45(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("rot45")
	err = vipsCall("rot45", options)
	return out, err
}

// Round executes the 'round' operation
func Round(in *C.VipsImage, round OperationRound, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("round", int(round)),
		OutputImage("out", &out),
	)
	incOpCounter("round")
	err = vipsCall("round", options)
	return out, err
}

// Scale executes the 'scale' operation
func Scale(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("scale")
	err = vipsCall("scale", options)
	return out, err
}

// Scrgb2Bw executes the 'scRGB2BW' operation
func Scrgb2Bw(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("scRGB2BW")
	err = vipsCall("scRGB2BW", options)
	return out, err
}

// Scrgb2Srgb executes the 'scRGB2sRGB' operation
func Scrgb2Srgb(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("scRGB2sRGB")
	err = vipsCall("scRGB2sRGB", options)
	return out, err
}

// Scrgb2Xyz executes the 'scRGB2XYZ' operation
func Scrgb2Xyz(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("scRGB2XYZ")
	err = vipsCall("scRGB2XYZ", options)
	return out, err
}

// Sequential executes the 'sequential' operation
func Sequential(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("sequential")
	err = vipsCall("sequential", options)
	return out, err
}

// Shrink executes the 'shrink' operation
func Shrink(in *C.VipsImage, hshrink float64, vshrink float64, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputDouble("hshrink", hshrink),
		InputDouble("vshrink", vshrink),
		OutputImage("out", &out),
	)
	incOpCounter("shrink")
	err = vipsCall("shrink", options)
	return out, err
}

// Shrinkh executes the 'shrinkh' operation
func Shrinkh(in *C.VipsImage, hshrink int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("hshrink", hshrink),
		OutputImage("out", &out),
	)
	incOpCounter("shrinkh")
	err = vipsCall("shrinkh", options)
	return out, err
}

// Shrinkv executes the 'shrinkv' operation
func Shrinkv(in *C.VipsImage, vshrink int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("vshrink", vshrink),
		OutputImage("out", &out),
	)
	incOpCounter("shrinkv")
	err = vipsCall("shrinkv", options)
	return out, err
}

// Sign executes the 'sign' operation
func Sign(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("sign")
	err = vipsCall("sign", options)
	return out, err
}

// Similarity executes the 'similarity' operation
func Similarity(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("similarity")
	err = vipsCall("similarity", options)
	return out, err
}

// Smartcrop executes the 'smartcrop' operation
func Smartcrop(input *C.VipsImage, width int, height int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("input", input),
		InputInt("width", width),
		InputInt("height", height),
		OutputImage("out", &out),
	)
	incOpCounter("smartcrop")
	err = vipsCall("smartcrop", options)
	return out, err
}

// Spectrum executes the 'spectrum' operation
func Spectrum(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("spectrum")
	err = vipsCall("spectrum", options)
	return out, err
}

// Srgb2Hsv executes the 'sRGB2HSV' operation
func Srgb2Hsv(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("sRGB2HSV")
	err = vipsCall("sRGB2HSV", options)
	return out, err
}

// Srgb2Scrgb executes the 'sRGB2scRGB' operation
func Srgb2Scrgb(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("sRGB2scRGB")
	err = vipsCall("sRGB2scRGB", options)
	return out, err
}

// Stats executes the 'stats' operation
func Stats(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("stats")
	err = vipsCall("stats", options)
	return out, err
}

// Stdif executes the 'stdif' operation
func Stdif(in *C.VipsImage, width int, height int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("width", width),
		InputInt("height", height),
		OutputImage("out", &out),
	)
	incOpCounter("stdif")
	err = vipsCall("stdif", options)
	return out, err
}

// Subsample executes the 'subsample' operation
func Subsample(input *C.VipsImage, xfac int, yfac int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("input", input),
		InputInt("xfac", xfac),
		InputInt("yfac", yfac),
		OutputImage("out", &out),
	)
	incOpCounter("subsample")
	err = vipsCall("subsample", options)
	return out, err
}

// ThumbnailImage executes the 'thumbnail_image' operation
func ThumbnailImage(in *C.VipsImage, width int, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		InputInt("width", width),
		OutputImage("out", &out),
	)
	incOpCounter("thumbnail_image")
	err = vipsCall("thumbnail_image", options)
	return out, err
}

// Tilecache executes the 'tilecache' operation
func Tilecache(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("tilecache")
	err = vipsCall("tilecache", options)
	return out, err
}

// Unpremultiply executes the 'unpremultiply' operation
func Unpremultiply(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("unpremultiply")
	err = vipsCall("unpremultiply", options)
	return out, err
}

// Wrap executes the 'wrap' operation
func Wrap(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("wrap")
	err = vipsCall("wrap", options)
	return out, err
}

// Xyz2Lab executes the 'XYZ2Lab' operation
func Xyz2Lab(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("XYZ2Lab")
	err = vipsCall("XYZ2Lab", options)
	return out, err
}

// Xyz2Scrgb executes the 'XYZ2scRGB' operation
func Xyz2Scrgb(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("XYZ2scRGB")
	err = vipsCall("XYZ2scRGB", options)
	return out, err
}

// Xyz2Yxy executes the 'XYZ2Yxy' operation
func Xyz2Yxy(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("XYZ2Yxy")
	err = vipsCall("XYZ2Yxy", options)
	return out, err
}

// Yxy2Xyz executes the 'Yxy2XYZ' operation
func Yxy2Xyz(in *C.VipsImage, options ...*Option) (*C.VipsImage, error) {
	var out *C.VipsImage
	var err error
	options = append(options,
		InputImage("in", in),
		OutputImage("out", &out),
	)
	incOpCounter("Yxy2XYZ")
	err = vipsCall("Yxy2XYZ", options)
	return out, err
}
