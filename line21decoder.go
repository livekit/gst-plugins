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

type Line21Decoder struct {
	*base.GstBaseTransform
}

func NewLine21Decoder() (*Line21Decoder, error) {
	e, err := gst.NewElement("line21decoder")
	if err != nil {
		return nil, err
	}
	return &Line21Decoder{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLine21DecoderWithName(name string) (*Line21Decoder, error) {
	e, err := gst.NewElementWithName("line21decoder", name)
	if err != nil {
		return nil, err
	}
	return &Line21Decoder{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// mode (GstLine21DecoderMode)
//
// Control whether and how detected CC meta should be inserted
// in the list of existing CC meta on a frame (if any).

func (e *Line21Decoder) GetMode() (GstLine21DecoderMode, error) {
	var value GstLine21DecoderMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLine21DecoderMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLine21DecoderMode")
	}
	return value, nil
}

func (e *Line21Decoder) SetMode(value GstLine21DecoderMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ntsc-only (bool)
//
// Whether line 21 decoding should only be attempted when the
// input resolution matches NTSC (720 x 525) or NTSC usable
// lines (720 x 486)

func (e *Line21Decoder) GetNtscOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ntsc-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Line21Decoder) SetNtscOnly(value bool) error {
	return e.Element.SetProperty("ntsc-only", value)
}

// ----- Constants -----

type GstLine21DecoderMode string

const (
	GstLine21DecoderModeAdd GstLine21DecoderMode = "add" // add (0) – add new CC meta on top of other CC meta, if any
	GstLine21DecoderModeDrop GstLine21DecoderMode = "drop" // drop (1) – ignore CC if a CC meta was already present
	GstLine21DecoderModeReplace GstLine21DecoderMode = "replace" // replace (2) – replace existing CC meta
)

