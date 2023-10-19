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

type Srtsink struct {
	*base.GstBaseSink
}

func NewSrtsink() (*Srtsink, error) {
	e, err := gst.NewElement("srtsink")
	if err != nil {
		return nil, err
	}
	return &Srtsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewSrtsinkWithName(name string) (*Srtsink, error) {
	e, err := gst.NewElementWithName("srtsink", name)
	if err != nil {
		return nil, err
	}
	return &Srtsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// latency (int)
//
// Minimum latency (milliseconds)

func (e *Srtsink) GetLatency() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Srtsink) SetLatency(value int) error {
	return e.Element.SetProperty("latency", value)
}

// localaddress (string)
//
// Local address to bind

func (e *Srtsink) GetLocaladdress() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("localaddress")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Srtsink) SetLocaladdress(value string) error {
	return e.Element.SetProperty("localaddress", value)
}

// localport (uint)
//
// Local port to bind

func (e *Srtsink) GetLocalport() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("localport")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Srtsink) SetLocalport(value uint) error {
	return e.Element.SetProperty("localport", value)
}

// mode (GstSRTConnectionMode)
//
// SRT connection mode

func (e *Srtsink) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *Srtsink) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// passphrase (string)
//
// Password for the encrypted transmission

func (e *Srtsink) GetPassphrase() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("passphrase")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Srtsink) SetPassphrase(value string) error {
	return e.Element.SetProperty("passphrase", value)
}

// pbkeylen (GstSRTKeyLength)
//
// Crypto key length in bytes

func (e *Srtsink) GetPbkeylen() (interface{}, error) {
	return e.Element.GetProperty("pbkeylen")
}

func (e *Srtsink) SetPbkeylen(value interface{}) error {
	return e.Element.SetProperty("pbkeylen", value)
}

// poll-timeout (int)
//
// Return poll wait after timeout milliseconds (-1 = infinite)

func (e *Srtsink) GetPollTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("poll-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Srtsink) SetPollTimeout(value int) error {
	return e.Element.SetProperty("poll-timeout", value)
}

// stats (GstStructure)
//
// SRT Statistics

func (e *Srtsink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// streamid (string)
//
// Stream ID for the SRT access control

func (e *Srtsink) GetStreamid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("streamid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Srtsink) SetStreamid(value string) error {
	return e.Element.SetProperty("streamid", value)
}

// uri (string)
//
// URI in the form of srt://address:port

func (e *Srtsink) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Srtsink) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// wait-for-connection (bool)
//
// Boolean to block streaming until a client connects.  If TRUE,
// `srtsink' will stream only when a client is connected.

func (e *Srtsink) GetWaitForConnection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-for-connection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Srtsink) SetWaitForConnection(value bool) error {
	return e.Element.SetProperty("wait-for-connection", value)
}

