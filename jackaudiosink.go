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

type Jackaudiosink struct {
	*base.GstBaseSink
}

func NewJackaudiosink() (*Jackaudiosink, error) {
	e, err := gst.NewElement("jackaudiosink")
	if err != nil {
		return nil, err
	}
	return &Jackaudiosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewJackaudiosinkWithName(name string) (*Jackaudiosink, error) {
	e, err := gst.NewElementWithName("jackaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &Jackaudiosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// client (GstJackClient)
//
// Handle for jack client

func (e *Jackaudiosink) GetClient() (interface{}, error) {
	return e.Element.GetProperty("client")
}

func (e *Jackaudiosink) SetClient(value interface{}) error {
	return e.Element.SetProperty("client", value)
}

// client-name (string)
//
// The client name to use.

func (e *Jackaudiosink) GetClientName() (string, error) {
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

func (e *Jackaudiosink) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

// connect (GstJackConnect)
//
// Specify how the output ports will be connected

func (e *Jackaudiosink) GetConnect() (interface{}, error) {
	return e.Element.GetProperty("connect")
}

func (e *Jackaudiosink) SetConnect(value interface{}) error {
	return e.Element.SetProperty("connect", value)
}

// low-latency (bool)
//
// Optimize all settings for lowest latency. When enabled,
// buffer-time and latency-time will be
// ignored.

func (e *Jackaudiosink) GetLowLatency() (bool, error) {
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

func (e *Jackaudiosink) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// port-names (string)
//
// Comma-separated list of port name including "client_name:" prefix

func (e *Jackaudiosink) GetPortNames() (string, error) {
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

func (e *Jackaudiosink) SetPortNames(value string) error {
	return e.Element.SetProperty("port-names", value)
}

// port-pattern (string)
//
// autoconnect to ports matching pattern, when NULL connect to physical ports

func (e *Jackaudiosink) GetPortPattern() (string, error) {
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

func (e *Jackaudiosink) SetPortPattern(value string) error {
	return e.Element.SetProperty("port-pattern", value)
}

// server (string)
//
// The Jack server to connect to (NULL = default)

func (e *Jackaudiosink) GetServer() (string, error) {
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

func (e *Jackaudiosink) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

// transport (GstJackTransport)
//
// The jack transport behaviour for the client.

func (e *Jackaudiosink) GetTransport() (interface{}, error) {
	return e.Element.GetProperty("transport")
}

func (e *Jackaudiosink) SetTransport(value interface{}) error {
	return e.Element.SetProperty("transport", value)
}

