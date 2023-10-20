// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// RMS/Peak/Decaying Peak Level messager for audio/raw
type GstLevel struct {
	*base.GstBaseTransform
}

func NewLevel() (*GstLevel, error) {
	e, err := gst.NewElement("level")
	if err != nil {
		return nil, err
	}
	return &GstLevel{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLevelWithName(name string) (*GstLevel, error) {
	e, err := gst.NewElementWithName("level", name)
	if err != nil {
		return nil, err
	}
	return &GstLevel{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Time To Live of decay peak before it falls back (in nanoseconds)
// Default: 300000000, Min: 0, Max: 18446744073709551615
func (e *GstLevel) SetPeakTtl(value uint64) error {
	return e.Element.SetProperty("peak-ttl", value)
}

func (e *GstLevel) GetPeakTtl() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("peak-ttl")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Whether to post a 'level' element message on the bus for each passed interval
// Default: true
func (e *GstLevel) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

func (e *GstLevel) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set GstAudioLevelMeta on buffers
// Default: false
func (e *GstLevel) SetAudioLevelMeta(value bool) error {
	return e.Element.SetProperty("audio-level-meta", value)
}

func (e *GstLevel) GetAudioLevelMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("audio-level-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Interval of time between message posts (in nanoseconds)
// Default: 100000000, Min: 1, Max: 18446744073709551615
func (e *GstLevel) SetInterval(value uint64) error {
	return e.Element.SetProperty("interval", value)
}

func (e *GstLevel) GetInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Post a 'level' message for each passed interval (deprecated, use the post-messages property instead)
// Default: true
func (e *GstLevel) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

func (e *GstLevel) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Decay rate of decay peak after TTL (in dB/sec)
// Default: 10, Min: 0, Max: 1.79769e+308
func (e *GstLevel) SetPeakFalloff(value float64) error {
	return e.Element.SetProperty("peak-falloff", value)
}

func (e *GstLevel) GetPeakFalloff() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("peak-falloff")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}
