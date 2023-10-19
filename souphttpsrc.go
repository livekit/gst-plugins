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

type Souphttpsrc struct {
	*base.GstPushSrc
}

func NewSouphttpsrc() (*Souphttpsrc, error) {
	e, err := gst.NewElement("souphttpsrc")
	if err != nil {
		return nil, err
	}
	return &Souphttpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewSouphttpsrcWithName(name string) (*Souphttpsrc, error) {
	e, err := gst.NewElementWithName("souphttpsrc", name)
	if err != nil {
		return nil, err
	}
	return &Souphttpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// automatic-redirect (bool)
//
// Automatically follow HTTP redirects (HTTP Status Code 3xx)

func (e *Souphttpsrc) GetAutomaticRedirect() (bool, error) {
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

func (e *Souphttpsrc) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

// compress (bool)
//
// Allow compressed content encodings

func (e *Souphttpsrc) GetCompress() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("compress")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Souphttpsrc) SetCompress(value bool) error {
	return e.Element.SetProperty("compress", value)
}

// cookies (GstGstrv)
//
// HTTP request cookies

func (e *Souphttpsrc) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

func (e *Souphttpsrc) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

// extra-headers (GstStructure)
//
// Extra headers to append to the HTTP request

func (e *Souphttpsrc) GetExtraHeaders() (interface{}, error) {
	return e.Element.GetProperty("extra-headers")
}

func (e *Souphttpsrc) SetExtraHeaders(value interface{}) error {
	return e.Element.SetProperty("extra-headers", value)
}

// http-log-level (GstSoupLoggerLogLevel)
//
// Set log level for soup's HTTP session log

func (e *Souphttpsrc) GetHttpLogLevel() (interface{}, error) {
	return e.Element.GetProperty("http-log-level")
}

func (e *Souphttpsrc) SetHttpLogLevel(value interface{}) error {
	return e.Element.SetProperty("http-log-level", value)
}

// iradio-mode (bool)
//
// Enable internet radio mode (ask server to send shoutcast/icecast metadata interleaved with the actual stream data)

func (e *Souphttpsrc) GetIradioMode() (bool, error) {
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

func (e *Souphttpsrc) SetIradioMode(value bool) error {
	return e.Element.SetProperty("iradio-mode", value)
}

// is-live (bool)
//
// Act like a live source

func (e *Souphttpsrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Souphttpsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// keep-alive (bool)
//
// Use HTTP persistent connections

func (e *Souphttpsrc) GetKeepAlive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-alive")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Souphttpsrc) SetKeepAlive(value bool) error {
	return e.Element.SetProperty("keep-alive", value)
}

// location (string)
//
// Location to read from

func (e *Souphttpsrc) GetLocation() (string, error) {
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

func (e *Souphttpsrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// method (string)
//
// The HTTP method to use (GET, HEAD, OPTIONS, etc)

func (e *Souphttpsrc) GetMethod() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Souphttpsrc) SetMethod(value string) error {
	return e.Element.SetProperty("method", value)
}

// proxy (string)
//
// HTTP proxy server URI

func (e *Souphttpsrc) GetProxy() (string, error) {
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

func (e *Souphttpsrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

// proxy-id (string)
//
// HTTP proxy URI user id for authentication

func (e *Souphttpsrc) GetProxyId() (string, error) {
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

func (e *Souphttpsrc) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

// proxy-pw (string)
//
// HTTP proxy URI user password for authentication

func (e *Souphttpsrc) GetProxyPw() (string, error) {
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

func (e *Souphttpsrc) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

// retries (int)
//
// Maximum number of retries until giving up (-1=infinite)

func (e *Souphttpsrc) GetRetries() (int, error) {
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

func (e *Souphttpsrc) SetRetries(value int) error {
	return e.Element.SetProperty("retries", value)
}

// ssl-ca-file (string)
//
// Location of a SSL anchor CA file to use

func (e *Souphttpsrc) GetSslCaFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssl-ca-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Souphttpsrc) SetSslCaFile(value string) error {
	return e.Element.SetProperty("ssl-ca-file", value)
}

// ssl-strict (bool)
//
// Strict SSL certificate checking

func (e *Souphttpsrc) GetSslStrict() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssl-strict")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Souphttpsrc) SetSslStrict(value bool) error {
	return e.Element.SetProperty("ssl-strict", value)
}

// ssl-use-system-ca-file (bool)
//
// Use system CA file

func (e *Souphttpsrc) GetSslUseSystemCaFile() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssl-use-system-ca-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Souphttpsrc) SetSslUseSystemCaFile(value bool) error {
	return e.Element.SetProperty("ssl-use-system-ca-file", value)
}

// timeout (uint)
//
// Value in seconds to timeout a blocking I/O (0 = No timeout).

func (e *Souphttpsrc) GetTimeout() (uint, error) {
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

func (e *Souphttpsrc) SetTimeout(value uint) error {
	return e.Element.SetProperty("timeout", value)
}

// tls-database (GstGtlsDatabase)
//
// TLS database with anchor certificate authorities used to validate the server certificate

func (e *Souphttpsrc) GetTlsDatabase() (interface{}, error) {
	return e.Element.GetProperty("tls-database")
}

func (e *Souphttpsrc) SetTlsDatabase(value interface{}) error {
	return e.Element.SetProperty("tls-database", value)
}

// tls-interaction (GstGtlsInteraction)
//
// A GTlsInteraction object to be used when the connection or certificate database need to interact with the user.

func (e *Souphttpsrc) GetTlsInteraction() (interface{}, error) {
	return e.Element.GetProperty("tls-interaction")
}

func (e *Souphttpsrc) SetTlsInteraction(value interface{}) error {
	return e.Element.SetProperty("tls-interaction", value)
}

// user-agent (string)
//
// Value of the User-Agent HTTP request header field

func (e *Souphttpsrc) GetUserAgent() (string, error) {
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

func (e *Souphttpsrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

// user-id (string)
//
// HTTP location URI user id for authentication

func (e *Souphttpsrc) GetUserId() (string, error) {
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

func (e *Souphttpsrc) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

// user-pw (string)
//
// HTTP location URI user password for authentication

func (e *Souphttpsrc) GetUserPw() (string, error) {
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

func (e *Souphttpsrc) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

