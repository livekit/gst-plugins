// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Detect the BPM of an audio stream
type GstBPMDetect struct {
	*GstAudioFilter
}

func NewBPMDetect() (*GstBPMDetect, error) {
	e, err := gst.NewElement("bpmdetect")
	if err != nil {
		return nil, err
	}
	return &GstBPMDetect{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewBPMDetectWithName(name string) (*GstBPMDetect, error) {
	e, err := gst.NewElementWithName("bpmdetect", name)
	if err != nil {
		return nil, err
	}
	return &GstBPMDetect{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Control the pitch of an audio stream
type GstPitch struct {
	*gst.Element
}

func NewPitch() (*GstPitch, error) {
	e, err := gst.NewElement("pitch")
	if err != nil {
		return nil, err
	}
	return &GstPitch{Element: e}, nil
}

func NewPitchWithName(name string) (*GstPitch, error) {
	e, err := gst.NewElementWithName("pitch", name)
	if err != nil {
		return nil, err
	}
	return &GstPitch{Element: e}, nil
}

// Output rate on downstream segment events
// Default: 1, Min: 0.1, Max: 10
func (e *GstPitch) SetOutputRate(value float32) error {
	return e.Element.SetProperty("output-rate", value)
}

func (e *GstPitch) GetOutputRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("output-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Audio stream pitch
// Default: 1, Min: 0.1, Max: 10
func (e *GstPitch) SetPitch(value float32) error {
	return e.Element.SetProperty("pitch", value)
}

func (e *GstPitch) GetPitch() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pitch")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Audio stream rate
// Default: 1, Min: 0.1, Max: 10
func (e *GstPitch) SetRate(value float32) error {
	return e.Element.SetProperty("rate", value)
}

func (e *GstPitch) GetRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Audio stream tempo
// Default: 1, Min: 0.1, Max: 10
func (e *GstPitch) SetTempo(value float32) error {
	return e.Element.SetProperty("tempo", value)
}

func (e *GstPitch) GetTempo() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tempo")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}
