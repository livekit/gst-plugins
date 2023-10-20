// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// High-quality free MP3 encoder
type GstLameMP3Enc struct {
	*GstAudioEncoder
}

func NewLameMP3Enc() (*GstLameMP3Enc, error) {
	e, err := gst.NewElement("lamemp3enc")
	if err != nil {
		return nil, err
	}
	return &GstLameMP3Enc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewLameMP3EncWithName(name string) (*GstLameMP3Enc, error) {
	e, err := gst.NewElementWithName("lamemp3enc", name)
	if err != nil {
		return nil, err
	}
	return &GstLameMP3Enc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Bitrate in kbit/sec (Only valid if target is bitrate, for CBR one of 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256 or 320)
// Default: 128, Min: 8, Max: 320
func (e *GstLameMP3Enc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstLameMP3Enc) GetBitrate() (int, error) {
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

// Enforce constant bitrate encoding (Only valid if target is bitrate)
// Default: false
func (e *GstLameMP3Enc) SetCbr(value bool) error {
	return e.Element.SetProperty("cbr", value)
}

func (e *GstLameMP3Enc) GetCbr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Quality/speed of the encoding engine, this does not affect the bitrate!
// Default: standard (1)
func (e *GstLameMP3Enc) SetEncodingEngineQuality(value GstLameMP3EncEncodingEngineQuality) error {
	e.Element.SetArg("encoding-engine-quality", string(value))
	return nil
}

func (e *GstLameMP3Enc) GetEncodingEngineQuality() (GstLameMP3EncEncodingEngineQuality, error) {
	var value GstLameMP3EncEncodingEngineQuality
	var ok bool
	v, err := e.Element.GetProperty("encoding-engine-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLameMP3EncEncodingEngineQuality)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLameMP3EncEncodingEngineQuality")
	}
	return value, nil
}

// Enforce mono encoding
// Default: false
func (e *GstLameMP3Enc) SetMono(value bool) error {
	return e.Element.SetProperty("mono", value)
}

func (e *GstLameMP3Enc) GetMono() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mono")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// VBR Quality from 0 to 10, 0 being the best (Only valid if target is quality)
// Default: 4, Min: 0, Max: 9.999
func (e *GstLameMP3Enc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstLameMP3Enc) GetQuality() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Optimize for quality or bitrate
// Default: quality (0)
func (e *GstLameMP3Enc) SetTarget(value GstLameMP3EncTarget) error {
	e.Element.SetArg("target", string(value))
	return nil
}

func (e *GstLameMP3Enc) GetTarget() (GstLameMP3EncTarget, error) {
	var value GstLameMP3EncTarget
	var ok bool
	v, err := e.Element.GetProperty("target")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLameMP3EncTarget)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLameMP3EncTarget")
	}
	return value, nil
}

type GstLameMP3EncTarget string

const (
	GstLameMP3EncTargetQuality GstLameMP3EncTarget = "quality" // Quality
	GstLameMP3EncTargetBitrate GstLameMP3EncTarget = "bitrate" // Bitrate
)

type GstLameMP3EncEncodingEngineQuality string

const (
	GstLameMP3EncEncodingEngineQualityFast     GstLameMP3EncEncodingEngineQuality = "fast"     // Fast
	GstLameMP3EncEncodingEngineQualityStandard GstLameMP3EncEncodingEngineQuality = "standard" // Standard
	GstLameMP3EncEncodingEngineQualityHigh     GstLameMP3EncEncodingEngineQuality = "high"     // High
)
