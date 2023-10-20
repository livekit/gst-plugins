// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decoders module files (MOD/S3M/XM/IT/MTM/...) using OpenMPT
type GstOpenMptDec struct {
	*GstNonstreamAudioDecoder
}

func NewOpenMptDec() (*GstOpenMptDec, error) {
	e, err := gst.NewElement("openmptdec")
	if err != nil {
		return nil, err
	}
	return &GstOpenMptDec{GstNonstreamAudioDecoder: &GstNonstreamAudioDecoder{Element: e}}, nil
}

func NewOpenMptDecWithName(name string) (*GstOpenMptDec, error) {
	e, err := gst.NewElementWithName("openmptdec", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenMptDec{GstNonstreamAudioDecoder: &GstNonstreamAudioDecoder{Element: e}}, nil
}

// Subsong that is currently selected for playback
// Default: 0, Min: 0, Max: -1
func (e *GstOpenMptDec) SetCurrentSubsong(value uint) error {
	return e.Element.SetProperty("current-subsong", value)
}

func (e *GstOpenMptDec) GetCurrentSubsong() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-subsong")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Length of interpolation filter to use for the samples (0 = internal default)
// Default: 0, Min: 0, Max: 8
func (e *GstOpenMptDec) SetFilterLength(value int) error {
	return e.Element.SetProperty("filter-length", value)
}

func (e *GstOpenMptDec) GetFilterLength() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("filter-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of times a playback loop shall be executed (special values: 0 = no looping; -1 = infinite loop)
// Default: 0, Min: -1, Max: 2147483647
func (e *GstOpenMptDec) SetNumLoops(value int) error {
	return e.Element.SetProperty("num-loops", value)
}

func (e *GstOpenMptDec) GetNumLoops() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-loops")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Which mode playback shall use when a loop is encountered; looping = reset position to start of loop, steady = do not reset position
// Default: steady (1)
func (e *GstOpenMptDec) SetOutputMode(value interface{}) error {
	return e.Element.SetProperty("output-mode", value)
}

func (e *GstOpenMptDec) GetOutputMode() (interface{}, error) {
	return e.Element.GetProperty("output-mode")
}

// Volume ramping strength; higher value -> slower ramping (-1 = internal default)
// Default: -1, Min: -1, Max: 10
func (e *GstOpenMptDec) SetVolumeRamping(value int) error {
	return e.Element.SetProperty("volume-ramping", value)
}

func (e *GstOpenMptDec) GetVolumeRamping() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("volume-ramping")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Gain to apply to the playback, in millibel
// Default: 0, Min: -2147483647, Max: 2147483647
func (e *GstOpenMptDec) SetMasterGain(value int) error {
	return e.Element.SetProperty("master-gain", value)
}

func (e *GstOpenMptDec) GetMasterGain() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("master-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Size of each output buffer, in samples (actual size can be smaller than this during flush or EOS)
// Default: 1024, Min: 1, Max: 268435455
func (e *GstOpenMptDec) SetOutputBufferSize(value uint) error {
	return e.Element.SetProperty("output-buffer-size", value)
}

func (e *GstOpenMptDec) GetOutputBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("output-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Degree of separation for stereo channels, in percent
// Default: 100, Min: 0, Max: 400
func (e *GstOpenMptDec) SetStereoSeparation(value int) error {
	return e.Element.SetProperty("stereo-separation", value)
}

func (e *GstOpenMptDec) GetStereoSeparation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("stereo-separation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Mode which defines how to treat subsongs
// Default: default (2)
func (e *GstOpenMptDec) SetSubsongMode(value interface{}) error {
	return e.Element.SetProperty("subsong-mode", value)
}

func (e *GstOpenMptDec) GetSubsongMode() (interface{}, error) {
	return e.Element.GetProperty("subsong-mode")
}
