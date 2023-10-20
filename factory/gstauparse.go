// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Parse an .au file into raw audio
type GstAuParse struct {
	*gst.Element
}

func NewAuParse() (*GstAuParse, error) {
	e, err := gst.NewElement("auparse")
	if err != nil {
		return nil, err
	}
	return &GstAuParse{Element: e}, nil
}

func NewAuParseWithName(name string) (*GstAuParse, error) {
	e, err := gst.NewElementWithName("auparse", name)
	if err != nil {
		return nil, err
	}
	return &GstAuParse{Element: e}, nil
}
