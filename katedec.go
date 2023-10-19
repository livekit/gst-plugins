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

type Katedec struct {
	Element *gst.Element
}

func NewKatedec() (*Katedec, error) {
	e, err := gst.NewElement("katedec")
	if err != nil {
		return nil, err
	}
	return &Katedec{Element: e}, nil
}

func NewKatedecWithName(name string) (*Katedec, error) {
	e, err := gst.NewElementWithName("katedec", name)
	if err != nil {
		return nil, err
	}
	return &Katedec{Element: e}, nil
}

// ----- Properties -----

// category (string)
//
// The category of the stream

func (e *Katedec) GetCategory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("category")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// language (string)
//
// The language of the stream

func (e *Katedec) GetLanguage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("language")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// original-canvas-height (int)
//
// The canvas height this stream was authored for (0 is unspecified)

func (e *Katedec) GetOriginalCanvasHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// original-canvas-width (int)
//
// The canvas width this stream was authored for

func (e *Katedec) GetOriginalCanvasWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// remove-markup (bool)
//
// Remove markup from decoded text ?

func (e *Katedec) GetRemoveMarkup() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-markup")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Katedec) SetRemoveMarkup(value bool) error {
	return e.Element.SetProperty("remove-markup", value)
}

