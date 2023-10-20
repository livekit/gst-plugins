// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Creates Audio/Video streams from a web page using WPE web engine
type GstWpeSrc struct {
	*gst.Bin
}

func NewWpeSrc() (*GstWpeSrc, error) {
	e, err := gst.NewElement("wpesrc")
	if err != nil {
		return nil, err
	}
	return &GstWpeSrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewWpeSrcWithName(name string) (*GstWpeSrc, error) {
	e, err := gst.NewElementWithName("wpesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstWpeSrc{Bin: &gst.Bin{Element: e}}, nil
}

// Whether to draw the WebView background
// Default: true
func (e *GstWpeSrc) SetDrawBackground(value bool) error {
	return e.Element.SetProperty("draw-background", value)
}

func (e *GstWpeSrc) GetDrawBackground() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-background")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The URL to display
// Default: NULL
func (e *GstWpeSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstWpeSrc) GetLocation() (string, error) {
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

// Creates a video stream from a WPE browser
type GstWpeVideoSrc struct {
	*GstGLBaseSrc
}

func NewWpeVideoSrc() (*GstWpeVideoSrc, error) {
	e, err := gst.NewElement("wpevideosrc")
	if err != nil {
		return nil, err
	}
	return &GstWpeVideoSrc{GstGLBaseSrc: &GstGLBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewWpeVideoSrcWithName(name string) (*GstWpeVideoSrc, error) {
	e, err := gst.NewElementWithName("wpevideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstWpeVideoSrc{GstGLBaseSrc: &GstGLBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// Whether to draw the WebView background
// Default: true
func (e *GstWpeVideoSrc) SetDrawBackground(value bool) error {
	return e.Element.SetProperty("draw-background", value)
}

func (e *GstWpeVideoSrc) GetDrawBackground() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-background")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The URL to display
// Default: NULL
func (e *GstWpeVideoSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstWpeVideoSrc) GetLocation() (string, error) {
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
