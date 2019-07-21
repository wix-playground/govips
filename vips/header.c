#include "header.h"

unsigned long has_profile_embed(VipsImage *in) {
	return vips_image_get_typeof(in, VIPS_META_ICC_NAME);
}
