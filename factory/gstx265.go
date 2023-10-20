// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// H265 Encoder
type GstX265Enc struct {
	*GstVideoEncoder
}

func NewX265Enc() (*GstX265Enc, error) {
	e, err := gst.NewElement("x265enc")
	if err != nil {
		return nil, err
	}
	return &GstX265Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewX265EncWithName(name string) (*GstX265Enc, error) {
	e, err := gst.NewElementWithName("x265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstX265Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Bitrate in kbit/sec
// Default: 2048, Min: 1, Max: 102400
func (e *GstX265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstX265Enc) GetBitrate() (uint, error) {
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

// Maximal distance between two key-frames (0 = x265 default / 250)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstX265Enc) SetKeyIntMax(value int) error {
	return e.Element.SetProperty("key-int-max", value)
}

func (e *GstX265Enc) GetKeyIntMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("key-int-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// x265 log level
// Default: none (-1)
func (e *GstX265Enc) SetLogLevel(value GstX265LogLevel) error {
	e.Element.SetArg("log-level", string(value))
	return nil
}

func (e *GstX265Enc) GetLogLevel() (GstX265LogLevel, error) {
	var value GstX265LogLevel
	var ok bool
	v, err := e.Element.GetProperty("log-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX265LogLevel)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX265LogLevel")
	}
	return value, nil
}

// String of x265 options (overridden by element properties) in the format "key1=value1:key2=value2".

func (e *GstX265Enc) SetOptionString(value string) error {
	return e.Element.SetProperty("option-string", value)
}

func (e *GstX265Enc) GetOptionString() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("option-string")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// QP for P slices in (implied) CQP mode (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstX265Enc) SetQp(value int) error {
	return e.Element.SetProperty("qp", value)
}

func (e *GstX265Enc) GetQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Preset name for speed/quality tradeoff options
// Default: medium (6)
func (e *GstX265Enc) SetSpeedPreset(value GstX265SpeedPreset) error {
	e.Element.SetArg("speed-preset", string(value))
	return nil
}

func (e *GstX265Enc) GetSpeedPreset() (GstX265SpeedPreset, error) {
	var value GstX265SpeedPreset
	var ok bool
	v, err := e.Element.GetProperty("speed-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX265SpeedPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX265SpeedPreset")
	}
	return value, nil
}

// Preset name for tuning options
// Default: ssim (2)
func (e *GstX265Enc) SetTune(value GstX265Tune) error {
	e.Element.SetArg("tune", string(value))
	return nil
}

func (e *GstX265Enc) GetTune() (GstX265Tune, error) {
	var value GstX265Tune
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX265Tune)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX265Tune")
	}
	return value, nil
}

type GstX265SpeedPreset string

const (
	GstX265SpeedPresetNoPreset  GstX265SpeedPreset = "No preset" // No preset
	GstX265SpeedPresetUltrafast GstX265SpeedPreset = "ultrafast" // ultrafast
	GstX265SpeedPresetSuperfast GstX265SpeedPreset = "superfast" // superfast
	GstX265SpeedPresetVeryfast  GstX265SpeedPreset = "veryfast"  // veryfast
	GstX265SpeedPresetFaster    GstX265SpeedPreset = "faster"    // faster
	GstX265SpeedPresetFast      GstX265SpeedPreset = "fast"      // fast
	GstX265SpeedPresetMedium    GstX265SpeedPreset = "medium"    // medium
	GstX265SpeedPresetSlow      GstX265SpeedPreset = "slow"      // slow
	GstX265SpeedPresetSlower    GstX265SpeedPreset = "slower"    // slower
	GstX265SpeedPresetVeryslow  GstX265SpeedPreset = "veryslow"  // veryslow
	GstX265SpeedPresetPlacebo   GstX265SpeedPreset = "placebo"   // placebo
)

type GstX265LogLevel string

const (
	GstX265LogLevelNone    GstX265LogLevel = "none"    // No logging
	GstX265LogLevelError   GstX265LogLevel = "error"   // Error
	GstX265LogLevelWarning GstX265LogLevel = "warning" // Warning
	GstX265LogLevelInfo    GstX265LogLevel = "info"    // Info
	GstX265LogLevelDebug   GstX265LogLevel = "debug"   // Debug
	GstX265LogLevelFull    GstX265LogLevel = "full"    // Full
)
