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

type Audiopanorama struct {
	*base.GstBaseTransform
}

func NewAudiopanorama() (*Audiopanorama, error) {
	e, err := gst.NewElement("audiopanorama")
	if err != nil {
		return nil, err
	}
	return &Audiopanorama{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiopanoramaWithName(name string) (*Audiopanorama, error) {
	e, err := gst.NewElementWithName("audiopanorama", name)
	if err != nil {
		return nil, err
	}
	return &Audiopanorama{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// method (GstAudioPanoramaMethod)
//
// Panning method: psychoacoustic mode keeps the same perceived loudness,
// while simple mode just controls the volume of one channel. It's merely
// a matter of taste which method should be chosen.

func (e *Audiopanorama) GetMethod() (GstAudioPanoramaMethod, error) {
	var value GstAudioPanoramaMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioPanoramaMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioPanoramaMethod")
	}
	return value, nil
}

func (e *Audiopanorama) SetMethod(value GstAudioPanoramaMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// panorama (float32)
//
// Position in stereo panorama (-1.0 left -> 1.0 right)

func (e *Audiopanorama) GetPanorama() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panorama")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiopanorama) SetPanorama(value float32) error {
	return e.Element.SetProperty("panorama", value)
}

// ----- Constants -----

type GstAudioPanoramaMethod string

const (
	GstAudioPanoramaMethodPsychoacoustic GstAudioPanoramaMethod = "psychoacoustic" // psychoacoustic (0) – Psychoacoustic Panning (default)
	GstAudioPanoramaMethodSimple GstAudioPanoramaMethod = "simple" // simple (1) – Simple Panning
)

