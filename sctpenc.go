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

type Sctpenc struct {
	Element *gst.Element
}

func NewSctpenc() (*Sctpenc, error) {
	e, err := gst.NewElement("sctpenc")
	if err != nil {
		return nil, err
	}
	return &Sctpenc{Element: e}, nil
}

func NewSctpencWithName(name string) (*Sctpenc, error) {
	e, err := gst.NewElementWithName("sctpenc", name)
	if err != nil {
		return nil, err
	}
	return &Sctpenc{Element: e}, nil
}

// ----- Properties -----

// remote-sctp-port (uint)
//
// Sctp remote sctp port for the sctp association. The local port is configured via the GstSctpDec element.

func (e *Sctpenc) GetRemoteSctpPort() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("remote-sctp-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Sctpenc) SetRemoteSctpPort(value uint) error {
	return e.Element.SetProperty("remote-sctp-port", value)
}

// sctp-association-id (uint)
//
// Every encoder/decoder pair should have the same, unique, sctp-association-id. This value must be set before any pads are requested.

func (e *Sctpenc) GetSctpAssociationId() (uint, error) {
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

func (e *Sctpenc) SetSctpAssociationId(value uint) error {
	return e.Element.SetProperty("sctp-association-id", value)
}

// use-sock-stream (bool)
//
// When set to TRUE, a sequenced, reliable, connection-based connection is used.When TRUE the partial reliability parameters of the channel are ignored.

func (e *Sctpenc) GetUseSockStream() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-sock-stream")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Sctpenc) SetUseSockStream(value bool) error {
	return e.Element.SetProperty("use-sock-stream", value)
}

