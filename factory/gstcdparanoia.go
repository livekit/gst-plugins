// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Read audio from CD in paranoid mode
type GstCdParanoiaSrc struct {
	*GstAudioCdSrc
}

func NewCdParanoiaSrc() (*GstCdParanoiaSrc, error) {
	e, err := gst.NewElement("cdparanoiasrc")
	if err != nil {
		return nil, err
	}
	return &GstCdParanoiaSrc{GstAudioCdSrc: &GstAudioCdSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewCdParanoiaSrcWithName(name string) (*GstCdParanoiaSrc, error) {
	e, err := gst.NewElementWithName("cdparanoiasrc", name)
	if err != nil {
		return nil, err
	}
	return &GstCdParanoiaSrc{GstAudioCdSrc: &GstAudioCdSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// Force minimum overlap search during verification to n sectors
// Default: -1, Min: -1, Max: 75
func (e *GstCdParanoiaSrc) SetSearchOverlap(value int) error {
	return e.Element.SetProperty("search-overlap", value)
}

func (e *GstCdParanoiaSrc) GetSearchOverlap() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("search-overlap")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Set CD cache size to n sectors (-1 = auto)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstCdParanoiaSrc) SetCacheSize(value int) error {
	return e.Element.SetProperty("cache-size", value)
}

func (e *GstCdParanoiaSrc) GetCacheSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cache-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Use specified generic scsi device
// Default: NULL
func (e *GstCdParanoiaSrc) SetGenericDevice(value string) error {
	return e.Element.SetProperty("generic-device", value)
}

func (e *GstCdParanoiaSrc) GetGenericDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("generic-device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Type of checking to perform
// Default: fragment
func (e *GstCdParanoiaSrc) SetParanoiaMode(value GstCdParanoiaMode) error {
	e.Element.SetArg("paranoia-mode", fmt.Sprint(value))
	return nil
}

func (e *GstCdParanoiaSrc) GetParanoiaMode() (GstCdParanoiaMode, error) {
	var value GstCdParanoiaMode
	var ok bool
	v, err := e.Element.GetProperty("paranoia-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCdParanoiaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCdParanoiaMode")
	}
	return value, nil
}

// Read from device at specified speed (-1 and 0 = full speed)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstCdParanoiaSrc) SetReadSpeed(value int) error {
	return e.Element.SetProperty("read-speed", value)
}

func (e *GstCdParanoiaSrc) GetReadSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("read-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstCdParanoiaMode int

const (
	GstCdParanoiaModeDisable  GstCdParanoiaMode = 0x00000000 // PARANOIA_MODE_DISABLE
	GstCdParanoiaModeFragment GstCdParanoiaMode = 0x00000002 // PARANOIA_MODE_FRAGMENT
	GstCdParanoiaModeOverlap  GstCdParanoiaMode = 0x00000004 // PARANOIA_MODE_OVERLAP
	GstCdParanoiaModeScratch  GstCdParanoiaMode = 0x00000008 // PARANOIA_MODE_SCRATCH
	GstCdParanoiaModeRepair   GstCdParanoiaMode = 0x00000010 // PARANOIA_MODE_REPAIR
	GstCdParanoiaModeFull     GstCdParanoiaMode = 0x000000ff // PARANOIA_MODE_FULL
)
