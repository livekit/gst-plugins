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

type Assrender struct {
	Element *gst.Element
}

func NewAssrender() (*Assrender, error) {
	e, err := gst.NewElement("assrender")
	if err != nil {
		return nil, err
	}
	return &Assrender{Element: e}, nil
}

func NewAssrenderWithName(name string) (*Assrender, error) {
	e, err := gst.NewElementWithName("assrender", name)
	if err != nil {
		return nil, err
	}
	return &Assrender{Element: e}, nil
}

// ----- Properties -----

// embeddedfonts (bool)
//
// Extract and use fonts embedded in the stream

func (e *Assrender) GetEmbeddedfonts() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("embeddedfonts")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Assrender) SetEmbeddedfonts(value bool) error {
	return e.Element.SetProperty("embeddedfonts", value)
}

// enable (bool)
//
// Enable rendering of subtitles

func (e *Assrender) GetEnable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Assrender) SetEnable(value bool) error {
	return e.Element.SetProperty("enable", value)
}

// wait-text (bool)
//
// Whether to wait for subtitles

func (e *Assrender) GetWaitText() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Assrender) SetWaitText(value bool) error {
	return e.Element.SetProperty("wait-text", value)
}

