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

type Debugspy struct {
	*base.GstBaseTransform
}

func NewDebugspy() (*Debugspy, error) {
	e, err := gst.NewElement("debugspy")
	if err != nil {
		return nil, err
	}
	return &Debugspy{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDebugspyWithName(name string) (*Debugspy, error) {
	e, err := gst.NewElementWithName("debugspy", name)
	if err != nil {
		return nil, err
	}
	return &Debugspy{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// checksum-type (GstGchecksumType)
//
// Checksum algorithm to use

func (e *Debugspy) GetChecksumType() (interface{}, error) {
	return e.Element.GetProperty("checksum-type")
}

func (e *Debugspy) SetChecksumType(value interface{}) error {
	return e.Element.SetProperty("checksum-type", value)
}

// silent (bool)
//
// Produce verbose output ?

func (e *Debugspy) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Debugspy) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

