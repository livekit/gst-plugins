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

type Sctpdec struct {
	Element *gst.Element
}

func NewSctpdec() (*Sctpdec, error) {
	e, err := gst.NewElement("sctpdec")
	if err != nil {
		return nil, err
	}
	return &Sctpdec{Element: e}, nil
}

func NewSctpdecWithName(name string) (*Sctpdec, error) {
	e, err := gst.NewElementWithName("sctpdec", name)
	if err != nil {
		return nil, err
	}
	return &Sctpdec{Element: e}, nil
}

// ----- Properties -----

// local-sctp-port (uint)
//
// Local sctp port for the sctp association. The remote port is configured via the GstSctpEnc element.

func (e *Sctpdec) GetLocalSctpPort() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("local-sctp-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Sctpdec) SetLocalSctpPort(value uint) error {
	return e.Element.SetProperty("local-sctp-port", value)
}

// sctp-association-id (uint)
//
// Every encoder/decoder pair should have the same, unique, sctp-association-id. This value must be set before any pads are requested.

func (e *Sctpdec) GetSctpAssociationId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sctp-association-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Sctpdec) SetSctpAssociationId(value uint) error {
	return e.Element.SetProperty("sctp-association-id", value)
}

