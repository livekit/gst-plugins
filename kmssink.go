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

type Kmssink struct {
	*base.GstBaseSink
}

func NewKmssink() (*Kmssink, error) {
	e, err := gst.NewElement("kmssink")
	if err != nil {
		return nil, err
	}
	return &Kmssink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewKmssinkWithName(name string) (*Kmssink, error) {
	e, err := gst.NewElementWithName("kmssink", name)
	if err != nil {
		return nil, err
	}
	return &Kmssink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// bus-id (string)
//
// If you have a system with multiple displays for the same driver-name,
// you can choose which display to use by setting the DRM bus ID. Otherwise,
// the driver decides which one.

func (e *Kmssink) GetBusId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("bus-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Kmssink) SetBusId(value string) error {
	return e.Element.SetProperty("bus-id", value)
}

// can-scale (bool)
//
// User can tell kmssink if the driver can support scale.

func (e *Kmssink) GetCanScale() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Kmssink) SetCanScale(value bool) error {
	return e.Element.SetProperty("can-scale", value)
}

// connector-id (int)
//
// A GPU has several output connectors, for example: LVDS, VGA,
// HDMI, etc. By default the first LVDS is tried, then the first
// eDP, and at the end, the first connected one.

func (e *Kmssink) GetConnectorId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("connector-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kmssink) SetConnectorId(value int) error {
	return e.Element.SetProperty("connector-id", value)
}

// connector-properties (GstStructure)
//
// Additional properties for the connector. Keys are strings and values
// unsigned 64 bits integers.

func (e *Kmssink) GetConnectorProperties() (interface{}, error) {
	return e.Element.GetProperty("connector-properties")
}

func (e *Kmssink) SetConnectorProperties(value interface{}) error {
	return e.Element.SetProperty("connector-properties", value)
}

// display-height (int)
//
// Actual height of the display. This is read only and only available in
// PAUSED and PLAYING state. It's meant to be used with
// gst_video_overlay_set_render_rectangle function.

func (e *Kmssink) GetDisplayHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("display-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// display-width (int)
//
// Actual width of the display. This is read only and only available in
// PAUSED and PLAYING state. It's meant to be used with
// gst_video_overlay_set_render_rectangle function.

func (e *Kmssink) GetDisplayWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("display-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// driver-name (string)
//
// If you have a system with multiple GPUs, you can choose which GPU
// to use setting the DRM device driver name. Otherwise, the first
// one from an internal list is used.

func (e *Kmssink) GetDriverName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("driver-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Kmssink) SetDriverName(value string) error {
	return e.Element.SetProperty("driver-name", value)
}

// fd (int)
//
// You can supply your own DRM file descriptor.  By default, the sink will
// open its own DRM file descriptor.

func (e *Kmssink) GetFd() (int, error) {
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

func (e *Kmssink) SetFd(value int) error {
	return e.Element.SetProperty("fd", value)
}

// force-modesetting (bool)
//
// If the output connector is already active, the sink automatically uses an
// overlay plane. Enforce mode setting in the kms sink and output to the
// base plane to override the automatic behavior.

func (e *Kmssink) GetForceModesetting() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-modesetting")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Kmssink) SetForceModesetting(value bool) error {
	return e.Element.SetProperty("force-modesetting", value)
}

// plane-id (int)
//
// There could be several planes associated with a CRTC.
// By default the first plane that's possible to use with a given
// CRTC is tried.

func (e *Kmssink) GetPlaneId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("plane-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kmssink) SetPlaneId(value int) error {
	return e.Element.SetProperty("plane-id", value)
}

// plane-properties (GstStructure)
//
// Additional properties for the plane. Keys are strings and values
// unsigned 64 bits integers.

func (e *Kmssink) GetPlaneProperties() (interface{}, error) {
	return e.Element.GetProperty("plane-properties")
}

func (e *Kmssink) SetPlaneProperties(value interface{}) error {
	return e.Element.SetProperty("plane-properties", value)
}

// render-rectangle (GstValueArray)
//
// The render rectangle ('<x, y, width, height>')

func (e *Kmssink) GetRenderRectangle() (interface{}, error) {
	return e.Element.GetProperty("render-rectangle")
}

func (e *Kmssink) SetRenderRectangle(value interface{}) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// restore-crtc (bool)
//
// Restore previous CRTC setting if new CRTC mode was set forcefully.
// By default this is enabled if user set CRTC with a new mode on an already
// active CRTC wich was having a valid mode.

func (e *Kmssink) GetRestoreCrtc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("restore-crtc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Kmssink) SetRestoreCrtc(value bool) error {
	return e.Element.SetProperty("restore-crtc", value)
}

// skip-vsync (bool)
//
// For some cases, to suppress internal vsync, which can drop framerate
// in half, set this to 1.

func (e *Kmssink) GetSkipVsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-vsync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Kmssink) SetSkipVsync(value bool) error {
	return e.Element.SetProperty("skip-vsync", value)
}

