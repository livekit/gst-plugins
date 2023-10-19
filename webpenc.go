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

type Webpenc struct {
	Element *gst.Element
}

func NewWebpenc() (*Webpenc, error) {
	e, err := gst.NewElement("webpenc")
	if err != nil {
		return nil, err
	}
	return &Webpenc{Element: e}, nil
}

func NewWebpencWithName(name string) (*Webpenc, error) {
	e, err := gst.NewElementWithName("webpenc", name)
	if err != nil {
		return nil, err
	}
	return &Webpenc{Element: e}, nil
}

// ----- Properties -----

// lossless (bool)
//
// Enable lossless encoding

func (e *Webpenc) GetLossless() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lossless")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webpenc) SetLossless(value bool) error {
	return e.Element.SetProperty("lossless", value)
}

// preset (GstWebpEncPreset)
//
// Preset name for visual tuning

func (e *Webpenc) GetPreset() (GstWebpEncPreset, error) {
	var value GstWebpEncPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWebpEncPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWebpEncPreset")
	}
	return value, nil
}

func (e *Webpenc) SetPreset(value GstWebpEncPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

// quality (float32)
//
// quality level, between 0 (smallest file) and 100 (biggest)

func (e *Webpenc) GetQuality() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Webpenc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

// speed (uint)
//
// quality/speed trade-off (0=fast, 6=slower-better)

func (e *Webpenc) GetSpeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Webpenc) SetSpeed(value uint) error {
	return e.Element.SetProperty("speed", value)
}

// ----- Constants -----

type GstWebpEncPreset string

const (
	GstWebpEncPresetNone GstWebpEncPreset = "none" // none (0) – Default
	GstWebpEncPresetPicture GstWebpEncPreset = "picture" // picture (1) – Digital picture,inner shot
	GstWebpEncPresetPhoto GstWebpEncPreset = "photo" // photo (2) – Outdoor photo, natural lighting
	GstWebpEncPresetDrawing GstWebpEncPreset = "drawing" // drawing (3) – Hand or Line drawing
	GstWebpEncPresetIcon GstWebpEncPreset = "icon" // icon (4) – Small-sized colorful images
	GstWebpEncPresetText GstWebpEncPreset = "text" // text (5) – text-like
)

