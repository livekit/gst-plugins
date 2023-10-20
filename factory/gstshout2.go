// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Sends data to an icecast server
type GstShout2send struct {
	*base.GstBaseSink
}

func NewShout2send() (*GstShout2send, error) {
	e, err := gst.NewElement("shout2send")
	if err != nil {
		return nil, err
	}
	return &GstShout2send{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewShout2sendWithName(name string) (*GstShout2send, error) {
	e, err := gst.NewElementWithName("shout2send", name)
	if err != nil {
		return nil, err
	}
	return &GstShout2send{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// IP address or hostname
// Default: 127.0.0.1
func (e *GstShout2send) SetIp(value string) error {
	return e.Element.SetProperty("ip", value)
}

func (e *GstShout2send) GetIp() (string, error) {
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

// Update stream metadata with song title and artist information
// Default: true
func (e *GstShout2send) SetSendTitleInfo(value bool) error {
	return e.Element.SetProperty("send-title-info", value)
}

func (e *GstShout2send) GetSendTitleInfo() (bool, error) {
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

// name of the stream

func (e *GstShout2send) SetStreamname(value string) error {
	return e.Element.SetProperty("streamname", value)
}

func (e *GstShout2send) GetStreamname() (string, error) {
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

// username
// Default: source
func (e *GstShout2send) SetUsername(value string) error {
	return e.Element.SetProperty("username", value)
}

func (e *GstShout2send) GetUsername() (string, error) {
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

// mount

func (e *GstShout2send) SetMount(value string) error {
	return e.Element.SetProperty("mount", value)
}

func (e *GstShout2send) GetMount() (string, error) {
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

// If the stream should be listed on the server's stream directory
// Default: false
func (e *GstShout2send) SetPublic(value bool) error {
	return e.Element.SetProperty("public", value)
}

func (e *GstShout2send) GetPublic() (bool, error) {
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

// Max amount of time to wait for network activity, in milliseconds
// Default: 10000, Min: 1, Max: -1
func (e *GstShout2send) SetTimeout(value uint) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstShout2send) GetTimeout() (uint, error) {
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

// Connection Protocol to use
// Default: http (3)
func (e *GstShout2send) SetProtocol(value GstShout2SendProtocol) error {
	e.Element.SetArg("protocol", string(value))
	return nil
}

func (e *GstShout2send) GetProtocol() (GstShout2SendProtocol, error) {
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

// the stream's homepage URL

func (e *GstShout2send) SetUrl(value string) error {
	return e.Element.SetProperty("url", value)
}

func (e *GstShout2send) GetUrl() (string, error) {
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

// description

func (e *GstShout2send) SetDescription(value string) error {
	return e.Element.SetProperty("description", value)
}

func (e *GstShout2send) GetDescription() (string, error) {
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

// genre

func (e *GstShout2send) SetGenre(value string) error {
	return e.Element.SetProperty("genre", value)
}

func (e *GstShout2send) GetGenre() (string, error) {
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

// password
// Default: hackme
func (e *GstShout2send) SetPassword(value string) error {
	return e.Element.SetProperty("password", value)
}

func (e *GstShout2send) GetPassword() (string, error) {
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

// port
// Default: 8000, Min: 1, Max: 65535
func (e *GstShout2send) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstShout2send) GetPort() (int, error) {
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

// User agent of the source
// Default: GStreamer {VERSION}
func (e *GstShout2send) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

func (e *GstShout2send) GetUserAgent() (string, error) {
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

type GstShout2SendProtocol string

const (
	GstShout2SendProtocolXaudiocast GstShout2SendProtocol = "xaudiocast" // Xaudiocast Protocol (icecast 1.3.x)
	GstShout2SendProtocolIcy        GstShout2SendProtocol = "icy"        // Icy Protocol (ShoutCast)
	GstShout2SendProtocolHttp       GstShout2SendProtocol = "http"       // Http Protocol (icecast 2.x)
)
