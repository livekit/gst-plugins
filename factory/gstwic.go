// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Jpeg image decoder using Windows Imaging Component API
type GstWicJpegDec struct {
	*GstWicDecoder
}

func NewWicJpegDec() (*GstWicJpegDec, error) {
	e, err := gst.NewElement("wicjpegdec")
	if err != nil {
		return nil, err
	}
	return &GstWicJpegDec{GstWicDecoder: &GstWicDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewWicJpegDecWithName(name string) (*GstWicJpegDec, error) {
	e, err := gst.NewElementWithName("wicjpegdec", name)
	if err != nil {
		return nil, err
	}
	return &GstWicJpegDec{GstWicDecoder: &GstWicDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Png image decoder using Windows Imaging Component API
type GstWicPngDec struct {
	*GstWicDecoder
}

func NewWicPngDec() (*GstWicPngDec, error) {
	e, err := gst.NewElement("wicpngdec")
	if err != nil {
		return nil, err
	}
	return &GstWicPngDec{GstWicDecoder: &GstWicDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewWicPngDecWithName(name string) (*GstWicPngDec, error) {
	e, err := gst.NewElementWithName("wicpngdec", name)
	if err != nil {
		return nil, err
	}
	return &GstWicPngDec{GstWicDecoder: &GstWicDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

type GstWicDecoder struct {
	*GstVideoDecoder
}
