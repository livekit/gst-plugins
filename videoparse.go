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

type Videoparse struct {
	Element *gst.Element
}

func NewVideoparse() (*Videoparse, error) {
	e, err := gst.NewElement("videoparse")
	if err != nil {
		return nil, err
	}
	return &Videoparse{Element: e}, nil
}

func NewVideoparseWithName(name string) (*Videoparse, error) {
	e, err := gst.NewElementWithName("videoparse", name)
	if err != nil {
		return nil, err
	}
	return &Videoparse{Element: e}, nil
}

// ----- Properties -----

// format (GstVideoFormat)
//
// Format of images in raw stream

func (e *Videoparse) GetFormat() (interface{}, error) {
	return e.Element.GetProperty("format")
}

func (e *Videoparse) SetFormat(value interface{}) error {
	return e.Element.SetProperty("format", value)
}

// framerate (GstFraction)
//
// Frame rate of images in raw stream

func (e *Videoparse) GetFramerate() (interface{}, error) {
	return e.Element.GetProperty("framerate")
}

func (e *Videoparse) SetFramerate(value interface{}) error {
	return e.Element.SetProperty("framerate", value)
}

// framesize (uint)
//
// Size of an image in raw stream (0: default)

func (e *Videoparse) GetFramesize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("framesize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Videoparse) SetFramesize(value uint) error {
	return e.Element.SetProperty("framesize", value)
}

// height (int)
//
// Height of images in raw stream

func (e *Videoparse) GetHeight() (int, error) {
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

func (e *Videoparse) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// interlaced (bool)
//
// True if video is interlaced

func (e *Videoparse) GetInterlaced() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("interlaced")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videoparse) SetInterlaced(value bool) error {
	return e.Element.SetProperty("interlaced", value)
}

// offsets (string)
//
// Offset of each planes in bytes using string format: 'o0,o1,o2,o3'

func (e *Videoparse) GetOffsets() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("offsets")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Videoparse) SetOffsets(value string) error {
	return e.Element.SetProperty("offsets", value)
}

// pixel-aspect-ratio (GstFraction)
//
// Pixel aspect ratio of images in raw stream

func (e *Videoparse) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

func (e *Videoparse) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// strides (string)
//
// Stride of each planes in bytes using string format: 's0,s1,s2,s3'

func (e *Videoparse) GetStrides() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("strides")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Videoparse) SetStrides(value string) error {
	return e.Element.SetProperty("strides", value)
}

// top-field-first (bool)
//
// True if top field is earlier than bottom field

func (e *Videoparse) GetTopFieldFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("top-field-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videoparse) SetTopFieldFirst(value bool) error {
	return e.Element.SetProperty("top-field-first", value)
}

// width (int)
//
// Width of images in raw stream

func (e *Videoparse) GetWidth() (int, error) {
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

func (e *Videoparse) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

