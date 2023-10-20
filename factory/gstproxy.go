// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Proxy source for internal process communication
type GstProxySink struct {
	*gst.Element
}

func NewProxySink() (*GstProxySink, error) {
	e, err := gst.NewElement("proxysink")
	if err != nil {
		return nil, err
	}
	return &GstProxySink{Element: e}, nil
}

func NewProxySinkWithName(name string) (*GstProxySink, error) {
	e, err := gst.NewElementWithName("proxysink", name)
	if err != nil {
		return nil, err
	}
	return &GstProxySink{Element: e}, nil
}

// Proxy source for internal process communication
type GstProxySrc struct {
	*gst.Bin
}

func NewProxySrc() (*GstProxySrc, error) {
	e, err := gst.NewElement("proxysrc")
	if err != nil {
		return nil, err
	}
	return &GstProxySrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewProxySrcWithName(name string) (*GstProxySrc, error) {
	e, err := gst.NewElementWithName("proxysrc", name)
	if err != nil {
		return nil, err
	}
	return &GstProxySrc{Bin: &gst.Bin{Element: e}}, nil
}

// Matching proxysink

func (e *GstProxySrc) SetProxysink(value interface{}) error {
	return e.Element.SetProperty("proxysink", value)
}

func (e *GstProxySrc) GetProxysink() (interface{}, error) {
	return e.Element.GetProperty("proxysink")
}
