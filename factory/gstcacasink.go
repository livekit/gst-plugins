// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A colored ASCII art videosink
type GstCACASink struct {
	*base.GstBaseSink
}

func NewCACASink() (*GstCACASink, error) {
	e, err := gst.NewElement("cacasink")
	if err != nil {
		return nil, err
	}
	return &GstCACASink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewCACASinkWithName(name string) (*GstCACASink, error) {
	e, err := gst.NewElementWithName("cacasink", name)
	if err != nil {
		return nil, err
	}
	return &GstCACASink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Enables Anti-Aliasing
// Default: true
func (e *GstCACASink) SetAntiAliasing(value bool) error {
	return e.Element.SetProperty("anti-aliasing", value)
}

func (e *GstCACASink) GetAntiAliasing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("anti-aliasing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set type of Dither
// Default: none (49)
func (e *GstCACASink) SetDither(value GstCACASinkDithering) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

func (e *GstCACASink) GetDither() (GstCACASinkDithering, error) {
	var value GstCACASinkDithering
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCACASinkDithering)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCACASinkDithering")
	}
	return value, nil
}

// The height of the screen
// Default: 25, Min: 0, Max: 2147483647
func (e *GstCACASink) GetScreenHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("screen-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The width of the screen
// Default: 80, Min: 0, Max: 2147483647
func (e *GstCACASink) GetScreenWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("screen-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Colored ASCII art effect
type GstCACATv struct {
	*GstVideoFilter
}

func NewCACATv() (*GstCACATv, error) {
	e, err := gst.NewElement("cacatv")
	if err != nil {
		return nil, err
	}
	return &GstCACATv{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewCACATvWithName(name string) (*GstCACATv, error) {
	e, err := gst.NewElementWithName("cacatv", name)
	if err != nil {
		return nil, err
	}
	return &GstCACATv{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// selected libcaca font
// Default: 0, Min: 0, Max: 2147483647
func (e *GstCACATv) SetFont(value int) error {
	return e.Element.SetProperty("font", value)
}

func (e *GstCACATv) GetFont() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("font")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Enables Anti-Aliasing
// Default: false
func (e *GstCACATv) SetAntiAliasing(value bool) error {
	return e.Element.SetProperty("anti-aliasing", value)
}

func (e *GstCACATv) GetAntiAliasing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("anti-aliasing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The height of the canvas in characters
// Default: 24, Min: 0, Max: 2147483647
func (e *GstCACATv) SetCanvasHeight(value int) error {
	return e.Element.SetProperty("canvas-height", value)
}

func (e *GstCACATv) GetCanvasHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("canvas-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The width of the canvas in characters
// Default: 80, Min: 0, Max: 2147483647
func (e *GstCACATv) SetCanvasWidth(value int) error {
	return e.Element.SetProperty("canvas-width", value)
}

func (e *GstCACATv) GetCanvasWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("canvas-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Set type of Dither
// Default:  (0)
func (e *GstCACATv) SetDither(value GstCACATvDithering) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

func (e *GstCACATv) GetDither() (GstCACATvDithering, error) {
	var value GstCACATvDithering
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCACATvDithering)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCACATvDithering")
	}
	return value, nil
}

type GstCACASinkDithering string

const (
	GstCACASinkDitheringNone   GstCACASinkDithering = "none"   // No dithering
	GstCACASinkDithering2x2    GstCACASinkDithering = "2x2"    // Ordered 2x2 Bayer dithering
	GstCACASinkDithering4x4    GstCACASinkDithering = "4x4"    // Ordered 4x4 Bayer dithering
	GstCACASinkDithering8x8    GstCACASinkDithering = "8x8"    // Ordered 8x8 Bayer dithering
	GstCACASinkDitheringRandom GstCACASinkDithering = "random" // Random dithering
)

type GstCACATvDithering string

const (
	GstCACATvDitheringNone   GstCACATvDithering = "none"   // No dither_mode
	GstCACATvDithering2x2    GstCACATvDithering = "2x2"    // Ordered 2x2 Bayer dither_mode
	GstCACATvDithering4x4    GstCACATvDithering = "4x4"    // Ordered 4x4 Bayer dither_mode
	GstCACATvDithering8x8    GstCACATvDithering = "8x8"    // Ordered 8x8 Bayer dither_mode
	GstCACATvDitheringRandom GstCACATvDithering = "random" // Random dither_mode
)
