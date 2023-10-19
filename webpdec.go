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

type Webpdec struct {
	Element *gst.Element
}

func NewWebpdec() (*Webpdec, error) {
	e, err := gst.NewElement("webpdec")
	if err != nil {
		return nil, err
	}
	return &Webpdec{Element: e}, nil
}

func NewWebpdecWithName(name string) (*Webpdec, error) {
	e, err := gst.NewElementWithName("webpdec", name)
	if err != nil {
		return nil, err
	}
	return &Webpdec{Element: e}, nil
}

// ----- Properties -----

// bypass-filtering (bool)
//
// When enabled, skip the in-loop filtering

func (e *Webpdec) GetBypassFiltering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bypass-filtering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webpdec) SetBypassFiltering(value bool) error {
	return e.Element.SetProperty("bypass-filtering", value)
}

// no-fancy-upsampling (bool)
//
// When enabled, use faster pointwise upsampler

func (e *Webpdec) GetNoFancyUpsampling() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("no-fancy-upsampling")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webpdec) SetNoFancyUpsampling(value bool) error {
	return e.Element.SetProperty("no-fancy-upsampling", value)
}

// use-threads (bool)
//
// When enabled, use multi-threaded decoding

func (e *Webpdec) GetUseThreads() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webpdec) SetUseThreads(value bool) error {
	return e.Element.SetProperty("use-threads", value)
}

