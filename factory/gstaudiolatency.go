// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Measures the audio latency between the source and the sink
type GstAudioLatency struct {
	*gst.Bin
}

func NewAudioLatency() (*GstAudioLatency, error) {
	e, err := gst.NewElement("audiolatency")
	if err != nil {
		return nil, err
	}
	return &GstAudioLatency{Bin: &gst.Bin{Element: e}}, nil
}

func NewAudioLatencyWithName(name string) (*GstAudioLatency, error) {
	e, err := gst.NewElementWithName("audiolatency", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioLatency{Bin: &gst.Bin{Element: e}}, nil
}

// The running average latency, in microseconds
// Default: 0, Min: 0, Max: 1000000
func (e *GstAudioLatency) GetAverageLatency() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("average-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The last latency that was measured, in microseconds
// Default: 0, Min: 0, Max: 1000000
func (e *GstAudioLatency) GetLastLatency() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("last-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Print the measured latencies on stdout
// Default: false
func (e *GstAudioLatency) SetPrintLatency(value bool) error {
	return e.Element.SetProperty("print-latency", value)
}

func (e *GstAudioLatency) GetPrintLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("print-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of samples in each outgoing buffer
// Default: 240, Min: 1, Max: 2147483647
func (e *GstAudioLatency) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

func (e *GstAudioLatency) GetSamplesperbuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samplesperbuffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
