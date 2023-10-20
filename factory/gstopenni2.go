// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Extract readings from an OpenNI supported device (Kinect etc).
type GstOpenni2Src struct {
	*base.GstPushSrc
}

func NewOpenni2Src() (*GstOpenni2Src, error) {
	e, err := gst.NewElement("openni2src")
	if err != nil {
		return nil, err
	}
	return &GstOpenni2Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewOpenni2SrcWithName(name string) (*GstOpenni2Src, error) {
	e, err := gst.NewElementWithName("openni2src", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenni2Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Run typefind before negotiating (deprecated, non-functional)
// Default: false
func (e *GstOpenni2Src) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

func (e *GstOpenni2Src) GetTypefind() (bool, error) {
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

// Size in bytes to read per buffer (-1 = default)
// Default: 4096, Min: 0, Max: -1
func (e *GstOpenni2Src) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstOpenni2Src) GetBlocksize() (uint, error) {
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

// Apply current stream time to buffers
// Default: false
func (e *GstOpenni2Src) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

func (e *GstOpenni2Src) GetDoTimestamp() (bool, error) {
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

// Source uri, can be a file or a device.
// Default: NULL
func (e *GstOpenni2Src) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstOpenni2Src) GetLocation() (string, error) {
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

// Number of buffers to output before sending EOS (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstOpenni2Src) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstOpenni2Src) GetNumBuffers() (int, error) {
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

// Type of readings to get from the source
// Default: depth (0)
func (e *GstOpenni2Src) SetSourcetype(value GstOpenni2SrcSourcetype) error {
	e.Element.SetArg("sourcetype", string(value))
	return nil
}

func (e *GstOpenni2Src) GetSourcetype() (GstOpenni2SrcSourcetype, error) {
	var value GstOpenni2SrcSourcetype
	var ok bool
	v, err := e.Element.GetProperty("sourcetype")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenni2SrcSourcetype)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenni2SrcSourcetype")
	}
	return value, nil
}

type GstOpenni2SrcSourcetype string

const (
	GstOpenni2SrcSourcetypeDepth GstOpenni2SrcSourcetype = "depth" // Get depth readings
	GstOpenni2SrcSourcetypeColor GstOpenni2SrcSourcetype = "color" // Get color readings
	GstOpenni2SrcSourcetypeBoth  GstOpenni2SrcSourcetype = "both"  // Get color and depth (as alpha) readings - EXPERIMENTAL
)
