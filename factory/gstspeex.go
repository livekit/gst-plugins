// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// decode speex streams to audio
type GstSpeexDec struct {
	*GstAudioDecoder
}

func NewSpeexDec() (*GstSpeexDec, error) {
	e, err := gst.NewElement("speexdec")
	if err != nil {
		return nil, err
	}
	return &GstSpeexDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewSpeexDecWithName(name string) (*GstSpeexDec, error) {
	e, err := gst.NewElementWithName("speexdec", name)
	if err != nil {
		return nil, err
	}
	return &GstSpeexDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Enable perceptual enhancement
// Default: true
func (e *GstSpeexDec) SetEnh(value bool) error {
	return e.Element.SetProperty("enh", value)
}

func (e *GstSpeexDec) GetEnh() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enh")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Encodes audio in Speex format
type GstSpeexEnc struct {
	*GstAudioEncoder
}

func NewSpeexEnc() (*GstSpeexEnc, error) {
	e, err := gst.NewElement("speexenc")
	if err != nil {
		return nil, err
	}
	return &GstSpeexEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewSpeexEncWithName(name string) (*GstSpeexEnc, error) {
	e, err := gst.NewElementWithName("speexenc", name)
	if err != nil {
		return nil, err
	}
	return &GstSpeexEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Enable discontinuous transmission
// Default: false
func (e *GstSpeexEnc) SetDtx(value bool) error {
	return e.Element.SetProperty("dtx", value)
}

func (e *GstSpeexEnc) GetDtx() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dtx")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The last status message
// Default: NULL
func (e *GstSpeexEnc) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The encoding mode
// Default: auto (0)
func (e *GstSpeexEnc) SetMode(value GstSpeexEncMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstSpeexEnc) GetMode() (GstSpeexEncMode, error) {
	var value GstSpeexEncMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSpeexEncMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSpeexEncMode")
	}
	return value, nil
}

// Number of frames per buffer
// Default: 1, Min: 0, Max: 2147483647
func (e *GstSpeexEnc) SetNframes(value int) error {
	return e.Element.SetProperty("nframes", value)
}

func (e *GstSpeexEnc) GetNframes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("nframes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Enable variable bit-rate
// Default: false
func (e *GstSpeexEnc) SetVbr(value bool) error {
	return e.Element.SetProperty("vbr", value)
}

func (e *GstSpeexEnc) GetVbr() (bool, error) {
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

// Enable average bit-rate (0 = disabled)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSpeexEnc) SetAbr(value int) error {
	return e.Element.SetProperty("abr", value)
}

func (e *GstSpeexEnc) GetAbr() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("abr")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Specify an encoding bit-rate (in bps). (0 = automatic)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSpeexEnc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstSpeexEnc) GetBitrate() (int, error) {
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

// Enable voice activity detection
// Default: false
func (e *GstSpeexEnc) SetVad(value bool) error {
	return e.Element.SetProperty("vad", value)
}

func (e *GstSpeexEnc) GetVad() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vad")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set encoding complexity
// Default: 3, Min: 0, Max: 2147483647
func (e *GstSpeexEnc) SetComplexity(value int) error {
	return e.Element.SetProperty("complexity", value)
}

func (e *GstSpeexEnc) GetComplexity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("complexity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding quality
// Default: 8, Min: 0, Max: 10
func (e *GstSpeexEnc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstSpeexEnc) GetQuality() (float32, error) {
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

type GstSpeexEncMode string

const (
	GstSpeexEncModeAuto GstSpeexEncMode = "auto" // Auto
	GstSpeexEncModeUwb  GstSpeexEncMode = "uwb"  // Ultra Wide Band
	GstSpeexEncModeWb   GstSpeexEncMode = "wb"   // Wide Band
	GstSpeexEncModeNb   GstSpeexEncMode = "nb"   // Narrow Band
)
