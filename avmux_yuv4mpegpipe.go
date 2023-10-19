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

type AvmuxYuv4Mpegpipe struct {
	Element *gst.Element
}

func NewAvmuxYuv4Mpegpipe() (*AvmuxYuv4Mpegpipe, error) {
	e, err := gst.NewElement("avmux_yuv4mpegpipe")
	if err != nil {
		return nil, err
	}
	return &AvmuxYuv4Mpegpipe{Element: e}, nil
}

func NewAvmuxYuv4MpegpipeWithName(name string) (*AvmuxYuv4Mpegpipe, error) {
	e, err := gst.NewElementWithName("avmux_yuv4mpegpipe", name)
	if err != nil {
		return nil, err
	}
	return &AvmuxYuv4Mpegpipe{Element: e}, nil
}

// ----- Properties -----

// maxdelay (int)
//
// Set the maximum demux-decode delay (in microseconds)

func (e *AvmuxYuv4Mpegpipe) GetMaxdelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("maxdelay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvmuxYuv4Mpegpipe) SetMaxdelay(value int) error {
	return e.Element.SetProperty("maxdelay", value)
}

// preload (int)
//
// Set the initial demux-decode delay (in microseconds)

func (e *AvmuxYuv4Mpegpipe) GetPreload() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("preload")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvmuxYuv4Mpegpipe) SetPreload(value int) error {
	return e.Element.SetProperty("preload", value)
}

