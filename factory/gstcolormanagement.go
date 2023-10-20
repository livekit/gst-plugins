// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Uses LittleCMS 2 to perform ICC profile correction
type GstLcms struct {
	*GstVideoFilter
}

func NewLcms() (*GstLcms, error) {
	e, err := gst.NewElement("lcms")
	if err != nil {
		return nil, err
	}
	return &GstLcms{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewLcmsWithName(name string) (*GstLcms, error) {
	e, err := gst.NewElementWithName("lcms", name)
	if err != nil {
		return nil, err
	}
	return &GstLcms{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Select the rendering intent of the color correction
// Default: perceptual (0)
func (e *GstLcms) SetIntent(value GstLcmsIntent) error {
	e.Element.SetArg("intent", string(value))
	return nil
}

func (e *GstLcms) GetIntent() (GstLcmsIntent, error) {
	var value GstLcmsIntent
	var ok bool
	v, err := e.Element.GetProperty("intent")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLcmsIntent)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLcmsIntent")
	}
	return value, nil
}

// Select the caching method for the color compensation calculations
// Default: cached (2)
func (e *GstLcms) SetLookup(value GstLcmsLookupMethod) error {
	e.Element.SetArg("lookup", string(value))
	return nil
}

func (e *GstLcms) GetLookup() (GstLcmsLookupMethod, error) {
	var value GstLcmsLookupMethod
	var ok bool
	v, err := e.Element.GetProperty("lookup")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLcmsLookupMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLcmsLookupMethod")
	}
	return value, nil
}

// Select whether purely black pixels should be preserved
// Default: false
func (e *GstLcms) SetPreserveBlack(value bool) error {
	return e.Element.SetProperty("preserve-black", value)
}

func (e *GstLcms) GetPreserveBlack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("preserve-black")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Specify the destination ICC profile file to apply
// Default: NULL
func (e *GstLcms) SetDestProfile(value string) error {
	return e.Element.SetProperty("dest-profile", value)
}

func (e *GstLcms) GetDestProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dest-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Extract and use source profiles embedded in images
// Default: true
func (e *GstLcms) SetEmbeddedProfile(value bool) error {
	return e.Element.SetProperty("embedded-profile", value)
}

func (e *GstLcms) GetEmbeddedProfile() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("embedded-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Specify the input ICC profile file to apply
// Default: NULL
func (e *GstLcms) SetInputProfile(value string) error {
	return e.Element.SetProperty("input-profile", value)
}

func (e *GstLcms) GetInputProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("input-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstLcmsIntent string

const (
	GstLcmsIntentPerceptual GstLcmsIntent = "perceptual" // Perceptual
	GstLcmsIntentRelative   GstLcmsIntent = "relative"   // Relative Colorimetric
	GstLcmsIntentSaturation GstLcmsIntent = "saturation" // Saturation
	GstLcmsIntentAbsolute   GstLcmsIntent = "absolute"   // Absolute Colorimetric
)

type GstLcmsLookupMethod string

const (
	GstLcmsLookupMethodUncached      GstLcmsLookupMethod = "uncached"      // Uncached, calculate every pixel on the fly (very slow playback)
	GstLcmsLookupMethodPrecalculated GstLcmsLookupMethod = "precalculated" // Precalculate lookup table (takes a long time getting READY)
	GstLcmsLookupMethodCached        GstLcmsLookupMethod = "cached"        // Calculate and cache color replacement values on first occurrence
)
