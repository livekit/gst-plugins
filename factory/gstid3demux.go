// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Read and output ID3v1 and ID3v2 tags while demuxing the contents
type GstID3Demux struct {
	*GstTagDemux
}

func NewID3Demux() (*GstID3Demux, error) {
	e, err := gst.NewElement("id3demux")
	if err != nil {
		return nil, err
	}
	return &GstID3Demux{GstTagDemux: &GstTagDemux{Element: e}}, nil
}

func NewID3DemuxWithName(name string) (*GstID3Demux, error) {
	e, err := gst.NewElementWithName("id3demux", name)
	if err != nil {
		return nil, err
	}
	return &GstID3Demux{GstTagDemux: &GstTagDemux{Element: e}}, nil
}

// Prefer tags from ID3v1 tag at end of file when both ID3v1 and ID3v2 tags are present
// Default: false
func (e *GstID3Demux) SetPreferV1(value bool) error {
	return e.Element.SetProperty("prefer-v1", value)
}

func (e *GstID3Demux) GetPreferV1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("prefer-v1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
