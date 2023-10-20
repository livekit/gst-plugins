// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// libdc1394 based source for IIDC cameras
type GstDC1394Src struct {
	*base.GstPushSrc
}

func NewDC1394Src() (*GstDC1394Src, error) {
	e, err := gst.NewElement("dc1394src")
	if err != nil {
		return nil, err
	}
	return &GstDC1394Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDC1394SrcWithName(name string) (*GstDC1394Src, error) {
	e, err := gst.NewElementWithName("dc1394src", name)
	if err != nil {
		return nil, err
	}
	return &GstDC1394Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// The hexadecimal representation of the GUID of the camera (use first camera available if null)
// Default: NULL
func (e *GstDC1394Src) SetGuid(value string) error {
	return e.Element.SetProperty("guid", value)
}

func (e *GstDC1394Src) GetGuid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("guid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The ISO bandwidth in Mbps
// Default: 400 (400)
func (e *GstDC1394Src) SetIso(value GstDC1394ISOSpeed) error {
	e.Element.SetArg("iso", string(value))
	return nil
}

func (e *GstDC1394Src) GetIso() (GstDC1394ISOSpeed, error) {
	var value GstDC1394ISOSpeed
	var ok bool
	v, err := e.Element.GetProperty("iso")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDC1394ISOSpeed)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDC1394ISOSpeed")
	}
	return value, nil
}

// Number of buffers to output before sending EOS (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstDC1394Src) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstDC1394Src) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Run typefind before negotiating (deprecated, non-functional)
// Default: false
func (e *GstDC1394Src) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

func (e *GstDC1394Src) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The unit number of the camera (-1 if no unit number is used)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstDC1394Src) SetUnit(value int) error {
	return e.Element.SetProperty("unit", value)
}

func (e *GstDC1394Src) GetUnit() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("unit")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Size in bytes to read per buffer (-1 = default)
// Default: 4096, Min: 0, Max: -1
func (e *GstDC1394Src) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstDC1394Src) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The number of frames in the Direct Memory Access ring buffer
// Default: 10, Min: 1, Max: -1
func (e *GstDC1394Src) SetDma(value uint) error {
	return e.Element.SetProperty("dma", value)
}

func (e *GstDC1394Src) GetDma() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dma")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Apply current stream time to buffers
// Default: true
func (e *GstDC1394Src) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

func (e *GstDC1394Src) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstDC1394ISOSpeed string

const (
	GstDC1394ISOSpeed100  GstDC1394ISOSpeed = "100"  // DC1394 ISO speed 100
	GstDC1394ISOSpeed200  GstDC1394ISOSpeed = "200"  // DC1394 ISO speed 200
	GstDC1394ISOSpeed400  GstDC1394ISOSpeed = "400"  // DC1394 ISO speed 400
	GstDC1394ISOSpeed800  GstDC1394ISOSpeed = "800"  // DC1394 ISO speed 800
	GstDC1394ISOSpeed1600 GstDC1394ISOSpeed = "1600" // DC1394 ISO speed 1600
	GstDC1394ISOSpeed3200 GstDC1394ISOSpeed = "3200" // DC1394 ISO speed 3200
)
