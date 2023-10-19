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

type Avdeinterlace struct {
	Element *gst.Element
}

func NewAvdeinterlace() (*Avdeinterlace, error) {
	e, err := gst.NewElement("avdeinterlace")
	if err != nil {
		return nil, err
	}
	return &Avdeinterlace{Element: e}, nil
}

func NewAvdeinterlaceWithName(name string) (*Avdeinterlace, error) {
	e, err := gst.NewElementWithName("avdeinterlace", name)
	if err != nil {
		return nil, err
	}
	return &Avdeinterlace{Element: e}, nil
}

// ----- Properties -----

// mode (GstLibAvdeinterlaceModes)
//
// This selects whether the deinterlacing methods should
// always be applied or if they should only be applied
// on content that has the "interlaced" flag on the caps.

func (e *Avdeinterlace) GetMode() (GstLibAvdeinterlaceModes, error) {
	var value GstLibAvdeinterlaceModes
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLibAvdeinterlaceModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLibAvdeinterlaceModes")
	}
	return value, nil
}

func (e *Avdeinterlace) SetMode(value GstLibAvdeinterlaceModes) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ----- Constants -----

type GstLibAvdeinterlaceModes string

const (
	GstLibAvdeinterlaceModesAuto GstLibAvdeinterlaceModes = "auto" // auto (0) – Auto detection
	GstLibAvdeinterlaceModesInterlaced GstLibAvdeinterlaceModes = "interlaced" // interlaced (1) – Force deinterlacing
	GstLibAvdeinterlaceModesDisabled GstLibAvdeinterlaceModes = "disabled" // disabled (2) – Run in passthrough mode
)

