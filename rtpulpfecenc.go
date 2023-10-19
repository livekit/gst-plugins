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

type Rtpulpfecenc struct {
	Element *gst.Element
}

func NewRtpulpfecenc() (*Rtpulpfecenc, error) {
	e, err := gst.NewElement("rtpulpfecenc")
	if err != nil {
		return nil, err
	}
	return &Rtpulpfecenc{Element: e}, nil
}

func NewRtpulpfecencWithName(name string) (*Rtpulpfecenc, error) {
	e, err := gst.NewElementWithName("rtpulpfecenc", name)
	if err != nil {
		return nil, err
	}
	return &Rtpulpfecenc{Element: e}, nil
}

// ----- Properties -----

// multipacket (bool)
//
// Apply FEC on multiple packets

func (e *Rtpulpfecenc) GetMultipacket() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("multipacket")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpulpfecenc) SetMultipacket(value bool) error {
	return e.Element.SetProperty("multipacket", value)
}

// percentage (uint)
//
// FEC overhead percentage for the whole stream

func (e *Rtpulpfecenc) GetPercentage() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("percentage")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpulpfecenc) SetPercentage(value uint) error {
	return e.Element.SetProperty("percentage", value)
}

// percentage-important (uint)
//
// FEC overhead percentage for important packets

func (e *Rtpulpfecenc) GetPercentageImportant() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("percentage-important")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpulpfecenc) SetPercentageImportant(value uint) error {
	return e.Element.SetProperty("percentage-important", value)
}

// protected (uint)
//
// Count of protected packets

func (e *Rtpulpfecenc) GetProtected() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("protected")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// pt (uint)
//
// The payload type of FEC packets

func (e *Rtpulpfecenc) GetPt() (uint, error) {
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

func (e *Rtpulpfecenc) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

