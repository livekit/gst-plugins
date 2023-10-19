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

type Curlhttpsink struct {
	*base.GstBaseSink
}

func NewCurlhttpsink() (*Curlhttpsink, error) {
	e, err := gst.NewElement("curlhttpsink")
	if err != nil {
		return nil, err
	}
	return &Curlhttpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewCurlhttpsinkWithName(name string) (*Curlhttpsink, error) {
	e, err := gst.NewElementWithName("curlhttpsink", name)
	if err != nil {
		return nil, err
	}
	return &Curlhttpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// content-type (string)
//
// Content Type to use for the Content-Type header. If not set, detected mime type will be used

func (e *Curlhttpsink) GetContentType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("content-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlhttpsink) SetContentType(value string) error {
	return e.Element.SetProperty("content-type", value)
}

// proxy (string)
//
// HTTP proxy server URI

func (e *Curlhttpsink) GetProxy() (string, error) {
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

func (e *Curlhttpsink) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

// proxy-passwd (string)
//
// Proxy user password to use for proxy authentication

func (e *Curlhttpsink) GetProxyPasswd() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-passwd")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlhttpsink) SetProxyPasswd(value string) error {
	return e.Element.SetProperty("proxy-passwd", value)
}

// proxy-port (int)
//
// HTTP proxy server port

func (e *Curlhttpsink) GetProxyPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("proxy-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Curlhttpsink) SetProxyPort(value int) error {
	return e.Element.SetProperty("proxy-port", value)
}

// proxy-user (string)
//
// Proxy user name to use for proxy authentication

func (e *Curlhttpsink) GetProxyUser() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-user")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlhttpsink) SetProxyUser(value string) error {
	return e.Element.SetProperty("proxy-user", value)
}

// use-content-length (bool)
//
// Use the Content-Length HTTP header instead of Transfer-Encoding header

func (e *Curlhttpsink) GetUseContentLength() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-content-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlhttpsink) SetUseContentLength(value bool) error {
	return e.Element.SetProperty("use-content-length", value)
}

