// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Demux MXF files
type GstMXFDemux struct {
	*gst.Element
}

func NewMXFDemux() (*GstMXFDemux, error) {
	e, err := gst.NewElement("mxfdemux")
	if err != nil {
		return nil, err
	}
	return &GstMXFDemux{Element: e}, nil
}

func NewMXFDemuxWithName(name string) (*GstMXFDemux, error) {
	e, err := gst.NewElementWithName("mxfdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstMXFDemux{Element: e}, nil
}

// Maximum number of nanoseconds by which tracks can differ
// Default: 100000000, Min: 100000000, Max: 18446744073709551615
func (e *GstMXFDemux) SetMaxDrift(value uint64) error {
	return e.Element.SetProperty("max-drift", value)
}

func (e *GstMXFDemux) GetMaxDrift() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-drift")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Material or Source package to use for playback
// Default: NULL
func (e *GstMXFDemux) SetPackage(value string) error {
	return e.Element.SetProperty("package", value)
}

func (e *GstMXFDemux) GetPackage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("package")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Structural metadata of the MXF file

func (e *GstMXFDemux) GetStructure() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("structure")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Muxes video/audio streams into a MXF stream
type GstMXFMux struct {
	*GstAggregator
}

func NewMXFMux() (*GstMXFMux, error) {
	e, err := gst.NewElement("mxfmux")
	if err != nil {
		return nil, err
	}
	return &GstMXFMux{GstAggregator: &GstAggregator{Element: e}}, nil
}

func NewMXFMuxWithName(name string) (*GstMXFMux, error) {
	e, err := gst.NewElementWithName("mxfmux", name)
	if err != nil {
		return nil, err
	}
	return &GstMXFMux{GstAggregator: &GstAggregator{Element: e}}, nil
}
