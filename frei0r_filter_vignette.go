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

type Frei0RFilterVignette struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterVignette() (*Frei0RFilterVignette, error) {
	e, err := gst.NewElement("frei0r-filter-vignette")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterVignette{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterVignetteWithName(name string) (*Frei0RFilterVignette, error) {
	e, err := gst.NewElementWithName("frei0r-filter-vignette", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterVignette{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// aspect (float64)
//
// Aspect ratio

func (e *Frei0RFilterVignette) GetAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterVignette) SetAspect(value float64) error {
	return e.Element.SetProperty("aspect", value)
}

// clearcenter (float64)
//
// Size of the unaffected center

func (e *Frei0RFilterVignette) GetClearcenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clearcenter")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterVignette) SetClearcenter(value float64) error {
	return e.Element.SetProperty("clearcenter", value)
}

// soft (float64)
//
// Softness

func (e *Frei0RFilterVignette) GetSoft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("soft")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterVignette) SetSoft(value float64) error {
	return e.Element.SetProperty("soft", value)
}

