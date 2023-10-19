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

type Openni2Src struct {
	*base.GstPushSrc
}

func NewOpenni2Src() (*Openni2Src, error) {
	e, err := gst.NewElement("openni2src")
	if err != nil {
		return nil, err
	}
	return &Openni2Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewOpenni2SrcWithName(name string) (*Openni2Src, error) {
	e, err := gst.NewElementWithName("openni2src", name)
	if err != nil {
		return nil, err
	}
	return &Openni2Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// blocksize (uint)
//
// Size in bytes to read per buffer (-1 = default)

func (e *Openni2Src) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openni2Src) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// do-timestamp (bool)
//
// Apply current stream time to buffers

func (e *Openni2Src) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Openni2Src) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

// location (string)
//
// Source uri, can be a file or a device.

func (e *Openni2Src) GetLocation() (string, error) {
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

func (e *Openni2Src) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// num-buffers (int)
//
// Number of buffers to output before sending EOS (-1 = unlimited)

func (e *Openni2Src) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openni2Src) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// sourcetype (GstOpenni2SrcSourcetype)
//
// Type of readings to get from the source

func (e *Openni2Src) GetSourcetype() (GstOpenni2SrcSourcetype, error) {
	var value GstOpenni2SrcSourcetype
	var ok bool
	v, err := e.Element.GetProperty("sourcetype")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenni2SrcSourcetype)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenni2SrcSourcetype")
	}
	return value, nil
}

func (e *Openni2Src) SetSourcetype(value GstOpenni2SrcSourcetype) error {
	e.Element.SetArg("sourcetype", string(value))
	return nil
}

// typefind (bool)
//
// Run typefind before negotiating (deprecated, non-functional)

func (e *Openni2Src) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Openni2Src) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

// ----- Constants -----

type GstOpenni2SrcSourcetype string

const (
	GstOpenni2SrcSourcetypeDepth GstOpenni2SrcSourcetype = "depth" // depth (0) – Get depth readings
	GstOpenni2SrcSourcetypeColor GstOpenni2SrcSourcetype = "color" // color (1) – Get color readings
	GstOpenni2SrcSourcetypeBoth GstOpenni2SrcSourcetype = "both" // both (2) – Get color and depth (as alpha) readings - EXPERIMENTAL
)

