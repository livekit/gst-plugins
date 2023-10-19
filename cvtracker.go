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

type Cvtracker struct {
	*base.GstBaseTransform
}

func NewCvtracker() (*Cvtracker, error) {
	e, err := gst.NewElement("cvtracker")
	if err != nil {
		return nil, err
	}
	return &Cvtracker{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCvtrackerWithName(name string) (*Cvtracker, error) {
	e, err := gst.NewElementWithName("cvtracker", name)
	if err != nil {
		return nil, err
	}
	return &Cvtracker{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// algorithm (GstOpenCvtrackerAlgorithm)
//
// Algorithm for tracking objects

func (e *Cvtracker) GetAlgorithm() (GstOpenCvtrackerAlgorithm, error) {
	var value GstOpenCvtrackerAlgorithm
	var ok bool
	v, err := e.Element.GetProperty("algorithm")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenCvtrackerAlgorithm)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenCvtrackerAlgorithm")
	}
	return value, nil
}

func (e *Cvtracker) SetAlgorithm(value GstOpenCvtrackerAlgorithm) error {
	e.Element.SetArg("algorithm", string(value))
	return nil
}

// draw-rect (bool)
//
// Draw rectangle around tracked object

func (e *Cvtracker) GetDrawRect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-rect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Cvtracker) SetDrawRect(value bool) error {
	return e.Element.SetProperty("draw-rect", value)
}

// object-initial-height (uint)
//
// Track object box's initial height

func (e *Cvtracker) GetObjectInitialHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Cvtracker) SetObjectInitialHeight(value uint) error {
	return e.Element.SetProperty("object-initial-height", value)
}

// object-initial-width (uint)
//
// Track object box's initial width

func (e *Cvtracker) GetObjectInitialWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Cvtracker) SetObjectInitialWidth(value uint) error {
	return e.Element.SetProperty("object-initial-width", value)
}

// object-initial-x (uint)
//
// Track object box's initial X coordinate

func (e *Cvtracker) GetObjectInitialX() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Cvtracker) SetObjectInitialX(value uint) error {
	return e.Element.SetProperty("object-initial-x", value)
}

// object-initial-y (uint)
//
// Track object box's initial Y coordinate

func (e *Cvtracker) GetObjectInitialY() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Cvtracker) SetObjectInitialY(value uint) error {
	return e.Element.SetProperty("object-initial-y", value)
}

// ----- Constants -----

type GstOpenCvtrackerAlgorithm string

const (
	GstOpenCvtrackerAlgorithmBoosting GstOpenCvtrackerAlgorithm = "Boosting" // Boosting (0) – the Boosting tracker
	GstOpenCvtrackerAlgorithmCsrt GstOpenCvtrackerAlgorithm = "CSRT" // CSRT (1) – the CSRT tracker
	GstOpenCvtrackerAlgorithmKcf GstOpenCvtrackerAlgorithm = "KCF" // KCF (2) – the KCF (Kernelized Correlation Filter) tracker
	GstOpenCvtrackerAlgorithmMedianFlow GstOpenCvtrackerAlgorithm = "MedianFlow" // MedianFlow (3) – the Median Flow tracker
	GstOpenCvtrackerAlgorithmMil GstOpenCvtrackerAlgorithm = "MIL" // MIL (4) – the MIL tracker
	GstOpenCvtrackerAlgorithmMosse GstOpenCvtrackerAlgorithm = "MOSSE" // MOSSE (5) – the MOSSE (Minimum Output Sum of Squared Error) tracker
	GstOpenCvtrackerAlgorithmTld GstOpenCvtrackerAlgorithm = "TLD" // TLD (6) – the TLD (Tracking, learning and detection) tracker
)

