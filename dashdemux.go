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

type Dashdemux struct {
	Element *gst.Element
}

func NewDashdemux() (*Dashdemux, error) {
	e, err := gst.NewElement("dashdemux")
	if err != nil {
		return nil, err
	}
	return &Dashdemux{Element: e}, nil
}

func NewDashdemuxWithName(name string) (*Dashdemux, error) {
	e, err := gst.NewElementWithName("dashdemux", name)
	if err != nil {
		return nil, err
	}
	return &Dashdemux{Element: e}, nil
}

// ----- Properties -----

// bandwidth-usage (float32)
//
// Percentage of the available bandwidth to use when selecting representations (deprecated)

func (e *Dashdemux) GetBandwidthUsage() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("bandwidth-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Dashdemux) SetBandwidthUsage(value float32) error {
	return e.Element.SetProperty("bandwidth-usage", value)
}

// max-bitrate (uint)
//
// Max of bitrate supported by target video decoder (0 = no maximum)

func (e *Dashdemux) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dashdemux) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-buffering-time (uint)
//
// Maximum number of seconds of buffer accumulated during playback(deprecated)

func (e *Dashdemux) GetMaxBufferingTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-buffering-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dashdemux) SetMaxBufferingTime(value uint) error {
	return e.Element.SetProperty("max-buffering-time", value)
}

// max-video-framerate (GstFraction)
//
// Max video framerate to select (0/1 = no maximum)

func (e *Dashdemux) GetMaxVideoFramerate() (interface{}, error) {
	return e.Element.GetProperty("max-video-framerate")
}

func (e *Dashdemux) SetMaxVideoFramerate(value interface{}) error {
	return e.Element.SetProperty("max-video-framerate", value)
}

// max-video-height (uint)
//
// Max video height to select (0 = no maximum)

func (e *Dashdemux) GetMaxVideoHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-video-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dashdemux) SetMaxVideoHeight(value uint) error {
	return e.Element.SetProperty("max-video-height", value)
}

// max-video-width (uint)
//
// Max video width to select (0 = no maximum)

func (e *Dashdemux) GetMaxVideoWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-video-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dashdemux) SetMaxVideoWidth(value uint) error {
	return e.Element.SetProperty("max-video-width", value)
}

// presentation-delay (string)
//
// Default presentation delay (in seconds, milliseconds or fragments) (e.g. 12s, 2500ms, 3f)

func (e *Dashdemux) GetPresentationDelay() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("presentation-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dashdemux) SetPresentationDelay(value string) error {
	return e.Element.SetProperty("presentation-delay", value)
}

