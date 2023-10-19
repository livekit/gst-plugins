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

type Curlhttpsrc struct {
	*base.GstPushSrc
}

func NewCurlhttpsrc() (*Curlhttpsrc, error) {
	e, err := gst.NewElement("curlhttpsrc")
	if err != nil {
		return nil, err
	}
	return &Curlhttpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewCurlhttpsrcWithName(name string) (*Curlhttpsrc, error) {
	e, err := gst.NewElementWithName("curlhttpsrc", name)
	if err != nil {
		return nil, err
	}
	return &Curlhttpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// automatic-redirect (bool)
//
// Allow HTTP Redirections (HTTP Status Code 300 series)

func (e *Curlhttpsrc) GetAutomaticRedirect() (bool, error) {
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

func (e *Curlhttpsrc) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

// compress (bool)
//
// Allow compressed content encodings

func (e *Curlhttpsrc) GetCompress() (bool, error) {
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

func (e *Curlhttpsrc) SetCompress(value bool) error {
	return e.Element.SetProperty("compress", value)
}

// cookies (GstGstrv)
//
// List of HTTP Cookies

func (e *Curlhttpsrc) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

func (e *Curlhttpsrc) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

// extra-headers (GstStructure)
//
// Extra headers to append to the HTTP request

func (e *Curlhttpsrc) GetExtraHeaders() (interface{}, error) {
	return e.Element.GetProperty("extra-headers")
}

func (e *Curlhttpsrc) SetExtraHeaders(value interface{}) error {
	return e.Element.SetProperty("extra-headers", value)
}

// http-version (GstCurlHttpVersionType)
//
// The preferred HTTP protocol version

func (e *Curlhttpsrc) GetHttpVersion() (GstCurlHttpVersionType, error) {
	var value GstCurlHttpVersionType
	var ok bool
	v, err := e.Element.GetProperty("http-version")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCurlHttpVersionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCurlHttpVersionType")
	}
	return value, nil
}

func (e *Curlhttpsrc) SetHttpVersion(value GstCurlHttpVersionType) error {
	e.Element.SetArg("http-version", string(value))
	return nil
}

// keep-alive (bool)
//
// Toggle keep-alive for connection reuse.

func (e *Curlhttpsrc) GetKeepAlive() (bool, error) {
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

func (e *Curlhttpsrc) SetKeepAlive(value bool) error {
	return e.Element.SetProperty("keep-alive", value)
}

// location (string)
//
// URI of resource to read

func (e *Curlhttpsrc) GetLocation() (string, error) {
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

func (e *Curlhttpsrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-connection-time (uint)
//
// Maximum amount of time to keep-alive HTTP connections

func (e *Curlhttpsrc) GetMaxConnectionTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connection-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Curlhttpsrc) SetMaxConnectionTime(value uint) error {
	return e.Element.SetProperty("max-connection-time", value)
}

// max-connections (uint)
//
// Maximum number of concurrent connections allowed for HTTP/1.x

func (e *Curlhttpsrc) GetMaxConnections() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connections")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Curlhttpsrc) SetMaxConnections(value uint) error {
	return e.Element.SetProperty("max-connections", value)
}

// max-connections-per-proxy (uint)
//
// Maximum number of concurrent connections allowed per proxy for HTTP/1.x

func (e *Curlhttpsrc) GetMaxConnectionsPerProxy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connections-per-proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Curlhttpsrc) SetMaxConnectionsPerProxy(value uint) error {
	return e.Element.SetProperty("max-connections-per-proxy", value)
}

// max-connections-per-server (uint)
//
// Maximum number of connections allowed per server for HTTP/1.x

func (e *Curlhttpsrc) GetMaxConnectionsPerServer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connections-per-server")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Curlhttpsrc) SetMaxConnectionsPerServer(value uint) error {
	return e.Element.SetProperty("max-connections-per-server", value)
}

// max-redirect (int)
//
// Maximum number of permitted redirections. -1 is unlimited.

func (e *Curlhttpsrc) GetMaxRedirect() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-redirect")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Curlhttpsrc) SetMaxRedirect(value int) error {
	return e.Element.SetProperty("max-redirect", value)
}

// proxy (string)
//
// URI of HTTP proxy server

func (e *Curlhttpsrc) GetProxy() (string, error) {
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

func (e *Curlhttpsrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

// proxy-id (string)
//
// HTTP proxy URI user id for authentication

func (e *Curlhttpsrc) GetProxyId() (string, error) {
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

func (e *Curlhttpsrc) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

// proxy-pw (string)
//
// HTTP proxy URI password for authentication

func (e *Curlhttpsrc) GetProxyPw() (string, error) {
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

func (e *Curlhttpsrc) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

// retries (int)
//
// Maximum number of retries until giving up (-1=infinite)

func (e *Curlhttpsrc) GetRetries() (int, error) {
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

func (e *Curlhttpsrc) SetRetries(value int) error {
	return e.Element.SetProperty("retries", value)
}

// ssl-ca-file (string)
//
// Location of an SSL CA file to use for checking SSL certificates

func (e *Curlhttpsrc) GetSslCaFile() (string, error) {
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

func (e *Curlhttpsrc) SetSslCaFile(value string) error {
	return e.Element.SetProperty("ssl-ca-file", value)
}

// ssl-strict (bool)
//
// Strict SSL certificate checking

func (e *Curlhttpsrc) GetSslStrict() (bool, error) {
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

func (e *Curlhttpsrc) SetSslStrict(value bool) error {
	return e.Element.SetProperty("ssl-strict", value)
}

// timeout (int)
//
// Value in seconds before timeout a blocking request (0 = no timeout)

func (e *Curlhttpsrc) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Curlhttpsrc) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

// user-agent (string)
//
// URI of resource requested

func (e *Curlhttpsrc) GetUserAgent() (string, error) {
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

func (e *Curlhttpsrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

// user-id (string)
//
// HTTP location URI user id for authentication

func (e *Curlhttpsrc) GetUserId() (string, error) {
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

func (e *Curlhttpsrc) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

// user-pw (string)
//
// HTTP location URI password for authentication

func (e *Curlhttpsrc) GetUserPw() (string, error) {
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

func (e *Curlhttpsrc) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

// ----- Constants -----

type GstCurlHttpVersionType string

const (
	GstCurlHttpVersionType10 GstCurlHttpVersionType = "1.0" // 1.0 (0) – HTTP Version 1.0
	GstCurlHttpVersionType11 GstCurlHttpVersionType = "1.1" // 1.1 (1) – HTTP Version 1.1
	GstCurlHttpVersionType20 GstCurlHttpVersionType = "2.0" // 2.0 (2) – HTTP Version 2.0
)

