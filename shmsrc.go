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

type Shmsrc struct {
	*base.GstPushSrc
}

func NewShmsrc() (*Shmsrc, error) {
	e, err := gst.NewElement("shmsrc")
	if err != nil {
		return nil, err
	}
	return &Shmsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewShmsrcWithName(name string) (*Shmsrc, error) {
	e, err := gst.NewElementWithName("shmsrc", name)
	if err != nil {
		return nil, err
	}
	return &Shmsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// is-live (bool)
//
// True if the element cannot produce data in PAUSED

func (e *Shmsrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Shmsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// shm-area-name (string)
//
// The name of the shared memory area used to get buffers

func (e *Shmsrc) GetShmAreaName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("shm-area-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// socket-path (string)
//
// The path to the control socket used to control the shared memory

func (e *Shmsrc) GetSocketPath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("socket-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shmsrc) SetSocketPath(value string) error {
	return e.Element.SetProperty("socket-path", value)
}

