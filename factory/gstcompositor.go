// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Composite multiple video streams
type GstCompositor struct {
	*GstVideoAggregator
}

func NewCompositor() (*GstCompositor, error) {
	e, err := gst.NewElement("compositor")
	if err != nil {
		return nil, err
	}
	return &GstCompositor{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewCompositorWithName(name string) (*GstCompositor, error) {
	e, err := gst.NewElementWithName("compositor", name)
	if err != nil {
		return nil, err
	}
	return &GstCompositor{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// If TRUE, then input video is unscaled in that dimension if width or height is 0 (for backwards compatibility)
// Default: true
func (e *GstCompositor) SetZeroSizeIsUnscaled(value bool) error {
	return e.Element.SetProperty("zero-size-is-unscaled", value)
}

func (e *GstCompositor) GetZeroSizeIsUnscaled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-size-is-unscaled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Background type
// Default: checker (0)
func (e *GstCompositor) SetBackground(value GstCompositorBackground) error {
	e.Element.SetArg("background", string(value))
	return nil
}

func (e *GstCompositor) GetBackground() (GstCompositorBackground, error) {
	var value GstCompositorBackground
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCompositorBackground)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCompositorBackground")
	}
	return value, nil
}

// Avoid timing out waiting for inactive pads
// Default: false
func (e *GstCompositor) SetIgnoreInactivePads(value bool) error {
	return e.Element.SetProperty("ignore-inactive-pads", value)
}

func (e *GstCompositor) GetIgnoreInactivePads() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-inactive-pads")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum number of blending/rendering worker threads to spawn (0 = auto)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstCompositor) SetMaxThreads(value uint) error {
	return e.Element.SetProperty("max-threads", value)
}

func (e *GstCompositor) GetMaxThreads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstCompositorBackground string

const (
	GstCompositorBackgroundChecker     GstCompositorBackground = "checker"     // Checker pattern
	GstCompositorBackgroundBlack       GstCompositorBackground = "black"       // Black
	GstCompositorBackgroundWhite       GstCompositorBackground = "white"       // White
	GstCompositorBackgroundTransparent GstCompositorBackground = "transparent" // Transparent Background to enable further compositing
)

type GstCompositorOperator string

const (
	GstCompositorOperatorSource GstCompositorOperator = "source" // Source
	GstCompositorOperatorOver   GstCompositorOperator = "over"   // Over
	GstCompositorOperatorAdd    GstCompositorOperator = "add"    // Add
)

type GstCompositorSizingPolicy string

const (
	GstCompositorSizingPolicyNone            GstCompositorSizingPolicy = "none"              // None: Image is scaled to fill configured destination rectangle without padding or keeping the aspect ratio
	GstCompositorSizingPolicyKeepAspectRatio GstCompositorSizingPolicy = "keep-aspect-ratio" // Keep Aspect Ratio: Image is scaled to fit destination rectangle specified by GstCompositorPad:{xpos, ypos, width, height} with preserved aspect ratio. Resulting image will be centered in the destination rectangle with padding if necessary
)
