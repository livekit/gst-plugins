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

type Xvimagesink struct {
	*base.GstBaseSink
}

func NewXvimagesink() (*Xvimagesink, error) {
	e, err := gst.NewElement("xvimagesink")
	if err != nil {
		return nil, err
	}
	return &Xvimagesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewXvimagesinkWithName(name string) (*Xvimagesink, error) {
	e, err := gst.NewElementWithName("xvimagesink", name)
	if err != nil {
		return nil, err
	}
	return &Xvimagesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// autopaint-colorkey (bool)
//
// Whether to autofill overlay with colorkey

func (e *Xvimagesink) GetAutopaintColorkey() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("autopaint-colorkey")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Xvimagesink) SetAutopaintColorkey(value bool) error {
	return e.Element.SetProperty("autopaint-colorkey", value)
}

// brightness (int)
//
// The brightness of the video

func (e *Xvimagesink) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Xvimagesink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

// colorkey (int)
//
// Color to use for the overlay mask.

func (e *Xvimagesink) GetColorkey() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorkey")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Xvimagesink) SetColorkey(value int) error {
	return e.Element.SetProperty("colorkey", value)
}

// contrast (int)
//
// The contrast of the video

func (e *Xvimagesink) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Xvimagesink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

// device (string)
//
// The number of the video adaptor

func (e *Xvimagesink) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Xvimagesink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// The name of the video adaptor

func (e *Xvimagesink) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// display (string)
//
// X Display name

func (e *Xvimagesink) GetDisplay() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Xvimagesink) SetDisplay(value string) error {
	return e.Element.SetProperty("display", value)
}

// double-buffer (bool)
//
// Whether to double-buffer the output.

func (e *Xvimagesink) GetDoubleBuffer() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("double-buffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Xvimagesink) SetDoubleBuffer(value bool) error {
	return e.Element.SetProperty("double-buffer", value)
}

// draw-borders (bool)
//
// Draw black borders when using GstXvImageSink:force-aspect-ratio to fill
// unused parts of the video area.

func (e *Xvimagesink) GetDrawBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Xvimagesink) SetDrawBorders(value bool) error {
	return e.Element.SetProperty("draw-borders", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Xvimagesink) GetForceAspectRatio() (bool, error) {
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

func (e *Xvimagesink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// handle-events (bool)
//
// When enabled, XEvents will be selected and handled

func (e *Xvimagesink) GetHandleEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Xvimagesink) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

// handle-expose (bool)
//
// When enabled, the current frame will always be drawn in response to X
// Expose.

func (e *Xvimagesink) GetHandleExpose() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-expose")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Xvimagesink) SetHandleExpose(value bool) error {
	return e.Element.SetProperty("handle-expose", value)
}

// hue (int)
//
// The hue of the video

func (e *Xvimagesink) GetHue() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Xvimagesink) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

// pixel-aspect-ratio (string)
//
// The pixel aspect ratio of the device

func (e *Xvimagesink) GetPixelAspectRatio() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pixel-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Xvimagesink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// render-rectangle (GstValueArray)
//
// The render rectangle ('<x, y, width, height>')

func (e *Xvimagesink) GetRenderRectangle() (interface{}, error) {
	return e.Element.GetProperty("render-rectangle")
}

func (e *Xvimagesink) SetRenderRectangle(value interface{}) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// saturation (int)
//
// The saturation of the video

func (e *Xvimagesink) GetSaturation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Xvimagesink) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

// synchronous (bool)
//
// When enabled, runs the X display in synchronous mode. (unrelated to A/V sync, used only for debugging)

func (e *Xvimagesink) GetSynchronous() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("synchronous")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Xvimagesink) SetSynchronous(value bool) error {
	return e.Element.SetProperty("synchronous", value)
}

// window-height (uint64)
//
// Actual height of the video window.

func (e *Xvimagesink) GetWindowHeight() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// window-width (uint64)
//
// Actual width of the video window.

func (e *Xvimagesink) GetWindowWidth() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

