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

type Dvbsuboverlay struct {
	Element *gst.Element
}

func NewDvbsuboverlay() (*Dvbsuboverlay, error) {
	e, err := gst.NewElement("dvbsuboverlay")
	if err != nil {
		return nil, err
	}
	return &Dvbsuboverlay{Element: e}, nil
}

func NewDvbsuboverlayWithName(name string) (*Dvbsuboverlay, error) {
	e, err := gst.NewElementWithName("dvbsuboverlay", name)
	if err != nil {
		return nil, err
	}
	return &Dvbsuboverlay{Element: e}, nil
}

// ----- Properties -----

// enable (bool)
//
// Enable rendering of subtitles

func (e *Dvbsuboverlay) GetEnable() (bool, error) {
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

func (e *Dvbsuboverlay) SetEnable(value bool) error {
	return e.Element.SetProperty("enable", value)
}

// force-end (bool)
//
// Assume PES-aligned subtitles and force end-of-display

func (e *Dvbsuboverlay) GetForceEnd() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-end")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dvbsuboverlay) SetForceEnd(value bool) error {
	return e.Element.SetProperty("force-end", value)
}

// max-page-timeout (int)
//
// Limit maximum display time of a subtitle page (0 - disabled, value in seconds)

func (e *Dvbsuboverlay) GetMaxPageTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-page-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dvbsuboverlay) SetMaxPageTimeout(value int) error {
	return e.Element.SetProperty("max-page-timeout", value)
}

