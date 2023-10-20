// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// demux multipart streams
type GstMultipartDemux struct {
	*gst.Element
}

func NewMultipartDemux() (*GstMultipartDemux, error) {
	e, err := gst.NewElement("multipartdemux")
	if err != nil {
		return nil, err
	}
	return &GstMultipartDemux{Element: e}, nil
}

func NewMultipartDemuxWithName(name string) (*GstMultipartDemux, error) {
	e, err := gst.NewElementWithName("multipartdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstMultipartDemux{Element: e}, nil
}

// The boundary string separating data, automatic if NULL
// Default: NULL
func (e *GstMultipartDemux) SetBoundary(value string) error {
	return e.Element.SetProperty("boundary", value)
}

func (e *GstMultipartDemux) GetBoundary() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("boundary")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Assume that there is only one stream whose content-type will not change and emit no-more-pads as soon as the first boundary content is parsed, decoded, and pads are linked
// Default: false
func (e *GstMultipartDemux) SetSingleStream(value bool) error {
	return e.Element.SetProperty("single-stream", value)
}

func (e *GstMultipartDemux) GetSingleStream() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("single-stream")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// mux multipart streams
type GstMultipartMux struct {
	*gst.Element
}

func NewMultipartMux() (*GstMultipartMux, error) {
	e, err := gst.NewElement("multipartmux")
	if err != nil {
		return nil, err
	}
	return &GstMultipartMux{Element: e}, nil
}

func NewMultipartMuxWithName(name string) (*GstMultipartMux, error) {
	e, err := gst.NewElementWithName("multipartmux", name)
	if err != nil {
		return nil, err
	}
	return &GstMultipartMux{Element: e}, nil
}

// Boundary string
// Default: ThisRandomString
func (e *GstMultipartMux) SetBoundary(value string) error {
	return e.Element.SetProperty("boundary", value)
}

func (e *GstMultipartMux) GetBoundary() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("boundary")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
