// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Srt subtitle encoder
type GstSrtEnc struct {
	*gst.Element
}

func NewSrtEnc() (*GstSrtEnc, error) {
	e, err := gst.NewElement("srtenc")
	if err != nil {
		return nil, err
	}
	return &GstSrtEnc{Element: e}, nil
}

func NewSrtEncWithName(name string) (*GstSrtEnc, error) {
	e, err := gst.NewElementWithName("srtenc", name)
	if err != nil {
		return nil, err
	}
	return &GstSrtEnc{Element: e}, nil
}

// Offset for the duration of the subtitles
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstSrtEnc) SetDuration(value int64) error {
	return e.Element.SetProperty("duration", value)
}

func (e *GstSrtEnc) GetDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Offset for the starttime for the subtitles
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstSrtEnc) SetTimestamp(value int64) error {
	return e.Element.SetProperty("timestamp", value)
}

func (e *GstSrtEnc) GetTimestamp() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// WebVTT subtitle encoder
type GstWebvttEnc struct {
	*gst.Element
}

func NewWebvttEnc() (*GstWebvttEnc, error) {
	e, err := gst.NewElement("webvttenc")
	if err != nil {
		return nil, err
	}
	return &GstWebvttEnc{Element: e}, nil
}

func NewWebvttEncWithName(name string) (*GstWebvttEnc, error) {
	e, err := gst.NewElementWithName("webvttenc", name)
	if err != nil {
		return nil, err
	}
	return &GstWebvttEnc{Element: e}, nil
}

// Offset for the duration of the subtitles
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstWebvttEnc) SetDuration(value int64) error {
	return e.Element.SetProperty("duration", value)
}

func (e *GstWebvttEnc) GetDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Offset for the starttime for the subtitles
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstWebvttEnc) SetTimestamp(value int64) error {
	return e.Element.SetProperty("timestamp", value)
}

func (e *GstWebvttEnc) GetTimestamp() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}
