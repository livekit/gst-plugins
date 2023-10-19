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

type Bs2B struct {
	*base.GstBaseTransform
}

func NewBs2B() (*Bs2B, error) {
	e, err := gst.NewElement("bs2b")
	if err != nil {
		return nil, err
	}
	return &Bs2B{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewBs2BWithName(name string) (*Bs2B, error) {
	e, err := gst.NewElementWithName("bs2b", name)
	if err != nil {
		return nil, err
	}
	return &Bs2B{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// fcut (int)
//
// Low-pass filter cut frequency (Hz)

func (e *Bs2B) GetFcut() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fcut")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Bs2B) SetFcut(value int) error {
	return e.Element.SetProperty("fcut", value)
}

// feed (int)
//
// Feed Level (dB/10)

func (e *Bs2B) GetFeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("feed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Bs2B) SetFeed(value int) error {
	return e.Element.SetProperty("feed", value)
}

