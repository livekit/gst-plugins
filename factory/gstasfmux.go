// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Muxes audio and video into an ASF stream
type GstAsfMux struct {
	*gst.Element
}

func NewAsfMux() (*GstAsfMux, error) {
	e, err := gst.NewElement("asfmux")
	if err != nil {
		return nil, err
	}
	return &GstAsfMux{Element: e}, nil
}

func NewAsfMuxWithName(name string) (*GstAsfMux, error) {
	e, err := gst.NewElementWithName("asfmux", name)
	if err != nil {
		return nil, err
	}
	return &GstAsfMux{Element: e}, nil
}

// If the stream metadata (received as events in the sink) should be merged to the main file metadata.
// Default: true
func (e *GstAsfMux) SetMergeStreamTags(value bool) error {
	return e.Element.SetProperty("merge-stream-tags", value)
}

func (e *GstAsfMux) GetMergeStreamTags() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("merge-stream-tags")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The ASF packets size (bytes)
// Default: 4800, Min: 18, Max: -1
func (e *GstAsfMux) SetPacketSize(value uint) error {
	return e.Element.SetProperty("packet-size", value)
}

func (e *GstAsfMux) GetPacketSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("packet-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Size of the padding object to be added to the end of the header. If this less than 24 (the smaller size of an ASF object), no padding is added.
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAsfMux) SetPadding(value uint64) error {
	return e.Element.SetProperty("padding", value)
}

func (e *GstAsfMux) GetPadding() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The preroll time (milliseconds)
// Default: 5000, Min: 0, Max: 18446744073709551615
func (e *GstAsfMux) SetPreroll(value uint64) error {
	return e.Element.SetProperty("preroll", value)
}

func (e *GstAsfMux) GetPreroll() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("preroll")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written.
// Default: false
func (e *GstAsfMux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *GstAsfMux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parses ASF
type GstAsfParse struct {
	*GstBaseParse
}

func NewAsfParse() (*GstAsfParse, error) {
	e, err := gst.NewElement("asfparse")
	if err != nil {
		return nil, err
	}
	return &GstAsfParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewAsfParseWithName(name string) (*GstAsfParse, error) {
	e, err := gst.NewElementWithName("asfparse", name)
	if err != nil {
		return nil, err
	}
	return &GstAsfParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Payload-encodes ASF into RTP packets (MS_RTSP)
type GstRtpAsfPay struct {
	*GstRTPBasePayload
}

func NewRtpAsfPay() (*GstRtpAsfPay, error) {
	e, err := gst.NewElement("rtpasfpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpAsfPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpAsfPayWithName(name string) (*GstRtpAsfPay, error) {
	e, err := gst.NewElementWithName("rtpasfpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpAsfPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}
