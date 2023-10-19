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

type Id3Demux struct {
	Element *gst.Element
}

func NewId3Demux() (*Id3Demux, error) {
	e, err := gst.NewElement("id3demux")
	if err != nil {
		return nil, err
	}
	return &Id3Demux{Element: e}, nil
}

func NewId3DemuxWithName(name string) (*Id3Demux, error) {
	e, err := gst.NewElementWithName("id3demux", name)
	if err != nil {
		return nil, err
	}
	return &Id3Demux{Element: e}, nil
}

// ----- Properties -----

// prefer-v1 (bool)
//
// Prefer tags from ID3v1 tag at end of file when both ID3v1 and ID3v2 tags are present

func (e *Id3Demux) GetPreferV1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("prefer-v1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Id3Demux) SetPreferV1(value bool) error {
	return e.Element.SetProperty("prefer-v1", value)
}

