// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decode VmWare video to raw (RGB) video
type GstVMncDec struct {
	*GstVideoDecoder
}

func NewVMncDec() (*GstVMncDec, error) {
	e, err := gst.NewElement("vmncdec")
	if err != nil {
		return nil, err
	}
	return &GstVMncDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewVMncDecWithName(name string) (*GstVMncDec, error) {
	e, err := gst.NewElementWithName("vmncdec", name)
	if err != nil {
		return nil, err
	}
	return &GstVMncDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}
