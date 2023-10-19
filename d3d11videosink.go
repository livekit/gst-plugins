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

type D3D11Videosink struct {
	*base.GstBaseSink
}

func NewD3D11Videosink() (*D3D11Videosink, error) {
	e, err := gst.NewElement("d3d11videosink")
	if err != nil {
		return nil, err
	}
	return &D3D11Videosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewD3D11VideosinkWithName(name string) (*D3D11Videosink, error) {
	e, err := gst.NewElementWithName("d3d11videosink", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Videosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// adapter (int)
//
// Adapter index for creating device (-1 for default)

func (e *D3D11Videosink) GetAdapter() (int, error) {
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

func (e *D3D11Videosink) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

// display-format (GstD3D11VideoSinkDisplayFormat)
//
// Swapchain display format

func (e *D3D11Videosink) GetDisplayFormat() (GstD3D11VideoSinkDisplayFormat, error) {
	var value GstD3D11VideoSinkDisplayFormat
	var ok bool
	v, err := e.Element.GetProperty("display-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11VideoSinkDisplayFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11VideoSinkDisplayFormat")
	}
	return value, nil
}

func (e *D3D11Videosink) SetDisplayFormat(value GstD3D11VideoSinkDisplayFormat) error {
	e.Element.SetArg("display-format", string(value))
	return nil
}

// draw-on-shared-texture (bool)
//
// Instruct the sink to draw on a shared texture provided by user.
// User must watch begin-draw signal and should call
// draw method on the begin-draw
// signal handler.

// emit-present (bool)
//
// Emits "present" signal

func (e *D3D11Videosink) GetEmitPresent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("emit-present")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Videosink) SetEmitPresent(value bool) error {
	return e.Element.SetProperty("emit-present", value)
}

// enable-navigation-events (bool)
//
// When enabled, navigation events are sent upstream

func (e *D3D11Videosink) GetEnableNavigationEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-navigation-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Videosink) SetEnableNavigationEvents(value bool) error {
	return e.Element.SetProperty("enable-navigation-events", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *D3D11Videosink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Videosink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// fullscreen (bool)
//
// Ignored when "fullscreen-toggle-mode" does not include "property"

func (e *D3D11Videosink) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Videosink) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

// fullscreen-toggle-mode (GstD3D11WindowFullscreenToggleMode)
//
// Full screen toggle mode used to trigger fullscreen mode change

func (e *D3D11Videosink) GetFullscreenToggleMode() (GstD3D11WindowFullscreenToggleMode, error) {
	var value GstD3D11WindowFullscreenToggleMode
	var ok bool
	v, err := e.Element.GetProperty("fullscreen-toggle-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11WindowFullscreenToggleMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11WindowFullscreenToggleMode")
	}
	return value, nil
}

func (e *D3D11Videosink) SetFullscreenToggleMode(value GstD3D11WindowFullscreenToggleMode) error {
	e.Element.SetArg("fullscreen-toggle-mode", string(value))
	return nil
}

// gamma-mode (GstVideoGammaMode)
//
// Gamma conversion mode

func (e *D3D11Videosink) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

func (e *D3D11Videosink) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

// primaries-mode (GstVideoPrimariesMode)
//
// Primaries conversion mode

func (e *D3D11Videosink) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

func (e *D3D11Videosink) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

// render-rectangle (GstValueArray)

func (e *D3D11Videosink) GetRenderRectangle() (interface{}, error) {
	return e.Element.GetProperty("render-rectangle")
}

func (e *D3D11Videosink) SetRenderRectangle(value interface{}) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// rotate-method (GstVideoOrientationMethod)
//
// Video rotation/flip method to use

func (e *D3D11Videosink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

func (e *D3D11Videosink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

// ----- Constants -----

type GstD3D11VideoSinkDisplayFormat string

const (
	GstD3D11VideoSinkDisplayFormatUnknown GstD3D11VideoSinkDisplayFormat = "unknown" // unknown (0) – DXGI_FORMAT_UNKNOWN
	GstD3D11VideoSinkDisplayFormatR10G10B10A2Unorm GstD3D11VideoSinkDisplayFormat = "r10g10b10a2-unorm" // r10g10b10a2-unorm (24) – DXGI_FORMAT_R10G10B10A2_UNORM
	GstD3D11VideoSinkDisplayFormatR8G8B8A8Unorm GstD3D11VideoSinkDisplayFormat = "r8g8b8a8-unorm" // r8g8b8a8-unorm (28) – DXGI_FORMAT_R8G8B8A8_UNORM
	GstD3D11VideoSinkDisplayFormatB8G8R8A8Unorm GstD3D11VideoSinkDisplayFormat = "b8g8r8a8-unorm" // b8g8r8a8-unorm (87) – DXGI_FORMAT_B8G8R8A8_UNORM
)

type GstD3D11WindowFullscreenToggleMode string

const (
	GstD3D11WindowFullscreenToggleModeNone GstD3D11WindowFullscreenToggleMode = "none" // none (0x00000000) – GST_D3D11_WINDOW_FULLSCREEN_TOGGLE_MODE_NONE
	GstD3D11WindowFullscreenToggleModeAltEnter GstD3D11WindowFullscreenToggleMode = "alt-enter" // alt-enter (0x00000002) – GST_D3D11_WINDOW_FULLSCREEN_TOGGLE_MODE_ALT_ENTER
	GstD3D11WindowFullscreenToggleModeProperty GstD3D11WindowFullscreenToggleMode = "property" // property (0x00000004) – GST_D3D11_WINDOW_FULLSCREEN_TOGGLE_MODE_PROPERTY
)

