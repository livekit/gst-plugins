// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Parses Sub-Picture command streams and renders the SPU overlay onto the video as it passes through
type GstDVDSpu struct {
	*gst.Element
}

func NewDVDSpu() (*GstDVDSpu, error) {
	e, err := gst.NewElement("dvdspu")
	if err != nil {
		return nil, err
	}
	return &GstDVDSpu{Element: e}, nil
}

func NewDVDSpuWithName(name string) (*GstDVDSpu, error) {
	e, err := gst.NewElementWithName("dvdspu", name)
	if err != nil {
		return nil, err
	}
	return &GstDVDSpu{Element: e}, nil
}
