// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Midi Synthesizer Element
type GstFluidDec struct {
	*gst.Element
}

func NewFluidDec() (*GstFluidDec, error) {
	e, err := gst.NewElement("fluiddec")
	if err != nil {
		return nil, err
	}
	return &GstFluidDec{Element: e}, nil
}

func NewFluidDecWithName(name string) (*GstFluidDec, error) {
	e, err := gst.NewElementWithName("fluiddec", name)
	if err != nil {
		return nil, err
	}
	return &GstFluidDec{Element: e}, nil
}

// the filename of a soundfont (NULL for default)
// Default: NULL
func (e *GstFluidDec) SetSoundfont(value string) error {
	return e.Element.SetProperty("soundfont", value)
}

func (e *GstFluidDec) GetSoundfont() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("soundfont")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Turn the chorus on or off
// Default: true
func (e *GstFluidDec) SetSynthChorus(value bool) error {
	return e.Element.SetProperty("synth-chorus", value)
}

func (e *GstFluidDec) GetSynthChorus() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("synth-chorus")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set the master gain
// Default: 0.2, Min: 0, Max: 10
func (e *GstFluidDec) SetSynthGain(value float64) error {
	return e.Element.SetProperty("synth-gain", value)
}

func (e *GstFluidDec) GetSynthGain() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("synth-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The number of simultaneous voices
// Default: 256, Min: 1, Max: 65535
func (e *GstFluidDec) SetSynthPolyphony(value int) error {
	return e.Element.SetProperty("synth-polyphony", value)
}

func (e *GstFluidDec) GetSynthPolyphony() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("synth-polyphony")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Turn the reverb on or off
// Default: true
func (e *GstFluidDec) SetSynthReverb(value bool) error {
	return e.Element.SetProperty("synth-reverb", value)
}

func (e *GstFluidDec) GetSynthReverb() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("synth-reverb")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
