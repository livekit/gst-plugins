// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Wrapper video sink for automatically detected video sink
type GstAutoVideoSink struct {
	*GstAutoDetect
}

func NewAutoVideoSink() (*GstAutoVideoSink, error) {
	e, err := gst.NewElement("autovideosink")
	if err != nil {
		return nil, err
	}
	return &GstAutoVideoSink{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

func NewAutoVideoSinkWithName(name string) (*GstAutoVideoSink, error) {
	e, err := gst.NewElementWithName("autovideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstAutoVideoSink{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstAutoVideoSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstAutoVideoSink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Wrapper video source for automatically detected video source
type GstAutoVideoSrc struct {
	*GstAutoDetect
}

func NewAutoVideoSrc() (*GstAutoVideoSrc, error) {
	e, err := gst.NewElement("autovideosrc")
	if err != nil {
		return nil, err
	}
	return &GstAutoVideoSrc{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

func NewAutoVideoSrcWithName(name string) (*GstAutoVideoSrc, error) {
	e, err := gst.NewElementWithName("autovideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAutoVideoSrc{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

// Wrapper audio sink for automatically detected audio sink
type GstAutoAudioSink struct {
	*GstAutoDetect
}

func NewAutoAudioSink() (*GstAutoAudioSink, error) {
	e, err := gst.NewElement("autoaudiosink")
	if err != nil {
		return nil, err
	}
	return &GstAutoAudioSink{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

func NewAutoAudioSinkWithName(name string) (*GstAutoAudioSink, error) {
	e, err := gst.NewElementWithName("autoaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &GstAutoAudioSink{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstAutoAudioSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstAutoAudioSink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Wrapper audio source for automatically detected audio source
type GstAutoAudioSrc struct {
	*GstAutoDetect
}

func NewAutoAudioSrc() (*GstAutoAudioSrc, error) {
	e, err := gst.NewElement("autoaudiosrc")
	if err != nil {
		return nil, err
	}
	return &GstAutoAudioSrc{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

func NewAutoAudioSrcWithName(name string) (*GstAutoAudioSrc, error) {
	e, err := gst.NewElementWithName("autoaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAutoAudioSrc{GstAutoDetect: &GstAutoDetect{Bin: &gst.Bin{Element: e}}}, nil
}

type GstAutoDetect struct {
	*gst.Bin
}

// Filter sink candidates using these caps.

func (e *GstAutoDetect) SetFilterCaps(value *gst.Caps) error {
	return e.Element.SetProperty("filter-caps", value)
}

func (e *GstAutoDetect) GetFilterCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("filter-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Sync on the clock
// Default: true
func (e *GstAutoDetect) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstAutoDetect) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
