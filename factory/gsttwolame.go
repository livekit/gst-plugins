// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// High-quality free MP2 encoder
type GstTwoLame struct {
	*GstAudioEncoder
}

func NewTwoLame() (*GstTwoLame, error) {
	e, err := gst.NewElement("twolamemp2enc")
	if err != nil {
		return nil, err
	}
	return &GstTwoLame{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewTwoLameWithName(name string) (*GstTwoLame, error) {
	e, err := gst.NewElementWithName("twolamemp2enc", name)
	if err != nil {
		return nil, err
	}
	return &GstTwoLame{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Write peak PCM level to each frame
// Default: false
func (e *GstTwoLame) SetEnergyLevelExtension(value bool) error {
	return e.Element.SetProperty("energy-level-extension", value)
}

func (e *GstTwoLame) GetEnergyLevelExtension() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("energy-level-extension")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adds checksum to every frame
// Default: false
func (e *GstTwoLame) SetErrorProtection(value bool) error {
	return e.Element.SetProperty("error-protection", value)
}

func (e *GstTwoLame) GetErrorProtection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("error-protection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Calculate Psymodel every n frames
// Default: 10, Min: 0, Max: 2147483647
func (e *GstTwoLame) SetQuickModeCount(value int) error {
	return e.Element.SetProperty("quick-mode-count", value)
}

func (e *GstTwoLame) GetQuickModeCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quick-mode-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding mode
// Default: joint (1)
func (e *GstTwoLame) SetMode(value GstTwoLameMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstTwoLame) GetMode() (GstTwoLameMode, error) {
	var value GstTwoLameMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTwoLameMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTwoLameMode")
	}
	return value, nil
}

// Calculate Psymodel every frames
// Default: false
func (e *GstTwoLame) SetQuickMode(value bool) error {
	return e.Element.SetProperty("quick-mode", value)
}

func (e *GstTwoLame) GetQuickMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("quick-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable variable bitrate mode
// Default: false
func (e *GstTwoLame) SetVbr(value bool) error {
	return e.Element.SetProperty("vbr", value)
}

func (e *GstTwoLame) GetVbr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vbr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// VBR Level
// Default: 5, Min: -10, Max: 10
func (e *GstTwoLame) SetVbrLevel(value float32) error {
	return e.Element.SetProperty("vbr-level", value)
}

func (e *GstTwoLame) GetVbrLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("vbr-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Mark as original
// Default: true
func (e *GstTwoLame) SetOriginal(value bool) error {
	return e.Element.SetProperty("original", value)
}

func (e *GstTwoLame) GetOriginal() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("original")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Padding type
// Default: never (0)
func (e *GstTwoLame) SetPadding(value GstTwoLamePadding) error {
	e.Element.SetArg("padding", string(value))
	return nil
}

func (e *GstTwoLame) GetPadding() (GstTwoLamePadding, error) {
	var value GstTwoLamePadding
	var ok bool
	v, err := e.Element.GetProperty("padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTwoLamePadding)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTwoLamePadding")
	}
	return value, nil
}

// Psychoacoustic model used to encode the audio
// Default: 3, Min: -1, Max: 4
func (e *GstTwoLame) SetPsymodel(value int) error {
	return e.Element.SetProperty("psymodel", value)
}

func (e *GstTwoLame) GetPsymodel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("psymodel")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ATH Level in dB
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstTwoLame) SetAthLevel(value float32) error {
	return e.Element.SetProperty("ath-level", value)
}

func (e *GstTwoLame) GetAthLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ath-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Bitrate in kbit/sec (8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160, 192, 224, 256, 320, 384)
// Default: 192, Min: 8, Max: 384
func (e *GstTwoLame) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstTwoLame) GetBitrate() (int, error) {
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

// Mark as copyright
// Default: false
func (e *GstTwoLame) SetCopyright(value bool) error {
	return e.Element.SetProperty("copyright", value)
}

func (e *GstTwoLame) GetCopyright() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("copyright")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Pre-emphasis to apply to the decoded audio
// Default: none (0)
func (e *GstTwoLame) SetEmphasis(value GstTwoLameEmphasis) error {
	e.Element.SetArg("emphasis", string(value))
	return nil
}

func (e *GstTwoLame) GetEmphasis() (GstTwoLameEmphasis, error) {
	var value GstTwoLameEmphasis
	var ok bool
	v, err := e.Element.GetProperty("emphasis")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTwoLameEmphasis)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTwoLameEmphasis")
	}
	return value, nil
}

// Specify maximum VBR bitrate (0=off, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160, 192, 224, 256, 320, 384)
// Default: 0, Min: 0, Max: 384
func (e *GstTwoLame) SetVbrMaxBitrate(value int) error {
	return e.Element.SetProperty("vbr-max-bitrate", value)
}

func (e *GstTwoLame) GetVbrMaxBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("vbr-max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstTwoLameEmphasis string

const (
	GstTwoLameEmphasisNone GstTwoLameEmphasis = "none" // No emphasis
	GstTwoLameEmphasis5    GstTwoLameEmphasis = "5"    // 50/15 ms
	GstTwoLameEmphasisCcit GstTwoLameEmphasis = "ccit" // CCIT J.17
)

type GstTwoLameMode string

const (
	GstTwoLameModeAuto   GstTwoLameMode = "auto"   // Auto
	GstTwoLameModeStereo GstTwoLameMode = "stereo" // Stereo
	GstTwoLameModeJoint  GstTwoLameMode = "joint"  // Joint Stereo
	GstTwoLameModeDual   GstTwoLameMode = "dual"   // Dual Channel
	GstTwoLameModeMono   GstTwoLameMode = "mono"   // Mono
)

type GstTwoLamePadding string

const (
	GstTwoLamePaddingNever  GstTwoLamePadding = "never"  // No Padding
	GstTwoLamePaddingAlways GstTwoLamePadding = "always" // Always Pad
)
