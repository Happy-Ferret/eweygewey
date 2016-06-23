// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package eweygewey

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

// constants used for polling the state of a mouse button
const (
	MouseDown        = 0
	MouseUp          = 1
	MouseClick       = 2
	MouseDoubleClick = 4
)

// Style defines parameters to the drawing functions that control the way
// the widgets are organized and drawn.
type Style struct {
	ButtonColor          mgl.Vec4 // button background color
	ButtonHoverColor     mgl.Vec4 // button background color with mouse hovering
	ButtonActiveColor    mgl.Vec4 // button background color when clicked
	ButtonTextColor      mgl.Vec4 // button text color
	ButtonMargin         mgl.Vec4 // [left,right,top,bottom] margin values for buttons
	ButtonPadding        mgl.Vec4 // [left,right,top,bottom] padding values for buttons
	FontName             string   // font name to use by default
	ImageMargin          mgl.Vec4 // margin for the image widgets
	ScrollBarCursorColor mgl.Vec4 // the color of the cursor of the scroll bar
	ScrollBarBgColor     mgl.Vec4 // the color of the background of the scroll bar
	ScrollBarWidth       float32  // the width of the scroll bar
	ScrollBarCursorWidth float32  // the width of the scroll bar cursor
	SeparatorColor       mgl.Vec4 // the color of the separator bar
	SeparatorHeight      float32  // the height of the separator rectangle
	SeparatorMargin      mgl.Vec4 // the margin for the separator rectangle
	SliderBgColor        mgl.Vec4 // slider background color
	SliderCursorColor    mgl.Vec4 // slider cursor color
	SliderFloatFormat    string   // formatting string for the float value in a slider
	SliderIntFormat      string   // formatting string for the int value in a slider
	SliderMargin         mgl.Vec4 // margin for the slider text strings
	SliderPadding        mgl.Vec4 // padding for the slider text strings
	SliderTextColor      mgl.Vec4 // slider text color
	SliderCursorWidth    float32  // slider cursor width
	TextColor            mgl.Vec4 // text color
	TextMargin           mgl.Vec4 // margin for text widgets
	TitleBarTextColor    mgl.Vec4 // text color
	TitleBarBgColor      mgl.Vec4 // window background color
	WindowBgColor        mgl.Vec4 // window background color
	WindowPadding        mgl.Vec4 // [left,right,top,bottom] padding values for windows
}

var (
	// VertShader330 is the GLSL vertex shader program for the user interface.
	VertShader330 = `#version 330
  uniform mat4 VIEW;
  in vec2 VERTEX_POSITION;
  in vec2 VERTEX_UV;
  in float VERTEX_TEXTURE_INDEX;
  in vec4 VERTEX_COLOR;
  out vec2 vs_uv;
  out vec4 vs_color;
  out float vs_tex_index;
  void main()
  {
    vs_uv = VERTEX_UV;
    vs_color = VERTEX_COLOR;
	vs_tex_index = VERTEX_TEXTURE_INDEX;
    gl_Position = VIEW * vec4(VERTEX_POSITION, 0.0, 1.0);
  }`

	// FragShader330 is the GLSL fragment shader program for the user interface.
	// NOTE: 4 samplers is a hardcoded value now, but there's no reason it has to be that specifically.
	FragShader330 = `#version 330
  uniform sampler2D TEX[4];
  in vec2 vs_uv;
  in vec4 vs_color;
  in float vs_tex_index;
  out vec4 frag_color;
  void main()
  {
  	int i = int(vs_tex_index);
	frag_color = vs_color * texture(TEX[i], vs_uv).rgba;
  }`

	// DefaultStyle is the default style to use for drawing widgets
	DefaultStyle = Style{
		ButtonColor:          ColorIToV(171, 102, 102, 153),
		ButtonActiveColor:    ColorIToV(204, 128, 120, 255),
		ButtonHoverColor:     ColorIToV(171, 102, 102, 255),
		ButtonTextColor:      ColorIToV(230, 230, 230, 255),
		ButtonMargin:         mgl.Vec4{2, 2, 4, 4},
		ButtonPadding:        mgl.Vec4{4, 4, 2, 2},
		FontName:             "Default",
		ImageMargin:          mgl.Vec4{0, 0, 0, 0},
		ScrollBarCursorColor: ColorIToV(102, 102, 204, 77),
		ScrollBarBgColor:     ColorIToV(51, 64, 77, 153),
		ScrollBarWidth:       16.0,
		ScrollBarCursorWidth: 10.0,
		SeparatorColor:       ColorIToV(230, 230, 230, 255),
		SeparatorHeight:      1.0,
		SeparatorMargin:      mgl.Vec4{4, 4, 4, 4},
		SliderBgColor:        ColorIToV(128, 128, 128, 179),
		SliderCursorColor:    ColorIToV(179, 179, 179, 179),
		SliderFloatFormat:    "%0.3f",
		SliderIntFormat:      "%d",
		SliderMargin:         mgl.Vec4{2, 2, 4, 4},
		SliderPadding:        mgl.Vec4{2, 2, 2, 2},
		SliderTextColor:      ColorIToV(230, 230, 230, 255),
		SliderCursorWidth:    15.0,
		TextMargin:           mgl.Vec4{0, 0, 0, 0},
		TextColor:            ColorIToV(230, 230, 230, 255),
		TitleBarTextColor:    ColorIToV(230, 230, 230, 255),
		TitleBarBgColor:      ColorIToV(69, 69, 138, 255),
		WindowBgColor:        ColorIToV(0, 0, 0, 179),
		WindowPadding:        mgl.Vec4{4, 4, 4, 4},
	}
)

// ColorIToV takes the color parameters as integers and returns them
// as a float vector.
func ColorIToV(r, g, b, a int) mgl.Vec4 {
	return mgl.Vec4{float32(r) / 255.0, float32(g) / 255.0, float32(b) / 255.0, float32(a) / 255.0}
}

// ClipF32 returns a value that is between the closed interval of [min .. max].
func ClipF32(min, max, value float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
