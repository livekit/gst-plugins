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

type Alpha struct {
	*base.GstBaseTransform
}

func NewAlpha() (*Alpha, error) {
	e, err := gst.NewElement("alpha")
	if err != nil {
		return nil, err
	}
	return &Alpha{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAlphaWithName(name string) (*Alpha, error) {
	e, err := gst.NewElementWithName("alpha", name)
	if err != nil {
		return nil, err
	}
	return &Alpha{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha (float64)
//
// The value for the alpha channel

func (e *Alpha) GetAlpha() (float64, error) {
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

func (e *Alpha) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

// angle (float32)
//
// Size of the colorcube to change

func (e *Alpha) GetAngle() (float32, error) {
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

func (e *Alpha) SetAngle(value float32) error {
	return e.Element.SetProperty("angle", value)
}

// black-sensitivity (uint)
//
// Sensitivity to dark colors

func (e *Alpha) GetBlackSensitivity() (uint, error) {
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

func (e *Alpha) SetBlackSensitivity(value uint) error {
	return e.Element.SetProperty("black-sensitivity", value)
}

// method (GstAlphaMethod)
//
// How the alpha channels should be created

func (e *Alpha) GetMethod() (GstAlphaMethod, error) {
	var value GstAlphaMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAlphaMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAlphaMethod")
	}
	return value, nil
}

func (e *Alpha) SetMethod(value GstAlphaMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// noise-level (float32)
//
// Size of noise radius

func (e *Alpha) GetNoiseLevel() (float32, error) {
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

func (e *Alpha) SetNoiseLevel(value float32) error {
	return e.Element.SetProperty("noise-level", value)
}

// prefer-passthrough (bool)
//
// Don't do any processing for alpha=1.0 if possible

func (e *Alpha) GetPreferPassthrough() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("prefer-passthrough")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Alpha) SetPreferPassthrough(value bool) error {
	return e.Element.SetProperty("prefer-passthrough", value)
}

// target-b (uint)
//
// The blue color value for custom RGB chroma keying

func (e *Alpha) GetTargetB() (uint, error) {
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

func (e *Alpha) SetTargetB(value uint) error {
	return e.Element.SetProperty("target-b", value)
}

// target-g (uint)
//
// The green color value for custom RGB chroma keying

func (e *Alpha) GetTargetG() (uint, error) {
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

func (e *Alpha) SetTargetG(value uint) error {
	return e.Element.SetProperty("target-g", value)
}

// target-r (uint)
//
// The red color value for custom RGB chroma keying

func (e *Alpha) GetTargetR() (uint, error) {
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

func (e *Alpha) SetTargetR(value uint) error {
	return e.Element.SetProperty("target-r", value)
}

// white-sensitivity (uint)
//
// Sensitivity to bright colors

func (e *Alpha) GetWhiteSensitivity() (uint, error) {
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

func (e *Alpha) SetWhiteSensitivity(value uint) error {
	return e.Element.SetProperty("white-sensitivity", value)
}

// ----- Constants -----

type GstAlphaMethod string

const (
	GstAlphaMethodSet GstAlphaMethod = "set" // set (0) – Set/adjust alpha channel
	GstAlphaMethodGreen GstAlphaMethod = "green" // green (1) – Chroma Key on pure green
	GstAlphaMethodBlue GstAlphaMethod = "blue" // blue (2) – Chroma Key on pure blue
	GstAlphaMethodCustom GstAlphaMethod = "custom" // custom (3) – Chroma Key on custom RGB values
)

