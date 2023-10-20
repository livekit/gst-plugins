// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Removes all the silence periods from the audio stream.
type GstRemoveSilence struct {
	*base.GstBaseTransform
}

func NewRemoveSilence() (*GstRemoveSilence, error) {
	e, err := gst.NewElement("removesilence")
	if err != nil {
		return nil, err
	}
	return &GstRemoveSilence{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRemoveSilenceWithName(name string) (*GstRemoveSilence, error) {
	e, err := gst.NewElementWithName("removesilence", name)
	if err != nil {
		return nil, err
	}
	return &GstRemoveSilence{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Set to true to retimestamp buffers when silence is removed and so avoid timestamp gap
// Default: false
func (e *GstRemoveSilence) SetSquash(value bool) error {
	return e.Element.SetProperty("squash", value)
}

func (e *GstRemoveSilence) GetSquash() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("squash")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set the silence threshold used on the internal VAD in dB
// Default: -60, Min: -70, Max: 70
func (e *GstRemoveSilence) SetThreshold(value int) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstRemoveSilence) GetThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Set the hysteresis (on samples) used on the internal VAD
// Default: 480, Min: 1, Max: 18446744073709551615
func (e *GstRemoveSilence) SetHysteresis(value uint64) error {
	return e.Element.SetProperty("hysteresis", value)
}

func (e *GstRemoveSilence) GetHysteresis() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("hysteresis")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Define the minimum number of consecutive silence buffers before removing silence, 0 means disabled. This will not introduce latency
// Default: 0, Min: 0, Max: 10000
func (e *GstRemoveSilence) SetMinimumSilenceBuffers(value uint) error {
	return e.Element.SetProperty("minimum-silence-buffers", value)
}

func (e *GstRemoveSilence) GetMinimumSilenceBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("minimum-silence-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Define the minimum silence time in nanoseconds before removing  silence, 0 means disabled. This will not introduce latency
// Default: 0, Min: 0, Max: 10000000000
func (e *GstRemoveSilence) SetMinimumSilenceTime(value uint64) error {
	return e.Element.SetProperty("minimum-silence-time", value)
}

func (e *GstRemoveSilence) GetMinimumSilenceTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("minimum-silence-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Set to true to remove silence from the stream, false otherwise
// Default: false
func (e *GstRemoveSilence) SetRemove(value bool) error {
	return e.Element.SetProperty("remove", value)
}

func (e *GstRemoveSilence) GetRemove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Disable/enable bus message notifications for silence detected/finished
// Default: true
func (e *GstRemoveSilence) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstRemoveSilence) GetSilent() (bool, error) {
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
