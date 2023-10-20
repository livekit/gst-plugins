// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Resamples audio
type GstAudioResample struct {
	*base.GstBaseTransform
}

func NewAudioResample() (*GstAudioResample, error) {
	e, err := gst.NewElement("audioresample")
	if err != nil {
		return nil, err
	}
	return &GstAudioResample{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioResampleWithName(name string) (*GstAudioResample, error) {
	e, err := gst.NewElementWithName("audioresample", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioResample{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Memory usage threshold to use if sinc filter mode is AUTO, given in bytes
// Default: 1048576, Min: 0, Max: -1
func (e *GstAudioResample) SetSincFilterAutoThreshold(value uint) error {
	return e.Element.SetProperty("sinc-filter-auto-threshold", value)
}

func (e *GstAudioResample) GetSincFilterAutoThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sinc-filter-auto-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// How to interpolate the sinc filter table
// Default: cubic (2)
func (e *GstAudioResample) SetSincFilterInterpolation(value GstAudioResamplerFilterInterpolation) error {
	e.Element.SetArg("sinc-filter-interpolation", string(value))
	return nil
}

func (e *GstAudioResample) GetSincFilterInterpolation() (GstAudioResamplerFilterInterpolation, error) {
	var value GstAudioResamplerFilterInterpolation
	var ok bool
	v, err := e.Element.GetProperty("sinc-filter-interpolation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioResamplerFilterInterpolation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioResamplerFilterInterpolation")
	}
	return value, nil
}

// What sinc filter table mode to use
// Default: auto (2)
func (e *GstAudioResample) SetSincFilterMode(value GstAudioResamplerFilterMode) error {
	e.Element.SetArg("sinc-filter-mode", string(value))
	return nil
}

func (e *GstAudioResample) GetSincFilterMode() (GstAudioResamplerFilterMode, error) {
	var value GstAudioResamplerFilterMode
	var ok bool
	v, err := e.Element.GetProperty("sinc-filter-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioResamplerFilterMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioResamplerFilterMode")
	}
	return value, nil
}

// Resample quality with 0 being the lowest and 10 being the best
// Default: 4, Min: 0, Max: 10
func (e *GstAudioResample) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstAudioResample) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// What resample method to use
// Default: kaiser (4)
func (e *GstAudioResample) SetResampleMethod(value GstAudioResamplerMethod) error {
	e.Element.SetArg("resample-method", string(value))
	return nil
}

func (e *GstAudioResample) GetResampleMethod() (GstAudioResamplerMethod, error) {
	var value GstAudioResamplerMethod
	var ok bool
	v, err := e.Element.GetProperty("resample-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioResamplerMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioResamplerMethod")
	}
	return value, nil
}

type GstAudioResamplerFilterInterpolation string

const (
	GstAudioResamplerFilterInterpolationNone   GstAudioResamplerFilterInterpolation = "none"   // GST_AUDIO_RESAMPLER_FILTER_INTERPOLATION_NONE
	GstAudioResamplerFilterInterpolationLinear GstAudioResamplerFilterInterpolation = "linear" // GST_AUDIO_RESAMPLER_FILTER_INTERPOLATION_LINEAR
	GstAudioResamplerFilterInterpolationCubic  GstAudioResamplerFilterInterpolation = "cubic"  // GST_AUDIO_RESAMPLER_FILTER_INTERPOLATION_CUBIC
)

type GstAudioResamplerFilterMode string

const (
	GstAudioResamplerFilterModeInterpolated GstAudioResamplerFilterMode = "interpolated" // GST_AUDIO_RESAMPLER_FILTER_MODE_INTERPOLATED
	GstAudioResamplerFilterModeFull         GstAudioResamplerFilterMode = "full"         // GST_AUDIO_RESAMPLER_FILTER_MODE_FULL
	GstAudioResamplerFilterModeAuto         GstAudioResamplerFilterMode = "auto"         // GST_AUDIO_RESAMPLER_FILTER_MODE_AUTO
)

type GstAudioResamplerMethod string

const (
	GstAudioResamplerMethodNearest         GstAudioResamplerMethod = "nearest"          // GST_AUDIO_RESAMPLER_METHOD_NEAREST
	GstAudioResamplerMethodLinear          GstAudioResamplerMethod = "linear"           // GST_AUDIO_RESAMPLER_METHOD_LINEAR
	GstAudioResamplerMethodCubic           GstAudioResamplerMethod = "cubic"            // GST_AUDIO_RESAMPLER_METHOD_CUBIC
	GstAudioResamplerMethodBlackmanNuttall GstAudioResamplerMethod = "blackman-nuttall" // GST_AUDIO_RESAMPLER_METHOD_BLACKMAN_NUTTALL
	GstAudioResamplerMethodKaiser          GstAudioResamplerMethod = "kaiser"           // GST_AUDIO_RESAMPLER_METHOD_KAISER
)
