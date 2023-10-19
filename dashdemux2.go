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

type Dashdemux2 struct {
	Element *gst.Element
}

func NewDashdemux2() (*Dashdemux2, error) {
	e, err := gst.NewElement("dashdemux2")
	if err != nil {
		return nil, err
	}
	return &Dashdemux2{Element: e}, nil
}

func NewDashdemux2WithName(name string) (*Dashdemux2, error) {
	e, err := gst.NewElementWithName("dashdemux2", name)
	if err != nil {
		return nil, err
	}
	return &Dashdemux2{Element: e}, nil
}

// ----- Properties -----

// max-bitrate (uint)
//
// Max of bitrate supported by target video decoder (0 = no maximum)

func (e *Dashdemux2) GetMaxBitrate() (uint, error) {
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

func (e *Dashdemux2) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-video-framerate (GstFraction)
//
// Max video framerate to select (0/1 = no maximum)

func (e *Dashdemux2) GetMaxVideoFramerate() (interface{}, error) {
	return e.Element.GetProperty("max-video-framerate")
}

func (e *Dashdemux2) SetMaxVideoFramerate(value interface{}) error {
	return e.Element.SetProperty("max-video-framerate", value)
}

// max-video-height (uint)
//
// Max video height to select (0 = no maximum)

func (e *Dashdemux2) GetMaxVideoHeight() (uint, error) {
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

func (e *Dashdemux2) SetMaxVideoHeight(value uint) error {
	return e.Element.SetProperty("max-video-height", value)
}

// max-video-width (uint)
//
// Max video width to select (0 = no maximum)

func (e *Dashdemux2) GetMaxVideoWidth() (uint, error) {
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

func (e *Dashdemux2) SetMaxVideoWidth(value uint) error {
	return e.Element.SetProperty("max-video-width", value)
}

// presentation-delay (string)
//
// Default presentation delay (in seconds, milliseconds or fragments) (e.g. 12s, 2500ms, 3f)

func (e *Dashdemux2) GetPresentationDelay() (string, error) {
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

func (e *Dashdemux2) SetPresentationDelay(value string) error {
	return e.Element.SetProperty("presentation-delay", value)
}

// start-bitrate (uint)

func (e *Dashdemux2) GetStartBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("start-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dashdemux2) SetStartBitrate(value uint) error {
	return e.Element.SetProperty("start-bitrate", value)
}

