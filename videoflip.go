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

type Videoflip struct {
	*base.GstBaseTransform
}

func NewVideoflip() (*Videoflip, error) {
	e, err := gst.NewElement("videoflip")
	if err != nil {
		return nil, err
	}
	return &Videoflip{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideoflipWithName(name string) (*Videoflip, error) {
	e, err := gst.NewElementWithName("videoflip", name)
	if err != nil {
		return nil, err
	}
	return &Videoflip{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// method (GstVideoFlipMethod)
//
// method (deprecated, use video-direction instead)

func (e *Videoflip) GetMethod() (GstVideoFlipMethod, error) {
	var value GstVideoFlipMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoFlipMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoFlipMethod")
	}
	return value, nil
}

func (e *Videoflip) SetMethod(value GstVideoFlipMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// ----- Constants -----

type GstVideoFlipMethod string

const (
	GstVideoFlipMethodNone GstVideoFlipMethod = "none" // none (0) – Identity (no rotation)
	GstVideoFlipMethodClockwise GstVideoFlipMethod = "clockwise" // clockwise (1) – Rotate clockwise 90 degrees
	GstVideoFlipMethodRotate180 GstVideoFlipMethod = "rotate-180" // rotate-180 (2) – Rotate 180 degrees
	GstVideoFlipMethodCounterclockwise GstVideoFlipMethod = "counterclockwise" // counterclockwise (3) – Rotate counter-clockwise 90 degrees
	GstVideoFlipMethodHorizontalFlip GstVideoFlipMethod = "horizontal-flip" // horizontal-flip (4) – Flip horizontally
	GstVideoFlipMethodVerticalFlip GstVideoFlipMethod = "vertical-flip" // vertical-flip (5) – Flip vertically
	GstVideoFlipMethodUpperLeftDiagonal GstVideoFlipMethod = "upper-left-diagonal" // upper-left-diagonal (6) – Flip across upper left/lower right diagonal
	GstVideoFlipMethodUpperRightDiagonal GstVideoFlipMethod = "upper-right-diagonal" // upper-right-diagonal (7) – Flip across upper right/lower left diagonal
	GstVideoFlipMethodAutomatic GstVideoFlipMethod = "automatic" // automatic (8) – Select flip method based on image-orientation tag
)

