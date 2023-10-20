// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// AAC audio encoder
type GstVoAacEnc struct {
	*GstAudioEncoder
}

func NewVoAacEnc() (*GstVoAacEnc, error) {
	e, err := gst.NewElement("voaacenc")
	if err != nil {
		return nil, err
	}
	return &GstVoAacEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewVoAacEncWithName(name string) (*GstVoAacEnc, error) {
	e, err := gst.NewElementWithName("voaacenc", name)
	if err != nil {
		return nil, err
	}
	return &GstVoAacEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Consider discontinuity if timestamp jitter/imperfection exceeds tolerance (ns)
// Default: 40000000, Min: 0, Max: 9223372036854775807
func (e *GstVoAacEnc) SetTolerance(value int64) error {
	return e.Element.SetProperty("tolerance", value)
}

func (e *GstVoAacEnc) GetTolerance() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Target Audio Bitrate (bits per second)
// Default: 128000, Min: 0, Max: 320000
func (e *GstVoAacEnc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstVoAacEnc) GetBitrate() (int, error) {
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

// Perform clipping and sample flushing upon discontinuity
// Default: false
func (e *GstVoAacEnc) SetHardResync(value bool) error {
	return e.Element.SetProperty("hard-resync", value)
}

func (e *GstVoAacEnc) GetHardResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hard-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Apply granule semantics to buffer metadata (implies perfect-timestamp)
// Default: false
func (e *GstVoAacEnc) GetMarkGranule() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mark-granule")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Favour perfect timestamps over tracking upstream timestamps
// Default: false
func (e *GstVoAacEnc) SetPerfectTimestamp(value bool) error {
	return e.Element.SetProperty("perfect-timestamp", value)
}

func (e *GstVoAacEnc) GetPerfectTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("perfect-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
