// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// decode raw vorbis streams to float audio
type GstVorbisDec struct {
	*GstAudioDecoder
}

func NewVorbisDec() (*GstVorbisDec, error) {
	e, err := gst.NewElement("vorbisdec")
	if err != nil {
		return nil, err
	}
	return &GstVorbisDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewVorbisDecWithName(name string) (*GstVorbisDec, error) {
	e, err := gst.NewElementWithName("vorbisdec", name)
	if err != nil {
		return nil, err
	}
	return &GstVorbisDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Encodes audio in Vorbis format
type GstVorbisEnc struct {
	*GstAudioEncoder
}

func NewVorbisEnc() (*GstVorbisEnc, error) {
	e, err := gst.NewElement("vorbisenc")
	if err != nil {
		return nil, err
	}
	return &GstVorbisEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewVorbisEncWithName(name string) (*GstVorbisEnc, error) {
	e, err := gst.NewElementWithName("vorbisenc", name)
	if err != nil {
		return nil, err
	}
	return &GstVorbisEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Attempt to encode at a bitrate averaging this (in bps). This uses the bitrate management engine, and is not recommended for most users. Quality is a better alternative. (-1 == disabled)
// Default: -1, Min: -1, Max: 250001
func (e *GstVorbisEnc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstVorbisEnc) GetBitrate() (int, error) {
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

// The last status message
// Default: NULL
func (e *GstVorbisEnc) GetLastMessage() (string, error) {
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

// Enable bitrate management engine
// Default: false
func (e *GstVorbisEnc) SetManaged(value bool) error {
	return e.Element.SetProperty("managed", value)
}

func (e *GstVorbisEnc) GetManaged() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("managed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Specify a maximum bitrate (in bps). Useful for streaming applications. (-1 == disabled)
// Default: -1, Min: -1, Max: 250001
func (e *GstVorbisEnc) SetMaxBitrate(value int) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstVorbisEnc) GetMaxBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Specify a minimum bitrate (in bps). Useful for encoding for a fixed-size channel. (-1 == disabled)
// Default: -1, Min: -1, Max: 250001
func (e *GstVorbisEnc) SetMinBitrate(value int) error {
	return e.Element.SetProperty("min-bitrate", value)
}

func (e *GstVorbisEnc) GetMinBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Specify quality instead of specifying a particular bitrate.
// Default: 0.3, Min: -0.1, Max: 1
func (e *GstVorbisEnc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstVorbisEnc) GetQuality() (float32, error) {
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

// parse raw vorbis streams
type GstVorbisParse struct {
	*gst.Element
}

func NewVorbisParse() (*GstVorbisParse, error) {
	e, err := gst.NewElement("vorbisparse")
	if err != nil {
		return nil, err
	}
	return &GstVorbisParse{Element: e}, nil
}

func NewVorbisParseWithName(name string) (*GstVorbisParse, error) {
	e, err := gst.NewElementWithName("vorbisparse", name)
	if err != nil {
		return nil, err
	}
	return &GstVorbisParse{Element: e}, nil
}

// Retags vorbis streams
type GstVorbisTag struct {
	*GstVorbisParse
}

func NewVorbisTag() (*GstVorbisTag, error) {
	e, err := gst.NewElement("vorbistag")
	if err != nil {
		return nil, err
	}
	return &GstVorbisTag{GstVorbisParse: &GstVorbisParse{Element: e}}, nil
}

func NewVorbisTagWithName(name string) (*GstVorbisTag, error) {
	e, err := gst.NewElementWithName("vorbistag", name)
	if err != nil {
		return nil, err
	}
	return &GstVorbisTag{GstVorbisParse: &GstVorbisParse{Element: e}}, nil
}
