// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Capture from a soundcard via DirectSound
type GstDirectSoundSrc struct {
	*GstAudioSrc
}

func NewDirectSoundSrc() (*GstDirectSoundSrc, error) {
	e, err := gst.NewElement("directsoundsrc")
	if err != nil {
		return nil, err
	}
	return &GstDirectSoundSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewDirectSoundSrcWithName(name string) (*GstDirectSoundSrc, error) {
	e, err := gst.NewElementWithName("directsoundsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDirectSoundSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// DirectSound playback device as a GUID string (volume and mute will not work!)
// Default: NULL
func (e *GstDirectSoundSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstDirectSoundSrc) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Human-readable name of the sound device
// Default: NULL
func (e *GstDirectSoundSrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

func (e *GstDirectSoundSrc) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Mute state of this stream
// Default: false
func (e *GstDirectSoundSrc) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstDirectSoundSrc) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Volume of this stream
// Default: 1, Min: 0, Max: 1
func (e *GstDirectSoundSrc) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstDirectSoundSrc) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}
