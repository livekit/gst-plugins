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

type Capsfilter struct {
	*base.GstBaseTransform
}

func NewCapsfilter() (*Capsfilter, error) {
	e, err := gst.NewElement("capsfilter")
	if err != nil {
		return nil, err
	}
	return &Capsfilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCapsfilterWithName(name string) (*Capsfilter, error) {
	e, err := gst.NewElementWithName("capsfilter", name)
	if err != nil {
		return nil, err
	}
	return &Capsfilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// Restrict the possible allowed capabilities (NULL means ANY). Setting this property takes a reference to the supplied GstCaps object.

func (e *Capsfilter) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Capsfilter) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// caps-change-mode (GstCapsFilterCapsChangeMode)
//
// Filter caps change behaviour

func (e *Capsfilter) GetCapsChangeMode() (GstCapsFilterCapsChangeMode, error) {
	var value GstCapsFilterCapsChangeMode
	var ok bool
	v, err := e.Element.GetProperty("caps-change-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCapsFilterCapsChangeMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCapsFilterCapsChangeMode")
	}
	return value, nil
}

func (e *Capsfilter) SetCapsChangeMode(value GstCapsFilterCapsChangeMode) error {
	e.Element.SetArg("caps-change-mode", string(value))
	return nil
}

// ----- Constants -----

type GstCapsFilterCapsChangeMode string

const (
	GstCapsFilterCapsChangeModeImmediate GstCapsFilterCapsChangeMode = "immediate" // immediate (0) – Only accept the current filter caps
	GstCapsFilterCapsChangeModeDelayed GstCapsFilterCapsChangeMode = "delayed" // delayed (1) – Temporarily accept previous filter caps
)

