// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Plays audio to a PulseAudio server
type GstPulseSink struct {
	*GstAudioBaseSink
}

func NewPulseSink() (*GstPulseSink, error) {
	e, err := gst.NewElement("pulsesink")
	if err != nil {
		return nil, err
	}
	return &GstPulseSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewPulseSinkWithName(name string) (*GstPulseSink, error) {
	e, err := gst.NewElementWithName("pulsesink", name)
	if err != nil {
		return nil, err
	}
	return &GstPulseSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Mute state of this stream
// Default: false
func (e *GstPulseSink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstPulseSink) GetMute() (bool, error) {
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

// The PulseAudio server to connect to
// Default: NULL
func (e *GstPulseSink) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

func (e *GstPulseSink) GetServer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("server")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// list of pulseaudio stream properties

func (e *GstPulseSink) SetStreamProperties(value *gst.Structure) error {
	return e.Element.SetProperty("stream-properties", value)
}

func (e *GstPulseSink) GetStreamProperties() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stream-properties")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Linear volume of this stream, 1.0=100%%
// Default: 1, Min: 0, Max: 10
func (e *GstPulseSink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstPulseSink) GetVolume() (float64, error) {
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

// The PulseAudio client name to use
// Default: gst-hotdoc-plugins-scanner
func (e *GstPulseSink) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

func (e *GstPulseSink) GetClientName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("client-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The current PulseAudio sink device

func (e *GstPulseSink) GetCurrentDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The PulseAudio sink device to connect to
// Default: NULL
func (e *GstPulseSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstPulseSink) GetDevice() (string, error) {
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
func (e *GstPulseSink) GetDeviceName() (string, error) {
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

// Captures audio from a PulseAudio server
type GstPulseSrc struct {
	*GstAudioSrc
}

func NewPulseSrc() (*GstPulseSrc, error) {
	e, err := gst.NewElement("pulsesrc")
	if err != nil {
		return nil, err
	}
	return &GstPulseSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewPulseSrcWithName(name string) (*GstPulseSrc, error) {
	e, err := gst.NewElementWithName("pulsesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstPulseSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// The current PulseAudio source device

func (e *GstPulseSrc) GetCurrentDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The PulseAudio source device to connect to
// Default: NULL
func (e *GstPulseSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstPulseSrc) GetDevice() (string, error) {
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

// Mute state of this stream
// Default: false
func (e *GstPulseSrc) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstPulseSrc) GetMute() (bool, error) {
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

// The index of the PulseAudio source output corresponding to this record stream
// Default: -1, Min: 0, Max: -1
func (e *GstPulseSrc) GetSourceOutputIndex() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("source-output-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// list of pulseaudio stream properties

func (e *GstPulseSrc) SetStreamProperties(value *gst.Structure) error {
	return e.Element.SetProperty("stream-properties", value)
}

func (e *GstPulseSrc) GetStreamProperties() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stream-properties")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// The PulseAudio client_name_to_use
// Default: gst-hotdoc-plugins-scanner
func (e *GstPulseSrc) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

func (e *GstPulseSrc) GetClientName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("client-name")
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
func (e *GstPulseSrc) GetDeviceName() (string, error) {
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

// The PulseAudio server to connect to
// Default: NULL
func (e *GstPulseSrc) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

func (e *GstPulseSrc) GetServer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("server")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Linear volume of this stream, 1.0=100%%
// Default: 1, Min: 0, Max: 10
func (e *GstPulseSrc) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstPulseSrc) GetVolume() (float64, error) {
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
