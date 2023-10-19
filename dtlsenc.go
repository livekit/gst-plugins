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

type Dtlsenc struct {
	Element *gst.Element
}

func NewDtlsenc() (*Dtlsenc, error) {
	e, err := gst.NewElement("dtlsenc")
	if err != nil {
		return nil, err
	}
	return &Dtlsenc{Element: e}, nil
}

func NewDtlsencWithName(name string) (*Dtlsenc, error) {
	e, err := gst.NewElementWithName("dtlsenc", name)
	if err != nil {
		return nil, err
	}
	return &Dtlsenc{Element: e}, nil
}

// ----- Properties -----

// connection-id (string)
//
// Every encoder/decoder pair should have the same, unique, connection-id

func (e *Dtlsenc) GetConnectionId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("connection-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dtlsenc) SetConnectionId(value string) error {
	return e.Element.SetProperty("connection-id", value)
}

// connection-state (GstDtlsConnectionState)
//
// Current connection state

func (e *Dtlsenc) GetConnectionState() (interface{}, error) {
	return e.Element.GetProperty("connection-state")
}

// encoder-key (GstBuffer)
//
// Master key that should be used by the SRTP encoder

func (e *Dtlsenc) GetEncoderKey() (interface{}, error) {
	return e.Element.GetProperty("encoder-key")
}

// is-client (bool)
//
// Set to true if the decoder should act as client and initiate the handshake

func (e *Dtlsenc) GetIsClient() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-client")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dtlsenc) SetIsClient(value bool) error {
	return e.Element.SetProperty("is-client", value)
}

// srtp-auth (uint)
//
// The SRTP authentication selected in the DTLS handshake. The value will be set to an GstDtlsSrtpAuth.

func (e *Dtlsenc) GetSrtpAuth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("srtp-auth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// srtp-cipher (uint)
//
// The SRTP cipher selected in the DTLS handshake. The value will be set to an GstDtlsSrtpCipher.

func (e *Dtlsenc) GetSrtpCipher() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("srtp-cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

