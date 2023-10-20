// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// AV1 to json element
type GstAV12json struct {
	*gst.Element
}

func NewAV12json() (*GstAV12json, error) {
	e, err := gst.NewElement("av12json")
	if err != nil {
		return nil, err
	}
	return &GstAV12json{Element: e}, nil
}

func NewAV12jsonWithName(name string) (*GstAV12json, error) {
	e, err := gst.NewElementWithName("av12json", name)
	if err != nil {
		return nil, err
	}
	return &GstAV12json{Element: e}, nil
}

// VP8 to json element
type GstVp82json struct {
	*gst.Element
}

func NewVp82json() (*GstVp82json, error) {
	e, err := gst.NewElement("vp82json")
	if err != nil {
		return nil, err
	}
	return &GstVp82json{Element: e}, nil
}

func NewVp82jsonWithName(name string) (*GstVp82json, error) {
	e, err := gst.NewElementWithName("vp82json", name)
	if err != nil {
		return nil, err
	}
	return &GstVp82json{Element: e}, nil
}
