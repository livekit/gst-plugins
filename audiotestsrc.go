// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Audiotestsrc struct {
	*base.GstBaseSrc
}

func NewAudiotestsrc() (*Audiotestsrc, error) {
	e, err := gst.NewElement("audiotestsrc")
	if err != nil {
		return nil, err
	}
	return &Audiotestsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewAudiotestsrcWithName(name string) (*Audiotestsrc, error) {
	e, err := gst.NewElementWithName("audiotestsrc", name)
	if err != nil {
		return nil, err
	}
	return &Audiotestsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// apply-tick-ramp (bool)
//
// Apply ramp to tick samples

func (e *Audiotestsrc) GetApplyTickRamp() (bool, error) {
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

func (e *Audiotestsrc) SetApplyTickRamp(value bool) error {
	return e.Element.SetProperty("apply-tick-ramp", value)
}

// can-activate-pull (bool)
//
// Can activate in pull mode

func (e *Audiotestsrc) GetCanActivatePull() (bool, error) {
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

func (e *Audiotestsrc) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

// can-activate-push (bool)
//
// Can activate in push mode

func (e *Audiotestsrc) GetCanActivatePush() (bool, error) {
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

func (e *Audiotestsrc) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

// freq (float64)
//
// Frequency of test signal. The sample rate needs to be at least 2 times higher.

func (e *Audiotestsrc) GetFreq() (float64, error) {
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

func (e *Audiotestsrc) SetFreq(value float64) error {
	return e.Element.SetProperty("freq", value)
}

// is-live (bool)
//
// Whether to act as a live source

func (e *Audiotestsrc) GetIsLive() (bool, error) {
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

func (e *Audiotestsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// marker-tick-period (uint)
//
// Make every Nth tick a marker tick (= a tick with different volume). Only used if wave = ticks. 0 = no marker ticks.

func (e *Audiotestsrc) GetMarkerTickPeriod() (uint, error) {
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

func (e *Audiotestsrc) SetMarkerTickPeriod(value uint) error {
	return e.Element.SetProperty("marker-tick-period", value)
}

// marker-tick-volume (float64)
//
// Volume of marker ticks. Only used if wave = ticks andmarker-tick-period is set to a nonzero value.

func (e *Audiotestsrc) GetMarkerTickVolume() (float64, error) {
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

func (e *Audiotestsrc) SetMarkerTickVolume(value float64) error {
	return e.Element.SetProperty("marker-tick-volume", value)
}

// samplesperbuffer (int)
//
// Number of samples in each outgoing buffer

func (e *Audiotestsrc) GetSamplesperbuffer() (int, error) {
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

func (e *Audiotestsrc) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

// sine-periods-per-tick (uint)
//
// Number of sine wave periods in one tick. Only used if wave = ticks.

func (e *Audiotestsrc) GetSinePeriodsPerTick() (uint, error) {
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

func (e *Audiotestsrc) SetSinePeriodsPerTick(value uint) error {
	return e.Element.SetProperty("sine-periods-per-tick", value)
}

// tick-interval (uint64)
//
// Distance between start of current and start of next tick, in nanoseconds.

func (e *Audiotestsrc) GetTickInterval() (uint64, error) {
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

func (e *Audiotestsrc) SetTickInterval(value uint64) error {
	return e.Element.SetProperty("tick-interval", value)
}

// timestamp-offset (int64)
//
// An offset added to timestamps set on buffers (in ns)

func (e *Audiotestsrc) GetTimestampOffset() (int64, error) {
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

func (e *Audiotestsrc) SetTimestampOffset(value int64) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

// volume (float64)
//
// Volume of test signal

func (e *Audiotestsrc) GetVolume() (float64, error) {
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

func (e *Audiotestsrc) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

// wave (GstAudioTestSrcWave)
//
// Oscillator waveform

func (e *Audiotestsrc) GetWave() (GstAudioTestSrcWave, error) {
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

func (e *Audiotestsrc) SetWave(value GstAudioTestSrcWave) error {
	e.Element.SetArg("wave", string(value))
	return nil
}

// ----- Constants -----

type GstAudioTestSrcWave string

const (
	GstAudioTestSrcWaveSine GstAudioTestSrcWave = "sine" // sine (0) – Sine
	GstAudioTestSrcWaveSquare GstAudioTestSrcWave = "square" // square (1) – Square
	GstAudioTestSrcWaveSaw GstAudioTestSrcWave = "saw" // saw (2) – Saw
	GstAudioTestSrcWaveTriangle GstAudioTestSrcWave = "triangle" // triangle (3) – Triangle
	GstAudioTestSrcWaveSilence GstAudioTestSrcWave = "silence" // silence (4) – Silence
	GstAudioTestSrcWaveWhiteNoise GstAudioTestSrcWave = "white-noise" // white-noise (5) – White uniform noise
	GstAudioTestSrcWavePinkNoise GstAudioTestSrcWave = "pink-noise" // pink-noise (6) – Pink noise
	GstAudioTestSrcWaveSineTable GstAudioTestSrcWave = "sine-table" // sine-table (7) – Sine table
	GstAudioTestSrcWaveTicks GstAudioTestSrcWave = "ticks" // ticks (8) – Periodic Ticks
	GstAudioTestSrcWaveGaussianNoise GstAudioTestSrcWave = "gaussian-noise" // gaussian-noise (9) – White Gaussian noise
	GstAudioTestSrcWaveRedNoise GstAudioTestSrcWave = "red-noise" // red-noise (10) – Red (brownian) noise
	GstAudioTestSrcWaveBlueNoise GstAudioTestSrcWave = "blue-noise" // blue-noise (11) – Blue noise
	GstAudioTestSrcWaveVioletNoise GstAudioTestSrcWave = "violet-noise" // violet-noise (12) – Violet noise
)

