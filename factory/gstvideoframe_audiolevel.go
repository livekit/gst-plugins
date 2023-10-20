// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Synchronized audio/video RMS Level messenger for audio/raw
type GstVideoFrameAudioLevel struct {
	*gst.Element
}

func NewVideoFrameAudioLevel() (*GstVideoFrameAudioLevel, error) {
	e, err := gst.NewElement("videoframe-audiolevel")
	if err != nil {
		return nil, err
	}
	return &GstVideoFrameAudioLevel{Element: e}, nil
}

func NewVideoFrameAudioLevelWithName(name string) (*GstVideoFrameAudioLevel, error) {
	e, err := gst.NewElementWithName("videoframe-audiolevel", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoFrameAudioLevel{Element: e}, nil
}
