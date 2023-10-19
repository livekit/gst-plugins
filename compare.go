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

type Compare struct {
	Element *gst.Element
}

func NewCompare() (*Compare, error) {
	e, err := gst.NewElement("compare")
	if err != nil {
		return nil, err
	}
	return &Compare{Element: e}, nil
}

func NewCompareWithName(name string) (*Compare, error) {
	e, err := gst.NewElementWithName("compare", name)
	if err != nil {
		return nil, err
	}
	return &Compare{Element: e}, nil
}

// ----- Properties -----

// meta (GstBufferCopyFlags)
//
// Indicates which metadata should be compared

func (e *Compare) GetMeta() (interface{}, error) {
	return e.Element.GetProperty("meta")
}

func (e *Compare) SetMeta(value interface{}) error {
	return e.Element.SetProperty("meta", value)
}

// method (GstCompareMethod)
//
// Method to compare buffer content

func (e *Compare) GetMethod() (GstCompareMethod, error) {
	var value GstCompareMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCompareMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCompareMethod")
	}
	return value, nil
}

func (e *Compare) SetMethod(value GstCompareMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// offset-ts (bool)
//
// Consider OFFSET and OFFSET_END part of timestamp metadata

func (e *Compare) GetOffsetTs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("offset-ts")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Compare) SetOffsetTs(value bool) error {
	return e.Element.SetProperty("offset-ts", value)
}

// threshold (float64)
//
// Threshold beyond which to consider content different as determined by content-method

func (e *Compare) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Compare) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

// upper (bool)
//
// Whether threshold value is upper bound or lower bound for difference measure

func (e *Compare) GetUpper() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("upper")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Compare) SetUpper(value bool) error {
	return e.Element.SetProperty("upper", value)
}

// ----- Constants -----

type GstCompareMethod string

const (
	GstCompareMethodMem GstCompareMethod = "mem" // mem (0) – Memory
	GstCompareMethodMax GstCompareMethod = "max" // max (1) – Maximum metric
	GstCompareMethodSsim GstCompareMethod = "ssim" // ssim (2) – SSIM (raw video)
)

