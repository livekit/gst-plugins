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

type Tcpserversrc struct {
	*base.GstPushSrc
}

func NewTcpserversrc() (*Tcpserversrc, error) {
	e, err := gst.NewElement("tcpserversrc")
	if err != nil {
		return nil, err
	}
	return &Tcpserversrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewTcpserversrcWithName(name string) (*Tcpserversrc, error) {
	e, err := gst.NewElementWithName("tcpserversrc", name)
	if err != nil {
		return nil, err
	}
	return &Tcpserversrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// current-port (int)
//
// The port number the socket is currently bound to. Applications can use
// this property to retrieve the port number actually bound to in case
// the port requested was 0 (=allocate a random available port).

func (e *Tcpserversrc) GetCurrentPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("current-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// host (string)
//
// The hostname to listen as

func (e *Tcpserversrc) GetHost() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("host")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Tcpserversrc) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

// port (int)
//
// The port to listen to (0=random available port)

func (e *Tcpserversrc) GetPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tcpserversrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

// stats (GstStructure)
//
// Retrieve a statistics structure

func (e *Tcpserversrc) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

