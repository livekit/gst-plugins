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

type Katetag struct {
	Element *gst.Element
}

func NewKatetag() (*Katetag, error) {
	e, err := gst.NewElement("katetag")
	if err != nil {
		return nil, err
	}
	return &Katetag{Element: e}, nil
}

func NewKatetagWithName(name string) (*Katetag, error) {
	e, err := gst.NewElementWithName("katetag", name)
	if err != nil {
		return nil, err
	}
	return &Katetag{Element: e}, nil
}

// ----- Properties -----

// category (string)
//
// Set the category of the stream

func (e *Katetag) GetCategory() (string, error) {
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

func (e *Katetag) SetCategory(value string) error {
	return e.Element.SetProperty("category", value)
}

// language (string)
//
// Set the language of the stream

func (e *Katetag) GetLanguage() (string, error) {
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

func (e *Katetag) SetLanguage(value string) error {
	return e.Element.SetProperty("language", value)
}

// original-canvas-height (int)
//
// Set the height of the canvas this stream was authored for (0 is unspecified)

func (e *Katetag) GetOriginalCanvasHeight() (int, error) {
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

func (e *Katetag) SetOriginalCanvasHeight(value int) error {
	return e.Element.SetProperty("original-canvas-height", value)
}

// original-canvas-width (int)
//
// Set the width of the canvas this stream was authored for (0 is unspecified)

func (e *Katetag) GetOriginalCanvasWidth() (int, error) {
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

func (e *Katetag) SetOriginalCanvasWidth(value int) error {
	return e.Element.SetProperty("original-canvas-width", value)
}

