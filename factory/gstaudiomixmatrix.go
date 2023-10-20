// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Mixes a number of input channels into a number of output channels according to a transformation matrix
type GstAudioMixMatrix struct {
	*base.GstBaseTransform
}

func NewAudioMixMatrix() (*GstAudioMixMatrix, error) {
	e, err := gst.NewElement("audiomixmatrix")
	if err != nil {
		return nil, err
	}
	return &GstAudioMixMatrix{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioMixMatrixWithName(name string) (*GstAudioMixMatrix, error) {
	e, err := gst.NewElementWithName("audiomixmatrix", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioMixMatrix{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// How many audio channels we have on the output side
// Default: 0, Min: 0, Max: 64
func (e *GstAudioMixMatrix) SetOutChannels(value uint) error {
	return e.Element.SetProperty("out-channels", value)
}

func (e *GstAudioMixMatrix) GetOutChannels() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("out-channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Output channel mask (-1 means "default for these channels")
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAudioMixMatrix) SetChannelMask(value uint64) error {
	return e.Element.SetProperty("channel-mask", value)
}

func (e *GstAudioMixMatrix) GetChannelMask() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("channel-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// How many audio channels we have on the input side
// Default: 0, Min: 0, Max: 64
func (e *GstAudioMixMatrix) SetInChannels(value uint) error {
	return e.Element.SetProperty("in-channels", value)
}

func (e *GstAudioMixMatrix) GetInChannels() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("in-channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Transformation matrix for input/output channels

func (e *GstAudioMixMatrix) SetMatrix(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("matrix", value)
}

func (e *GstAudioMixMatrix) GetMatrix() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("matrix")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Whether to auto-negotiate input/output channels and matrix
// Default: manual (0)
func (e *GstAudioMixMatrix) SetMode(value GstAudioMixMatrixModeType) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstAudioMixMatrix) GetMode() (GstAudioMixMatrixModeType, error) {
	var value GstAudioMixMatrixModeType
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioMixMatrixModeType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioMixMatrixModeType")
	}
	return value, nil
}

type GstAudioMixMatrixModeType string

const (
	GstAudioMixMatrixModeTypeManual        GstAudioMixMatrixModeType = "manual"         // Manual mode: please specify input/output channels and transformation matrix
	GstAudioMixMatrixModeTypeFirstChannels GstAudioMixMatrixModeType = "first-channels" // First channels mode: input/output channels are auto-negotiated, transformation matrix is a truncated identity matrix
)
