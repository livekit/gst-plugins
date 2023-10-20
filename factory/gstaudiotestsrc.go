// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Creates audio test signals of given frequency and volume
type GstAudioTestSrc struct {
	*base.GstBaseSrc
}

func NewAudioTestSrc() (*GstAudioTestSrc, error) {
	e, err := gst.NewElement("audiotestsrc")
	if err != nil {
		return nil, err
	}
	return &GstAudioTestSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewAudioTestSrcWithName(name string) (*GstAudioTestSrc, error) {
	e, err := gst.NewElementWithName("audiotestsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioTestSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Volume of marker ticks. Only used if wave = ticks andmarker-tick-period is set to a nonzero value.
// Default: 1, Min: 0, Max: 1
func (e *GstAudioTestSrc) SetMarkerTickVolume(value float64) error {
	return e.Element.SetProperty("marker-tick-volume", value)
}

func (e *GstAudioTestSrc) GetMarkerTickVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("marker-tick-volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Number of samples in each outgoing buffer
// Default: 1024, Min: 1, Max: 2147483647
func (e *GstAudioTestSrc) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

func (e *GstAudioTestSrc) GetSamplesperbuffer() (int, error) {
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

// Number of sine wave periods in one tick. Only used if wave = ticks.
// Default: 10, Min: 1, Max: -1
func (e *GstAudioTestSrc) SetSinePeriodsPerTick(value uint) error {
	return e.Element.SetProperty("sine-periods-per-tick", value)
}

func (e *GstAudioTestSrc) GetSinePeriodsPerTick() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sine-periods-per-tick")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Can activate in push mode
// Default: true
func (e *GstAudioTestSrc) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

func (e *GstAudioTestSrc) GetCanActivatePush() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-push")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Frequency of test signal. The sample rate needs to be at least 2 times higher.
// Default: 440, Min: 0, Max: 1.07374e+09
func (e *GstAudioTestSrc) SetFreq(value float64) error {
	return e.Element.SetProperty("freq", value)
}

func (e *GstAudioTestSrc) GetFreq() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Whether to act as a live source
// Default: false
func (e *GstAudioTestSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstAudioTestSrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Make every Nth tick a marker tick (= a tick with different volume). Only used if wave = ticks. 0 = no marker ticks.
// Default: 0, Min: 0, Max: -1
func (e *GstAudioTestSrc) SetMarkerTickPeriod(value uint) error {
	return e.Element.SetProperty("marker-tick-period", value)
}

func (e *GstAudioTestSrc) GetMarkerTickPeriod() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("marker-tick-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Volume of test signal
// Default: 0.8, Min: 0, Max: 1
func (e *GstAudioTestSrc) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstAudioTestSrc) GetVolume() (float64, error) {
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

// Oscillator waveform
// Default: sine (0)
func (e *GstAudioTestSrc) SetWave(value GstAudioTestSrcWave) error {
	e.Element.SetArg("wave", string(value))
	return nil
}

func (e *GstAudioTestSrc) GetWave() (GstAudioTestSrcWave, error) {
	var value GstAudioTestSrcWave
	var ok bool
	v, err := e.Element.GetProperty("wave")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioTestSrcWave)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioTestSrcWave")
	}
	return value, nil
}

// Apply ramp to tick samples
// Default: false
func (e *GstAudioTestSrc) SetApplyTickRamp(value bool) error {
	return e.Element.SetProperty("apply-tick-ramp", value)
}

func (e *GstAudioTestSrc) GetApplyTickRamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("apply-tick-ramp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Can activate in pull mode
// Default: false
func (e *GstAudioTestSrc) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *GstAudioTestSrc) GetCanActivatePull() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-pull")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Distance between start of current and start of next tick, in nanoseconds.
// Default: 1000000000, Min: 1, Max: 18446744073709551615
func (e *GstAudioTestSrc) SetTickInterval(value uint64) error {
	return e.Element.SetProperty("tick-interval", value)
}

func (e *GstAudioTestSrc) GetTickInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("tick-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// An offset added to timestamps set on buffers (in ns)
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstAudioTestSrc) SetTimestampOffset(value int64) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

func (e *GstAudioTestSrc) GetTimestampOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

type GstAudioTestSrcWave string

const (
	GstAudioTestSrcWaveSine          GstAudioTestSrcWave = "sine"           // Sine
	GstAudioTestSrcWaveSquare        GstAudioTestSrcWave = "square"         // Square
	GstAudioTestSrcWaveSaw           GstAudioTestSrcWave = "saw"            // Saw
	GstAudioTestSrcWaveTriangle      GstAudioTestSrcWave = "triangle"       // Triangle
	GstAudioTestSrcWaveSilence       GstAudioTestSrcWave = "silence"        // Silence
	GstAudioTestSrcWaveWhiteNoise    GstAudioTestSrcWave = "white-noise"    // White uniform noise
	GstAudioTestSrcWavePinkNoise     GstAudioTestSrcWave = "pink-noise"     // Pink noise
	GstAudioTestSrcWaveSineTable     GstAudioTestSrcWave = "sine-table"     // Sine table
	GstAudioTestSrcWaveTicks         GstAudioTestSrcWave = "ticks"          // Periodic Ticks
	GstAudioTestSrcWaveGaussianNoise GstAudioTestSrcWave = "gaussian-noise" // White Gaussian noise
	GstAudioTestSrcWaveRedNoise      GstAudioTestSrcWave = "red-noise"      // Red (brownian) noise
	GstAudioTestSrcWaveBlueNoise     GstAudioTestSrcWave = "blue-noise"     // Blue noise
	GstAudioTestSrcWaveVioletNoise   GstAudioTestSrcWave = "violet-noise"   // Violet noise
)
