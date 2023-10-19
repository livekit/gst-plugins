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

type Templatematch struct {
	*base.GstBaseTransform
}

func NewTemplatematch() (*Templatematch, error) {
	e, err := gst.NewElement("templatematch")
	if err != nil {
		return nil, err
	}
	return &Templatematch{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewTemplatematchWithName(name string) (*Templatematch, error) {
	e, err := gst.NewElementWithName("templatematch", name)
	if err != nil {
		return nil, err
	}
	return &Templatematch{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// display (bool)
//
// Sets whether the detected template should be highlighted in the output

func (e *Templatematch) GetDisplay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Templatematch) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

// method (int)
//
// Specifies the way the template must be compared with image regions. 0=SQDIFF, 1=SQDIFF_NORMED, 2=CCOR, 3=CCOR_NORMED, 4=CCOEFF, 5=CCOEFF_NORMED.

func (e *Templatematch) GetMethod() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Templatematch) SetMethod(value int) error {
	return e.Element.SetProperty("method", value)
}

// template (string)
//
// Filename of template image

func (e *Templatematch) GetTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Templatematch) SetTemplate(value string) error {
	return e.Element.SetProperty("template", value)
}

