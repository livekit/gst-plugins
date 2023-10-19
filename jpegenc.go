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

type Jpegenc struct {
	Element *gst.Element
}

func NewJpegenc() (*Jpegenc, error) {
	e, err := gst.NewElement("jpegenc")
	if err != nil {
		return nil, err
	}
	return &Jpegenc{Element: e}, nil
}

func NewJpegencWithName(name string) (*Jpegenc, error) {
	e, err := gst.NewElementWithName("jpegenc", name)
	if err != nil {
		return nil, err
	}
	return &Jpegenc{Element: e}, nil
}

// ----- Properties -----

// idct-method (GstIDCTMethod)
//
// The IDCT algorithm to use

func (e *Jpegenc) GetIdctMethod() (interface{}, error) {
	return e.Element.GetProperty("idct-method")
}

func (e *Jpegenc) SetIdctMethod(value interface{}) error {
	return e.Element.SetProperty("idct-method", value)
}

// quality (int)
//
// Quality of encoding

func (e *Jpegenc) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Jpegenc) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

// snapshot (bool)
//
// Send EOS after encoding a frame, useful for snapshots.

func (e *Jpegenc) GetSnapshot() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("snapshot")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Jpegenc) SetSnapshot(value bool) error {
	return e.Element.SetProperty("snapshot", value)
}

// ----- Constants -----

type GstIdctmethod string

const (
	GstIdctmethodIslow GstIdctmethod = "islow" // islow (0) – Slow but accurate integer algorithm
	GstIdctmethodIfast GstIdctmethod = "ifast" // ifast (1) – Faster, less accurate integer method
	GstIdctmethodFloat GstIdctmethod = "float" // float (2) – Floating-point: accurate, fast on fast HW
)

