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
	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Multisocketsink struct {
	*base.GstBaseSink
}

func NewMultisocketsink() (*Multisocketsink, error) {
	e, err := gst.NewElement("multisocketsink")
	if err != nil {
		return nil, err
	}
	return &Multisocketsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewMultisocketsinkWithName(name string) (*Multisocketsink, error) {
	e, err := gst.NewElementWithName("multisocketsink", name)
	if err != nil {
		return nil, err
	}
	return &Multisocketsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// send-dispatched (bool)
//
// Sends a GstNetworkMessageDispatched event upstream whenever a buffer
// is sent to a client.
// The event is a CUSTOM event name GstNetworkMessageDispatched and
// contains:

// send-messages (bool)
//
// Sends a GstNetworkMessage event upstream whenever a buffer
// is received from a client.
// The event is a CUSTOM event name GstNetworkMessage and contains:

// ----- Constants -----

type GstMultiHandleSinkClientStatus string

const (
	GstMultiHandleSinkClientStatusOk GstMultiHandleSinkClientStatus = "ok" // ok (0) – ok
	GstMultiHandleSinkClientStatusClosed GstMultiHandleSinkClientStatus = "closed" // closed (1) – Closed
	GstMultiHandleSinkClientStatusRemoved GstMultiHandleSinkClientStatus = "removed" // removed (2) – Removed
	GstMultiHandleSinkClientStatusSlow GstMultiHandleSinkClientStatus = "slow" // slow (3) – Too slow
	GstMultiHandleSinkClientStatusError GstMultiHandleSinkClientStatus = "error" // error (4) – Error
	GstMultiHandleSinkClientStatusDuplicate GstMultiHandleSinkClientStatus = "duplicate" // duplicate (5) – Duplicate
	GstMultiHandleSinkClientStatusFlushing GstMultiHandleSinkClientStatus = "flushing" // flushing (6) – Flushing
)

