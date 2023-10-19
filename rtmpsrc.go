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

type Rtmpsrc struct {
	*base.GstPushSrc
}

func NewRtmpsrc() (*Rtmpsrc, error) {
	e, err := gst.NewElement("rtmpsrc")
	if err != nil {
		return nil, err
	}
	return &Rtmpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewRtmpsrcWithName(name string) (*Rtmpsrc, error) {
	e, err := gst.NewElementWithName("rtmpsrc", name)
	if err != nil {
		return nil, err
	}
	return &Rtmpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// location (string)
//
// Location of the RTMP url to read

func (e *Rtmpsrc) GetLocation() (string, error) {
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

func (e *Rtmpsrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// timeout (int)
//
// Time without receiving any data from the server to wait before to timeout the session

func (e *Rtmpsrc) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtmpsrc) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

