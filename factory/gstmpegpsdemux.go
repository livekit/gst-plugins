// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Demultiplexes MPEG Program Streams
type GstMpegPSDemux struct {
	*gst.Element
}

func NewMpegPSDemux() (*GstMpegPSDemux, error) {
	e, err := gst.NewElement("mpegpsdemux")
	if err != nil {
		return nil, err
	}
	return &GstMpegPSDemux{Element: e}, nil
}

func NewMpegPSDemuxWithName(name string) (*GstMpegPSDemux, error) {
	e, err := gst.NewElementWithName("mpegpsdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstMpegPSDemux{Element: e}, nil
}

// Ignore SCR data for timing
// Default: false
func (e *GstMpegPSDemux) SetIgnoreScr(value bool) error {
	return e.Element.SetProperty("ignore-scr", value)
}

func (e *GstMpegPSDemux) GetIgnoreScr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-scr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
