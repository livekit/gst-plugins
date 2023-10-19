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

type Uritranscodebin struct {
	Element *gst.Element
}

func NewUritranscodebin() (*Uritranscodebin, error) {
	e, err := gst.NewElement("uritranscodebin")
	if err != nil {
		return nil, err
	}
	return &Uritranscodebin{Element: e}, nil
}

func NewUritranscodebinWithName(name string) (*Uritranscodebin, error) {
	e, err := gst.NewElementWithName("uritranscodebin", name)
	if err != nil {
		return nil, err
	}
	return &Uritranscodebin{Element: e}, nil
}

// ----- Properties -----

// audio-filter (GstElement)
//
// Set the audio filter element/bin to use.

func (e *Uritranscodebin) GetAudioFilter() (interface{}, error) {
	return e.Element.GetProperty("audio-filter")
}

func (e *Uritranscodebin) SetAudioFilter(value interface{}) error {
	return e.Element.SetProperty("audio-filter", value)
}

// avoid-reencoding (bool)
//
// See avoid-reencoding

func (e *Uritranscodebin) GetAvoidReencoding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("avoid-reencoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Uritranscodebin) SetAvoidReencoding(value bool) error {
	return e.Element.SetProperty("avoid-reencoding", value)
}

// cpu-usage (uint)
//
// The percentage of CPU to try to use with the processus running the pipeline driven by the clock

func (e *Uritranscodebin) GetCpuUsage() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cpu-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uritranscodebin) SetCpuUsage(value uint) error {
	return e.Element.SetProperty("cpu-usage", value)
}

// dest-uri (string)
//
// The destination URI to which the stream should be encoded.

func (e *Uritranscodebin) GetDestUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dest-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uritranscodebin) SetDestUri(value string) error {
	return e.Element.SetProperty("dest-uri", value)
}

// profile (GstEncodingProfile)
//
// The GstEncodingProfile to use. This property must be set before going
// to GST_STATE_PAUSED or higher.

func (e *Uritranscodebin) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}

func (e *Uritranscodebin) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

// source-uri (string)
//
// The URI of the stream to encode

func (e *Uritranscodebin) GetSourceUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("source-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uritranscodebin) SetSourceUri(value string) error {
	return e.Element.SetProperty("source-uri", value)
}

// video-filter (GstElement)
//
// Set the video filter element/bin to use.

func (e *Uritranscodebin) GetVideoFilter() (interface{}, error) {
	return e.Element.GetProperty("video-filter")
}

func (e *Uritranscodebin) SetVideoFilter(value interface{}) error {
	return e.Element.SetProperty("video-filter", value)
}

