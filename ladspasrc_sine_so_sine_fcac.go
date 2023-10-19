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

type LadspasrcSineSoSineFcac struct {
	*base.GstBaseSrc
}

func NewLadspasrcSineSoSineFcac() (*LadspasrcSineSoSineFcac, error) {
	e, err := gst.NewElement("ladspasrc-sine-so-sine-fcac")
	if err != nil {
		return nil, err
	}
	return &LadspasrcSineSoSineFcac{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewLadspasrcSineSoSineFcacWithName(name string) (*LadspasrcSineSoSineFcac, error) {
	e, err := gst.NewElementWithName("ladspasrc-sine-so-sine-fcac", name)
	if err != nil {
		return nil, err
	}
	return &LadspasrcSineSoSineFcac{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// amplitude (float32)
//
// Amplitude

func (e *LadspasrcSineSoSineFcac) GetAmplitude() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LadspasrcSineSoSineFcac) SetAmplitude(value float32) error {
	return e.Element.SetProperty("amplitude", value)
}

// can-activate-pull (bool)
//
// Can activate in pull mode

func (e *LadspasrcSineSoSineFcac) GetCanActivatePull() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-pull")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LadspasrcSineSoSineFcac) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

// can-activate-push (bool)
//
// Can activate in push mode

func (e *LadspasrcSineSoSineFcac) GetCanActivatePush() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-push")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LadspasrcSineSoSineFcac) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

// frequency (float32)
//
// Frequency (Hz)

func (e *LadspasrcSineSoSineFcac) GetFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LadspasrcSineSoSineFcac) SetFrequency(value float32) error {
	return e.Element.SetProperty("frequency", value)
}

// is-live (bool)
//
// Whether to act as a live source

func (e *LadspasrcSineSoSineFcac) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LadspasrcSineSoSineFcac) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// samplesperbuffer (int)
//
// Number of samples in each outgoing buffer

func (e *LadspasrcSineSoSineFcac) GetSamplesperbuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samplesperbuffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LadspasrcSineSoSineFcac) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

// timestamp-offset (int64)
//
// An offset added to timestamps set on buffers (in ns)

func (e *LadspasrcSineSoSineFcac) GetTimestampOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *LadspasrcSineSoSineFcac) SetTimestampOffset(value int64) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

