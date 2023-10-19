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

type Frei0RFilterSigmoidaltransfer struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterSigmoidaltransfer() (*Frei0RFilterSigmoidaltransfer, error) {
	e, err := gst.NewElement("frei0r-filter-sigmoidaltransfer")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSigmoidaltransfer{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterSigmoidaltransferWithName(name string) (*Frei0RFilterSigmoidaltransfer, error) {
	e, err := gst.NewElementWithName("frei0r-filter-sigmoidaltransfer", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSigmoidaltransfer{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// brightness (float64)
//
// Brightnesss of image

func (e *Frei0RFilterSigmoidaltransfer) GetBrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSigmoidaltransfer) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

// sharpness (float64)
//
// Sharpness of transfer

func (e *Frei0RFilterSigmoidaltransfer) GetSharpness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sharpness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSigmoidaltransfer) SetSharpness(value float64) error {
	return e.Element.SetProperty("sharpness", value)
}

