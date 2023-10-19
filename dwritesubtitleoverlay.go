// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type Dwritesubtitleoverlay struct {
	Element *gst.Element
}

func NewDwritesubtitleoverlay() (*Dwritesubtitleoverlay, error) {
	e, err := gst.NewElement("dwritesubtitleoverlay")
	if err != nil {
		return nil, err
	}
	return &Dwritesubtitleoverlay{Element: e}, nil
}

func NewDwritesubtitleoverlayWithName(name string) (*Dwritesubtitleoverlay, error) {
	e, err := gst.NewElementWithName("dwritesubtitleoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Dwritesubtitleoverlay{Element: e}, nil
}

// ----- Properties -----

// auto-resize (bool)
//
// Calculate font size to be equivalent to "font-size" at "reference-frame-size"

func (e *Dwritesubtitleoverlay) GetAutoResize() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-resize")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetAutoResize(value bool) error {
	return e.Element.SetProperty("auto-resize", value)
}

// background-color (uint)
//
// Background color to use (big-endian ARGB)

func (e *Dwritesubtitleoverlay) GetBackgroundColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("background-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetBackgroundColor(value uint) error {
	return e.Element.SetProperty("background-color", value)
}

// cc-field (int)
//
// The closed caption field to render when available, (-1 = automatic)

func (e *Dwritesubtitleoverlay) GetCcField() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cc-field")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetCcField(value int) error {
	return e.Element.SetProperty("cc-field", value)
}

// cc-timeout (uint64)
//
// Duration after which to erase overlay when no cc data has arrived for the selected field, in nanoseconds unit

func (e *Dwritesubtitleoverlay) GetCcTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("cc-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetCcTimeout(value uint64) error {
	return e.Element.SetProperty("cc-timeout", value)
}

// color-font (bool)
//
// Enable color font, requires Windows 10 or newer

func (e *Dwritesubtitleoverlay) GetColorFont() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("color-font")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetColorFont(value bool) error {
	return e.Element.SetProperty("color-font", value)
}

// enable-cc (bool)
//
// Enable closed caption rendering

func (e *Dwritesubtitleoverlay) GetEnableCc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-cc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetEnableCc(value bool) error {
	return e.Element.SetProperty("enable-cc", value)
}

// font-family (string)
//
// Font family to use

func (e *Dwritesubtitleoverlay) GetFontFamily() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-family")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetFontFamily(value string) error {
	return e.Element.SetProperty("font-family", value)
}

// font-size (float32)
//
// Font size to use

func (e *Dwritesubtitleoverlay) GetFontSize() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("font-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetFontSize(value float32) error {
	return e.Element.SetProperty("font-size", value)
}

// font-stretch (GstDWriteFontStretch)
//
// Font Stretch

func (e *Dwritesubtitleoverlay) GetFontStretch() (interface{}, error) {
	return e.Element.GetProperty("font-stretch")
}

func (e *Dwritesubtitleoverlay) SetFontStretch(value interface{}) error {
	return e.Element.SetProperty("font-stretch", value)
}

// font-style (GstDWriteFontStyle)
//
// Font Style

func (e *Dwritesubtitleoverlay) GetFontStyle() (interface{}, error) {
	return e.Element.GetProperty("font-style")
}

func (e *Dwritesubtitleoverlay) SetFontStyle(value interface{}) error {
	return e.Element.SetProperty("font-style", value)
}

// font-weight (GstDWriteFontWeight)
//
// Font Weight

func (e *Dwritesubtitleoverlay) GetFontWeight() (interface{}, error) {
	return e.Element.GetProperty("font-weight")
}

func (e *Dwritesubtitleoverlay) SetFontWeight(value interface{}) error {
	return e.Element.SetProperty("font-weight", value)
}

// foreground-color (uint)
//
// Foreground color to use (big-endian ARGB)

func (e *Dwritesubtitleoverlay) GetForegroundColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("foreground-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetForegroundColor(value uint) error {
	return e.Element.SetProperty("foreground-color", value)
}

// layout-height (float64)
//
// Normalized height of text layout

func (e *Dwritesubtitleoverlay) GetLayoutHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetLayoutHeight(value float64) error {
	return e.Element.SetProperty("layout-height", value)
}

// layout-width (float64)
//
// Normalized width of text layout

func (e *Dwritesubtitleoverlay) GetLayoutWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetLayoutWidth(value float64) error {
	return e.Element.SetProperty("layout-width", value)
}

// layout-x (float64)
//
// Normalized X coordinate of text layout

func (e *Dwritesubtitleoverlay) GetLayoutX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetLayoutX(value float64) error {
	return e.Element.SetProperty("layout-x", value)
}

// layout-y (float64)
//
// Normalized Y coordinate of text layout

func (e *Dwritesubtitleoverlay) GetLayoutY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetLayoutY(value float64) error {
	return e.Element.SetProperty("layout-y", value)
}

// outline-color (uint)
//
// Text outline color to use (big-endian ARGB)

func (e *Dwritesubtitleoverlay) GetOutlineColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("outline-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetOutlineColor(value uint) error {
	return e.Element.SetProperty("outline-color", value)
}

// paragraph-alignment (GstDWriteParagraphAlignment)
//
// Paragraph Alignment

func (e *Dwritesubtitleoverlay) GetParagraphAlignment() (interface{}, error) {
	return e.Element.GetProperty("paragraph-alignment")
}

func (e *Dwritesubtitleoverlay) SetParagraphAlignment(value interface{}) error {
	return e.Element.SetProperty("paragraph-alignment", value)
}

// remove-cc-meta (bool)
//
// Remove caption meta from output buffers when closed caption rendering is enabled

func (e *Dwritesubtitleoverlay) GetRemoveCcMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-cc-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetRemoveCcMeta(value bool) error {
	return e.Element.SetProperty("remove-cc-meta", value)
}

// shadow-color (uint)
//
// Shadow color to use (big-endian ARGB)

func (e *Dwritesubtitleoverlay) GetShadowColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shadow-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetShadowColor(value uint) error {
	return e.Element.SetProperty("shadow-color", value)
}

// text (string)
//
// Text to render

func (e *Dwritesubtitleoverlay) GetText() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("text")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetText(value string) error {
	return e.Element.SetProperty("text", value)
}

// text-alignment (GstDWriteTextAlignment)
//
// Text Alignment

func (e *Dwritesubtitleoverlay) GetTextAlignment() (interface{}, error) {
	return e.Element.GetProperty("text-alignment")
}

func (e *Dwritesubtitleoverlay) SetTextAlignment(value interface{}) error {
	return e.Element.SetProperty("text-alignment", value)
}

// visible (bool)
//
// Whether to draw text

func (e *Dwritesubtitleoverlay) GetVisible() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("visible")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritesubtitleoverlay) SetVisible(value bool) error {
	return e.Element.SetProperty("visible", value)
}

