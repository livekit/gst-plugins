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

type Lcms struct {
	*base.GstBaseTransform
}

func NewLcms() (*Lcms, error) {
	e, err := gst.NewElement("lcms")
	if err != nil {
		return nil, err
	}
	return &Lcms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLcmsWithName(name string) (*Lcms, error) {
	e, err := gst.NewElementWithName("lcms", name)
	if err != nil {
		return nil, err
	}
	return &Lcms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// dest-profile (string)
//
// Specify the destination ICC profile file to apply

func (e *Lcms) GetDestProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dest-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Lcms) SetDestProfile(value string) error {
	return e.Element.SetProperty("dest-profile", value)
}

// embedded-profile (bool)
//
// Extract and use source profiles embedded in images

func (e *Lcms) GetEmbeddedProfile() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("embedded-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Lcms) SetEmbeddedProfile(value bool) error {
	return e.Element.SetProperty("embedded-profile", value)
}

// input-profile (string)
//
// Specify the input ICC profile file to apply

func (e *Lcms) GetInputProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("input-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Lcms) SetInputProfile(value string) error {
	return e.Element.SetProperty("input-profile", value)
}

// intent (GstLcmsIntent)
//
// Select the rendering intent of the color correction

func (e *Lcms) GetIntent() (GstLcmsIntent, error) {
	var value GstLcmsIntent
	var ok bool
	v, err := e.Element.GetProperty("intent")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLcmsIntent)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLcmsIntent")
	}
	return value, nil
}

func (e *Lcms) SetIntent(value GstLcmsIntent) error {
	e.Element.SetArg("intent", string(value))
	return nil
}

// lookup (GstLcmsLookupMethod)
//
// Select the caching method for the color compensation calculations

func (e *Lcms) GetLookup() (GstLcmsLookupMethod, error) {
	var value GstLcmsLookupMethod
	var ok bool
	v, err := e.Element.GetProperty("lookup")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLcmsLookupMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLcmsLookupMethod")
	}
	return value, nil
}

func (e *Lcms) SetLookup(value GstLcmsLookupMethod) error {
	e.Element.SetArg("lookup", string(value))
	return nil
}

// preserve-black (bool)
//
// Select whether purely black pixels should be preserved

func (e *Lcms) GetPreserveBlack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("preserve-black")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Lcms) SetPreserveBlack(value bool) error {
	return e.Element.SetProperty("preserve-black", value)
}

// ----- Constants -----

type GstLcmsIntent string

const (
	GstLcmsIntentPerceptual GstLcmsIntent = "perceptual" // perceptual (0) – Perceptual
	GstLcmsIntentRelative GstLcmsIntent = "relative" // relative (1) – Relative Colorimetric
	GstLcmsIntentSaturation GstLcmsIntent = "saturation" // saturation (2) – Saturation
	GstLcmsIntentAbsolute GstLcmsIntent = "absolute" // absolute (3) – Absolute Colorimetric
)

type GstLcmsLookupMethod string

const (
	GstLcmsLookupMethodUncached GstLcmsLookupMethod = "uncached" // uncached (0) – Uncached, calculate every pixel on the fly (very slow playback)
	GstLcmsLookupMethodPrecalculated GstLcmsLookupMethod = "precalculated" // precalculated (1) – Precalculate lookup table (takes a long time getting READY)
	GstLcmsLookupMethodCached GstLcmsLookupMethod = "cached" // cached (2) – Calculate and cache color replacement values on first occurrence
)

