// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Deprecated by compositor. Mix multiple video streams
type GstVideoMixer2 struct {
	*gst.Element
}

func NewVideoMixer2() (*GstVideoMixer2, error) {
	e, err := gst.NewElement("videomixer")
	if err != nil {
		return nil, err
	}
	return &GstVideoMixer2{Element: e}, nil
}

func NewVideoMixer2WithName(name string) (*GstVideoMixer2, error) {
	e, err := gst.NewElementWithName("videomixer", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoMixer2{Element: e}, nil
}

// Background type
// Default: checker (0)
func (e *GstVideoMixer2) SetBackground(value GstVideoMixer2Background) error {
	e.Element.SetArg("background", string(value))
	return nil
}

func (e *GstVideoMixer2) GetBackground() (GstVideoMixer2Background, error) {
	var value GstVideoMixer2Background
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoMixer2Background)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoMixer2Background")
	}
	return value, nil
}

type GstVideoMixer2Background string

const (
	GstVideoMixer2BackgroundChecker     GstVideoMixer2Background = "checker"     // Checker pattern
	GstVideoMixer2BackgroundBlack       GstVideoMixer2Background = "black"       // Black
	GstVideoMixer2BackgroundWhite       GstVideoMixer2Background = "white"       // White
	GstVideoMixer2BackgroundTransparent GstVideoMixer2Background = "transparent" // Transparent Background to enable further mixing
)
