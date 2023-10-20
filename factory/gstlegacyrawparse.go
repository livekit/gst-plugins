// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Converts stream into audio frames (deprecated: use rawaudioparse instead)
type GstAudioParse struct {
	*gst.Bin
}

func NewAudioParse() (*GstAudioParse, error) {
	e, err := gst.NewElement("audioparse")
	if err != nil {
		return nil, err
	}
	return &GstAudioParse{Bin: &gst.Bin{Element: e}}, nil
}

func NewAudioParseWithName(name string) (*GstAudioParse, error) {
	e, err := gst.NewElementWithName("audioparse", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioParse{Bin: &gst.Bin{Element: e}}, nil
}

// Channel positions used on the output

func (e *GstAudioParse) SetChannelPositions(value interface{}) error {
	return e.Element.SetProperty("channel-positions", value)
}

func (e *GstAudioParse) GetChannelPositions() (interface{}, error) {
	return e.Element.GetProperty("channel-positions")
}

// Number of channels in raw stream
// Default: 2, Min: 1, Max: 64
func (e *GstAudioParse) SetChannels(value int) error {
	return e.Element.SetProperty("channels", value)
}

func (e *GstAudioParse) GetChannels() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Format of audio samples in raw stream
// Default: raw (0)
func (e *GstAudioParse) SetFormat(value GstAudioParseFormat) error {
	e.Element.SetArg("format", string(value))
	return nil
}

func (e *GstAudioParse) GetFormat() (GstAudioParseFormat, error) {
	var value GstAudioParseFormat
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioParseFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioParseFormat")
	}
	return value, nil
}

// True if audio has interleaved layout
// Default: true
func (e *GstAudioParse) SetInterleaved(value bool) error {
	return e.Element.SetProperty("interleaved", value)
}

func (e *GstAudioParse) GetInterleaved() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("interleaved")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Rate of audio samples in raw stream
// Default: 44100, Min: 1, Max: 2147483647
func (e *GstAudioParse) SetRate(value int) error {
	return e.Element.SetProperty("rate", value)
}

func (e *GstAudioParse) GetRate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Format of audio samples in raw stream
// Default: s16le (4)
func (e *GstAudioParse) SetRawFormat(value interface{}) error {
	return e.Element.SetProperty("raw-format", value)
}

func (e *GstAudioParse) GetRawFormat() (interface{}, error) {
	return e.Element.GetProperty("raw-format")
}

// Use the sink caps for the format, only performing timestamping
// Default: false
func (e *GstAudioParse) SetUseSinkCaps(value bool) error {
	return e.Element.SetProperty("use-sink-caps", value)
}

func (e *GstAudioParse) GetUseSinkCaps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-sink-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Converts stream into video frames (deprecated: use rawvideoparse instead)
type GstVideoParse struct {
	*gst.Bin
}

func NewVideoParse() (*GstVideoParse, error) {
	e, err := gst.NewElement("videoparse")
	if err != nil {
		return nil, err
	}
	return &GstVideoParse{Bin: &gst.Bin{Element: e}}, nil
}

func NewVideoParseWithName(name string) (*GstVideoParse, error) {
	e, err := gst.NewElementWithName("videoparse", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoParse{Bin: &gst.Bin{Element: e}}, nil
}

// Format of images in raw stream
// Default: i420 (2)
func (e *GstVideoParse) SetFormat(value interface{}) error {
	return e.Element.SetProperty("format", value)
}

func (e *GstVideoParse) GetFormat() (interface{}, error) {
	return e.Element.GetProperty("format")
}

// Frame rate of images in raw stream
// Default: 25/1, Min: 0/1, Max: 2147483647/1
func (e *GstVideoParse) SetFramerate(value interface{}) error {
	return e.Element.SetProperty("framerate", value)
}

func (e *GstVideoParse) GetFramerate() (interface{}, error) {
	return e.Element.GetProperty("framerate")
}

// Size of an image in raw stream (0: default)
// Default: 0, Min: 0, Max: -1
func (e *GstVideoParse) SetFramesize(value uint) error {
	return e.Element.SetProperty("framesize", value)
}

func (e *GstVideoParse) GetFramesize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("framesize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// True if video is interlaced
// Default: false
func (e *GstVideoParse) SetInterlaced(value bool) error {
	return e.Element.SetProperty("interlaced", value)
}

func (e *GstVideoParse) GetInterlaced() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("interlaced")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Pixel aspect ratio of images in raw stream
// Default: 1/1, Min: 1/100, Max: 100/1
func (e *GstVideoParse) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstVideoParse) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

// True if top field is earlier than bottom field
// Default: false
func (e *GstVideoParse) SetTopFieldFirst(value bool) error {
	return e.Element.SetProperty("top-field-first", value)
}

func (e *GstVideoParse) GetTopFieldFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("top-field-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Width of images in raw stream
// Default: 320, Min: 0, Max: 2147483647
func (e *GstVideoParse) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstVideoParse) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Height of images in raw stream
// Default: 240, Min: 0, Max: 2147483647
func (e *GstVideoParse) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstVideoParse) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Offset of each planes in bytes using string format: 'o0,o1,o2,o3'
// Default: 0,76800,96000
func (e *GstVideoParse) SetOffsets(value string) error {
	return e.Element.SetProperty("offsets", value)
}

func (e *GstVideoParse) GetOffsets() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("offsets")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Stride of each planes in bytes using string format: 's0,s1,s2,s3'
// Default: 320,160,160
func (e *GstVideoParse) SetStrides(value string) error {
	return e.Element.SetProperty("strides", value)
}

func (e *GstVideoParse) GetStrides() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("strides")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstAudioParseFormat string

const (
	GstAudioParseFormatRaw   GstAudioParseFormat = "raw"   // Raw
	GstAudioParseFormatAlaw  GstAudioParseFormat = "alaw"  // A-Law
	GstAudioParseFormatMulaw GstAudioParseFormat = "mulaw" // µ-Law
)
