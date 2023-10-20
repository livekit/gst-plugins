// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Sends FLV content to a server via RTMP
type GstRTMPSink struct {
	*base.GstBaseSink
}

func NewRTMPSink() (*GstRTMPSink, error) {
	e, err := gst.NewElement("rtmpsink")
	if err != nil {
		return nil, err
	}
	return &GstRTMPSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewRTMPSinkWithName(name string) (*GstRTMPSink, error) {
	e, err := gst.NewElementWithName("rtmpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstRTMPSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// RTMP url
// Default: NULL
func (e *GstRTMPSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstRTMPSink) GetLocation() (string, error) {
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

// Read RTMP streams
type GstRTMPSrc struct {
	*base.GstPushSrc
}

func NewRTMPSrc() (*GstRTMPSrc, error) {
	e, err := gst.NewElement("rtmpsrc")
	if err != nil {
		return nil, err
	}
	return &GstRTMPSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewRTMPSrcWithName(name string) (*GstRTMPSrc, error) {
	e, err := gst.NewElementWithName("rtmpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstRTMPSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Location of the RTMP url to read
// Default: NULL
func (e *GstRTMPSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstRTMPSrc) GetLocation() (string, error) {
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

// Time without receiving any data from the server to wait before to timeout the session
// Default: 120, Min: 0, Max: 2147483647
func (e *GstRTMPSrc) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstRTMPSrc) GetTimeout() (int, error) {
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
