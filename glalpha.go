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

type Glalpha struct {
	*base.GstBaseTransform
}

func NewGlalpha() (*Glalpha, error) {
	e, err := gst.NewElement("glalpha")
	if err != nil {
		return nil, err
	}
	return &Glalpha{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGlalphaWithName(name string) (*Glalpha, error) {
	e, err := gst.NewElementWithName("glalpha", name)
	if err != nil {
		return nil, err
	}
	return &Glalpha{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha (float64)
//
// The value for the alpha channel

func (e *Glalpha) GetAlpha() (float64, error) {
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

func (e *Glalpha) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

// angle (float32)
//
// Size of the colorcube to change

func (e *Glalpha) GetAngle() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Glalpha) SetAngle(value float32) error {
	return e.Element.SetProperty("angle", value)
}

// black-sensitivity (uint)
//
// Sensitivity to dark colors

func (e *Glalpha) GetBlackSensitivity() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("black-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Glalpha) SetBlackSensitivity(value uint) error {
	return e.Element.SetProperty("black-sensitivity", value)
}

// method (GstGlalphaMethod)
//
// How the alpha channels should be created

func (e *Glalpha) GetMethod() (GstGlalphaMethod, error) {
	var value GstGlalphaMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGlalphaMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGlalphaMethod")
	}
	return value, nil
}

func (e *Glalpha) SetMethod(value GstGlalphaMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// noise-level (float32)
//
// Size of noise radius

func (e *Glalpha) GetNoiseLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("noise-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Glalpha) SetNoiseLevel(value float32) error {
	return e.Element.SetProperty("noise-level", value)
}

// target-b (uint)
//
// The blue color value for custom RGB chroma keying

func (e *Glalpha) GetTargetB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Glalpha) SetTargetB(value uint) error {
	return e.Element.SetProperty("target-b", value)
}

// target-g (uint)
//
// The green color value for custom RGB chroma keying

func (e *Glalpha) GetTargetG() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Glalpha) SetTargetG(value uint) error {
	return e.Element.SetProperty("target-g", value)
}

// target-r (uint)
//
// The red color value for custom RGB chroma keying

func (e *Glalpha) GetTargetR() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Glalpha) SetTargetR(value uint) error {
	return e.Element.SetProperty("target-r", value)
}

// white-sensitivity (uint)
//
// Sensitivity to bright colors

func (e *Glalpha) GetWhiteSensitivity() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("white-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Glalpha) SetWhiteSensitivity(value uint) error {
	return e.Element.SetProperty("white-sensitivity", value)
}

// ----- Constants -----

type GstGlalphaMethod string

const (
	GstGlalphaMethodSet GstGlalphaMethod = "set" // set (0) – Set/adjust alpha channel
	GstGlalphaMethodGreen GstGlalphaMethod = "green" // green (1) – Chroma Key on pure green
	GstGlalphaMethodBlue GstGlalphaMethod = "blue" // blue (2) – Chroma Key on pure blue
	GstGlalphaMethodCustom GstGlalphaMethod = "custom" // custom (3) – Chroma Key on custom RGB values
)

