// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Apply ReplayGain volume adjustment
type GstRgVolume struct {
	*gst.Bin
}

func NewRgVolume() (*GstRgVolume, error) {
	e, err := gst.NewElement("rgvolume")
	if err != nil {
		return nil, err
	}
	return &GstRgVolume{Bin: &gst.Bin{Element: e}}, nil
}

func NewRgVolumeWithName(name string) (*GstRgVolume, error) {
	e, err := gst.NewElementWithName("rgvolume", name)
	if err != nil {
		return nil, err
	}
	return &GstRgVolume{Bin: &gst.Bin{Element: e}}, nil
}

// Applicable gain [dB]
// Default: 0, Min: -120, Max: 120
func (e *GstRgVolume) GetTargetGain() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("target-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Prefer album over track gain
// Default: true
func (e *GstRgVolume) SetAlbumMode(value bool) error {
	return e.Element.SetProperty("album-mode", value)
}

func (e *GstRgVolume) GetAlbumMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("album-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Gain for streams missing tags [dB]
// Default: 0, Min: -60, Max: 60
func (e *GstRgVolume) SetFallbackGain(value float64) error {
	return e.Element.SetProperty("fallback-gain", value)
}

func (e *GstRgVolume) GetFallbackGain() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("fallback-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Extra headroom [dB]
// Default: 0, Min: 0, Max: 60
func (e *GstRgVolume) SetHeadroom(value float64) error {
	return e.Element.SetProperty("headroom", value)
}

func (e *GstRgVolume) GetHeadroom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("headroom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Extra gain [dB]
// Default: 0, Min: -60, Max: 60
func (e *GstRgVolume) SetPreAmp(value float64) error {
	return e.Element.SetProperty("pre-amp", value)
}

func (e *GstRgVolume) GetPreAmp() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("pre-amp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Applied gain [dB]
// Default: 0, Min: -120, Max: 120
func (e *GstRgVolume) GetResultGain() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("result-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Perform the ReplayGain analysis
type GstRgAnalysis struct {
	*base.GstBaseTransform
}

func NewRgAnalysis() (*GstRgAnalysis, error) {
	e, err := gst.NewElement("rganalysis")
	if err != nil {
		return nil, err
	}
	return &GstRgAnalysis{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRgAnalysisWithName(name string) (*GstRgAnalysis, error) {
	e, err := gst.NewElementWithName("rganalysis", name)
	if err != nil {
		return nil, err
	}
	return &GstRgAnalysis{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Analyze even if ReplayGain tags exist
// Default: true
func (e *GstRgAnalysis) SetForced(value bool) error {
	return e.Element.SetProperty("forced", value)
}

func (e *GstRgAnalysis) GetForced() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("forced")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Post statics messages
// Default: false
func (e *GstRgAnalysis) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

func (e *GstRgAnalysis) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of remaining album tracks
// Default: 0, Min: 0, Max: 2147483647
func (e *GstRgAnalysis) SetNumTracks(value int) error {
	return e.Element.SetProperty("num-tracks", value)
}

func (e *GstRgAnalysis) GetNumTracks() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-tracks")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Reference level [dB]
// Default: 89, Min: 0, Max: 150
func (e *GstRgAnalysis) SetReferenceLevel(value float64) error {
	return e.Element.SetProperty("reference-level", value)
}

func (e *GstRgAnalysis) GetReferenceLevel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("reference-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Apply signal compression to raw audio data
type GstRgLimiter struct {
	*base.GstBaseTransform
}

func NewRgLimiter() (*GstRgLimiter, error) {
	e, err := gst.NewElement("rglimiter")
	if err != nil {
		return nil, err
	}
	return &GstRgLimiter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRgLimiterWithName(name string) (*GstRgLimiter, error) {
	e, err := gst.NewElementWithName("rglimiter", name)
	if err != nil {
		return nil, err
	}
	return &GstRgLimiter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Enable processing
// Default: true
func (e *GstRgLimiter) SetEnabled(value bool) error {
	return e.Element.SetProperty("enabled", value)
}

func (e *GstRgLimiter) GetEnabled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enabled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
