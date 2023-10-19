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

type Dx9Screencapsrc struct {
	*base.GstPushSrc
}

func NewDx9Screencapsrc() (*Dx9Screencapsrc, error) {
	e, err := gst.NewElement("dx9screencapsrc")
	if err != nil {
		return nil, err
	}
	return &Dx9Screencapsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDx9ScreencapsrcWithName(name string) (*Dx9Screencapsrc, error) {
	e, err := gst.NewElementWithName("dx9screencapsrc", name)
	if err != nil {
		return nil, err
	}
	return &Dx9Screencapsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// cursor (bool)
//
// Whether to show mouse cursor (default off)

func (e *Dx9Screencapsrc) GetCursor() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cursor")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dx9Screencapsrc) SetCursor(value bool) error {
	return e.Element.SetProperty("cursor", value)
}

// height (int)
//
// Height of screen capture area (0 = maximum)

func (e *Dx9Screencapsrc) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dx9Screencapsrc) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// monitor (int)
//
// Which monitor to use (0 = 1st monitor and default)

func (e *Dx9Screencapsrc) GetMonitor() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("monitor")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dx9Screencapsrc) SetMonitor(value int) error {
	return e.Element.SetProperty("monitor", value)
}

// width (int)
//
// Width of screen capture area (0 = maximum)

func (e *Dx9Screencapsrc) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dx9Screencapsrc) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

// x (int)
//
// Horizontal coordinate of top left corner for the screen capture area

func (e *Dx9Screencapsrc) GetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dx9Screencapsrc) SetX(value int) error {
	return e.Element.SetProperty("x", value)
}

// y (int)
//
// Vertical coordinate of top left corner for the screen capture area

func (e *Dx9Screencapsrc) GetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dx9Screencapsrc) SetY(value int) error {
	return e.Element.SetProperty("y", value)
}

