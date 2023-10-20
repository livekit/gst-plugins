// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Handle navigation events showing black squares following mouse pointer and touch points
type GstNavigationtest struct {
	*GstVideoFilter
}

func NewNavigationtest() (*GstNavigationtest, error) {
	e, err := gst.NewElement("navigationtest")
	if err != nil {
		return nil, err
	}
	return &GstNavigationtest{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewNavigationtestWithName(name string) (*GstNavigationtest, error) {
	e, err := gst.NewElementWithName("navigationtest", name)
	if err != nil {
		return nil, err
	}
	return &GstNavigationtest{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Toggles display of mouse events
// Default: true
func (e *GstNavigationtest) SetDisplayMouse(value bool) error {
	return e.Element.SetProperty("display-mouse", value)
}

func (e *GstNavigationtest) GetDisplayMouse() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-mouse")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Toggles display of touchscreen events
// Default: true
func (e *GstNavigationtest) SetDisplayTouch(value bool) error {
	return e.Element.SetProperty("display-touch", value)
}

func (e *GstNavigationtest) GetDisplayTouch() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-touch")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
