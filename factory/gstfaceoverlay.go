// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Overlays SVG graphics over a detected face in a video stream
type GstFaceOverlay struct {
	*gst.Bin
}

func NewFaceOverlay() (*GstFaceOverlay, error) {
	e, err := gst.NewElement("faceoverlay")
	if err != nil {
		return nil, err
	}
	return &GstFaceOverlay{Bin: &gst.Bin{Element: e}}, nil
}

func NewFaceOverlayWithName(name string) (*GstFaceOverlay, error) {
	e, err := gst.NewElementWithName("faceoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstFaceOverlay{Bin: &gst.Bin{Element: e}}, nil
}

// Specify image height relative to face height.
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *GstFaceOverlay) SetH(value float32) error {
	return e.Element.SetProperty("h", value)
}

func (e *GstFaceOverlay) GetH() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("h")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Location of SVG file to use for face overlay
// Default: NULL
func (e *GstFaceOverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstFaceOverlay) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Specify image width relative to face width.
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *GstFaceOverlay) SetW(value float32) error {
	return e.Element.SetProperty("w", value)
}

func (e *GstFaceOverlay) GetW() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("w")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Specify image x relative to detected face x.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstFaceOverlay) SetX(value float32) error {
	return e.Element.SetProperty("x", value)
}

func (e *GstFaceOverlay) GetX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Specify image y relative to detected face y.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstFaceOverlay) SetY(value float32) error {
	return e.Element.SetProperty("y", value)
}

func (e *GstFaceOverlay) GetY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}
