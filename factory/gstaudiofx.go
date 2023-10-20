// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Positions audio streams in the stereo panorama
type GstAudioPanorama struct {
	*base.GstBaseTransform
}

func NewAudioPanorama() (*GstAudioPanorama, error) {
	e, err := gst.NewElement("audiopanorama")
	if err != nil {
		return nil, err
	}
	return &GstAudioPanorama{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioPanoramaWithName(name string) (*GstAudioPanorama, error) {
	e, err := gst.NewElementWithName("audiopanorama", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioPanorama{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Psychoacoustic mode keeps same perceived loudness, simple mode just controls volume of one channel.
// Default: psychoacoustic (0)
func (e *GstAudioPanorama) SetMethod(value GstAudioPanoramaMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstAudioPanorama) GetMethod() (GstAudioPanoramaMethod, error) {
	var value GstAudioPanoramaMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioPanoramaMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioPanoramaMethod")
	}
	return value, nil
}

// Position in stereo panorama (-1.0 left -> 1.0 right)
// Default: 0, Min: -1, Max: 1
func (e *GstAudioPanorama) SetPanorama(value float32) error {
	return e.Element.SetProperty("panorama", value)
}

func (e *GstAudioPanorama) GetPanorama() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panorama")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Compressor and Expander
type GstAudioDynamic struct {
	*GstAudioFilter
}

func NewAudioDynamic() (*GstAudioDynamic, error) {
	e, err := gst.NewElement("audiodynamic")
	if err != nil {
		return nil, err
	}
	return &GstAudioDynamic{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAudioDynamicWithName(name string) (*GstAudioDynamic, error) {
	e, err := gst.NewElementWithName("audiodynamic", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioDynamic{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Selects whether the ratio should be applied smooth (soft-knee) or hard (hard-knee).
// Default: hard-knee (0)
func (e *GstAudioDynamic) SetCharacteristics(value GstAudioDynamicCharacteristics) error {
	e.Element.SetArg("characteristics", string(value))
	return nil
}

func (e *GstAudioDynamic) GetCharacteristics() (GstAudioDynamicCharacteristics, error) {
	var value GstAudioDynamicCharacteristics
	var ok bool
	v, err := e.Element.GetProperty("characteristics")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioDynamicCharacteristics)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioDynamicCharacteristics")
	}
	return value, nil
}

// Selects whether the filter should work on loud samples (compressor) orquiet samples (expander).
// Default: compressor (0)
func (e *GstAudioDynamic) SetMode(value GstAudioDynamicMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstAudioDynamic) GetMode() (GstAudioDynamicMode, error) {
	var value GstAudioDynamicMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioDynamicMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioDynamicMode")
	}
	return value, nil
}

// Ratio that should be applied
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *GstAudioDynamic) SetRatio(value float32) error {
	return e.Element.SetProperty("ratio", value)
}

func (e *GstAudioDynamic) GetRatio() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Threshold until the filter is activated
// Default: 0, Min: 0, Max: 1
func (e *GstAudioDynamic) SetThreshold(value float32) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstAudioDynamic) GetThreshold() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Generic audio IIR filter with custom filter kernel
type GstAudioIIRFilter struct {
	*GstAudioFXBaseIIRFilter
}

func NewAudioIIRFilter() (*GstAudioIIRFilter, error) {
	e, err := gst.NewElement("audioiirfilter")
	if err != nil {
		return nil, err
	}
	return &GstAudioIIRFilter{GstAudioFXBaseIIRFilter: &GstAudioFXBaseIIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewAudioIIRFilterWithName(name string) (*GstAudioIIRFilter, error) {
	e, err := gst.NewElementWithName("audioiirfilter", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioIIRFilter{GstAudioFXBaseIIRFilter: &GstAudioFXBaseIIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Filter coefficients (denominator of transfer function)

func (e *GstAudioIIRFilter) SetA(value interface{}) error {
	return e.Element.SetProperty("a", value)
}

func (e *GstAudioIIRFilter) GetA() (interface{}, error) {
	return e.Element.GetProperty("a")
}

// Filter coefficients (numerator of transfer function)

func (e *GstAudioIIRFilter) SetB(value interface{}) error {
	return e.Element.SetProperty("b", value)
}

func (e *GstAudioIIRFilter) GetB() (interface{}, error) {
	return e.Element.GetProperty("b")
}

// Removes voice from sound
type GstAudioKaraoke struct {
	*GstAudioFilter
}

func NewAudioKaraoke() (*GstAudioKaraoke, error) {
	e, err := gst.NewElement("audiokaraoke")
	if err != nil {
		return nil, err
	}
	return &GstAudioKaraoke{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAudioKaraokeWithName(name string) (*GstAudioKaraoke, error) {
	e, err := gst.NewElementWithName("audiokaraoke", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioKaraoke{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The Frequency band of the filter
// Default: 220, Min: 0, Max: 441
func (e *GstAudioKaraoke) SetFilterBand(value float32) error {
	return e.Element.SetProperty("filter-band", value)
}

func (e *GstAudioKaraoke) GetFilterBand() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("filter-band")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The Frequency width of the filter
// Default: 100, Min: 0, Max: 100
func (e *GstAudioKaraoke) SetFilterWidth(value float32) error {
	return e.Element.SetProperty("filter-width", value)
}

func (e *GstAudioKaraoke) GetFilterWidth() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("filter-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Level of the effect (1.0 = full)
// Default: 1, Min: 0, Max: 1
func (e *GstAudioKaraoke) SetLevel(value float32) error {
	return e.Element.SetProperty("level", value)
}

func (e *GstAudioKaraoke) GetLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Level of the mono channel (1.0 = full)
// Default: 1, Min: 0, Max: 1
func (e *GstAudioKaraoke) SetMonoLevel(value float32) error {
	return e.Element.SetProperty("mono-level", value)
}

func (e *GstAudioKaraoke) GetMonoLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mono-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Band pass and band reject windowed sinc filter
type GstAudioWSincBand struct {
	*GstAudioFXBaseFIRFilter
}

func NewAudioWSincBand() (*GstAudioWSincBand, error) {
	e, err := gst.NewElement("audiowsincband")
	if err != nil {
		return nil, err
	}
	return &GstAudioWSincBand{GstAudioFXBaseFIRFilter: &GstAudioFXBaseFIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewAudioWSincBandWithName(name string) (*GstAudioWSincBand, error) {
	e, err := gst.NewElementWithName("audiowsincband", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioWSincBand{GstAudioFXBaseFIRFilter: &GstAudioFXBaseFIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Filter kernel length, will be rounded to the next odd number
// Default: 101, Min: 3, Max: 256000
func (e *GstAudioWSincBand) SetLength(value int) error {
	return e.Element.SetProperty("length", value)
}

func (e *GstAudioWSincBand) GetLength() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Cut-off lower frequency (Hz)
// Default: 0, Min: 0, Max: 100000
func (e *GstAudioWSincBand) SetLowerFrequency(value float32) error {
	return e.Element.SetProperty("lower-frequency", value)
}

func (e *GstAudioWSincBand) GetLowerFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lower-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Band pass or band reject mode
// Default: band-pass (0)
func (e *GstAudioWSincBand) SetMode(value GstAudioWSincBandMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstAudioWSincBand) GetMode() (GstAudioWSincBandMode, error) {
	var value GstAudioWSincBandMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWSincBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWSincBandMode")
	}
	return value, nil
}

// Cut-off upper frequency (Hz)
// Default: 0, Min: 0, Max: 100000
func (e *GstAudioWSincBand) SetUpperFrequency(value float32) error {
	return e.Element.SetProperty("upper-frequency", value)
}

func (e *GstAudioWSincBand) GetUpperFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("upper-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Window function to use
// Default: hamming (0)
func (e *GstAudioWSincBand) SetWindow(value GstAudioWSincBandWindow) error {
	e.Element.SetArg("window", string(value))
	return nil
}

func (e *GstAudioWSincBand) GetWindow() (GstAudioWSincBandWindow, error) {
	var value GstAudioWSincBandWindow
	var ok bool
	v, err := e.Element.GetProperty("window")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWSincBandWindow)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWSincBandWindow")
	}
	return value, nil
}

// Low pass and high pass windowed sinc filter
type GstAudioWSincLimit struct {
	*GstAudioFXBaseFIRFilter
}

func NewAudioWSincLimit() (*GstAudioWSincLimit, error) {
	e, err := gst.NewElement("audiowsinclimit")
	if err != nil {
		return nil, err
	}
	return &GstAudioWSincLimit{GstAudioFXBaseFIRFilter: &GstAudioFXBaseFIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewAudioWSincLimitWithName(name string) (*GstAudioWSincLimit, error) {
	e, err := gst.NewElementWithName("audiowsinclimit", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioWSincLimit{GstAudioFXBaseFIRFilter: &GstAudioFXBaseFIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Low pass or high pass mode
// Default: low-pass (0)
func (e *GstAudioWSincLimit) SetMode(value GstAudioWSincLimitMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstAudioWSincLimit) GetMode() (GstAudioWSincLimitMode, error) {
	var value GstAudioWSincLimitMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWSincLimitMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWSincLimitMode")
	}
	return value, nil
}

// Window function to use
// Default: hamming (0)
func (e *GstAudioWSincLimit) SetWindow(value GstAudioWSincLimitWindow) error {
	e.Element.SetArg("window", string(value))
	return nil
}

func (e *GstAudioWSincLimit) GetWindow() (GstAudioWSincLimitWindow, error) {
	var value GstAudioWSincLimitWindow
	var ok bool
	v, err := e.Element.GetProperty("window")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWSincLimitWindow)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWSincLimitWindow")
	}
	return value, nil
}

// Cut-off Frequency (Hz)
// Default: 0, Min: 0, Max: 100000
func (e *GstAudioWSincLimit) SetCutoff(value float32) error {
	return e.Element.SetProperty("cutoff", value)
}

func (e *GstAudioWSincLimit) GetCutoff() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cutoff")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Filter kernel length, will be rounded to the next odd number
// Default: 101, Min: 3, Max: 256000
func (e *GstAudioWSincLimit) SetLength(value int) error {
	return e.Element.SetProperty("length", value)
}

func (e *GstAudioWSincLimit) GetLength() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sync audio tempo with playback rate
type GstScaletempo struct {
	*base.GstBaseTransform
}

func NewScaletempo() (*GstScaletempo, error) {
	e, err := gst.NewElement("scaletempo")
	if err != nil {
		return nil, err
	}
	return &GstScaletempo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewScaletempoWithName(name string) (*GstScaletempo, error) {
	e, err := gst.NewElementWithName("scaletempo", name)
	if err != nil {
		return nil, err
	}
	return &GstScaletempo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Percentage of stride to overlap
// Default: 0.2, Min: 0, Max: 1
func (e *GstScaletempo) SetOverlap(value float64) error {
	return e.Element.SetProperty("overlap", value)
}

func (e *GstScaletempo) GetOverlap() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("overlap")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Current playback rate
// Default: 0, Min: -2.14748e+09, Max: 2.14748e+09
func (e *GstScaletempo) GetRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Length in milliseconds to search for best overlap position
// Default: 14, Min: 0, Max: 500
func (e *GstScaletempo) SetSearch(value uint) error {
	return e.Element.SetProperty("search", value)
}

func (e *GstScaletempo) GetSearch() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("search")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Length in milliseconds to output each stride
// Default: 30, Min: 1, Max: 5000
func (e *GstScaletempo) SetStride(value uint) error {
	return e.Element.SetProperty("stride", value)
}

func (e *GstScaletempo) GetStride() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("stride")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Chebyshev band pass and band reject filter
type GstAudioChebBand struct {
	*GstAudioFXBaseIIRFilter
}

func NewAudioChebBand() (*GstAudioChebBand, error) {
	e, err := gst.NewElement("audiochebband")
	if err != nil {
		return nil, err
	}
	return &GstAudioChebBand{GstAudioFXBaseIIRFilter: &GstAudioFXBaseIIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewAudioChebBandWithName(name string) (*GstAudioChebBand, error) {
	e, err := gst.NewElementWithName("audiochebband", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioChebBand{GstAudioFXBaseIIRFilter: &GstAudioFXBaseIIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Start frequency of the band (Hz)
// Default: 0, Min: 0, Max: 100000
func (e *GstAudioChebBand) SetLowerFrequency(value float32) error {
	return e.Element.SetProperty("lower-frequency", value)
}

func (e *GstAudioChebBand) GetLowerFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lower-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Low pass or high pass mode
// Default: band-pass (0)
func (e *GstAudioChebBand) SetMode(value GstAudioChebBandMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstAudioChebBand) GetMode() (GstAudioChebBandMode, error) {
	var value GstAudioChebBandMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioChebBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioChebBandMode")
	}
	return value, nil
}

// Number of poles to use, will be rounded up to the next multiply of four
// Default: 4, Min: 4, Max: 32
func (e *GstAudioChebBand) SetPoles(value int) error {
	return e.Element.SetProperty("poles", value)
}

func (e *GstAudioChebBand) GetPoles() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("poles")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Amount of ripple (dB)
// Default: 0.25, Min: 0, Max: 200
func (e *GstAudioChebBand) SetRipple(value float32) error {
	return e.Element.SetProperty("ripple", value)
}

func (e *GstAudioChebBand) GetRipple() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ripple")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Type of the chebychev filter
// Default: 1, Min: 1, Max: 2
func (e *GstAudioChebBand) SetType(value int) error {
	return e.Element.SetProperty("type", value)
}

func (e *GstAudioChebBand) GetType() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Stop frequency of the band (Hz)
// Default: 0, Min: 0, Max: 100000
func (e *GstAudioChebBand) SetUpperFrequency(value float32) error {
	return e.Element.SetProperty("upper-frequency", value)
}

func (e *GstAudioChebBand) GetUpperFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("upper-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Swaps upper and lower half of audio samples
type GstAudioInvert struct {
	*GstAudioFilter
}

func NewAudioInvert() (*GstAudioInvert, error) {
	e, err := gst.NewElement("audioinvert")
	if err != nil {
		return nil, err
	}
	return &GstAudioInvert{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAudioInvertWithName(name string) (*GstAudioInvert, error) {
	e, err := gst.NewElementWithName("audioinvert", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioInvert{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Degree of inversion
// Default: 0, Min: 0, Max: 1
func (e *GstAudioInvert) SetDegree(value float32) error {
	return e.Element.SetProperty("degree", value)
}

func (e *GstAudioInvert) GetDegree() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("degree")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Generic audio FIR filter with custom filter kernel
type GstAudioFIRFilter struct {
	*GstAudioFXBaseFIRFilter
}

func NewAudioFIRFilter() (*GstAudioFIRFilter, error) {
	e, err := gst.NewElement("audiofirfilter")
	if err != nil {
		return nil, err
	}
	return &GstAudioFIRFilter{GstAudioFXBaseFIRFilter: &GstAudioFXBaseFIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewAudioFIRFilterWithName(name string) (*GstAudioFIRFilter, error) {
	e, err := gst.NewElementWithName("audiofirfilter", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioFIRFilter{GstAudioFXBaseFIRFilter: &GstAudioFXBaseFIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Filter latency in samples
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAudioFIRFilter) SetLatency(value uint64) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstAudioFIRFilter) GetLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Filter kernel for the FIR filter

func (e *GstAudioFIRFilter) SetKernel(value interface{}) error {
	return e.Element.SetProperty("kernel", value)
}

func (e *GstAudioFIRFilter) GetKernel() (interface{}, error) {
	return e.Element.GetProperty("kernel")
}

// Muck with the stereo signal to enhance its 'stereo-ness'
type GstStereo struct {
	*GstAudioFilter
}

func NewStereo() (*GstStereo, error) {
	e, err := gst.NewElement("stereo")
	if err != nil {
		return nil, err
	}
	return &GstStereo{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewStereoWithName(name string) (*GstStereo, error) {
	e, err := gst.NewElementWithName("stereo", name)
	if err != nil {
		return nil, err
	}
	return &GstStereo{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// stereo
// Default: 0.01, Min: 0, Max: 1
func (e *GstStereo) SetStereo(value float32) error {
	return e.Element.SetProperty("stereo", value)
}

func (e *GstStereo) GetStereo() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("stereo")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// active
// Default: true
func (e *GstStereo) SetActive(value bool) error {
	return e.Element.SetProperty("active", value)
}

func (e *GstStereo) GetActive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("active")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adds an echo or reverb effect to an audio stream
type GstAudioEcho struct {
	*GstAudioFilter
}

func NewAudioEcho() (*GstAudioEcho, error) {
	e, err := gst.NewElement("audioecho")
	if err != nil {
		return nil, err
	}
	return &GstAudioEcho{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAudioEchoWithName(name string) (*GstAudioEcho, error) {
	e, err := gst.NewElementWithName("audioecho", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioEcho{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Delay of the echo in nanoseconds
// Default: 1, Min: 1, Max: 18446744073709551615
func (e *GstAudioEcho) SetDelay(value uint64) error {
	return e.Element.SetProperty("delay", value)
}

func (e *GstAudioEcho) GetDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Amount of feedback
// Default: 0, Min: 0, Max: 1
func (e *GstAudioEcho) SetFeedback(value float32) error {
	return e.Element.SetProperty("feedback", value)
}

func (e *GstAudioEcho) GetFeedback() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("feedback")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Intensity of the echo
// Default: 0, Min: 0, Max: 1
func (e *GstAudioEcho) SetIntensity(value float32) error {
	return e.Element.SetProperty("intensity", value)
}

func (e *GstAudioEcho) GetIntensity() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Maximum delay of the echo in nanoseconds (can't be changed in PLAYING or PAUSED state)
// Default: 1, Min: 1, Max: 18446744073709551615
func (e *GstAudioEcho) SetMaxDelay(value uint64) error {
	return e.Element.SetProperty("max-delay", value)
}

func (e *GstAudioEcho) GetMaxDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Delay Surround Channels when TRUE instead of applying an echo effect
// Default: false
func (e *GstAudioEcho) SetSurroundDelay(value bool) error {
	return e.Element.SetProperty("surround-delay", value)
}

func (e *GstAudioEcho) GetSurroundDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("surround-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// A bitmask of channels that are considered surround and delayed when surround-delay = TRUE
// Default: 18446744073709551612, Min: 1, Max: 18446744073709551615
func (e *GstAudioEcho) SetSurroundMask(value uint64) error {
	return e.Element.SetProperty("surround-mask", value)
}

func (e *GstAudioEcho) GetSurroundMask() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("surround-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Amplifies an audio stream by a given factor
type GstAudioAmplify struct {
	*GstAudioFilter
}

func NewAudioAmplify() (*GstAudioAmplify, error) {
	e, err := gst.NewElement("audioamplify")
	if err != nil {
		return nil, err
	}
	return &GstAudioAmplify{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAudioAmplifyWithName(name string) (*GstAudioAmplify, error) {
	e, err := gst.NewElementWithName("audioamplify", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioAmplify{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Factor of amplification
// Default: 1, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstAudioAmplify) SetAmplification(value float32) error {
	return e.Element.SetProperty("amplification", value)
}

func (e *GstAudioAmplify) GetAmplification() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("amplification")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Selects how to handle values higher than the maximum
// Default: clip (0)
func (e *GstAudioAmplify) SetClippingMethod(value GstAudioAmplifyClippingMethod) error {
	e.Element.SetArg("clipping-method", string(value))
	return nil
}

func (e *GstAudioAmplify) GetClippingMethod() (GstAudioAmplifyClippingMethod, error) {
	var value GstAudioAmplifyClippingMethod
	var ok bool
	v, err := e.Element.GetProperty("clipping-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioAmplifyClippingMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioAmplifyClippingMethod")
	}
	return value, nil
}

// Chebyshev low pass and high pass filter
type GstAudioChebLimit struct {
	*GstAudioFXBaseIIRFilter
}

func NewAudioChebLimit() (*GstAudioChebLimit, error) {
	e, err := gst.NewElement("audiocheblimit")
	if err != nil {
		return nil, err
	}
	return &GstAudioChebLimit{GstAudioFXBaseIIRFilter: &GstAudioFXBaseIIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewAudioChebLimitWithName(name string) (*GstAudioChebLimit, error) {
	e, err := gst.NewElementWithName("audiocheblimit", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioChebLimit{GstAudioFXBaseIIRFilter: &GstAudioFXBaseIIRFilter{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Cut off frequency (Hz)
// Default: 0, Min: 0, Max: 100000
func (e *GstAudioChebLimit) SetCutoff(value float32) error {
	return e.Element.SetProperty("cutoff", value)
}

func (e *GstAudioChebLimit) GetCutoff() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cutoff")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Low pass or high pass mode
// Default: low-pass (0)
func (e *GstAudioChebLimit) SetMode(value GstAudioChebLimitMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstAudioChebLimit) GetMode() (GstAudioChebLimitMode, error) {
	var value GstAudioChebLimitMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioChebLimitMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioChebLimitMode")
	}
	return value, nil
}

// Number of poles to use, will be rounded up to the next even number
// Default: 4, Min: 2, Max: 32
func (e *GstAudioChebLimit) SetPoles(value int) error {
	return e.Element.SetProperty("poles", value)
}

func (e *GstAudioChebLimit) GetPoles() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("poles")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Amount of ripple (dB)
// Default: 0.25, Min: 0, Max: 200
func (e *GstAudioChebLimit) SetRipple(value float32) error {
	return e.Element.SetProperty("ripple", value)
}

func (e *GstAudioChebLimit) GetRipple() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ripple")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Type of the chebychev filter
// Default: 1, Min: 1, Max: 2
func (e *GstAudioChebLimit) SetType(value int) error {
	return e.Element.SetProperty("type", value)
}

func (e *GstAudioChebLimit) GetType() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstAudioFXBaseFIRFilter struct {
	*GstAudioFilter
}

// Drains the filter when its coefficients change
// Default: true
func (e *GstAudioFXBaseFIRFilter) SetDrainOnChanges(value bool) error {
	return e.Element.SetProperty("drain-on-changes", value)
}

func (e *GstAudioFXBaseFIRFilter) GetDrainOnChanges() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drain-on-changes")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Operate in low latency mode. This mode is slower but the latency will only be the filter pre-latency. Can only be changed in states < PAUSED!
// Default: false
func (e *GstAudioFXBaseFIRFilter) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstAudioFXBaseFIRFilter) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstAudioFXBaseIIRFilter struct {
	*GstAudioFilter
}

type GstAudioWSincBandWindow string

const (
	GstAudioWSincBandWindowHamming  GstAudioWSincBandWindow = "hamming"  // Hamming window (default)
	GstAudioWSincBandWindowBlackman GstAudioWSincBandWindow = "blackman" // Blackman window
	GstAudioWSincBandWindowGaussian GstAudioWSincBandWindow = "gaussian" // Gaussian window
	GstAudioWSincBandWindowCosine   GstAudioWSincBandWindow = "cosine"   // Cosine window
	GstAudioWSincBandWindowHann     GstAudioWSincBandWindow = "hann"     // Hann window
)

type GstAudioWSincLimitWindow string

const (
	GstAudioWSincLimitWindowHamming  GstAudioWSincLimitWindow = "hamming"  // Hamming window (default)
	GstAudioWSincLimitWindowBlackman GstAudioWSincLimitWindow = "blackman" // Blackman window
	GstAudioWSincLimitWindowGaussian GstAudioWSincLimitWindow = "gaussian" // Gaussian window
	GstAudioWSincLimitWindowCosine   GstAudioWSincLimitWindow = "cosine"   // Cosine window
	GstAudioWSincLimitWindowHann     GstAudioWSincLimitWindow = "hann"     // Hann window
)

type GstAudioAmplifyClippingMethod string

const (
	GstAudioAmplifyClippingMethodClip         GstAudioAmplifyClippingMethod = "clip"          // Normal clipping (default)
	GstAudioAmplifyClippingMethodWrapNegative GstAudioAmplifyClippingMethod = "wrap-negative" // Push overdriven values back from the opposite side
	GstAudioAmplifyClippingMethodWrapPositive GstAudioAmplifyClippingMethod = "wrap-positive" // Push overdriven values back from the same side
	GstAudioAmplifyClippingMethodNone         GstAudioAmplifyClippingMethod = "none"          // No clipping
)

type GstAudioChebBandMode string

const (
	GstAudioChebBandModeBandPass   GstAudioChebBandMode = "band-pass"   // Band pass (default)
	GstAudioChebBandModeBandReject GstAudioChebBandMode = "band-reject" // Band reject
)

type GstAudioChebLimitMode string

const (
	GstAudioChebLimitModeLowPass  GstAudioChebLimitMode = "low-pass"  // Low pass (default)
	GstAudioChebLimitModeHighPass GstAudioChebLimitMode = "high-pass" // High pass
)

type GstAudioDynamicMode string

const (
	GstAudioDynamicModeCompressor GstAudioDynamicMode = "compressor" // Compressor (default)
	GstAudioDynamicModeExpander   GstAudioDynamicMode = "expander"   // Expander
)

type GstAudioDynamicCharacteristics string

const (
	GstAudioDynamicCharacteristicsHardKnee GstAudioDynamicCharacteristics = "hard-knee" // Hard Knee (default)
	GstAudioDynamicCharacteristicsSoftKnee GstAudioDynamicCharacteristics = "soft-knee" // Soft Knee (smooth)
)

type GstAudioPanoramaMethod string

const (
	GstAudioPanoramaMethodPsychoacoustic GstAudioPanoramaMethod = "psychoacoustic" // Psychoacoustic Panning (default)
	GstAudioPanoramaMethodSimple         GstAudioPanoramaMethod = "simple"         // Simple Panning
)

type GstAudioWSincBandMode string

const (
	GstAudioWSincBandModeBandPass   GstAudioWSincBandMode = "band-pass"   // Band pass (default)
	GstAudioWSincBandModeBandReject GstAudioWSincBandMode = "band-reject" // Band reject
)

type GstAudioWSincLimitMode string

const (
	GstAudioWSincLimitModeLowPass  GstAudioWSincLimitMode = "low-pass"  // Low pass (default)
	GstAudioWSincLimitModeHighPass GstAudioWSincLimitMode = "high-pass" // High pass
)
