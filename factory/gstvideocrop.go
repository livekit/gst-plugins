// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Crops video into a user-defined aspect-ratio
type GstAspectRatioCrop struct {
	*gst.Bin
}

func NewAspectRatioCrop() (*GstAspectRatioCrop, error) {
	e, err := gst.NewElement("aspectratiocrop")
	if err != nil {
		return nil, err
	}
	return &GstAspectRatioCrop{Bin: &gst.Bin{Element: e}}, nil
}

func NewAspectRatioCropWithName(name string) (*GstAspectRatioCrop, error) {
	e, err := gst.NewElementWithName("aspectratiocrop", name)
	if err != nil {
		return nil, err
	}
	return &GstAspectRatioCrop{Bin: &gst.Bin{Element: e}}, nil
}

// Target aspect-ratio of video
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstAspectRatioCrop) SetAspectRatio(value interface{}) error {
	return e.Element.SetProperty("aspect-ratio", value)
}

func (e *GstAspectRatioCrop) GetAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("aspect-ratio")
}

// Crops video into a user-defined region
type GstVideoCrop struct {
	*GstVideoFilter
}

func NewVideoCrop() (*GstVideoCrop, error) {
	e, err := gst.NewElement("videocrop")
	if err != nil {
		return nil, err
	}
	return &GstVideoCrop{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoCropWithName(name string) (*GstVideoCrop, error) {
	e, err := gst.NewElementWithName("videocrop", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoCrop{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Pixels to crop at bottom (-1 to auto-crop)
// Default: 0, Min: -1, Max: 2147483647
func (e *GstVideoCrop) SetBottom(value int) error {
	return e.Element.SetProperty("bottom", value)
}

func (e *GstVideoCrop) GetBottom() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pixels to crop at left (-1 to auto-crop)
// Default: 0, Min: -1, Max: 2147483647
func (e *GstVideoCrop) SetLeft(value int) error {
	return e.Element.SetProperty("left", value)
}

func (e *GstVideoCrop) GetLeft() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("left")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pixels to crop at right (-1 to auto-crop)
// Default: 0, Min: -1, Max: 2147483647
func (e *GstVideoCrop) SetRight(value int) error {
	return e.Element.SetProperty("right", value)
}

func (e *GstVideoCrop) GetRight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("right")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pixels to crop at top (-1 to auto-crop)
// Default: 0, Min: -1, Max: 2147483647
func (e *GstVideoCrop) SetTop(value int) error {
	return e.Element.SetProperty("top", value)
}

func (e *GstVideoCrop) GetTop() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("top")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
