// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Detect combing artifacts in video stream
type GstCombDetect struct {
	*GstVideoFilter
}

func NewCombDetect() (*GstCombDetect, error) {
	e, err := gst.NewElement("combdetect")
	if err != nil {
		return nil, err
	}
	return &GstCombDetect{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewCombDetectWithName(name string) (*GstCombDetect, error) {
	e, err := gst.NewElementWithName("combdetect", name)
	if err != nil {
		return nil, err
	}
	return &GstCombDetect{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Inverse Telecine Filter
type GstIvtc struct {
	*base.GstBaseTransform
}

func NewIvtc() (*GstIvtc, error) {
	e, err := gst.NewElement("ivtc")
	if err != nil {
		return nil, err
	}
	return &GstIvtc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewIvtcWithName(name string) (*GstIvtc, error) {
	e, err := gst.NewElementWithName("ivtc", name)
	if err != nil {
		return nil, err
	}
	return &GstIvtc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}
