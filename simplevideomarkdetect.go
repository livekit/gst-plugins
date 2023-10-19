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

type Simplevideomarkdetect struct {
	*base.GstBaseTransform
}

func NewSimplevideomarkdetect() (*Simplevideomarkdetect, error) {
	e, err := gst.NewElement("simplevideomarkdetect")
	if err != nil {
		return nil, err
	}
	return &Simplevideomarkdetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSimplevideomarkdetectWithName(name string) (*Simplevideomarkdetect, error) {
	e, err := gst.NewElementWithName("simplevideomarkdetect", name)
	if err != nil {
		return nil, err
	}
	return &Simplevideomarkdetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bottom-offset (int)
//
// The offset from the bottom border where the pattern starts

func (e *Simplevideomarkdetect) GetBottomOffset() (int, error) {
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

func (e *Simplevideomarkdetect) SetBottomOffset(value int) error {
	return e.Element.SetProperty("bottom-offset", value)
}

// left-offset (int)
//
// The offset from the left border where the pattern starts

func (e *Simplevideomarkdetect) GetLeftOffset() (int, error) {
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

func (e *Simplevideomarkdetect) SetLeftOffset(value int) error {
	return e.Element.SetProperty("left-offset", value)
}

// message (bool)
//
// Post detected data as bus messages

func (e *Simplevideomarkdetect) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Simplevideomarkdetect) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

// pattern-center (float64)
//
// The center of the black/white separation (0.0 = lowest, 1.0 highest)

func (e *Simplevideomarkdetect) GetPatternCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("pattern-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Simplevideomarkdetect) SetPatternCenter(value float64) error {
	return e.Element.SetProperty("pattern-center", value)
}

// pattern-count (int)
//
// The number of pattern markers

func (e *Simplevideomarkdetect) GetPatternCount() (int, error) {
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

func (e *Simplevideomarkdetect) SetPatternCount(value int) error {
	return e.Element.SetProperty("pattern-count", value)
}

// pattern-data-count (int)
//
// The number of extra data pattern markers

func (e *Simplevideomarkdetect) GetPatternDataCount() (int, error) {
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

func (e *Simplevideomarkdetect) SetPatternDataCount(value int) error {
	return e.Element.SetProperty("pattern-data-count", value)
}

// pattern-height (int)
//
// The height of the pattern markers

func (e *Simplevideomarkdetect) GetPatternHeight() (int, error) {
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

func (e *Simplevideomarkdetect) SetPatternHeight(value int) error {
	return e.Element.SetProperty("pattern-height", value)
}

// pattern-sensitivity (float64)
//
// The sensitivity around the center for detecting the markers (0.0 = lowest, 1.0 highest)

func (e *Simplevideomarkdetect) GetPatternSensitivity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("pattern-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Simplevideomarkdetect) SetPatternSensitivity(value float64) error {
	return e.Element.SetProperty("pattern-sensitivity", value)
}

// pattern-width (int)
//
// The width of the pattern markers

func (e *Simplevideomarkdetect) GetPatternWidth() (int, error) {
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

func (e *Simplevideomarkdetect) SetPatternWidth(value int) error {
	return e.Element.SetProperty("pattern-width", value)
}

