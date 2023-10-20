// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Adjusts gamma on a video stream
type GstGamma struct {
	*GstVideoFilter
}

func NewGamma() (*GstGamma, error) {
	e, err := gst.NewElement("gamma")
	if err != nil {
		return nil, err
	}
	return &GstGamma{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewGammaWithName(name string) (*GstGamma, error) {
	e, err := gst.NewElementWithName("gamma", name)
	if err != nil {
		return nil, err
	}
	return &GstGamma{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// gamma
// Default: 1, Min: 0.01, Max: 10
func (e *GstGamma) SetGamma(value float64) error {
	return e.Element.SetProperty("gamma", value)
}

func (e *GstGamma) GetGamma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("gamma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adjusts brightness, contrast, hue, saturation on a video stream
type GstVideoBalance struct {
	*GstVideoFilter
}

func NewVideoBalance() (*GstVideoBalance, error) {
	e, err := gst.NewElement("videobalance")
	if err != nil {
		return nil, err
	}
	return &GstVideoBalance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoBalanceWithName(name string) (*GstVideoBalance, error) {
	e, err := gst.NewElementWithName("videobalance", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoBalance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// brightness
// Default: 0, Min: -1, Max: 1
func (e *GstVideoBalance) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstVideoBalance) GetBrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// contrast
// Default: 1, Min: 0, Max: 2
func (e *GstVideoBalance) SetContrast(value float64) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstVideoBalance) GetContrast() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// hue
// Default: 0, Min: -1, Max: 1
func (e *GstVideoBalance) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstVideoBalance) GetHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// saturation
// Default: 1, Min: 0, Max: 2
func (e *GstVideoBalance) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstVideoBalance) GetSaturation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Flips and rotates video
type GstVideoFlip struct {
	*GstVideoFilter
}

func NewVideoFlip() (*GstVideoFlip, error) {
	e, err := gst.NewElement("videoflip")
	if err != nil {
		return nil, err
	}
	return &GstVideoFlip{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoFlipWithName(name string) (*GstVideoFlip, error) {
	e, err := gst.NewElementWithName("videoflip", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoFlip{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// method (deprecated, use video-direction instead)
// Default: none (0)
func (e *GstVideoFlip) SetMethod(value GstVideoFlipMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstVideoFlip) GetMethod() (GstVideoFlipMethod, error) {
	var value GstVideoFlipMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoFlipMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoFlipMethod")
	}
	return value, nil
}

// Apply a median filter to an image
type GstVideoMedian struct {
	*GstVideoFilter
}

func NewVideoMedian() (*GstVideoMedian, error) {
	e, err := gst.NewElement("videomedian")
	if err != nil {
		return nil, err
	}
	return &GstVideoMedian{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoMedianWithName(name string) (*GstVideoMedian, error) {
	e, err := gst.NewElementWithName("videomedian", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoMedian{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The size of the filter
// Default: 5 (5)
func (e *GstVideoMedian) SetFiltersize(value GstVideoMedianSize) error {
	e.Element.SetArg("filtersize", string(value))
	return nil
}

func (e *GstVideoMedian) GetFiltersize() (GstVideoMedianSize, error) {
	var value GstVideoMedianSize
	var ok bool
	v, err := e.Element.GetProperty("filtersize")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoMedianSize)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoMedianSize")
	}
	return value, nil
}

// Only apply filter on luminance
// Default: true
func (e *GstVideoMedian) SetLumOnly(value bool) error {
	return e.Element.SetProperty("lum-only", value)
}

func (e *GstVideoMedian) GetLumOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lum-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstVideoMedianSize string

const (
	GstVideoMedianSize5 GstVideoMedianSize = "5" // Median of 5 neighbour pixels
	GstVideoMedianSize9 GstVideoMedianSize = "9" // Median of 9 neighbour pixels
)

type GstVideoFlipMethod string

const (
	GstVideoFlipMethodNone               GstVideoFlipMethod = "none"                 // Identity (no rotation)
	GstVideoFlipMethodClockwise          GstVideoFlipMethod = "clockwise"            // Rotate clockwise 90 degrees
	GstVideoFlipMethodRotate180          GstVideoFlipMethod = "rotate-180"           // Rotate 180 degrees
	GstVideoFlipMethodCounterclockwise   GstVideoFlipMethod = "counterclockwise"     // Rotate counter-clockwise 90 degrees
	GstVideoFlipMethodHorizontalFlip     GstVideoFlipMethod = "horizontal-flip"      // Flip horizontally
	GstVideoFlipMethodVerticalFlip       GstVideoFlipMethod = "vertical-flip"        // Flip vertically
	GstVideoFlipMethodUpperLeftDiagonal  GstVideoFlipMethod = "upper-left-diagonal"  // Flip across upper left/lower right diagonal
	GstVideoFlipMethodUpperRightDiagonal GstVideoFlipMethod = "upper-right-diagonal" // Flip across upper right/lower left diagonal
	GstVideoFlipMethodAutomatic          GstVideoFlipMethod = "automatic"            // Select flip method based on image-orientation tag
)
