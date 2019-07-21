#include "bridge.h"


int is_16bit(VipsInterpretation interpretation) {
	return interpretation == VIPS_INTERPRETATION_RGB16 || interpretation == VIPS_INTERPRETATION_GREY16;
}

int text(VipsImage **out, const char *text, const char *font, int width, int height, VipsAlign align, int dpi) {
	return vips_text(out, text, "font", font, "width", width, "height", height, "align", align, "dpi", dpi, NULL);
}

void gobject_set_property(VipsObject *object, const char *name, const GValue *value) {
  VipsObjectClass *object_class = VIPS_OBJECT_GET_CLASS( object );
  GType type = G_VALUE_TYPE( value );

  GParamSpec *pspec;
  VipsArgumentClass *argument_class;
  VipsArgumentInstance *argument_instance;

  if( vips_object_get_argument( object, name,
    &pspec, &argument_class, &argument_instance ) ) {
    g_warning( "gobject warning: %s", vips_error_buffer() );
    vips_error_clear();
    return;
  }

  if( G_IS_PARAM_SPEC_ENUM( pspec ) &&
    type == G_TYPE_STRING ) {
    GType pspec_type = G_PARAM_SPEC_VALUE_TYPE( pspec );

    int enum_value;
    GValue value2 = { 0 };

    if( (enum_value = vips_enum_from_nick( object_class->nickname,
      pspec_type, g_value_get_string( value ) )) < 0 ) {
      g_warning( "gobject warning: %s", vips_error_buffer() );
      vips_error_clear();
      return;
    }

    g_value_init( &value2, pspec_type );
    g_value_set_enum( &value2, enum_value );
    g_object_set_property( G_OBJECT( object ), name, &value2 );
    g_value_unset( &value2 );
  } else {
    g_object_set_property( G_OBJECT( object ), name, value );
  }
}

int label(VipsImage *in, VipsImage **out, LabelOptions *o) {
	double ones[3] = { 1, 1, 1 };
	VipsImage *base = vips_image_new();
	VipsImage **t = (VipsImage **) vips_object_local_array(VIPS_OBJECT(base), 10);
	t[0] = in;
	if (
		vips_text(&t[1], o->Text,
			"font", o->Font,
			"width", o->Width,
			"height", o->Height,
			"align", o->Align,
			NULL) ||
		vips_linear1(t[1], &t[2], o->Opacity, 0.0, NULL) ||
		vips_cast(t[2], &t[3], VIPS_FORMAT_UCHAR, NULL) ||
		vips_embed(t[3], &t[4], o->OffsetX, o->OffsetY, t[3]->Xsize + o->OffsetX, t[3]->Ysize + o->OffsetY, NULL)
		) {
		g_object_unref(base);
		return 1;
	}
	if (
		vips_black(&t[5], 1, 1, NULL) ||
		vips_linear(t[5], &t[6], ones, o->Color, 3, NULL) ||
		vips_cast(t[6], &t[7], VIPS_FORMAT_UCHAR, NULL) ||
		vips_copy(t[7], &t[8], "interpretation", t[0]->Type, NULL) ||
		vips_embed(t[8], &t[9], 0, 0, t[0]->Xsize, t[0]->Ysize, "extend", VIPS_EXTEND_COPY, NULL)
		) {
		g_object_unref(base);
		return 1;
	}
	if (vips_ifthenelse(t[4], t[9], t[0], out, "blend", TRUE, NULL)) {
		g_object_unref(base);
		return 1;
	}
	g_object_unref(base);
	return 0;
}




int remove_icc_profile(VipsImage *in) {
  return vips_image_remove(in, VIPS_META_ICC_NAME);
}
