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

type Glfiltercube struct {
	*base.GstBaseTransform
}

func NewGlfiltercube() (*Glfiltercube, error) {
	e, err := gst.NewElement("glfiltercube")
	if err != nil {
		return nil, err
	}
	return &Glfiltercube{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGlfiltercubeWithName(name string) (*Glfiltercube, error) {
	e, err := gst.NewElementWithName("glfiltercube", name)
	if err != nil {
		return nil, err
	}
	return &Glfiltercube{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// aspect (float64)
//
// Field of view in the x direction

func (e *Glfiltercube) GetAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glfiltercube) SetAspect(value float64) error {
	return e.Element.SetProperty("aspect", value)
}

// blue (float32)
//
// Background blue color

func (e *Glfiltercube) GetBlue() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Glfiltercube) SetBlue(value float32) error {
	return e.Element.SetProperty("blue", value)
}

// fovy (float64)
//
// Field of view angle in degrees

func (e *Glfiltercube) GetFovy() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("fovy")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glfiltercube) SetFovy(value float64) error {
	return e.Element.SetProperty("fovy", value)
}

// green (float32)
//
// Background green color

func (e *Glfiltercube) GetGreen() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("green")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Glfiltercube) SetGreen(value float32) error {
	return e.Element.SetProperty("green", value)
}

// red (float32)
//
// Background red color

func (e *Glfiltercube) GetRed() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("red")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Glfiltercube) SetRed(value float32) error {
	return e.Element.SetProperty("red", value)
}

// zfar (float64)
//
// Specifies the distance from the viewer to the far clipping plane

func (e *Glfiltercube) GetZfar() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zfar")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glfiltercube) SetZfar(value float64) error {
	return e.Element.SetProperty("zfar", value)
}

// znear (float64)
//
// Specifies the distance from the viewer to the near clipping plane

func (e *Glfiltercube) GetZnear() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("znear")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glfiltercube) SetZnear(value float64) error {
	return e.Element.SetProperty("znear", value)
}

