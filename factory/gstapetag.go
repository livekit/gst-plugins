// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Read and output APE tags while demuxing the contents
type GstApeDemux struct {
	*GstTagDemux
}

func NewApeDemux() (*GstApeDemux, error) {
	e, err := gst.NewElement("apedemux")
	if err != nil {
		return nil, err
	}
	return &GstApeDemux{GstTagDemux: &GstTagDemux{Element: e}}, nil
}

func NewApeDemuxWithName(name string) (*GstApeDemux, error) {
	e, err := gst.NewElementWithName("apedemux", name)
	if err != nil {
		return nil, err
	}
	return &GstApeDemux{GstTagDemux: &GstTagDemux{Element: e}}, nil
}
