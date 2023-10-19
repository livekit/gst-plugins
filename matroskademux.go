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

type Matroskademux struct {
	Element *gst.Element
}

func NewMatroskademux() (*Matroskademux, error) {
	e, err := gst.NewElement("matroskademux")
	if err != nil {
		return nil, err
	}
	return &Matroskademux{Element: e}, nil
}

func NewMatroskademuxWithName(name string) (*Matroskademux, error) {
	e, err := gst.NewElementWithName("matroskademux", name)
	if err != nil {
		return nil, err
	}
	return &Matroskademux{Element: e}, nil
}

// ----- Properties -----

// max-backtrack-distance (uint)
//
// Maximum backtrack distance in seconds when seeking without and index in pull mode and search for a keyframe (0 = disable backtracking).

func (e *Matroskademux) GetMaxBacktrackDistance() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-backtrack-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Matroskademux) SetMaxBacktrackDistance(value uint) error {
	return e.Element.SetProperty("max-backtrack-distance", value)
}

// max-gap-time (uint64)
//
// The demuxer sends out segment events for skipping gaps longer than this (0 = disabled).

func (e *Matroskademux) GetMaxGapTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-gap-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Matroskademux) SetMaxGapTime(value uint64) error {
	return e.Element.SetProperty("max-gap-time", value)
}

