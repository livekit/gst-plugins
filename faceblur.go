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

type Faceblur struct {
	*base.GstBaseTransform
}

func NewFaceblur() (*Faceblur, error) {
	e, err := gst.NewElement("faceblur")
	if err != nil {
		return nil, err
	}
	return &Faceblur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFaceblurWithName(name string) (*Faceblur, error) {
	e, err := gst.NewElementWithName("faceblur", name)
	if err != nil {
		return nil, err
	}
	return &Faceblur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// flags (GstOpencvFaceBlurFlags)
//
// Flags to cvHaarDetectObjects

func (e *Faceblur) GetFlags() (GstOpencvFaceBlurFlags, error) {
	var value GstOpencvFaceBlurFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpencvFaceBlurFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpencvFaceBlurFlags")
	}
	return value, nil
}

func (e *Faceblur) SetFlags(value GstOpencvFaceBlurFlags) error {
	e.Element.SetArg("flags", string(value))
	return nil
}

// min-neighbors (int)
//
// Minimum number (minus 1) of neighbor rectangles that makes up an object

func (e *Faceblur) GetMinNeighbors() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-neighbors")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Faceblur) SetMinNeighbors(value int) error {
	return e.Element.SetProperty("min-neighbors", value)
}

// min-size-height (int)
//
// Minimum window height size

func (e *Faceblur) GetMinSizeHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Faceblur) SetMinSizeHeight(value int) error {
	return e.Element.SetProperty("min-size-height", value)
}

// min-size-width (int)
//
// Minimum window width size

func (e *Faceblur) GetMinSizeWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Faceblur) SetMinSizeWidth(value int) error {
	return e.Element.SetProperty("min-size-width", value)
}

// profile (string)
//
// Location of Haar cascade file to use for face blurion

func (e *Faceblur) GetProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Faceblur) SetProfile(value string) error {
	return e.Element.SetProperty("profile", value)
}

// scale-factor (float64)
//
// Factor by which the windows is scaled after each scan

func (e *Faceblur) GetScaleFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Faceblur) SetScaleFactor(value float64) error {
	return e.Element.SetProperty("scale-factor", value)
}

// ----- Constants -----

type GstOpencvFaceBlurFlags string

const (
	GstOpencvFaceBlurFlagsDoCannyPruning GstOpencvFaceBlurFlags = "do-canny-pruning" // do-canny-pruning (0x00000001) – Do Canny edge detection to discard some regions
)

