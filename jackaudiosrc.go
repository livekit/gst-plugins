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

type Jackaudiosrc struct {
	*base.GstPushSrc
}

func NewJackaudiosrc() (*Jackaudiosrc, error) {
	e, err := gst.NewElement("jackaudiosrc")
	if err != nil {
		return nil, err
	}
	return &Jackaudiosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewJackaudiosrcWithName(name string) (*Jackaudiosrc, error) {
	e, err := gst.NewElementWithName("jackaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &Jackaudiosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// client (GstJackClient)
//
// Handle for jack client

func (e *Jackaudiosrc) GetClient() (interface{}, error) {
	return e.Element.GetProperty("client")
}

func (e *Jackaudiosrc) SetClient(value interface{}) error {
	return e.Element.SetProperty("client", value)
}

// client-name (string)
//
// The client name to use.

func (e *Jackaudiosrc) GetClientName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("client-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Jackaudiosrc) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

// connect (GstJackConnect)
//
// Specify how the input ports will be connected

func (e *Jackaudiosrc) GetConnect() (GstJackConnect, error) {
	var value GstJackConnect
	var ok bool
	v, err := e.Element.GetProperty("connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstJackConnect)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstJackConnect")
	}
	return value, nil
}

func (e *Jackaudiosrc) SetConnect(value GstJackConnect) error {
	e.Element.SetArg("connect", string(value))
	return nil
}

// low-latency (bool)
//
// Optimize all settings for lowest latency. When enabled,
// buffer-time and latency-time will be
// ignored.

func (e *Jackaudiosrc) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Jackaudiosrc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// port-names (string)
//
// Comma-separated list of port name including "client_name:" prefix

func (e *Jackaudiosrc) GetPortNames() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("port-names")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Jackaudiosrc) SetPortNames(value string) error {
	return e.Element.SetProperty("port-names", value)
}

// port-pattern (string)
//
// autoconnect to ports matching pattern, when NULL connect to physical ports

func (e *Jackaudiosrc) GetPortPattern() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("port-pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Jackaudiosrc) SetPortPattern(value string) error {
	return e.Element.SetProperty("port-pattern", value)
}

// server (string)
//
// The Jack server to connect to (NULL = default)

func (e *Jackaudiosrc) GetServer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("server")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Jackaudiosrc) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

// transport (GstJackTransport)
//
// Jack transport behaviour of the client

func (e *Jackaudiosrc) GetTransport() (GstJackTransport, error) {
	var value GstJackTransport
	var ok bool
	v, err := e.Element.GetProperty("transport")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstJackTransport)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstJackTransport")
	}
	return value, nil
}

func (e *Jackaudiosrc) SetTransport(value GstJackTransport) error {
	e.Element.SetArg("transport", string(value))
	return nil
}

// ----- Constants -----

type GstJackConnect string

const (
	GstJackConnectNone GstJackConnect = "none" // none (0) – Don't automatically connect ports to physical ports
	GstJackConnectAuto GstJackConnect = "auto" // auto (1) – Automatically connect ports to physical ports
	GstJackConnectAutoForced GstJackConnect = "auto-forced" // auto-forced (2) – Automatically connect ports to as many physical ports as possible
	GstJackConnectExplicit GstJackConnect = "explicit" // explicit (3) – Connect ports to explicitly requested physical ports
)

type GstJackTransport string

const (
	GstJackTransportMaster GstJackTransport = "master" // master (0x00000001) – Start and stop transport with state changes
	GstJackTransportSlave GstJackTransport = "slave" // slave (0x00000002) – Follow transport state changes
)

