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

type Dtlssrtpdec struct {
	Element *gst.Element
}

func NewDtlssrtpdec() (*Dtlssrtpdec, error) {
	e, err := gst.NewElement("dtlssrtpdec")
	if err != nil {
		return nil, err
	}
	return &Dtlssrtpdec{Element: e}, nil
}

func NewDtlssrtpdecWithName(name string) (*Dtlssrtpdec, error) {
	e, err := gst.NewElementWithName("dtlssrtpdec", name)
	if err != nil {
		return nil, err
	}
	return &Dtlssrtpdec{Element: e}, nil
}

// ----- Properties -----

// connection-state (GstDtlsConnectionState)
//
// Current connection state

func (e *Dtlssrtpdec) GetConnectionState() (interface{}, error) {
	return e.Element.GetProperty("connection-state")
}

// peer-pem (string)
//
// The X509 certificate received in the DTLS handshake, in PEM format

func (e *Dtlssrtpdec) GetPeerPem() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("peer-pem")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// pem (string)
//
// A string containing a X509 certificate and RSA private key in PEM format

func (e *Dtlssrtpdec) GetPem() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pem")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dtlssrtpdec) SetPem(value string) error {
	return e.Element.SetProperty("pem", value)
}

