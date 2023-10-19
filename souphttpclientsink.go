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

type Souphttpclientsink struct {
	*base.GstBaseSink
}

func NewSouphttpclientsink() (*Souphttpclientsink, error) {
	e, err := gst.NewElement("souphttpclientsink")
	if err != nil {
		return nil, err
	}
	return &Souphttpclientsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewSouphttpclientsinkWithName(name string) (*Souphttpclientsink, error) {
	e, err := gst.NewElementWithName("souphttpclientsink", name)
	if err != nil {
		return nil, err
	}
	return &Souphttpclientsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// automatic-redirect (bool)
//
// Automatically follow HTTP redirects (HTTP Status Code 3xx)

func (e *Souphttpclientsink) GetAutomaticRedirect() (bool, error) {
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

func (e *Souphttpclientsink) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

// cookies (GstGstrv)
//
// HTTP request cookies

func (e *Souphttpclientsink) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

func (e *Souphttpclientsink) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

// http-log-level (GstSoupLoggerLogLevel)
//
// Set log level for soup's HTTP session log

func (e *Souphttpclientsink) GetHttpLogLevel() (interface{}, error) {
	return e.Element.GetProperty("http-log-level")
}

func (e *Souphttpclientsink) SetHttpLogLevel(value interface{}) error {
	return e.Element.SetProperty("http-log-level", value)
}

// location (string)
//
// URI to send to

func (e *Souphttpclientsink) GetLocation() (string, error) {
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

func (e *Souphttpclientsink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// proxy (string)
//
// HTTP proxy server URI

func (e *Souphttpclientsink) GetProxy() (string, error) {
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

func (e *Souphttpclientsink) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

// proxy-id (string)
//
// user id for proxy authentication

func (e *Souphttpclientsink) GetProxyId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Souphttpclientsink) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

// proxy-pw (string)
//
// user password for proxy authentication

func (e *Souphttpclientsink) GetProxyPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Souphttpclientsink) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

// retries (int)
//
// Maximum number of retries, zero to disable, -1 to retry forever

func (e *Souphttpclientsink) GetRetries() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("retries")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Souphttpclientsink) SetRetries(value int) error {
	return e.Element.SetProperty("retries", value)
}

// retry-delay (int)
//
// Delay in seconds between retries after a failure

func (e *Souphttpclientsink) GetRetryDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("retry-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Souphttpclientsink) SetRetryDelay(value int) error {
	return e.Element.SetProperty("retry-delay", value)
}

// session (GstSoupSession)
//
// SoupSession object to use for communication

func (e *Souphttpclientsink) GetSession() (interface{}, error) {
	return e.Element.GetProperty("session")
}

func (e *Souphttpclientsink) SetSession(value interface{}) error {
	return e.Element.SetProperty("session", value)
}

// user-agent (string)
//
// Value of the User-Agent HTTP request header field

func (e *Souphttpclientsink) GetUserAgent() (string, error) {
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

func (e *Souphttpclientsink) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

// user-id (string)
//
// user id for authentication

func (e *Souphttpclientsink) GetUserId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Souphttpclientsink) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

// user-pw (string)
//
// user password for authentication

func (e *Souphttpclientsink) GetUserPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Souphttpclientsink) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

