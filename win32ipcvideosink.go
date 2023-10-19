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

type Win32Ipcvideosink struct {
	*base.GstBaseSink
}

func NewWin32Ipcvideosink() (*Win32Ipcvideosink, error) {
	e, err := gst.NewElement("win32ipcvideosink")
	if err != nil {
		return nil, err
	}
	return &Win32Ipcvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewWin32IpcvideosinkWithName(name string) (*Win32Ipcvideosink, error) {
	e, err := gst.NewElementWithName("win32ipcvideosink", name)
	if err != nil {
		return nil, err
	}
	return &Win32Ipcvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// pipe-name (string)
//
// The name of Win32 named pipe to communicate with clients. Validation of the pipe name is caller's responsibility

func (e *Win32Ipcvideosink) GetPipeName() (string, error) {
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

func (e *Win32Ipcvideosink) SetPipeName(value string) error {
	return e.Element.SetProperty("pipe-name", value)
}

