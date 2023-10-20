// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Splits raw audio buffers into equal sized chunks
type GstAudioBufferSplit struct {
	*gst.Element
}

func NewAudioBufferSplit() (*GstAudioBufferSplit, error) {
	e, err := gst.NewElement("audiobuffersplit")
	if err != nil {
		return nil, err
	}
	return &GstAudioBufferSplit{Element: e}, nil
}

func NewAudioBufferSplitWithName(name string) (*GstAudioBufferSplit, error) {
	e, err := gst.NewElementWithName("audiobuffersplit", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioBufferSplit{Element: e}, nil
}

// Discard the last samples at EOS or discont if they are too small to fill a buffer
// Default: false
func (e *GstAudioBufferSplit) SetStrictBufferSize(value bool) error {
	return e.Element.SetProperty("strict-buffer-size", value)
}

func (e *GstAudioBufferSplit) GetStrictBufferSize() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Timestamp alignment threshold in nanoseconds
// Default: 40000000, Min: 0, Max: 18446744073709551614
func (e *GstAudioBufferSplit) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

func (e *GstAudioBufferSplit) GetAlignmentThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("alignment-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Window of time in nanoseconds to wait before creating a discontinuity
// Default: 1000000000, Min: 0, Max: 18446744073709551614
func (e *GstAudioBufferSplit) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

func (e *GstAudioBufferSplit) GetDiscontWait() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("discont-wait")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Insert silence/drop samples instead of creating a discontinuity
// Default: false
func (e *GstAudioBufferSplit) SetGapless(value bool) error {
	return e.Element.SetProperty("gapless", value)
}

func (e *GstAudioBufferSplit) GetGapless() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gapless")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Do not insert silence in gapless mode if the gap exceeds this period (in ns) (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAudioBufferSplit) SetMaxSilenceTime(value uint64) error {
	return e.Element.SetProperty("max-silence-time", value)
}

func (e *GstAudioBufferSplit) GetMaxSilenceTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-silence-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Output block size in seconds
// Default: 1/50, Min: 1/2147483647, Max: 2147483647/1
func (e *GstAudioBufferSplit) SetOutputBufferDuration(value interface{}) error {
	return e.Element.SetProperty("output-buffer-duration", value)
}

func (e *GstAudioBufferSplit) GetOutputBufferDuration() (interface{}, error) {
	return e.Element.GetProperty("output-buffer-duration")
}

// Output block size in bytes, takes precedence over buffer duration when set to non zero
// Default: 0, Min: 0, Max: 2147483647
func (e *GstAudioBufferSplit) SetOutputBufferSize(value uint) error {
	return e.Element.SetProperty("output-buffer-size", value)
}

func (e *GstAudioBufferSplit) GetOutputBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("output-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
