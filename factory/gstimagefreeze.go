// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Generates a still frame stream from an image
type GstImageFreeze struct {
	*gst.Element
}

func NewImageFreeze() (*GstImageFreeze, error) {
	e, err := gst.NewElement("imagefreeze")
	if err != nil {
		return nil, err
	}
	return &GstImageFreeze{Element: e}, nil
}

func NewImageFreezeWithName(name string) (*GstImageFreeze, error) {
	e, err := gst.NewElementWithName("imagefreeze", name)
	if err != nil {
		return nil, err
	}
	return &GstImageFreeze{Element: e}, nil
}

// Number of buffers to output before sending EOS (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstImageFreeze) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstImageFreeze) GetNumBuffers() (int, error) {
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

// Allow replacing the input buffer and always output the latest
// Default: false
func (e *GstImageFreeze) SetAllowReplace(value bool) error {
	return e.Element.SetProperty("allow-replace", value)
}

func (e *GstImageFreeze) GetAllowReplace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-replace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to output a live video stream
// Default: false
func (e *GstImageFreeze) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstImageFreeze) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
