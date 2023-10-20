// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Read audio streams using libsndfile
type GstSFDec struct {
	*gst.Element
}

func NewSFDec() (*GstSFDec, error) {
	e, err := gst.NewElement("sfdec")
	if err != nil {
		return nil, err
	}
	return &GstSFDec{Element: e}, nil
}

func NewSFDecWithName(name string) (*GstSFDec, error) {
	e, err := gst.NewElementWithName("sfdec", name)
	if err != nil {
		return nil, err
	}
	return &GstSFDec{Element: e}, nil
}
