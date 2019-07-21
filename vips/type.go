package vips

// #cgo pkg-config: vips
// #include "bridge.h"
import "C"

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

// Color represents an RGB
type Color struct {
	R, G, B uint8
}

// ColorBlack is shorthand for black RGB
//noinspection GoUnusedGlobalVariable
var ColorBlack = Color{0, 0, 0}

const DefaultFont = "sans 10"

// LabelParams represents a text-based label
type LabelParams struct {
	Text      string
	Font      string
	Width     Scalar
	Height    Scalar
	OffsetX   Scalar
	OffsetY   Scalar
	Opacity   float32
	Color     Color
	Alignment Align
}

// Interpolator represents the vips interpolator types
type Interpolator string

// Interpolator enum
const (
	InterpolateBicubic  Interpolator = "bicubic"
	InterpolateBilinear Interpolator = "bilinear"
	InterpolateNoHalo   Interpolator = "nohalo"
)

// String returns the canonical name of the interpolator
func (i Interpolator) String() string {
	return string(i)
}

// OperationMath represents VIPS_OPERATION_MATH type
type OperationMath int

// OperationMath enum
const (
	OperationMathSin   OperationMath = C.VIPS_OPERATION_MATH_SIN
	OperationMathCos   OperationMath = C.VIPS_OPERATION_MATH_COS
	OperationMathTan   OperationMath = C.VIPS_OPERATION_MATH_TAN
	OperationMathAsin  OperationMath = C.VIPS_OPERATION_MATH_ASIN
	OperationMathAcos  OperationMath = C.VIPS_OPERATION_MATH_ACOS
	OperationMathAtan  OperationMath = C.VIPS_OPERATION_MATH_ATAN
	OperationMathLog   OperationMath = C.VIPS_OPERATION_MATH_LOG
	OperationMathLog10 OperationMath = C.VIPS_OPERATION_MATH_LOG10
	OperationMathExp   OperationMath = C.VIPS_OPERATION_MATH_EXP
	OperationMathExp10 OperationMath = C.VIPS_OPERATION_MATH_EXP10
)

// OperationMath2 represents VIPS_OPERATION_MATH2 type
type OperationMath2 int

// OperationMath2 enum
const (
	OperationMath2Pow OperationMath2 = C.VIPS_OPERATION_MATH2_POW
	OperationMath2Wop OperationMath2 = C.VIPS_OPERATION_MATH2_WOP
)

// OperationRound represents VIPS_OPERATION_ROUND type
type OperationRound int

// OperationRound enum
const (
	OperationRoundRint  OperationRound = C.VIPS_OPERATION_ROUND_RINT
	OperationRoundCeil  OperationRound = C.VIPS_OPERATION_ROUND_CEIL
	OperationRoundFloor OperationRound = C.VIPS_OPERATION_ROUND_FLOOR
)

// OperationRelational represents VIPS_OPERATION_RELATIONAL type
type OperationRelational int

// OperationRelational enum
const (
	OperationRelationalEqual  OperationRelational = C.VIPS_OPERATION_RELATIONAL_EQUAL
	OperationRelationalNotEq  OperationRelational = C.VIPS_OPERATION_RELATIONAL_NOTEQ
	OperationRelationalLess   OperationRelational = C.VIPS_OPERATION_RELATIONAL_LESS
	OperationRelationalLessEq OperationRelational = C.VIPS_OPERATION_RELATIONAL_LESSEQ
	OperationRelationalMore   OperationRelational = C.VIPS_OPERATION_RELATIONAL_MORE
	OperationRelationalMoreEq OperationRelational = C.VIPS_OPERATION_RELATIONAL_MOREEQ
)

// OperationBoolean represents VIPS_OPERATION_BOOLEAN type
type OperationBoolean int

// OperationBoolean enum
const (
	OperationBooleanAnd    OperationBoolean = C.VIPS_OPERATION_BOOLEAN_AND
	OperationBooleanOr     OperationBoolean = C.VIPS_OPERATION_BOOLEAN_OR
	OperationBooleanEOr    OperationBoolean = C.VIPS_OPERATION_BOOLEAN_EOR
	OperationBooleanLShift OperationBoolean = C.VIPS_OPERATION_BOOLEAN_LSHIFT
	OperationBooleanRShift OperationBoolean = C.VIPS_OPERATION_BOOLEAN_RSHIFT
)

// OperationComplex represents VIPS_OPERATION_COMPLEX type
type OperationComplex int

// OperationComplex enum
const (
	OperationComplexPolar OperationComplex = C.VIPS_OPERATION_COMPLEX_POLAR
	OperationComplexRect  OperationComplex = C.VIPS_OPERATION_COMPLEX_RECT
	OperationComplexConj  OperationComplex = C.VIPS_OPERATION_COMPLEX_CONJ
)

// OperationComplex2 represents VIPS_OPERATION_COMPLEX2 type
type OperationComplex2 int

// OperationComplex2 enum
const (
	OperationComplex2CrossPhase OperationComplex2 = C.VIPS_OPERATION_COMPLEX2_CROSS_PHASE
)

// OperationComplexGet represents VIPS_OPERATION_COMPLEXGET type
type OperationComplexGet int

// OperationComplexGet enum
const (
	OperationComplexReal OperationComplexGet = C.VIPS_OPERATION_COMPLEXGET_REAL
	OperationComplexImag OperationComplexGet = C.VIPS_OPERATION_COMPLEXGET_IMAG
)

// Align represents VIPS_ALIGN
type Align int

// Direction enum
const (
	AlignLow    Align = C.VIPS_ALIGN_LOW
	AlignCenter Align = C.VIPS_ALIGN_CENTRE
	AlignHigh   Align = C.VIPS_ALIGN_HIGH
)

// Interpretation represents VIPS_INTERPRETATION type
type Interpretation int

// Interpretation enum
const (
	InterpretationError     Interpretation = C.VIPS_INTERPRETATION_ERROR
	InterpretationMultiband Interpretation = C.VIPS_INTERPRETATION_MULTIBAND
	InterpretationBW        Interpretation = C.VIPS_INTERPRETATION_B_W
	InterpretationHistogram Interpretation = C.VIPS_INTERPRETATION_HISTOGRAM
	InterpretationXYZ       Interpretation = C.VIPS_INTERPRETATION_XYZ
	InterpretationLAB       Interpretation = C.VIPS_INTERPRETATION_LAB
	InterpretationCMYK      Interpretation = C.VIPS_INTERPRETATION_CMYK
	InterpretationLABQ      Interpretation = C.VIPS_INTERPRETATION_LABQ
	InterpretationRGB       Interpretation = C.VIPS_INTERPRETATION_RGB
	InterpretationCMC       Interpretation = C.VIPS_INTERPRETATION_CMC
	InterpretationLCH       Interpretation = C.VIPS_INTERPRETATION_LCH
	InterpretationLABS      Interpretation = C.VIPS_INTERPRETATION_LABS
	InterpretationSRGB      Interpretation = C.VIPS_INTERPRETATION_sRGB
	InterpretationYXY       Interpretation = C.VIPS_INTERPRETATION_YXY
	InterpretationFourier   Interpretation = C.VIPS_INTERPRETATION_FOURIER
	InterpretationGB16      Interpretation = C.VIPS_INTERPRETATION_RGB16
	InterpretationGrey16    Interpretation = C.VIPS_INTERPRETATION_GREY16
	InterpretationMatrix    Interpretation = C.VIPS_INTERPRETATION_MATRIX
	InterpretationScRGB     Interpretation = C.VIPS_INTERPRETATION_scRGB
	InterpretationHSV       Interpretation = C.VIPS_INTERPRETATION_HSV
)

// BandFormat represents VIPS_FORMAT type
type BandFormat int

// BandFormat enum
const (
	BandFormatNotSet    BandFormat = C.VIPS_FORMAT_NOTSET
	BandFormatUchar     BandFormat = C.VIPS_FORMAT_UCHAR
	BandFormatChar      BandFormat = C.VIPS_FORMAT_CHAR
	BandFormatUshort    BandFormat = C.VIPS_FORMAT_USHORT
	BandFormatShort     BandFormat = C.VIPS_FORMAT_SHORT
	BandFormatUint      BandFormat = C.VIPS_FORMAT_UINT
	BandFormatInt       BandFormat = C.VIPS_FORMAT_INT
	BandFormatFloat     BandFormat = C.VIPS_FORMAT_FLOAT
	BandFormatComplex   BandFormat = C.VIPS_FORMAT_COMPLEX
	BandFormatDouble    BandFormat = C.VIPS_FORMAT_DOUBLE
	BandFormatDpComplex BandFormat = C.VIPS_FORMAT_DPCOMPLEX
)

// Coding represents VIPS_CODING type
type Coding int

// Coding enum
const (
	CodingError Coding = C.VIPS_CODING_ERROR
	CodingNone  Coding = C.VIPS_CODING_NONE
	CodingLABQ  Coding = C.VIPS_CODING_LABQ
	CodingRAD   Coding = C.VIPS_CODING_RAD
)

// OperationMorphology represents VIPS_OPERATION_MORPHOLOGY
type OperationMorphology int

// OperationMorphology enum
const (
	MorphologyErode  OperationMorphology = C.VIPS_OPERATION_MORPHOLOGY_ERODE
	MorphologyDilate OperationMorphology = C.VIPS_OPERATION_MORPHOLOGY_DILATE
)

type Composite struct {
	Image     *ImageRef
	BlendMode BlendMode
}

// Size represents VIPS_SIZE type
type Size int

// Size enum
const (
	SizeBoth  Size = C.VIPS_SIZE_BOTH
	SizeUp    Size = C.VIPS_SIZE_UP
	SizeDown  Size = C.VIPS_SIZE_DOWN
	SizeForce Size = C.VIPS_SIZE_FORCE
	SizeLast  Size = C.VIPS_SIZE_LAST
)
