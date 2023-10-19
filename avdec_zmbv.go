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

type AvdecZmbv struct {
	Element *gst.Element
}

func NewAvdecZmbv() (*AvdecZmbv, error) {
	e, err := gst.NewElement("avdec_zmbv")
	if err != nil {
		return nil, err
	}
	return &AvdecZmbv{Element: e}, nil
}

func NewAvdecZmbvWithName(name string) (*AvdecZmbv, error) {
	e, err := gst.NewElementWithName("avdec_zmbv", name)
	if err != nil {
		return nil, err
	}
	return &AvdecZmbv{Element: e}, nil
}

// ----- Properties -----

// debug-mv (bool)
//
// Whether to print motion vectors on top of the image (deprecated, non-functional)

func (e *AvdecZmbv) GetDebugMv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("debug-mv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvdecZmbv) SetDebugMv(value bool) error {
	return e.Element.SetProperty("debug-mv", value)
}

// direct-rendering (bool)
//
// Enable direct rendering

func (e *AvdecZmbv) GetDirectRendering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("direct-rendering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvdecZmbv) SetDirectRendering(value bool) error {
	return e.Element.SetProperty("direct-rendering", value)
}

// lowres (GstLibAVVidDecLowres)
//
// At which resolution to decode images

func (e *AvdecZmbv) GetLowres() (interface{}, error) {
	return e.Element.GetProperty("lowres")
}

func (e *AvdecZmbv) SetLowres(value interface{}) error {
	return e.Element.SetProperty("lowres", value)
}

// output-corrupt (bool)
//
// Whether libav should output frames even if corrupted

func (e *AvdecZmbv) GetOutputCorrupt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-corrupt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvdecZmbv) SetOutputCorrupt(value bool) error {
	return e.Element.SetProperty("output-corrupt", value)
}

// skip-frame (GstLibAVVidDecSkipFrame)
//
// Which types of frames to skip during decoding

func (e *AvdecZmbv) GetSkipFrame() (interface{}, error) {
	return e.Element.GetProperty("skip-frame")
}

func (e *AvdecZmbv) SetSkipFrame(value interface{}) error {
	return e.Element.SetProperty("skip-frame", value)
}

// ----- Constants -----

type GstLibAvvidDecLowres string

const (
	GstLibAvvidDecLowresFull GstLibAvvidDecLowres = "full" // full (0) – 0
	GstLibAvvidDecLowres12Size GstLibAvvidDecLowres = "1/2-size" // 1/2-size (1) – 1
	GstLibAvvidDecLowres14Size GstLibAvvidDecLowres = "1/4-size" // 1/4-size (2) – 2
)

type GstLibAvvidDecSkipFrame string

const (
	GstLibAvvidDecSkipFrameSkipNothing GstLibAvvidDecSkipFrame = "Skip nothing" // Skip nothing (0) – 0
	GstLibAvvidDecSkipFrameSkipBFrames GstLibAvvidDecSkipFrame = "Skip B-frames" // Skip B-frames (1) – 1
	GstLibAvvidDecSkipFrameSkipIdctDequantization GstLibAvvidDecSkipFrame = "Skip IDCT/Dequantization" // Skip IDCT/Dequantization (2) – 2
	GstLibAvvidDecSkipFrameSkipEverything GstLibAvvidDecSkipFrame = "Skip everything" // Skip everything (5) – 5
)

