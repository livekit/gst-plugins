// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Direct Form IIR equalizer
type GstIirEqualizerNBands struct {
	*GstIirEqualizer
}

func NewIirEqualizerNBands() (*GstIirEqualizerNBands, error) {
	e, err := gst.NewElement("equalizer-nbands")
	if err != nil {
		return nil, err
	}
	return &GstIirEqualizerNBands{GstIirEqualizer: &GstIirEqualizer{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewIirEqualizerNBandsWithName(name string) (*GstIirEqualizerNBands, error) {
	e, err := gst.NewElementWithName("equalizer-nbands", name)
	if err != nil {
		return nil, err
	}
	return &GstIirEqualizerNBands{GstIirEqualizer: &GstIirEqualizer{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// number of different bands to use
// Default: 10, Min: 1, Max: 64
func (e *GstIirEqualizerNBands) SetNumBands(value uint) error {
	return e.Element.SetProperty("num-bands", value)
}

func (e *GstIirEqualizerNBands) GetNumBands() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-bands")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Direct Form 10 band IIR equalizer
type GstIirEqualizer10Bands struct {
	*GstIirEqualizer
}

func NewIirEqualizer10Bands() (*GstIirEqualizer10Bands, error) {
	e, err := gst.NewElement("equalizer-10bands")
	if err != nil {
		return nil, err
	}
	return &GstIirEqualizer10Bands{GstIirEqualizer: &GstIirEqualizer{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewIirEqualizer10BandsWithName(name string) (*GstIirEqualizer10Bands, error) {
	e, err := gst.NewElementWithName("equalizer-10bands", name)
	if err != nil {
		return nil, err
	}
	return &GstIirEqualizer10Bands{GstIirEqualizer: &GstIirEqualizer{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// gain for the frequency band 29 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand0(value float64) error {
	return e.Element.SetProperty("band0", value)
}

func (e *GstIirEqualizer10Bands) GetBand0() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 59 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand1(value float64) error {
	return e.Element.SetProperty("band1", value)
}

func (e *GstIirEqualizer10Bands) GetBand1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 237 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand3(value float64) error {
	return e.Element.SetProperty("band3", value)
}

func (e *GstIirEqualizer10Bands) GetBand3() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 947 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand5(value float64) error {
	return e.Element.SetProperty("band5", value)
}

func (e *GstIirEqualizer10Bands) GetBand5() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 3770 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand7(value float64) error {
	return e.Element.SetProperty("band7", value)
}

func (e *GstIirEqualizer10Bands) GetBand7() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 119 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand2(value float64) error {
	return e.Element.SetProperty("band2", value)
}

func (e *GstIirEqualizer10Bands) GetBand2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 474 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand4(value float64) error {
	return e.Element.SetProperty("band4", value)
}

func (e *GstIirEqualizer10Bands) GetBand4() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 1889 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand6(value float64) error {
	return e.Element.SetProperty("band6", value)
}

func (e *GstIirEqualizer10Bands) GetBand6() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 7523 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand8(value float64) error {
	return e.Element.SetProperty("band8", value)
}

func (e *GstIirEqualizer10Bands) GetBand8() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 15011 Hz, ranging from -24 dB to +12 dB
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer10Bands) SetBand9(value float64) error {
	return e.Element.SetProperty("band9", value)
}

func (e *GstIirEqualizer10Bands) GetBand9() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Direct Form 3 band IIR equalizer
type GstIirEqualizer3Bands struct {
	*GstIirEqualizer
}

func NewIirEqualizer3Bands() (*GstIirEqualizer3Bands, error) {
	e, err := gst.NewElement("equalizer-3bands")
	if err != nil {
		return nil, err
	}
	return &GstIirEqualizer3Bands{GstIirEqualizer: &GstIirEqualizer{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewIirEqualizer3BandsWithName(name string) (*GstIirEqualizer3Bands, error) {
	e, err := gst.NewElementWithName("equalizer-3bands", name)
	if err != nil {
		return nil, err
	}
	return &GstIirEqualizer3Bands{GstIirEqualizer: &GstIirEqualizer{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// gain for the frequency band 1100 Hz, ranging from -24.0 to +12.0
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer3Bands) SetBand1(value float64) error {
	return e.Element.SetProperty("band1", value)
}

func (e *GstIirEqualizer3Bands) GetBand1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 11 kHz, ranging from -24.0 to +12.0
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer3Bands) SetBand2(value float64) error {
	return e.Element.SetProperty("band2", value)
}

func (e *GstIirEqualizer3Bands) GetBand2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// gain for the frequency band 100 Hz, ranging from -24.0 to +12.0
// Default: 0, Min: -24, Max: 12
func (e *GstIirEqualizer3Bands) SetBand0(value float64) error {
	return e.Element.SetProperty("band0", value)
}

func (e *GstIirEqualizer3Bands) GetBand0() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

type GstIirEqualizer struct {
	*GstAudioFilter
}
