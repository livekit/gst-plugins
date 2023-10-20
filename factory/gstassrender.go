// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Renders ASS/SSA subtitles with libass
type GstAssRender struct {
	*gst.Element
}

func NewAssRender() (*GstAssRender, error) {
	e, err := gst.NewElement("assrender")
	if err != nil {
		return nil, err
	}
	return &GstAssRender{Element: e}, nil
}

func NewAssRenderWithName(name string) (*GstAssRender, error) {
	e, err := gst.NewElementWithName("assrender", name)
	if err != nil {
		return nil, err
	}
	return &GstAssRender{Element: e}, nil
}

// Enable rendering of subtitles
// Default: true
func (e *GstAssRender) SetEnable(value bool) error {
	return e.Element.SetProperty("enable", value)
}

func (e *GstAssRender) GetEnable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to wait for subtitles
// Default: false
func (e *GstAssRender) SetWaitText(value bool) error {
	return e.Element.SetProperty("wait-text", value)
}

func (e *GstAssRender) GetWaitText() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extract and use fonts embedded in the stream
// Default: true
func (e *GstAssRender) SetEmbeddedfonts(value bool) error {
	return e.Element.SetProperty("embeddedfonts", value)
}

func (e *GstAssRender) GetEmbeddedfonts() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("embeddedfonts")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
