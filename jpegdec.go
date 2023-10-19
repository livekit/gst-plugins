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

type Jpegdec struct {
	Element *gst.Element
}

func NewJpegdec() (*Jpegdec, error) {
	e, err := gst.NewElement("jpegdec")
	if err != nil {
		return nil, err
	}
	return &Jpegdec{Element: e}, nil
}

func NewJpegdecWithName(name string) (*Jpegdec, error) {
	e, err := gst.NewElementWithName("jpegdec", name)
	if err != nil {
		return nil, err
	}
	return &Jpegdec{Element: e}, nil
}

// ----- Properties -----

// idct-method (GstIDCTMethod)
//
// The IDCT algorithm to use

func (e *Jpegdec) GetIdctMethod() (interface{}, error) {
	return e.Element.GetProperty("idct-method")
}

func (e *Jpegdec) SetIdctMethod(value interface{}) error {
	return e.Element.SetProperty("idct-method", value)
}

// max-errors (int)
//
// Error out after receiving N consecutive decoding errors
// (-1 = never error out, 0 = automatic, 1 = fail on first error, etc.)

func (e *Jpegdec) GetMaxErrors() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-errors")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Jpegdec) SetMaxErrors(value int) error {
	return e.Element.SetProperty("max-errors", value)
}

