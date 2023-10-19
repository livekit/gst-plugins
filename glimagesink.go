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

type Glimagesink struct {
	Element *gst.Element
}

func NewGlimagesink() (*Glimagesink, error) {
	e, err := gst.NewElement("glimagesink")
	if err != nil {
		return nil, err
	}
	return &Glimagesink{Element: e}, nil
}

func NewGlimagesinkWithName(name string) (*Glimagesink, error) {
	e, err := gst.NewElementWithName("glimagesink", name)
	if err != nil {
		return nil, err
	}
	return &Glimagesink{Element: e}, nil
}

// ----- Properties -----

// context (GstGLContext)
//
// Get OpenGL context

func (e *Glimagesink) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Glimagesink) GetForceAspectRatio() (bool, error) {
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

func (e *Glimagesink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// handle-events (bool)
//
// When enabled, XEvents will be selected and handled

func (e *Glimagesink) GetHandleEvents() (bool, error) {
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

func (e *Glimagesink) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

// ignore-alpha (bool)
//
// When enabled, alpha will be ignored and converted to black

func (e *Glimagesink) GetIgnoreAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Glimagesink) SetIgnoreAlpha(value bool) error {
	return e.Element.SetProperty("ignore-alpha", value)
}

// output-multiview-downmix-mode (GstGLStereoDownmix)
//
// Output anaglyph type to generate when downmixing to mono

func (e *Glimagesink) GetOutputMultiviewDownmixMode() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-downmix-mode")
}

func (e *Glimagesink) SetOutputMultiviewDownmixMode(value interface{}) error {
	return e.Element.SetProperty("output-multiview-downmix-mode", value)
}

// output-multiview-flags (GstVideoMultiviewFlags)
//
// Output multiview layout modifier flags

func (e *Glimagesink) GetOutputMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-flags")
}

func (e *Glimagesink) SetOutputMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("output-multiview-flags", value)
}

// output-multiview-mode (GstVideoMultiviewMode)
//
// Choose output mode for multiview/3D video

func (e *Glimagesink) GetOutputMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-mode")
}

func (e *Glimagesink) SetOutputMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("output-multiview-mode", value)
}

// pixel-aspect-ratio (GstFraction)
//
// The pixel aspect ratio of the device

func (e *Glimagesink) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

func (e *Glimagesink) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// render-rectangle (GstValueArray)
//
// The render rectangle ('<x, y, width, height>')

func (e *Glimagesink) GetRenderRectangle() (interface{}, error) {
	return e.Element.GetProperty("render-rectangle")
}

func (e *Glimagesink) SetRenderRectangle(value interface{}) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// rotate-method (GstGLRotateMethod)
//
// rotate method

func (e *Glimagesink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

func (e *Glimagesink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

// show-preroll-frame (bool)
//
// Whether to render video frames during preroll

func (e *Glimagesink) GetShowPrerollFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-preroll-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Glimagesink) SetShowPrerollFrame(value bool) error {
	return e.Element.SetProperty("show-preroll-frame", value)
}

