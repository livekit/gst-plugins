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

type Opencvtextoverlay struct {
	*base.GstBaseTransform
}

func NewOpencvtextoverlay() (*Opencvtextoverlay, error) {
	e, err := gst.NewElement("opencvtextoverlay")
	if err != nil {
		return nil, err
	}
	return &Opencvtextoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewOpencvtextoverlayWithName(name string) (*Opencvtextoverlay, error) {
	e, err := gst.NewElementWithName("opencvtextoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Opencvtextoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// colorB (int)
//
// Sets the color -B

func (e *Opencvtextoverlay) GetColorB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorB")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetColorB(value int) error {
	return e.Element.SetProperty("colorB", value)
}

// colorG (int)
//
// Sets the color -G

func (e *Opencvtextoverlay) GetColorG() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorG")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetColorG(value int) error {
	return e.Element.SetProperty("colorG", value)
}

// colorR (int)
//
// Sets the color -R

func (e *Opencvtextoverlay) GetColorR() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorR")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetColorR(value int) error {
	return e.Element.SetProperty("colorR", value)
}

// height (float64)
//
// Sets the height of fonts

func (e *Opencvtextoverlay) GetHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetHeight(value float64) error {
	return e.Element.SetProperty("height", value)
}

// text (string)
//
// Text to be display.

func (e *Opencvtextoverlay) GetText() (string, error) {
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

func (e *Opencvtextoverlay) SetText(value string) error {
	return e.Element.SetProperty("text", value)
}

// thickness (int)
//
// Sets the Thickness of Font

func (e *Opencvtextoverlay) GetThickness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("thickness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetThickness(value int) error {
	return e.Element.SetProperty("thickness", value)
}

// width (float64)
//
// Sets the width of fonts

func (e *Opencvtextoverlay) GetWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetWidth(value float64) error {
	return e.Element.SetProperty("width", value)
}

// xpos (int)
//
// Sets the Horizontal position

func (e *Opencvtextoverlay) GetXpos() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xpos")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetXpos(value int) error {
	return e.Element.SetProperty("xpos", value)
}

// ypos (int)
//
// Sets the Vertical position

func (e *Opencvtextoverlay) GetYpos() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ypos")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opencvtextoverlay) SetYpos(value int) error {
	return e.Element.SetProperty("ypos", value)
}

