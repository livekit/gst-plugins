// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// This element detects DTMF tones
type GstDtmfDetect struct {
	*base.GstBaseTransform
}

func NewDtmfDetect() (*GstDtmfDetect, error) {
	e, err := gst.NewElement("dtmfdetect")
	if err != nil {
		return nil, err
	}
	return &GstDtmfDetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDtmfDetectWithName(name string) (*GstDtmfDetect, error) {
	e, err := gst.NewElementWithName("dtmfdetect", name)
	if err != nil {
		return nil, err
	}
	return &GstDtmfDetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Handle Quality-of-Service events
// Default: false
func (e *GstDtmfDetect) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstDtmfDetect) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adds packet loss concealment to audio
type GstSpanPlc struct {
	*gst.Element
}

func NewSpanPlc() (*GstSpanPlc, error) {
	e, err := gst.NewElement("spanplc")
	if err != nil {
		return nil, err
	}
	return &GstSpanPlc{Element: e}, nil
}

func NewSpanPlcWithName(name string) (*GstSpanPlc, error) {
	e, err := gst.NewElementWithName("spanplc", name)
	if err != nil {
		return nil, err
	}
	return &GstSpanPlc{Element: e}, nil
}

// Various statistics

func (e *GstSpanPlc) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Creates telephony signals of given frequency, volume, cadence
type GstToneGenerateSrc struct {
	*base.GstPushSrc
}

func NewToneGenerateSrc() (*GstToneGenerateSrc, error) {
	e, err := gst.NewElement("tonegeneratesrc")
	if err != nil {
		return nil, err
	}
	return &GstToneGenerateSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewToneGenerateSrcWithName(name string) (*GstToneGenerateSrc, error) {
	e, err := gst.NewElementWithName("tonegeneratesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstToneGenerateSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Whether to repeat specified tone indefinitely
// Default: false
func (e *GstToneGenerateSrc) SetRepeat(value bool) error {
	return e.Element.SetProperty("repeat", value)
}

func (e *GstToneGenerateSrc) GetRepeat() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Frequency of test signal
// Default: 0, Min: 0, Max: 20000
func (e *GstToneGenerateSrc) SetFreq(value int) error {
	return e.Element.SetProperty("freq", value)
}

func (e *GstToneGenerateSrc) GetFreq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Frequency of second telephony tone component
// Default: 0, Min: 0, Max: 20000
func (e *GstToneGenerateSrc) SetFreq2(value int) error {
	return e.Element.SetProperty("freq2", value)
}

func (e *GstToneGenerateSrc) GetFreq2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("freq2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Time of the second period  when the tone signal is present
// Default: 1000, Min: 1, Max: 2147483647
func (e *GstToneGenerateSrc) SetOnTime2(value int) error {
	return e.Element.SetProperty("on-time2", value)
}

func (e *GstToneGenerateSrc) GetOnTime2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("on-time2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Time of the first period  when the tone signal is present
// Default: 1000, Min: 1, Max: 2147483647
func (e *GstToneGenerateSrc) SetOnTime(value int) error {
	return e.Element.SetProperty("on-time", value)
}

func (e *GstToneGenerateSrc) GetOnTime() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("on-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of samples in each outgoing buffer
// Default: 1024, Min: 1, Max: 2147483647
func (e *GstToneGenerateSrc) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

func (e *GstToneGenerateSrc) GetSamplesperbuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samplesperbuffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of buffers to output before sending EOS (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstToneGenerateSrc) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstToneGenerateSrc) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Time of the first period  when the tone signal is off
// Default: 1000, Min: 0, Max: 2147483647
func (e *GstToneGenerateSrc) SetOffTime(value int) error {
	return e.Element.SetProperty("off-time", value)
}

func (e *GstToneGenerateSrc) GetOffTime() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("off-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Time of the second period  when the tone signal is off
// Default: 1000, Min: 0, Max: 2147483647
func (e *GstToneGenerateSrc) SetOffTime2(value int) error {
	return e.Element.SetProperty("off-time2", value)
}

func (e *GstToneGenerateSrc) GetOffTime2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("off-time2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Apply current stream time to buffers
// Default: false
func (e *GstToneGenerateSrc) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

func (e *GstToneGenerateSrc) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Volume of first signal
// Default: 0, Min: -50, Max: 0
func (e *GstToneGenerateSrc) SetVolume(value int) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstToneGenerateSrc) GetVolume() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Volume of second tone signal
// Default: 0, Min: -50, Max: 0
func (e *GstToneGenerateSrc) SetVolume2(value int) error {
	return e.Element.SetProperty("volume2", value)
}

func (e *GstToneGenerateSrc) GetVolume2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("volume2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Size in bytes to read per buffer (-1 = default)
// Default: 2048, Min: 0, Max: -1
func (e *GstToneGenerateSrc) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstToneGenerateSrc) GetBlocksize() (uint, error) {
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

// Run typefind before negotiating (deprecated, non-functional)
// Default: false
func (e *GstToneGenerateSrc) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

func (e *GstToneGenerateSrc) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
