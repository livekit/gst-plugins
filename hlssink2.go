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

type Hlssink2 struct {
	Element *gst.Element
}

func NewHlssink2() (*Hlssink2, error) {
	e, err := gst.NewElement("hlssink2")
	if err != nil {
		return nil, err
	}
	return &Hlssink2{Element: e}, nil
}

func NewHlssink2WithName(name string) (*Hlssink2, error) {
	e, err := gst.NewElementWithName("hlssink2", name)
	if err != nil {
		return nil, err
	}
	return &Hlssink2{Element: e}, nil
}

// ----- Properties -----

// location (string)
//
// Location of the file to write

func (e *Hlssink2) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Hlssink2) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-files (uint)
//
// Maximum number of files to keep on disk. Once the maximum is reached,old files start to be deleted to make room for new ones.

func (e *Hlssink2) GetMaxFiles() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-files")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Hlssink2) SetMaxFiles(value uint) error {
	return e.Element.SetProperty("max-files", value)
}

// playlist-length (uint)
//
// Length of HLS playlist. To allow players to conform to section 6.3.3 of the HLS specification, this should be at least 3. If set to 0, the playlist will be infinite.

func (e *Hlssink2) GetPlaylistLength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("playlist-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Hlssink2) SetPlaylistLength(value uint) error {
	return e.Element.SetProperty("playlist-length", value)
}

// playlist-location (string)
//
// Location of the playlist to write

func (e *Hlssink2) GetPlaylistLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("playlist-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Hlssink2) SetPlaylistLocation(value string) error {
	return e.Element.SetProperty("playlist-location", value)
}

// playlist-root (string)
//
// Location of the playlist to write

func (e *Hlssink2) GetPlaylistRoot() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("playlist-root")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Hlssink2) SetPlaylistRoot(value string) error {
	return e.Element.SetProperty("playlist-root", value)
}

// send-keyframe-requests (bool)
//
// Send keyframe requests to ensure correct fragmentation. If this is disabled then the input must have keyframes in regular intervals

func (e *Hlssink2) GetSendKeyframeRequests() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-keyframe-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Hlssink2) SetSendKeyframeRequests(value bool) error {
	return e.Element.SetProperty("send-keyframe-requests", value)
}

// target-duration (uint)
//
// The target duration in seconds of a segment/file. (0 - disabled, useful for management of segment duration by the streaming server)

func (e *Hlssink2) GetTargetDuration() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Hlssink2) SetTargetDuration(value uint) error {
	return e.Element.SetProperty("target-duration", value)
}

