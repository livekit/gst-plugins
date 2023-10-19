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

type Osxvideosink struct {
	*base.GstBaseSink
}

func NewOsxvideosink() (*Osxvideosink, error) {
	e, err := gst.NewElement("osxvideosink")
	if err != nil {
		return nil, err
	}
	return &Osxvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewOsxvideosinkWithName(name string) (*Osxvideosink, error) {
	e, err := gst.NewElementWithName("osxvideosink", name)
	if err != nil {
		return nil, err
	}
	return &Osxvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// embed (bool)
//
// For ABI compatibility only, do not use

func (e *Osxvideosink) GetEmbed() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("embed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Osxvideosink) SetEmbed(value bool) error {
	return e.Element.SetProperty("embed", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ration

func (e *Osxvideosink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Osxvideosink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

