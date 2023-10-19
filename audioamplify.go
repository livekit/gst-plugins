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

type Audioamplify struct {
	*base.GstBaseTransform
}

func NewAudioamplify() (*Audioamplify, error) {
	e, err := gst.NewElement("audioamplify")
	if err != nil {
		return nil, err
	}
	return &Audioamplify{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioamplifyWithName(name string) (*Audioamplify, error) {
	e, err := gst.NewElementWithName("audioamplify", name)
	if err != nil {
		return nil, err
	}
	return &Audioamplify{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// amplification (float32)
//
// Factor of amplification

func (e *Audioamplify) GetAmplification() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("amplification")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audioamplify) SetAmplification(value float32) error {
	return e.Element.SetProperty("amplification", value)
}

// clipping-method (GstAudioAmplifyClippingMethod)
//
// Clipping method: clip mode set values higher than the maximum to the
// maximum. The wrap-negative mode pushes those values back from the
// opposite side, wrap-positive pushes them back from the same side.

func (e *Audioamplify) GetClippingMethod() (GstAudioAmplifyClippingMethod, error) {
	var value GstAudioAmplifyClippingMethod
	var ok bool
	v, err := e.Element.GetProperty("clipping-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioAmplifyClippingMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioAmplifyClippingMethod")
	}
	return value, nil
}

func (e *Audioamplify) SetClippingMethod(value GstAudioAmplifyClippingMethod) error {
	e.Element.SetArg("clipping-method", string(value))
	return nil
}

// ----- Constants -----

type GstAudioAmplifyClippingMethod string

const (
	GstAudioAmplifyClippingMethodClip GstAudioAmplifyClippingMethod = "clip" // clip (0) – Normal clipping (default)
	GstAudioAmplifyClippingMethodWrapNegative GstAudioAmplifyClippingMethod = "wrap-negative" // wrap-negative (1) – Push overdriven values back from the opposite side
	GstAudioAmplifyClippingMethodWrapPositive GstAudioAmplifyClippingMethod = "wrap-positive" // wrap-positive (2) – Push overdriven values back from the same side
	GstAudioAmplifyClippingMethodNone GstAudioAmplifyClippingMethod = "none" // none (3) – No clipping
)

