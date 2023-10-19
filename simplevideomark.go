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

type Simplevideomark struct {
	*base.GstBaseTransform
}

func NewSimplevideomark() (*Simplevideomark, error) {
	e, err := gst.NewElement("simplevideomark")
	if err != nil {
		return nil, err
	}
	return &Simplevideomark{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSimplevideomarkWithName(name string) (*Simplevideomark, error) {
	e, err := gst.NewElementWithName("simplevideomark", name)
	if err != nil {
		return nil, err
	}
	return &Simplevideomark{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bottom-offset (int)
//
// The offset from the bottom border where the pattern starts

func (e *Simplevideomark) GetBottomOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bottom-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Simplevideomark) SetBottomOffset(value int) error {
	return e.Element.SetProperty("bottom-offset", value)
}

// enabled (bool)
//
// Enable or disable the filter

func (e *Simplevideomark) GetEnabled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enabled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Simplevideomark) SetEnabled(value bool) error {
	return e.Element.SetProperty("enabled", value)
}

// left-offset (int)
//
// The offset from the left border where the pattern starts

func (e *Simplevideomark) GetLeftOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("left-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Simplevideomark) SetLeftOffset(value int) error {
	return e.Element.SetProperty("left-offset", value)
}

// pattern-count (int)
//
// The number of pattern markers

func (e *Simplevideomark) GetPatternCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Simplevideomark) SetPatternCount(value int) error {
	return e.Element.SetProperty("pattern-count", value)
}

// pattern-data (uint64)
//
// The extra data pattern markers

func (e *Simplevideomark) GetPatternData() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("pattern-data")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Simplevideomark) SetPatternData(value uint64) error {
	return e.Element.SetProperty("pattern-data", value)
}

// pattern-data-count (int)
//
// The number of extra data pattern markers

func (e *Simplevideomark) GetPatternDataCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-data-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Simplevideomark) SetPatternDataCount(value int) error {
	return e.Element.SetProperty("pattern-data-count", value)
}

// pattern-height (int)
//
// The height of the pattern markers

func (e *Simplevideomark) GetPatternHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Simplevideomark) SetPatternHeight(value int) error {
	return e.Element.SetProperty("pattern-height", value)
}

// pattern-width (int)
//
// The width of the pattern markers

func (e *Simplevideomark) GetPatternWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Simplevideomark) SetPatternWidth(value int) error {
	return e.Element.SetProperty("pattern-width", value)
}

