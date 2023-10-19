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

type Rawvideoparse struct {
	*base.GstBaseParse
}

func NewRawvideoparse() (*Rawvideoparse, error) {
	e, err := gst.NewElement("rawvideoparse")
	if err != nil {
		return nil, err
	}
	return &Rawvideoparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

func NewRawvideoparseWithName(name string) (*Rawvideoparse, error) {
	e, err := gst.NewElementWithName("rawvideoparse", name)
	if err != nil {
		return nil, err
	}
	return &Rawvideoparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

// ----- Properties -----

// colorimetry (string)
//
// The video source colorimetry

func (e *Rawvideoparse) GetColorimetry() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("colorimetry")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rawvideoparse) SetColorimetry(value string) error {
	return e.Element.SetProperty("colorimetry", value)
}

// format (GstVideoFormat)
//
// Format of frames in raw stream

func (e *Rawvideoparse) GetFormat() (interface{}, error) {
	return e.Element.GetProperty("format")
}

func (e *Rawvideoparse) SetFormat(value interface{}) error {
	return e.Element.SetProperty("format", value)
}

// frame-size (uint)
//
// Size of a frame (0 = frames are tightly packed together)

func (e *Rawvideoparse) GetFrameSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rawvideoparse) SetFrameSize(value uint) error {
	return e.Element.SetProperty("frame-size", value)
}

// framerate (GstFraction)
//
// Rate of frames in raw stream

func (e *Rawvideoparse) GetFramerate() (interface{}, error) {
	return e.Element.GetProperty("framerate")
}

func (e *Rawvideoparse) SetFramerate(value interface{}) error {
	return e.Element.SetProperty("framerate", value)
}

// height (int)
//
// Height of frames in raw stream

func (e *Rawvideoparse) GetHeight() (int, error) {
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

func (e *Rawvideoparse) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// interlaced (bool)
//
// True if frames in raw stream are interlaced

func (e *Rawvideoparse) GetInterlaced() (bool, error) {
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

func (e *Rawvideoparse) SetInterlaced(value bool) error {
	return e.Element.SetProperty("interlaced", value)
}

// pixel-aspect-ratio (GstFraction)
//
// Pixel aspect ratio of frames in raw stream

func (e *Rawvideoparse) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

func (e *Rawvideoparse) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// plane-offsets (GstValueArray)
//
// Offsets of the planes in bytes (e.g. plane-offsets="<0,76800>")

func (e *Rawvideoparse) GetPlaneOffsets() (interface{}, error) {
	return e.Element.GetProperty("plane-offsets")
}

func (e *Rawvideoparse) SetPlaneOffsets(value interface{}) error {
	return e.Element.SetProperty("plane-offsets", value)
}

// plane-strides (GstValueArray)
//
// Strides of the planes in bytes (e.g. plane-strides="<320,320>")

func (e *Rawvideoparse) GetPlaneStrides() (interface{}, error) {
	return e.Element.GetProperty("plane-strides")
}

func (e *Rawvideoparse) SetPlaneStrides(value interface{}) error {
	return e.Element.SetProperty("plane-strides", value)
}

// top-field-first (bool)
//
// True if top field in frames in raw stream come first (not used if frames aren't interlaced)

func (e *Rawvideoparse) GetTopFieldFirst() (bool, error) {
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

func (e *Rawvideoparse) SetTopFieldFirst(value bool) error {
	return e.Element.SetProperty("top-field-first", value)
}

// width (int)
//
// Width of frames in raw stream

func (e *Rawvideoparse) GetWidth() (int, error) {
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

func (e *Rawvideoparse) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

