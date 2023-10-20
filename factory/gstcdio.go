// source: gst-plugins-ugly

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Read audio from CD using libcdio
type GstCdioCddaSrc struct {
	*GstAudioCdSrc
}

func NewCdioCddaSrc() (*GstCdioCddaSrc, error) {
	e, err := gst.NewElement("cdiocddasrc")
	if err != nil {
		return nil, err
	}
	return &GstCdioCddaSrc{GstAudioCdSrc: &GstAudioCdSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewCdioCddaSrcWithName(name string) (*GstCdioCddaSrc, error) {
	e, err := gst.NewElementWithName("cdiocddasrc", name)
	if err != nil {
		return nil, err
	}
	return &GstCdioCddaSrc{GstAudioCdSrc: &GstAudioCdSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// Read from device at the specified speed (-1 = default)
// Default: -1, Min: -1, Max: 100
func (e *GstCdioCddaSrc) SetReadSpeed(value int) error {
	return e.Element.SetProperty("read-speed", value)
}

func (e *GstCdioCddaSrc) GetReadSpeed() (int, error) {
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
