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

type D3D11Ipcsrc struct {
	*base.GstBaseSrc
}

func NewD3D11Ipcsrc() (*D3D11Ipcsrc, error) {
	e, err := gst.NewElement("d3d11ipcsrc")
	if err != nil {
		return nil, err
	}
	return &D3D11Ipcsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewD3D11IpcsrcWithName(name string) (*D3D11Ipcsrc, error) {
	e, err := gst.NewElementWithName("d3d11ipcsrc", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Ipcsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// connection-timeout (uint)
//
// Connection timeout in seconds (0 = never timeout)

func (e *D3D11Ipcsrc) GetConnectionTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("connection-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *D3D11Ipcsrc) SetConnectionTimeout(value uint) error {
	return e.Element.SetProperty("connection-timeout", value)
}

// io-mode (GstD3D11IpcIomode)
//
// Memory I/O mode to use

func (e *D3D11Ipcsrc) GetIoMode() (GstD3D11IpcIomode, error) {
	var value GstD3D11IpcIomode
	var ok bool
	v, err := e.Element.GetProperty("io-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11IpcIomode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11IpcIomode")
	}
	return value, nil
}

func (e *D3D11Ipcsrc) SetIoMode(value GstD3D11IpcIomode) error {
	e.Element.SetArg("io-mode", string(value))
	return nil
}

// pipe-name (string)
//
// The name of Win32 named pipe to communicate with clients. Validation of the pipe name is caller's responsibility

func (e *D3D11Ipcsrc) GetPipeName() (string, error) {
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

func (e *D3D11Ipcsrc) SetPipeName(value string) error {
	return e.Element.SetProperty("pipe-name", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *D3D11Ipcsrc) GetProcessingDeadline() (uint64, error) {
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

func (e *D3D11Ipcsrc) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// ----- Constants -----

type GstD3D11IpcIomode string

const (
	GstD3D11IpcIomodeCopy GstD3D11IpcIomode = "copy" // copy (0) – Copy remote texture
	GstD3D11IpcIomodeImport GstD3D11IpcIomode = "import" // import (1) – Import remote texture
)

