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

type Webrtcdsp struct {
	*base.GstBaseTransform
}

func NewWebrtcdsp() (*Webrtcdsp, error) {
	e, err := gst.NewElement("webrtcdsp")
	if err != nil {
		return nil, err
	}
	return &Webrtcdsp{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewWebrtcdspWithName(name string) (*Webrtcdsp, error) {
	e, err := gst.NewElementWithName("webrtcdsp", name)
	if err != nil {
		return nil, err
	}
	return &Webrtcdsp{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// compression-gain-db (int)
//
// Sets the maximum |gain| the digital compression stage may apply, in dB.

func (e *Webrtcdsp) GetCompressionGainDb() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("compression-gain-db")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Webrtcdsp) SetCompressionGainDb(value int) error {
	return e.Element.SetProperty("compression-gain-db", value)
}

// delay-agnostic (bool)
//
// Enable or disable the delay agnostic mode.

func (e *Webrtcdsp) GetDelayAgnostic() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("delay-agnostic")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetDelayAgnostic(value bool) error {
	return e.Element.SetProperty("delay-agnostic", value)
}

// echo-cancel (bool)
//
// Enable or disable echo canceller

func (e *Webrtcdsp) GetEchoCancel() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("echo-cancel")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetEchoCancel(value bool) error {
	return e.Element.SetProperty("echo-cancel", value)
}

// echo-suppression-level (GstWebrtcEchoSuppressionLevel)
//
// Controls the aggressiveness of the suppressor. A higher level trades off double-talk performance for increased echo suppression.

func (e *Webrtcdsp) GetEchoSuppressionLevel() (GstWebrtcEchoSuppressionLevel, error) {
	var value GstWebrtcEchoSuppressionLevel
	var ok bool
	v, err := e.Element.GetProperty("echo-suppression-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWebrtcEchoSuppressionLevel)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWebrtcEchoSuppressionLevel")
	}
	return value, nil
}

func (e *Webrtcdsp) SetEchoSuppressionLevel(value GstWebrtcEchoSuppressionLevel) error {
	e.Element.SetArg("echo-suppression-level", string(value))
	return nil
}

// experimental-agc (bool)
//
// Enable or disable experimental automatic gain control.

func (e *Webrtcdsp) GetExperimentalAgc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("experimental-agc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetExperimentalAgc(value bool) error {
	return e.Element.SetProperty("experimental-agc", value)
}

// extended-filter (bool)
//
// Enable or disable the extended filter.

func (e *Webrtcdsp) GetExtendedFilter() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("extended-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetExtendedFilter(value bool) error {
	return e.Element.SetProperty("extended-filter", value)
}

// gain-control (bool)
//
// Enable or disable automatic digital gain control

func (e *Webrtcdsp) GetGainControl() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gain-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetGainControl(value bool) error {
	return e.Element.SetProperty("gain-control", value)
}

// gain-control-mode (GstWebrtcGainControlMode)
//
// Controls the mode of the compression stage

func (e *Webrtcdsp) GetGainControlMode() (GstWebrtcGainControlMode, error) {
	var value GstWebrtcGainControlMode
	var ok bool
	v, err := e.Element.GetProperty("gain-control-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWebrtcGainControlMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWebrtcGainControlMode")
	}
	return value, nil
}

func (e *Webrtcdsp) SetGainControlMode(value GstWebrtcGainControlMode) error {
	e.Element.SetArg("gain-control-mode", string(value))
	return nil
}

// high-pass-filter (bool)
//
// Enable or disable high pass filtering

func (e *Webrtcdsp) GetHighPassFilter() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("high-pass-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetHighPassFilter(value bool) error {
	return e.Element.SetProperty("high-pass-filter", value)
}

// limiter (bool)
//
// When enabled, the compression stage will hard limit the signal to the target level. Otherwise, the signal will be compressed but not limited above the target level.

func (e *Webrtcdsp) GetLimiter() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("limiter")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetLimiter(value bool) error {
	return e.Element.SetProperty("limiter", value)
}

// noise-suppression (bool)
//
// Enable or disable noise suppression

func (e *Webrtcdsp) GetNoiseSuppression() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("noise-suppression")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetNoiseSuppression(value bool) error {
	return e.Element.SetProperty("noise-suppression", value)
}

// noise-suppression-level (GstWebrtcNoiseSuppressionLevel)
//
// Controls the aggressiveness of the suppression. Increasing the level will reduce the noise level at the expense of a higher speech distortion.

func (e *Webrtcdsp) GetNoiseSuppressionLevel() (GstWebrtcNoiseSuppressionLevel, error) {
	var value GstWebrtcNoiseSuppressionLevel
	var ok bool
	v, err := e.Element.GetProperty("noise-suppression-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWebrtcNoiseSuppressionLevel)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWebrtcNoiseSuppressionLevel")
	}
	return value, nil
}

func (e *Webrtcdsp) SetNoiseSuppressionLevel(value GstWebrtcNoiseSuppressionLevel) error {
	e.Element.SetArg("noise-suppression-level", string(value))
	return nil
}

// probe (string)
//
// The name of the webrtcechoprobe element that record the audio being played through loud speakers. Must be set before PAUSED state.

func (e *Webrtcdsp) GetProbe() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("probe")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Webrtcdsp) SetProbe(value string) error {
	return e.Element.SetProperty("probe", value)
}

// startup-min-volume (int)
//
// At startup the experimental AGC moves the microphone volume up to |startup_min_volume| if the current microphone volume is set too low. No effect if experimental-agc isn't enabled.

func (e *Webrtcdsp) GetStartupMinVolume() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("startup-min-volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Webrtcdsp) SetStartupMinVolume(value int) error {
	return e.Element.SetProperty("startup-min-volume", value)
}

// target-level-dbfs (int)
//
// Sets the target peak |level| (or envelope) of the gain control in dBFS (decibels from digital full-scale).

func (e *Webrtcdsp) GetTargetLevelDbfs() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("target-level-dbfs")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Webrtcdsp) SetTargetLevelDbfs(value int) error {
	return e.Element.SetProperty("target-level-dbfs", value)
}

// voice-detection (bool)
//
// Enable or disable the voice activity detector

func (e *Webrtcdsp) GetVoiceDetection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("voice-detection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Webrtcdsp) SetVoiceDetection(value bool) error {
	return e.Element.SetProperty("voice-detection", value)
}

// voice-detection-frame-size-ms (int)
//
// Sets the |size| of the frames in ms on which the VAD will operate. Larger frames will improve detection accuracy, but reduce the frequency of updates

func (e *Webrtcdsp) GetVoiceDetectionFrameSizeMs() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("voice-detection-frame-size-ms")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Webrtcdsp) SetVoiceDetectionFrameSizeMs(value int) error {
	return e.Element.SetProperty("voice-detection-frame-size-ms", value)
}

// voice-detection-likelihood (GstWebrtcVoiceDetectionLikelihood)
//
// Specifies the likelihood that a frame will be declared to contain voice.

func (e *Webrtcdsp) GetVoiceDetectionLikelihood() (GstWebrtcVoiceDetectionLikelihood, error) {
	var value GstWebrtcVoiceDetectionLikelihood
	var ok bool
	v, err := e.Element.GetProperty("voice-detection-likelihood")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWebrtcVoiceDetectionLikelihood)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWebrtcVoiceDetectionLikelihood")
	}
	return value, nil
}

func (e *Webrtcdsp) SetVoiceDetectionLikelihood(value GstWebrtcVoiceDetectionLikelihood) error {
	e.Element.SetArg("voice-detection-likelihood", string(value))
	return nil
}

// ----- Constants -----

type GstWebrtcGainControlMode string

const (
	GstWebrtcGainControlModeAdaptiveDigital GstWebrtcGainControlMode = "adaptive-digital" // adaptive-digital (1) – Adaptive Digital
	GstWebrtcGainControlModeFixedDigital GstWebrtcGainControlMode = "fixed-digital" // fixed-digital (2) – Fixed Digital
)

type GstWebrtcNoiseSuppressionLevel string

const (
	GstWebrtcNoiseSuppressionLevelLow GstWebrtcNoiseSuppressionLevel = "low" // low (0) – Low Suppression
	GstWebrtcNoiseSuppressionLevelModerate GstWebrtcNoiseSuppressionLevel = "moderate" // moderate (1) – Moderate Suppression
	GstWebrtcNoiseSuppressionLevelHigh GstWebrtcNoiseSuppressionLevel = "high" // high (2) – High Suppression
	GstWebrtcNoiseSuppressionLevelVeryHigh GstWebrtcNoiseSuppressionLevel = "very-high" // very-high (3) – Very High Suppression
)

type GstWebrtcVoiceDetectionLikelihood string

const (
	GstWebrtcVoiceDetectionLikelihoodVeryLow GstWebrtcVoiceDetectionLikelihood = "very-low" // very-low (0) – Very Low Likelihood
	GstWebrtcVoiceDetectionLikelihoodLow GstWebrtcVoiceDetectionLikelihood = "low" // low (1) – Low Likelihood
	GstWebrtcVoiceDetectionLikelihoodModerate GstWebrtcVoiceDetectionLikelihood = "moderate" // moderate (2) – Moderate Likelihood
	GstWebrtcVoiceDetectionLikelihoodHigh GstWebrtcVoiceDetectionLikelihood = "high" // high (3) – High Likelihood
)

type GstWebrtcEchoSuppressionLevel string

const (
	GstWebrtcEchoSuppressionLevelLow GstWebrtcEchoSuppressionLevel = "low" // low (0) – Low Suppression
	GstWebrtcEchoSuppressionLevelModerate GstWebrtcEchoSuppressionLevel = "moderate" // moderate (1) – Moderate Suppression
	GstWebrtcEchoSuppressionLevelHigh GstWebrtcEchoSuppressionLevel = "high" // high (2) – high Suppression
)

