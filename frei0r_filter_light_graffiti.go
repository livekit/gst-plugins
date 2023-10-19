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

type Frei0RFilterLightGraffiti struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterLightGraffiti() (*Frei0RFilterLightGraffiti, error) {
	e, err := gst.NewElement("frei0r-filter-light-graffiti")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLightGraffiti{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterLightGraffitiWithName(name string) (*Frei0RFilterLightGraffiti, error) {
	e, err := gst.NewElementWithName("frei0r-filter-light-graffiti", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLightGraffiti{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// backgroundweight (float64)
//
// Describes how strong the (accumulated) background should shine through

func (e *Frei0RFilterLightGraffiti) GetBackgroundweight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("backgroundweight")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetBackgroundweight(value float64) error {
	return e.Element.SetProperty("backgroundweight", value)
}

// blackreference (bool)
//
// Uses black as background image instead of the first frame.

func (e *Frei0RFilterLightGraffiti) GetBlackreference() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("blackreference")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetBlackreference(value bool) error {
	return e.Element.SetProperty("blackreference", value)
}

// dim (float64)
//
// Dimming of the light mask

func (e *Frei0RFilterLightGraffiti) GetDim() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("dim")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetDim(value float64) error {
	return e.Element.SetProperty("dim", value)
}

// longalpha (float64)
//
// Alpha value for moving average

func (e *Frei0RFilterLightGraffiti) GetLongalpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("longalpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetLongalpha(value float64) error {
	return e.Element.SetProperty("longalpha", value)
}

// loweroverexposure (float64)
//
// Prevents some overexposure if the light source stays steady too long (varying speed)

func (e *Frei0RFilterLightGraffiti) GetLoweroverexposure() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("loweroverexposure")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetLoweroverexposure(value float64) error {
	return e.Element.SetProperty("loweroverexposure", value)
}

// nonlineardim (bool)
//
// Nonlinear dimming (may look more natural)

func (e *Frei0RFilterLightGraffiti) GetNonlineardim() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonlineardim")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetNonlineardim(value bool) error {
	return e.Element.SetProperty("nonlineardim", value)
}

// reset (bool)
//
// Reset filter masks

func (e *Frei0RFilterLightGraffiti) GetReset() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reset")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetReset(value bool) error {
	return e.Element.SetProperty("reset", value)
}

// saturation (float64)
//
// Saturation of lights

func (e *Frei0RFilterLightGraffiti) GetSaturation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

// sensitivity (float64)
//
// Sensitivity of the effect for light (higher sensitivity will lead to brighter lights)

func (e *Frei0RFilterLightGraffiti) GetSensitivity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetSensitivity(value float64) error {
	return e.Element.SetProperty("sensitivity", value)
}

// statsbrightness (bool)
//
// Display the brightness and threshold, for adjusting the brightness threshold parameter

func (e *Frei0RFilterLightGraffiti) GetStatsbrightness() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("statsbrightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetStatsbrightness(value bool) error {
	return e.Element.SetProperty("statsbrightness", value)
}

// statsdifference (bool)
//
// Display the background difference and threshold

func (e *Frei0RFilterLightGraffiti) GetStatsdifference() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("statsdifference")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetStatsdifference(value bool) error {
	return e.Element.SetProperty("statsdifference", value)
}

// statsdiffsum (bool)
//
// Display the sum of the background difference and the threshold

func (e *Frei0RFilterLightGraffiti) GetStatsdiffsum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("statsdiffsum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetStatsdiffsum(value bool) error {
	return e.Element.SetProperty("statsdiffsum", value)
}

// thresholdbrightness (float64)
//
// Brightness threshold to distinguish between foreground and background

func (e *Frei0RFilterLightGraffiti) GetThresholdbrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("thresholdbrightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetThresholdbrightness(value float64) error {
	return e.Element.SetProperty("thresholdbrightness", value)
}

// thresholddifference (float64)
//
// Threshold: Difference to background to distinguish between fore- and background

func (e *Frei0RFilterLightGraffiti) GetThresholddifference() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("thresholddifference")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetThresholddifference(value float64) error {
	return e.Element.SetProperty("thresholddifference", value)
}

// thresholddiffsum (float64)
//
// Threshold for sum of differences. Can in most cases be ignored (set to 0).

func (e *Frei0RFilterLightGraffiti) GetThresholddiffsum() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("thresholddiffsum")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetThresholddiffsum(value float64) error {
	return e.Element.SetProperty("thresholddiffsum", value)
}

// transparentbackground (bool)
//
// Make the background transparent

func (e *Frei0RFilterLightGraffiti) GetTransparentbackground() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("transparentbackground")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLightGraffiti) SetTransparentbackground(value bool) error {
	return e.Element.SetProperty("transparentbackground", value)
}

