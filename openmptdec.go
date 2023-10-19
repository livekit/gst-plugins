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

type Openmptdec struct {
	Element *gst.Element
}

func NewOpenmptdec() (*Openmptdec, error) {
	e, err := gst.NewElement("openmptdec")
	if err != nil {
		return nil, err
	}
	return &Openmptdec{Element: e}, nil
}

func NewOpenmptdecWithName(name string) (*Openmptdec, error) {
	e, err := gst.NewElementWithName("openmptdec", name)
	if err != nil {
		return nil, err
	}
	return &Openmptdec{Element: e}, nil
}

// ----- Properties -----

// current-subsong (uint)
//
// Subsong that is currently selected for playback

func (e *Openmptdec) GetCurrentSubsong() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-subsong")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openmptdec) SetCurrentSubsong(value uint) error {
	return e.Element.SetProperty("current-subsong", value)
}

// filter-length (int)
//
// Length of interpolation filter to use for the samples (0 = internal default)

func (e *Openmptdec) GetFilterLength() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("filter-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openmptdec) SetFilterLength(value int) error {
	return e.Element.SetProperty("filter-length", value)
}

// master-gain (int)
//
// Gain to apply to the playback, in millibel

func (e *Openmptdec) GetMasterGain() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("master-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openmptdec) SetMasterGain(value int) error {
	return e.Element.SetProperty("master-gain", value)
}

// num-loops (int)
//
// Number of times a playback loop shall be executed (special values: 0 = no looping; -1 = infinite loop)

func (e *Openmptdec) GetNumLoops() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-loops")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openmptdec) SetNumLoops(value int) error {
	return e.Element.SetProperty("num-loops", value)
}

// output-buffer-size (uint)
//
// Size of each output buffer, in samples (actual size can be smaller than this during flush or EOS)

func (e *Openmptdec) GetOutputBufferSize() (uint, error) {
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

func (e *Openmptdec) SetOutputBufferSize(value uint) error {
	return e.Element.SetProperty("output-buffer-size", value)
}

// output-mode (GstNonstreamAudioOutputMode)
//
// Which mode playback shall use when a loop is encountered; looping = reset position to start of loop, steady = do not reset position

func (e *Openmptdec) GetOutputMode() (interface{}, error) {
	return e.Element.GetProperty("output-mode")
}

func (e *Openmptdec) SetOutputMode(value interface{}) error {
	return e.Element.SetProperty("output-mode", value)
}

// stereo-separation (int)
//
// Degree of separation for stereo channels, in percent

func (e *Openmptdec) GetStereoSeparation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("stereo-separation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openmptdec) SetStereoSeparation(value int) error {
	return e.Element.SetProperty("stereo-separation", value)
}

// subsong-mode (GstNonstreamAudioSubsongMode)
//
// Mode which defines how to treat subsongs

func (e *Openmptdec) GetSubsongMode() (interface{}, error) {
	return e.Element.GetProperty("subsong-mode")
}

func (e *Openmptdec) SetSubsongMode(value interface{}) error {
	return e.Element.SetProperty("subsong-mode", value)
}

// volume-ramping (int)
//
// Volume ramping strength; higher value -> slower ramping (-1 = internal default)

func (e *Openmptdec) GetVolumeRamping() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("volume-ramping")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openmptdec) SetVolumeRamping(value int) error {
	return e.Element.SetProperty("volume-ramping", value)
}

