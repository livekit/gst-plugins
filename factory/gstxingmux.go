// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Adds a Xing header to the beginning of a VBR MP3 file
type GstXingMux struct {
	*gst.Element
}

func NewXingMux() (*GstXingMux, error) {
	e, err := gst.NewElement("xingmux")
	if err != nil {
		return nil, err
	}
	return &GstXingMux{Element: e}, nil
}

func NewXingMuxWithName(name string) (*GstXingMux, error) {
	e, err := gst.NewElementWithName("xingmux", name)
	if err != nil {
		return nil, err
	}
	return &GstXingMux{Element: e}, nil
}
