// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Exclusion exclodes the colors in the video signal.
type GstExclusion struct {
	*GstVideoFilter
}

func NewExclusion() (*GstExclusion, error) {
	e, err := gst.NewElement("exclusion")
	if err != nil {
		return nil, err
	}
	return &GstExclusion{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewExclusionWithName(name string) (*GstExclusion, error) {
	e, err := gst.NewElementWithName("exclusion", name)
	if err != nil {
		return nil, err
	}
	return &GstExclusion{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Exclusion factor parameter
// Default: 175, Min: 1, Max: 175
func (e *GstExclusion) SetFactor(value uint) error {
	return e.Element.SetProperty("factor", value)
}

func (e *GstExclusion) GetFactor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Perform Gaussian blur/sharpen on a video
type GstGaussianBlur struct {
	*GstVideoFilter
}

func NewGaussianBlur() (*GstGaussianBlur, error) {
	e, err := gst.NewElement("gaussianblur")
	if err != nil {
		return nil, err
	}
	return &GstGaussianBlur{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewGaussianBlurWithName(name string) (*GstGaussianBlur, error) {
	e, err := gst.NewElementWithName("gaussianblur", name)
	if err != nil {
		return nil, err
	}
	return &GstGaussianBlur{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Sigma value for gaussian blur (negative for sharpen)
// Default: 1.2, Min: -20, Max: 20
func (e *GstGaussianBlur) SetSigma(value float64) error {
	return e.Element.SetProperty("sigma", value)
}

func (e *GstGaussianBlur) GetSigma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sigma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Solarize tunable inverse in the video signal.
type GstSolarize struct {
	*GstVideoFilter
}

func NewSolarize() (*GstSolarize, error) {
	e, err := gst.NewElement("solarize")
	if err != nil {
		return nil, err
	}
	return &GstSolarize{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewSolarizeWithName(name string) (*GstSolarize, error) {
	e, err := gst.NewElementWithName("solarize", name)
	if err != nil {
		return nil, err
	}
	return &GstSolarize{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// End parameter
// Default: 185, Min: 0, Max: 256
func (e *GstSolarize) SetEnd(value uint) error {
	return e.Element.SetProperty("end", value)
}

func (e *GstSolarize) GetEnd() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("end")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Start parameter
// Default: 50, Min: 0, Max: 256
func (e *GstSolarize) SetStart(value uint) error {
	return e.Element.SetProperty("start", value)
}

func (e *GstSolarize) GetStart() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("start")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Threshold parameter
// Default: 127, Min: 0, Max: 256
func (e *GstSolarize) SetThreshold(value uint) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstSolarize) GetThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Burn adjusts the colors in the video signal.
type GstBurn struct {
	*GstVideoFilter
}

func NewBurn() (*GstBurn, error) {
	e, err := gst.NewElement("burn")
	if err != nil {
		return nil, err
	}
	return &GstBurn{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewBurnWithName(name string) (*GstBurn, error) {
	e, err := gst.NewElementWithName("burn", name)
	if err != nil {
		return nil, err
	}
	return &GstBurn{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Adjustment parameter
// Default: 175, Min: 0, Max: 256
func (e *GstBurn) SetAdjustment(value uint) error {
	return e.Element.SetProperty("adjustment", value)
}

func (e *GstBurn) GetAdjustment() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("adjustment")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Chromium breaks the colors of the video signal.
type GstChromium struct {
	*GstVideoFilter
}

func NewChromium() (*GstChromium, error) {
	e, err := gst.NewElement("chromium")
	if err != nil {
		return nil, err
	}
	return &GstChromium{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewChromiumWithName(name string) (*GstChromium, error) {
	e, err := gst.NewElementWithName("chromium", name)
	if err != nil {
		return nil, err
	}
	return &GstChromium{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// First edge parameter
// Default: 200, Min: 0, Max: 256
func (e *GstChromium) SetEdgeA(value uint) error {
	return e.Element.SetProperty("edge-a", value)
}

func (e *GstChromium) GetEdgeA() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("edge-a")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Second edge parameter
// Default: 1, Min: 0, Max: 256
func (e *GstChromium) SetEdgeB(value uint) error {
	return e.Element.SetProperty("edge-b", value)
}

func (e *GstChromium) GetEdgeB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("edge-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Dilate copies the brightest pixel around.
type GstDilate struct {
	*GstVideoFilter
}

func NewDilate() (*GstDilate, error) {
	e, err := gst.NewElement("dilate")
	if err != nil {
		return nil, err
	}
	return &GstDilate{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewDilateWithName(name string) (*GstDilate, error) {
	e, err := gst.NewElementWithName("dilate", name)
	if err != nil {
		return nil, err
	}
	return &GstDilate{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Erode parameter
// Default: false
func (e *GstDilate) SetErode(value bool) error {
	return e.Element.SetProperty("erode", value)
}

func (e *GstDilate) GetErode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("erode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Dodge saturates the colors in the video signal.
type GstDodge struct {
	*GstVideoFilter
}

func NewDodge() (*GstDodge, error) {
	e, err := gst.NewElement("dodge")
	if err != nil {
		return nil, err
	}
	return &GstDodge{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewDodgeWithName(name string) (*GstDodge, error) {
	e, err := gst.NewElementWithName("dodge", name)
	if err != nil {
		return nil, err
	}
	return &GstDodge{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}
