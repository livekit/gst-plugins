// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Converts unformatted data streams into timestamped raw audio frames
type GstRawAudioParse struct {
	*GstRawBaseParse
}

func NewRawAudioParse() (*GstRawAudioParse, error) {
	e, err := gst.NewElement("rawaudioparse")
	if err != nil {
		return nil, err
	}
	return &GstRawAudioParse{GstRawBaseParse: &GstRawBaseParse{GstBaseParse: &GstBaseParse{Element: e}}}, nil
}

func NewRawAudioParseWithName(name string) (*GstRawAudioParse, error) {
	e, err := gst.NewElementWithName("rawaudioparse", name)
	if err != nil {
		return nil, err
	}
	return &GstRawAudioParse{GstRawBaseParse: &GstRawBaseParse{GstBaseParse: &GstBaseParse{Element: e}}}, nil
}

// Format of the raw audio stream
// Default: pcm (0)
func (e *GstRawAudioParse) SetFormat(value GstRawAudioParseFormat) error {
	e.Element.SetArg("format", string(value))
	return nil
}

func (e *GstRawAudioParse) GetFormat() (GstRawAudioParseFormat, error) {
	var value GstRawAudioParseFormat
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRawAudioParseFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRawAudioParseFormat")
	}
	return value, nil
}

// True if audio has interleaved layout
// Default: true
func (e *GstRawAudioParse) SetInterleaved(value bool) error {
	return e.Element.SetProperty("interleaved", value)
}

func (e *GstRawAudioParse) GetInterleaved() (bool, error) {
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

// Number of channels in raw stream
// Default: 2, Min: 1, Max: 2147483647
func (e *GstRawAudioParse) SetNumChannels(value int) error {
	return e.Element.SetProperty("num-channels", value)
}

func (e *GstRawAudioParse) GetNumChannels() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Format of audio samples in PCM stream (ignored if format property is not set to pcm)
// Default: s16le (4)
func (e *GstRawAudioParse) SetPcmFormat(value interface{}) error {
	return e.Element.SetProperty("pcm-format", value)
}

func (e *GstRawAudioParse) GetPcmFormat() (interface{}, error) {
	return e.Element.GetProperty("pcm-format")
}

// Rate of audio samples in raw stream
// Default: 44100, Min: 1, Max: 2147483647
func (e *GstRawAudioParse) SetSampleRate(value int) error {
	return e.Element.SetProperty("sample-rate", value)
}

func (e *GstRawAudioParse) GetSampleRate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sample-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Channel positions used on the output

func (e *GstRawAudioParse) SetChannelPositions(value interface{}) error {
	return e.Element.SetProperty("channel-positions", value)
}

func (e *GstRawAudioParse) GetChannelPositions() (interface{}, error) {
	return e.Element.GetProperty("channel-positions")
}

// Converts unformatted data streams into timestamped raw video frames
type GstRawVideoParse struct {
	*GstRawBaseParse
}

func NewRawVideoParse() (*GstRawVideoParse, error) {
	e, err := gst.NewElement("rawvideoparse")
	if err != nil {
		return nil, err
	}
	return &GstRawVideoParse{GstRawBaseParse: &GstRawBaseParse{GstBaseParse: &GstBaseParse{Element: e}}}, nil
}

func NewRawVideoParseWithName(name string) (*GstRawVideoParse, error) {
	e, err := gst.NewElementWithName("rawvideoparse", name)
	if err != nil {
		return nil, err
	}
	return &GstRawVideoParse{GstRawBaseParse: &GstRawBaseParse{GstBaseParse: &GstBaseParse{Element: e}}}, nil
}

// Width of frames in raw stream
// Default: 320, Min: 0, Max: 2147483647
func (e *GstRawVideoParse) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstRawVideoParse) GetWidth() (int, error) {
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

// The video source colorimetry
// Default: NULL
func (e *GstRawVideoParse) SetColorimetry(value string) error {
	return e.Element.SetProperty("colorimetry", value)
}

// Format of frames in raw stream
// Default: i420 (2)
func (e *GstRawVideoParse) SetFormat(value interface{}) error {
	return e.Element.SetProperty("format", value)
}

func (e *GstRawVideoParse) GetFormat() (interface{}, error) {
	return e.Element.GetProperty("format")
}

// Size of a frame (0 = frames are tightly packed together)
// Default: 0, Min: 0, Max: -1
func (e *GstRawVideoParse) SetFrameSize(value uint) error {
	return e.Element.SetProperty("frame-size", value)
}

func (e *GstRawVideoParse) GetFrameSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Offsets of the planes in bytes (e.g. plane-offsets="<0,76800>")

func (e *GstRawVideoParse) SetPlaneOffsets(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("plane-offsets", value)
}

func (e *GstRawVideoParse) GetPlaneOffsets() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("plane-offsets")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Strides of the planes in bytes (e.g. plane-strides="<320,320>")

func (e *GstRawVideoParse) SetPlaneStrides(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("plane-strides", value)
}

func (e *GstRawVideoParse) GetPlaneStrides() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("plane-strides")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// True if top field in frames in raw stream come first (not used if frames aren't interlaced)
// Default: false
func (e *GstRawVideoParse) SetTopFieldFirst(value bool) error {
	return e.Element.SetProperty("top-field-first", value)
}

func (e *GstRawVideoParse) GetTopFieldFirst() (bool, error) {
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

// Rate of frames in raw stream
// Default: 25/1, Min: 0/1, Max: 2147483647/1
func (e *GstRawVideoParse) SetFramerate(value interface{}) error {
	return e.Element.SetProperty("framerate", value)
}

func (e *GstRawVideoParse) GetFramerate() (interface{}, error) {
	return e.Element.GetProperty("framerate")
}

// Height of frames in raw stream
// Default: 240, Min: 0, Max: 2147483647
func (e *GstRawVideoParse) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstRawVideoParse) GetHeight() (int, error) {
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

// True if frames in raw stream are interlaced
// Default: false
func (e *GstRawVideoParse) SetInterlaced(value bool) error {
	return e.Element.SetProperty("interlaced", value)
}

func (e *GstRawVideoParse) GetInterlaced() (bool, error) {
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

// Pixel aspect ratio of frames in raw stream
// Default: 1/1, Min: 1/100, Max: 100/1
func (e *GstRawVideoParse) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstRawVideoParse) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

// Parse unaligned raw audio data
type GstUnalignedAudioParse struct {
	*gst.Bin
}

func NewUnalignedAudioParse() (*GstUnalignedAudioParse, error) {
	e, err := gst.NewElement("unalignedaudioparse")
	if err != nil {
		return nil, err
	}
	return &GstUnalignedAudioParse{Bin: &gst.Bin{Element: e}}, nil
}

func NewUnalignedAudioParseWithName(name string) (*GstUnalignedAudioParse, error) {
	e, err := gst.NewElementWithName("unalignedaudioparse", name)
	if err != nil {
		return nil, err
	}
	return &GstUnalignedAudioParse{Bin: &gst.Bin{Element: e}}, nil
}

// Parse unaligned raw video data
type GstUnalignedVideoParse struct {
	*gst.Bin
}

func NewUnalignedVideoParse() (*GstUnalignedVideoParse, error) {
	e, err := gst.NewElement("unalignedvideoparse")
	if err != nil {
		return nil, err
	}
	return &GstUnalignedVideoParse{Bin: &gst.Bin{Element: e}}, nil
}

func NewUnalignedVideoParseWithName(name string) (*GstUnalignedVideoParse, error) {
	e, err := gst.NewElementWithName("unalignedvideoparse", name)
	if err != nil {
		return nil, err
	}
	return &GstUnalignedVideoParse{Bin: &gst.Bin{Element: e}}, nil
}

type GstRawAudioParseFormat string

const (
	GstRawAudioParseFormatPcm   GstRawAudioParseFormat = "pcm"   // PCM
	GstRawAudioParseFormatAlaw  GstRawAudioParseFormat = "alaw"  // A-Law
	GstRawAudioParseFormatMulaw GstRawAudioParseFormat = "mulaw" // µ-Law
)

type GstRawBaseParse struct {
	*GstBaseParse
}

// Use the sink caps for defining the output format
// Default: false
func (e *GstRawBaseParse) SetUseSinkCaps(value bool) error {
	return e.Element.SetProperty("use-sink-caps", value)
}

func (e *GstRawBaseParse) GetUseSinkCaps() (bool, error) {
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
