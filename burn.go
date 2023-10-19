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

type Burn struct {
	*base.GstBaseTransform
}

func NewBurn() (*Burn, error) {
	e, err := gst.NewElement("burn")
	if err != nil {
		return nil, err
	}
	return &Burn{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewBurnWithName(name string) (*Burn, error) {
	e, err := gst.NewElementWithName("burn", name)
	if err != nil {
		return nil, err
	}
	return &Burn{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// adjustment (uint)
//
// Adjustment parameter

func (e *Burn) GetAdjustment() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("adjustment")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Burn) SetAdjustment(value uint) error {
	return e.Element.SetProperty("adjustment", value)
}

