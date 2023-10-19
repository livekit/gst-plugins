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

type Transcodebin struct {
	Element *gst.Element
}

func NewTranscodebin() (*Transcodebin, error) {
	e, err := gst.NewElement("transcodebin")
	if err != nil {
		return nil, err
	}
	return &Transcodebin{Element: e}, nil
}

func NewTranscodebinWithName(name string) (*Transcodebin, error) {
	e, err := gst.NewElementWithName("transcodebin", name)
	if err != nil {
		return nil, err
	}
	return &Transcodebin{Element: e}, nil
}

// ----- Properties -----

// audio-filter (GstElement)
//
// Set the audio filter element/bin to use.

func (e *Transcodebin) GetAudioFilter() (interface{}, error) {
	return e.Element.GetProperty("audio-filter")
}

func (e *Transcodebin) SetAudioFilter(value interface{}) error {
	return e.Element.SetProperty("audio-filter", value)
}

// avoid-reencoding (bool)
//
// See avoid-reencoding

func (e *Transcodebin) GetAvoidReencoding() (bool, error) {
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

func (e *Transcodebin) SetAvoidReencoding(value bool) error {
	return e.Element.SetProperty("avoid-reencoding", value)
}

// profile (GstEncodingProfile)
//
// The GstEncodingProfile to use. This property must be set before going
// to GST_STATE_PAUSED or higher.

func (e *Transcodebin) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}

func (e *Transcodebin) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

// video-filter (GstElement)
//
// Set the video filter element/bin to use.

func (e *Transcodebin) GetVideoFilter() (interface{}, error) {
	return e.Element.GetProperty("video-filter")
}

func (e *Transcodebin) SetVideoFilter(value interface{}) error {
	return e.Element.SetProperty("video-filter", value)
}

