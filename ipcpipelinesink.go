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

type Ipcpipelinesink struct {
	Element *gst.Element
}

func NewIpcpipelinesink() (*Ipcpipelinesink, error) {
	e, err := gst.NewElement("ipcpipelinesink")
	if err != nil {
		return nil, err
	}
	return &Ipcpipelinesink{Element: e}, nil
}

func NewIpcpipelinesinkWithName(name string) (*Ipcpipelinesink, error) {
	e, err := gst.NewElementWithName("ipcpipelinesink", name)
	if err != nil {
		return nil, err
	}
	return &Ipcpipelinesink{Element: e}, nil
}

// ----- Properties -----

// ack-time (uint64)
//
// Maximum time to wait for a response to a message

func (e *Ipcpipelinesink) GetAckTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ack-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Ipcpipelinesink) SetAckTime(value uint64) error {
	return e.Element.SetProperty("ack-time", value)
}

// fdin (int)
//
// File descriptor to received data from

func (e *Ipcpipelinesink) GetFdin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fdin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Ipcpipelinesink) SetFdin(value int) error {
	return e.Element.SetProperty("fdin", value)
}

// fdout (int)
//
// File descriptor to send data through

func (e *Ipcpipelinesink) GetFdout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fdout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Ipcpipelinesink) SetFdout(value int) error {
	return e.Element.SetProperty("fdout", value)
}

// read-chunk-size (uint)
//
// Read chunk size

func (e *Ipcpipelinesink) GetReadChunkSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("read-chunk-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ipcpipelinesink) SetReadChunkSize(value uint) error {
	return e.Element.SetProperty("read-chunk-size", value)
}

