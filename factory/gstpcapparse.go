// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Parses a raw interleaved RTSP stream
type GstIRTSPParse struct {
	*GstBaseParse
}

func NewIRTSPParse() (*GstIRTSPParse, error) {
	e, err := gst.NewElement("irtspparse")
	if err != nil {
		return nil, err
	}
	return &GstIRTSPParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewIRTSPParseWithName(name string) (*GstIRTSPParse, error) {
	e, err := gst.NewElementWithName("irtspparse", name)
	if err != nil {
		return nil, err
	}
	return &GstIRTSPParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Channel Identifier
// Default: 0, Min: 0, Max: 255
func (e *GstIRTSPParse) SetChannelId(value int) error {
	return e.Element.SetProperty("channel-id", value)
}

func (e *GstIRTSPParse) GetChannelId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("channel-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Parses a raw pcap stream
type GstPcapParse struct {
	*gst.Element
}

func NewPcapParse() (*GstPcapParse, error) {
	e, err := gst.NewElement("pcapparse")
	if err != nil {
		return nil, err
	}
	return &GstPcapParse{Element: e}, nil
}

func NewPcapParseWithName(name string) (*GstPcapParse, error) {
	e, err := gst.NewElementWithName("pcapparse", name)
	if err != nil {
		return nil, err
	}
	return &GstPcapParse{Element: e}, nil
}

// The caps of the source pad

func (e *GstPcapParse) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstPcapParse) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Destination IP to restrict to

func (e *GstPcapParse) SetDstIp(value string) error {
	return e.Element.SetProperty("dst-ip", value)
}

func (e *GstPcapParse) GetDstIp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dst-ip")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Destination port to restrict to
// Default: -1, Min: -1, Max: 65535
func (e *GstPcapParse) SetDstPort(value int) error {
	return e.Element.SetProperty("dst-port", value)
}

func (e *GstPcapParse) GetDstPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("dst-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Source IP to restrict to

func (e *GstPcapParse) SetSrcIp(value string) error {
	return e.Element.SetProperty("src-ip", value)
}

func (e *GstPcapParse) GetSrcIp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("src-ip")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Source port to restrict to
// Default: -1, Min: -1, Max: 65535
func (e *GstPcapParse) SetSrcPort(value int) error {
	return e.Element.SetProperty("src-port", value)
}

func (e *GstPcapParse) GetSrcPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("src-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Relative timestamp offset (ns) to apply (-1 = use absolute packet time)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstPcapParse) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstPcapParse) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}
