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

type Gltransformation struct {
	*base.GstBaseTransform
}

func NewGltransformation() (*Gltransformation, error) {
	e, err := gst.NewElement("gltransformation")
	if err != nil {
		return nil, err
	}
	return &Gltransformation{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGltransformationWithName(name string) (*Gltransformation, error) {
	e, err := gst.NewElementWithName("gltransformation", name)
	if err != nil {
		return nil, err
	}
	return &Gltransformation{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// fov (float32)
//
// Field of view angle in degrees

func (e *Gltransformation) GetFov() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fov")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetFov(value float32) error {
	return e.Element.SetProperty("fov", value)
}

// mvp-matrix (GstGrapheneMatrix)
//
// The final Graphene 4x4 Matrix for transformation

func (e *Gltransformation) GetMvpMatrix() (interface{}, error) {
	return e.Element.GetProperty("mvp-matrix")
}

func (e *Gltransformation) SetMvpMatrix(value interface{}) error {
	return e.Element.SetProperty("mvp-matrix", value)
}

// ortho (bool)
//
// Use orthographic projection

func (e *Gltransformation) GetOrtho() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ortho")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Gltransformation) SetOrtho(value bool) error {
	return e.Element.SetProperty("ortho", value)
}

// pivot-x (float32)
//
// Rotation pivot point X coordinate, where 0 is the center, -1 the left border, +1 the right border and <-1, >1 outside.

func (e *Gltransformation) GetPivotX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pivot-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetPivotX(value float32) error {
	return e.Element.SetProperty("pivot-x", value)
}

// pivot-y (float32)
//
// Rotation pivot point X coordinate, where 0 is the center, -1 the left border, +1 the right border and <-1, >1 outside.

func (e *Gltransformation) GetPivotY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pivot-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetPivotY(value float32) error {
	return e.Element.SetProperty("pivot-y", value)
}

// pivot-z (float32)
//
// Relevant for rotation in 3D space. You look into the negative Z axis direction

func (e *Gltransformation) GetPivotZ() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pivot-z")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetPivotZ(value float32) error {
	return e.Element.SetProperty("pivot-z", value)
}

// rotation-x (float32)
//
// Rotates the video around the X-Axis in degrees.

func (e *Gltransformation) GetRotationX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rotation-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetRotationX(value float32) error {
	return e.Element.SetProperty("rotation-x", value)
}

// rotation-y (float32)
//
// Rotates the video around the Y-Axis in degrees.

func (e *Gltransformation) GetRotationY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rotation-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetRotationY(value float32) error {
	return e.Element.SetProperty("rotation-y", value)
}

// rotation-z (float32)
//
// Rotates the video around the Z-Axis in degrees.

func (e *Gltransformation) GetRotationZ() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rotation-z")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetRotationZ(value float32) error {
	return e.Element.SetProperty("rotation-z", value)
}

// scale-x (float32)
//
// Scale multiplier for the X-Axis.

func (e *Gltransformation) GetScaleX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scale-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetScaleX(value float32) error {
	return e.Element.SetProperty("scale-x", value)
}

// scale-y (float32)
//
// Scale multiplier for the Y-Axis.

func (e *Gltransformation) GetScaleY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scale-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetScaleY(value float32) error {
	return e.Element.SetProperty("scale-y", value)
}

// translation-x (float32)
//
// Translates the video at the X-Axis, in universal [0-1] coordinate.

func (e *Gltransformation) GetTranslationX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("translation-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetTranslationX(value float32) error {
	return e.Element.SetProperty("translation-x", value)
}

// translation-y (float32)
//
// Translates the video at the Y-Axis, in universal [0-1] coordinate.

func (e *Gltransformation) GetTranslationY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("translation-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetTranslationY(value float32) error {
	return e.Element.SetProperty("translation-y", value)
}

// translation-z (float32)
//
// Translates the video at the Z-Axis, in universal [0-1] coordinate.

func (e *Gltransformation) GetTranslationZ() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("translation-z")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Gltransformation) SetTranslationZ(value float32) error {
	return e.Element.SetProperty("translation-z", value)
}

