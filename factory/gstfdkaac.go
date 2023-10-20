// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// FDK AAC audio decoder
type GstFdkAacDec struct {
	*GstAudioDecoder
}

func NewFdkAacDec() (*GstFdkAacDec, error) {
	e, err := gst.NewElement("fdkaacdec")
	if err != nil {
		return nil, err
	}
	return &GstFdkAacDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewFdkAacDecWithName(name string) (*GstFdkAacDec, error) {
	e, err := gst.NewElementWithName("fdkaacdec", name)
	if err != nil {
		return nil, err
	}
	return &GstFdkAacDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// FDK AAC audio encoder
type GstFdkAacEnc struct {
	*GstAudioEncoder
}

func NewFdkAacEnc() (*GstFdkAacEnc, error) {
	e, err := gst.NewElement("fdkaacenc")
	if err != nil {
		return nil, err
	}
	return &GstFdkAacEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewFdkAacEncWithName(name string) (*GstFdkAacEnc, error) {
	e, err := gst.NewElementWithName("fdkaacenc", name)
	if err != nil {
		return nil, err
	}
	return &GstFdkAacEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Additional quality control parameter. Can cause workload increase.
// Default: false
func (e *GstFdkAacEnc) SetAfterburner(value bool) error {
	return e.Element.SetProperty("afterburner", value)
}

func (e *GstFdkAacEnc) GetAfterburner() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("afterburner")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Target Audio Bitrate. Only applicable if rate-control=cbr. (0 = fixed value based on sample rate and channel count)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstFdkAacEnc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstFdkAacEnc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Peak Bitrate to adjust maximum bits per audio frame. Bitrate is in bits/second. Only applicable if rate-control=vbr. (0 = Not set)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstFdkAacEnc) SetPeakBitrate(value int) error {
	return e.Element.SetProperty("peak-bitrate", value)
}

func (e *GstFdkAacEnc) GetPeakBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("peak-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Whether Constant or Variable Bitrate should be used.
// Default: cbr (0)
func (e *GstFdkAacEnc) SetRateControl(value GstFdkAacRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstFdkAacEnc) GetRateControl() (GstFdkAacRateControl, error) {
	var value GstFdkAacRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFdkAacRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFdkAacRateControl")
	}
	return value, nil
}

// AAC Variable Bitrate configurations. Requires rate-control as vbr.
// Default: medium (3)
func (e *GstFdkAacEnc) SetVbrPreset(value GstFdkAacVbrPreset) error {
	e.Element.SetArg("vbr-preset", string(value))
	return nil
}

func (e *GstFdkAacEnc) GetVbrPreset() (GstFdkAacVbrPreset, error) {
	var value GstFdkAacVbrPreset
	var ok bool
	v, err := e.Element.GetProperty("vbr-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFdkAacVbrPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFdkAacVbrPreset")
	}
	return value, nil
}

type GstFdkAacRateControl string

const (
	GstFdkAacRateControlCbr GstFdkAacRateControl = "cbr" // Constant Bitrate
	GstFdkAacRateControlVbr GstFdkAacRateControl = "vbr" // Variable Bitrate
)

type GstFdkAacVbrPreset string

const (
	GstFdkAacVbrPresetVeryLow  GstFdkAacVbrPreset = "very-low"  // Very Low Variable Bitrate
	GstFdkAacVbrPresetLow      GstFdkAacVbrPreset = "low"       // Low Variable Bitrate
	GstFdkAacVbrPresetMedium   GstFdkAacVbrPreset = "medium"    // Medium Variable Bitrate
	GstFdkAacVbrPresetHigh     GstFdkAacVbrPreset = "high"      // High Variable Bitrate
	GstFdkAacVbrPresetVeryHigh GstFdkAacVbrPreset = "very-high" // Very High Variable Bitrate
)
