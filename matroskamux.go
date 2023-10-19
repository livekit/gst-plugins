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

type Matroskamux struct {
	Element *gst.Element
}

func NewMatroskamux() (*Matroskamux, error) {
	e, err := gst.NewElement("matroskamux")
	if err != nil {
		return nil, err
	}
	return &Matroskamux{Element: e}, nil
}

func NewMatroskamuxWithName(name string) (*Matroskamux, error) {
	e, err := gst.NewElementWithName("matroskamux", name)
	if err != nil {
		return nil, err
	}
	return &Matroskamux{Element: e}, nil
}

// ----- Properties -----

// cluster-timestamp-offset (uint64)
//
// An offset to add to all clusters/blocks (in nanoseconds)

func (e *Matroskamux) GetClusterTimestampOffset() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("cluster-timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Matroskamux) SetClusterTimestampOffset(value uint64) error {
	return e.Element.SetProperty("cluster-timestamp-offset", value)
}

// creation-time (GstGdateTime)
//
// Date and time of creation. This will be used for the DateUTC field. NULL means that the current time will be used.

func (e *Matroskamux) GetCreationTime() (interface{}, error) {
	return e.Element.GetProperty("creation-time")
}

func (e *Matroskamux) SetCreationTime(value interface{}) error {
	return e.Element.SetProperty("creation-time", value)
}

// max-cluster-duration (int64)
//
// A new cluster will be created if its duration exceeds this value. 0 means no maximum duration.

func (e *Matroskamux) GetMaxClusterDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-cluster-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Matroskamux) SetMaxClusterDuration(value int64) error {
	return e.Element.SetProperty("max-cluster-duration", value)
}

// min-cluster-duration (int64)
//
// Desired cluster duration as nanoseconds. A new cluster will be created irrespective of this property if a force key unit event is received. 0 means create a new cluster for each video keyframe or for each audio buffer in audio only streams.

func (e *Matroskamux) GetMinClusterDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("min-cluster-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Matroskamux) SetMinClusterDuration(value int64) error {
	return e.Element.SetProperty("min-cluster-duration", value)
}

// min-index-interval (int64)
//
// An index entry is created every so many nanoseconds.

func (e *Matroskamux) GetMinIndexInterval() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("min-index-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Matroskamux) SetMinIndexInterval(value int64) error {
	return e.Element.SetProperty("min-index-interval", value)
}

// offset-to-zero (bool)
//
// Offsets all streams so that the earliest stream starts at 0.

func (e *Matroskamux) GetOffsetToZero() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("offset-to-zero")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Matroskamux) SetOffsetToZero(value bool) error {
	return e.Element.SetProperty("offset-to-zero", value)
}

// streamable (bool)
//
// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written.

func (e *Matroskamux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Matroskamux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

// timecodescale (int64)
//
// TimecodeScale used to calculate the Raw Timecode of a Block

func (e *Matroskamux) GetTimecodescale() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timecodescale")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Matroskamux) SetTimecodescale(value int64) error {
	return e.Element.SetProperty("timecodescale", value)
}

// version (int)
//
// This parameter determines what Matroska features can be used.

func (e *Matroskamux) GetVersion() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("version")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Matroskamux) SetVersion(value int) error {
	return e.Element.SetProperty("version", value)
}

// writing-app (string)
//
// The name the application that creates the matroska file.

func (e *Matroskamux) GetWritingApp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("writing-app")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Matroskamux) SetWritingApp(value string) error {
	return e.Element.SetProperty("writing-app", value)
}

