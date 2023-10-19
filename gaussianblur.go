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

type Gaussianblur struct {
	*base.GstBaseTransform
}

func NewGaussianblur() (*Gaussianblur, error) {
	e, err := gst.NewElement("gaussianblur")
	if err != nil {
		return nil, err
	}
	return &Gaussianblur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGaussianblurWithName(name string) (*Gaussianblur, error) {
	e, err := gst.NewElementWithName("gaussianblur", name)
	if err != nil {
		return nil, err
	}
	return &Gaussianblur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// sigma (float64)
//
// Sigma value for gaussian blur (negative for sharpen)

func (e *Gaussianblur) GetSigma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sigma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Gaussianblur) SetSigma(value float64) error {
	return e.Element.SetProperty("sigma", value)
}

