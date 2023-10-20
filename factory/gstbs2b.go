// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Improve headphone listening of stereo audio records using the bs2b library.
type GstBs2b struct {
	*GstAudioFilter
}

func NewBs2b() (*GstBs2b, error) {
	e, err := gst.NewElement("bs2b")
	if err != nil {
		return nil, err
	}
	return &GstBs2b{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewBs2bWithName(name string) (*GstBs2b, error) {
	e, err := gst.NewElementWithName("bs2b", name)
	if err != nil {
		return nil, err
	}
	return &GstBs2b{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Low-pass filter cut frequency (Hz)
// Default: 700, Min: 300, Max: 2000
func (e *GstBs2b) SetFcut(value int) error {
	return e.Element.SetProperty("fcut", value)
}

func (e *GstBs2b) GetFcut() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fcut")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Feed Level (dB/10)
// Default: 45, Min: 10, Max: 150
func (e *GstBs2b) SetFeed(value int) error {
	return e.Element.SetProperty("feed", value)
}

func (e *GstBs2b) GetFeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("feed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
