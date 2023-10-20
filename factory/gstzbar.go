// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Detect bar codes in the video streams
type GstZBar struct {
	*GstVideoFilter
}

func NewZBar() (*GstZBar, error) {
	e, err := gst.NewElement("zbar")
	if err != nil {
		return nil, err
	}
	return &GstZBar{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewZBarWithName(name string) (*GstZBar, error) {
	e, err := gst.NewElementWithName("zbar", name)
	if err != nil {
		return nil, err
	}
	return &GstZBar{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Attach a frame dump to each barcode message
// Default: false
func (e *GstZBar) SetAttachFrame(value bool) error {
	return e.Element.SetProperty("attach-frame", value)
}

func (e *GstZBar) GetAttachFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("attach-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable or disable the inter-image result cache
// Default: false
func (e *GstZBar) SetCache(value bool) error {
	return e.Element.SetProperty("cache", value)
}

func (e *GstZBar) GetCache() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cache")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Post a barcode message for each detected code
// Default: true
func (e *GstZBar) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

func (e *GstZBar) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
