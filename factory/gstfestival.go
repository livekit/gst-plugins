// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Synthesizes plain text into audio
type GstFestival struct {
	*gst.Element
}

func NewFestival() (*GstFestival, error) {
	e, err := gst.NewElement("festival")
	if err != nil {
		return nil, err
	}
	return &GstFestival{Element: e}, nil
}

func NewFestivalWithName(name string) (*GstFestival, error) {
	e, err := gst.NewElementWithName("festival", name)
	if err != nil {
		return nil, err
	}
	return &GstFestival{Element: e}, nil
}
