// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Run an FFT on the audio signal, output spectrum data
type GstSpectrum struct {
	*GstAudioFilter
}

func NewSpectrum() (*GstSpectrum, error) {
	e, err := gst.NewElement("spectrum")
	if err != nil {
		return nil, err
	}
	return &GstSpectrum{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewSpectrumWithName(name string) (*GstSpectrum, error) {
	e, err := gst.NewElementWithName("spectrum", name)
	if err != nil {
		return nil, err
	}
	return &GstSpectrum{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Send separate results for each channel
// Default: false
func (e *GstSpectrum) SetMultiChannel(value bool) error {
	return e.Element.SetProperty("multi-channel", value)
}

func (e *GstSpectrum) GetMultiChannel() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("multi-channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to post a 'spectrum' element message on the bus for each passed interval
// Default: true
func (e *GstSpectrum) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

func (e *GstSpectrum) GetPostMessages() (bool, error) {
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

// dB threshold for result. All lower values will be set to this
// Default: -60, Min: -2147483648, Max: 0
func (e *GstSpectrum) SetThreshold(value int) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstSpectrum) GetThreshold() (int, error) {
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

// Number of frequency bands
// Default: 128, Min: 2, Max: 1073741824
func (e *GstSpectrum) SetBands(value uint) error {
	return e.Element.SetProperty("bands", value)
}

func (e *GstSpectrum) GetBands() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bands")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Interval of time between message posts (in nanoseconds)
// Default: 100000000, Min: 1, Max: 18446744073709551615
func (e *GstSpectrum) SetInterval(value uint64) error {
	return e.Element.SetProperty("interval", value)
}

func (e *GstSpectrum) GetInterval() (uint64, error) {
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

// Whether to add a 'magnitude' field to the structure of any 'spectrum' element messages posted on the bus
// Default: true
func (e *GstSpectrum) SetMessageMagnitude(value bool) error {
	return e.Element.SetProperty("message-magnitude", value)
}

func (e *GstSpectrum) GetMessageMagnitude() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message-magnitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to add a 'phase' field to the structure of any 'spectrum' element messages posted on the bus
// Default: false
func (e *GstSpectrum) SetMessagePhase(value bool) error {
	return e.Element.SetProperty("message-phase", value)
}

func (e *GstSpectrum) GetMessagePhase() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message-phase")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
