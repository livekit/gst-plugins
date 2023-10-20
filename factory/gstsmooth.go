// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Apply a smooth filter to an image
type GstSmooth struct {
	*GstVideoFilter
}

func NewSmooth() (*GstSmooth, error) {
	e, err := gst.NewElement("smooth")
	if err != nil {
		return nil, err
	}
	return &GstSmooth{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewSmoothWithName(name string) (*GstSmooth, error) {
	e, err := gst.NewElementWithName("smooth", name)
	if err != nil {
		return nil, err
	}
	return &GstSmooth{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// process video
// Default: true
func (e *GstSmooth) SetActive(value bool) error {
	return e.Element.SetProperty("active", value)
}

func (e *GstSmooth) GetActive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("active")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// size of media filter
// Default: 3, Min: -2147483648, Max: 2147483647
func (e *GstSmooth) SetFilterSize(value int) error {
	return e.Element.SetProperty("filter-size", value)
}

func (e *GstSmooth) GetFilterSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("filter-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// only filter luma part
// Default: true
func (e *GstSmooth) SetLumaOnly(value bool) error {
	return e.Element.SetProperty("luma-only", value)
}

func (e *GstSmooth) GetLumaOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("luma-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// contrast tolerance for smoothing
// Default: 8, Min: -2147483648, Max: 2147483647
func (e *GstSmooth) SetTolerance(value int) error {
	return e.Element.SetProperty("tolerance", value)
}

func (e *GstSmooth) GetTolerance() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
