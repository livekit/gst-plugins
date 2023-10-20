// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Read and output ICY tags while demuxing the contents
type GstICYDemux struct {
	*gst.Element
}

func NewICYDemux() (*GstICYDemux, error) {
	e, err := gst.NewElement("icydemux")
	if err != nil {
		return nil, err
	}
	return &GstICYDemux{Element: e}, nil
}

func NewICYDemuxWithName(name string) (*GstICYDemux, error) {
	e, err := gst.NewElementWithName("icydemux", name)
	if err != nil {
		return nil, err
	}
	return &GstICYDemux{Element: e}, nil
}
