// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Virtual audio sink for internal process communication
type GstInterAudioSink struct {
	*base.GstBaseSink
}

func NewInterAudioSink() (*GstInterAudioSink, error) {
	e, err := gst.NewElement("interaudiosink")
	if err != nil {
		return nil, err
	}
	return &GstInterAudioSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewInterAudioSinkWithName(name string) (*GstInterAudioSink, error) {
	e, err := gst.NewElementWithName("interaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &GstInterAudioSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Channel name to match inter src and sink elements
// Default: default
func (e *GstInterAudioSink) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstInterAudioSink) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Virtual audio source for internal process communication
type GstInterAudioSrc struct {
	*base.GstBaseSrc
}

func NewInterAudioSrc() (*GstInterAudioSrc, error) {
	e, err := gst.NewElement("interaudiosrc")
	if err != nil {
		return nil, err
	}
	return &GstInterAudioSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewInterAudioSrcWithName(name string) (*GstInterAudioSrc, error) {
	e, err := gst.NewElementWithName("interaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstInterAudioSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Size of audio buffer
// Default: 1000000000, Min: 1, Max: 18446744073709551615
func (e *GstInterAudioSrc) SetBufferTime(value uint64) error {
	return e.Element.SetProperty("buffer-time", value)
}

func (e *GstInterAudioSrc) GetBufferTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Channel name to match inter src and sink elements
// Default: default
func (e *GstInterAudioSrc) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstInterAudioSrc) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Latency as reported by the source
// Default: 100000000, Min: 1, Max: 18446744073709551615
func (e *GstInterAudioSrc) SetLatencyTime(value uint64) error {
	return e.Element.SetProperty("latency-time", value)
}

func (e *GstInterAudioSrc) GetLatencyTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("latency-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The minimum amount of data to read in each iteration
// Default: 25000000, Min: 1, Max: 18446744073709551615
func (e *GstInterAudioSrc) SetPeriodTime(value uint64) error {
	return e.Element.SetProperty("period-time", value)
}

func (e *GstInterAudioSrc) GetPeriodTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("period-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Virtual subtitle sink for internal process communication
type GstInterSubSink struct {
	*base.GstBaseSink
}

func NewInterSubSink() (*GstInterSubSink, error) {
	e, err := gst.NewElement("intersubsink")
	if err != nil {
		return nil, err
	}
	return &GstInterSubSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewInterSubSinkWithName(name string) (*GstInterSubSink, error) {
	e, err := gst.NewElementWithName("intersubsink", name)
	if err != nil {
		return nil, err
	}
	return &GstInterSubSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Channel name to match inter src and sink elements
// Default: default
func (e *GstInterSubSink) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstInterSubSink) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Virtual subtitle source for internal process communication
type GstInterSubSrc struct {
	*base.GstBaseSrc
}

func NewInterSubSrc() (*GstInterSubSrc, error) {
	e, err := gst.NewElement("intersubsrc")
	if err != nil {
		return nil, err
	}
	return &GstInterSubSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewInterSubSrcWithName(name string) (*GstInterSubSrc, error) {
	e, err := gst.NewElementWithName("intersubsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstInterSubSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Channel name to match inter src and sink elements
// Default: default
func (e *GstInterSubSrc) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstInterSubSrc) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Virtual video sink for internal process communication
type GstInterVideoSink struct {
	*GstVideoSink
}

func NewInterVideoSink() (*GstInterVideoSink, error) {
	e, err := gst.NewElement("intervideosink")
	if err != nil {
		return nil, err
	}
	return &GstInterVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewInterVideoSinkWithName(name string) (*GstInterVideoSink, error) {
	e, err := gst.NewElementWithName("intervideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstInterVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Channel name to match inter src and sink elements
// Default: default
func (e *GstInterVideoSink) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstInterVideoSink) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Virtual video source for internal process communication
type GstInterVideoSrc struct {
	*base.GstBaseSrc
}

func NewInterVideoSrc() (*GstInterVideoSrc, error) {
	e, err := gst.NewElement("intervideosrc")
	if err != nil {
		return nil, err
	}
	return &GstInterVideoSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewInterVideoSrcWithName(name string) (*GstInterVideoSrc, error) {
	e, err := gst.NewElementWithName("intervideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstInterVideoSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Channel name to match inter src and sink elements
// Default: default
func (e *GstInterVideoSrc) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstInterVideoSrc) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Timeout after which to start outputting black frames
// Default: 1000000000, Min: 0, Max: 18446744073709551615
func (e *GstInterVideoSrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstInterVideoSrc) GetTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}
