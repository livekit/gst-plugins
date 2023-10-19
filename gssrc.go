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

type Gssrc struct {
	*base.GstBaseSrc
}

func NewGssrc() (*Gssrc, error) {
	e, err := gst.NewElement("gssrc")
	if err != nil {
		return nil, err
	}
	return &Gssrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewGssrcWithName(name string) (*Gssrc, error) {
	e, err := gst.NewElementWithName("gssrc", name)
	if err != nil {
		return nil, err
	}
	return &Gssrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// location (string)
//
// Location of the file to read

func (e *Gssrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Gssrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// service-account-email (string)
//
// Service Account Email to use for credentials.

func (e *Gssrc) GetServiceAccountEmail() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("service-account-email")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Gssrc) SetServiceAccountEmail(value string) error {
	return e.Element.SetProperty("service-account-email", value)
}

