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
)

type Disparity struct {
	Element *gst.Element
}

func NewDisparity() (*Disparity, error) {
	e, err := gst.NewElement("disparity")
	if err != nil {
		return nil, err
	}
	return &Disparity{Element: e}, nil
}

func NewDisparityWithName(name string) (*Disparity, error) {
	e, err := gst.NewElementWithName("disparity", name)
	if err != nil {
		return nil, err
	}
	return &Disparity{Element: e}, nil
}

// ----- Properties -----

// method (GstDisparityMethod)
//
// Stereo matching method to use

func (e *Disparity) GetMethod() (GstDisparityMethod, error) {
	var value GstDisparityMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDisparityMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDisparityMethod")
	}
	return value, nil
}

func (e *Disparity) SetMethod(value GstDisparityMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// ----- Constants -----

type GstDisparityMethod string

const (
	GstDisparityMethodSbm GstDisparityMethod = "sbm" // sbm (0) – Global block matching algorithm
	GstDisparityMethodSgbm GstDisparityMethod = "sgbm" // sgbm (1) – Semi-global block matching algorithm
)

