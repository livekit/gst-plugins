// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Adds a shape wipe transition to a video stream
type GstShapeWipe struct {
	*gst.Element
}

func NewShapeWipe() (*GstShapeWipe, error) {
	e, err := gst.NewElement("shapewipe")
	if err != nil {
		return nil, err
	}
	return &GstShapeWipe{Element: e}, nil
}

func NewShapeWipeWithName(name string) (*GstShapeWipe, error) {
	e, err := gst.NewElementWithName("shapewipe", name)
	if err != nil {
		return nil, err
	}
	return &GstShapeWipe{Element: e}, nil
}

// Border of the mask
// Default: 0, Min: 0, Max: 1
func (e *GstShapeWipe) SetBorder(value float32) error {
	return e.Element.SetProperty("border", value)
}

func (e *GstShapeWipe) GetBorder() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("border")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Position of the mask
// Default: 0, Min: 0, Max: 1
func (e *GstShapeWipe) SetPosition(value float32) error {
	return e.Element.SetProperty("position", value)
}

func (e *GstShapeWipe) GetPosition() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}
