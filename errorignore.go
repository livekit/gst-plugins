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

type Errorignore struct {
	Element *gst.Element
}

func NewErrorignore() (*Errorignore, error) {
	e, err := gst.NewElement("errorignore")
	if err != nil {
		return nil, err
	}
	return &Errorignore{Element: e}, nil
}

func NewErrorignoreWithName(name string) (*Errorignore, error) {
	e, err := gst.NewElementWithName("errorignore", name)
	if err != nil {
		return nil, err
	}
	return &Errorignore{Element: e}, nil
}

// ----- Properties -----

// convert-to (GstFlowReturn)
//
// Which GstFlowReturn value we should convert to when ignoring

func (e *Errorignore) GetConvertTo() (interface{}, error) {
	return e.Element.GetProperty("convert-to")
}

func (e *Errorignore) SetConvertTo(value interface{}) error {
	return e.Element.SetProperty("convert-to", value)
}

// ignore-eos (bool)

func (e *Errorignore) GetIgnoreEos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-eos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Errorignore) SetIgnoreEos(value bool) error {
	return e.Element.SetProperty("ignore-eos", value)
}

// ignore-error (bool)
//
// Whether to ignore GST_FLOW_ERROR

func (e *Errorignore) GetIgnoreError() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Errorignore) SetIgnoreError(value bool) error {
	return e.Element.SetProperty("ignore-error", value)
}

// ignore-notlinked (bool)
//
// Whether to ignore GST_FLOW_NOT_LINKED

func (e *Errorignore) GetIgnoreNotlinked() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-notlinked")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Errorignore) SetIgnoreNotlinked(value bool) error {
	return e.Element.SetProperty("ignore-notlinked", value)
}

// ignore-notnegotiated (bool)
//
// Whether to ignore GST_FLOW_NOT_NEGOTIATED

func (e *Errorignore) GetIgnoreNotnegotiated() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-notnegotiated")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Errorignore) SetIgnoreNotnegotiated(value bool) error {
	return e.Element.SetProperty("ignore-notnegotiated", value)
}

