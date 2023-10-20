// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Output audio to a JACK server
type GstJackAudioSink struct {
	*GstAudioBaseSink
}

func NewJackAudioSink() (*GstJackAudioSink, error) {
	e, err := gst.NewElement("jackaudiosink")
	if err != nil {
		return nil, err
	}
	return &GstJackAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewJackAudioSinkWithName(name string) (*GstJackAudioSink, error) {
	e, err := gst.NewElementWithName("jackaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &GstJackAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Handle for jack client

func (e *GstJackAudioSink) SetClient(value interface{}) error {
	return e.Element.SetProperty("client", value)
}

func (e *GstJackAudioSink) GetClient() (interface{}, error) {
	return e.Element.GetProperty("client")
}

// The client name of the Jack instance (NULL = default)
// Default: NULL
func (e *GstJackAudioSink) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

func (e *GstJackAudioSink) GetClientName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("client-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Specify how the output ports will be connected
// Default: auto (1)
func (e *GstJackAudioSink) SetConnect(value GstJackConnect) error {
	e.Element.SetArg("connect", string(value))
	return nil
}

func (e *GstJackAudioSink) GetConnect() (GstJackConnect, error) {
	var value GstJackConnect
	var ok bool
	v, err := e.Element.GetProperty("connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstJackConnect)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstJackConnect")
	}
	return value, nil
}

// Optimize all settings for lowest latency. When enabled, "buffer-time" and "latency-time" will be ignored
// Default: false
func (e *GstJackAudioSink) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstJackAudioSink) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Comma-separated list of port name including "client_name:" prefix
// Default: NULL
func (e *GstJackAudioSink) SetPortNames(value string) error {
	return e.Element.SetProperty("port-names", value)
}

func (e *GstJackAudioSink) GetPortNames() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("port-names")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// A pattern to select which ports to connect to (NULL = first physical ports)
// Default: NULL
func (e *GstJackAudioSink) SetPortPattern(value string) error {
	return e.Element.SetProperty("port-pattern", value)
}

func (e *GstJackAudioSink) GetPortPattern() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("port-pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The Jack server to connect to (NULL = default)
// Default: NULL
func (e *GstJackAudioSink) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

func (e *GstJackAudioSink) GetServer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("server")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Jack transport behaviour of the client
// Default: (none)
func (e *GstJackAudioSink) SetTransport(value GstJackTransport) error {
	e.Element.SetArg("transport", fmt.Sprint(value))
	return nil
}

func (e *GstJackAudioSink) GetTransport() (GstJackTransport, error) {
	var value GstJackTransport
	var ok bool
	v, err := e.Element.GetProperty("transport")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstJackTransport)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstJackTransport")
	}
	return value, nil
}

// Captures audio from a JACK server
type GstJackAudioSrc struct {
	*GstAudioBaseSrc
}

func NewJackAudioSrc() (*GstJackAudioSrc, error) {
	e, err := gst.NewElement("jackaudiosrc")
	if err != nil {
		return nil, err
	}
	return &GstJackAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewJackAudioSrcWithName(name string) (*GstJackAudioSrc, error) {
	e, err := gst.NewElementWithName("jackaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstJackAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// The client name of the Jack instance (NULL = default)
// Default: NULL
func (e *GstJackAudioSrc) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

func (e *GstJackAudioSrc) GetClientName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("client-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Specify how the input ports will be connected
// Default: auto (1)
func (e *GstJackAudioSrc) SetConnect(value GstJackConnect) error {
	e.Element.SetArg("connect", string(value))
	return nil
}

func (e *GstJackAudioSrc) GetConnect() (GstJackConnect, error) {
	var value GstJackConnect
	var ok bool
	v, err := e.Element.GetProperty("connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstJackConnect)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstJackConnect")
	}
	return value, nil
}

// Optimize all settings for lowest latency. When enabled, "buffer-time" and "latency-time" will be ignored
// Default: false
func (e *GstJackAudioSrc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstJackAudioSrc) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Comma-separated list of port name including "client_name:" prefix
// Default: NULL
func (e *GstJackAudioSrc) SetPortNames(value string) error {
	return e.Element.SetProperty("port-names", value)
}

func (e *GstJackAudioSrc) GetPortNames() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("port-names")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// A pattern to select which ports to connect to (NULL = first physical ports)
// Default: NULL
func (e *GstJackAudioSrc) SetPortPattern(value string) error {
	return e.Element.SetProperty("port-pattern", value)
}

func (e *GstJackAudioSrc) GetPortPattern() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("port-pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The Jack server to connect to (NULL = default)
// Default: NULL
func (e *GstJackAudioSrc) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

func (e *GstJackAudioSrc) GetServer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("server")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Jack transport behaviour of the client
// Default: (none)
func (e *GstJackAudioSrc) SetTransport(value GstJackTransport) error {
	e.Element.SetArg("transport", fmt.Sprint(value))
	return nil
}

func (e *GstJackAudioSrc) GetTransport() (GstJackTransport, error) {
	var value GstJackTransport
	var ok bool
	v, err := e.Element.GetProperty("transport")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstJackTransport)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstJackTransport")
	}
	return value, nil
}

// Handle for jack client

func (e *GstJackAudioSrc) SetClient(value interface{}) error {
	return e.Element.SetProperty("client", value)
}

func (e *GstJackAudioSrc) GetClient() (interface{}, error) {
	return e.Element.GetProperty("client")
}

type GstJackConnect string

const (
	GstJackConnectNone       GstJackConnect = "none"        // Don't automatically connect ports to physical ports
	GstJackConnectAuto       GstJackConnect = "auto"        // Automatically connect ports to physical ports
	GstJackConnectAutoForced GstJackConnect = "auto-forced" // Automatically connect ports to as many physical ports as possible
	GstJackConnectExplicit   GstJackConnect = "explicit"    // Connect ports to explicitly requested physical ports
)

type GstJackTransport int

const (
	GstJackTransportMaster GstJackTransport = 0x00000001 // Start and stop transport with state changes
	GstJackTransportSlave  GstJackTransport = 0x00000002 // Follow transport state changes
)
