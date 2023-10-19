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

type Win32Ipcvideosrc struct {
	*base.GstBaseSrc
}

func NewWin32Ipcvideosrc() (*Win32Ipcvideosrc, error) {
	e, err := gst.NewElement("win32ipcvideosrc")
	if err != nil {
		return nil, err
	}
	return &Win32Ipcvideosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewWin32IpcvideosrcWithName(name string) (*Win32Ipcvideosrc, error) {
	e, err := gst.NewElementWithName("win32ipcvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &Win32Ipcvideosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// pipe-name (string)
//
// The name of Win32 named pipe to communicate with server. Validation of the pipe name is caller's responsibility

func (e *Win32Ipcvideosrc) GetPipeName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pipe-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Win32Ipcvideosrc) SetPipeName(value string) error {
	return e.Element.SetProperty("pipe-name", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *Win32Ipcvideosrc) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Win32Ipcvideosrc) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

