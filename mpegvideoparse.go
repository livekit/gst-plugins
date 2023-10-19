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

type Mpegvideoparse struct {
	*base.GstBaseParse
}

func NewMpegvideoparse() (*Mpegvideoparse, error) {
	e, err := gst.NewElement("mpegvideoparse")
	if err != nil {
		return nil, err
	}
	return &Mpegvideoparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

func NewMpegvideoparseWithName(name string) (*Mpegvideoparse, error) {
	e, err := gst.NewElementWithName("mpegvideoparse", name)
	if err != nil {
		return nil, err
	}
	return &Mpegvideoparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

// ----- Properties -----

// drop (bool)
//
// Drop data until valid configuration data is received either in the stream or through caps

func (e *Mpegvideoparse) GetDrop() (bool, error) {
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

func (e *Mpegvideoparse) SetDrop(value bool) error {
	return e.Element.SetProperty("drop", value)
}

// gop-split (bool)
//
// Split frame when encountering GOP

func (e *Mpegvideoparse) GetGopSplit() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gop-split")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpegvideoparse) SetGopSplit(value bool) error {
	return e.Element.SetProperty("gop-split", value)
}

