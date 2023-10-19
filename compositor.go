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

type Compositor struct {
	*base.GstAggregator
}

func NewCompositor() (*Compositor, error) {
	e, err := gst.NewElement("compositor")
	if err != nil {
		return nil, err
	}
	return &Compositor{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewCompositorWithName(name string) (*Compositor, error) {
	e, err := gst.NewElementWithName("compositor", name)
	if err != nil {
		return nil, err
	}
	return &Compositor{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// background (GstCompositorBackground)
//
// Background type

func (e *Compositor) GetBackground() (GstCompositorBackground, error) {
	var value GstCompositorBackground
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCompositorBackground)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCompositorBackground")
	}
	return value, nil
}

func (e *Compositor) SetBackground(value GstCompositorBackground) error {
	e.Element.SetArg("background", string(value))
	return nil
}

// ignore-inactive-pads (bool)
//
// Don't wait for inactive pads when live. An inactive pad
// is a pad that hasn't yet received a buffer, but that has
// been waited on at least once.

// max-threads (uint)
//
// Maximum number of blending/rendering worker threads to spawn (0 = auto)

func (e *Compositor) GetMaxThreads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Compositor) SetMaxThreads(value uint) error {
	return e.Element.SetProperty("max-threads", value)
}

// zero-size-is-unscaled (bool)
//
// Whether a pad with height or width 0 should be left unscaled
// in that dimension, or simply not composited in. Setting it to
// FALSE might be useful when animating those properties.

func (e *Compositor) GetZeroSizeIsUnscaled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-size-is-unscaled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Compositor) SetZeroSizeIsUnscaled(value bool) error {
	return e.Element.SetProperty("zero-size-is-unscaled", value)
}

// alpha (float64)
//
// Alpha of the picture

func (e *Compositor) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Compositor) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

// height (int)
//
// Height of the picture

func (e *Compositor) GetHeight() (int, error) {
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

func (e *Compositor) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// operator (GstCompositorOperator)
//
// Blending operator to use for blending this pad over the previous ones

func (e *Compositor) GetOperator() (GstCompositorOperator, error) {
	var value GstCompositorOperator
	var ok bool
	v, err := e.Element.GetProperty("operator")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCompositorOperator)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCompositorOperator")
	}
	return value, nil
}

func (e *Compositor) SetOperator(value GstCompositorOperator) error {
	e.Element.SetArg("operator", string(value))
	return nil
}

// sizing-policy (GstCompositorSizingPolicy)
//
// Specifies sizing policy to use. Depending on selected sizing policy,
// scaled image might not fully cover the configured target rectangle area
// (e.g., "keep-aspect-ratio"). In that case, any uncovered area will be
// filled with background unless the uncovered area is drawn by other image.

func (e *Compositor) GetSizingPolicy() (GstCompositorSizingPolicy, error) {
	var value GstCompositorSizingPolicy
	var ok bool
	v, err := e.Element.GetProperty("sizing-policy")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCompositorSizingPolicy)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCompositorSizingPolicy")
	}
	return value, nil
}

func (e *Compositor) SetSizingPolicy(value GstCompositorSizingPolicy) error {
	e.Element.SetArg("sizing-policy", string(value))
	return nil
}

// width (int)
//
// Width of the picture

func (e *Compositor) GetWidth() (int, error) {
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

func (e *Compositor) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

// xpos (int)
//
// X Position of the picture

func (e *Compositor) GetXpos() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xpos")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Compositor) SetXpos(value int) error {
	return e.Element.SetProperty("xpos", value)
}

// ypos (int)
//
// Y Position of the picture

func (e *Compositor) GetYpos() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ypos")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Compositor) SetYpos(value int) error {
	return e.Element.SetProperty("ypos", value)
}

// ----- Constants -----

type GstCompositorSizingPolicy string

const (
	GstCompositorSizingPolicyNone GstCompositorSizingPolicy = "none" // none (0) – None: Image is scaled to fill configured destination rectangle without padding or keeping the aspect ratio
	GstCompositorSizingPolicyKeepAspectRatio GstCompositorSizingPolicy = "keep-aspect-ratio" // keep-aspect-ratio (1) – Keep Aspect Ratio: Image is scaled to fit destination rectangle specified by GstCompositorPad:{xpos, ypos, width, height} with preserved aspect ratio. Resulting image will be centered in the destination rectangle with padding if necessary
)

type GstCompositorBackground string

const (
	GstCompositorBackgroundChecker GstCompositorBackground = "checker" // checker (0) – Checker pattern
	GstCompositorBackgroundBlack GstCompositorBackground = "black" // black (1) – Black
	GstCompositorBackgroundWhite GstCompositorBackground = "white" // white (2) – White
	GstCompositorBackgroundTransparent GstCompositorBackground = "transparent" // transparent (3) – Transparent Background to enable further compositing
)

type GstCompositorOperator string

const (
	GstCompositorOperatorSource GstCompositorOperator = "source" // source (0) – Source
	GstCompositorOperatorOver GstCompositorOperator = "over" // over (1) – Over
	GstCompositorOperatorAdd GstCompositorOperator = "add" // add (2) – Add
)

