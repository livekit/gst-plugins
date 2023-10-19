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

type Shout2Send struct {
	*base.GstBaseSink
}

func NewShout2Send() (*Shout2Send, error) {
	e, err := gst.NewElement("shout2send")
	if err != nil {
		return nil, err
	}
	return &Shout2Send{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewShout2SendWithName(name string) (*Shout2Send, error) {
	e, err := gst.NewElementWithName("shout2send", name)
	if err != nil {
		return nil, err
	}
	return &Shout2Send{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// description (string)
//
// description

func (e *Shout2Send) GetDescription() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("description")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetDescription(value string) error {
	return e.Element.SetProperty("description", value)
}

// genre (string)
//
// genre

func (e *Shout2Send) GetGenre() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("genre")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetGenre(value string) error {
	return e.Element.SetProperty("genre", value)
}

// ip (string)
//
// IP address or hostname

func (e *Shout2Send) GetIp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ip")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetIp(value string) error {
	return e.Element.SetProperty("ip", value)
}

// mount (string)
//
// mount

func (e *Shout2Send) GetMount() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mount")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetMount(value string) error {
	return e.Element.SetProperty("mount", value)
}

// password (string)
//
// password

func (e *Shout2Send) GetPassword() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("password")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetPassword(value string) error {
	return e.Element.SetProperty("password", value)
}

// port (int)
//
// port

func (e *Shout2Send) GetPort() (int, error) {
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

func (e *Shout2Send) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

// protocol (GstShout2SendProtocol)
//
// Connection Protocol to use

func (e *Shout2Send) GetProtocol() (GstShout2SendProtocol, error) {
	var value GstShout2SendProtocol
	var ok bool
	v, err := e.Element.GetProperty("protocol")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstShout2SendProtocol)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstShout2SendProtocol")
	}
	return value, nil
}

func (e *Shout2Send) SetProtocol(value GstShout2SendProtocol) error {
	e.Element.SetArg("protocol", string(value))
	return nil
}

// public (bool)
//
// If the stream should be listed on the server's stream directory

func (e *Shout2Send) GetPublic() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("public")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Shout2Send) SetPublic(value bool) error {
	return e.Element.SetProperty("public", value)
}

// send-title-info (bool)
//
// Update stream metadata with song title and artist information

func (e *Shout2Send) GetSendTitleInfo() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-title-info")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Shout2Send) SetSendTitleInfo(value bool) error {
	return e.Element.SetProperty("send-title-info", value)
}

// streamname (string)
//
// name of the stream

func (e *Shout2Send) GetStreamname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("streamname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetStreamname(value string) error {
	return e.Element.SetProperty("streamname", value)
}

// timeout (uint)
//
// Max amount of time to wait for network activity, in milliseconds

func (e *Shout2Send) GetTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Shout2Send) SetTimeout(value uint) error {
	return e.Element.SetProperty("timeout", value)
}

// url (string)
//
// the stream's homepage URL

func (e *Shout2Send) GetUrl() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("url")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetUrl(value string) error {
	return e.Element.SetProperty("url", value)
}

// user-agent (string)
//
// User agent of the source

// username (string)
//
// username

func (e *Shout2Send) GetUsername() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("username")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Shout2Send) SetUsername(value string) error {
	return e.Element.SetProperty("username", value)
}

// ----- Constants -----

type GstShout2SendProtocol string

const (
	GstShout2SendProtocolXaudiocast GstShout2SendProtocol = "xaudiocast" // xaudiocast (1) – Xaudiocast Protocol (icecast 1.3.x)
	GstShout2SendProtocolIcy GstShout2SendProtocol = "icy" // icy (2) – Icy Protocol (ShoutCast)
	GstShout2SendProtocolHttp GstShout2SendProtocol = "http" // http (3) – Http Protocol (icecast 2.x)
)

