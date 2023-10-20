// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Encodes Kate streams from text or subpictures
type GstKateEnc struct {
	*gst.Element
}

func NewKateEnc() (*GstKateEnc, error) {
	e, err := gst.NewElement("kateenc")
	if err != nil {
		return nil, err
	}
	return &GstKateEnc{Element: e}, nil
}

func NewKateEncWithName(name string) (*GstKateEnc, error) {
	e, err := gst.NewElementWithName("kateenc", name)
	if err != nil {
		return nil, err
	}
	return &GstKateEnc{Element: e}, nil
}

// The numerator of the granule rate
// Default: 1000, Min: 1, Max: 2147483647
func (e *GstKateEnc) SetGranuleRateNumerator(value int) error {
	return e.Element.SetProperty("granule-rate-numerator", value)
}

func (e *GstKateEnc) GetGranuleRateNumerator() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("granule-rate-numerator")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum time to emit keepalive packets (0 disables keepalive packets)
// Default: 2.5, Min: 0, Max: 3.40282e+38
func (e *GstKateEnc) SetKeepaliveMinTime(value float32) error {
	return e.Element.SetProperty("keepalive-min-time", value)
}

func (e *GstKateEnc) GetKeepaliveMinTime() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("keepalive-min-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The language of the stream (e.g. "fr" or "fr_FR" for French)

func (e *GstKateEnc) SetLanguage(value string) error {
	return e.Element.SetProperty("language", value)
}

func (e *GstKateEnc) GetLanguage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("language")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The height of the canvas this stream was authored for (0 is unspecified)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstKateEnc) SetOriginalCanvasHeight(value int) error {
	return e.Element.SetProperty("original-canvas-height", value)
}

func (e *GstKateEnc) GetOriginalCanvasHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The width of the canvas this stream was authored for (0 is unspecified)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstKateEnc) SetOriginalCanvasWidth(value int) error {
	return e.Element.SetProperty("original-canvas-width", value)
}

func (e *GstKateEnc) GetOriginalCanvasWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The category of the stream

func (e *GstKateEnc) SetCategory(value string) error {
	return e.Element.SetProperty("category", value)
}

func (e *GstKateEnc) GetCategory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("category")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The assumed max duration (in seconds) of SPUs with no duration specified
// Default: 1.5, Min: 0, Max: 3.40282e+38
func (e *GstKateEnc) SetDefaultSpuDuration(value float32) error {
	return e.Element.SetProperty("default-spu-duration", value)
}

func (e *GstKateEnc) GetDefaultSpuDuration() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("default-spu-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The denominator of the granule rate
// Default: 1, Min: 1, Max: 2147483647
func (e *GstKateEnc) SetGranuleRateDenominator(value int) error {
	return e.Element.SetProperty("granule-rate-denominator", value)
}

func (e *GstKateEnc) GetGranuleRateDenominator() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("granule-rate-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The granule shift
// Default: 32, Min: 0, Max: 64
func (e *GstKateEnc) SetGranuleShift(value int) error {
	return e.Element.SetProperty("granule-shift", value)
}

func (e *GstKateEnc) GetGranuleShift() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("granule-shift")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// parse raw kate streams
type GstKateParse struct {
	*gst.Element
}

func NewKateParse() (*GstKateParse, error) {
	e, err := gst.NewElement("kateparse")
	if err != nil {
		return nil, err
	}
	return &GstKateParse{Element: e}, nil
}

func NewKateParseWithName(name string) (*GstKateParse, error) {
	e, err := gst.NewElementWithName("kateparse", name)
	if err != nil {
		return nil, err
	}
	return &GstKateParse{Element: e}, nil
}

// Retags kate streams
type GstKateTag struct {
	*GstKateParse
}

func NewKateTag() (*GstKateTag, error) {
	e, err := gst.NewElement("katetag")
	if err != nil {
		return nil, err
	}
	return &GstKateTag{GstKateParse: &GstKateParse{Element: e}}, nil
}

func NewKateTagWithName(name string) (*GstKateTag, error) {
	e, err := gst.NewElementWithName("katetag", name)
	if err != nil {
		return nil, err
	}
	return &GstKateTag{GstKateParse: &GstKateParse{Element: e}}, nil
}

// Set the height of the canvas this stream was authored for (0 is unspecified)
// Default: -1, Min: 0, Max: 2147483647
func (e *GstKateTag) SetOriginalCanvasHeight(value int) error {
	return e.Element.SetProperty("original-canvas-height", value)
}

func (e *GstKateTag) GetOriginalCanvasHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Set the width of the canvas this stream was authored for (0 is unspecified)
// Default: -1, Min: 0, Max: 2147483647
func (e *GstKateTag) SetOriginalCanvasWidth(value int) error {
	return e.Element.SetProperty("original-canvas-width", value)
}

func (e *GstKateTag) GetOriginalCanvasWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Set the category of the stream

func (e *GstKateTag) SetCategory(value string) error {
	return e.Element.SetProperty("category", value)
}

func (e *GstKateTag) GetCategory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("category")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Set the language of the stream

func (e *GstKateTag) SetLanguage(value string) error {
	return e.Element.SetProperty("language", value)
}

func (e *GstKateTag) GetLanguage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("language")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Decodes Kate text streams
type GstKateDec struct {
	*gst.Element
}

func NewKateDec() (*GstKateDec, error) {
	e, err := gst.NewElement("katedec")
	if err != nil {
		return nil, err
	}
	return &GstKateDec{Element: e}, nil
}

func NewKateDecWithName(name string) (*GstKateDec, error) {
	e, err := gst.NewElementWithName("katedec", name)
	if err != nil {
		return nil, err
	}
	return &GstKateDec{Element: e}, nil
}

// Remove markup from decoded text ?
// Default: false
func (e *GstKateDec) SetRemoveMarkup(value bool) error {
	return e.Element.SetProperty("remove-markup", value)
}

func (e *GstKateDec) GetRemoveMarkup() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-markup")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The category of the stream
// Default: NULL
func (e *GstKateDec) GetCategory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("category")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The language of the stream
// Default: NULL
func (e *GstKateDec) GetLanguage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("language")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The canvas height this stream was authored for (0 is unspecified)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstKateDec) GetOriginalCanvasHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The canvas width this stream was authored for
// Default: 0, Min: 0, Max: 2147483647
func (e *GstKateDec) GetOriginalCanvasWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
