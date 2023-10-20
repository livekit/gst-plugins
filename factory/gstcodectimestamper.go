// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Timestamp H.264 streams
type GstH264Timestamper struct {
	*GstCodecTimestamper
}

func NewH264Timestamper() (*GstH264Timestamper, error) {
	e, err := gst.NewElement("h264timestamper")
	if err != nil {
		return nil, err
	}
	return &GstH264Timestamper{GstCodecTimestamper: &GstCodecTimestamper{Element: e}}, nil
}

func NewH264TimestamperWithName(name string) (*GstH264Timestamper, error) {
	e, err := gst.NewElementWithName("h264timestamper", name)
	if err != nil {
		return nil, err
	}
	return &GstH264Timestamper{GstCodecTimestamper: &GstCodecTimestamper{Element: e}}, nil
}

// Timestamp H.265 streams
type GstH265Timestamper struct {
	*GstCodecTimestamper
}

func NewH265Timestamper() (*GstH265Timestamper, error) {
	e, err := gst.NewElement("h265timestamper")
	if err != nil {
		return nil, err
	}
	return &GstH265Timestamper{GstCodecTimestamper: &GstCodecTimestamper{Element: e}}, nil
}

func NewH265TimestamperWithName(name string) (*GstH265Timestamper, error) {
	e, err := gst.NewElementWithName("h265timestamper", name)
	if err != nil {
		return nil, err
	}
	return &GstH265Timestamper{GstCodecTimestamper: &GstCodecTimestamper{Element: e}}, nil
}

type GstCodecTimestamper struct {
	*gst.Element
}
