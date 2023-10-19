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

type Glimagesinkelement struct {
	*base.GstBaseSink
}

func NewGlimagesinkelement() (*Glimagesinkelement, error) {
	e, err := gst.NewElement("glimagesinkelement")
	if err != nil {
		return nil, err
	}
	return &Glimagesinkelement{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewGlimagesinkelementWithName(name string) (*Glimagesinkelement, error) {
	e, err := gst.NewElementWithName("glimagesinkelement", name)
	if err != nil {
		return nil, err
	}
	return &Glimagesinkelement{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// context (GstGLContext)
//
// Get OpenGL context

func (e *Glimagesinkelement) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Glimagesinkelement) GetForceAspectRatio() (bool, error) {
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

func (e *Glimagesinkelement) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// handle-events (bool)
//
// When enabled, XEvents will be selected and handled

func (e *Glimagesinkelement) GetHandleEvents() (bool, error) {
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

func (e *Glimagesinkelement) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

// ignore-alpha (bool)
//
// When enabled, alpha will be ignored and converted to black

func (e *Glimagesinkelement) GetIgnoreAlpha() (bool, error) {
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

func (e *Glimagesinkelement) SetIgnoreAlpha(value bool) error {
	return e.Element.SetProperty("ignore-alpha", value)
}

// output-multiview-downmix-mode (GstGLStereoDownmix)
//
// Output anaglyph type to generate when downmixing to mono

func (e *Glimagesinkelement) GetOutputMultiviewDownmixMode() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-downmix-mode")
}

func (e *Glimagesinkelement) SetOutputMultiviewDownmixMode(value interface{}) error {
	return e.Element.SetProperty("output-multiview-downmix-mode", value)
}

// output-multiview-flags (GstVideoMultiviewFlags)
//
// Output multiview layout modifier flags

func (e *Glimagesinkelement) GetOutputMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-flags")
}

func (e *Glimagesinkelement) SetOutputMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("output-multiview-flags", value)
}

// output-multiview-mode (GstVideoMultiviewMode)
//
// Choose output mode for multiview/3D video

func (e *Glimagesinkelement) GetOutputMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-mode")
}

func (e *Glimagesinkelement) SetOutputMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("output-multiview-mode", value)
}

// pixel-aspect-ratio (GstFraction)
//
// The pixel aspect ratio of the device

func (e *Glimagesinkelement) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

func (e *Glimagesinkelement) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// render-rectangle (GstValueArray)
//
// The render rectangle ('<x, y, width, height>')

func (e *Glimagesinkelement) GetRenderRectangle() (interface{}, error) {
	return e.Element.GetProperty("render-rectangle")
}

func (e *Glimagesinkelement) SetRenderRectangle(value interface{}) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// rotate-method (GstGLRotateMethod)
//
// rotate method

func (e *Glimagesinkelement) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

func (e *Glimagesinkelement) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

// ----- Constants -----

type GstGlrotateMethod string

const (
	GstGlrotateMethodNone GstGlrotateMethod = "none" // none (0) – Identity (no rotation)
	GstGlrotateMethodClockwise GstGlrotateMethod = "clockwise" // clockwise (1) – Rotate clockwise 90 degrees
	GstGlrotateMethodRotate180 GstGlrotateMethod = "rotate-180" // rotate-180 (2) – Rotate 180 degrees
	GstGlrotateMethodCounterclockwise GstGlrotateMethod = "counterclockwise" // counterclockwise (3) – Rotate counter-clockwise 90 degrees
	GstGlrotateMethodHorizontalFlip GstGlrotateMethod = "horizontal-flip" // horizontal-flip (4) – Flip horizontally
	GstGlrotateMethodVerticalFlip GstGlrotateMethod = "vertical-flip" // vertical-flip (5) – Flip vertically
	GstGlrotateMethodUpperLeftDiagonal GstGlrotateMethod = "upper-left-diagonal" // upper-left-diagonal (6) – Flip across upper left/lower right diagonal
	GstGlrotateMethodUpperRightDiagonal GstGlrotateMethod = "upper-right-diagonal" // upper-right-diagonal (7) – Flip across upper right/lower left diagonal
	GstGlrotateMethodAutomatic GstGlrotateMethod = "automatic" // automatic (8) – Select rotate method based on image-orientation tag
)

