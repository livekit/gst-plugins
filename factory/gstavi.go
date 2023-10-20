// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Demultiplex an avi file into audio and video
type GstAviDemux struct {
	*gst.Element
}

func NewAviDemux() (*GstAviDemux, error) {
	e, err := gst.NewElement("avidemux")
	if err != nil {
		return nil, err
	}
	return &GstAviDemux{Element: e}, nil
}

func NewAviDemuxWithName(name string) (*GstAviDemux, error) {
	e, err := gst.NewElementWithName("avidemux", name)
	if err != nil {
		return nil, err
	}
	return &GstAviDemux{Element: e}, nil
}

// Muxes audio and video into an avi stream
type GstAviMux struct {
	*gst.Element
}

func NewAviMux() (*GstAviMux, error) {
	e, err := gst.NewElement("avimux")
	if err != nil {
		return nil, err
	}
	return &GstAviMux{Element: e}, nil
}

func NewAviMuxWithName(name string) (*GstAviMux, error) {
	e, err := gst.NewElementWithName("avimux", name)
	if err != nil {
		return nil, err
	}
	return &GstAviMux{Element: e}, nil
}

// Support for openDML-2.0 (big) AVI files
// Default: true
func (e *GstAviMux) SetBigfile(value bool) error {
	return e.Element.SetProperty("bigfile", value)
}

func (e *GstAviMux) GetBigfile() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bigfile")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parse avi subtitle stream
type GstAviSubtitle struct {
	*gst.Element
}

func NewAviSubtitle() (*GstAviSubtitle, error) {
	e, err := gst.NewElement("avisubtitle")
	if err != nil {
		return nil, err
	}
	return &GstAviSubtitle{Element: e}, nil
}

func NewAviSubtitleWithName(name string) (*GstAviSubtitle, error) {
	e, err := gst.NewElementWithName("avisubtitle", name)
	if err != nil {
		return nil, err
	}
	return &GstAviSubtitle{Element: e}, nil
}
