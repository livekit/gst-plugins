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

type Fpsdisplaysink struct {
	Element *gst.Element
}

func NewFpsdisplaysink() (*Fpsdisplaysink, error) {
	e, err := gst.NewElement("fpsdisplaysink")
	if err != nil {
		return nil, err
	}
	return &Fpsdisplaysink{Element: e}, nil
}

func NewFpsdisplaysinkWithName(name string) (*Fpsdisplaysink, error) {
	e, err := gst.NewElementWithName("fpsdisplaysink", name)
	if err != nil {
		return nil, err
	}
	return &Fpsdisplaysink{Element: e}, nil
}

// ----- Properties -----

// fps-update-interval (int)
//
// Time between consecutive frames per second measures and update  (in ms). Should be set on NULL state

func (e *Fpsdisplaysink) GetFpsUpdateInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fps-update-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fpsdisplaysink) SetFpsUpdateInterval(value int) error {
	return e.Element.SetProperty("fps-update-interval", value)
}

// frames-dropped (uint)
//
// Number of frames dropped by the sink

func (e *Fpsdisplaysink) GetFramesDropped() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frames-dropped")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// frames-rendered (uint)
//
// Number of frames rendered

func (e *Fpsdisplaysink) GetFramesRendered() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frames-rendered")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// last-message (string)
//
// The message describing current status

func (e *Fpsdisplaysink) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// max-fps (float64)
//
// Maximum fps rate measured. Reset when going from NULL to READY.-1 means no measurement has yet been done

func (e *Fpsdisplaysink) GetMaxFps() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("max-fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// min-fps (float64)
//
// Minimum fps rate measured. Reset when going from NULL to READY.-1 means no measurement has yet been done

func (e *Fpsdisplaysink) GetMinFps() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("min-fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// signal-fps-measurements (bool)
//
// If the fps-measurements signal should be emitted.

func (e *Fpsdisplaysink) GetSignalFpsMeasurements() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal-fps-measurements")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fpsdisplaysink) SetSignalFpsMeasurements(value bool) error {
	return e.Element.SetProperty("signal-fps-measurements", value)
}

// silent (bool)
//
// Don't produce last_message events

func (e *Fpsdisplaysink) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fpsdisplaysink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// sync (bool)
//
// Sync on the clock (if the internally used sink doesn't have this property it will be ignored

func (e *Fpsdisplaysink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fpsdisplaysink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// text-overlay (bool)
//
// Whether to use text-overlay

func (e *Fpsdisplaysink) GetTextOverlay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("text-overlay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fpsdisplaysink) SetTextOverlay(value bool) error {
	return e.Element.SetProperty("text-overlay", value)
}

// video-sink (GstElement)
//
// Video sink to use (Must only be called on NULL state)

func (e *Fpsdisplaysink) GetVideoSink() (interface{}, error) {
	return e.Element.GetProperty("video-sink")
}

func (e *Fpsdisplaysink) SetVideoSink(value interface{}) error {
	return e.Element.SetProperty("video-sink", value)
}

