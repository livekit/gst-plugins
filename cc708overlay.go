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

type Cc708Overlay struct {
	Element *gst.Element
}

func NewCc708Overlay() (*Cc708Overlay, error) {
	e, err := gst.NewElement("cc708overlay")
	if err != nil {
		return nil, err
	}
	return &Cc708Overlay{Element: e}, nil
}

func NewCc708OverlayWithName(name string) (*Cc708Overlay, error) {
	e, err := gst.NewElementWithName("cc708overlay", name)
	if err != nil {
		return nil, err
	}
	return &Cc708Overlay{Element: e}, nil
}

// ----- Properties -----

// font-desc (string)
//
// Pango font description of font to be used for rendering.
// See documentation of pango_font_description_from_string for syntax.
// this will override closed caption stream specified font style/pen size.

func (e *Cc708Overlay) GetFontDesc() (string, error) {
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

func (e *Cc708Overlay) SetFontDesc(value string) error {
	return e.Element.SetProperty("font-desc", value)
}

// service-number (int)
//
// Service number. Service 1 is designated as the Primary Caption Service, Service 2 is the Secondary Language Service.

func (e *Cc708Overlay) GetServiceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("service-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cc708Overlay) SetServiceNumber(value int) error {
	return e.Element.SetProperty("service-number", value)
}

// silent (bool)
//
// If set, no text is rendered. Useful to switch off text rendering
// temporarily without removing the textoverlay element from the pipeline.

func (e *Cc708Overlay) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Cc708Overlay) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// window-h-pos (GstCeaCcOverlayWinHpos)
//
// Window's Horizontal position

func (e *Cc708Overlay) GetWindowHPos() (GstCeaCcOverlayWinHpos, error) {
	var value GstCeaCcOverlayWinHpos
	var ok bool
	v, err := e.Element.GetProperty("window-h-pos")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCeaCcOverlayWinHpos)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCeaCcOverlayWinHpos")
	}
	return value, nil
}

func (e *Cc708Overlay) SetWindowHPos(value GstCeaCcOverlayWinHpos) error {
	e.Element.SetArg("window-h-pos", string(value))
	return nil
}

// ----- Constants -----

type GstCeaCcOverlayWinHpos string

const (
	GstCeaCcOverlayWinHposLeft GstCeaCcOverlayWinHpos = "left" // left (0) – left
	GstCeaCcOverlayWinHposCenter GstCeaCcOverlayWinHpos = "center" // center (1) – center
	GstCeaCcOverlayWinHposRight GstCeaCcOverlayWinHpos = "right" // right (2) – right
	GstCeaCcOverlayWinHposAuto GstCeaCcOverlayWinHpos = "auto" // auto (3) – auto
)

