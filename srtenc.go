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

type Srtenc struct {
	Element *gst.Element
}

func NewSrtenc() (*Srtenc, error) {
	e, err := gst.NewElement("srtenc")
	if err != nil {
		return nil, err
	}
	return &Srtenc{Element: e}, nil
}

func NewSrtencWithName(name string) (*Srtenc, error) {
	e, err := gst.NewElementWithName("srtenc", name)
	if err != nil {
		return nil, err
	}
	return &Srtenc{Element: e}, nil
}

// ----- Properties -----

// duration (int64)
//
// Offset for the duration of the subtitles

func (e *Srtenc) GetDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Srtenc) SetDuration(value int64) error {
	return e.Element.SetProperty("duration", value)
}

// timestamp (int64)
//
// Offset for the starttime for the subtitles

func (e *Srtenc) GetTimestamp() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Srtenc) SetTimestamp(value int64) error {
	return e.Element.SetProperty("timestamp", value)
}

