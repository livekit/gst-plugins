// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes Wavpack audio data
type GstWavpackDec struct {
	*GstAudioDecoder
}

func NewWavpackDec() (*GstWavpackDec, error) {
	e, err := gst.NewElement("wavpackdec")
	if err != nil {
		return nil, err
	}
	return &GstWavpackDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewWavpackDecWithName(name string) (*GstWavpackDec, error) {
	e, err := gst.NewElementWithName("wavpackdec", name)
	if err != nil {
		return nil, err
	}
	return &GstWavpackDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Encodes audio with the Wavpack lossless/lossy audio codec
type GstWavpackEnc struct {
	*GstAudioEncoder
}

func NewWavpackEnc() (*GstWavpackEnc, error) {
	e, err := gst.NewElement("wavpackenc")
	if err != nil {
		return nil, err
	}
	return &GstWavpackEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewWavpackEncWithName(name string) (*GstWavpackEnc, error) {
	e, err := gst.NewElementWithName("wavpackenc", name)
	if err != nil {
		return nil, err
	}
	return &GstWavpackEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Try to encode with this amount of bits per sample. This enables lossy encoding, values smaller than 2.0 disable it again.
// Default: 0, Min: 0, Max: 24
func (e *GstWavpackEnc) SetBitsPerSample(value float64) error {
	return e.Element.SetProperty("bits-per-sample", value)
}

func (e *GstWavpackEnc) GetBitsPerSample() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bits-per-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Use this mode for the correction stream. Only works in lossy mode!
// Default: off (0)
func (e *GstWavpackEnc) SetCorrectionMode(value GstWavpackEncCorrectionMode) error {
	e.Element.SetArg("correction-mode", string(value))
	return nil
}

func (e *GstWavpackEnc) GetCorrectionMode() (GstWavpackEncCorrectionMode, error) {
	var value GstWavpackEncCorrectionMode
	var ok bool
	v, err := e.Element.GetProperty("correction-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWavpackEncCorrectionMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWavpackEncCorrectionMode")
	}
	return value, nil
}

// Use better but slower filters for better compression/quality.
// Default: 0, Min: 0, Max: 6
func (e *GstWavpackEnc) SetExtraProcessing(value uint) error {
	return e.Element.SetProperty("extra-processing", value)
}

func (e *GstWavpackEnc) GetExtraProcessing() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("extra-processing")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Use this joint-stereo mode.
// Default: auto (0)
func (e *GstWavpackEnc) SetJointStereoMode(value GstWavpackEncJSMode) error {
	e.Element.SetArg("joint-stereo-mode", string(value))
	return nil
}

func (e *GstWavpackEnc) GetJointStereoMode() (GstWavpackEncJSMode, error) {
	var value GstWavpackEncJSMode
	var ok bool
	v, err := e.Element.GetProperty("joint-stereo-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWavpackEncJSMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWavpackEncJSMode")
	}
	return value, nil
}

// Store MD5 hash of raw samples within the file.
// Default: false
func (e *GstWavpackEnc) SetMd5(value bool) error {
	return e.Element.SetProperty("md5", value)
}

func (e *GstWavpackEnc) GetMd5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Speed versus compression tradeoff.
// Default: normal (2)
func (e *GstWavpackEnc) SetMode(value GstWavpackEncMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstWavpackEnc) GetMode() (GstWavpackEncMode, error) {
	var value GstWavpackEncMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWavpackEncMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWavpackEncMode")
	}
	return value, nil
}

// Try to encode with this average bitrate (bits/sec). This enables lossy encoding, values smaller than 24000 disable it again.
// Default: 0, Min: 0, Max: 9600000
func (e *GstWavpackEnc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstWavpackEnc) GetBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstWavpackEncCorrectionMode string

const (
	GstWavpackEncCorrectionModeOff       GstWavpackEncCorrectionMode = "off"       // Create no correction file
	GstWavpackEncCorrectionModeOn        GstWavpackEncCorrectionMode = "on"        // Create correction file
	GstWavpackEncCorrectionModeOptimized GstWavpackEncCorrectionMode = "optimized" // Create optimized correction file
)

type GstWavpackEncJSMode string

const (
	GstWavpackEncJSModeAuto      GstWavpackEncJSMode = "auto"      // auto
	GstWavpackEncJSModeLeftright GstWavpackEncJSMode = "leftright" // left/right
	GstWavpackEncJSModeMidside   GstWavpackEncJSMode = "midside"   // mid/side
)

type GstWavpackEncMode string

const (
	GstWavpackEncModeFast     GstWavpackEncMode = "fast"     // Fast Compression
	GstWavpackEncModeNormal   GstWavpackEncMode = "normal"   // Normal Compression
	GstWavpackEncModeHigh     GstWavpackEncMode = "high"     // High Compression
	GstWavpackEncModeVeryhigh GstWavpackEncMode = "veryhigh" // Very High Compression
)
