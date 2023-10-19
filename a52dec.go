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

type A52Dec struct {
	Element *gst.Element
}

func NewA52Dec() (*A52Dec, error) {
	e, err := gst.NewElement("a52dec")
	if err != nil {
		return nil, err
	}
	return &A52Dec{Element: e}, nil
}

func NewA52DecWithName(name string) (*A52Dec, error) {
	e, err := gst.NewElementWithName("a52dec", name)
	if err != nil {
		return nil, err
	}
	return &A52Dec{Element: e}, nil
}

// ----- Properties -----

// drc (bool)
//
// Use Dynamic Range Compression

func (e *A52Dec) GetDrc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *A52Dec) SetDrc(value bool) error {
	return e.Element.SetProperty("drc", value)
}

// lfe (bool)
//
// LFE

func (e *A52Dec) GetLfe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *A52Dec) SetLfe(value bool) error {
	return e.Element.SetProperty("lfe", value)
}

// mode (GstA52DecMode)
//
// Decoding Mode (default 3f2r)

func (e *A52Dec) GetMode() (GstA52DecMode, error) {
	var value GstA52DecMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstA52DecMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstA52DecMode")
	}
	return value, nil
}

func (e *A52Dec) SetMode(value GstA52DecMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ----- Constants -----

type GstA52DecMode string

const (
	GstA52DecModeMono GstA52DecMode = "mono" // mono (1) – Mono
	GstA52DecModeStereo GstA52DecMode = "stereo" // stereo (2) – Stereo
	GstA52DecMode3F GstA52DecMode = "3f" // 3f (3) – 3 Front
	GstA52DecMode2F1R GstA52DecMode = "2f1r" // 2f1r (4) – 2 Front, 1 Rear
	GstA52DecMode3F1R GstA52DecMode = "3f1r" // 3f1r (5) – 3 Front, 1 Rear
	GstA52DecMode2F2R GstA52DecMode = "2f2r" // 2f2r (6) – 2 Front, 2 Rear
	GstA52DecMode3F2R GstA52DecMode = "3f2r" // 3f2r (7) – 3 Front, 2 Rear
	GstA52DecModeDolby GstA52DecMode = "dolby" // dolby (10) – Dolby
)

