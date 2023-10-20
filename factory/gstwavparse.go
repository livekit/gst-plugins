// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Parse a .wav file into raw audio
type GstWavParse struct {
	*gst.Element
}

func NewWavParse() (*GstWavParse, error) {
	e, err := gst.NewElement("wavparse")
	if err != nil {
		return nil, err
	}
	return &GstWavParse{Element: e}, nil
}

func NewWavParseWithName(name string) (*GstWavParse, error) {
	e, err := gst.NewElementWithName("wavparse", name)
	if err != nil {
		return nil, err
	}
	return &GstWavParse{Element: e}, nil
}

// Ignore length from the Wave header
// Default: false
func (e *GstWavParse) SetIgnoreLength(value bool) error {
	return e.Element.SetProperty("ignore-length", value)
}

func (e *GstWavParse) GetIgnoreLength() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
