// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes DTS audio streams
type GstDtsDec struct {
	*GstAudioDecoder
}

func NewDtsDec() (*GstDtsDec, error) {
	e, err := gst.NewElement("dtsdec")
	if err != nil {
		return nil, err
	}
	return &GstDtsDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewDtsDecWithName(name string) (*GstDtsDec, error) {
	e, err := gst.NewElementWithName("dtsdec", name)
	if err != nil {
		return nil, err
	}
	return &GstDtsDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Use Dynamic Range Compression
// Default: false
func (e *GstDtsDec) SetDrc(value bool) error {
	return e.Element.SetProperty("drc", value)
}

func (e *GstDtsDec) GetDrc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
