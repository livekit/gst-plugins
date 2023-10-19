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

type Glvideoflip struct {
	Element *gst.Element
}

func NewGlvideoflip() (*Glvideoflip, error) {
	e, err := gst.NewElement("glvideoflip")
	if err != nil {
		return nil, err
	}
	return &Glvideoflip{Element: e}, nil
}

func NewGlvideoflipWithName(name string) (*Glvideoflip, error) {
	e, err := gst.NewElementWithName("glvideoflip", name)
	if err != nil {
		return nil, err
	}
	return &Glvideoflip{Element: e}, nil
}

// ----- Properties -----

// method (GstGlvideoFlipMethod)
//
// method (deprecated, use video-direction instead)

func (e *Glvideoflip) GetMethod() (GstGlvideoFlipMethod, error) {
	var value GstGlvideoFlipMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGlvideoFlipMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGlvideoFlipMethod")
	}
	return value, nil
}

func (e *Glvideoflip) SetMethod(value GstGlvideoFlipMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// ----- Constants -----

type GstGlvideoFlipMethod string

const (
	GstGlvideoFlipMethodNone GstGlvideoFlipMethod = "none" // none (0) – Identity (no rotation)
	GstGlvideoFlipMethodClockwise GstGlvideoFlipMethod = "clockwise" // clockwise (1) – Rotate clockwise 90 degrees
	GstGlvideoFlipMethodRotate180 GstGlvideoFlipMethod = "rotate-180" // rotate-180 (2) – Rotate 180 degrees
	GstGlvideoFlipMethodCounterclockwise GstGlvideoFlipMethod = "counterclockwise" // counterclockwise (3) – Rotate counter-clockwise 90 degrees
	GstGlvideoFlipMethodHorizontalFlip GstGlvideoFlipMethod = "horizontal-flip" // horizontal-flip (4) – Flip horizontally
	GstGlvideoFlipMethodVerticalFlip GstGlvideoFlipMethod = "vertical-flip" // vertical-flip (5) – Flip vertically
	GstGlvideoFlipMethodUpperLeftDiagonal GstGlvideoFlipMethod = "upper-left-diagonal" // upper-left-diagonal (6) – Flip across upper left/lower right diagonal
	GstGlvideoFlipMethodUpperRightDiagonal GstGlvideoFlipMethod = "upper-right-diagonal" // upper-right-diagonal (7) – Flip across upper right/lower left diagonal
	GstGlvideoFlipMethodAutomatic GstGlvideoFlipMethod = "automatic" // automatic (8) – Select flip method based on image-orientation tag
)

