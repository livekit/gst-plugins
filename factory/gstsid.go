// source: gst-plugins-ugly

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Use libsidplay to decode SID audio tunes
type GstSidDec struct {
	*gst.Element
}

func NewSidDec() (*GstSidDec, error) {
	e, err := gst.NewElement("siddec")
	if err != nil {
		return nil, err
	}
	return &GstSidDec{Element: e}, nil
}

func NewSidDecWithName(name string) (*GstSidDec, error) {
	e, err := gst.NewElementWithName("siddec", name)
	if err != nil {
		return nil, err
	}
	return &GstSidDec{Element: e}, nil
}

// mos8580
// Default: false
func (e *GstSidDec) SetMos8580(value bool) error {
	return e.Element.SetProperty("mos8580", value)
}

func (e *GstSidDec) GetMos8580() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mos8580")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// tune
// Default: 0, Min: 0, Max: 100
func (e *GstSidDec) SetTune(value int) error {
	return e.Element.SetProperty("tune", value)
}

func (e *GstSidDec) GetTune() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Size in bytes to output per buffer
// Default: 4096, Min: 1, Max: -1
func (e *GstSidDec) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstSidDec) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// clock
// Default: pal (1)
func (e *GstSidDec) SetClock(value GstSidClock) error {
	e.Element.SetArg("clock", string(value))
	return nil
}

func (e *GstSidDec) GetClock() (GstSidClock, error) {
	var value GstSidClock
	var ok bool
	v, err := e.Element.GetProperty("clock")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSidClock)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSidClock")
	}
	return value, nil
}

// measured_volume
// Default: true
func (e *GstSidDec) SetMeasuredVolume(value bool) error {
	return e.Element.SetProperty("measured-volume", value)
}

func (e *GstSidDec) GetMeasuredVolume() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("measured-volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// memory
// Default: bank-switching (32)
func (e *GstSidDec) SetMemory(value GstSidMemory) error {
	e.Element.SetArg("memory", string(value))
	return nil
}

func (e *GstSidDec) GetMemory() (GstSidMemory, error) {
	var value GstSidMemory
	var ok bool
	v, err := e.Element.GetProperty("memory")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSidMemory)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSidMemory")
	}
	return value, nil
}

// filter
// Default: true
func (e *GstSidDec) SetFilter(value bool) error {
	return e.Element.SetProperty("filter", value)
}

func (e *GstSidDec) GetFilter() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// force_speed
// Default: false
func (e *GstSidDec) SetForceSpeed(value bool) error {
	return e.Element.SetProperty("force-speed", value)
}

func (e *GstSidDec) GetForceSpeed() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Metadata

func (e *GstSidDec) GetMetadata() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("metadata")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

type GstSidClock string

const (
	GstSidClockPal  GstSidClock = "pal"  // PAL
	GstSidClockNtsc GstSidClock = "ntsc" // NTSC
)

type GstSidMemory string

const (
	GstSidMemoryBankSwitching      GstSidMemory = "bank-switching"      // Bank Switching
	GstSidMemoryTransparentRom     GstSidMemory = "transparent-rom"     // Transparent ROM
	GstSidMemoryPlaysidEnvironment GstSidMemory = "playsid-environment" // Playsid Environment
)
