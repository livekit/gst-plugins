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

type Neonhttpsrc struct {
	*base.GstPushSrc
}

func NewNeonhttpsrc() (*Neonhttpsrc, error) {
	e, err := gst.NewElement("neonhttpsrc")
	if err != nil {
		return nil, err
	}
	return &Neonhttpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewNeonhttpsrcWithName(name string) (*Neonhttpsrc, error) {
	e, err := gst.NewElementWithName("neonhttpsrc", name)
	if err != nil {
		return nil, err
	}
	return &Neonhttpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// accept-self-signed (bool)
//
// Accept self-signed SSL/TLS certificates

func (e *Neonhttpsrc) GetAcceptSelfSigned() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("accept-self-signed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetAcceptSelfSigned(value bool) error {
	return e.Element.SetProperty("accept-self-signed", value)
}

// automatic-redirect (bool)
//
// Automatically follow HTTP redirects (HTTP Status Code 3xx)

func (e *Neonhttpsrc) GetAutomaticRedirect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("automatic-redirect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

// connect-timeout (uint)
//
// Value in seconds to timeout a blocking connection (0 = default).

func (e *Neonhttpsrc) GetConnectTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("connect-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetConnectTimeout(value uint) error {
	return e.Element.SetProperty("connect-timeout", value)
}

// cookies (GstGstrv)
//
// HTTP request cookies

func (e *Neonhttpsrc) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

func (e *Neonhttpsrc) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

// iradio-mode (bool)
//
// Enable internet radio mode (ask server to send shoutcast/icecast metadata interleaved with the actual stream data)

func (e *Neonhttpsrc) GetIradioMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iradio-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetIradioMode(value bool) error {
	return e.Element.SetProperty("iradio-mode", value)
}

// location (string)
//
// Location to read from

func (e *Neonhttpsrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// neon-http-debug (bool)
//
// Enable Neon HTTP debug messages

func (e *Neonhttpsrc) GetNeonHttpDebug() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("neon-http-debug")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetNeonHttpDebug(value bool) error {
	return e.Element.SetProperty("neon-http-debug", value)
}

// proxy (string)
//
// Proxy server to use, in the form HOSTNAME:PORT. Defaults to the http_proxy environment variable

func (e *Neonhttpsrc) GetProxy() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

// read-timeout (uint)
//
// Value in seconds to timeout a blocking read (0 = default).

func (e *Neonhttpsrc) GetReadTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("read-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetReadTimeout(value uint) error {
	return e.Element.SetProperty("read-timeout", value)
}

// user-agent (string)
//
// Value of the User-Agent HTTP request header field

func (e *Neonhttpsrc) GetUserAgent() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-agent")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Neonhttpsrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

