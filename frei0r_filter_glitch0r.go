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

type Frei0RFilterGlitch0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterGlitch0R() (*Frei0RFilterGlitch0R, error) {
	e, err := gst.NewElement("frei0r-filter-glitch0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterGlitch0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterGlitch0RWithName(name string) (*Frei0RFilterGlitch0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-glitch0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterGlitch0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// block-height (float64)
//
// Height range of the block that will be shifted/glitched

func (e *Frei0RFilterGlitch0R) GetBlockHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("block-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterGlitch0R) SetBlockHeight(value float64) error {
	return e.Element.SetProperty("block-height", value)
}

// color-glitching-intensity (float64)
//
// How intensive should be color distortion

func (e *Frei0RFilterGlitch0R) GetColorGlitchingIntensity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color-glitching-intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterGlitch0R) SetColorGlitchingIntensity(value float64) error {
	return e.Element.SetProperty("color-glitching-intensity", value)
}

// glitch-frequency (float64)
//
// How frequently the glitch should be applied

func (e *Frei0RFilterGlitch0R) GetGlitchFrequency() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("glitch-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterGlitch0R) SetGlitchFrequency(value float64) error {
	return e.Element.SetProperty("glitch-frequency", value)
}

// shift-intensity (float64)
//
// How much we should move blocks when glitching

func (e *Frei0RFilterGlitch0R) GetShiftIntensity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shift-intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterGlitch0R) SetShiftIntensity(value float64) error {
	return e.Element.SetProperty("shift-intensity", value)
}

