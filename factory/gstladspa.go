// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Mono Amplifier
type LadspaAmpSoAmpMono struct {
	*GstLADSPAFilter
}

func NewLadspaAmpSoAmpMono() (*LadspaAmpSoAmpMono, error) {
	e, err := gst.NewElement("ladspa-amp-so-amp-mono")
	if err != nil {
		return nil, err
	}
	return &LadspaAmpSoAmpMono{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaAmpSoAmpMonoWithName(name string) (*LadspaAmpSoAmpMono, error) {
	e, err := gst.NewElementWithName("ladspa-amp-so-amp-mono", name)
	if err != nil {
		return nil, err
	}
	return &LadspaAmpSoAmpMono{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Gain
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *LadspaAmpSoAmpMono) SetGain(value float32) error {
	return e.Element.SetProperty("gain", value)
}

func (e *LadspaAmpSoAmpMono) GetGain() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Stereo Amplifier
type LadspaAmpSoAmpStereo struct {
	*GstLADSPAFilter
}

func NewLadspaAmpSoAmpStereo() (*LadspaAmpSoAmpStereo, error) {
	e, err := gst.NewElement("ladspa-amp-so-amp-stereo")
	if err != nil {
		return nil, err
	}
	return &LadspaAmpSoAmpStereo{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaAmpSoAmpStereoWithName(name string) (*LadspaAmpSoAmpStereo, error) {
	e, err := gst.NewElementWithName("ladspa-amp-so-amp-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LadspaAmpSoAmpStereo{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Gain
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *LadspaAmpSoAmpStereo) SetGain(value float32) error {
	return e.Element.SetProperty("gain", value)
}

func (e *LadspaAmpSoAmpStereo) GetGain() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Simple High Pass Filter
type LadspaFilterSoHpf struct {
	*GstLADSPAFilter
}

func NewLadspaFilterSoHpf() (*LadspaFilterSoHpf, error) {
	e, err := gst.NewElement("ladspa-filter-so-hpf")
	if err != nil {
		return nil, err
	}
	return &LadspaFilterSoHpf{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaFilterSoHpfWithName(name string) (*LadspaFilterSoHpf, error) {
	e, err := gst.NewElementWithName("ladspa-filter-so-hpf", name)
	if err != nil {
		return nil, err
	}
	return &LadspaFilterSoHpf{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Cutoff Frequency (Hz)
// Default: 440, Min: 0, Max: 22050
func (e *LadspaFilterSoHpf) SetCutoffFrequency(value float32) error {
	return e.Element.SetProperty("cutoff-frequency", value)
}

func (e *LadspaFilterSoHpf) GetCutoffFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cutoff-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Sine Oscillator (Freq:audio, Amp:control)
type LadspaSineSoSineFaac struct {
	*GstLADSPAFilter
}

func NewLadspaSineSoSineFaac() (*LadspaSineSoSineFaac, error) {
	e, err := gst.NewElement("ladspa-sine-so-sine-faac")
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFaac{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaSineSoSineFaacWithName(name string) (*LadspaSineSoSineFaac, error) {
	e, err := gst.NewElementWithName("ladspa-sine-so-sine-faac", name)
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFaac{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Amplitude
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *LadspaSineSoSineFaac) SetAmplitude(value float32) error {
	return e.Element.SetProperty("amplitude", value)
}

func (e *LadspaSineSoSineFaac) GetAmplitude() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Sine Oscillator (Freq:control, Amp:audio)
type LadspaSineSoSineFcaa struct {
	*GstLADSPAFilter
}

func NewLadspaSineSoSineFcaa() (*LadspaSineSoSineFcaa, error) {
	e, err := gst.NewElement("ladspa-sine-so-sine-fcaa")
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFcaa{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaSineSoSineFcaaWithName(name string) (*LadspaSineSoSineFcaa, error) {
	e, err := gst.NewElementWithName("ladspa-sine-so-sine-fcaa", name)
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFcaa{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Frequency (Hz)
// Default: 440, Min: 0, Max: 22050
func (e *LadspaSineSoSineFcaa) SetFrequency(value float32) error {
	return e.Element.SetProperty("frequency", value)
}

func (e *LadspaSineSoSineFcaa) GetFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// White Noise Source
type LadspasrcNoiseSoNoiseWhite struct {
	*GstLADSPASource
}

func NewLadspasrcNoiseSoNoiseWhite() (*LadspasrcNoiseSoNoiseWhite, error) {
	e, err := gst.NewElement("ladspasrc-noise-so-noise-white")
	if err != nil {
		return nil, err
	}
	return &LadspasrcNoiseSoNoiseWhite{GstLADSPASource: &GstLADSPASource{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewLadspasrcNoiseSoNoiseWhiteWithName(name string) (*LadspasrcNoiseSoNoiseWhite, error) {
	e, err := gst.NewElementWithName("ladspasrc-noise-so-noise-white", name)
	if err != nil {
		return nil, err
	}
	return &LadspasrcNoiseSoNoiseWhite{GstLADSPASource: &GstLADSPASource{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Can activate in push mode
// Default: true
func (e *LadspasrcNoiseSoNoiseWhite) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

func (e *LadspasrcNoiseSoNoiseWhite) GetCanActivatePush() (bool, error) {
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

// Whether to act as a live source
// Default: false
func (e *LadspasrcNoiseSoNoiseWhite) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *LadspasrcNoiseSoNoiseWhite) GetIsLive() (bool, error) {
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

// Number of samples in each outgoing buffer
// Default: 1024, Min: 1, Max: 2147483647
func (e *LadspasrcNoiseSoNoiseWhite) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

func (e *LadspasrcNoiseSoNoiseWhite) GetSamplesperbuffer() (int, error) {
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

// An offset added to timestamps set on buffers (in ns)
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *LadspasrcNoiseSoNoiseWhite) SetTimestampOffset(value int64) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

func (e *LadspasrcNoiseSoNoiseWhite) GetTimestampOffset() (int64, error) {
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

// Amplitude
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *LadspasrcNoiseSoNoiseWhite) SetAmplitude(value float32) error {
	return e.Element.SetProperty("amplitude", value)
}

func (e *LadspasrcNoiseSoNoiseWhite) GetAmplitude() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Can activate in pull mode
// Default: false
func (e *LadspasrcNoiseSoNoiseWhite) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *LadspasrcNoiseSoNoiseWhite) GetCanActivatePull() (bool, error) {
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

// Sine Oscillator (Freq:control, Amp:control)
type LadspasrcSineSoSineFcac struct {
	*GstLADSPASource
}

func NewLadspasrcSineSoSineFcac() (*LadspasrcSineSoSineFcac, error) {
	e, err := gst.NewElement("ladspasrc-sine-so-sine-fcac")
	if err != nil {
		return nil, err
	}
	return &LadspasrcSineSoSineFcac{GstLADSPASource: &GstLADSPASource{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewLadspasrcSineSoSineFcacWithName(name string) (*LadspasrcSineSoSineFcac, error) {
	e, err := gst.NewElementWithName("ladspasrc-sine-so-sine-fcac", name)
	if err != nil {
		return nil, err
	}
	return &LadspasrcSineSoSineFcac{GstLADSPASource: &GstLADSPASource{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Frequency (Hz)
// Default: 440, Min: 0, Max: 22050
func (e *LadspasrcSineSoSineFcac) SetFrequency(value float32) error {
	return e.Element.SetProperty("frequency", value)
}

func (e *LadspasrcSineSoSineFcac) GetFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Whether to act as a live source
// Default: false
func (e *LadspasrcSineSoSineFcac) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *LadspasrcSineSoSineFcac) GetIsLive() (bool, error) {
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

// Number of samples in each outgoing buffer
// Default: 1024, Min: 1, Max: 2147483647
func (e *LadspasrcSineSoSineFcac) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

func (e *LadspasrcSineSoSineFcac) GetSamplesperbuffer() (int, error) {
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

// An offset added to timestamps set on buffers (in ns)
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *LadspasrcSineSoSineFcac) SetTimestampOffset(value int64) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

func (e *LadspasrcSineSoSineFcac) GetTimestampOffset() (int64, error) {
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

// Amplitude
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *LadspasrcSineSoSineFcac) SetAmplitude(value float32) error {
	return e.Element.SetProperty("amplitude", value)
}

func (e *LadspasrcSineSoSineFcac) GetAmplitude() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Can activate in pull mode
// Default: false
func (e *LadspasrcSineSoSineFcac) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *LadspasrcSineSoSineFcac) GetCanActivatePull() (bool, error) {
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

// Can activate in push mode
// Default: true
func (e *LadspasrcSineSoSineFcac) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

func (e *LadspasrcSineSoSineFcac) GetCanActivatePush() (bool, error) {
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

// Simple Delay Line
type LadspaDelaySoDelay5s struct {
	*GstLADSPAFilter
}

func NewLadspaDelaySoDelay5s() (*LadspaDelaySoDelay5s, error) {
	e, err := gst.NewElement("ladspa-delay-so-delay-5s")
	if err != nil {
		return nil, err
	}
	return &LadspaDelaySoDelay5s{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaDelaySoDelay5sWithName(name string) (*LadspaDelaySoDelay5s, error) {
	e, err := gst.NewElementWithName("ladspa-delay-so-delay-5s", name)
	if err != nil {
		return nil, err
	}
	return &LadspaDelaySoDelay5s{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Delay (Seconds)
// Default: 1, Min: 0, Max: 5
func (e *LadspaDelaySoDelay5s) SetDelay(value float32) error {
	return e.Element.SetProperty("delay", value)
}

func (e *LadspaDelaySoDelay5s) GetDelay() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Dry/Wet Balance
// Default: 0.5, Min: 0, Max: 1
func (e *LadspaDelaySoDelay5s) SetDryWetBalance(value float32) error {
	return e.Element.SetProperty("dry-wet-balance", value)
}

func (e *LadspaDelaySoDelay5s) GetDryWetBalance() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dry-wet-balance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Simple Low Pass Filter
type LadspaFilterSoLpf struct {
	*GstLADSPAFilter
}

func NewLadspaFilterSoLpf() (*LadspaFilterSoLpf, error) {
	e, err := gst.NewElement("ladspa-filter-so-lpf")
	if err != nil {
		return nil, err
	}
	return &LadspaFilterSoLpf{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaFilterSoLpfWithName(name string) (*LadspaFilterSoLpf, error) {
	e, err := gst.NewElementWithName("ladspa-filter-so-lpf", name)
	if err != nil {
		return nil, err
	}
	return &LadspaFilterSoLpf{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Cutoff Frequency (Hz)
// Default: 440, Min: 0, Max: 22050
func (e *LadspaFilterSoLpf) SetCutoffFrequency(value float32) error {
	return e.Element.SetProperty("cutoff-frequency", value)
}

func (e *LadspaFilterSoLpf) GetCutoffFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cutoff-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Sine Oscillator (Freq:audio, Amp:audio)
type LadspaSineSoSineFaaa struct {
	*GstLADSPAFilter
}

func NewLadspaSineSoSineFaaa() (*LadspaSineSoSineFaaa, error) {
	e, err := gst.NewElement("ladspa-sine-so-sine-faaa")
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFaaa{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewLadspaSineSoSineFaaaWithName(name string) (*LadspaSineSoSineFaaa, error) {
	e, err := gst.NewElementWithName("ladspa-sine-so-sine-faaa", name)
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFaaa{GstLADSPAFilter: &GstLADSPAFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

type GstLADSPAFilter struct {
	*GstAudioFilter
}

type GstLADSPASource struct {
	*base.GstBaseSrc
}
