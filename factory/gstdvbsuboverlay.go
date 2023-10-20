// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Renders DVB subtitles
type GstDVBSubOverlay struct {
	*gst.Element
}

func NewDVBSubOverlay() (*GstDVBSubOverlay, error) {
	e, err := gst.NewElement("dvbsuboverlay")
	if err != nil {
		return nil, err
	}
	return &GstDVBSubOverlay{Element: e}, nil
}

func NewDVBSubOverlayWithName(name string) (*GstDVBSubOverlay, error) {
	e, err := gst.NewElementWithName("dvbsuboverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstDVBSubOverlay{Element: e}, nil
}

// Enable rendering of subtitles
// Default: true
func (e *GstDVBSubOverlay) SetEnable(value bool) error {
	return e.Element.SetProperty("enable", value)
}

func (e *GstDVBSubOverlay) GetEnable() (bool, error) {
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

// Assume PES-aligned subtitles and force end-of-display
// Default: false
func (e *GstDVBSubOverlay) SetForceEnd(value bool) error {
	return e.Element.SetProperty("force-end", value)
}

func (e *GstDVBSubOverlay) GetForceEnd() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-end")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Limit maximum display time of a subtitle page (0 - disabled, value in seconds)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDVBSubOverlay) SetMaxPageTimeout(value int) error {
	return e.Element.SetProperty("max-page-timeout", value)
}

func (e *GstDVBSubOverlay) GetMaxPageTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-page-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
