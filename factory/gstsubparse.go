// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Parses SSA subtitle streams
type GstSsaParse struct {
	*gst.Element
}

func NewSsaParse() (*GstSsaParse, error) {
	e, err := gst.NewElement("ssaparse")
	if err != nil {
		return nil, err
	}
	return &GstSsaParse{Element: e}, nil
}

func NewSsaParseWithName(name string) (*GstSsaParse, error) {
	e, err := gst.NewElementWithName("ssaparse", name)
	if err != nil {
		return nil, err
	}
	return &GstSsaParse{Element: e}, nil
}

// Parses subtitle (.sub) files into text streams
type GstSubParse struct {
	*gst.Element
}

func NewSubParse() (*GstSubParse, error) {
	e, err := gst.NewElement("subparse")
	if err != nil {
		return nil, err
	}
	return &GstSubParse{Element: e}, nil
}

func NewSubParseWithName(name string) (*GstSubParse, error) {
	e, err := gst.NewElementWithName("subparse", name)
	if err != nil {
		return nil, err
	}
	return &GstSubParse{Element: e}, nil
}

// Framerate of the video stream. This is needed by some subtitle formats to synchronize subtitles and video properly. If not set and the subtitle format requires it subtitles may be out of sync.
// Default: 24000/1001, Min: 0/1, Max: 2147483647/1
func (e *GstSubParse) SetVideoFps(value interface{}) error {
	return e.Element.SetProperty("video-fps", value)
}

func (e *GstSubParse) GetVideoFps() (interface{}, error) {
	return e.Element.GetProperty("video-fps")
}

// Encoding to assume if input subtitles are not in UTF-8 or any other Unicode encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstSubParse) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstSubParse) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
