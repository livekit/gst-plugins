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

type Srtsrc struct {
	*base.GstPushSrc
}

func NewSrtsrc() (*Srtsrc, error) {
	e, err := gst.NewElement("srtsrc")
	if err != nil {
		return nil, err
	}
	return &Srtsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewSrtsrcWithName(name string) (*Srtsrc, error) {
	e, err := gst.NewElementWithName("srtsrc", name)
	if err != nil {
		return nil, err
	}
	return &Srtsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// latency (int)
//
// The maximum accepted transmission latency.

func (e *Srtsrc) GetLatency() (int, error) {
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

func (e *Srtsrc) SetLatency(value int) error {
	return e.Element.SetProperty("latency", value)
}

// localaddress (string)
//
// The address to bind when mode is listener or rendezvous.
// This property can be set by URI parameters.

func (e *Srtsrc) GetLocaladdress() (string, error) {
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

func (e *Srtsrc) SetLocaladdress(value string) error {
	return e.Element.SetProperty("localaddress", value)
}

// localport (uint)
//
// The local port to bind when mode is listener or rendezvous.
// This property can be set by URI parameters.

func (e *Srtsrc) GetLocalport() (uint, error) {
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

func (e *Srtsrc) SetLocalport(value uint) error {
	return e.Element.SetProperty("localport", value)
}

// mode (GstSRTConnectionMode)
//
// The SRT connection mode.
// This property can be set by URI parameters.

func (e *Srtsrc) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *Srtsrc) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// passphrase (string)
//
// The password for the encrypted transmission.
// This property can be set by URI parameters.

func (e *Srtsrc) GetPassphrase() (string, error) {
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

func (e *Srtsrc) SetPassphrase(value string) error {
	return e.Element.SetProperty("passphrase", value)
}

// pbkeylen (GstSRTKeyLength)
//
// The crypto key length.
// This property can be set by URI parameters.

func (e *Srtsrc) GetPbkeylen() (interface{}, error) {
	return e.Element.GetProperty("pbkeylen")
}

func (e *Srtsrc) SetPbkeylen(value interface{}) error {
	return e.Element.SetProperty("pbkeylen", value)
}

// poll-timeout (int)
//
// The polling timeout used when srt poll is started.
// Even if the default value indicates infinite waiting, it can be cancellable according to GstState
// This property can be set by URI parameters.

func (e *Srtsrc) GetPollTimeout() (int, error) {
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

func (e *Srtsrc) SetPollTimeout(value int) error {
	return e.Element.SetProperty("poll-timeout", value)
}

// stats (GstStructure)
//
// The statistics from SRT.

func (e *Srtsrc) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// streamid (string)
//
// The stream id for the SRT access control.

func (e *Srtsrc) GetStreamid() (string, error) {
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

func (e *Srtsrc) SetStreamid(value string) error {
	return e.Element.SetProperty("streamid", value)
}

// uri (string)
//
// The URI used by SRT connection. User can specify SRT specific options by URI parameters.
// Refer to Mediun: SRT

func (e *Srtsrc) GetUri() (string, error) {
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

func (e *Srtsrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// wait-for-connection (bool)
//
// Block the stream until a client connects

func (e *Srtsrc) GetWaitForConnection() (bool, error) {
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

func (e *Srtsrc) SetWaitForConnection(value bool) error {
	return e.Element.SetProperty("wait-for-connection", value)
}

// ----- Constants -----

type GstSrtconnectionMode string

const (
	GstSrtconnectionModeNone GstSrtconnectionMode = "none" // none (0) – GST_SRT_CONNECTION_MODE_NONE
	GstSrtconnectionModeCaller GstSrtconnectionMode = "caller" // caller (1) – GST_SRT_CONNECTION_MODE_CALLER
	GstSrtconnectionModeListener GstSrtconnectionMode = "listener" // listener (2) – GST_SRT_CONNECTION_MODE_LISTENER
	GstSrtconnectionModeRendezvous GstSrtconnectionMode = "rendezvous" // rendezvous (3) – GST_SRT_CONNECTION_MODE_RENDEZVOUS
)

type GstSrtkeyLength string

const (
	GstSrtkeyLengthNoKey GstSrtkeyLength = "no-key" // no-key (0) – GST_SRT_KEY_LENGTH_NO_KEY
	GstSrtkeyLength0 GstSrtkeyLength = "0" // 0 (0) – GST_SRT_KEY_LENGTH_0
	GstSrtkeyLength16 GstSrtkeyLength = "16" // 16 (16) – GST_SRT_KEY_LENGTH_16
	GstSrtkeyLength24 GstSrtkeyLength = "24" // 24 (24) – GST_SRT_KEY_LENGTH_24
	GstSrtkeyLength32 GstSrtkeyLength = "32" // 32 (32) – GST_SRT_KEY_LENGTH_32
)

