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

type Shmsink struct {
	*base.GstBaseSink
}

func NewShmsink() (*Shmsink, error) {
	e, err := gst.NewElement("shmsink")
	if err != nil {
		return nil, err
	}
	return &Shmsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewShmsinkWithName(name string) (*Shmsink, error) {
	e, err := gst.NewElementWithName("shmsink", name)
	if err != nil {
		return nil, err
	}
	return &Shmsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// buffer-time (int64)
//
// Maximum Size of the shm buffer in nanoseconds (-1 to disable)

func (e *Shmsink) GetBufferTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Shmsink) SetBufferTime(value int64) error {
	return e.Element.SetProperty("buffer-time", value)
}

// perms (uint)
//
// Permissions to set on the shm area

func (e *Shmsink) GetPerms() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("perms")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Shmsink) SetPerms(value uint) error {
	return e.Element.SetProperty("perms", value)
}

// shm-size (uint)
//
// Size of the shared memory area

func (e *Shmsink) GetShmSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shm-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Shmsink) SetShmSize(value uint) error {
	return e.Element.SetProperty("shm-size", value)
}

// socket-path (string)
//
// The path to the control socket used to control the shared memory transport. This may be modified during the NULL->READY transition

func (e *Shmsink) GetSocketPath() (string, error) {
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

func (e *Shmsink) SetSocketPath(value string) error {
	return e.Element.SetProperty("socket-path", value)
}

// wait-for-connection (bool)
//
// Block the stream until the shm pipe is connected

func (e *Shmsink) GetWaitForConnection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-for-connection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Shmsink) SetWaitForConnection(value bool) error {
	return e.Element.SetProperty("wait-for-connection", value)
}

