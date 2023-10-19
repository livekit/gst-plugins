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

type Ximagesink struct {
	*base.GstBaseSink
}

func NewXimagesink() (*Ximagesink, error) {
	e, err := gst.NewElement("ximagesink")
	if err != nil {
		return nil, err
	}
	return &Ximagesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewXimagesinkWithName(name string) (*Ximagesink, error) {
	e, err := gst.NewElementWithName("ximagesink", name)
	if err != nil {
		return nil, err
	}
	return &Ximagesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// display (string)
//
// X Display name

func (e *Ximagesink) GetDisplay() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Ximagesink) SetDisplay(value string) error {
	return e.Element.SetProperty("display", value)
}

// force-aspect-ratio (bool)
//
// When enabled, reverse caps negotiation (scaling) will respect original aspect ratio

func (e *Ximagesink) GetForceAspectRatio() (bool, error) {
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

func (e *Ximagesink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// handle-events (bool)
//
// When enabled, XEvents will be selected and handled

func (e *Ximagesink) GetHandleEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ximagesink) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

// handle-expose (bool)
//
// When enabled, the current frame will always be drawn in response to X Expose events

func (e *Ximagesink) GetHandleExpose() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-expose")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ximagesink) SetHandleExpose(value bool) error {
	return e.Element.SetProperty("handle-expose", value)
}

// pixel-aspect-ratio (string)
//
// The pixel aspect ratio of the device

func (e *Ximagesink) GetPixelAspectRatio() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pixel-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Ximagesink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// synchronous (bool)
//
// When enabled, runs the X display in synchronous mode. (unrelated to A/V sync, used only for debugging)

func (e *Ximagesink) GetSynchronous() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("synchronous")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ximagesink) SetSynchronous(value bool) error {
	return e.Element.SetProperty("synchronous", value)
}

// window-height (uint64)
//
// Actual height of the video window.

func (e *Ximagesink) GetWindowHeight() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// window-width (uint64)
//
// Actual width of the video window.

func (e *Ximagesink) GetWindowWidth() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

