// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Clips audio buffers to the configured segment
type GstAudioSegmentClip struct {
	*GstSegmentClip
}

func NewAudioSegmentClip() (*GstAudioSegmentClip, error) {
	e, err := gst.NewElement("audiosegmentclip")
	if err != nil {
		return nil, err
	}
	return &GstAudioSegmentClip{GstSegmentClip: &GstSegmentClip{Element: e}}, nil
}

func NewAudioSegmentClipWithName(name string) (*GstAudioSegmentClip, error) {
	e, err := gst.NewElementWithName("audiosegmentclip", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioSegmentClip{GstSegmentClip: &GstSegmentClip{Element: e}}, nil
}

// Clips video buffers to the configured segment
type GstVideoSegmentClip struct {
	*GstSegmentClip
}

func NewVideoSegmentClip() (*GstVideoSegmentClip, error) {
	e, err := gst.NewElement("videosegmentclip")
	if err != nil {
		return nil, err
	}
	return &GstVideoSegmentClip{GstSegmentClip: &GstSegmentClip{Element: e}}, nil
}

func NewVideoSegmentClipWithName(name string) (*GstVideoSegmentClip, error) {
	e, err := gst.NewElementWithName("videosegmentclip", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoSegmentClip{GstSegmentClip: &GstSegmentClip{Element: e}}, nil
}

type GstSegmentClip struct {
	*gst.Element
}
