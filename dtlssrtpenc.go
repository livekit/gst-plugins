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

type Dtlssrtpenc struct {
	Element *gst.Element
}

func NewDtlssrtpenc() (*Dtlssrtpenc, error) {
	e, err := gst.NewElement("dtlssrtpenc")
	if err != nil {
		return nil, err
	}
	return &Dtlssrtpenc{Element: e}, nil
}

func NewDtlssrtpencWithName(name string) (*Dtlssrtpenc, error) {
	e, err := gst.NewElementWithName("dtlssrtpenc", name)
	if err != nil {
		return nil, err
	}
	return &Dtlssrtpenc{Element: e}, nil
}

// ----- Properties -----

// connection-state (GstDtlsConnectionState)
//
// Current connection state

func (e *Dtlssrtpenc) GetConnectionState() (GstDtlsConnectionState, error) {
	var value GstDtlsConnectionState
	var ok bool
	v, err := e.Element.GetProperty("connection-state")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDtlsConnectionState)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDtlsConnectionState")
	}
	return value, nil
}

// is-client (bool)
//
// Set to true if the decoder should act as client and initiate the handshake

func (e *Dtlssrtpenc) GetIsClient() (bool, error) {
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

func (e *Dtlssrtpenc) SetIsClient(value bool) error {
	return e.Element.SetProperty("is-client", value)
}

// rtp-sync (bool)
//
// Synchronize RTP to the pipeline clock before merging with RTCP

func (e *Dtlssrtpenc) GetRtpSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtp-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dtlssrtpenc) SetRtpSync(value bool) error {
	return e.Element.SetProperty("rtp-sync", value)
}

// ----- Constants -----

type GstDtlsConnectionState string

const (
	GstDtlsConnectionStateNew GstDtlsConnectionState = "new" // new (0) – New connection
	GstDtlsConnectionStateClosed GstDtlsConnectionState = "closed" // closed (1) – Closed connection on either side
	GstDtlsConnectionStateFailed GstDtlsConnectionState = "failed" // failed (2) – Failed connection
	GstDtlsConnectionStateConnecting GstDtlsConnectionState = "connecting" // connecting (3) – Connecting
	GstDtlsConnectionStateConnected GstDtlsConnectionState = "connected" // connected (4) – Successfully connected
)

