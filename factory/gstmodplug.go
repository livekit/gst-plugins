// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Module decoder based on modplug engine
type GstModPlug struct {
	*gst.Element
}

func NewModPlug() (*GstModPlug, error) {
	e, err := gst.NewElement("modplug")
	if err != nil {
		return nil, err
	}
	return &GstModPlug{Element: e}, nil
}

func NewModPlugWithName(name string) (*GstModPlug, error) {
	e, err := gst.NewElementWithName("modplug", name)
	if err != nil {
		return nil, err
	}
	return &GstModPlug{Element: e}, nil
}

// Reverb depth
// Default: 30, Min: 0, Max: 100
func (e *GstModPlug) SetReverbDepth(value int) error {
	return e.Element.SetProperty("reverb-depth", value)
}

func (e *GstModPlug) GetReverbDepth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("reverb-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The song name
// Default: NULL
func (e *GstModPlug) GetSongname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("songname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Surround delay
// Default: 20, Min: 0, Max: 40
func (e *GstModPlug) SetSurroundDelay(value int) error {
	return e.Element.SetProperty("surround-delay", value)
}

func (e *GstModPlug) GetSurroundDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("surround-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Megabass
// Default: false
func (e *GstModPlug) SetMegabass(value bool) error {
	return e.Element.SetProperty("megabass", value)
}

func (e *GstModPlug) GetMegabass() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("megabass")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Megabass amount
// Default: 40, Min: 0, Max: 100
func (e *GstModPlug) SetMegabassAmount(value int) error {
	return e.Element.SetProperty("megabass-amount", value)
}

func (e *GstModPlug) GetMegabassAmount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("megabass-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// noise reduction
// Default: true
func (e *GstModPlug) SetNoiseReduction(value bool) error {
	return e.Element.SetProperty("noise-reduction", value)
}

func (e *GstModPlug) GetNoiseReduction() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("noise-reduction")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// oversamp
// Default: false
func (e *GstModPlug) SetOversamp(value bool) error {
	return e.Element.SetProperty("oversamp", value)
}

func (e *GstModPlug) GetOversamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("oversamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Reverb delay
// Default: 100, Min: 0, Max: 200
func (e *GstModPlug) SetReverbDelay(value int) error {
	return e.Element.SetProperty("reverb-delay", value)
}

func (e *GstModPlug) GetReverbDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("reverb-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Surround depth
// Default: 20, Min: 0, Max: 100
func (e *GstModPlug) SetSurroundDepth(value int) error {
	return e.Element.SetProperty("surround-depth", value)
}

func (e *GstModPlug) GetSurroundDepth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("surround-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Megabass range
// Default: 30, Min: 0, Max: 100
func (e *GstModPlug) SetMegabassRange(value int) error {
	return e.Element.SetProperty("megabass-range", value)
}

func (e *GstModPlug) GetMegabassRange() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("megabass-range")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Reverb
// Default: false
func (e *GstModPlug) SetReverb(value bool) error {
	return e.Element.SetProperty("reverb", value)
}

func (e *GstModPlug) GetReverb() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reverb")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Surround
// Default: true
func (e *GstModPlug) SetSurround(value bool) error {
	return e.Element.SetProperty("surround", value)
}

func (e *GstModPlug) GetSurround() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("surround")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
