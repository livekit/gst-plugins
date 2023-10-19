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

	"github.com/livekit/gstplugins/base"
)

type Videobox struct {
	*base.GstBaseTransform
}

func NewVideobox() (*Videobox, error) {
	e, err := gst.NewElement("videobox")
	if err != nil {
		return nil, err
	}
	return &Videobox{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideoboxWithName(name string) (*Videobox, error) {
	e, err := gst.NewElementWithName("videobox", name)
	if err != nil {
		return nil, err
	}
	return &Videobox{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha (float64)
//
// Alpha value picture

func (e *Videobox) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videobox) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

// autocrop (bool)
//
// If set to TRUE videobox will automatically crop/pad the input
// video to be centered in the output.

func (e *Videobox) GetAutocrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("autocrop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videobox) SetAutocrop(value bool) error {
	return e.Element.SetProperty("autocrop", value)
}

// border-alpha (float64)
//
// Alpha value of the border

func (e *Videobox) GetBorderAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("border-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videobox) SetBorderAlpha(value float64) error {
	return e.Element.SetProperty("border-alpha", value)
}

// bottom (int)
//
// Pixels to box at bottom (<0 = add a border)

func (e *Videobox) GetBottom() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videobox) SetBottom(value int) error {
	return e.Element.SetProperty("bottom", value)
}

// fill (GstVideoBoxFill)
//
// How to fill the borders

func (e *Videobox) GetFill() (GstVideoBoxFill, error) {
	var value GstVideoBoxFill
	var ok bool
	v, err := e.Element.GetProperty("fill")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoBoxFill)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoBoxFill")
	}
	return value, nil
}

func (e *Videobox) SetFill(value GstVideoBoxFill) error {
	e.Element.SetArg("fill", string(value))
	return nil
}

// left (int)
//
// Pixels to box at left (<0  = add a border)

func (e *Videobox) GetLeft() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("left")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videobox) SetLeft(value int) error {
	return e.Element.SetProperty("left", value)
}

// right (int)
//
// Pixels to box at right (<0 = add a border)

func (e *Videobox) GetRight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("right")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videobox) SetRight(value int) error {
	return e.Element.SetProperty("right", value)
}

// top (int)
//
// Pixels to box at top (<0 = add a border)

func (e *Videobox) GetTop() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("top")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videobox) SetTop(value int) error {
	return e.Element.SetProperty("top", value)
}

// ----- Constants -----

type GstVideoBoxFill string

const (
	GstVideoBoxFillBlack GstVideoBoxFill = "black" // black (0) – Black
	GstVideoBoxFillGreen GstVideoBoxFill = "green" // green (1) – Green
	GstVideoBoxFillBlue GstVideoBoxFill = "blue" // blue (2) – Blue
	GstVideoBoxFillRed GstVideoBoxFill = "red" // red (3) – Red
	GstVideoBoxFillYellow GstVideoBoxFill = "yellow" // yellow (4) – Yellow
	GstVideoBoxFillWhite GstVideoBoxFill = "white" // white (5) – White
)

