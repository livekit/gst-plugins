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

type D3D11Screencapturesrc struct {
	*base.GstBaseSrc
}

func NewD3D11Screencapturesrc() (*D3D11Screencapturesrc, error) {
	e, err := gst.NewElement("d3d11screencapturesrc")
	if err != nil {
		return nil, err
	}
	return &D3D11Screencapturesrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewD3D11ScreencapturesrcWithName(name string) (*D3D11Screencapturesrc, error) {
	e, err := gst.NewElementWithName("d3d11screencapturesrc", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Screencapturesrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// adapter (int)
//
// DXGI Adapter index for creating device when WGC mode is selected

func (e *D3D11Screencapturesrc) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

// capture-api (GstD3D11ScreenCaptureApi)
//
// Capture API to use

func (e *D3D11Screencapturesrc) GetCaptureApi() (GstD3D11ScreenCaptureApi, error) {
	var value GstD3D11ScreenCaptureApi
	var ok bool
	v, err := e.Element.GetProperty("capture-api")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11ScreenCaptureApi)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11ScreenCaptureApi")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetCaptureApi(value GstD3D11ScreenCaptureApi) error {
	e.Element.SetArg("capture-api", string(value))
	return nil
}

// crop-height (uint)
//
// Height of screen capture area (0 = maximum)

func (e *D3D11Screencapturesrc) GetCropHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetCropHeight(value uint) error {
	return e.Element.SetProperty("crop-height", value)
}

// crop-width (uint)
//
// Width of screen capture area (0 = maximum)

func (e *D3D11Screencapturesrc) GetCropWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetCropWidth(value uint) error {
	return e.Element.SetProperty("crop-width", value)
}

// crop-x (uint)
//
// Horizontal coordinate of top left corner for the screen capture area

func (e *D3D11Screencapturesrc) GetCropX() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetCropX(value uint) error {
	return e.Element.SetProperty("crop-x", value)
}

// crop-y (uint)
//
// Vertical coordinate of top left corner for the screen capture area

func (e *D3D11Screencapturesrc) GetCropY() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetCropY(value uint) error {
	return e.Element.SetProperty("crop-y", value)
}

// monitor-handle (uint64)
//
// A HMONITOR handle of monitor to capture

func (e *D3D11Screencapturesrc) GetMonitorHandle() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("monitor-handle")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetMonitorHandle(value uint64) error {
	return e.Element.SetProperty("monitor-handle", value)
}

// monitor-index (int)
//
// Zero-based index for monitor to capture (-1 = primary monitor)

func (e *D3D11Screencapturesrc) GetMonitorIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("monitor-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetMonitorIndex(value int) error {
	return e.Element.SetProperty("monitor-index", value)
}

// show-border (bool)
//
// Show border lines to capture area when WGC mode is selected.
// This feature requires Windows11 or newer

func (e *D3D11Screencapturesrc) GetShowBorder() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-border")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetShowBorder(value bool) error {
	return e.Element.SetProperty("show-border", value)
}

// show-cursor (bool)
//
// Whether to show mouse cursor

func (e *D3D11Screencapturesrc) GetShowCursor() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-cursor")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetShowCursor(value bool) error {
	return e.Element.SetProperty("show-cursor", value)
}

// window-capture-mode (GstD3D11WindowCaptureMode)
//
// Window capture mode to use

func (e *D3D11Screencapturesrc) GetWindowCaptureMode() (GstD3D11WindowCaptureMode, error) {
	var value GstD3D11WindowCaptureMode
	var ok bool
	v, err := e.Element.GetProperty("window-capture-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11WindowCaptureMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11WindowCaptureMode")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetWindowCaptureMode(value GstD3D11WindowCaptureMode) error {
	e.Element.SetArg("window-capture-mode", string(value))
	return nil
}

// window-handle (uint64)
//
// HWND window handle to capture

func (e *D3D11Screencapturesrc) GetWindowHandle() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-handle")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *D3D11Screencapturesrc) SetWindowHandle(value uint64) error {
	return e.Element.SetProperty("window-handle", value)
}

// ----- Constants -----

type GstD3D11ScreenCaptureApi string

const (
	GstD3D11ScreenCaptureApiDxgi GstD3D11ScreenCaptureApi = "dxgi" // dxgi (0) – DXGI Desktop Duplication
	GstD3D11ScreenCaptureApiWgc GstD3D11ScreenCaptureApi = "wgc" // wgc (1) – Windows Graphics Capture
)

type GstD3D11WindowCaptureMode string

const (
	GstD3D11WindowCaptureModeDefault GstD3D11WindowCaptureMode = "default" // default (0) – Capture entire window area
	GstD3D11WindowCaptureModeClient GstD3D11WindowCaptureMode = "client" // client (1) – Capture client area
)

