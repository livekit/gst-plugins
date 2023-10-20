// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Resizes a video by adding borders or cropping
type GstVideoBox struct {
	*GstVideoFilter
}

func NewVideoBox() (*GstVideoBox, error) {
	e, err := gst.NewElement("videobox")
	if err != nil {
		return nil, err
	}
	return &GstVideoBox{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoBoxWithName(name string) (*GstVideoBox, error) {
	e, err := gst.NewElementWithName("videobox", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoBox{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Alpha value of the border
// Default: 1, Min: 0, Max: 1
func (e *GstVideoBox) SetBorderAlpha(value float64) error {
	return e.Element.SetProperty("border-alpha", value)
}

func (e *GstVideoBox) GetBorderAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("border-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pixels to box at bottom (<0 = add a border)
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoBox) SetBottom(value int) error {
	return e.Element.SetProperty("bottom", value)
}

func (e *GstVideoBox) GetBottom() (int, error) {
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

// How to fill the borders
// Default: black (0)
func (e *GstVideoBox) SetFill(value GstVideoBoxFill) error {
	e.Element.SetArg("fill", string(value))
	return nil
}

func (e *GstVideoBox) GetFill() (GstVideoBoxFill, error) {
	var value GstVideoBoxFill
	var ok bool
	v, err := e.Element.GetProperty("fill")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoBoxFill)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoBoxFill")
	}
	return value, nil
}

// Pixels to box at left (<0  = add a border)
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoBox) SetLeft(value int) error {
	return e.Element.SetProperty("left", value)
}

func (e *GstVideoBox) GetLeft() (int, error) {
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

// Pixels to box at right (<0 = add a border)
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoBox) SetRight(value int) error {
	return e.Element.SetProperty("right", value)
}

func (e *GstVideoBox) GetRight() (int, error) {
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

// Pixels to box at top (<0 = add a border)
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoBox) SetTop(value int) error {
	return e.Element.SetProperty("top", value)
}

func (e *GstVideoBox) GetTop() (int, error) {
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

// Alpha value picture
// Default: 1, Min: 0, Max: 1
func (e *GstVideoBox) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *GstVideoBox) GetAlpha() (float64, error) {
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

// Auto crop
// Default: false
func (e *GstVideoBox) SetAutocrop(value bool) error {
	return e.Element.SetProperty("autocrop", value)
}

func (e *GstVideoBox) GetAutocrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("autocrop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstVideoBoxFill string

const (
	GstVideoBoxFillBlack  GstVideoBoxFill = "black"  // Black
	GstVideoBoxFillGreen  GstVideoBoxFill = "green"  // Green
	GstVideoBoxFillBlue   GstVideoBoxFill = "blue"   // Blue
	GstVideoBoxFillRed    GstVideoBoxFill = "red"    // Red
	GstVideoBoxFillYellow GstVideoBoxFill = "yellow" // Yellow
	GstVideoBoxFillWhite  GstVideoBoxFill = "white"  // White
)
