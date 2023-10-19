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

type Exclusion struct {
	*base.GstBaseTransform
}

func NewExclusion() (*Exclusion, error) {
	e, err := gst.NewElement("exclusion")
	if err != nil {
		return nil, err
	}
	return &Exclusion{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewExclusionWithName(name string) (*Exclusion, error) {
	e, err := gst.NewElementWithName("exclusion", name)
	if err != nil {
		return nil, err
	}
	return &Exclusion{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// factor (uint)
//
// Exclusion factor parameter

func (e *Exclusion) GetFactor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Exclusion) SetFactor(value uint) error {
	return e.Element.SetProperty("factor", value)
}

