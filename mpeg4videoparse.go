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

	"github.com/livekit/gstplugins/base"
)

type Mpeg4Videoparse struct {
	*base.GstBaseParse
}

func NewMpeg4Videoparse() (*Mpeg4Videoparse, error) {
	e, err := gst.NewElement("mpeg4videoparse")
	if err != nil {
		return nil, err
	}
	return &Mpeg4Videoparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

func NewMpeg4VideoparseWithName(name string) (*Mpeg4Videoparse, error) {
	e, err := gst.NewElementWithName("mpeg4videoparse", name)
	if err != nil {
		return nil, err
	}
	return &Mpeg4Videoparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

// ----- Properties -----

// config-interval (int)
//
// Send Configuration Insertion Interval in seconds (configuration headers will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)

func (e *Mpeg4Videoparse) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg4Videoparse) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

// drop (bool)
//
// Drop data until valid configuration data is received either in the stream or through caps

func (e *Mpeg4Videoparse) GetDrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg4Videoparse) SetDrop(value bool) error {
	return e.Element.SetProperty("drop", value)
}

