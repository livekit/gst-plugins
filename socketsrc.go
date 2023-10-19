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

type Socketsrc struct {
	*base.GstPushSrc
}

func NewSocketsrc() (*Socketsrc, error) {
	e, err := gst.NewElement("socketsrc")
	if err != nil {
		return nil, err
	}
	return &Socketsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewSocketsrcWithName(name string) (*Socketsrc, error) {
	e, err := gst.NewElementWithName("socketsrc", name)
	if err != nil {
		return nil, err
	}
	return &Socketsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// The caps of the source pad

func (e *Socketsrc) GetCaps() (*gst.Caps, error) {
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

func (e *Socketsrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// send-messages (bool)
//
// Control if the source will handle GstNetworkMessage events.
// The event is a CUSTOM event named 'GstNetworkMessage' and contains:

// socket (GstGsocket)
//
// The socket to receive packets from

func (e *Socketsrc) GetSocket() (interface{}, error) {
	return e.Element.GetProperty("socket")
}

func (e *Socketsrc) SetSocket(value interface{}) error {
	return e.Element.SetProperty("socket", value)
}

