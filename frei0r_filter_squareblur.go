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

type Frei0RFilterSquareblur struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterSquareblur() (*Frei0RFilterSquareblur, error) {
	e, err := gst.NewElement("frei0r-filter-squareblur")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSquareblur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterSquareblurWithName(name string) (*Frei0RFilterSquareblur, error) {
	e, err := gst.NewElementWithName("frei0r-filter-squareblur", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSquareblur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// kernel-size (float64)
//
// The size of the kernel, as a proportion to its coverage of the image

func (e *Frei0RFilterSquareblur) GetKernelSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("kernel-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSquareblur) SetKernelSize(value float64) error {
	return e.Element.SetProperty("kernel-size", value)
}

