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

type Subparse struct {
	Element *gst.Element
}

func NewSubparse() (*Subparse, error) {
	e, err := gst.NewElement("subparse")
	if err != nil {
		return nil, err
	}
	return &Subparse{Element: e}, nil
}

func NewSubparseWithName(name string) (*Subparse, error) {
	e, err := gst.NewElementWithName("subparse", name)
	if err != nil {
		return nil, err
	}
	return &Subparse{Element: e}, nil
}

// ----- Properties -----

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 or any other Unicode encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Subparse) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Subparse) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

// video-fps (GstFraction)
//
// Framerate of the video stream. This is needed by some subtitle formats to synchronize subtitles and video properly. If not set and the subtitle format requires it subtitles may be out of sync.

func (e *Subparse) GetVideoFps() (interface{}, error) {
	return e.Element.GetProperty("video-fps")
}

func (e *Subparse) SetVideoFps(value interface{}) error {
	return e.Element.SetProperty("video-fps", value)
}

