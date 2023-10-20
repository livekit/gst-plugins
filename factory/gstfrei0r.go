// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// RGB splitting and shifting
type Frei0rFilterRgbsplit0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterRgbsplit0r() (*Frei0rFilterRgbsplit0r, error) {
	e, err := gst.NewElement("frei0r-filter-rgbsplit0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterRgbsplit0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterRgbsplit0rWithName(name string) (*Frei0rFilterRgbsplit0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-rgbsplit0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterRgbsplit0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// How far should layers be moved horizontally from each other
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterRgbsplit0r) SetHorizontalSplitDistance(value float64) error {
	return e.Element.SetProperty("horizontal-split-distance", value)
}

func (e *Frei0rFilterRgbsplit0r) GetHorizontalSplitDistance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("horizontal-split-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// How far should layers be moved vertically from each other
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterRgbsplit0r) SetVerticalSplitDistance(value float64) error {
	return e.Element.SetProperty("vertical-split-distance", value)
}

func (e *Frei0rFilterRgbsplit0r) GetVerticalSplitDistance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("vertical-split-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Do simple color correction, in a physically meaningful way
type Frei0rFilterWhiteBalanceLmsSpace struct {
	*GstVideoFilter
}

func NewFrei0rFilterWhiteBalanceLmsSpace() (*Frei0rFilterWhiteBalanceLmsSpace, error) {
	e, err := gst.NewElement("frei0r-filter-white-balance--lms-space-")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterWhiteBalanceLmsSpace{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterWhiteBalanceLmsSpaceWithName(name string) (*Frei0rFilterWhiteBalanceLmsSpace, error) {
	e, err := gst.NewElementWithName("frei0r-filter-white-balance--lms-space-", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterWhiteBalanceLmsSpace{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Choose a color from the source image that should be white.
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalanceLmsSpace) SetNeutralColorB(value float32) error {
	return e.Element.SetProperty("neutral-color-b", value)
}

func (e *Frei0rFilterWhiteBalanceLmsSpace) GetNeutralColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Choose a color from the source image that should be white.
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalanceLmsSpace) SetNeutralColorG(value float32) error {
	return e.Element.SetProperty("neutral-color-g", value)
}

func (e *Frei0rFilterWhiteBalanceLmsSpace) GetNeutralColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Choose a color from the source image that should be white.
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalanceLmsSpace) SetNeutralColorR(value float32) error {
	return e.Element.SetProperty("neutral-color-r", value)
}

func (e *Frei0rFilterWhiteBalanceLmsSpace) GetNeutralColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Choose an output color temperature, if different from 6500 K.
// Default: 0.433333, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalanceLmsSpace) SetColorTemperature(value float64) error {
	return e.Element.SetProperty("color-temperature", value)
}

func (e *Frei0rFilterWhiteBalanceLmsSpace) GetColorTemperature() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color-temperature")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// the alpha IN operation
type Frei0rMixerAlphain struct {
	*gst.Element
}

func NewFrei0rMixerAlphain() (*Frei0rMixerAlphain, error) {
	e, err := gst.NewElement("frei0r-mixer-alphain")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphain{Element: e}, nil
}

func NewFrei0rMixerAlphainWithName(name string) (*Frei0rMixerAlphain, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-alphain", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphain{Element: e}, nil
}

// Composites second input on the first input with user-defined blend mode and opacity.
type Frei0rMixerCairoblend struct {
	*gst.Element
}

func NewFrei0rMixerCairoblend() (*Frei0rMixerCairoblend, error) {
	e, err := gst.NewElement("frei0r-mixer-cairoblend")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerCairoblend{Element: e}, nil
}

func NewFrei0rMixerCairoblendWithName(name string) (*Frei0rMixerCairoblend, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-cairoblend", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerCairoblend{Element: e}, nil
}

// Blend mode used to compose image. Accepted values: 'normal', 'add', 'saturate', 'multiply', 'screen', 'overlay', 'darken', 'lighten', 'colordodge', 'colorburn', 'hardlight', 'softlight', 'difference', 'exclusion', 'hslhue', 'hslsaturation', 'hslcolor', 'hslluminosity'
// Default: normal
func (e *Frei0rMixerCairoblend) SetBlendMode(value string) error {
	return e.Element.SetProperty("blend-mode", value)
}

func (e *Frei0rMixerCairoblend) GetBlendMode() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("blend-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Opacity of composited image
// Default: 1, Min: 0, Max: 1
func (e *Frei0rMixerCairoblend) SetOpacity(value float64) error {
	return e.Element.SetProperty("opacity", value)
}

func (e *Frei0rMixerCairoblend) GetOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform a conversion to value only of the source input1 using the value of input2.
type Frei0rMixerValue struct {
	*gst.Element
}

func NewFrei0rMixerValue() (*Frei0rMixerValue, error) {
	e, err := gst.NewElement("frei0r-mixer-value")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerValue{Element: e}, nil
}

func NewFrei0rMixerValueWithName(name string) (*Frei0rMixerValue, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-value", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerValue{Element: e}, nil
}

// Clusters of a source image by color and spatial distance
type Frei0rFilterKMeansClustering struct {
	*GstVideoFilter
}

func NewFrei0rFilterKMeansClustering() (*Frei0rFilterKMeansClustering, error) {
	e, err := gst.NewElement("frei0r-filter-k-means-clustering")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterKMeansClustering{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterKMeansClusteringWithName(name string) (*Frei0rFilterKMeansClustering, error) {
	e, err := gst.NewElementWithName("frei0r-filter-k-means-clustering", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterKMeansClustering{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The weight on distance
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterKMeansClustering) SetDistWeight(value float64) error {
	return e.Element.SetProperty("dist-weight", value)
}

func (e *Frei0rFilterKMeansClustering) GetDistWeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("dist-weight")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The number of clusters
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterKMeansClustering) SetNum(value float64) error {
	return e.Element.SetProperty("num", value)
}

func (e *Frei0rFilterKMeansClustering) GetNum() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("num")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Creates embossed relief image of source image
type Frei0rFilterEmboss struct {
	*GstVideoFilter
}

func NewFrei0rFilterEmboss() (*Frei0rFilterEmboss, error) {
	e, err := gst.NewElement("frei0r-filter-emboss")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterEmboss{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterEmbossWithName(name string) (*Frei0rFilterEmboss, error) {
	e, err := gst.NewElementWithName("frei0r-filter-emboss", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterEmboss{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Light direction
// Default: 0.375, Min: 0, Max: 1
func (e *Frei0rFilterEmboss) SetAzimuth(value float64) error {
	return e.Element.SetProperty("azimuth", value)
}

func (e *Frei0rFilterEmboss) GetAzimuth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("azimuth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Background lightness
// Default: 0.333333, Min: 0, Max: 1
func (e *Frei0rFilterEmboss) SetElevation(value float64) error {
	return e.Element.SetProperty("elevation", value)
}

func (e *Frei0rFilterEmboss) GetElevation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("elevation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Bump height
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterEmboss) SetWidth45(value float64) error {
	return e.Element.SetProperty("width45", value)
}

func (e *Frei0rFilterEmboss) GetWidth45() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("width45")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Composites second input on first input applying user-defined transformation, opacity and blend mode
type Frei0rMixerCairoaffineblend struct {
	*gst.Element
}

func NewFrei0rMixerCairoaffineblend() (*Frei0rMixerCairoaffineblend, error) {
	e, err := gst.NewElement("frei0r-mixer-cairoaffineblend")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerCairoaffineblend{Element: e}, nil
}

func NewFrei0rMixerCairoaffineblendWithName(name string) (*Frei0rMixerCairoaffineblend, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-cairoaffineblend", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerCairoaffineblend{Element: e}, nil
}

// Y position of second input, value interperted as range -2*height - 3*height
// Default: 0.4, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetY(value float64) error {
	return e.Element.SetProperty("y", value)
}

func (e *Frei0rMixerCairoaffineblend) GetY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X position of rotation center within the second input
// Default: 0, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetAnchorX(value float64) error {
	return e.Element.SetProperty("anchor-x", value)
}

func (e *Frei0rMixerCairoaffineblend) GetAnchorX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("anchor-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Opacity of second input
// Default: 1, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetOpacity(value float64) error {
	return e.Element.SetProperty("opacity", value)
}

func (e *Frei0rMixerCairoaffineblend) GetOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X scale of second input, value interperted as range 0 - 5
// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetXScale(value float64) error {
	return e.Element.SetProperty("x-scale", value)
}

func (e *Frei0rMixerCairoaffineblend) GetXScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X position of second input, value interperted as range -2*width - 3*width
// Default: 0.4, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetX(value float64) error {
	return e.Element.SetProperty("x", value)
}

func (e *Frei0rMixerCairoaffineblend) GetX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y scale of second input, value interperted as range 0 - 5
// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetYScale(value float64) error {
	return e.Element.SetProperty("y-scale", value)
}

func (e *Frei0rMixerCairoaffineblend) GetYScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y position of rotation center within the second input
// Default: 0, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetAnchorY(value float64) error {
	return e.Element.SetProperty("anchor-y", value)
}

func (e *Frei0rMixerCairoaffineblend) GetAnchorY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("anchor-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Blend mode used to compose image. Accepted values: 'normal', 'add', 'saturate', 'multiply', 'screen', 'overlay', 'darken', 'lighten', 'colordodge', 'colorburn', 'hardlight', 'softlight', 'difference', 'exclusion', 'hslhue', 'hslsaturation', 'hslcolor', 'hslluminosity'
// Default: normal
func (e *Frei0rMixerCairoaffineblend) SetBlendMode(value string) error {
	return e.Element.SetProperty("blend-mode", value)
}

func (e *Frei0rMixerCairoaffineblend) GetBlendMode() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("blend-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Rotation of second input, value interperted as range 0 - 360
// Default: 0, Min: 0, Max: 1
func (e *Frei0rMixerCairoaffineblend) SetRotation(value float64) error {
	return e.Element.SetProperty("rotation", value)
}

func (e *Frei0rMixerCairoaffineblend) GetRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform an RGB[A] multiply operation between the pixel sources.
type Frei0rMixerMultiply struct {
	*gst.Element
}

func NewFrei0rMixerMultiply() (*Frei0rMixerMultiply, error) {
	e, err := gst.NewElement("frei0r-mixer-multiply")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerMultiply{Element: e}, nil
}

func NewFrei0rMixerMultiplyWithName(name string) (*Frei0rMixerMultiply, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-multiply", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerMultiply{Element: e}, nil
}

// Uses Input 1 as UV Map to distort Input 2
type Frei0rMixerUvMap struct {
	*gst.Element
}

func NewFrei0rMixerUvMap() (*Frei0rMixerUvMap, error) {
	e, err := gst.NewElement("frei0r-mixer-uv-map")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerUvMap{Element: e}, nil
}

func NewFrei0rMixerUvMapWithName(name string) (*Frei0rMixerUvMap, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-uv-map", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerUvMap{Element: e}, nil
}

// Display and manipulation of the alpha channel
type Frei0rFilterAlpha0ps struct {
	*GstVideoFilter
}

func NewFrei0rFilterAlpha0ps() (*Frei0rFilterAlpha0ps, error) {
	e, err := gst.NewElement("frei0r-filter-alpha0ps")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterAlpha0ps{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterAlpha0psWithName(name string) (*Frei0rFilterAlpha0ps, error) {
	e, err := gst.NewElementWithName("frei0r-filter-alpha0ps", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterAlpha0ps{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterAlpha0ps) SetDisplay(value float64) error {
	return e.Element.SetProperty("display", value)
}

func (e *Frei0rFilterAlpha0ps) GetDisplay() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: false
func (e *Frei0rFilterAlpha0ps) SetDisplayInputAlpha(value bool) error {
	return e.Element.SetProperty("display-input-alpha", value)
}

func (e *Frei0rFilterAlpha0ps) GetDisplayInputAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-input-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: false
func (e *Frei0rFilterAlpha0ps) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

func (e *Frei0rFilterAlpha0ps) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterAlpha0ps) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

func (e *Frei0rFilterAlpha0ps) GetOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.333344, Min: 0, Max: 1
func (e *Frei0rFilterAlpha0ps) SetShrinkGrowBlurAmount(value float64) error {
	return e.Element.SetProperty("shrink-grow-blur-amount", value)
}

func (e *Frei0rFilterAlpha0ps) GetShrinkGrowBlurAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shrink-grow-blur-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterAlpha0ps) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *Frei0rFilterAlpha0ps) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pixelize input image.
type Frei0rFilterPixeliz0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterPixeliz0r() (*Frei0rFilterPixeliz0r, error) {
	e, err := gst.NewElement("frei0r-filter-pixeliz0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPixeliz0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterPixeliz0rWithName(name string) (*Frei0rFilterPixeliz0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-pixeliz0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPixeliz0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Vertical size of one "pixel"
// Default: 0.0291667, Min: 0, Max: 1
func (e *Frei0rFilterPixeliz0r) SetBlockHeight(value float64) error {
	return e.Element.SetProperty("block-height", value)
}

func (e *Frei0rFilterPixeliz0r) GetBlockHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("block-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Horizontal size of one "pixel"
// Default: 0.021875, Min: 0, Max: 1
func (e *Frei0rFilterPixeliz0r) SetBlockWidth(value float64) error {
	return e.Element.SetProperty("block-width", value)
}

func (e *Frei0rFilterPixeliz0r) GetBlockWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("block-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Lens vignetting effect, applies natural vignetting
type Frei0rFilterVignette struct {
	*GstVideoFilter
}

func NewFrei0rFilterVignette() (*Frei0rFilterVignette, error) {
	e, err := gst.NewElement("frei0r-filter-vignette")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterVignette{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterVignetteWithName(name string) (*Frei0rFilterVignette, error) {
	e, err := gst.NewElementWithName("frei0r-filter-vignette", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterVignette{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Aspect ratio
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterVignette) SetAspect(value float64) error {
	return e.Element.SetProperty("aspect", value)
}

func (e *Frei0rFilterVignette) GetAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Size of the unaffected center
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterVignette) SetClearcenter(value float64) error {
	return e.Element.SetProperty("clearcenter", value)
}

func (e *Frei0rFilterVignette) GetClearcenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clearcenter")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Softness
// Default: 0.6, Min: 0, Max: 1
func (e *Frei0rFilterVignette) SetSoft(value float64) error {
	return e.Element.SetProperty("soft", value)
}

func (e *Frei0rFilterVignette) GetSoft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("soft")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// detect faces and draw shapes on them
type Frei0rFilterOpencvfacedetect struct {
	*GstVideoFilter
}

func NewFrei0rFilterOpencvfacedetect() (*Frei0rFilterOpencvfacedetect, error) {
	e, err := gst.NewElement("frei0r-filter-opencvfacedetect")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterOpencvfacedetect{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterOpencvfacedetectWithName(name string) (*Frei0rFilterOpencvfacedetect, error) {
	e, err := gst.NewElementWithName("frei0r-filter-opencvfacedetect", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterOpencvfacedetect{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The color of the third object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor3G(value float32) error {
	return e.Element.SetProperty("color-3-g", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor3G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-3-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the fourth object
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor4R(value float32) error {
	return e.Element.SetProperty("color-4-r", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-4-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the fifth object
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor5B(value float32) error {
	return e.Element.SetProperty("color-5-b", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor5B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-5-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the fifth object
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor5G(value float32) error {
	return e.Element.SetProperty("color-5-g", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor5G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-5-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Down scale the image prior detection
// Default: 0.666667, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetScale(value float64) error {
	return e.Element.SetProperty("scale", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The color of the second object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor2B(value float32) error {
	return e.Element.SetProperty("color-2-b", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor2B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-2-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the second object
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor2G(value float32) error {
	return e.Element.SetProperty("color-2-g", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor2G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-2-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the third object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor3B(value float32) error {
	return e.Element.SetProperty("color-3-b", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor3B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-3-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Minimum window size in pixels, divided by 1000
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetSmallest(value float64) error {
	return e.Element.SetProperty("smallest", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetSmallest() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("smallest")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Line width, divided by 100, or fill if 0
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetStroke(value float64) error {
	return e.Element.SetProperty("stroke", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetStroke() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("stroke")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The alpha channel value for the shapes
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetAlpha() (float64, error) {
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

// How many threads to use divided by 100; 0 uses CPU count
// Default: 0.01, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetThreads(value float64) error {
	return e.Element.SetProperty("threads", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetThreads() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The color of the third object
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor3R(value float32) error {
	return e.Element.SetProperty("color-3-r", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the fourth object
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor4B(value float32) error {
	return e.Element.SetProperty("color-4-b", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor4B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-4-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the fourth object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor4G(value float32) error {
	return e.Element.SetProperty("color-4-g", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor4G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-4-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the fifth object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor5R(value float32) error {
	return e.Element.SetProperty("color-5-r", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-5-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Minimum number of rectangles that makes up an object, divided by 100
// Default: 0.02, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetNeighbors(value float64) error {
	return e.Element.SetProperty("neighbors", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetNeighbors() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("neighbors")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The color of the first object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor1B(value float32) error {
	return e.Element.SetProperty("color-1-b", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor1B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-1-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the first object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor1G(value float32) error {
	return e.Element.SetProperty("color-1-g", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor1G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-1-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color of the second object
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor2R(value float32) error {
	return e.Element.SetProperty("color-2-r", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// How often to detect an object in number of frames, divided by 1000
// Default: 0.025, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetRecheck(value float64) error {
	return e.Element.SetProperty("recheck", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetRecheck() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("recheck")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The shape to draw: 0=circle, 0.1=ellipse, 0.2=rectangle, 1=random
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetShape(value float64) error {
	return e.Element.SetProperty("shape", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetShape() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shape")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The search window scale factor, divided by 10
// Default: 0.12, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetSearchScale(value float64) error {
	return e.Element.SetProperty("search-scale", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetSearchScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("search-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Draw with antialiasing
// Default: false
func (e *Frei0rFilterOpencvfacedetect) SetAntialias(value bool) error {
	return e.Element.SetProperty("antialias", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetAntialias() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("antialias")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Full path to the XML pattern model for recognition; look in /usr/share/opencv/haarcascades
// Default: /usr/share/opencv/haarcascades/haarcascade_frontalface_default.xml
func (e *Frei0rFilterOpencvfacedetect) SetClassifier(value string) error {
	return e.Element.SetProperty("classifier", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetClassifier() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("classifier")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The color of the first object
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterOpencvfacedetect) SetColor1R(value float32) error {
	return e.Element.SetProperty("color-1-r", value)
}

func (e *Frei0rFilterOpencvfacedetect) GetColor1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Draws a grid of input images.
type Frei0rFilterCairoimagegrid struct {
	*GstVideoFilter
}

func NewFrei0rFilterCairoimagegrid() (*Frei0rFilterCairoimagegrid, error) {
	e, err := gst.NewElement("frei0r-filter-cairoimagegrid")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCairoimagegrid{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterCairoimagegridWithName(name string) (*Frei0rFilterCairoimagegrid, error) {
	e, err := gst.NewElementWithName("frei0r-filter-cairoimagegrid", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCairoimagegrid{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Number of columns in the image grid. Input range 0 - 1 is interpreted as range 1 - 20
// Default: 0.105263, Min: 0, Max: 1
func (e *Frei0rFilterCairoimagegrid) SetColumns(value float64) error {
	return e.Element.SetProperty("columns", value)
}

func (e *Frei0rFilterCairoimagegrid) GetColumns() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Number of rows in the image grid. Input range 0 - 1 is interpreted as range 1 - 20
// Default: 0.105263, Min: 0, Max: 1
func (e *Frei0rFilterCairoimagegrid) SetRows(value float64) error {
	return e.Element.SetProperty("rows", value)
}

func (e *Frei0rFilterCairoimagegrid) GetRows() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Plasma
type Frei0rFilterDistort0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterDistort0r() (*Frei0rFilterDistort0r, error) {
	e, err := gst.NewElement("frei0r-filter-distort0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDistort0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterDistort0rWithName(name string) (*Frei0rFilterDistort0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-distort0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDistort0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// 'Time Based' or 'Adjustable Velocity'
// Default: false
func (e *Frei0rFilterDistort0r) SetUseVelocity(value bool) error {
	return e.Element.SetProperty("use-velocity", value)
}

func (e *Frei0rFilterDistort0r) GetUseVelocity() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-velocity")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Changing speed of the plasma signal
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterDistort0r) SetVelocity(value float64) error {
	return e.Element.SetProperty("velocity", value)
}

func (e *Frei0rFilterDistort0r) GetVelocity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("velocity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The amplitude of the plasma signal
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterDistort0r) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

func (e *Frei0rFilterDistort0r) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The frequency of the plasma signal
// Default: 0.005, Min: 0, Max: 1
func (e *Frei0rFilterDistort0r) SetFrequency(value float64) error {
	return e.Element.SetProperty("frequency", value)
}

func (e *Frei0rFilterDistort0r) GetFrequency() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adds RGB noise to image.
type Frei0rFilterRgbnoise struct {
	*GstVideoFilter
}

func NewFrei0rFilterRgbnoise() (*Frei0rFilterRgbnoise, error) {
	e, err := gst.NewElement("frei0r-filter-rgbnoise")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterRgbnoise{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterRgbnoiseWithName(name string) (*Frei0rFilterRgbnoise, error) {
	e, err := gst.NewElementWithName("frei0r-filter-rgbnoise", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterRgbnoise{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Amount of noise added
// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterRgbnoise) SetNoise(value float64) error {
	return e.Element.SetProperty("noise", value)
}

func (e *Frei0rFilterRgbnoise) GetNoise() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("noise")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adjust the white balance / color temperature
type Frei0rFilterWhiteBalance struct {
	*GstVideoFilter
}

func NewFrei0rFilterWhiteBalance() (*Frei0rFilterWhiteBalance, error) {
	e, err := gst.NewElement("frei0r-filter-white-balance")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterWhiteBalance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterWhiteBalanceWithName(name string) (*Frei0rFilterWhiteBalance, error) {
	e, err := gst.NewElementWithName("frei0r-filter-white-balance", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterWhiteBalance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Adjust the level of green.
// Default: 0.133333, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalance) SetGreenTint(value float64) error {
	return e.Element.SetProperty("green-tint", value)
}

func (e *Frei0rFilterWhiteBalance) GetGreenTint() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("green-tint")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Choose a color from the source image that should be white.
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalance) SetNeutralColorB(value float32) error {
	return e.Element.SetProperty("neutral-color-b", value)
}

func (e *Frei0rFilterWhiteBalance) GetNeutralColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Choose a color from the source image that should be white.
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalance) SetNeutralColorG(value float32) error {
	return e.Element.SetProperty("neutral-color-g", value)
}

func (e *Frei0rFilterWhiteBalance) GetNeutralColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Choose a color from the source image that should be white.
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterWhiteBalance) SetNeutralColorR(value float32) error {
	return e.Element.SetProperty("neutral-color-r", value)
}

func (e *Frei0rFilterWhiteBalance) GetNeutralColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Perform an RGB[A] grain-merge operation between the pixel sources.
type Frei0rMixerGrainMerge struct {
	*gst.Element
}

func NewFrei0rMixerGrainMerge() (*Frei0rMixerGrainMerge, error) {
	e, err := gst.NewElement("frei0r-mixer-grain-merge")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerGrainMerge{Element: e}, nil
}

func NewFrei0rMixerGrainMergeWithName(name string) (*Frei0rMixerGrainMerge, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-grain-merge", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerGrainMerge{Element: e}, nil
}

// Demo scene 8bit plasma
type Frei0rSrcPlasma struct {
	*base.GstPushSrc
}

func NewFrei0rSrcPlasma() (*Frei0rSrcPlasma, error) {
	e, err := gst.NewElement("frei0r-src-plasma")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcPlasma{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcPlasmaWithName(name string) (*Frei0rSrcPlasma, error) {
	e, err := gst.NewElementWithName("frei0r-src-plasma", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcPlasma{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcPlasma) SetParam1Move(value float64) error {
	return e.Element.SetProperty("param-1-move", value)
}

func (e *Frei0rSrcPlasma) GetParam1Move() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-1-move")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcPlasma) SetParam1Speed(value float64) error {
	return e.Element.SetProperty("param-1-speed", value)
}

func (e *Frei0rSrcPlasma) GetParam1Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-1-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcPlasma) SetParam2Move(value float64) error {
	return e.Element.SetProperty("param-2-move", value)
}

func (e *Frei0rSrcPlasma) GetParam2Move() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-2-move")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcPlasma) SetParam2Speed(value float64) error {
	return e.Element.SetProperty("param-2-speed", value)
}

func (e *Frei0rSrcPlasma) GetParam2Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-2-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcPlasma) SetParam3Speed(value float64) error {
	return e.Element.SetProperty("param-3-speed", value)
}

func (e *Frei0rSrcPlasma) GetParam3Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-3-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcPlasma) SetParam4Speed(value float64) error {
	return e.Element.SetProperty("param-4-speed", value)
}

func (e *Frei0rSrcPlasma) GetParam4Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-4-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Generates resolution test patterns
type Frei0rSrcTestPatR struct {
	*base.GstPushSrc
}

func NewFrei0rSrcTestPatR() (*Frei0rSrcTestPatR, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-r")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatR{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcTestPatRWithName(name string) (*Frei0rSrcTestPatR, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatR{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Type of test pattern
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatR) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rSrcTestPatR) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amplitude (contrast) of the pattern
// Default: 0.8, Min: 0, Max: 1
func (e *Frei0rSrcTestPatR) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

func (e *Frei0rSrcTestPatR) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pixel aspect ratio presets
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatR) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

func (e *Frei0rSrcTestPatR) GetAspectType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Into which color channel to draw
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatR) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

func (e *Frei0rSrcTestPatR) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pattern 7 H frequency
// Default: 0.03, Min: 0, Max: 1
func (e *Frei0rSrcTestPatR) SetFreq1(value float64) error {
	return e.Element.SetProperty("freq-1", value)
}

func (e *Frei0rSrcTestPatR) GetFreq1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("freq-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pattern 7 V frequency
// Default: 0.03, Min: 0, Max: 1
func (e *Frei0rSrcTestPatR) SetFreq2(value float64) error {
	return e.Element.SetProperty("freq-2", value)
}

func (e *Frei0rSrcTestPatR) GetFreq2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("freq-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Use linear period sweep
// Default: false
func (e *Frei0rSrcTestPatR) SetLinPSwp(value bool) error {
	return e.Element.SetProperty("lin-p-swp", value)
}

func (e *Frei0rSrcTestPatR) GetLinPSwp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lin-p-swp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Manual pixel aspect ratio
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rSrcTestPatR) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

func (e *Frei0rSrcTestPatR) GetManualAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Fills alpha channel with a gradient
type Frei0rFilterAlphagrad struct {
	*GstVideoFilter
}

func NewFrei0rFilterAlphagrad() (*Frei0rFilterAlphagrad, error) {
	e, err := gst.NewElement("frei0r-filter-alphagrad")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterAlphagrad{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterAlphagradWithName(name string) (*Frei0rFilterAlphagrad, error) {
	e, err := gst.NewElementWithName("frei0r-filter-alphagrad", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterAlphagrad{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterAlphagrad) SetTransitionWidth(value float64) error {
	return e.Element.SetProperty("transition-width", value)
}

func (e *Frei0rFilterAlphagrad) GetTransitionWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("transition-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterAlphagrad) SetMax(value float64) error {
	return e.Element.SetProperty("max", value)
}

func (e *Frei0rFilterAlphagrad) GetMax() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("max")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterAlphagrad) SetMin(value float64) error {
	return e.Element.SetProperty("min", value)
}

func (e *Frei0rFilterAlphagrad) GetMin() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("min")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterAlphagrad) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

func (e *Frei0rFilterAlphagrad) GetOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterAlphagrad) SetPosition(value float64) error {
	return e.Element.SetProperty("position", value)
}

func (e *Frei0rFilterAlphagrad) GetPosition() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterAlphagrad) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

func (e *Frei0rFilterAlphagrad) GetTilt() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Generates spatial impulse and step test patterns
type Frei0rSrcTestPatI struct {
	*base.GstPushSrc
}

func NewFrei0rSrcTestPatI() (*Frei0rSrcTestPatI, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-i")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatI{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcTestPatIWithName(name string) (*Frei0rSrcTestPatI, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-i", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatI{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Type of test pattern
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatI) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rSrcTestPatI) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Width of impulse
// Default: 0.040404, Min: 0, Max: 1
func (e *Frei0rSrcTestPatI) SetWidth(value float64) error {
	return e.Element.SetProperty("width", value)
}

func (e *Frei0rSrcTestPatI) GetWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amplitude (contrast) of the pattern
// Default: 0.8, Min: 0, Max: 1
func (e *Frei0rSrcTestPatI) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

func (e *Frei0rSrcTestPatI) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Into which color channel to draw
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatI) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

func (e *Frei0rSrcTestPatI) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Change polarity of impulse/step
// Default: false
func (e *Frei0rSrcTestPatI) SetNegative(value bool) error {
	return e.Element.SetProperty("negative", value)
}

func (e *Frei0rSrcTestPatI) GetNegative() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("negative")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Angle of step function
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rSrcTestPatI) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

func (e *Frei0rSrcTestPatI) GetTilt() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Tunes the alpha channel.
type Frei0rFilterTransparency struct {
	*GstVideoFilter
}

func NewFrei0rFilterTransparency() (*Frei0rFilterTransparency, error) {
	e, err := gst.NewElement("frei0r-filter-transparency")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTransparency{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterTransparencyWithName(name string) (*Frei0rFilterTransparency, error) {
	e, err := gst.NewElementWithName("frei0r-filter-transparency", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTransparency{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The transparency value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTransparency) SetTransparency(value float64) error {
	return e.Element.SetProperty("transparency", value)
}

func (e *Frei0rFilterTransparency) GetTransparency() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("transparency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Applies a pre-made color effect to image
type Frei0rFilterColortap struct {
	*GstVideoFilter
}

func NewFrei0rFilterColortap() (*Frei0rFilterColortap, error) {
	e, err := gst.NewElement("frei0r-filter-colortap")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColortap{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterColortapWithName(name string) (*Frei0rFilterColortap, error) {
	e, err := gst.NewElementWithName("frei0r-filter-colortap", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColortap{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Lookup table used to filter colors. One of: xpro, sepia, heat, red_green, old_photo, xray, esses, yellow_blue
// Default: esses
func (e *Frei0rFilterColortap) SetTable(value string) error {
	return e.Element.SetProperty("table", value)
}

func (e *Frei0rFilterColortap) GetTable() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("table")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Non rectilinear lens mappings
type Frei0rFilterDefish0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterDefish0r() (*Frei0rFilterDefish0r, error) {
	e, err := gst.NewElement("frei0r-filter-defish0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDefish0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterDefish0rWithName(name string) (*Frei0rFilterDefish0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-defish0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDefish0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Scaling method
// Default: 0.666667, Min: 0, Max: 1
func (e *Frei0rFilterDefish0r) SetScaling(value float64) error {
	return e.Element.SetProperty("scaling", value)
}

func (e *Frei0rFilterDefish0r) GetScaling() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scaling")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Mapping function
// Default: 0.666667, Min: 0, Max: 1
func (e *Frei0rFilterDefish0r) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rFilterDefish0r) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Focal Ratio
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterDefish0r) SetAmount(value float64) error {
	return e.Element.SetProperty("amount", value)
}

func (e *Frei0rFilterDefish0r) GetAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pixel aspect ratio presets
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterDefish0r) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

func (e *Frei0rFilterDefish0r) GetAspectType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Fish or Defish
// Default: false
func (e *Frei0rFilterDefish0r) SetDefish(value bool) error {
	return e.Element.SetProperty("defish", value)
}

func (e *Frei0rFilterDefish0r) GetDefish() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("defish")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Quality of interpolation
// Default: 0.166667, Min: 0, Max: 1
func (e *Frei0rFilterDefish0r) SetInterpolator(value float64) error {
	return e.Element.SetProperty("interpolator", value)
}

func (e *Frei0rFilterDefish0r) GetInterpolator() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("interpolator")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Manual Pixel Aspect ratio
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterDefish0r) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

func (e *Frei0rFilterDefish0r) GetManualAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Manual Scale
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterDefish0r) SetManualScale(value float64) error {
	return e.Element.SetProperty("manual-scale", value)
}

func (e *Frei0rFilterDefish0r) GetManualScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Three types of fast IIR blurring
type Frei0rFilterIirBlur struct {
	*GstVideoFilter
}

func NewFrei0rFilterIirBlur() (*Frei0rFilterIirBlur, error) {
	e, err := gst.NewElement("frei0r-filter-iir-blur")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterIirBlur{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterIirBlurWithName(name string) (*Frei0rFilterIirBlur, error) {
	e, err := gst.NewElementWithName("frei0r-filter-iir-blur", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterIirBlur{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Blur type
// Default: 0.333344, Min: 0, Max: 1
func (e *Frei0rFilterIirBlur) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rFilterIirBlur) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount of blur
// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterIirBlur) SetAmount(value float64) error {
	return e.Element.SetProperty("amount", value)
}

func (e *Frei0rFilterIirBlur) GetAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Edge compensation
// Default: true
func (e *Frei0rFilterIirBlur) SetEdge(value bool) error {
	return e.Element.SetProperty("edge", value)
}

func (e *Frei0rFilterIirBlur) GetEdge() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("edge")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adjust color balance with 3 color points
type Frei0rFilter3PointColorBalance struct {
	*GstVideoFilter
}

func NewFrei0rFilter3PointColorBalance() (*Frei0rFilter3PointColorBalance, error) {
	e, err := gst.NewElement("frei0r-filter-3-point-color-balance")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilter3PointColorBalance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilter3PointColorBalanceWithName(name string) (*Frei0rFilter3PointColorBalance, error) {
	e, err := gst.NewElementWithName("frei0r-filter-3-point-color-balance", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilter3PointColorBalance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Black color
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetBlackColorG(value float32) error {
	return e.Element.SetProperty("black-color-g", value)
}

func (e *Frei0rFilter3PointColorBalance) GetBlackColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("black-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Black color
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetBlackColorR(value float32) error {
	return e.Element.SetProperty("black-color-r", value)
}

func (e *Frei0rFilter3PointColorBalance) GetBlackColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("black-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Gray color
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetGrayColorB(value float32) error {
	return e.Element.SetProperty("gray-color-b", value)
}

func (e *Frei0rFilter3PointColorBalance) GetGrayColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gray-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Gray color
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetGrayColorG(value float32) error {
	return e.Element.SetProperty("gray-color-g", value)
}

func (e *Frei0rFilter3PointColorBalance) GetGrayColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gray-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Gray color
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetGrayColorR(value float32) error {
	return e.Element.SetProperty("gray-color-r", value)
}

func (e *Frei0rFilter3PointColorBalance) GetGrayColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gray-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// White color
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetWhiteColorG(value float32) error {
	return e.Element.SetProperty("white-color-g", value)
}

func (e *Frei0rFilter3PointColorBalance) GetWhiteColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("white-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// White color
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetWhiteColorR(value float32) error {
	return e.Element.SetProperty("white-color-r", value)
}

func (e *Frei0rFilter3PointColorBalance) GetWhiteColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("white-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Black color
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetBlackColorB(value float32) error {
	return e.Element.SetProperty("black-color-b", value)
}

func (e *Frei0rFilter3PointColorBalance) GetBlackColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("black-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Source image on left side
// Default: true
func (e *Frei0rFilter3PointColorBalance) SetSourceImageOnLeftSide(value bool) error {
	return e.Element.SetProperty("source-image-on-left-side", value)
}

func (e *Frei0rFilter3PointColorBalance) GetSourceImageOnLeftSide() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("source-image-on-left-side")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Split privew
// Default: true
func (e *Frei0rFilter3PointColorBalance) SetSplitPreview(value bool) error {
	return e.Element.SetProperty("split-preview", value)
}

func (e *Frei0rFilter3PointColorBalance) GetSplitPreview() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("split-preview")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// White color
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilter3PointColorBalance) SetWhiteColorB(value float32) error {
	return e.Element.SetProperty("white-color-b", value)
}

func (e *Frei0rFilter3PointColorBalance) GetWhiteColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("white-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Cartoonify video, do a form of edge detect
type Frei0rFilterCartoon struct {
	*GstVideoFilter
}

func NewFrei0rFilterCartoon() (*Frei0rFilterCartoon, error) {
	e, err := gst.NewElement("frei0r-filter-cartoon")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCartoon{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterCartoonWithName(name string) (*Frei0rFilterCartoon, error) {
	e, err := gst.NewElementWithName("frei0r-filter-cartoon", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCartoon{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// difference space: a value from 0 to 256 (mapped to [0,1])
// Default: 0.00390625, Min: 0, Max: 1
func (e *Frei0rFilterCartoon) SetDiffspace(value float64) error {
	return e.Element.SetProperty("diffspace", value)
}

func (e *Frei0rFilterCartoon) GetDiffspace() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("diffspace")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// level of trip: mapped to [0,1] asymptotical
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterCartoon) SetTriplevel(value float64) error {
	return e.Element.SetProperty("triplevel", value)
}

func (e *Frei0rFilterCartoon) GetTriplevel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("triplevel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Particles generated on prime number sinusoidal blossoming
type Frei0rSrcPartik0l struct {
	*base.GstPushSrc
}

func NewFrei0rSrcPartik0l() (*Frei0rSrcPartik0l, error) {
	e, err := gst.NewElement("frei0r-src-partik0l")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcPartik0l{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcPartik0lWithName(name string) (*Frei0rSrcPartik0l, error) {
	e, err := gst.NewElementWithName("frei0r-src-partik0l", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcPartik0l{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// blossom on a lower prime number
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcPartik0l) SetDown(value float64) error {
	return e.Element.SetProperty("down", value)
}

func (e *Frei0rSrcPartik0l) GetDown() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("down")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// blossom on a higher prime number
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcPartik0l) SetUp(value float64) error {
	return e.Element.SetProperty("up", value)
}

func (e *Frei0rSrcPartik0l) GetUp() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("up")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Generates linearity checking patterns
type Frei0rSrcTestPatL struct {
	*base.GstPushSrc
}

func NewFrei0rSrcTestPatL() (*Frei0rSrcTestPatL, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-l")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatL{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcTestPatLWithName(name string) (*Frei0rSrcTestPatL, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-l", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatL{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Into which color channel to draw
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatL) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

func (e *Frei0rSrcTestPatL) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Type of test pattern
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatL) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rSrcTestPatL) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Turns image black/white.
type Frei0rFilterBw0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterBw0r() (*Frei0rFilterBw0r, error) {
	e, err := gst.NewElement("frei0r-filter-bw0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBw0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterBw0rWithName(name string) (*Frei0rFilterBw0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-bw0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBw0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Draws a gradient on top of image. Filter is given gradient start and end points, colors and opacities.
type Frei0rFilterCairogradient struct {
	*GstVideoFilter
}

func NewFrei0rFilterCairogradient() (*Frei0rFilterCairogradient, error) {
	e, err := gst.NewElement("frei0r-filter-cairogradient")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCairogradient{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterCairogradientWithName(name string) (*Frei0rFilterCairogradient, error) {
	e, err := gst.NewElementWithName("frei0r-filter-cairogradient", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCairogradient{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Opacity of the first color of the gradient
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetStartOpacity(value float64) error {
	return e.Element.SetProperty("start-opacity", value)
}

func (e *Frei0rFilterCairogradient) GetStartOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("start-opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Blend mode used to compose gradient on image. Accepted values: 'normal', 'add', 'saturate', 'multiply', 'screen', 'overlay', 'darken', 'lighten', 'colordodge', 'colorburn', 'hardlight', 'softlight', 'difference', 'exclusion', 'hslhue', 'hslsaturation', 'hslcolor', 'hslluminosity'
// Default: normal
func (e *Frei0rFilterCairogradient) SetBlendMode(value string) error {
	return e.Element.SetProperty("blend-mode", value)
}

func (e *Frei0rFilterCairogradient) GetBlendMode() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("blend-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Second color of the gradient
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetEndColorB(value float32) error {
	return e.Element.SetProperty("end-color-b", value)
}

func (e *Frei0rFilterCairogradient) GetEndColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("end-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Y position of the end point of the gradient
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetEndY(value float64) error {
	return e.Element.SetProperty("end-y", value)
}

func (e *Frei0rFilterCairogradient) GetEndY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("end-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Position of first color in the line connecting gradient ends, really useful only for radial gradient
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetOffset(value float64) error {
	return e.Element.SetProperty("offset", value)
}

func (e *Frei0rFilterCairogradient) GetOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// First color of the gradient
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetStartColorB(value float32) error {
	return e.Element.SetProperty("start-color-b", value)
}

func (e *Frei0rFilterCairogradient) GetStartColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("start-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// First color of the gradient
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetStartColorG(value float32) error {
	return e.Element.SetProperty("start-color-g", value)
}

func (e *Frei0rFilterCairogradient) GetStartColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("start-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// First color of the gradient
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetStartColorR(value float32) error {
	return e.Element.SetProperty("start-color-r", value)
}

func (e *Frei0rFilterCairogradient) GetStartColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("start-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Y position of the start point of the gradient
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetStartY(value float64) error {
	return e.Element.SetProperty("start-y", value)
}

func (e *Frei0rFilterCairogradient) GetStartY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("start-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Second color of the gradient
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetEndColorR(value float32) error {
	return e.Element.SetProperty("end-color-r", value)
}

func (e *Frei0rFilterCairogradient) GetEndColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("end-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Opacity of the second color of the gradient
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetEndOpacity(value float64) error {
	return e.Element.SetProperty("end-opacity", value)
}

func (e *Frei0rFilterCairogradient) GetEndOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("end-opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Linear or radial gradient
// Default: gradient_linear
func (e *Frei0rFilterCairogradient) SetPattern(value string) error {
	return e.Element.SetProperty("pattern", value)
}

func (e *Frei0rFilterCairogradient) GetPattern() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// X position of the start point of the gradient
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetStartX(value float64) error {
	return e.Element.SetProperty("start-x", value)
}

func (e *Frei0rFilterCairogradient) GetStartX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("start-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Second color of the gradient
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetEndColorG(value float32) error {
	return e.Element.SetProperty("end-color-g", value)
}

func (e *Frei0rFilterCairogradient) GetEndColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("end-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// X position of the end point of the gradient
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterCairogradient) SetEndX(value float64) error {
	return e.Element.SetProperty("end-x", value)
}

func (e *Frei0rFilterCairogradient) GetEndX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("end-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Remove green or blue spill light from subjects shot in front of green or blue screen
type Frei0rFilterSpillsupress struct {
	*GstVideoFilter
}

func NewFrei0rFilterSpillsupress() (*Frei0rFilterSpillsupress, error) {
	e, err := gst.NewElement("frei0r-filter-spillsupress")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSpillsupress{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSpillsupressWithName(name string) (*Frei0rFilterSpillsupress, error) {
	e, err := gst.NewElementWithName("frei0r-filter-spillsupress", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSpillsupress{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Defines if green or blue screen spill suppress is applied
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSpillsupress) SetSupresstype(value float64) error {
	return e.Element.SetProperty("supresstype", value)
}

func (e *Frei0rFilterSpillsupress) GetSupresstype() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("supresstype")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform an RGB[A] addition operation of the pixel sources.
type Frei0rMixerAddition struct {
	*gst.Element
}

func NewFrei0rMixerAddition() (*Frei0rMixerAddition, error) {
	e, err := gst.NewElement("frei0r-mixer-addition")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAddition{Element: e}, nil
}

func NewFrei0rMixerAdditionWithName(name string) (*Frei0rMixerAddition, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-addition", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAddition{Element: e}, nil
}

// Averages Input 1 and uses this as Alpha Channel on Input 2
type Frei0rMixerAlphaInjection struct {
	*gst.Element
}

func NewFrei0rMixerAlphaInjection() (*Frei0rMixerAlphaInjection, error) {
	e, err := gst.NewElement("frei0r-mixer-alpha-injection")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaInjection{Element: e}, nil
}

func NewFrei0rMixerAlphaInjectionWithName(name string) (*Frei0rMixerAlphaInjection, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-alpha-injection", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaInjection{Element: e}, nil
}

// the alpha OUT operation
type Frei0rMixerAlphaout struct {
	*gst.Element
}

func NewFrei0rMixerAlphaout() (*Frei0rMixerAlphaout, error) {
	e, err := gst.NewElement("frei0r-mixer-alphaout")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaout{Element: e}, nil
}

func NewFrei0rMixerAlphaoutWithName(name string) (*Frei0rMixerAlphaout, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-alphaout", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaout{Element: e}, nil
}

// Perform a lighten operation between two sources (maximum value of both sources).
type Frei0rMixerLighten struct {
	*gst.Element
}

func NewFrei0rMixerLighten() (*Frei0rMixerLighten, error) {
	e, err := gst.NewElement("frei0r-mixer-lighten")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerLighten{Element: e}, nil
}

func NewFrei0rMixerLightenWithName(name string) (*Frei0rMixerLighten, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-lighten", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerLighten{Element: e}, nil
}

// delayed alpha smoothed blit of time
type Frei0rFilterBaltan struct {
	*GstVideoFilter
}

func NewFrei0rFilterBaltan() (*Frei0rFilterBaltan, error) {
	e, err := gst.NewElement("frei0r-filter-baltan")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBaltan{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterBaltanWithName(name string) (*Frei0rFilterBaltan, error) {
	e, err := gst.NewElementWithName("frei0r-filter-baltan", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBaltan{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Adds glitches and block shifting
type Frei0rFilterGlitch0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterGlitch0r() (*Frei0rFilterGlitch0r, error) {
	e, err := gst.NewElement("frei0r-filter-glitch0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterGlitch0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterGlitch0rWithName(name string) (*Frei0rFilterGlitch0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-glitch0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterGlitch0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// How intensive should be color distortion
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterGlitch0r) SetColorGlitchingIntensity(value float64) error {
	return e.Element.SetProperty("color-glitching-intensity", value)
}

func (e *Frei0rFilterGlitch0r) GetColorGlitchingIntensity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color-glitching-intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// How frequently the glitch should be applied
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterGlitch0r) SetGlitchFrequency(value float64) error {
	return e.Element.SetProperty("glitch-frequency", value)
}

func (e *Frei0rFilterGlitch0r) GetGlitchFrequency() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("glitch-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// How much we should move blocks when glitching
// Default: 0.499218, Min: 0, Max: 1
func (e *Frei0rFilterGlitch0r) SetShiftIntensity(value float64) error {
	return e.Element.SetProperty("shift-intensity", value)
}

func (e *Frei0rFilterGlitch0r) GetShiftIntensity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shift-intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Height range of the block that will be shifted/glitched
// Default: 0.498956, Min: 0, Max: 1
func (e *Frei0rFilterGlitch0r) SetBlockHeight(value float64) error {
	return e.Element.SetProperty("block-height", value)
}

func (e *Frei0rFilterGlitch0r) GetBlockHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("block-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Distorts the image for a pseudo perspective
type Frei0rFilterPerspective struct {
	*GstVideoFilter
}

func NewFrei0rFilterPerspective() (*Frei0rFilterPerspective, error) {
	e, err := gst.NewElement("frei0r-filter-perspective")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPerspective{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterPerspectiveWithName(name string) (*Frei0rFilterPerspective, error) {
	e, err := gst.NewElementWithName("frei0r-filter-perspective", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPerspective{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetTopRightX(value float64) error {
	return e.Element.SetProperty("top-right-x", value)
}

func (e *Frei0rFilterPerspective) GetTopRightX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-right-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetBottomLeftY(value float64) error {
	return e.Element.SetProperty("bottom-left-Y", value)
}

func (e *Frei0rFilterPerspective) GetBottomLeftY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-left-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetBottomLeftX(value float64) error {
	return e.Element.SetProperty("bottom-left-x", value)
}

func (e *Frei0rFilterPerspective) GetBottomLeftX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-left-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetBottomRightY(value float64) error {
	return e.Element.SetProperty("bottom-right-Y", value)
}

func (e *Frei0rFilterPerspective) GetBottomRightY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-right-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetBottomRightX(value float64) error {
	return e.Element.SetProperty("bottom-right-x", value)
}

func (e *Frei0rFilterPerspective) GetBottomRightX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-right-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetTopLeftY(value float64) error {
	return e.Element.SetProperty("top-left-Y", value)
}

func (e *Frei0rFilterPerspective) GetTopLeftY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-left-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetTopLeftX(value float64) error {
	return e.Element.SetProperty("top-left-x", value)
}

func (e *Frei0rFilterPerspective) GetTopLeftX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-left-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPerspective) SetTopRightY(value float64) error {
	return e.Element.SetProperty("top-right-Y", value)
}

func (e *Frei0rFilterPerspective) GetTopRightY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-right-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Variable-size square blur
type Frei0rFilterSquareblur struct {
	*GstVideoFilter
}

func NewFrei0rFilterSquareblur() (*Frei0rFilterSquareblur, error) {
	e, err := gst.NewElement("frei0r-filter-squareblur")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSquareblur{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSquareblurWithName(name string) (*Frei0rFilterSquareblur, error) {
	e, err := gst.NewElementWithName("frei0r-filter-squareblur", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSquareblur{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The size of the kernel, as a proportion to its coverage of the image
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSquareblur) SetKernelSize(value float64) error {
	return e.Element.SetProperty("kernel-size", value)
}

func (e *Frei0rFilterSquareblur) GetKernelSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("kernel-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Thresholds a source image
type Frei0rFilterThreshold0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterThreshold0r() (*Frei0rFilterThreshold0r, error) {
	e, err := gst.NewElement("frei0r-filter-threshold0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterThreshold0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterThreshold0rWithName(name string) (*Frei0rFilterThreshold0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-threshold0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterThreshold0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The threshold
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterThreshold0r) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *Frei0rFilterThreshold0r) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform a conversion to color only of the source input1 using the hue and saturation values of input2.
type Frei0rMixerColorOnly struct {
	*gst.Element
}

func NewFrei0rMixerColorOnly() (*Frei0rMixerColorOnly, error) {
	e, err := gst.NewElement("frei0r-mixer-color-only")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerColorOnly{Element: e}, nil
}

func NewFrei0rMixerColorOnlyWithName(name string) (*Frei0rMixerColorOnly, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-color-only", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerColorOnly{Element: e}, nil
}

// Extracts Green from Image
type Frei0rFilterG struct {
	*GstVideoFilter
}

func NewFrei0rFilterG() (*Frei0rFilterG, error) {
	e, err := gst.NewElement("frei0r-filter-g")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterG{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterGWithName(name string) (*Frei0rFilterG, error) {
	e, err := gst.NewElementWithName("frei0r-filter-g", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterG{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Normalize (aka histogram stretch, contrast stretch)
type Frei0rFilterNormaliz0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterNormaliz0r() (*Frei0rFilterNormaliz0r, error) {
	e, err := gst.NewElement("frei0r-filter-normaliz0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNormaliz0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterNormaliz0rWithName(name string) (*Frei0rFilterNormaliz0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-normaliz0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNormaliz0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Amount of temporal smoothing of the input range, to reduce flicker (default 0.0)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetSmoothing(value float64) error {
	return e.Element.SetProperty("smoothing", value)
}

func (e *Frei0rFilterNormaliz0r) GetSmoothing() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("smoothing")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Strength of filter, from no effect to full normalization (default 1.0)
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetStrength(value float64) error {
	return e.Element.SetProperty("strength", value)
}

func (e *Frei0rFilterNormaliz0r) GetStrength() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Output color to which darkest input color is mapped (default black)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetBlackptB(value float32) error {
	return e.Element.SetProperty("blackpt-b", value)
}

func (e *Frei0rFilterNormaliz0r) GetBlackptB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blackpt-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Output color to which darkest input color is mapped (default black)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetBlackptG(value float32) error {
	return e.Element.SetProperty("blackpt-g", value)
}

func (e *Frei0rFilterNormaliz0r) GetBlackptG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blackpt-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Output color to which darkest input color is mapped (default black)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetBlackptR(value float32) error {
	return e.Element.SetProperty("blackpt-r", value)
}

func (e *Frei0rFilterNormaliz0r) GetBlackptR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blackpt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Proportion of independent to linked channel normalization (default 1.0)
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetIndependence(value float64) error {
	return e.Element.SetProperty("independence", value)
}

func (e *Frei0rFilterNormaliz0r) GetIndependence() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("independence")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Output color to which brightest input color is mapped (default white)
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetWhiteptB(value float32) error {
	return e.Element.SetProperty("whitept-b", value)
}

func (e *Frei0rFilterNormaliz0r) GetWhiteptB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("whitept-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Output color to which brightest input color is mapped (default white)
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetWhiteptG(value float32) error {
	return e.Element.SetProperty("whitept-g", value)
}

func (e *Frei0rFilterNormaliz0r) GetWhiteptG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("whitept-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Output color to which brightest input color is mapped (default white)
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterNormaliz0r) SetWhiteptR(value float32) error {
	return e.Element.SetProperty("whitept-r", value)
}

func (e *Frei0rFilterNormaliz0r) GetWhiteptR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("whitept-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// broken tv
type Frei0rFilterNosync0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterNosync0r() (*Frei0rFilterNosync0r, error) {
	e, err := gst.NewElement("frei0r-filter-nosync0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNosync0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterNosync0rWithName(name string) (*Frei0rFilterNosync0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-nosync0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNosync0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// the hsync offset
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterNosync0r) SetHsync(value float64) error {
	return e.Element.SetProperty("hsync", value)
}

func (e *Frei0rFilterNosync0r) GetHsync() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hsync")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Generates cross sections of color spaces
type Frei0rSrcTestPatC struct {
	*base.GstPushSrc
}

func NewFrei0rSrcTestPatC() (*Frei0rSrcTestPatC, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-c")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatC{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcTestPatCWithName(name string) (*Frei0rSrcTestPatC, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-c", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatC{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatC) SetColorSpace(value float64) error {
	return e.Element.SetProperty("color-space", value)
}

func (e *Frei0rSrcTestPatC) GetColorSpace() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color-space")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatC) SetCrossSection(value float64) error {
	return e.Element.SetProperty("cross-section", value)
}

func (e *Frei0rSrcTestPatC) GetCrossSection() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("cross-section")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: true
func (e *Frei0rSrcTestPatC) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

func (e *Frei0rSrcTestPatC) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rSrcTestPatC) SetThirdAxisValue(value float64) error {
	return e.Element.SetProperty("third-axis-value", value)
}

func (e *Frei0rSrcTestPatC) GetThirdAxisValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("third-axis-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Simple color adjustment
type Frei0rFilterColoradjRgb struct {
	*GstVideoFilter
}

func NewFrei0rFilterColoradjRgb() (*Frei0rFilterColoradjRgb, error) {
	e, err := gst.NewElement("frei0r-filter-coloradj-rgb")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColoradjRgb{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterColoradjRgbWithName(name string) (*Frei0rFilterColoradjRgb, error) {
	e, err := gst.NewElementWithName("frei0r-filter-coloradj-rgb", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColoradjRgb{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Type of color adjustment
// Default: 0.333344, Min: 0, Max: 1
func (e *Frei0rFilterColoradjRgb) SetAction(value float64) error {
	return e.Element.SetProperty("action", value)
}

func (e *Frei0rFilterColoradjRgb) GetAction() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("action")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adjust only areas with nonzero alpha
// Default: false
func (e *Frei0rFilterColoradjRgb) SetAlphaControlled(value bool) error {
	return e.Element.SetProperty("alpha-controlled", value)
}

func (e *Frei0rFilterColoradjRgb) GetAlphaControlled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("alpha-controlled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Amount of blue
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColoradjRgb) SetB(value float64) error {
	return e.Element.SetProperty("b", value)
}

func (e *Frei0rFilterColoradjRgb) GetB() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount of green
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColoradjRgb) SetG(value float64) error {
	return e.Element.SetProperty("g", value)
}

func (e *Frei0rFilterColoradjRgb) GetG() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Don't change brightness
// Default: true
func (e *Frei0rFilterColoradjRgb) SetKeepLuma(value bool) error {
	return e.Element.SetProperty("keep-luma", value)
}

func (e *Frei0rFilterColoradjRgb) GetKeepLuma() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-luma")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: 0.500025, Min: 0, Max: 1
func (e *Frei0rFilterColoradjRgb) SetLumaFormula(value float64) error {
	return e.Element.SetProperty("luma-formula", value)
}

func (e *Frei0rFilterColoradjRgb) GetLumaFormula() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("luma-formula")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount of red
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColoradjRgb) SetR(value float64) error {
	return e.Element.SetProperty("r", value)
}

func (e *Frei0rFilterColoradjRgb) GetR() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform a darken operation between two sources (minimum value of both sources).
type Frei0rMixerDarken struct {
	*gst.Element
}

func NewFrei0rMixerDarken() (*Frei0rMixerDarken, error) {
	e, err := gst.NewElement("frei0r-mixer-darken")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDarken{Element: e}, nil
}

func NewFrei0rMixerDarkenWithName(name string) (*Frei0rMixerDarken, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-darken", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDarken{Element: e}, nil
}

// This filter creates a false image from a visible + infrared source.
type Frei0rFilterNdviFilter struct {
	*GstVideoFilter
}

func NewFrei0rFilterNdviFilter() (*Frei0rFilterNdviFilter, error) {
	e, err := gst.NewElement("frei0r-filter-ndvi-filter")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNdviFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterNdviFilterWithName(name string) (*Frei0rFilterNdviFilter, error) {
	e, err := gst.NewElementWithName("frei0r-filter-ndvi-filter", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNdviFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The index calculation to use. One of 'ndvi' or 'vi'.
// Default: ndvi
func (e *Frei0rFilterNdviFilter) SetIndexCalculation(value string) error {
	return e.Element.SetProperty("index-calculation", value)
}

func (e *Frei0rFilterNdviFilter) GetIndexCalculation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("index-calculation")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Control legend display. One of 'off' or 'bottom'.
// Default: off
func (e *Frei0rFilterNdviFilter) SetLegend(value string) error {
	return e.Element.SetProperty("legend", value)
}

func (e *Frei0rFilterNdviFilter) GetLegend() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("legend")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The channel to use for the near-infrared component. One of 'r', 'g', or 'b'.
// Default: r
func (e *Frei0rFilterNdviFilter) SetNirChannel(value string) error {
	return e.Element.SetProperty("nir-channel", value)
}

func (e *Frei0rFilterNdviFilter) GetNirChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("nir-channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The channel to use for the visible component. One of 'r', 'g', or 'b'.
// Default: b
func (e *Frei0rFilterNdviFilter) SetVisibleChannel(value string) error {
	return e.Element.SetProperty("visible-channel", value)
}

func (e *Frei0rFilterNdviFilter) GetVisibleChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("visible-channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The color map to use. One of 'earth', 'grayscale', 'heat' or 'rainbow'.
// Default: grayscale
func (e *Frei0rFilterNdviFilter) SetColorMap(value string) error {
	return e.Element.SetProperty("color-map", value)
}

func (e *Frei0rFilterNdviFilter) GetColorMap() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("color-map")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The number of color levels to use in the false image (divided by 1000).
// Default: 0.256, Min: 0, Max: 1
func (e *Frei0rFilterNdviFilter) SetLevels(value float64) error {
	return e.Element.SetProperty("levels", value)
}

func (e *Frei0rFilterNdviFilter) GetLevels() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("levels")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// An offset to be applied to the near-infrared component (mapped to [-100%%, 100%%].
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterNdviFilter) SetNirOffset(value float64) error {
	return e.Element.SetProperty("nir-offset", value)
}

func (e *Frei0rFilterNdviFilter) GetNirOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("nir-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// A scaling factor to be applied to the near-infrared component (divided by 10).
// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterNdviFilter) SetNirScale(value float64) error {
	return e.Element.SetProperty("nir-scale", value)
}

func (e *Frei0rFilterNdviFilter) GetNirScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("nir-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// An offset to be applied to the visible component (mapped to [-100%%, 100%%].
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterNdviFilter) SetVisOffset(value float64) error {
	return e.Element.SetProperty("vis-offset", value)
}

func (e *Frei0rFilterNdviFilter) GetVisOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("vis-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// A scaling factor to be applied to the visible component (divided by 10).
// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterNdviFilter) SetVisScale(value float64) error {
	return e.Element.SetProperty("vis-scale", value)
}

func (e *Frei0rFilterNdviFilter) GetVisScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("vis-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Color based alpha selection
type Frei0rFilterSelect0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterSelect0r() (*Frei0rFilterSelect0r, error) {
	e, err := gst.NewElement("frei0r-filter-select0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSelect0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSelect0rWithName(name string) (*Frei0rFilterSelect0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-select0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSelect0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetColorToSelectB(value float32) error {
	return e.Element.SetProperty("color-to-select-b", value)
}

func (e *Frei0rFilterSelect0r) GetColorToSelectB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-to-select-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Default: 0.8, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetColorToSelectG(value float32) error {
	return e.Element.SetProperty("color-to-select-g", value)
}

func (e *Frei0rFilterSelect0r) GetColorToSelectG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-to-select-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetColorToSelectR(value float32) error {
	return e.Element.SetProperty("color-to-select-r", value)
}

func (e *Frei0rFilterSelect0r) GetColorToSelectR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-to-select-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetDeltaBII(value float64) error {
	return e.Element.SetProperty("delta-b---i---i", value)
}

func (e *Frei0rFilterSelect0r) GetDeltaBII() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("delta-b---i---i")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: false
func (e *Frei0rFilterSelect0r) SetInvertSelection(value bool) error {
	return e.Element.SetProperty("invert-selection", value)
}

func (e *Frei0rFilterSelect0r) GetInvertSelection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert-selection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetSelectionSubspace(value float64) error {
	return e.Element.SetProperty("selection-subspace", value)
}

func (e *Frei0rFilterSelect0r) GetSelectionSubspace() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("selection-subspace")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetSlope(value float64) error {
	return e.Element.SetProperty("slope", value)
}

func (e *Frei0rFilterSelect0r) GetSlope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("slope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetDeltaGBChroma(value float64) error {
	return e.Element.SetProperty("delta-g---b---chroma", value)
}

func (e *Frei0rFilterSelect0r) GetDeltaGBChroma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("delta-g---b---chroma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetDeltaRAHue(value float64) error {
	return e.Element.SetProperty("delta-r---a---hue", value)
}

func (e *Frei0rFilterSelect0r) GetDeltaRAHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("delta-r---a---hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetEdgeMode(value float64) error {
	return e.Element.SetProperty("edge-mode", value)
}

func (e *Frei0rFilterSelect0r) GetEdgeMode() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("edge-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

func (e *Frei0rFilterSelect0r) GetOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSelect0r) SetSubspaceShape(value float64) error {
	return e.Element.SetProperty("subspace-shape", value)
}

func (e *Frei0rFilterSelect0r) GetSubspaceShape() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("subspace-shape")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Creates a Glamorous Glow
type Frei0rFilterGlow struct {
	*GstVideoFilter
}

func NewFrei0rFilterGlow() (*Frei0rFilterGlow, error) {
	e, err := gst.NewElement("frei0r-filter-glow")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterGlow{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterGlowWithName(name string) (*Frei0rFilterGlow, error) {
	e, err := gst.NewElementWithName("frei0r-filter-glow", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterGlow{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Blur of the glow
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterGlow) SetBlur(value float64) error {
	return e.Element.SetProperty("blur", value)
}

func (e *Frei0rFilterGlow) GetBlur() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blur")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// flushes frames in time in a nervous way
type Frei0rFilterNervous struct {
	*GstVideoFilter
}

func NewFrei0rFilterNervous() (*Frei0rFilterNervous, error) {
	e, err := gst.NewElement("frei0r-filter-nervous")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNervous{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterNervousWithName(name string) (*Frei0rFilterNervous, error) {
	e, err := gst.NewElementWithName("frei0r-filter-nervous", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNervous{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// High quality 3D denoiser from Mplayer
type Frei0rFilterHqdn3d struct {
	*GstVideoFilter
}

func NewFrei0rFilterHqdn3d() (*Frei0rFilterHqdn3d, error) {
	e, err := gst.NewElement("frei0r-filter-hqdn3d")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterHqdn3d{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterHqdn3dWithName(name string) (*Frei0rFilterHqdn3d, error) {
	e, err := gst.NewElementWithName("frei0r-filter-hqdn3d", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterHqdn3d{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Amount of temporal filtering
// Default: 0.06, Min: 0, Max: 1
func (e *Frei0rFilterHqdn3d) SetTemporal(value float64) error {
	return e.Element.SetProperty("temporal", value)
}

func (e *Frei0rFilterHqdn3d) GetTemporal() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("temporal")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount of spatial filtering
// Default: 0.04, Min: 0, Max: 1
func (e *Frei0rFilterHqdn3d) SetSpatial(value float64) error {
	return e.Element.SetProperty("spatial", value)
}

func (e *Frei0rFilterHqdn3d) GetSpatial() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("spatial")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform an RGB[A] softlight operation between the pixel sources.
type Frei0rMixerSoftlight struct {
	*gst.Element
}

func NewFrei0rMixerSoftlight() (*Frei0rMixerSoftlight, error) {
	e, err := gst.NewElement("frei0r-mixer-softlight")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerSoftlight{Element: e}, nil
}

func NewFrei0rMixerSoftlightWithName(name string) (*Frei0rMixerSoftlight, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-softlight", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerSoftlight{Element: e}, nil
}

// Four corners geometry engine
type Frei0rFilterC0rners struct {
	*GstVideoFilter
}

func NewFrei0rFilterC0rners() (*Frei0rFilterC0rners, error) {
	e, err := gst.NewElement("frei0r-filter-c0rners")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterC0rners{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterC0rnersWithName(name string) (*Frei0rFilterC0rners, error) {
	e, err := gst.NewElementWithName("frei0r-filter-c0rners", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterC0rners{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Y coordinate of corner 2
// Default: 0.333333, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner2Y(value float64) error {
	return e.Element.SetProperty("corner-2-y", value)
}

func (e *Frei0rFilterC0rners) GetCorner2Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-2-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Makes background transparent
// Default: false
func (e *Frei0rFilterC0rners) SetTransparentBackground(value bool) error {
	return e.Element.SetProperty("transparent-background", value)
}

func (e *Frei0rFilterC0rners) GetTransparentBackground() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("transparent-background")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetAlphaOperation(value float64) error {
	return e.Element.SetProperty("alpha-operation", value)
}

func (e *Frei0rFilterC0rners) GetAlphaOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha-operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X coordinate of corner 1
// Default: 0.333333, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner1X(value float64) error {
	return e.Element.SetProperty("corner-1-x", value)
}

func (e *Frei0rFilterC0rners) GetCorner1X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-1-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X coordinate of corner 3
// Default: 0.666666, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner3X(value float64) error {
	return e.Element.SetProperty("corner-3-x", value)
}

func (e *Frei0rFilterC0rners) GetCorner3X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-3-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X coordinate of corner 4
// Default: 0.333333, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner4X(value float64) error {
	return e.Element.SetProperty("corner-4-x", value)
}

func (e *Frei0rFilterC0rners) GetCorner4X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-4-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Enable stretching
// Default: false
func (e *Frei0rFilterC0rners) SetEnableStretch(value bool) error {
	return e.Element.SetProperty("enable-stretch", value)
}

func (e *Frei0rFilterC0rners) GetEnableStretch() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-stretch")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Quality of interpolation
// Default: 0.166667, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetInterpolator(value float64) error {
	return e.Element.SetProperty("interpolator", value)
}

func (e *Frei0rFilterC0rners) GetInterpolator() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("interpolator")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount of stretching in X direction
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetStretchX(value float64) error {
	return e.Element.SetProperty("stretch-x", value)
}

func (e *Frei0rFilterC0rners) GetStretchX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("stretch-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y coordinate of corner 1
// Default: 0.333333, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner1Y(value float64) error {
	return e.Element.SetProperty("corner-1-y", value)
}

func (e *Frei0rFilterC0rners) GetCorner1Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-1-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y coordinate of corner 4
// Default: 0.666666, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner4Y(value float64) error {
	return e.Element.SetProperty("corner-4-y", value)
}

func (e *Frei0rFilterC0rners) GetCorner4Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-4-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Makes smooth transition into transparent
// Default: 0.01, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetFeatherAlpha(value float64) error {
	return e.Element.SetProperty("feather-alpha", value)
}

func (e *Frei0rFilterC0rners) GetFeatherAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("feather-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X coordinate of corner 2
// Default: 0.666666, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner2X(value float64) error {
	return e.Element.SetProperty("corner-2-x", value)
}

func (e *Frei0rFilterC0rners) GetCorner2X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-2-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y coordinate of corner 3
// Default: 0.666666, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetCorner3Y(value float64) error {
	return e.Element.SetProperty("corner-3-y", value)
}

func (e *Frei0rFilterC0rners) GetCorner3Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-3-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount of stretching in Y direction
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterC0rners) SetStretchY(value float64) error {
	return e.Element.SetProperty("stretch-y", value)
}

func (e *Frei0rFilterC0rners) GetStretchY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("stretch-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adjust luminance or color channel intensity
type Frei0rFilterLevels struct {
	*GstVideoFilter
}

func NewFrei0rFilterLevels() (*Frei0rFilterLevels, error) {
	e, err := gst.NewElement("frei0r-filter-levels")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLevels{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterLevelsWithName(name string) (*Frei0rFilterLevels, error) {
	e, err := gst.NewElementWithName("frei0r-filter-levels", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLevels{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Histogram position. 0%%=TL, 10%%=TR, 20%%=BL, 30%%=BR
// Default: 0.3, Min: 0, Max: 1
func (e *Frei0rFilterLevels) SetHistogramPosition(value float64) error {
	return e.Element.SetProperty("histogram-position", value)
}

func (e *Frei0rFilterLevels) GetHistogramPosition() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("histogram-position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Input black level
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLevels) SetInputBlackLevel(value float64) error {
	return e.Element.SetProperty("input-black-level", value)
}

func (e *Frei0rFilterLevels) GetInputBlackLevel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("input-black-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Input white level
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterLevels) SetInputWhiteLevel(value float64) error {
	return e.Element.SetProperty("input-white-level", value)
}

func (e *Frei0rFilterLevels) GetInputWhiteLevel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("input-white-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Show histogram
// Default: true
func (e *Frei0rFilterLevels) SetShowHistogram(value bool) error {
	return e.Element.SetProperty("show-histogram", value)
}

func (e *Frei0rFilterLevels) GetShowHistogram() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-histogram")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// White output
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterLevels) SetWhiteOutput(value float64) error {
	return e.Element.SetProperty("white-output", value)
}

func (e *Frei0rFilterLevels) GetWhiteOutput() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("white-output")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Black output
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLevels) SetBlackOutput(value float64) error {
	return e.Element.SetProperty("black-output", value)
}

func (e *Frei0rFilterLevels) GetBlackOutput() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("black-output")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Channel to adjust levels. 0%%=R, 10%%=G, 20%%=B, 30%%=Luma
// Default: 0.3, Min: 0, Max: 1
func (e *Frei0rFilterLevels) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

func (e *Frei0rFilterLevels) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Gamma
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterLevels) SetGamma(value float64) error {
	return e.Element.SetProperty("gamma", value)
}

func (e *Frei0rFilterLevels) GetGamma() (float64, error) {
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

// Implements several median-type filters
type Frei0rFilterMedians struct {
	*GstVideoFilter
}

func NewFrei0rFilterMedians() (*Frei0rFilterMedians, error) {
	e, err := gst.NewElement("frei0r-filter-medians")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterMedians{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterMediansWithName(name string) (*Frei0rFilterMedians, error) {
	e, err := gst.NewElementWithName("frei0r-filter-medians", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterMedians{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Size for 'var size' type filter
// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterMedians) SetSize(value float64) error {
	return e.Element.SetProperty("size", value)
}

func (e *Frei0rFilterMedians) GetSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Choose type of median: Cross5, Square3x3, Bilevel, Diamond3x3, Square5x5, Temp3, Temp5, ArceBI, ML3D, ML3dEX, VarSize
// Default: Square3x3
func (e *Frei0rFilterMedians) SetType(value string) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rFilterMedians) GetType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// 2D video oscilloscope
type Frei0rFilterPr0file struct {
	*GstVideoFilter
}

func NewFrei0rFilterPr0file() (*Frei0rFilterPr0file, error) {
	e, err := gst.NewElement("frei0r-filter-pr0file")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPr0file{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterPr0fileWithName(name string) (*Frei0rFilterPr0file, error) {
	e, err := gst.NewElementWithName("frei0r-filter-pr0file", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPr0file{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// rec 601 or rec 709
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetColor(value float64) error {
	return e.Element.SetProperty("color", value)
}

func (e *Frei0rFilterPr0file) GetColor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: false
func (e *Frei0rFilterPr0file) SetDisplayMinimum(value bool) error {
	return e.Element.SetProperty("display-minimum", value)
}

func (e *Frei0rFilterPr0file) GetDisplayMinimum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-minimum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Position of marker 2
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetMarker2(value float64) error {
	return e.Element.SetProperty("marker-2", value)
}

func (e *Frei0rFilterPr0file) GetMarker2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("marker-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Show B trace on scope
// Default: true
func (e *Frei0rFilterPr0file) SetBTrace(value bool) error {
	return e.Element.SetProperty("b-trace", value)
}

func (e *Frei0rFilterPr0file) GetBTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Channel to numerically display
// Default: 0.375005, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

func (e *Frei0rFilterPr0file) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// e
// Default: true
func (e *Frei0rFilterPr0file) SetDisplayAverage(value bool) error {
	return e.Element.SetProperty("display-average", value)
}

func (e *Frei0rFilterPr0file) GetDisplayAverage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-average")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: true
func (e *Frei0rFilterPr0file) SetDisplayRms(value bool) error {
	return e.Element.SetProperty("display-rms", value)
}

func (e *Frei0rFilterPr0file) GetDisplayRms() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-rms")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Tilt of profile
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

func (e *Frei0rFilterPr0file) GetTilt() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X position of profile
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetX(value float64) error {
	return e.Element.SetProperty("x", value)
}

func (e *Frei0rFilterPr0file) GetX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y position of profile
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetY(value float64) error {
	return e.Element.SetProperty("y", value)
}

func (e *Frei0rFilterPr0file) GetY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Show Alpha trace on scope
// Default: false
func (e *Frei0rFilterPr0file) SetAlphaTrace(value bool) error {
	return e.Element.SetProperty("alpha-trace", value)
}

func (e *Frei0rFilterPr0file) GetAlphaTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("alpha-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Color of the profile marker
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetCrosshairColor(value float64) error {
	return e.Element.SetProperty("crosshair-color", value)
}

func (e *Frei0rFilterPr0file) GetCrosshairColor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("crosshair-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// use 0-255 instead of 0.0-1.0
// Default: false
func (e *Frei0rFilterPr0file) SetParam256Scale(value bool) error {
	return e.Element.SetProperty("param-256-scale", value)
}

func (e *Frei0rFilterPr0file) GetParam256Scale() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("param-256-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Show R trace on scope
// Default: true
func (e *Frei0rFilterPr0file) SetRTrace(value bool) error {
	return e.Element.SetProperty("r-trace", value)
}

func (e *Frei0rFilterPr0file) GetRTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("r-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Show Y' trace on scope
// Default: false
func (e *Frei0rFilterPr0file) SetYTrace(value bool) error {
	return e.Element.SetProperty("y-trace", value)
}

func (e *Frei0rFilterPr0file) GetYTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("y-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Show G trace on scope
// Default: true
func (e *Frei0rFilterPr0file) SetGTrace(value bool) error {
	return e.Element.SetProperty("g-trace", value)
}

func (e *Frei0rFilterPr0file) GetGTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("g-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Position of marker 1
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetMarker1(value float64) error {
	return e.Element.SetProperty("marker-1", value)
}

func (e *Frei0rFilterPr0file) GetMarker1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("marker-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Show Pb trace on scope
// Default: false
func (e *Frei0rFilterPr0file) SetPbTrace(value bool) error {
	return e.Element.SetProperty("pb-trace", value)
}

func (e *Frei0rFilterPr0file) GetPbTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pb-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Show Pr trace on scope
// Default: false
func (e *Frei0rFilterPr0file) SetPrTrace(value bool) error {
	return e.Element.SetProperty("pr-trace", value)
}

func (e *Frei0rFilterPr0file) GetPrTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pr-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default: false
func (e *Frei0rFilterPr0file) SetDisplayMaximum(value bool) error {
	return e.Element.SetProperty("display-maximum", value)
}

func (e *Frei0rFilterPr0file) GetDisplayMaximum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-maximum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Length of profile
// Default: 0.589744, Min: 0, Max: 1
func (e *Frei0rFilterPr0file) SetLength(value float64) error {
	return e.Element.SetProperty("length", value)
}

func (e *Frei0rFilterPr0file) GetLength() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("length")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Something videowall-ish
type Frei0rFilterTehroxx0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterTehroxx0r() (*Frei0rFilterTehroxx0r, error) {
	e, err := gst.NewElement("frei0r-filter-tehroxx0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTehroxx0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterTehroxx0rWithName(name string) (*Frei0rFilterTehroxx0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-tehroxx0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTehroxx0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Changing speed of small blocks
// Default: 0.01, Min: 0, Max: 1
func (e *Frei0rFilterTehroxx0r) SetInterval(value float64) error {
	return e.Element.SetProperty("interval", value)
}

func (e *Frei0rFilterTehroxx0r) GetInterval() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform an RGB[A] overlay operation between the pixel sources, using the generalised algorithm:
// D =  A * (B + (2 * B) * (255 - A))
type Frei0rMixerOverlay struct {
	*gst.Element
}

func NewFrei0rMixerOverlay() (*Frei0rMixerOverlay, error) {
	e, err := gst.NewElement("frei0r-mixer-overlay")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerOverlay{Element: e}, nil
}

func NewFrei0rMixerOverlayWithName(name string) (*Frei0rMixerOverlay, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-overlay", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerOverlay{Element: e}, nil
}

// Generates white noise images
type Frei0rSrcNois0r struct {
	*base.GstPushSrc
}

func NewFrei0rSrcNois0r() (*Frei0rSrcNois0r, error) {
	e, err := gst.NewElement("frei0r-src-nois0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcNois0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcNois0rWithName(name string) (*Frei0rSrcNois0r, error) {
	e, err := gst.NewElementWithName("frei0r-src-nois0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcNois0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Adds Black Borders at top and bottom for Cinema Look
type Frei0rFilterLetterb0xed struct {
	*GstVideoFilter
}

func NewFrei0rFilterLetterb0xed() (*Frei0rFilterLetterb0xed, error) {
	e, err := gst.NewElement("frei0r-filter-letterb0xed")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLetterb0xed{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterLetterb0xedWithName(name string) (*Frei0rFilterLetterb0xed, error) {
	e, err := gst.NewElementWithName("frei0r-filter-letterb0xed", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLetterb0xed{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0.4, Min: 0, Max: 1
func (e *Frei0rFilterLetterb0xed) SetBorderWidth(value float64) error {
	return e.Element.SetProperty("border-width", value)
}

func (e *Frei0rFilterLetterb0xed) GetBorderWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("border-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: false
func (e *Frei0rFilterLetterb0xed) SetTransparency(value bool) error {
	return e.Element.SetProperty("transparency", value)
}

func (e *Frei0rFilterLetterb0xed) GetTransparency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("transparency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adjusts the saturation of a source image
type Frei0rFilterSaturat0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterSaturat0r() (*Frei0rFilterSaturat0r, error) {
	e, err := gst.NewElement("frei0r-filter-saturat0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSaturat0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSaturat0rWithName(name string) (*Frei0rFilterSaturat0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-saturat0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSaturat0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The saturation value
// Default: 0.125, Min: 0, Max: 1
func (e *Frei0rFilterSaturat0r) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *Frei0rFilterSaturat0r) GetSaturation() (float64, error) {
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

// Perform an RGB[A] divide operation between the pixel sources: input1 is the numerator, input2 the denominator
type Frei0rMixerDivide struct {
	*gst.Element
}

func NewFrei0rMixerDivide() (*Frei0rMixerDivide, error) {
	e, err := gst.NewElement("frei0r-mixer-divide")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDivide{Element: e}, nil
}

func NewFrei0rMixerDivideWithName(name string) (*Frei0rMixerDivide, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-divide", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDivide{Element: e}, nil
}

// Adjusts the brightness of a source image
type Frei0rFilterBrightness struct {
	*GstVideoFilter
}

func NewFrei0rFilterBrightness() (*Frei0rFilterBrightness, error) {
	e, err := gst.NewElement("frei0r-filter-brightness")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBrightness{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterBrightnessWithName(name string) (*Frei0rFilterBrightness, error) {
	e, err := gst.NewElementWithName("frei0r-filter-brightness", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBrightness{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The brightness value
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterBrightness) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *Frei0rFilterBrightness) GetBrightness() (float64, error) {
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

// Extracts Red from Image
type Frei0rFilterR struct {
	*GstVideoFilter
}

func NewFrei0rFilterR() (*Frei0rFilterR, error) {
	e, err := gst.NewElement("frei0r-filter-r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterR{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterRWithName(name string) (*Frei0rFilterR, error) {
	e, err := gst.NewElementWithName("frei0r-filter-r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterR{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Unsharp masking (port from Mplayer)
type Frei0rFilterSharpness struct {
	*GstVideoFilter
}

func NewFrei0rFilterSharpness() (*Frei0rFilterSharpness, error) {
	e, err := gst.NewElement("frei0r-filter-sharpness")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSharpness{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSharpnessWithName(name string) (*Frei0rFilterSharpness, error) {
	e, err := gst.NewElementWithName("frei0r-filter-sharpness", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSharpness{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0.3, Min: 0, Max: 1
func (e *Frei0rFilterSharpness) SetAmount(value float64) error {
	return e.Element.SetProperty("amount", value)
}

func (e *Frei0rFilterSharpness) GetAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSharpness) SetSize(value float64) error {
	return e.Element.SetProperty("size", value)
}

func (e *Frei0rFilterSharpness) GetSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Slope/Offset/Power and Saturation color corrections according to the ASC CDL (Color Decision List)
type Frei0rFilterSopSat struct {
	*GstVideoFilter
}

func NewFrei0rFilterSopSat() (*Frei0rFilterSopSat, error) {
	e, err := gst.NewElement("frei0r-filter-sop-sat")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSopSat{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSopSatWithName(name string) (*Frei0rFilterSopSat, error) {
	e, err := gst.NewElementWithName("frei0r-filter-sop-sat", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSopSat{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Offset of the alpha component
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetAoffset(value float64) error {
	return e.Element.SetProperty("aoffset", value)
}

func (e *Frei0rFilterSopSat) GetAoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Power (Gamma) of the alpha component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetApower(value float64) error {
	return e.Element.SetProperty("apower", value)
}

func (e *Frei0rFilterSopSat) GetApower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("apower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Slope of the alpha component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetAslope(value float64) error {
	return e.Element.SetProperty("aslope", value)
}

func (e *Frei0rFilterSopSat) GetAslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Overall saturation
// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *Frei0rFilterSopSat) GetSaturation() (float64, error) {
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

// Offset of the red color component
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetRoffset(value float64) error {
	return e.Element.SetProperty("roffset", value)
}

func (e *Frei0rFilterSopSat) GetRoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("roffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Power (Gamma) of the red color component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetRpower(value float64) error {
	return e.Element.SetProperty("rpower", value)
}

func (e *Frei0rFilterSopSat) GetRpower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rpower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Offset of the blue color component
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetBoffset(value float64) error {
	return e.Element.SetProperty("boffset", value)
}

func (e *Frei0rFilterSopSat) GetBoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("boffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Power (Gamma) of the blue color component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetBpower(value float64) error {
	return e.Element.SetProperty("bpower", value)
}

func (e *Frei0rFilterSopSat) GetBpower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bpower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Slope of the blue color component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetBslope(value float64) error {
	return e.Element.SetProperty("bslope", value)
}

func (e *Frei0rFilterSopSat) GetBslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Offset of the green color component
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetGoffset(value float64) error {
	return e.Element.SetProperty("goffset", value)
}

func (e *Frei0rFilterSopSat) GetGoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("goffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Power (Gamma) of the green color component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetGpower(value float64) error {
	return e.Element.SetProperty("gpower", value)
}

func (e *Frei0rFilterSopSat) GetGpower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("gpower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Slope of the green color component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetGslope(value float64) error {
	return e.Element.SetProperty("gslope", value)
}

func (e *Frei0rFilterSopSat) GetGslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("gslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Slope of the red color component
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterSopSat) SetRslope(value float64) error {
	return e.Element.SetProperty("rslope", value)
}

func (e *Frei0rFilterSopSat) GetRslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Timeout indicators e.g. for slides.
type Frei0rFilterTimeoutIndicator struct {
	*GstVideoFilter
}

func NewFrei0rFilterTimeoutIndicator() (*Frei0rFilterTimeoutIndicator, error) {
	e, err := gst.NewElement("frei0r-filter-timeout-indicator")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTimeoutIndicator{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterTimeoutIndicatorWithName(name string) (*Frei0rFilterTimeoutIndicator, error) {
	e, err := gst.NewElementWithName("frei0r-filter-timeout-indicator", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTimeoutIndicator{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Indicator transparency
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTimeoutIndicator) SetTransparency(value float64) error {
	return e.Element.SetProperty("transparency", value)
}

func (e *Frei0rFilterTimeoutIndicator) GetTransparency() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("transparency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Indicator colour
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTimeoutIndicator) SetColorB(value float32) error {
	return e.Element.SetProperty("color-b", value)
}

func (e *Frei0rFilterTimeoutIndicator) GetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Indicator colour
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTimeoutIndicator) SetColorG(value float32) error {
	return e.Element.SetProperty("color-g", value)
}

func (e *Frei0rFilterTimeoutIndicator) GetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Indicator colour
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTimeoutIndicator) SetColorR(value float32) error {
	return e.Element.SetProperty("color-r", value)
}

func (e *Frei0rFilterTimeoutIndicator) GetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Current time
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTimeoutIndicator) SetTime(value float64) error {
	return e.Element.SetProperty("time", value)
}

func (e *Frei0rFilterTimeoutIndicator) GetTime() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("time")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Averages each Input and uses each as R, G or B channel of the Output
type Frei0rMixerRgb struct {
	*gst.Element
}

func NewFrei0rMixerRgb() (*Frei0rMixerRgb, error) {
	e, err := gst.NewElement("frei0r-mixer-rgb")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerRgb{Element: e}, nil
}

func NewFrei0rMixerRgbWithName(name string) (*Frei0rMixerRgb, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-rgb", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerRgb{Element: e}, nil
}

// Draws simple shapes into the alpha channel
type Frei0rFilterAlphaspot struct {
	*GstVideoFilter
}

func NewFrei0rFilterAlphaspot() (*Frei0rFilterAlphaspot, error) {
	e, err := gst.NewElement("frei0r-filter-alphaspot")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterAlphaspot{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterAlphaspotWithName(name string) (*Frei0rFilterAlphaspot, error) {
	e, err := gst.NewElementWithName("frei0r-filter-alphaspot", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterAlphaspot{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetSizeY(value float64) error {
	return e.Element.SetProperty("size-y", value)
}

func (e *Frei0rFilterAlphaspot) GetSizeY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetTransitionWidth(value float64) error {
	return e.Element.SetProperty("transition-width", value)
}

func (e *Frei0rFilterAlphaspot) GetTransitionWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("transition-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetMax(value float64) error {
	return e.Element.SetProperty("max", value)
}

func (e *Frei0rFilterAlphaspot) GetMax() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("max")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetMin(value float64) error {
	return e.Element.SetProperty("min", value)
}

func (e *Frei0rFilterAlphaspot) GetMin() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("min")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetShape(value float64) error {
	return e.Element.SetProperty("shape", value)
}

func (e *Frei0rFilterAlphaspot) GetShape() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shape")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetSizeX(value float64) error {
	return e.Element.SetProperty("size-x", value)
}

func (e *Frei0rFilterAlphaspot) GetSizeX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

func (e *Frei0rFilterAlphaspot) GetOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetPositionX(value float64) error {
	return e.Element.SetProperty("position-x", value)
}

func (e *Frei0rFilterAlphaspot) GetPositionX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("position-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetPositionY(value float64) error {
	return e.Element.SetProperty("position-y", value)
}

func (e *Frei0rFilterAlphaspot) GetPositionY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("position-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterAlphaspot) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

func (e *Frei0rFilterAlphaspot) GetTilt() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Measure video values
type Frei0rFilterPr0be struct {
	*GstVideoFilter
}

func NewFrei0rFilterPr0be() (*Frei0rFilterPr0be, error) {
	e, err := gst.NewElement("frei0r-filter-pr0be")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPr0be{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterPr0beWithName(name string) (*Frei0rFilterPr0be, error) {
	e, err := gst.NewElementWithName("frei0r-filter-pr0be", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPr0be{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Y size of probe
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterPr0be) SetYSize(value float64) error {
	return e.Element.SetProperty("y-size", value)
}

func (e *Frei0rFilterPr0be) GetYSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Display more data
// Default: false
func (e *Frei0rFilterPr0be) SetBigWindow(value bool) error {
	return e.Element.SetProperty("big-window", value)
}

func (e *Frei0rFilterPr0be) GetBigWindow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("big-window")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// What measurement to display
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterPr0be) SetMeasurement(value float64) error {
	return e.Element.SetProperty("measurement", value)
}

func (e *Frei0rFilterPr0be) GetMeasurement() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("measurement")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// use 0-255 instead of 0.0-1.0
// Default: false
func (e *Frei0rFilterPr0be) SetParam256Scale(value bool) error {
	return e.Element.SetProperty("param-256-scale", value)
}

func (e *Frei0rFilterPr0be) GetParam256Scale() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("param-256-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Display alpha value too
// Default: false
func (e *Frei0rFilterPr0be) SetShowAlpha(value bool) error {
	return e.Element.SetProperty("show-alpha", value)
}

func (e *Frei0rFilterPr0be) GetShowAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// X position of probe
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterPr0be) SetX(value float64) error {
	return e.Element.SetProperty("x", value)
}

func (e *Frei0rFilterPr0be) GetX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X size of probe
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterPr0be) SetXSize(value float64) error {
	return e.Element.SetProperty("x-size", value)
}

func (e *Frei0rFilterPr0be) GetXSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y position of probe
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterPr0be) SetY(value float64) error {
	return e.Element.SetProperty("y", value)
}

func (e *Frei0rFilterPr0be) GetY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Does softglow effect on highlights
type Frei0rFilterSoftglow struct {
	*GstVideoFilter
}

func NewFrei0rFilterSoftglow() (*Frei0rFilterSoftglow, error) {
	e, err := gst.NewElement("frei0r-filter-softglow")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSoftglow{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSoftglowWithName(name string) (*Frei0rFilterSoftglow, error) {
	e, err := gst.NewElementWithName("frei0r-filter-softglow", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSoftglow{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Brightness of highlight areas
// Default: 0.75, Min: 0, Max: 1
func (e *Frei0rFilterSoftglow) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *Frei0rFilterSoftglow) GetBrightness() (float64, error) {
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

// Sharpness of highlight areas
// Default: 0.85, Min: 0, Max: 1
func (e *Frei0rFilterSoftglow) SetSharpness(value float64) error {
	return e.Element.SetProperty("sharpness", value)
}

func (e *Frei0rFilterSoftglow) GetSharpness() (float64, error) {
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

// Blur of the glow
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterSoftglow) SetBlur(value float64) error {
	return e.Element.SetProperty("blur", value)
}

func (e *Frei0rFilterSoftglow) GetBlur() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blur")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Blend mode used to blend highlight blur with input image
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterSoftglow) SetBlurblend(value float64) error {
	return e.Element.SetProperty("blurblend", value)
}

func (e *Frei0rFilterSoftglow) GetBlurblend() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blurblend")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform an RGB[A] difference operation between the pixel sources.
type Frei0rMixerDifference struct {
	*gst.Element
}

func NewFrei0rMixerDifference() (*Frei0rMixerDifference, error) {
	e, err := gst.NewElement("frei0r-mixer-difference")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDifference{Element: e}, nil
}

func NewFrei0rMixerDifferenceWithName(name string) (*Frei0rMixerDifference, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-difference", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDifference{Element: e}, nil
}

// Bluescreen the background of a static video.
type Frei0rFilterBgsubtract0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterBgsubtract0r() (*Frei0rFilterBgsubtract0r, error) {
	e, err := gst.NewElement("frei0r-filter-bgsubtract0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBgsubtract0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterBgsubtract0rWithName(name string) (*Frei0rFilterBgsubtract0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-bgsubtract0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBgsubtract0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Blur alpha channel by given radius (to remove sharp edges)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterBgsubtract0r) SetBlur(value float64) error {
	return e.Element.SetProperty("blur", value)
}

func (e *Frei0rFilterBgsubtract0r) GetBlur() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blur")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Remove noise
// Default: true
func (e *Frei0rFilterBgsubtract0r) SetDenoise(value bool) error {
	return e.Element.SetProperty("denoise", value)
}

func (e *Frei0rFilterBgsubtract0r) GetDenoise() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("denoise")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Threshold for difference
// Default: 0.101961, Min: 0, Max: 1
func (e *Frei0rFilterBgsubtract0r) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *Frei0rFilterBgsubtract0r) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adjust luminance or color channel intensity with curve level mapping
type Frei0rFilterCurves struct {
	*GstVideoFilter
}

func NewFrei0rFilterCurves() (*Frei0rFilterCurves, error) {
	e, err := gst.NewElement("frei0r-filter-curves")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCurves{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterCurvesWithName(name string) (*Frei0rFilterCurves, error) {
	e, err := gst.NewElementWithName("frei0r-filter-curves", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterCurves{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Output image corner where curve graph will be drawn (0.1 = TOP,LEFT; 0.2 = TOP,RIGHT; 0.3 = BOTTOM,LEFT; 0.4 = BOTTOM, RIGHT)
// Default: 0.3, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetGraphPosition(value float64) error {
	return e.Element.SetProperty("graph-position", value)
}

func (e *Frei0rFilterCurves) GetGraphPosition() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("graph-position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 4 input value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint4InputValue(value float64) error {
	return e.Element.SetProperty("point-4-input-value", value)
}

func (e *Frei0rFilterCurves) GetPoint4InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-4-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 4 output value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint4OutputValue(value float64) error {
	return e.Element.SetProperty("point-4-output-value", value)
}

func (e *Frei0rFilterCurves) GetPoint4OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-4-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 3 input value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint3InputValue(value float64) error {
	return e.Element.SetProperty("point-3-input-value", value)
}

func (e *Frei0rFilterCurves) GetPoint3InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-3-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 3 output value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint3OutputValue(value float64) error {
	return e.Element.SetProperty("point-3-output-value", value)
}

func (e *Frei0rFilterCurves) GetPoint3OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-3-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Use cubic Bézier spline. Has to be a sorted list of points in the format "handle1x;handle1y#pointx;pointy#handle2x;handle2y"(pointx = in, pointy = out). Points are separated by a "|".The values can have "double" precision. x, y for points should be in the range 0-1. x,y for handles might also be out of this range.

func (e *Frei0rFilterCurves) SetBZierSpline(value string) error {
	return e.Element.SetProperty("b--zier-spline", value)
}

func (e *Frei0rFilterCurves) GetBZierSpline() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("b--zier-spline")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Number of points to use to build curve (/10 to fit [0,1] parameter range). Minimum 2 (0.2), Maximum 5 (0.5). Not relevant for Bézier spline.
// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetCurvePointNumber(value float64) error {
	return e.Element.SetProperty("curve-point-number", value)
}

func (e *Frei0rFilterCurves) GetCurvePointNumber() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("curve-point-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 1 input value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint1InputValue(value float64) error {
	return e.Element.SetProperty("point-1-input-value", value)
}

func (e *Frei0rFilterCurves) GetPoint1InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-1-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 2 input value
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint2InputValue(value float64) error {
	return e.Element.SetProperty("point-2-input-value", value)
}

func (e *Frei0rFilterCurves) GetPoint2InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-2-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 2 output value
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint2OutputValue(value float64) error {
	return e.Element.SetProperty("point-2-output-value", value)
}

func (e *Frei0rFilterCurves) GetPoint2OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-2-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Use Rec. 601 (false) or Rec. 709 (true)
// Default: true
func (e *Frei0rFilterCurves) SetLumaFormula(value bool) error {
	return e.Element.SetProperty("luma-formula", value)
}

func (e *Frei0rFilterCurves) GetLumaFormula() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("luma-formula")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Point 5 input value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint5InputValue(value float64) error {
	return e.Element.SetProperty("point-5-input-value", value)
}

func (e *Frei0rFilterCurves) GetPoint5InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-5-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Channel to adjust (0 = red, 0.1 = green, 0.2 = blue, 0.3 = alpha, 0.4 = luma, 0.5 = rgb, 0.6 = hue, 0.7 = saturation)
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

func (e *Frei0rFilterCurves) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 1 output value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint1OutputValue(value float64) error {
	return e.Element.SetProperty("point-1-output-value", value)
}

func (e *Frei0rFilterCurves) GetPoint1OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-1-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Point 5 output value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterCurves) SetPoint5OutputValue(value float64) error {
	return e.Element.SetProperty("point-5-output-value", value)
}

func (e *Frei0rFilterCurves) GetPoint5OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-5-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Draw curve graph on output image
// Default: true
func (e *Frei0rFilterCurves) SetShowCurves(value bool) error {
	return e.Element.SetProperty("show-curves", value)
}

func (e *Frei0rFilterCurves) GetShowCurves() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-curves")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// video delay
type Frei0rFilterDelay0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterDelay0r() (*Frei0rFilterDelay0r, error) {
	e, err := gst.NewElement("frei0r-filter-delay0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDelay0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterDelay0rWithName(name string) (*Frei0rFilterDelay0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-delay0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDelay0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// the delay time
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterDelay0r) SetDelaytime(value float64) error {
	return e.Element.SetProperty("delaytime", value)
}

func (e *Frei0rFilterDelay0r) GetDelaytime() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("delaytime")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Displays a histogram of R, G and B of the video-data
type Frei0rFilterRgbParade struct {
	*GstVideoFilter
}

func NewFrei0rFilterRgbParade() (*Frei0rFilterRgbParade, error) {
	e, err := gst.NewElement("frei0r-filter-rgb-parade")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterRgbParade{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterRgbParadeWithName(name string) (*Frei0rFilterRgbParade, error) {
	e, err := gst.NewElementWithName("frei0r-filter-rgb-parade", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterRgbParade{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The amount of source image mixed into background of display
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterRgbParade) SetMix(value float64) error {
	return e.Element.SetProperty("mix", value)
}

func (e *Frei0rFilterRgbParade) GetMix() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("mix")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// If false, the sides of image are shown without overlay
// Default: true
func (e *Frei0rFilterRgbParade) SetOverlaySides(value bool) error {
	return e.Element.SetProperty("overlay-sides", value)
}

func (e *Frei0rFilterRgbParade) GetOverlaySides() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("overlay-sides")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Desaturates image and creates a particular look that could be called Stamp, Newspaper or Photocopy
type Frei0rFilterSigmoidaltransfer struct {
	*GstVideoFilter
}

func NewFrei0rFilterSigmoidaltransfer() (*Frei0rFilterSigmoidaltransfer, error) {
	e, err := gst.NewElement("frei0r-filter-sigmoidaltransfer")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSigmoidaltransfer{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSigmoidaltransferWithName(name string) (*Frei0rFilterSigmoidaltransfer, error) {
	e, err := gst.NewElementWithName("frei0r-filter-sigmoidaltransfer", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSigmoidaltransfer{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Brightnesss of image
// Default: 0.75, Min: 0, Max: 1
func (e *Frei0rFilterSigmoidaltransfer) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *Frei0rFilterSigmoidaltransfer) GetBrightness() (float64, error) {
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

// Sharpness of transfer
// Default: 0.85, Min: 0, Max: 1
func (e *Frei0rFilterSigmoidaltransfer) SetSharpness(value float64) error {
	return e.Element.SetProperty("sharpness", value)
}

func (e *Frei0rFilterSigmoidaltransfer) GetSharpness() (float64, error) {
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

// a simple xfader
type Frei0rMixerXfade0r struct {
	*gst.Element
}

func NewFrei0rMixerXfade0r() (*Frei0rMixerXfade0r, error) {
	e, err := gst.NewElement("frei0r-mixer-xfade0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerXfade0r{Element: e}, nil
}

func NewFrei0rMixerXfade0rWithName(name string) (*Frei0rMixerXfade0r, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-xfade0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerXfade0r{Element: e}, nil
}

// the fader position
// Default: 0, Min: 0, Max: 1
func (e *Frei0rMixerXfade0r) SetFader(value float64) error {
	return e.Element.SetProperty("fader", value)
}

func (e *Frei0rMixerXfade0r) GetFader() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("fader")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Color to alpha (blit SRCALPHA)
type Frei0rFilterBluescreen0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterBluescreen0r() (*Frei0rFilterBluescreen0r, error) {
	e, err := gst.NewElement("frei0r-filter-bluescreen0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBluescreen0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterBluescreen0rWithName(name string) (*Frei0rFilterBluescreen0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-bluescreen0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterBluescreen0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The color to make transparent (B G R)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterBluescreen0r) SetColorB(value float32) error {
	return e.Element.SetProperty("color-b", value)
}

func (e *Frei0rFilterBluescreen0r) GetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color to make transparent (B G R)
// Default: 0.94, Min: 0, Max: 1
func (e *Frei0rFilterBluescreen0r) SetColorG(value float32) error {
	return e.Element.SetProperty("color-g", value)
}

func (e *Frei0rFilterBluescreen0r) GetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color to make transparent (B G R)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterBluescreen0r) SetColorR(value float32) error {
	return e.Element.SetProperty("color-r", value)
}

func (e *Frei0rFilterBluescreen0r) GetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Distance to Color (127 is good)
// Default: 0.288, Min: 0, Max: 1
func (e *Frei0rFilterBluescreen0r) SetDistance(value float64) error {
	return e.Element.SetProperty("distance", value)
}

func (e *Frei0rFilterBluescreen0r) GetDistance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Scales, Tilts and Crops an Image
type Frei0rFilterScale0tilt struct {
	*GstVideoFilter
}

func NewFrei0rFilterScale0tilt() (*Frei0rFilterScale0tilt, error) {
	e, err := gst.NewElement("frei0r-filter-scale0tilt")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterScale0tilt{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterScale0tiltWithName(name string) (*Frei0rFilterScale0tilt, error) {
	e, err := gst.NewElementWithName("frei0r-filter-scale0tilt", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterScale0tilt{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetTiltY(value float64) error {
	return e.Element.SetProperty("tilt-y", value)
}

func (e *Frei0rFilterScale0tilt) GetTiltY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetClipBottom(value float64) error {
	return e.Element.SetProperty("clip-bottom", value)
}

func (e *Frei0rFilterScale0tilt) GetClipBottom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetClipLeft(value float64) error {
	return e.Element.SetProperty("clip-left", value)
}

func (e *Frei0rFilterScale0tilt) GetClipLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetClipRight(value float64) error {
	return e.Element.SetProperty("clip-right", value)
}

func (e *Frei0rFilterScale0tilt) GetClipRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetClipTop(value float64) error {
	return e.Element.SetProperty("clip-top", value)
}

func (e *Frei0rFilterScale0tilt) GetClipTop() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetScaleX(value float64) error {
	return e.Element.SetProperty("scale-x", value)
}

func (e *Frei0rFilterScale0tilt) GetScaleX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetScaleY(value float64) error {
	return e.Element.SetProperty("scale-y", value)
}

func (e *Frei0rFilterScale0tilt) GetScaleY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterScale0tilt) SetTiltX(value float64) error {
	return e.Element.SetProperty("tilt-x", value)
}

func (e *Frei0rFilterScale0tilt) GetTiltX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// interlaced dark lines
type Frei0rFilterScanline0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterScanline0r() (*Frei0rFilterScanline0r, error) {
	e, err := gst.NewElement("frei0r-filter-scanline0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterScanline0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterScanline0rWithName(name string) (*Frei0rFilterScanline0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-scanline0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterScanline0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Composites Image 2 onto Image 1 according to its Alpha Channel
type Frei0rMixerComposition struct {
	*gst.Element
}

func NewFrei0rMixerComposition() (*Frei0rMixerComposition, error) {
	e, err := gst.NewElement("frei0r-mixer-composition")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerComposition{Element: e}, nil
}

func NewFrei0rMixerCompositionWithName(name string) (*Frei0rMixerComposition, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-composition", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerComposition{Element: e}, nil
}

// Perform an RGB[A] subtract operation of the pixel source input2 from input1.
type Frei0rMixerSubtract struct {
	*gst.Element
}

func NewFrei0rMixerSubtract() (*Frei0rMixerSubtract, error) {
	e, err := gst.NewElement("frei0r-mixer-subtract")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerSubtract{Element: e}, nil
}

func NewFrei0rMixerSubtractWithName(name string) (*Frei0rMixerSubtract, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-subtract", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerSubtract{Element: e}, nil
}

// Filters image to resemble a halftone print in which tones are represented as variable sized dots
type Frei0rFilterColorhalftone struct {
	*GstVideoFilter
}

func NewFrei0rFilterColorhalftone() (*Frei0rFilterColorhalftone, error) {
	e, err := gst.NewElement("frei0r-filter-colorhalftone")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColorhalftone{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterColorhalftoneWithName(name string) (*Frei0rFilterColorhalftone, error) {
	e, err := gst.NewElementWithName("frei0r-filter-colorhalftone", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColorhalftone{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Cyan dots angle
// Default: 0.3, Min: 0, Max: 1
func (e *Frei0rFilterColorhalftone) SetCyanAngle(value float64) error {
	return e.Element.SetProperty("cyan-angle", value)
}

func (e *Frei0rFilterColorhalftone) GetCyanAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("cyan-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Halftone pattern dot size
// Default: 0.4, Min: 0, Max: 1
func (e *Frei0rFilterColorhalftone) SetDotRadius(value float64) error {
	return e.Element.SetProperty("dot-radius", value)
}

func (e *Frei0rFilterColorhalftone) GetDotRadius() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("dot-radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Magenta dots angle
// Default: 0.45, Min: 0, Max: 1
func (e *Frei0rFilterColorhalftone) SetMagentaAngle(value float64) error {
	return e.Element.SetProperty("magenta-angle", value)
}

func (e *Frei0rFilterColorhalftone) GetMagentaAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("magenta-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Yellow dots angle
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterColorhalftone) SetYellowAngle(value float64) error {
	return e.Element.SetProperty("yellow-angle", value)
}

func (e *Frei0rFilterColorhalftone) GetYellowAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("yellow-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Edgeglow filter
type Frei0rFilterEdgeglow struct {
	*GstVideoFilter
}

func NewFrei0rFilterEdgeglow() (*Frei0rFilterEdgeglow, error) {
	e, err := gst.NewElement("frei0r-filter-edgeglow")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterEdgeglow{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterEdgeglowWithName(name string) (*Frei0rFilterEdgeglow, error) {
	e, err := gst.NewElementWithName("frei0r-filter-edgeglow", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterEdgeglow{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// multiplier for downscaling non-edge brightness
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterEdgeglow) SetLredscale(value float64) error {
	return e.Element.SetProperty("lredscale", value)
}

func (e *Frei0rFilterEdgeglow) GetLredscale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("lredscale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// threshold for edge lightening
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterEdgeglow) SetLthresh(value float64) error {
	return e.Element.SetProperty("lthresh", value)
}

func (e *Frei0rFilterEdgeglow) GetLthresh() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("lthresh")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// multiplier for upscaling edge brightness
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterEdgeglow) SetLupscale(value float64) error {
	return e.Element.SetProperty("lupscale", value)
}

func (e *Frei0rFilterEdgeglow) GetLupscale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("lupscale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Colorizes image to selected hue, saturation and lightness
type Frei0rFilterColorize struct {
	*GstVideoFilter
}

func NewFrei0rFilterColorize() (*Frei0rFilterColorize, error) {
	e, err := gst.NewElement("frei0r-filter-colorize")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColorize{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterColorizeWithName(name string) (*Frei0rFilterColorize, error) {
	e, err := gst.NewElementWithName("frei0r-filter-colorize", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColorize{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Color shade of the colorized image
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColorize) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

func (e *Frei0rFilterColorize) GetHue() (float64, error) {
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

// Lightness of the colorized image
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColorize) SetLightness(value float64) error {
	return e.Element.SetProperty("lightness", value)
}

func (e *Frei0rFilterColorize) GetLightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("lightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount of color in the colorized image
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColorize) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *Frei0rFilterColorize) GetSaturation() (float64, error) {
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

// This is a frei0r filter which allows to scale video footage non-linearly.
type Frei0rFilterElasticScaleFilter struct {
	*GstVideoFilter
}

func NewFrei0rFilterElasticScaleFilter() (*Frei0rFilterElasticScaleFilter, error) {
	e, err := gst.NewElement("frei0r-filter-elastic-scale-filter")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterElasticScaleFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterElasticScaleFilterWithName(name string) (*Frei0rFilterElasticScaleFilter, error) {
	e, err := gst.NewElementWithName("frei0r-filter-elastic-scale-filter", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterElasticScaleFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Horizontal center position of the linear area
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterElasticScaleFilter) SetCenter(value float64) error {
	return e.Element.SetProperty("center", value)
}

func (e *Frei0rFilterElasticScaleFilter) GetCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount how much the linear area is scaled
// Default: 0.7, Min: 0, Max: 1
func (e *Frei0rFilterElasticScaleFilter) SetLinearScaleFactor(value float64) error {
	return e.Element.SetProperty("linear-scale-factor", value)
}

func (e *Frei0rFilterElasticScaleFilter) GetLinearScaleFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("linear-scale-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Width of the linear area
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterElasticScaleFilter) SetLinearWidth(value float64) error {
	return e.Element.SetProperty("linear-width", value)
}

func (e *Frei0rFilterElasticScaleFilter) GetLinearWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("linear-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Amount how much the outer left and outer right areas are scaled non linearly
// Default: 0.7125, Min: 0, Max: 1
func (e *Frei0rFilterElasticScaleFilter) SetNonLinearScaleFactor(value float64) error {
	return e.Element.SetProperty("non-linear-scale-factor", value)
}

func (e *Frei0rFilterElasticScaleFilter) GetNonLinearScaleFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("non-linear-scale-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Reduce image to primary colors
type Frei0rFilterPrimaries struct {
	*GstVideoFilter
}

func NewFrei0rFilterPrimaries() (*Frei0rFilterPrimaries, error) {
	e, err := gst.NewElement("frei0r-filter-primaries")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPrimaries{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterPrimariesWithName(name string) (*Frei0rFilterPrimaries, error) {
	e, err := gst.NewElementWithName("frei0r-filter-primaries", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPrimaries{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// influence of mean px value. > 32 = 0
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterPrimaries) SetFactor(value float64) error {
	return e.Element.SetProperty("factor", value)
}

func (e *Frei0rFilterPrimaries) GetFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// alpha blending with zoomed and rotated images
type Frei0rFilterVertigo struct {
	*GstVideoFilter
}

func NewFrei0rFilterVertigo() (*Frei0rFilterVertigo, error) {
	e, err := gst.NewElement("frei0r-filter-vertigo")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterVertigo{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterVertigoWithName(name string) (*Frei0rFilterVertigo, error) {
	e, err := gst.NewElementWithName("frei0r-filter-vertigo", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterVertigo{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Phase increment
// Default: 0.02, Min: 0, Max: 1
func (e *Frei0rFilterVertigo) SetPhaseincrement(value float64) error {
	return e.Element.SetProperty("phaseincrement", value)
}

func (e *Frei0rFilterVertigo) GetPhaseincrement() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("phaseincrement")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Zoomrate
// Default: 0.202, Min: 0, Max: 1
func (e *Frei0rFilterVertigo) SetZoomrate(value float64) error {
	return e.Element.SetProperty("zoomrate", value)
}

func (e *Frei0rFilterVertigo) GetZoomrate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zoomrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// the alpha XOR operation
type Frei0rMixerAlphaxor struct {
	*gst.Element
}

func NewFrei0rMixerAlphaxor() (*Frei0rMixerAlphaxor, error) {
	e, err := gst.NewElement("frei0r-mixer-alphaxor")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaxor{Element: e}, nil
}

func NewFrei0rMixerAlphaxorWithName(name string) (*Frei0rMixerAlphaxor, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-alphaxor", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaxor{Element: e}, nil
}

// Generates ising noise
type Frei0rSrcIsing0r struct {
	*base.GstPushSrc
}

func NewFrei0rSrcIsing0r() (*Frei0rSrcIsing0r, error) {
	e, err := gst.NewElement("frei0r-src-ising0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcIsing0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcIsing0rWithName(name string) (*Frei0rSrcIsing0r, error) {
	e, err := gst.NewElementWithName("frei0r-src-ising0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcIsing0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Border Growth
// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcIsing0r) SetBorderGrowth(value float64) error {
	return e.Element.SetProperty("border-growth", value)
}

func (e *Frei0rSrcIsing0r) GetBorderGrowth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("border-growth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Spontaneous Growth
// Default: 1, Min: 0, Max: 1
func (e *Frei0rSrcIsing0r) SetSpontaneousGrowth(value float64) error {
	return e.Element.SetProperty("spontaneous-growth", value)
}

func (e *Frei0rSrcIsing0r) GetSpontaneousGrowth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("spontaneous-growth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Noise Temperature
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcIsing0r) SetTemperature(value float64) error {
	return e.Element.SetProperty("temperature", value)
}

func (e *Frei0rSrcIsing0r) GetTemperature() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("temperature")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Frame rotation in 3d-space
type Frei0rFilter3dflippo struct {
	*GstVideoFilter
}

func NewFrei0rFilter3dflippo() (*Frei0rFilter3dflippo, error) {
	e, err := gst.NewElement("frei0r-filter-3dflippo")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilter3dflippo{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilter3dflippoWithName(name string) (*Frei0rFilter3dflippo, error) {
	e, err := gst.NewElementWithName("frei0r-filter-3dflippo", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilter3dflippo{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Position of the center of rotation on the X axis
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetCenterPositionX(value float64) error {
	return e.Element.SetProperty("center-position--x-", value)
}

func (e *Frei0rFilter3dflippo) GetCenterPositionX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("center-position--x-")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// If true, pixels that are not transposed are black, otherwise, they are copied with the original
// Default: false
func (e *Frei0rFilter3dflippo) SetFillWithImageOrBlack(value bool) error {
	return e.Element.SetProperty("fill-with-image-or-black", value)
}

func (e *Frei0rFilter3dflippo) GetFillWithImageOrBlack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fill-with-image-or-black")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Rotation on the X axis
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetXAxisRotation(value float64) error {
	return e.Element.SetProperty("x-axis-rotation", value)
}

func (e *Frei0rFilter3dflippo) GetXAxisRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-axis-rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Rotation rate on the X axis
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetXAxisRotationRate(value float64) error {
	return e.Element.SetProperty("x-axis-rotation-rate", value)
}

func (e *Frei0rFilter3dflippo) GetXAxisRotationRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-axis-rotation-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Rotation on the Z axis
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetZAxisRotation(value float64) error {
	return e.Element.SetProperty("z-axis-rotation", value)
}

func (e *Frei0rFilter3dflippo) GetZAxisRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("z-axis-rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Position of the center of rotation on the Y axis
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetCenterPositionY(value float64) error {
	return e.Element.SetProperty("center-position--y-", value)
}

func (e *Frei0rFilter3dflippo) GetCenterPositionY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("center-position--y-")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Mask for frame transposition is not blanked, so a trace of old transpositions is maintained
// Default: false
func (e *Frei0rFilter3dflippo) SetDonTBlankMask(value bool) error {
	return e.Element.SetProperty("don-t-blank-mask", value)
}

func (e *Frei0rFilter3dflippo) GetDonTBlankMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("don-t-blank-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// If true, when mapping rotation, make inverted (wrong) assignment
// Default: false
func (e *Frei0rFilter3dflippo) SetInvertRotationAssignment(value bool) error {
	return e.Element.SetProperty("invert-rotation-assignment", value)
}

func (e *Frei0rFilter3dflippo) GetInvertRotationAssignment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert-rotation-assignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Rotation on the Y axis
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetYAxisRotation(value float64) error {
	return e.Element.SetProperty("y-axis-rotation", value)
}

func (e *Frei0rFilter3dflippo) GetYAxisRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-axis-rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Rotation rate on the Y axis
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetYAxisRotationRate(value float64) error {
	return e.Element.SetProperty("y-axis-rotation-rate", value)
}

func (e *Frei0rFilter3dflippo) GetYAxisRotationRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-axis-rotation-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Rotation rate on the Z axis
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilter3dflippo) SetZAxisRotationRate(value float64) error {
	return e.Element.SetProperty("z-axis-rotation-rate", value)
}

func (e *Frei0rFilter3dflippo) GetZAxisRotationRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("z-axis-rotation-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Sobel filter
type Frei0rFilterSobel struct {
	*GstVideoFilter
}

func NewFrei0rFilterSobel() (*Frei0rFilterSobel, error) {
	e, err := gst.NewElement("frei0r-filter-sobel")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSobel{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterSobelWithName(name string) (*Frei0rFilterSobel, error) {
	e, err := gst.NewElementWithName("frei0r-filter-sobel", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterSobel{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Dithers the image and reduces the number of available colors
type Frei0rFilterDither struct {
	*GstVideoFilter
}

func NewFrei0rFilterDither() (*Frei0rFilterDither, error) {
	e, err := gst.NewElement("frei0r-filter-dither")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDither{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterDitherWithName(name string) (*Frei0rFilterDither, error) {
	e, err := gst.NewElementWithName("frei0r-filter-dither", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDither{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Id of matrix used for dithering
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterDither) SetMatrixid(value float64) error {
	return e.Element.SetProperty("matrixid", value)
}

func (e *Frei0rFilterDither) GetMatrixid() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("matrixid")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Number of values per channel
// Default: 0.104167, Min: 0, Max: 1
func (e *Frei0rFilterDither) SetLevels(value float64) error {
	return e.Element.SetProperty("levels", value)
}

func (e *Frei0rFilterDither) GetLevels() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("levels")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Creates a luminance map of the image
type Frei0rFilterLuminance struct {
	*GstVideoFilter
}

func NewFrei0rFilterLuminance() (*Frei0rFilterLuminance, error) {
	e, err := gst.NewElement("frei0r-filter-luminance")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLuminance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterLuminanceWithName(name string) (*Frei0rFilterLuminance, error) {
	e, err := gst.NewElementWithName("frei0r-filter-luminance", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLuminance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Removes the Stairstepping from Nikon D90 videos (720p only) by interpolation
type Frei0rFilterNikonD90StairsteppingFix struct {
	*GstVideoFilter
}

func NewFrei0rFilterNikonD90StairsteppingFix() (*Frei0rFilterNikonD90StairsteppingFix, error) {
	e, err := gst.NewElement("frei0r-filter-nikon-d90-stairstepping-fix")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNikonD90StairsteppingFix{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterNikonD90StairsteppingFixWithName(name string) (*Frei0rFilterNikonD90StairsteppingFix, error) {
	e, err := gst.NewElementWithName("frei0r-filter-nikon-d90-stairstepping-fix", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterNikonD90StairsteppingFix{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// dynamic 3 level thresholding
type Frei0rFilterThreelay0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterThreelay0r() (*Frei0rFilterThreelay0r, error) {
	e, err := gst.NewElement("frei0r-filter-threelay0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterThreelay0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterThreelay0rWithName(name string) (*Frei0rFilterThreelay0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-threelay0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterThreelay0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Generates Lissajous0r images
type Frei0rSrcLissajous0r struct {
	*base.GstPushSrc
}

func NewFrei0rSrcLissajous0r() (*Frei0rSrcLissajous0r, error) {
	e, err := gst.NewElement("frei0r-src-lissajous0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcLissajous0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcLissajous0rWithName(name string) (*Frei0rSrcLissajous0r, error) {
	e, err := gst.NewElementWithName("frei0r-src-lissajous0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcLissajous0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// x-ratio
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcLissajous0r) SetRatiox(value float64) error {
	return e.Element.SetProperty("ratiox", value)
}

func (e *Frei0rSrcLissajous0r) GetRatiox() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("ratiox")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// y-ratio
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcLissajous0r) SetRatioy(value float64) error {
	return e.Element.SetProperty("ratioy", value)
}

func (e *Frei0rSrcLissajous0r) GetRatioy() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("ratioy")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Inverts all colors of a source image
type Frei0rFilterInvert0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterInvert0r() (*Frei0rFilterInvert0r, error) {
	e, err := gst.NewElement("frei0r-filter-invert0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterInvert0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterInvert0rWithName(name string) (*Frei0rFilterInvert0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-invert0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterInvert0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Perform a blend operation between two sources
type Frei0rMixerBlend struct {
	*gst.Element
}

func NewFrei0rMixerBlend() (*Frei0rMixerBlend, error) {
	e, err := gst.NewElement("frei0r-mixer-blend")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerBlend{Element: e}, nil
}

func NewFrei0rMixerBlendWithName(name string) (*Frei0rMixerBlend, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-blend", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerBlend{Element: e}, nil
}

// blend factor
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rMixerBlend) SetBlend(value float64) error {
	return e.Element.SetProperty("blend", value)
}

func (e *Frei0rMixerBlend) GetBlend() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blend")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform an RGB[A] hardlight operation between the pixel sources
type Frei0rMixerHardlight struct {
	*gst.Element
}

func NewFrei0rMixerHardlight() (*Frei0rMixerHardlight, error) {
	e, err := gst.NewElement("frei0r-mixer-hardlight")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerHardlight{Element: e}, nil
}

func NewFrei0rMixerHardlightWithName(name string) (*Frei0rMixerHardlight, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-hardlight", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerHardlight{Element: e}, nil
}

// automatic face blur
type Frei0rFilterFacebl0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterFacebl0r() (*Frei0rFilterFacebl0r, error) {
	e, err := gst.NewElement("frei0r-filter-facebl0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterFacebl0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterFacebl0rWithName(name string) (*Frei0rFilterFacebl0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-facebl0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterFacebl0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Maximum object size in pixels, divided by 10000
// Default: 0.05, Min: 0, Max: 1
func (e *Frei0rFilterFacebl0r) SetLargest(value float64) error {
	return e.Element.SetProperty("largest", value)
}

func (e *Frei0rFilterFacebl0r) GetLargest() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("largest")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Minimum number of rectangles that makes up an object, divided by 100
// Default: 0.02, Min: 0, Max: 1
func (e *Frei0rFilterFacebl0r) SetNeighbors(value float64) error {
	return e.Element.SetProperty("neighbors", value)
}

func (e *Frei0rFilterFacebl0r) GetNeighbors() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("neighbors")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// How often to detect an object in number of frames, divided by 1000
// Default: 0.025, Min: 0, Max: 1
func (e *Frei0rFilterFacebl0r) SetRecheck(value float64) error {
	return e.Element.SetProperty("recheck", value)
}

func (e *Frei0rFilterFacebl0r) GetRecheck() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("recheck")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The search window scale factor, divided by 10
// Default: 0.12, Min: 0, Max: 1
func (e *Frei0rFilterFacebl0r) SetSearchScale(value float64) error {
	return e.Element.SetProperty("search-scale", value)
}

func (e *Frei0rFilterFacebl0r) GetSearchScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("search-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Minimum window size in pixels, divided by 1000
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterFacebl0r) SetSmallest(value float64) error {
	return e.Element.SetProperty("smallest", value)
}

func (e *Frei0rFilterFacebl0r) GetSmallest() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("smallest")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// How many threads to use divided by 100; 0 uses CPU count
// Default: 0.01, Min: 0, Max: 1
func (e *Frei0rFilterFacebl0r) SetThreads(value float64) error {
	return e.Element.SetProperty("threads", value)
}

func (e *Frei0rFilterFacebl0r) GetThreads() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Full path to the XML pattern model for recognition; look in /usr/share/opencv/haarcascades
// Default: /usr/share/opencv/haarcascades/haarcascade_frontalface_default.xml
func (e *Frei0rFilterFacebl0r) SetClassifier(value string) error {
	return e.Element.SetProperty("classifier", value)
}

func (e *Frei0rFilterFacebl0r) GetClassifier() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("classifier")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Draw a red ellipse around the object
// Default: false
func (e *Frei0rFilterFacebl0r) SetEllipse(value bool) error {
	return e.Element.SetProperty("ellipse", value)
}

func (e *Frei0rFilterFacebl0r) GetEllipse() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ellipse")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Tint a source image with specified color
type Frei0rFilterTint0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterTint0r() (*Frei0rFilterTint0r, error) {
	e, err := gst.NewElement("frei0r-filter-tint0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTint0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterTint0rWithName(name string) (*Frei0rFilterTint0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-tint0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTint0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The color to map source color with null luminance
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTint0r) SetMapBlackToB(value float32) error {
	return e.Element.SetProperty("map-black-to-b", value)
}

func (e *Frei0rFilterTint0r) GetMapBlackToB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-black-to-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color to map source color with null luminance
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTint0r) SetMapBlackToG(value float32) error {
	return e.Element.SetProperty("map-black-to-g", value)
}

func (e *Frei0rFilterTint0r) GetMapBlackToG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-black-to-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color to map source color with null luminance
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterTint0r) SetMapBlackToR(value float32) error {
	return e.Element.SetProperty("map-black-to-r", value)
}

func (e *Frei0rFilterTint0r) GetMapBlackToR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-black-to-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color to map source color with full luminance
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterTint0r) SetMapWhiteToB(value float32) error {
	return e.Element.SetProperty("map-white-to-b", value)
}

func (e *Frei0rFilterTint0r) GetMapWhiteToB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-white-to-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color to map source color with full luminance
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterTint0r) SetMapWhiteToG(value float32) error {
	return e.Element.SetProperty("map-white-to-g", value)
}

func (e *Frei0rFilterTint0r) GetMapWhiteToG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-white-to-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The color to map source color with full luminance
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterTint0r) SetMapWhiteToR(value float32) error {
	return e.Element.SetProperty("map-white-to-r", value)
}

func (e *Frei0rFilterTint0r) GetMapWhiteToR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-white-to-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Amount of color
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterTint0r) SetTintAmount(value float64) error {
	return e.Element.SetProperty("tint-amount", value)
}

func (e *Frei0rFilterTint0r) GetTintAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tint-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// dynamic thresholding
type Frei0rFilterTwolay0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterTwolay0r() (*Frei0rFilterTwolay0r, error) {
	e, err := gst.NewElement("frei0r-filter-twolay0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTwolay0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterTwolay0rWithName(name string) (*Frei0rFilterTwolay0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-twolay0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterTwolay0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Perform a conversion to hue only of the source input1 using the hue of input2.
type Frei0rMixerHue struct {
	*gst.Element
}

func NewFrei0rMixerHue() (*Frei0rMixerHue, error) {
	e, err := gst.NewElement("frei0r-mixer-hue")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerHue{Element: e}, nil
}

func NewFrei0rMixerHueWithName(name string) (*Frei0rMixerHue, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-hue", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerHue{Element: e}, nil
}

// image with just one color
type Frei0rSrcOnecol0r struct {
	*base.GstPushSrc
}

func NewFrei0rSrcOnecol0r() (*Frei0rSrcOnecol0r, error) {
	e, err := gst.NewElement("frei0r-src-onecol0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcOnecol0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcOnecol0rWithName(name string) (*Frei0rSrcOnecol0r, error) {
	e, err := gst.NewElementWithName("frei0r-src-onecol0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcOnecol0r{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// the color of the image
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcOnecol0r) SetColorB(value float32) error {
	return e.Element.SetProperty("color-b", value)
}

func (e *Frei0rSrcOnecol0r) GetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// the color of the image
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcOnecol0r) SetColorG(value float32) error {
	return e.Element.SetProperty("color-g", value)
}

func (e *Frei0rSrcOnecol0r) GetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// the color of the image
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcOnecol0r) SetColorR(value float32) error {
	return e.Element.SetProperty("color-r", value)
}

func (e *Frei0rSrcOnecol0r) GetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Posterizes image by reducing the number of colors used in image
type Frei0rFilterPosterize struct {
	*GstVideoFilter
}

func NewFrei0rFilterPosterize() (*Frei0rFilterPosterize, error) {
	e, err := gst.NewElement("frei0r-filter-posterize")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPosterize{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterPosterizeWithName(name string) (*Frei0rFilterPosterize, error) {
	e, err := gst.NewElementWithName("frei0r-filter-posterize", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPosterize{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Number of values per channel
// Default: 0.104167, Min: 0, Max: 1
func (e *Frei0rFilterPosterize) SetLevels(value float64) error {
	return e.Element.SetProperty("levels", value)
}

func (e *Frei0rFilterPosterize) GetLevels() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("levels")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Flipping in x and y axis
type Frei0rFilterFlippo struct {
	*GstVideoFilter
}

func NewFrei0rFilterFlippo() (*Frei0rFilterFlippo, error) {
	e, err := gst.NewElement("frei0r-filter-flippo")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterFlippo{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterFlippoWithName(name string) (*Frei0rFilterFlippo, error) {
	e, err := gst.NewElementWithName("frei0r-filter-flippo", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterFlippo{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Flipping on the horizontal axis
// Default: false
func (e *Frei0rFilterFlippo) SetXAxis(value bool) error {
	return e.Element.SetProperty("x-axis", value)
}

func (e *Frei0rFilterFlippo) GetXAxis() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("x-axis")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Flipping on the vertical axis
// Default: false
func (e *Frei0rFilterFlippo) SetYAxis(value bool) error {
	return e.Element.SetProperty("y-axis", value)
}

func (e *Frei0rFilterFlippo) GetYAxis() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("y-axis")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adjusts the gamma value of a source image
type Frei0rFilterGamma struct {
	*GstVideoFilter
}

func NewFrei0rFilterGamma() (*Frei0rFilterGamma, error) {
	e, err := gst.NewElement("frei0r-filter-gamma")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterGamma{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterGammaWithName(name string) (*Frei0rFilterGamma, error) {
	e, err := gst.NewElementWithName("frei0r-filter-gamma", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterGamma{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The gamma value
// Default: 1, Min: 0, Max: 1
func (e *Frei0rFilterGamma) SetGamma(value float64) error {
	return e.Element.SetProperty("gamma", value)
}

func (e *Frei0rFilterGamma) GetGamma() (float64, error) {
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

// Multiply (or divide) each color component by the pixel's alpha value
type Frei0rFilterPremultiplyOrUnpremultiply struct {
	*GstVideoFilter
}

func NewFrei0rFilterPremultiplyOrUnpremultiply() (*Frei0rFilterPremultiplyOrUnpremultiply, error) {
	e, err := gst.NewElement("frei0r-filter-premultiply-or-unpremultiply")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPremultiplyOrUnpremultiply{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterPremultiplyOrUnpremultiplyWithName(name string) (*Frei0rFilterPremultiplyOrUnpremultiply, error) {
	e, err := gst.NewElementWithName("frei0r-filter-premultiply-or-unpremultiply", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterPremultiplyOrUnpremultiply{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Whether to unpremultiply instead
// Default: true
func (e *Frei0rFilterPremultiplyOrUnpremultiply) SetUnpremultiply(value bool) error {
	return e.Element.SetProperty("unpremultiply", value)
}

func (e *Frei0rFilterPremultiplyOrUnpremultiply) GetUnpremultiply() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("unpremultiply")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Displays the vectorscope of the video-data
type Frei0rFilterVectorscope struct {
	*GstVideoFilter
}

func NewFrei0rFilterVectorscope() (*Frei0rFilterVectorscope, error) {
	e, err := gst.NewElement("frei0r-filter-vectorscope")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterVectorscope{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterVectorscopeWithName(name string) (*Frei0rFilterVectorscope, error) {
	e, err := gst.NewElementWithName("frei0r-filter-vectorscope", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterVectorscope{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The amount of source image mixed into background of display
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterVectorscope) SetMix(value float64) error {
	return e.Element.SetProperty("mix", value)
}

func (e *Frei0rFilterVectorscope) GetMix() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("mix")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// If false, the sides of image are shown without overlay
// Default: true
func (e *Frei0rFilterVectorscope) SetOverlaySides(value bool) error {
	return e.Element.SetProperty("overlay-sides", value)
}

func (e *Frei0rFilterVectorscope) GetOverlaySides() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("overlay-sides")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Perform an RGB[A] dodge operation between the pixel sources, using the generalised algorithm:
// D = saturation of 255 or (A * 256) / (256 - B)
type Frei0rMixerDodge struct {
	*gst.Element
}

func NewFrei0rMixerDodge() (*Frei0rMixerDodge, error) {
	e, err := gst.NewElement("frei0r-mixer-dodge")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDodge{Element: e}, nil
}

func NewFrei0rMixerDodgeWithName(name string) (*Frei0rMixerDodge, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-dodge", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerDodge{Element: e}, nil
}

// Perform a conversion to saturation only of the source input1 using the saturation level of input2.
type Frei0rMixerSaturation struct {
	*gst.Element
}

func NewFrei0rMixerSaturation() (*Frei0rMixerSaturation, error) {
	e, err := gst.NewElement("frei0r-mixer-saturation")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerSaturation{Element: e}, nil
}

func NewFrei0rMixerSaturationWithName(name string) (*Frei0rMixerSaturation, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-saturation", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerSaturation{Element: e}, nil
}

// Perform an RGB[A] screen operation between the pixel sources, using the generalised algorithm:
// D = 255 - (255 - A) * (255 - B)
type Frei0rMixerScreen struct {
	*gst.Element
}

func NewFrei0rMixerScreen() (*Frei0rMixerScreen, error) {
	e, err := gst.NewElement("frei0r-mixer-screen")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerScreen{Element: e}, nil
}

func NewFrei0rMixerScreenWithName(name string) (*Frei0rMixerScreen, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-screen", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerScreen{Element: e}, nil
}

// Calculates the distance between the selected color and the current pixel and uses that value as new pixel value
type Frei0rFilterColorDistance struct {
	*GstVideoFilter
}

func NewFrei0rFilterColorDistance() (*Frei0rFilterColorDistance, error) {
	e, err := gst.NewElement("frei0r-filter-color-distance")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColorDistance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterColorDistanceWithName(name string) (*Frei0rFilterColorDistance, error) {
	e, err := gst.NewElementWithName("frei0r-filter-color-distance", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterColorDistance{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The Source Color
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColorDistance) SetColorR(value float32) error {
	return e.Element.SetProperty("color-r", value)
}

func (e *Frei0rFilterColorDistance) GetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The Source Color
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColorDistance) SetColorB(value float32) error {
	return e.Element.SetProperty("color-b", value)
}

func (e *Frei0rFilterColorDistance) GetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The Source Color
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterColorDistance) SetColorG(value float32) error {
	return e.Element.SetProperty("color-g", value)
}

func (e *Frei0rFilterColorDistance) GetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Equalizes the intensity histograms
type Frei0rFilterEqualiz0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterEqualiz0r() (*Frei0rFilterEqualiz0r, error) {
	e, err := gst.NewElement("frei0r-filter-equaliz0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterEqualiz0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterEqualiz0rWithName(name string) (*Frei0rFilterEqualiz0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-equaliz0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterEqualiz0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Creates an square alpha-channel mask
type Frei0rFilterMask0mate struct {
	*GstVideoFilter
}

func NewFrei0rFilterMask0mate() (*Frei0rFilterMask0mate, error) {
	e, err := gst.NewElement("frei0r-filter-mask0mate")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterMask0mate{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterMask0mateWithName(name string) (*Frei0rFilterMask0mate, error) {
	e, err := gst.NewElementWithName("frei0r-filter-mask0mate", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterMask0mate{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterMask0mate) SetLeft(value float64) error {
	return e.Element.SetProperty("left", value)
}

func (e *Frei0rFilterMask0mate) GetLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterMask0mate) SetRight(value float64) error {
	return e.Element.SetProperty("right", value)
}

func (e *Frei0rFilterMask0mate) GetRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterMask0mate) SetTop(value float64) error {
	return e.Element.SetProperty("top", value)
}

func (e *Frei0rFilterMask0mate) GetTop() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Blur the outline of the mask
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterMask0mate) SetBlur(value float64) error {
	return e.Element.SetProperty("blur", value)
}

func (e *Frei0rFilterMask0mate) GetBlur() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blur")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterMask0mate) SetBottom(value float64) error {
	return e.Element.SetProperty("bottom", value)
}

func (e *Frei0rFilterMask0mate) GetBottom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Invert the mask, creates a hole in the frame.
// Default: false
func (e *Frei0rFilterMask0mate) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

func (e *Frei0rFilterMask0mate) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// the alpha OVER operation
type Frei0rMixerAlphaover struct {
	*gst.Element
}

func NewFrei0rMixerAlphaover() (*Frei0rMixerAlphaover, error) {
	e, err := gst.NewElement("frei0r-mixer-alphaover")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaover{Element: e}, nil
}

func NewFrei0rMixerAlphaoverWithName(name string) (*Frei0rMixerAlphaover, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-alphaover", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaover{Element: e}, nil
}

// Perform an RGB[A] dodge operation between the pixel sources, using the generalised algorithm:
// D = saturation of 255 or depletion of 0, of ((255 - A) * 256) / (B + 1)
type Frei0rMixerBurn struct {
	*gst.Element
}

func NewFrei0rMixerBurn() (*Frei0rMixerBurn, error) {
	e, err := gst.NewElement("frei0r-mixer-burn")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerBurn{Element: e}, nil
}

func NewFrei0rMixerBurnWithName(name string) (*Frei0rMixerBurn, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-burn", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerBurn{Element: e}, nil
}

// Perform an RGB[A] grain-extract operation between the pixel sources.
type Frei0rMixerGrainExtract struct {
	*gst.Element
}

func NewFrei0rMixerGrainExtract() (*Frei0rMixerGrainExtract, error) {
	e, err := gst.NewElement("frei0r-mixer-grain-extract")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerGrainExtract{Element: e}, nil
}

func NewFrei0rMixerGrainExtractWithName(name string) (*Frei0rMixerGrainExtract, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-grain-extract", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerGrainExtract{Element: e}, nil
}

// Extracts Blue from Image
type Frei0rFilterB struct {
	*GstVideoFilter
}

func NewFrei0rFilterB() (*Frei0rFilterB, error) {
	e, err := gst.NewElement("frei0r-filter-b")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterB{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterBWithName(name string) (*Frei0rFilterB, error) {
	e, err := gst.NewElementWithName("frei0r-filter-b", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterB{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// delayed frame blitting mapped on a time bitmap
type Frei0rFilterDelaygrab struct {
	*GstVideoFilter
}

func NewFrei0rFilterDelaygrab() (*Frei0rFilterDelaygrab, error) {
	e, err := gst.NewElement("frei0r-filter-delaygrab")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDelaygrab{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterDelaygrabWithName(name string) (*Frei0rFilterDelaygrab, error) {
	e, err := gst.NewElementWithName("frei0r-filter-delaygrab", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterDelaygrab{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Shifts the hue of a source image
type Frei0rFilterHueshift0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterHueshift0r() (*Frei0rFilterHueshift0r, error) {
	e, err := gst.NewElement("frei0r-filter-hueshift0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterHueshift0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterHueshift0rWithName(name string) (*Frei0rFilterHueshift0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-hueshift0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterHueshift0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The shift value
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterHueshift0r) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

func (e *Frei0rFilterHueshift0r) GetHue() (float64, error) {
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

// Reduces the visibility of key color spill in chroma keying
type Frei0rFilterKeyspillm0pup struct {
	*GstVideoFilter
}

func NewFrei0rFilterKeyspillm0pup() (*Frei0rFilterKeyspillm0pup, error) {
	e, err := gst.NewElement("frei0r-filter-keyspillm0pup")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterKeyspillm0pup{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterKeyspillm0pupWithName(name string) (*Frei0rFilterKeyspillm0pup, error) {
	e, err := gst.NewElementWithName("frei0r-filter-keyspillm0pup", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterKeyspillm0pup{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Which mask to apply [0,1,2,3]
// Default: 0
func (e *Frei0rFilterKeyspillm0pup) SetMaskType(value string) error {
	return e.Element.SetProperty("mask-type", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetMaskType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mask-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Restrict mask to saturated colors
// Default: 0.15, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetSaturationThreshold(value float64) error {
	return e.Element.SetProperty("saturation-threshold", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetSaturationThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Replace image with the mask
// Default: false
func (e *Frei0rFilterKeyspillm0pup) SetShowMask(value bool) error {
	return e.Element.SetProperty("show-mask", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetShowMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Desired color to replace key residue with
// Default: 0.4, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetTargetColorB(value float32) error {
	return e.Element.SetProperty("target-color-b", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetTargetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("target-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Desired color to replace key residue with
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetTargetColorG(value float32) error {
	return e.Element.SetProperty("target-color-g", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetTargetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("target-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Restrict mask to hues close to key
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetHueGate(value float64) error {
	return e.Element.SetProperty("hue-gate", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetHueGate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue-gate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Replace alpha channel with the mask
// Default: false
func (e *Frei0rFilterKeyspillm0pup) SetMaskToAlpha(value bool) error {
	return e.Element.SetProperty("mask-to-alpha", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetMaskToAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mask-to-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Key color that was used for chroma keying
// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetKeyColorR(value float32) error {
	return e.Element.SetProperty("key-color-r", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetKeyColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("key-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Second operation 2 [0,1,2]
// Default: 0
func (e *Frei0rFilterKeyspillm0pup) SetOperation2(value string) error {
	return e.Element.SetProperty("operation-2", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetOperation2() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("operation-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Default: 0.55, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetAmount1(value float64) error {
	return e.Element.SetProperty("amount-1", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetAmount1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Key color that was used for chroma keying
// Default: 0.1, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetKeyColorB(value float32) error {
	return e.Element.SetProperty("key-color-b", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetKeyColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("key-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Key color that was used for chroma keying
// Default: 0.8, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetKeyColorG(value float32) error {
	return e.Element.SetProperty("key-color-g", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetKeyColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("key-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Range of colors around the key, where effect is full strength
// Default: 0.24, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetTolerance(value float64) error {
	return e.Element.SetProperty("tolerance", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetTolerance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Range of colors around the key where effect gradually decreases
// Default: 0.4, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetSlope(value float64) error {
	return e.Element.SetProperty("slope", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetSlope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("slope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Desired color to replace key residue with
// Default: 0.78, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetTargetColorR(value float32) error {
	return e.Element.SetProperty("target-color-r", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetTargetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("target-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterKeyspillm0pup) SetAmount2(value float64) error {
	return e.Element.SetProperty("amount-2", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetAmount2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// First operation 1 [0,1,2]
// Default: 1
func (e *Frei0rFilterKeyspillm0pup) SetOperation1(value string) error {
	return e.Element.SetProperty("operation-1", value)
}

func (e *Frei0rFilterKeyspillm0pup) GetOperation1() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("operation-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Allows compensation of lens distortion
type Frei0rFilterLensCorrection struct {
	*GstVideoFilter
}

func NewFrei0rFilterLensCorrection() (*Frei0rFilterLensCorrection, error) {
	e, err := gst.NewElement("frei0r-filter-lens-correction")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLensCorrection{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterLensCorrectionWithName(name string) (*Frei0rFilterLensCorrection, error) {
	e, err := gst.NewElementWithName("frei0r-filter-lens-correction", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLensCorrection{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterLensCorrection) SetCorrectionNearCenter(value float64) error {
	return e.Element.SetProperty("correction-near-center", value)
}

func (e *Frei0rFilterLensCorrection) GetCorrectionNearCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("correction-near-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterLensCorrection) SetCorrectionNearEdges(value float64) error {
	return e.Element.SetProperty("correction-near-edges", value)
}

func (e *Frei0rFilterLensCorrection) GetCorrectionNearEdges() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("correction-near-edges")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterLensCorrection) SetXCenter(value float64) error {
	return e.Element.SetProperty("x-center", value)
}

func (e *Frei0rFilterLensCorrection) GetXCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterLensCorrection) SetYCenter(value float64) error {
	return e.Element.SetProperty("y-center", value)
}

func (e *Frei0rFilterLensCorrection) GetYCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterLensCorrection) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *Frei0rFilterLensCorrection) GetBrightness() (float64, error) {
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

// Generates test card lookalikes
type Frei0rSrcTestPatB struct {
	*base.GstPushSrc
}

func NewFrei0rSrcTestPatB() (*Frei0rSrcTestPatB, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-b")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatB{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcTestPatBWithName(name string) (*Frei0rSrcTestPatB, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-b", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatB{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// 8 choices, select test pattern
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatB) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rSrcTestPatB) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// 7 choices, pixel aspect ratio
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatB) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

func (e *Frei0rSrcTestPatB) GetAspectType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Manual pixel aspect ratio (Aspect type 6)
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rSrcTestPatB) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

func (e *Frei0rSrcTestPatB) GetManualAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adjusts the contrast of a source image
type Frei0rFilterContrast0r struct {
	*GstVideoFilter
}

func NewFrei0rFilterContrast0r() (*Frei0rFilterContrast0r, error) {
	e, err := gst.NewElement("frei0r-filter-contrast0r")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterContrast0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterContrast0rWithName(name string) (*Frei0rFilterContrast0r, error) {
	e, err := gst.NewElementWithName("frei0r-filter-contrast0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterContrast0r{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The contrast value
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rFilterContrast0r) SetContrast(value float64) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *Frei0rFilterContrast0r) GetContrast() (float64, error) {
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

// Perform an RGB[A] addition_alpha operation of the pixel sources.
type Frei0rMixerAdditionAlpha struct {
	*gst.Element
}

func NewFrei0rMixerAdditionAlpha() (*Frei0rMixerAdditionAlpha, error) {
	e, err := gst.NewElement("frei0r-mixer-addition-alpha")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAdditionAlpha{Element: e}, nil
}

func NewFrei0rMixerAdditionAlphaWithName(name string) (*Frei0rMixerAdditionAlpha, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-addition-alpha", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAdditionAlpha{Element: e}, nil
}

// the alpha ATOP operation
type Frei0rMixerAlphaatop struct {
	*gst.Element
}

func NewFrei0rMixerAlphaatop() (*Frei0rMixerAlphaatop, error) {
	e, err := gst.NewElement("frei0r-mixer-alphaatop")
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaatop{Element: e}, nil
}

func NewFrei0rMixerAlphaatopWithName(name string) (*Frei0rMixerAlphaatop, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-alphaatop", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rMixerAlphaatop{Element: e}, nil
}

// Generates geometry test pattern images
type Frei0rSrcTestPatG struct {
	*base.GstPushSrc
}

func NewFrei0rSrcTestPatG() (*Frei0rSrcTestPatG, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-g")
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatG{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFrei0rSrcTestPatGWithName(name string) (*Frei0rSrcTestPatG, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-g", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rSrcTestPatG{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Manual pixel aspect ratio
// Default: 0.5, Min: 0, Max: 1
func (e *Frei0rSrcTestPatG) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

func (e *Frei0rSrcTestPatG) GetManualAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Polarity of image
// Default: false
func (e *Frei0rSrcTestPatG) SetNegative(value bool) error {
	return e.Element.SetProperty("negative", value)
}

func (e *Frei0rSrcTestPatG) GetNegative() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("negative")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Size of major features
// Default: 0.28125, Min: 0, Max: 1
func (e *Frei0rSrcTestPatG) SetSize1(value float64) error {
	return e.Element.SetProperty("size-1", value)
}

func (e *Frei0rSrcTestPatG) GetSize1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Size of minor features
// Default: 0.0625, Min: 0, Max: 1
func (e *Frei0rSrcTestPatG) SetSize2(value float64) error {
	return e.Element.SetProperty("size-2", value)
}

func (e *Frei0rSrcTestPatG) GetSize2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Type of test pattern
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatG) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

func (e *Frei0rSrcTestPatG) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pixel aspect ratio presets
// Default: 0, Min: 0, Max: 1
func (e *Frei0rSrcTestPatG) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

func (e *Frei0rSrcTestPatG) GetAspectType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Creates light graffitis from a video by keeping the brightest spots.
type Frei0rFilterLightGraffiti struct {
	*GstVideoFilter
}

func NewFrei0rFilterLightGraffiti() (*Frei0rFilterLightGraffiti, error) {
	e, err := gst.NewElement("frei0r-filter-light-graffiti")
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLightGraffiti{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewFrei0rFilterLightGraffitiWithName(name string) (*Frei0rFilterLightGraffiti, error) {
	e, err := gst.NewElementWithName("frei0r-filter-light-graffiti", name)
	if err != nil {
		return nil, err
	}
	return &Frei0rFilterLightGraffiti{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Uses black as background image instead of the first frame.
// Default: false
func (e *Frei0rFilterLightGraffiti) SetBlackreference(value bool) error {
	return e.Element.SetProperty("blackreference", value)
}

func (e *Frei0rFilterLightGraffiti) GetBlackreference() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("blackreference")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Dimming of the light mask
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetDim(value float64) error {
	return e.Element.SetProperty("dim", value)
}

func (e *Frei0rFilterLightGraffiti) GetDim() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("dim")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Display the brightness and threshold, for adjusting the brightness threshold parameter
// Default: false
func (e *Frei0rFilterLightGraffiti) SetStatsbrightness(value bool) error {
	return e.Element.SetProperty("statsbrightness", value)
}

func (e *Frei0rFilterLightGraffiti) GetStatsbrightness() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("statsbrightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Display the background difference and threshold
// Default: true
func (e *Frei0rFilterLightGraffiti) SetStatsdifference(value bool) error {
	return e.Element.SetProperty("statsdifference", value)
}

func (e *Frei0rFilterLightGraffiti) GetStatsdifference() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("statsdifference")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Display the sum of the background difference and the threshold
// Default: false
func (e *Frei0rFilterLightGraffiti) SetStatsdiffsum(value bool) error {
	return e.Element.SetProperty("statsdiffsum", value)
}

func (e *Frei0rFilterLightGraffiti) GetStatsdiffsum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("statsdiffsum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Brightness threshold to distinguish between foreground and background
// Default: 0.588235, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetThresholdbrightness(value float64) error {
	return e.Element.SetProperty("thresholdbrightness", value)
}

func (e *Frei0rFilterLightGraffiti) GetThresholdbrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("thresholdbrightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Make the background transparent
// Default: false
func (e *Frei0rFilterLightGraffiti) SetTransparentbackground(value bool) error {
	return e.Element.SetProperty("transparentbackground", value)
}

func (e *Frei0rFilterLightGraffiti) GetTransparentbackground() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("transparentbackground")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Alpha value for moving average
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetLongalpha(value float64) error {
	return e.Element.SetProperty("longalpha", value)
}

func (e *Frei0rFilterLightGraffiti) GetLongalpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("longalpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Prevents some overexposure if the light source stays steady too long (varying speed)
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetLoweroverexposure(value float64) error {
	return e.Element.SetProperty("loweroverexposure", value)
}

func (e *Frei0rFilterLightGraffiti) GetLoweroverexposure() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("loweroverexposure")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Nonlinear dimming (may look more natural)
// Default: false
func (e *Frei0rFilterLightGraffiti) SetNonlineardim(value bool) error {
	return e.Element.SetProperty("nonlineardim", value)
}

func (e *Frei0rFilterLightGraffiti) GetNonlineardim() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonlineardim")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Threshold for sum of differences. Can in most cases be ignored (set to 0).
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetThresholddiffsum(value float64) error {
	return e.Element.SetProperty("thresholddiffsum", value)
}

func (e *Frei0rFilterLightGraffiti) GetThresholddiffsum() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("thresholddiffsum")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Describes how strong the (accumulated) background should shine through
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetBackgroundweight(value float64) error {
	return e.Element.SetProperty("backgroundweight", value)
}

func (e *Frei0rFilterLightGraffiti) GetBackgroundweight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("backgroundweight")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Reset filter masks
// Default: false
func (e *Frei0rFilterLightGraffiti) SetReset(value bool) error {
	return e.Element.SetProperty("reset", value)
}

func (e *Frei0rFilterLightGraffiti) GetReset() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reset")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Saturation of lights
// Default: 0.25, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *Frei0rFilterLightGraffiti) GetSaturation() (float64, error) {
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

// Sensitivity of the effect for light (higher sensitivity will lead to brighter lights)
// Default: 0.2, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetSensitivity(value float64) error {
	return e.Element.SetProperty("sensitivity", value)
}

func (e *Frei0rFilterLightGraffiti) GetSensitivity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Threshold: Difference to background to distinguish between fore- and background
// Default: 0, Min: 0, Max: 1
func (e *Frei0rFilterLightGraffiti) SetThresholddifference(value float64) error {
	return e.Element.SetProperty("thresholddifference", value)
}

func (e *Frei0rFilterLightGraffiti) GetThresholddifference() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("thresholddifference")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}
