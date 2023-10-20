// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Removes information from JPEG2000 streams without recompression
type GstJP2kDecimator struct {
	*gst.Element
}

func NewJP2kDecimator() (*GstJP2kDecimator, error) {
	e, err := gst.NewElement("jp2kdecimator")
	if err != nil {
		return nil, err
	}
	return &GstJP2kDecimator{Element: e}, nil
}

func NewJP2kDecimatorWithName(name string) (*GstJP2kDecimator, error) {
	e, err := gst.NewElementWithName("jp2kdecimator", name)
	if err != nil {
		return nil, err
	}
	return &GstJP2kDecimator{Element: e}, nil
}

// Maximum number of decomposition levels to keep (-1 == all)
// Default: -1, Min: -1, Max: 32
func (e *GstJP2kDecimator) SetMaxDecompositionLevels(value int) error {
	return e.Element.SetProperty("max-decomposition-levels", value)
}

func (e *GstJP2kDecimator) GetMaxDecompositionLevels() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-decomposition-levels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum number of layers to keep (0 == all)
// Default: 0, Min: 0, Max: 65535
func (e *GstJP2kDecimator) SetMaxLayers(value int) error {
	return e.Element.SetProperty("max-layers", value)
}

func (e *GstJP2kDecimator) GetMaxLayers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-layers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
