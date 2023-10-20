// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Adds an alpha channel to video - uniform or via chroma-keying
type GstAlpha struct {
	*GstVideoFilter
}

func NewAlpha() (*GstAlpha, error) {
	e, err := gst.NewElement("alpha")
	if err != nil {
		return nil, err
	}
	return &GstAlpha{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAlphaWithName(name string) (*GstAlpha, error) {
	e, err := gst.NewElementWithName("alpha", name)
	if err != nil {
		return nil, err
	}
	return &GstAlpha{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Size of noise radius
// Default: 2, Min: 0, Max: 64
func (e *GstAlpha) SetNoiseLevel(value float32) error {
	return e.Element.SetProperty("noise-level", value)
}

func (e *GstAlpha) GetNoiseLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("noise-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Don't do any processing for alpha=1.0 if possible
// Default: false
func (e *GstAlpha) SetPreferPassthrough(value bool) error {
	return e.Element.SetProperty("prefer-passthrough", value)
}

func (e *GstAlpha) GetPreferPassthrough() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("prefer-passthrough")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The blue color value for custom RGB chroma keying
// Default: 0, Min: 0, Max: 255
func (e *GstAlpha) SetTargetB(value uint) error {
	return e.Element.SetProperty("target-b", value)
}

func (e *GstAlpha) GetTargetB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The green color value for custom RGB chroma keying
// Default: 255, Min: 0, Max: 255
func (e *GstAlpha) SetTargetG(value uint) error {
	return e.Element.SetProperty("target-g", value)
}

func (e *GstAlpha) GetTargetG() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The red color value for custom RGB chroma keying
// Default: 0, Min: 0, Max: 255
func (e *GstAlpha) SetTargetR(value uint) error {
	return e.Element.SetProperty("target-r", value)
}

func (e *GstAlpha) GetTargetR() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Sensitivity to bright colors
// Default: 100, Min: 0, Max: 128
func (e *GstAlpha) SetWhiteSensitivity(value uint) error {
	return e.Element.SetProperty("white-sensitivity", value)
}

func (e *GstAlpha) GetWhiteSensitivity() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("white-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Sensitivity to dark colors
// Default: 100, Min: 0, Max: 128
func (e *GstAlpha) SetBlackSensitivity(value uint) error {
	return e.Element.SetProperty("black-sensitivity", value)
}

func (e *GstAlpha) GetBlackSensitivity() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("black-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Size of the colorcube to change
// Default: 20, Min: 0, Max: 90
func (e *GstAlpha) SetAngle(value float32) error {
	return e.Element.SetProperty("angle", value)
}

func (e *GstAlpha) GetAngle() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// How the alpha channels should be created
// Default: set (0)
func (e *GstAlpha) SetMethod(value GstAlphaMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstAlpha) GetMethod() (GstAlphaMethod, error) {
	var value GstAlphaMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAlphaMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAlphaMethod")
	}
	return value, nil
}

// The value for the alpha channel
// Default: 1, Min: 0, Max: 1
func (e *GstAlpha) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *GstAlpha) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

type GstAlphaMethod string

const (
	GstAlphaMethodSet    GstAlphaMethod = "set"    // Set/adjust alpha channel
	GstAlphaMethodGreen  GstAlphaMethod = "green"  // Chroma Key on pure green
	GstAlphaMethodBlue   GstAlphaMethod = "blue"   // Chroma Key on pure blue
	GstAlphaMethodCustom GstAlphaMethod = "custom" // Chroma Key on custom RGB values
)
