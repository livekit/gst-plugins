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

type Dshowvideosink struct {
	*base.GstBaseSink
}

func NewDshowvideosink() (*Dshowvideosink, error) {
	e, err := gst.NewElement("dshowvideosink")
	if err != nil {
		return nil, err
	}
	return &Dshowvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewDshowvideosinkWithName(name string) (*Dshowvideosink, error) {
	e, err := gst.NewElementWithName("dshowvideosink", name)
	if err != nil {
		return nil, err
	}
	return &Dshowvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Dshowvideosink) GetForceAspectRatio() (bool, error) {
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

func (e *Dshowvideosink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// fullscreen (bool)
//
// Use full-screen mode (not available when using XOverlay)

func (e *Dshowvideosink) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dshowvideosink) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

// renderer (string)
//
// Force usage of specific DirectShow renderer (EVR, VMR9 or VMR7)

func (e *Dshowvideosink) GetRenderer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("renderer")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dshowvideosink) SetRenderer(value string) error {
	return e.Element.SetProperty("renderer", value)
}

