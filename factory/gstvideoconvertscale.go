// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Converts video from one colorspace to another
type GstVideoConvert struct {
	*GstVideoConvertScale
}

func NewVideoConvert() (*GstVideoConvert, error) {
	e, err := gst.NewElement("videoconvert")
	if err != nil {
		return nil, err
	}
	return &GstVideoConvert{GstVideoConvertScale: &GstVideoConvertScale{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewVideoConvertWithName(name string) (*GstVideoConvert, error) {
	e, err := gst.NewElementWithName("videoconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoConvert{GstVideoConvertScale: &GstVideoConvertScale{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Resizes video and converts from one colorspace to another
type GstVideoConvertScale struct {
	*GstVideoFilter
}

func NewVideoConvertScale() (*GstVideoConvertScale, error) {
	e, err := gst.NewElement("videoconvertscale")
	if err != nil {
		return nil, err
	}
	return &GstVideoConvertScale{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoConvertScaleWithName(name string) (*GstVideoConvertScale, error) {
	e, err := gst.NewElementWithName("videoconvertscale", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoConvertScale{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Chroma resampler method
// Default: linear (1)
func (e *GstVideoConvertScale) SetChromaResampler(value interface{}) error {
	return e.Element.SetProperty("chroma-resampler", value)
}

func (e *GstVideoConvertScale) GetChromaResampler() (interface{}, error) {
	return e.Element.GetProperty("chroma-resampler")
}

// A GstStructure describing the configuration that should be used. This configuration, if set, takes precedence over the other similar conversion properties.

func (e *GstVideoConvertScale) SetConverterConfig(value *gst.Structure) error {
	return e.Element.SetProperty("converter-config", value)
}

func (e *GstVideoConvertScale) GetConverterConfig() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("converter-config")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Matrix Conversion Mode
// Default: full (0)
func (e *GstVideoConvertScale) SetMatrixMode(value interface{}) error {
	return e.Element.SetProperty("matrix-mode", value)
}

func (e *GstVideoConvertScale) GetMatrixMode() (interface{}, error) {
	return e.Element.GetProperty("matrix-mode")
}

// Maximum number of threads to use
// Default: 1, Min: 0, Max: -1
func (e *GstVideoConvertScale) SetNThreads(value uint) error {
	return e.Element.SetProperty("n-threads", value)
}

func (e *GstVideoConvertScale) GetNThreads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("n-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Primaries Conversion Mode
// Default: none (0)
func (e *GstVideoConvertScale) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

func (e *GstVideoConvertScale) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

// Apply dithering while converting
// Default: bayer (4)
func (e *GstVideoConvertScale) SetDither(value interface{}) error {
	return e.Element.SetProperty("dither", value)
}

func (e *GstVideoConvertScale) GetDither() (interface{}, error) {
	return e.Element.GetProperty("dither")
}

// Sharpening
// Default: 0, Min: 0, Max: 1
func (e *GstVideoConvertScale) SetSharpen(value float64) error {
	return e.Element.SetProperty("sharpen", value)
}

func (e *GstVideoConvertScale) GetSharpen() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sharpen")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Alpha Value to use
// Default: 1, Min: 0, Max: 1
func (e *GstVideoConvertScale) SetAlphaValue(value float64) error {
	return e.Element.SetProperty("alpha-value", value)
}

func (e *GstVideoConvertScale) GetAlphaValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Gamma Conversion Mode
// Default: none (0)
func (e *GstVideoConvertScale) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

func (e *GstVideoConvertScale) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

// method
// Default: bilinear (1)
func (e *GstVideoConvertScale) SetMethod(value GstVideoScaleMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstVideoConvertScale) GetMethod() (GstVideoScaleMethod, error) {
	var value GstVideoScaleMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoScaleMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoScaleMethod")
	}
	return value, nil
}

// Add black borders if necessary to keep the display aspect ratio
// Default: true
func (e *GstVideoConvertScale) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

func (e *GstVideoConvertScale) GetAddBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Alpha Mode to use
// Default: copy (0)
func (e *GstVideoConvertScale) SetAlphaMode(value interface{}) error {
	return e.Element.SetProperty("alpha-mode", value)
}

func (e *GstVideoConvertScale) GetAlphaMode() (interface{}, error) {
	return e.Element.GetProperty("alpha-mode")
}

// Chroma Resampling Mode
// Default: full (0)
func (e *GstVideoConvertScale) SetChromaMode(value interface{}) error {
	return e.Element.SetProperty("chroma-mode", value)
}

func (e *GstVideoConvertScale) GetChromaMode() (interface{}, error) {
	return e.Element.GetProperty("chroma-mode")
}

// Quantizer to use
// Default: 1, Min: 0, Max: -1
func (e *GstVideoConvertScale) SetDitherQuantization(value uint) error {
	return e.Element.SetProperty("dither-quantization", value)
}

func (e *GstVideoConvertScale) GetDitherQuantization() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dither-quantization")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Size of filter envelope
// Default: 2, Min: 1, Max: 5
func (e *GstVideoConvertScale) SetEnvelope(value float64) error {
	return e.Element.SetProperty("envelope", value)
}

func (e *GstVideoConvertScale) GetEnvelope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("envelope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Sharpness of filter
// Default: 1, Min: 0.5, Max: 1.5
func (e *GstVideoConvertScale) SetSharpness(value float64) error {
	return e.Element.SetProperty("sharpness", value)
}

func (e *GstVideoConvertScale) GetSharpness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sharpness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Resizes video
type GstVideoScale struct {
	*GstVideoConvertScale
}

func NewVideoScale() (*GstVideoScale, error) {
	e, err := gst.NewElement("videoscale")
	if err != nil {
		return nil, err
	}
	return &GstVideoScale{GstVideoConvertScale: &GstVideoConvertScale{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewVideoScaleWithName(name string) (*GstVideoScale, error) {
	e, err := gst.NewElementWithName("videoscale", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoScale{GstVideoConvertScale: &GstVideoConvertScale{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Decode gamma before scaling
// Default: false
func (e *GstVideoScale) SetGammaDecode(value bool) error {
	return e.Element.SetProperty("gamma-decode", value)
}

func (e *GstVideoScale) GetGammaDecode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gamma-decode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstVideoScaleMethod string

const (
	GstVideoScaleMethodNearestNeighbour GstVideoScaleMethod = "nearest-neighbour" // Nearest Neighbour
	GstVideoScaleMethodBilinear         GstVideoScaleMethod = "bilinear"          // Bilinear (2-tap)
	GstVideoScaleMethod4Tap             GstVideoScaleMethod = "4-tap"             // 4-tap Sinc
	GstVideoScaleMethodLanczos          GstVideoScaleMethod = "lanczos"           // Lanczos
	GstVideoScaleMethodBilinear2        GstVideoScaleMethod = "bilinear2"         // Bilinear (multi-tap)
	GstVideoScaleMethodSinc             GstVideoScaleMethod = "sinc"              // Sinc (multi-tap)
	GstVideoScaleMethodHermite          GstVideoScaleMethod = "hermite"           // Hermite (multi-tap)
	GstVideoScaleMethodSpline           GstVideoScaleMethod = "spline"            // Spline (multi-tap)
	GstVideoScaleMethodCatrom           GstVideoScaleMethod = "catrom"            // Catmull-Rom (multi-tap)
	GstVideoScaleMethodMitchell         GstVideoScaleMethod = "mitchell"          // Mitchell (multi-tap)
)
