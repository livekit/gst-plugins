// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Mixes multiple audio streams
type GstAudioInterleave struct {
	*GstAudioAggregator
}

func NewAudioInterleave() (*GstAudioInterleave, error) {
	e, err := gst.NewElement("audiointerleave")
	if err != nil {
		return nil, err
	}
	return &GstAudioInterleave{GstAudioAggregator: &GstAudioAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewAudioInterleaveWithName(name string) (*GstAudioInterleave, error) {
	e, err := gst.NewElementWithName("audiointerleave", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioInterleave{GstAudioAggregator: &GstAudioAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// Channel positions used on the output

func (e *GstAudioInterleave) SetChannelPositions(value interface{}) error {
	return e.Element.SetProperty("channel-positions", value)
}

func (e *GstAudioInterleave) GetChannelPositions() (interface{}, error) {
	return e.Element.GetProperty("channel-positions")
}

// Take channel positions from the input
// Default: true
func (e *GstAudioInterleave) SetChannelPositionsFromInput(value bool) error {
	return e.Element.SetProperty("channel-positions-from-input", value)
}

func (e *GstAudioInterleave) GetChannelPositionsFromInput() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("channel-positions-from-input")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Mixes multiple audio streams
type GstAudioMixer struct {
	*GstAudioAggregator
}

func NewAudioMixer() (*GstAudioMixer, error) {
	e, err := gst.NewElement("audiomixer")
	if err != nil {
		return nil, err
	}
	return &GstAudioMixer{GstAudioAggregator: &GstAudioAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewAudioMixerWithName(name string) (*GstAudioMixer, error) {
	e, err := gst.NewElementWithName("audiomixer", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioMixer{GstAudioAggregator: &GstAudioAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// Mixes multiple audio streams
type GstLiveAdder struct {
	*GstAudioMixer
}

func NewLiveAdder() (*GstLiveAdder, error) {
	e, err := gst.NewElement("liveadder")
	if err != nil {
		return nil, err
	}
	return &GstLiveAdder{GstAudioMixer: &GstAudioMixer{GstAudioAggregator: &GstAudioAggregator{GstAggregator: &GstAggregator{Element: e}}}}, nil
}

func NewLiveAdderWithName(name string) (*GstLiveAdder, error) {
	e, err := gst.NewElementWithName("liveadder", name)
	if err != nil {
		return nil, err
	}
	return &GstLiveAdder{GstAudioMixer: &GstAudioMixer{GstAudioAggregator: &GstAudioAggregator{GstAggregator: &GstAggregator{Element: e}}}}, nil
}

// Additional latency in live mode to allow upstream to take longer to produce buffers for the current position (in milliseconds)
// Default: 30, Min: 0, Max: -1
func (e *GstLiveAdder) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstLiveAdder) GetLatency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
