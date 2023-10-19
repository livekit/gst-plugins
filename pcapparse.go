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

type Pcapparse struct {
	Element *gst.Element
}

func NewPcapparse() (*Pcapparse, error) {
	e, err := gst.NewElement("pcapparse")
	if err != nil {
		return nil, err
	}
	return &Pcapparse{Element: e}, nil
}

func NewPcapparseWithName(name string) (*Pcapparse, error) {
	e, err := gst.NewElementWithName("pcapparse", name)
	if err != nil {
		return nil, err
	}
	return &Pcapparse{Element: e}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// The caps of the source pad

func (e *Pcapparse) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Pcapparse) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// dst-ip (string)
//
// Destination IP to restrict to

func (e *Pcapparse) GetDstIp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dst-ip")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Pcapparse) SetDstIp(value string) error {
	return e.Element.SetProperty("dst-ip", value)
}

// dst-port (int)
//
// Destination port to restrict to

func (e *Pcapparse) GetDstPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("dst-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Pcapparse) SetDstPort(value int) error {
	return e.Element.SetProperty("dst-port", value)
}

// src-ip (string)
//
// Source IP to restrict to

func (e *Pcapparse) GetSrcIp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("src-ip")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Pcapparse) SetSrcIp(value string) error {
	return e.Element.SetProperty("src-ip", value)
}

// src-port (int)
//
// Source port to restrict to

func (e *Pcapparse) GetSrcPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("src-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Pcapparse) SetSrcPort(value int) error {
	return e.Element.SetProperty("src-port", value)
}

// ts-offset (int64)
//
// Relative timestamp offset (ns) to apply (-1 = use absolute packet time)

func (e *Pcapparse) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Pcapparse) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

