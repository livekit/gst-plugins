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

type Amrnbenc struct {
	Element *gst.Element
}

func NewAmrnbenc() (*Amrnbenc, error) {
	e, err := gst.NewElement("amrnbenc")
	if err != nil {
		return nil, err
	}
	return &Amrnbenc{Element: e}, nil
}

func NewAmrnbencWithName(name string) (*Amrnbenc, error) {
	e, err := gst.NewElementWithName("amrnbenc", name)
	if err != nil {
		return nil, err
	}
	return &Amrnbenc{Element: e}, nil
}

// ----- Properties -----

// band-mode (GstAmrnbEncBandMode)
//
// Encoding Band Mode (Kbps)

func (e *Amrnbenc) GetBandMode() (GstAmrnbEncBandMode, error) {
	var value GstAmrnbEncBandMode
	var ok bool
	v, err := e.Element.GetProperty("band-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmrnbEncBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmrnbEncBandMode")
	}
	return value, nil
}

func (e *Amrnbenc) SetBandMode(value GstAmrnbEncBandMode) error {
	e.Element.SetArg("band-mode", string(value))
	return nil
}

// ----- Constants -----

type GstAmrnbEncBandMode string

const (
	GstAmrnbEncBandModeMr475 GstAmrnbEncBandMode = "MR475" // MR475 (0) – MR475
	GstAmrnbEncBandModeMr515 GstAmrnbEncBandMode = "MR515" // MR515 (1) – MR515
	GstAmrnbEncBandModeMr59 GstAmrnbEncBandMode = "MR59" // MR59 (2) – MR59
	GstAmrnbEncBandModeMr67 GstAmrnbEncBandMode = "MR67" // MR67 (3) – MR67
	GstAmrnbEncBandModeMr74 GstAmrnbEncBandMode = "MR74" // MR74 (4) – MR74
	GstAmrnbEncBandModeMr795 GstAmrnbEncBandMode = "MR795" // MR795 (5) – MR795
	GstAmrnbEncBandModeMr102 GstAmrnbEncBandMode = "MR102" // MR102 (6) – MR102
	GstAmrnbEncBandModeMr122 GstAmrnbEncBandMode = "MR122" // MR122 (7) – MR122
	GstAmrnbEncBandModeMrdtx GstAmrnbEncBandMode = "MRDTX" // MRDTX (8) – MRDTX
)

