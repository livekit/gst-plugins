// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Write to any GIO-supported location
type GstGioSink struct {
	*GstGioBaseSink
}

func NewGioSink() (*GstGioSink, error) {
	e, err := gst.NewElement("giosink")
	if err != nil {
		return nil, err
	}
	return &GstGioSink{GstGioBaseSink: &GstGioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewGioSinkWithName(name string) (*GstGioSink, error) {
	e, err := gst.NewElementWithName("giosink", name)
	if err != nil {
		return nil, err
	}
	return &GstGioSink{GstGioBaseSink: &GstGioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// GFile to write to

func (e *GstGioSink) SetFile(value interface{}) error {
	return e.Element.SetProperty("file", value)
}

func (e *GstGioSink) GetFile() (interface{}, error) {
	return e.Element.GetProperty("file")
}

// URI location to write to
// Default: NULL
func (e *GstGioSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstGioSink) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Read from any GIO-supported location
type GstGioSrc struct {
	*GstGioBaseSrc
}

func NewGioSrc() (*GstGioSrc, error) {
	e, err := gst.NewElement("giosrc")
	if err != nil {
		return nil, err
	}
	return &GstGioSrc{GstGioBaseSrc: &GstGioBaseSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewGioSrcWithName(name string) (*GstGioSrc, error) {
	e, err := gst.NewElementWithName("giosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstGioSrc{GstGioBaseSrc: &GstGioBaseSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// GFile to read from

func (e *GstGioSrc) SetFile(value interface{}) error {
	return e.Element.SetProperty("file", value)
}

func (e *GstGioSrc) GetFile() (interface{}, error) {
	return e.Element.GetProperty("file")
}

// Whether the file is growing, ignoring its end
// Default: false
func (e *GstGioSrc) SetIsGrowing(value bool) error {
	return e.Element.SetProperty("is-growing", value)
}

func (e *GstGioSrc) GetIsGrowing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-growing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// URI location to read from
// Default: NULL
func (e *GstGioSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstGioSrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Write to any GIO stream
type GstGioStreamSink struct {
	*GstGioBaseSink
}

func NewGioStreamSink() (*GstGioStreamSink, error) {
	e, err := gst.NewElement("giostreamsink")
	if err != nil {
		return nil, err
	}
	return &GstGioStreamSink{GstGioBaseSink: &GstGioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewGioStreamSinkWithName(name string) (*GstGioStreamSink, error) {
	e, err := gst.NewElementWithName("giostreamsink", name)
	if err != nil {
		return nil, err
	}
	return &GstGioStreamSink{GstGioBaseSink: &GstGioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Stream to write to

func (e *GstGioStreamSink) SetStream(value interface{}) error {
	return e.Element.SetProperty("stream", value)
}

func (e *GstGioStreamSink) GetStream() (interface{}, error) {
	return e.Element.GetProperty("stream")
}

// Read from any GIO stream
type GstGioStreamSrc struct {
	*GstGioBaseSrc
}

func NewGioStreamSrc() (*GstGioStreamSrc, error) {
	e, err := gst.NewElement("giostreamsrc")
	if err != nil {
		return nil, err
	}
	return &GstGioStreamSrc{GstGioBaseSrc: &GstGioBaseSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewGioStreamSrcWithName(name string) (*GstGioStreamSrc, error) {
	e, err := gst.NewElementWithName("giostreamsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstGioStreamSrc{GstGioBaseSrc: &GstGioBaseSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Stream to read from

func (e *GstGioStreamSrc) SetStream(value interface{}) error {
	return e.Element.SetProperty("stream", value)
}

func (e *GstGioStreamSrc) GetStream() (interface{}, error) {
	return e.Element.GetProperty("stream")
}

type GstGioBaseSink struct {
	*base.GstBaseSink
}

// Close the stream when the element stops (i.e. goes from READY to NULL) rather than when the element is disposed)
// Default: true
func (e *GstGioBaseSink) SetCloseOnStop(value bool) error {
	return e.Element.SetProperty("close-on-stop", value)
}

func (e *GstGioBaseSink) GetCloseOnStop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("close-on-stop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstGioBaseSrc struct {
	*base.GstBaseSrc
}
