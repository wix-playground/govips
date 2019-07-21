#include <stdlib.h>
#include <vips/vips.h>
#include <vips/foreign.h>

#if (VIPS_MAJOR_VERSION < 8)
  error_requires_version_8
#endif

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
