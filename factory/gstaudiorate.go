// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Drops/duplicates/adjusts timestamps on audio samples to make a perfect stream
type GstAudioRate struct {
	*gst.Element
}

func NewAudioRate() (*GstAudioRate, error) {
	e, err := gst.NewElement("audiorate")
	if err != nil {
		return nil, err
	}
	return &GstAudioRate{Element: e}, nil
}

func NewAudioRateWithName(name string) (*GstAudioRate, error) {
	e, err := gst.NewElementWithName("audiorate", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioRate{Element: e}, nil
}

// Number of added samples
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAudioRate) GetAdd() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("add")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Number of dropped samples
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAudioRate) GetDrop() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Number of input samples
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAudioRate) GetIn() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("in")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Number of output samples
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAudioRate) GetOut() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("out")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Don't emit notify for dropped and duplicated frames
// Default: true
func (e *GstAudioRate) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstAudioRate) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Don't produce buffers before the first one we receive
// Default: false
func (e *GstAudioRate) SetSkipToFirst(value bool) error {
	return e.Element.SetProperty("skip-to-first", value)
}

func (e *GstAudioRate) GetSkipToFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-to-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Only act if timestamp jitter/imperfection exceeds indicated tolerance (ns)
// Default: 40000000, Min: 0, Max: 18446744073709551615
func (e *GstAudioRate) SetTolerance(value uint64) error {
	return e.Element.SetProperty("tolerance", value)
}

func (e *GstAudioRate) GetTolerance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}
