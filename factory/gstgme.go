// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Uses libgme to emulate a gaming console sound processors
type GstGmeDec struct {
	*gst.Element
}

func NewGmeDec() (*GstGmeDec, error) {
	e, err := gst.NewElement("gmedec")
	if err != nil {
		return nil, err
	}
	return &GstGmeDec{Element: e}, nil
}

func NewGmeDecWithName(name string) (*GstGmeDec, error) {
	e, err := gst.NewElementWithName("gmedec", name)
	if err != nil {
		return nil, err
	}
	return &GstGmeDec{Element: e}, nil
}
