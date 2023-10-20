// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Musepack decoder
type GstMusepackDec struct {
	*gst.Element
}

func NewMusepackDec() (*GstMusepackDec, error) {
	e, err := gst.NewElement("musepackdec")
	if err != nil {
		return nil, err
	}
	return &GstMusepackDec{Element: e}, nil
}

func NewMusepackDecWithName(name string) (*GstMusepackDec, error) {
	e, err := gst.NewElementWithName("musepackdec", name)
	if err != nil {
		return nil, err
	}
	return &GstMusepackDec{Element: e}, nil
}
