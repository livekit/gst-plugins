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

type Textrender struct {
	Element *gst.Element
}

func NewTextrender() (*Textrender, error) {
	e, err := gst.NewElement("textrender")
	if err != nil {
		return nil, err
	}
	return &Textrender{Element: e}, nil
}

func NewTextrenderWithName(name string) (*Textrender, error) {
	e, err := gst.NewElementWithName("textrender", name)
	if err != nil {
		return nil, err
	}
	return &Textrender{Element: e}, nil
}

// ----- Properties -----

// font-desc (string)
//
// Pango font description of font to be used for rendering. See documentation of pango_font_description_from_string for syntax.

func (e *Textrender) GetFontDesc() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-desc")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Textrender) SetFontDesc(value string) error {
	return e.Element.SetProperty("font-desc", value)
}

// halignment (GstTextRenderHalign)
//
// Horizontal alignment of the text

func (e *Textrender) GetHalignment() (GstTextRenderHalign, error) {
	var value GstTextRenderHalign
	var ok bool
	v, err := e.Element.GetProperty("halignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTextRenderHalign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTextRenderHalign")
	}
	return value, nil
}

func (e *Textrender) SetHalignment(value GstTextRenderHalign) error {
	e.Element.SetArg("halignment", string(value))
	return nil
}

// line-alignment (GstTextRenderLineAlign)
//
// Alignment of text lines relative to each other.

func (e *Textrender) GetLineAlignment() (GstTextRenderLineAlign, error) {
	var value GstTextRenderLineAlign
	var ok bool
	v, err := e.Element.GetProperty("line-alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTextRenderLineAlign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTextRenderLineAlign")
	}
	return value, nil
}

func (e *Textrender) SetLineAlignment(value GstTextRenderLineAlign) error {
	e.Element.SetArg("line-alignment", string(value))
	return nil
}

// valignment (GstTextRenderValign)
//
// Vertical alignment of the text

func (e *Textrender) GetValignment() (GstTextRenderValign, error) {
	var value GstTextRenderValign
	var ok bool
	v, err := e.Element.GetProperty("valignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTextRenderValign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTextRenderValign")
	}
	return value, nil
}

func (e *Textrender) SetValignment(value GstTextRenderValign) error {
	e.Element.SetArg("valignment", string(value))
	return nil
}

// xpad (int)
//
// Horizontal paddding when using left/right alignment

func (e *Textrender) GetXpad() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xpad")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Textrender) SetXpad(value int) error {
	return e.Element.SetProperty("xpad", value)
}

// ypad (int)
//
// Vertical padding when using top/bottom alignment

func (e *Textrender) GetYpad() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ypad")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Textrender) SetYpad(value int) error {
	return e.Element.SetProperty("ypad", value)
}

// ----- Constants -----

type GstTextRenderHalign string

const (
	GstTextRenderHalignLeft GstTextRenderHalign = "left" // left (0) – left
	GstTextRenderHalignCenter GstTextRenderHalign = "center" // center (1) – center
	GstTextRenderHalignRight GstTextRenderHalign = "right" // right (2) – right
)

type GstTextRenderLineAlign string

const (
	GstTextRenderLineAlignLeft GstTextRenderLineAlign = "left" // left (0) – left
	GstTextRenderLineAlignCenter GstTextRenderLineAlign = "center" // center (1) – center
	GstTextRenderLineAlignRight GstTextRenderLineAlign = "right" // right (2) – right
)

type GstTextRenderValign string

const (
	GstTextRenderValignBaseline GstTextRenderValign = "baseline" // baseline (0) – baseline
	GstTextRenderValignBottom GstTextRenderValign = "bottom" // bottom (1) – bottom
	GstTextRenderValignTop GstTextRenderValign = "top" // top (2) – top
)

