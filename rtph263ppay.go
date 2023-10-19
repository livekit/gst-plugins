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

type Rtph263Ppay struct {
	Element *gst.Element
}

func NewRtph263Ppay() (*Rtph263Ppay, error) {
	e, err := gst.NewElement("rtph263ppay")
	if err != nil {
		return nil, err
	}
	return &Rtph263Ppay{Element: e}, nil
}

func NewRtph263PpayWithName(name string) (*Rtph263Ppay, error) {
	e, err := gst.NewElementWithName("rtph263ppay", name)
	if err != nil {
		return nil, err
	}
	return &Rtph263Ppay{Element: e}, nil
}

// ----- Properties -----

// fragmentation-mode (GstFragmentationMode)
//
// Packet Fragmentation Mode

func (e *Rtph263Ppay) GetFragmentationMode() (GstFragmentationMode, error) {
	var value GstFragmentationMode
	var ok bool
	v, err := e.Element.GetProperty("fragmentation-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFragmentationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFragmentationMode")
	}
	return value, nil
}

func (e *Rtph263Ppay) SetFragmentationMode(value GstFragmentationMode) error {
	e.Element.SetArg("fragmentation-mode", string(value))
	return nil
}

// ----- Constants -----

type GstFragmentationMode string

const (
	GstFragmentationModeNormal GstFragmentationMode = "normal" // normal (0) – Normal
	GstFragmentationModeSync GstFragmentationMode = "sync" // sync (1) – Fragment at sync points
)

