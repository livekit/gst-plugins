// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Adds an ID3v2 header and ID3v1 footer to a file
type GstId3Mux struct {
	*GstTagMux
}

func NewId3Mux() (*GstId3Mux, error) {
	e, err := gst.NewElement("id3mux")
	if err != nil {
		return nil, err
	}
	return &GstId3Mux{GstTagMux: &GstTagMux{Element: e}}, nil
}

func NewId3MuxWithName(name string) (*GstId3Mux, error) {
	e, err := gst.NewElementWithName("id3mux", name)
	if err != nil {
		return nil, err
	}
	return &GstId3Mux{GstTagMux: &GstTagMux{Element: e}}, nil
}

// Set version (3 for id3v2.3, 4 for id3v2.4) of id3v2 tags
// Default: 3, Min: 3, Max: 4
func (e *GstId3Mux) SetV2Version(value int) error {
	return e.Element.SetProperty("v2-version", value)
}

func (e *GstId3Mux) GetV2Version() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("v2-version")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Write an id3v1 tag at the end of the file
// Default: false
func (e *GstId3Mux) SetWriteV1(value bool) error {
	return e.Element.SetProperty("write-v1", value)
}

func (e *GstId3Mux) GetWriteV1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("write-v1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Write an id3v2 tag at the start of the file
// Default: true
func (e *GstId3Mux) SetWriteV2(value bool) error {
	return e.Element.SetProperty("write-v2", value)
}

func (e *GstId3Mux) GetWriteV2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("write-v2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
