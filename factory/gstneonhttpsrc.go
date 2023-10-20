// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Receive data as a client over the network via HTTP using NEON
type GstNeonhttpSrc struct {
	*base.GstPushSrc
}

func NewNeonhttpSrc() (*GstNeonhttpSrc, error) {
	e, err := gst.NewElement("neonhttpsrc")
	if err != nil {
		return nil, err
	}
	return &GstNeonhttpSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewNeonhttpSrcWithName(name string) (*GstNeonhttpSrc, error) {
	e, err := gst.NewElementWithName("neonhttpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstNeonhttpSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Value in seconds to timeout a blocking read (0 = default).
// Default: 0, Min: 0, Max: 3600
func (e *GstNeonhttpSrc) SetReadTimeout(value uint) error {
	return e.Element.SetProperty("read-timeout", value)
}

func (e *GstNeonhttpSrc) GetReadTimeout() (uint, error) {
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

// Automatically follow HTTP redirects (HTTP Status Code 3xx)
// Default: true
func (e *GstNeonhttpSrc) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

func (e *GstNeonhttpSrc) GetAutomaticRedirect() (bool, error) {
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

// HTTP request cookies

func (e *GstNeonhttpSrc) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

func (e *GstNeonhttpSrc) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

// Location to read from
// Default: http://localhost/
func (e *GstNeonhttpSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstNeonhttpSrc) GetLocation() (string, error) {
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

// Proxy server to use, in the form HOSTNAME:PORT. Defaults to the http_proxy environment variable

func (e *GstNeonhttpSrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

func (e *GstNeonhttpSrc) GetProxy() (string, error) {
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

// Value of the User-Agent HTTP request header field
// Default: GStreamer neonhttpsrc
func (e *GstNeonhttpSrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

func (e *GstNeonhttpSrc) GetUserAgent() (string, error) {
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

// Accept self-signed SSL/TLS certificates
// Default: false
func (e *GstNeonhttpSrc) SetAcceptSelfSigned(value bool) error {
	return e.Element.SetProperty("accept-self-signed", value)
}

func (e *GstNeonhttpSrc) GetAcceptSelfSigned() (bool, error) {
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

// Value in seconds to timeout a blocking connection (0 = default).
// Default: 0, Min: 0, Max: 3600
func (e *GstNeonhttpSrc) SetConnectTimeout(value uint) error {
	return e.Element.SetProperty("connect-timeout", value)
}

func (e *GstNeonhttpSrc) GetConnectTimeout() (uint, error) {
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

// Enable internet radio mode (ask server to send shoutcast/icecast metadata interleaved with the actual stream data)
// Default: true
func (e *GstNeonhttpSrc) SetIradioMode(value bool) error {
	return e.Element.SetProperty("iradio-mode", value)
}

func (e *GstNeonhttpSrc) GetIradioMode() (bool, error) {
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

// Enable Neon HTTP debug messages
// Default: false
func (e *GstNeonhttpSrc) SetNeonHttpDebug(value bool) error {
	return e.Element.SetProperty("neon-http-debug", value)
}

func (e *GstNeonhttpSrc) GetNeonHttpDebug() (bool, error) {
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
