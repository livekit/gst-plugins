// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Send data over the network via SRT
type GstSRTSink struct {
	*base.GstBaseSink
}

func NewSRTSink() (*GstSRTSink, error) {
	e, err := gst.NewElement("srtsink")
	if err != nil {
		return nil, err
	}
	return &GstSRTSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewSRTSinkWithName(name string) (*GstSRTSink, error) {
	e, err := gst.NewElementWithName("srtsink", name)
	if err != nil {
		return nil, err
	}
	return &GstSRTSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Return poll wait after timeout milliseconds (-1 = infinite)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstSRTSink) SetPollTimeout(value int) error {
	return e.Element.SetProperty("poll-timeout", value)
}

func (e *GstSRTSink) GetPollTimeout() (int, error) {
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

// SRT Statistics
// Default: application/x-srt-statistics, bytes-sent-total=(guint64)0;
func (e *GstSRTSink) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Stream ID for the SRT access control
// Default: NULL
func (e *GstSRTSink) SetStreamid(value string) error {
	return e.Element.SetProperty("streamid", value)
}

func (e *GstSRTSink) GetStreamid() (string, error) {
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

// URI in the form of srt://address:port
// Default: srt://127.0.0.1:7001
func (e *GstSRTSink) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstSRTSink) GetUri() (string, error) {
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

// Local address to bind
// Default: NULL
func (e *GstSRTSink) SetLocaladdress(value string) error {
	return e.Element.SetProperty("localaddress", value)
}

func (e *GstSRTSink) GetLocaladdress() (string, error) {
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

// Local port to bind
// Default: 7001, Min: 0, Max: 65535
func (e *GstSRTSink) SetLocalport(value uint) error {
	return e.Element.SetProperty("localport", value)
}

func (e *GstSRTSink) GetLocalport() (uint, error) {
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

// SRT connection mode
// Default: caller (1)
func (e *GstSRTSink) SetMode(value GstSRTConnectionMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstSRTSink) GetMode() (GstSRTConnectionMode, error) {
	var value GstSRTConnectionMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSRTConnectionMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSRTConnectionMode")
	}
	return value, nil
}

// Password for the encrypted transmission

func (e *GstSRTSink) SetPassphrase(value string) error {
	return e.Element.SetProperty("passphrase", value)
}

// Minimum latency (milliseconds)
// Default: 125, Min: 0, Max: 2147483647
func (e *GstSRTSink) SetLatency(value int) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstSRTSink) GetLatency() (int, error) {
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

// Crypto key length in bytes
// Default: no-key (0)
func (e *GstSRTSink) SetPbkeylen(value GstSRTKeyLength) error {
	e.Element.SetArg("pbkeylen", string(value))
	return nil
}

func (e *GstSRTSink) GetPbkeylen() (GstSRTKeyLength, error) {
	var value GstSRTKeyLength
	var ok bool
	v, err := e.Element.GetProperty("pbkeylen")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSRTKeyLength)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSRTKeyLength")
	}
	return value, nil
}

// Block the stream until a client connects
// Default: true
func (e *GstSRTSink) SetWaitForConnection(value bool) error {
	return e.Element.SetProperty("wait-for-connection", value)
}

func (e *GstSRTSink) GetWaitForConnection() (bool, error) {
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

// Receive data over the network via SRT
type GstSRTSrc struct {
	*base.GstPushSrc
}

func NewSRTSrc() (*GstSRTSrc, error) {
	e, err := gst.NewElement("srtsrc")
	if err != nil {
		return nil, err
	}
	return &GstSRTSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewSRTSrcWithName(name string) (*GstSRTSrc, error) {
	e, err := gst.NewElementWithName("srtsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSRTSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// SRT Statistics
// Default: application/x-srt-statistics, bytes-received-total=(guint64)0;
func (e *GstSRTSrc) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Stream ID for the SRT access control
// Default: NULL
func (e *GstSRTSrc) SetStreamid(value string) error {
	return e.Element.SetProperty("streamid", value)
}

func (e *GstSRTSrc) GetStreamid() (string, error) {
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

// URI in the form of srt://address:port
// Default: srt://127.0.0.1:7001
func (e *GstSRTSrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstSRTSrc) GetUri() (string, error) {
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

// SRT connection mode
// Default: caller (1)
func (e *GstSRTSrc) SetMode(value GstSRTConnectionMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstSRTSrc) GetMode() (GstSRTConnectionMode, error) {
	var value GstSRTConnectionMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSRTConnectionMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSRTConnectionMode")
	}
	return value, nil
}

// Password for the encrypted transmission

func (e *GstSRTSrc) SetPassphrase(value string) error {
	return e.Element.SetProperty("passphrase", value)
}

// Crypto key length in bytes
// Default: no-key (0)
func (e *GstSRTSrc) SetPbkeylen(value GstSRTKeyLength) error {
	e.Element.SetArg("pbkeylen", string(value))
	return nil
}

func (e *GstSRTSrc) GetPbkeylen() (GstSRTKeyLength, error) {
	var value GstSRTKeyLength
	var ok bool
	v, err := e.Element.GetProperty("pbkeylen")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSRTKeyLength)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSRTKeyLength")
	}
	return value, nil
}

// Return poll wait after timeout milliseconds (-1 = infinite)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstSRTSrc) SetPollTimeout(value int) error {
	return e.Element.SetProperty("poll-timeout", value)
}

func (e *GstSRTSrc) GetPollTimeout() (int, error) {
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

// Block the stream until a client connects
// Default: true
func (e *GstSRTSrc) SetWaitForConnection(value bool) error {
	return e.Element.SetProperty("wait-for-connection", value)
}

func (e *GstSRTSrc) GetWaitForConnection() (bool, error) {
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

// Minimum latency (milliseconds)
// Default: 125, Min: 0, Max: 2147483647
func (e *GstSRTSrc) SetLatency(value int) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstSRTSrc) GetLatency() (int, error) {
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

// Local address to bind
// Default: NULL
func (e *GstSRTSrc) SetLocaladdress(value string) error {
	return e.Element.SetProperty("localaddress", value)
}

func (e *GstSRTSrc) GetLocaladdress() (string, error) {
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

// Local port to bind
// Default: 7001, Min: 0, Max: 65535
func (e *GstSRTSrc) SetLocalport(value uint) error {
	return e.Element.SetProperty("localport", value)
}

func (e *GstSRTSrc) GetLocalport() (uint, error) {
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

// Send data over the network via SRT
type GstSRTClientSink struct {
	*GstSRTSink
}

func NewSRTClientSink() (*GstSRTClientSink, error) {
	e, err := gst.NewElement("srtclientsink")
	if err != nil {
		return nil, err
	}
	return &GstSRTClientSink{GstSRTSink: &GstSRTSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewSRTClientSinkWithName(name string) (*GstSRTClientSink, error) {
	e, err := gst.NewElementWithName("srtclientsink", name)
	if err != nil {
		return nil, err
	}
	return &GstSRTClientSink{GstSRTSink: &GstSRTSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Receive data over the network via SRT
type GstSRTClientSrc struct {
	*GstSRTSrc
}

func NewSRTClientSrc() (*GstSRTClientSrc, error) {
	e, err := gst.NewElement("srtclientsrc")
	if err != nil {
		return nil, err
	}
	return &GstSRTClientSrc{GstSRTSrc: &GstSRTSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewSRTClientSrcWithName(name string) (*GstSRTClientSrc, error) {
	e, err := gst.NewElementWithName("srtclientsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSRTClientSrc{GstSRTSrc: &GstSRTSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// Send data over the network via SRT
type GstSRTServerSink struct {
	*GstSRTSink
}

func NewSRTServerSink() (*GstSRTServerSink, error) {
	e, err := gst.NewElement("srtserversink")
	if err != nil {
		return nil, err
	}
	return &GstSRTServerSink{GstSRTSink: &GstSRTSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewSRTServerSinkWithName(name string) (*GstSRTServerSink, error) {
	e, err := gst.NewElementWithName("srtserversink", name)
	if err != nil {
		return nil, err
	}
	return &GstSRTServerSink{GstSRTSink: &GstSRTSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Receive data over the network via SRT
type GstSRTServerSrc struct {
	*GstSRTSrc
}

func NewSRTServerSrc() (*GstSRTServerSrc, error) {
	e, err := gst.NewElement("srtserversrc")
	if err != nil {
		return nil, err
	}
	return &GstSRTServerSrc{GstSRTSrc: &GstSRTSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewSRTServerSrcWithName(name string) (*GstSRTServerSrc, error) {
	e, err := gst.NewElementWithName("srtserversrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSRTServerSrc{GstSRTSrc: &GstSRTSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

type GstSRTKeyLength string

const (
	GstSRTKeyLengthNoKey GstSRTKeyLength = "no-key" // GST_SRT_KEY_LENGTH_NO_KEY
	GstSRTKeyLength0     GstSRTKeyLength = "0"      // GST_SRT_KEY_LENGTH_0
	GstSRTKeyLength16    GstSRTKeyLength = "16"     // GST_SRT_KEY_LENGTH_16
	GstSRTKeyLength24    GstSRTKeyLength = "24"     // GST_SRT_KEY_LENGTH_24
	GstSRTKeyLength32    GstSRTKeyLength = "32"     // GST_SRT_KEY_LENGTH_32
)

type GstSRTConnectionMode string

const (
	GstSRTConnectionModeNone       GstSRTConnectionMode = "none"       // GST_SRT_CONNECTION_MODE_NONE
	GstSRTConnectionModeCaller     GstSRTConnectionMode = "caller"     // GST_SRT_CONNECTION_MODE_CALLER
	GstSRTConnectionModeListener   GstSRTConnectionMode = "listener"   // GST_SRT_CONNECTION_MODE_LISTENER
	GstSRTConnectionModeRendezvous GstSRTConnectionMode = "rendezvous" // GST_SRT_CONNECTION_MODE_RENDEZVOUS
)
