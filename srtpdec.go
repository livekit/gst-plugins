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

type Srtpdec struct {
	Element *gst.Element
}

func NewSrtpdec() (*Srtpdec, error) {
	e, err := gst.NewElement("srtpdec")
	if err != nil {
		return nil, err
	}
	return &Srtpdec{Element: e}, nil
}

func NewSrtpdecWithName(name string) (*Srtpdec, error) {
	e, err := gst.NewElementWithName("srtpdec", name)
	if err != nil {
		return nil, err
	}
	return &Srtpdec{Element: e}, nil
}

// ----- Properties -----

// replay-window-size (uint)
//
// Size of the replay protection window

func (e *Srtpdec) GetReplayWindowSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("replay-window-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Srtpdec) SetReplayWindowSize(value uint) error {
	return e.Element.SetProperty("replay-window-size", value)
}

// stats (GstStructure)
//
// Various statistics

func (e *Srtpdec) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

