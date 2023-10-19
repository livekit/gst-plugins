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

type Audiokaraoke struct {
	*base.GstBaseTransform
}

func NewAudiokaraoke() (*Audiokaraoke, error) {
	e, err := gst.NewElement("audiokaraoke")
	if err != nil {
		return nil, err
	}
	return &Audiokaraoke{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiokaraokeWithName(name string) (*Audiokaraoke, error) {
	e, err := gst.NewElementWithName("audiokaraoke", name)
	if err != nil {
		return nil, err
	}
	return &Audiokaraoke{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// filter-band (float32)
//
// The Frequency band of the filter

func (e *Audiokaraoke) GetFilterBand() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("filter-band")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiokaraoke) SetFilterBand(value float32) error {
	return e.Element.SetProperty("filter-band", value)
}

// filter-width (float32)
//
// The Frequency width of the filter

func (e *Audiokaraoke) GetFilterWidth() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("filter-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiokaraoke) SetFilterWidth(value float32) error {
	return e.Element.SetProperty("filter-width", value)
}

// level (float32)
//
// Level of the effect (1.0 = full)

func (e *Audiokaraoke) GetLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiokaraoke) SetLevel(value float32) error {
	return e.Element.SetProperty("level", value)
}

// mono-level (float32)
//
// Level of the mono channel (1.0 = full)

func (e *Audiokaraoke) GetMonoLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mono-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiokaraoke) SetMonoLevel(value float32) error {
	return e.Element.SetProperty("mono-level", value)
}

