// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Depayloads GStreamer Data Protocol buffers
type GstGDPDepay struct {
	*gst.Element
}

func NewGDPDepay() (*GstGDPDepay, error) {
	e, err := gst.NewElement("gdpdepay")
	if err != nil {
		return nil, err
	}
	return &GstGDPDepay{Element: e}, nil
}

func NewGDPDepayWithName(name string) (*GstGDPDepay, error) {
	e, err := gst.NewElementWithName("gdpdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstGDPDepay{Element: e}, nil
}

// Timestamp Offset
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstGDPDepay) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstGDPDepay) GetTsOffset() (int64, error) {
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

// Payloads GStreamer Data Protocol buffers
type GstGDPPay struct {
	*gst.Element
}

func NewGDPPay() (*GstGDPPay, error) {
	e, err := gst.NewElement("gdppay")
	if err != nil {
		return nil, err
	}
	return &GstGDPPay{Element: e}, nil
}

func NewGDPPayWithName(name string) (*GstGDPPay, error) {
	e, err := gst.NewElementWithName("gdppay", name)
	if err != nil {
		return nil, err
	}
	return &GstGDPPay{Element: e}, nil
}

// Calculate and store a CRC checksum on the header
// Default: true
func (e *GstGDPPay) SetCrcHeader(value bool) error {
	return e.Element.SetProperty("crc-header", value)
}

func (e *GstGDPPay) GetCrcHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("crc-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Calculate and store a CRC checksum on the payload
// Default: false
func (e *GstGDPPay) SetCrcPayload(value bool) error {
	return e.Element.SetProperty("crc-payload", value)
}

func (e *GstGDPPay) GetCrcPayload() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("crc-payload")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
