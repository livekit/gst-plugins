// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Parses TTML subtitle files
type GstTtmlParse struct {
	*gst.Element
}

func NewTtmlParse() (*GstTtmlParse, error) {
	e, err := gst.NewElement("ttmlparse")
	if err != nil {
		return nil, err
	}
	return &GstTtmlParse{Element: e}, nil
}

func NewTtmlParseWithName(name string) (*GstTtmlParse, error) {
	e, err := gst.NewElementWithName("ttmlparse", name)
	if err != nil {
		return nil, err
	}
	return &GstTtmlParse{Element: e}, nil
}

// Renders timed-text subtitles on top of video buffers
type GstTtmlRender struct {
	*gst.Element
}

func NewTtmlRender() (*GstTtmlRender, error) {
	e, err := gst.NewElement("ttmlrender")
	if err != nil {
		return nil, err
	}
	return &GstTtmlRender{Element: e}, nil
}

func NewTtmlRenderWithName(name string) (*GstTtmlRender, error) {
	e, err := gst.NewElementWithName("ttmlrender", name)
	if err != nil {
		return nil, err
	}
	return &GstTtmlRender{Element: e}, nil
}
