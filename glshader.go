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

type Glshader struct {
	*base.GstBaseTransform
}

func NewGlshader() (*Glshader, error) {
	e, err := gst.NewElement("glshader")
	if err != nil {
		return nil, err
	}
	return &Glshader{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGlshaderWithName(name string) (*Glshader, error) {
	e, err := gst.NewElementWithName("glshader", name)
	if err != nil {
		return nil, err
	}
	return &Glshader{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// fragment (string)
//
// GLSL fragment source

func (e *Glshader) GetFragment() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("fragment")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Glshader) SetFragment(value string) error {
	return e.Element.SetProperty("fragment", value)
}

// shader (GstGLShader)
//
// GstGLShader to use

func (e *Glshader) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

func (e *Glshader) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

// uniforms (GstStructure)
//
// GLSL Uniforms

func (e *Glshader) GetUniforms() (interface{}, error) {
	return e.Element.GetProperty("uniforms")
}

func (e *Glshader) SetUniforms(value interface{}) error {
	return e.Element.SetProperty("uniforms", value)
}

// update-shader (bool)
//
// Emit the 'create-shader' signal for the next frame

func (e *Glshader) GetUpdateShader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("update-shader")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Glshader) SetUpdateShader(value bool) error {
	return e.Element.SetProperty("update-shader", value)
}

// vertex (string)
//
// GLSL vertex source

func (e *Glshader) GetVertex() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("vertex")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Glshader) SetVertex(value string) error {
	return e.Element.SetProperty("vertex", value)
}

