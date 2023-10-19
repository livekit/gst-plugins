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

type Rgvolume struct {
	Element *gst.Element
}

func NewRgvolume() (*Rgvolume, error) {
	e, err := gst.NewElement("rgvolume")
	if err != nil {
		return nil, err
	}
	return &Rgvolume{Element: e}, nil
}

func NewRgvolumeWithName(name string) (*Rgvolume, error) {
	e, err := gst.NewElementWithName("rgvolume", name)
	if err != nil {
		return nil, err
	}
	return &Rgvolume{Element: e}, nil
}

// ----- Properties -----

// album-mode (bool)
//
// Whether to prefer album gain over track gain.

// fallback-gain (float64)
//
// Fallback gain [dB] for streams missing ReplayGain tags.

func (e *Rgvolume) GetFallbackGain() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("fallback-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Rgvolume) SetFallbackGain(value float64) error {
	return e.Element.SetProperty("fallback-gain", value)
}

// headroom (float64)
//
// Extra headroom [dB].  This controls the amount by which the output can
// exceed digital full scale.

// pre-amp (float64)
//
// Additional gain to apply globally [dB].  This controls the trade-off
// between uniformity of normalization and utilization of available dynamic
// range.

// result-gain (float64)
//
// Applied gain [dB].  This gain is applied to processed buffer data.

// target-gain (float64)
//
// Applicable gain [dB].  This gain is supposed to be applied.

