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

type Wildmididec struct {
	Element *gst.Element
}

func NewWildmididec() (*Wildmididec, error) {
	e, err := gst.NewElement("wildmididec")
	if err != nil {
		return nil, err
	}
	return &Wildmididec{Element: e}, nil
}

func NewWildmididecWithName(name string) (*Wildmididec, error) {
	e, err := gst.NewElementWithName("wildmididec", name)
	if err != nil {
		return nil, err
	}
	return &Wildmididec{Element: e}, nil
}

// ----- Properties -----

// enhanced-resampling (bool)
//
// Use enhanced resampling if set to TRUE, or linear interpolation if set to FALSE

func (e *Wildmididec) GetEnhancedResampling() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enhanced-resampling")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wildmididec) SetEnhancedResampling(value bool) error {
	return e.Element.SetProperty("enhanced-resampling", value)
}

// log-volume-scale (bool)
//
// Use a logarithmic volume scale if set to TRUE, or a linear scale if set to FALSE

func (e *Wildmididec) GetLogVolumeScale() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("log-volume-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wildmididec) SetLogVolumeScale(value bool) error {
	return e.Element.SetProperty("log-volume-scale", value)
}

// output-buffer-size (uint)
//
// Size of each output buffer, in samples (actual size can be smaller than this during flush or EOS)

func (e *Wildmididec) GetOutputBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("output-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Wildmididec) SetOutputBufferSize(value uint) error {
	return e.Element.SetProperty("output-buffer-size", value)
}

// reverb (bool)
//
// Whether or not to enable the WildMidi 8 reflection reverb engine to add more depth to the sound

func (e *Wildmididec) GetReverb() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reverb")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wildmididec) SetReverb(value bool) error {
	return e.Element.SetProperty("reverb", value)
}

