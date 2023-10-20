// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Pre-processes voice with WebRTC Audio Processing Library
type GstWebrtcDsp struct {
	*GstAudioFilter
}

func NewWebrtcDsp() (*GstWebrtcDsp, error) {
	e, err := gst.NewElement("webrtcdsp")
	if err != nil {
		return nil, err
	}
	return &GstWebrtcDsp{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewWebrtcDspWithName(name string) (*GstWebrtcDsp, error) {
	e, err := gst.NewElementWithName("webrtcdsp", name)
	if err != nil {
		return nil, err
	}
	return &GstWebrtcDsp{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Controls the mode of the compression stage
// Default: adaptive-digital (1)
func (e *GstWebrtcDsp) SetGainControlMode(value GstWebrtcGainControlMode) error {
	e.Element.SetArg("gain-control-mode", string(value))
	return nil
}

func (e *GstWebrtcDsp) GetGainControlMode() (GstWebrtcGainControlMode, error) {
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

// The name of the webrtcechoprobe element that record the audio being played through loud speakers. Must be set before PAUSED state.
// Default: webrtcechoprobe0
func (e *GstWebrtcDsp) SetProbe(value string) error {
	return e.Element.SetProperty("probe", value)
}

func (e *GstWebrtcDsp) GetProbe() (string, error) {
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

// Sets the maximum |gain| the digital compression stage may apply, in dB.
// Default: 9, Min: 0, Max: 90
func (e *GstWebrtcDsp) SetCompressionGainDb(value int) error {
	return e.Element.SetProperty("compression-gain-db", value)
}

func (e *GstWebrtcDsp) GetCompressionGainDb() (int, error) {
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

// Controls the aggressiveness of the suppressor. A higher level trades off double-talk performance for increased echo suppression.
// Default: moderate (1)
func (e *GstWebrtcDsp) SetEchoSuppressionLevel(value GstWebrtcEchoSuppressionLevel) error {
	e.Element.SetArg("echo-suppression-level", string(value))
	return nil
}

func (e *GstWebrtcDsp) GetEchoSuppressionLevel() (GstWebrtcEchoSuppressionLevel, error) {
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

// Enable or disable experimental automatic gain control.
// Default: false
func (e *GstWebrtcDsp) SetExperimentalAgc(value bool) error {
	return e.Element.SetProperty("experimental-agc", value)
}

func (e *GstWebrtcDsp) GetExperimentalAgc() (bool, error) {
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

// Controls the aggressiveness of the suppression. Increasing the level will reduce the noise level at the expense of a higher speech distortion.
// Default: moderate (1)
func (e *GstWebrtcDsp) SetNoiseSuppressionLevel(value GstWebrtcNoiseSuppressionLevel) error {
	e.Element.SetArg("noise-suppression-level", string(value))
	return nil
}

func (e *GstWebrtcDsp) GetNoiseSuppressionLevel() (GstWebrtcNoiseSuppressionLevel, error) {
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

// Sets the target peak |level| (or envelope) of the gain control in dBFS (decibels from digital full-scale).
// Default: 3, Min: 0, Max: 31
func (e *GstWebrtcDsp) SetTargetLevelDbfs(value int) error {
	return e.Element.SetProperty("target-level-dbfs", value)
}

func (e *GstWebrtcDsp) GetTargetLevelDbfs() (int, error) {
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

// Enable or disable echo canceller
// Default: true
func (e *GstWebrtcDsp) SetEchoCancel(value bool) error {
	return e.Element.SetProperty("echo-cancel", value)
}

func (e *GstWebrtcDsp) GetEchoCancel() (bool, error) {
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

// Enable or disable the extended filter.
// Default: true
func (e *GstWebrtcDsp) SetExtendedFilter(value bool) error {
	return e.Element.SetProperty("extended-filter", value)
}

func (e *GstWebrtcDsp) GetExtendedFilter() (bool, error) {
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

// Enable or disable automatic digital gain control
// Default: true
func (e *GstWebrtcDsp) SetGainControl(value bool) error {
	return e.Element.SetProperty("gain-control", value)
}

func (e *GstWebrtcDsp) GetGainControl() (bool, error) {
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

// Enable or disable the voice activity detector
// Default: false
func (e *GstWebrtcDsp) SetVoiceDetection(value bool) error {
	return e.Element.SetProperty("voice-detection", value)
}

func (e *GstWebrtcDsp) GetVoiceDetection() (bool, error) {
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

// Specifies the likelihood that a frame will be declared to contain voice.
// Default: low (1)
func (e *GstWebrtcDsp) SetVoiceDetectionLikelihood(value GstWebrtcVoiceDetectionLikelihood) error {
	e.Element.SetArg("voice-detection-likelihood", string(value))
	return nil
}

func (e *GstWebrtcDsp) GetVoiceDetectionLikelihood() (GstWebrtcVoiceDetectionLikelihood, error) {
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

// Enable or disable the delay agnostic mode.
// Default: false
func (e *GstWebrtcDsp) SetDelayAgnostic(value bool) error {
	return e.Element.SetProperty("delay-agnostic", value)
}

func (e *GstWebrtcDsp) GetDelayAgnostic() (bool, error) {
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

// Enable or disable high pass filtering
// Default: true
func (e *GstWebrtcDsp) SetHighPassFilter(value bool) error {
	return e.Element.SetProperty("high-pass-filter", value)
}

func (e *GstWebrtcDsp) GetHighPassFilter() (bool, error) {
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

// When enabled, the compression stage will hard limit the signal to the target level. Otherwise, the signal will be compressed but not limited above the target level.
// Default: true
func (e *GstWebrtcDsp) SetLimiter(value bool) error {
	return e.Element.SetProperty("limiter", value)
}

func (e *GstWebrtcDsp) GetLimiter() (bool, error) {
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

// Enable or disable noise suppression
// Default: true
func (e *GstWebrtcDsp) SetNoiseSuppression(value bool) error {
	return e.Element.SetProperty("noise-suppression", value)
}

func (e *GstWebrtcDsp) GetNoiseSuppression() (bool, error) {
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

// At startup the experimental AGC moves the microphone volume up to |startup_min_volume| if the current microphone volume is set too low. No effect if experimental-agc isn't enabled.
// Default: 12, Min: 12, Max: 255
func (e *GstWebrtcDsp) SetStartupMinVolume(value int) error {
	return e.Element.SetProperty("startup-min-volume", value)
}

func (e *GstWebrtcDsp) GetStartupMinVolume() (int, error) {
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

// Sets the |size| of the frames in ms on which the VAD will operate. Larger frames will improve detection accuracy, but reduce the frequency of updates
// Default: 10, Min: 10, Max: 30
func (e *GstWebrtcDsp) SetVoiceDetectionFrameSizeMs(value int) error {
	return e.Element.SetProperty("voice-detection-frame-size-ms", value)
}

func (e *GstWebrtcDsp) GetVoiceDetectionFrameSizeMs() (int, error) {
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

// Gathers playback buffers for webrtcdsp
type GstWebrtcEchoProbe struct {
	*GstAudioFilter
}

func NewWebrtcEchoProbe() (*GstWebrtcEchoProbe, error) {
	e, err := gst.NewElement("webrtcechoprobe")
	if err != nil {
		return nil, err
	}
	return &GstWebrtcEchoProbe{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewWebrtcEchoProbeWithName(name string) (*GstWebrtcEchoProbe, error) {
	e, err := gst.NewElementWithName("webrtcechoprobe", name)
	if err != nil {
		return nil, err
	}
	return &GstWebrtcEchoProbe{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

type GstWebrtcVoiceDetectionLikelihood string

const (
	GstWebrtcVoiceDetectionLikelihoodVeryLow  GstWebrtcVoiceDetectionLikelihood = "very-low" // Very Low Likelihood
	GstWebrtcVoiceDetectionLikelihoodLow      GstWebrtcVoiceDetectionLikelihood = "low"      // Low Likelihood
	GstWebrtcVoiceDetectionLikelihoodModerate GstWebrtcVoiceDetectionLikelihood = "moderate" // Moderate Likelihood
	GstWebrtcVoiceDetectionLikelihoodHigh     GstWebrtcVoiceDetectionLikelihood = "high"     // High Likelihood
)

type GstWebrtcEchoSuppressionLevel string

const (
	GstWebrtcEchoSuppressionLevelLow      GstWebrtcEchoSuppressionLevel = "low"      // Low Suppression
	GstWebrtcEchoSuppressionLevelModerate GstWebrtcEchoSuppressionLevel = "moderate" // Moderate Suppression
	GstWebrtcEchoSuppressionLevelHigh     GstWebrtcEchoSuppressionLevel = "high"     // high Suppression
)

type GstWebrtcGainControlMode string

const (
	GstWebrtcGainControlModeAdaptiveDigital GstWebrtcGainControlMode = "adaptive-digital" // Adaptive Digital
	GstWebrtcGainControlModeFixedDigital    GstWebrtcGainControlMode = "fixed-digital"    // Fixed Digital
)

type GstWebrtcNoiseSuppressionLevel string

const (
	GstWebrtcNoiseSuppressionLevelLow      GstWebrtcNoiseSuppressionLevel = "low"       // Low Suppression
	GstWebrtcNoiseSuppressionLevelModerate GstWebrtcNoiseSuppressionLevel = "moderate"  // Moderate Suppression
	GstWebrtcNoiseSuppressionLevelHigh     GstWebrtcNoiseSuppressionLevel = "high"      // High Suppression
	GstWebrtcNoiseSuppressionLevelVeryHigh GstWebrtcNoiseSuppressionLevel = "very-high" // Very High Suppression
)
