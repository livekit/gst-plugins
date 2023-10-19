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

type Rtpulpfecdec struct {
	Element *gst.Element
}

func NewRtpulpfecdec() (*Rtpulpfecdec, error) {
	e, err := gst.NewElement("rtpulpfecdec")
	if err != nil {
		return nil, err
	}
	return &Rtpulpfecdec{Element: e}, nil
}

func NewRtpulpfecdecWithName(name string) (*Rtpulpfecdec, error) {
	e, err := gst.NewElementWithName("rtpulpfecdec", name)
	if err != nil {
		return nil, err
	}
	return &Rtpulpfecdec{Element: e}, nil
}

// ----- Properties -----

// passthrough (bool)
//
// Whether to push data through without any modification.  If passthrough is
// enabled, then no packets will ever be recovered.

func (e *Rtpulpfecdec) GetPassthrough() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("passthrough")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpulpfecdec) SetPassthrough(value bool) error {
	return e.Element.SetProperty("passthrough", value)
}

// pt (uint)
//
// FEC packets payload type

func (e *Rtpulpfecdec) GetPt() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpulpfecdec) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

// recovered (uint)
//
// The number of recovered packets

func (e *Rtpulpfecdec) GetRecovered() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("recovered")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// storage (GstGobject)
//
// RTP storage

func (e *Rtpulpfecdec) GetStorage() (interface{}, error) {
	return e.Element.GetProperty("storage")
}

func (e *Rtpulpfecdec) SetStorage(value interface{}) error {
	return e.Element.SetProperty("storage", value)
}

// unrecovered (uint)
//
// The number of unrecovered packets

func (e *Rtpulpfecdec) GetUnrecovered() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("unrecovered")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

