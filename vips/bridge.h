#include <stdlib.h>
#include <vips/vips.h>
#include <vips/foreign.h>

#if (VIPS_MAJOR_VERSION < 8)
  error_requires_version_8
#endif

#define INT_TO_GBOOLEAN(bool) (bool > 0 ? TRUE : FALSE)

enum types {
	UNKNOWN = 0,
	JPEG,
	WEBP,
	PNG,
	TIFF,
	GIF,
	PDF,
	SVG,
	MAGICK,
	HEIF
};
typedef struct {
  const char *Text;
	const char *Font;
	int Width;
  int Height;
  int OffsetX;
  int OffsetY;
  VipsAlign Align;
	int DPI;
  int Margin;
	float Opacity;
  double Color[3];
} LabelOptions;


int remove_icc_profile(VipsImage *in);

// Operations
int label(VipsImage *in, VipsImage **out, LabelOptions *o);


void gobject_set_property(VipsObject* object, const char* name, const GValue* value);
