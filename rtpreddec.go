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

type Rtpreddec struct {
	Element *gst.Element
}

func NewRtpreddec() (*Rtpreddec, error) {
	e, err := gst.NewElement("rtpreddec")
	if err != nil {
		return nil, err
	}
	return &Rtpreddec{Element: e}, nil
}

func NewRtpreddecWithName(name string) (*Rtpreddec, error) {
	e, err := gst.NewElementWithName("rtpreddec", name)
	if err != nil {
		return nil, err
	}
	return &Rtpreddec{Element: e}, nil
}

// ----- Properties -----

// payloads (GstValueArray)
//
// All the RED payloads this decoder may encounter

func (e *Rtpreddec) GetPayloads() (interface{}, error) {
	return e.Element.GetProperty("payloads")
}

func (e *Rtpreddec) SetPayloads(value interface{}) error {
	return e.Element.SetProperty("payloads", value)
}

// pt (int)
//
// Payload type FEC packets

func (e *Rtpreddec) GetPt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpreddec) SetPt(value int) error {
	return e.Element.SetProperty("pt", value)
}

// received (uint)
//
// Count of received packets

func (e *Rtpreddec) GetReceived() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("received")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

