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

type Fdsrc struct {
	*base.GstPushSrc
}

func NewFdsrc() (*Fdsrc, error) {
	e, err := gst.NewElement("fdsrc")
	if err != nil {
		return nil, err
	}
	return &Fdsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFdsrcWithName(name string) (*Fdsrc, error) {
	e, err := gst.NewElementWithName("fdsrc", name)
	if err != nil {
		return nil, err
	}
	return &Fdsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// fd (int)
//
// An open file descriptor to read from

func (e *Fdsrc) GetFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fdsrc) SetFd(value int) error {
	return e.Element.SetProperty("fd", value)
}

// timeout (uint64)
//
// Post a message after timeout microseconds

func (e *Fdsrc) GetTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fdsrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

