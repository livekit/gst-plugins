// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes FLAC lossless audio streams
type GstFlacDec struct {
	*GstAudioDecoder
}

func NewFlacDec() (*GstFlacDec, error) {
	e, err := gst.NewElement("flacdec")
	if err != nil {
		return nil, err
	}
	return &GstFlacDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewFlacDecWithName(name string) (*GstFlacDec, error) {
	e, err := gst.NewElementWithName("flacdec", name)
	if err != nil {
		return nil, err
	}
	return &GstFlacDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Encodes audio with the FLAC lossless audio encoder
type GstFlacEnc struct {
	*GstAudioEncoder
}

func NewFlacEnc() (*GstFlacEnc, error) {
	e, err := gst.NewElement("flacenc")
	if err != nil {
		return nil, err
	}
	return &GstFlacEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewFlacEncWithName(name string) (*GstFlacEnc, error) {
	e, err := gst.NewElementWithName("flacenc", name)
	if err != nil {
		return nil, err
	}
	return &GstFlacEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Max LPC order; 0 => use only fixed predictors
// Default: 8, Min: 0, Max: 32
func (e *GstFlacEnc) SetMaxLpcOrder(value uint) error {
	return e.Element.SetProperty("max-lpc-order", value)
}

func (e *GstFlacEnc) GetMaxLpcOrder() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-lpc-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Add SEEKTABLE metadata (if > 0, number of entries, if < 0, interval in sec)
// Default: -10, Min: -2147483647, Max: 2147483647
func (e *GstFlacEnc) SetSeekpoints(value int) error {
	return e.Element.SetProperty("seekpoints", value)
}

func (e *GstFlacEnc) GetSeekpoints() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("seekpoints")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Blocksize in samples
// Default: 4608, Min: 16, Max: 65535
func (e *GstFlacEnc) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstFlacEnc) GetBlocksize() (uint, error) {
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

// Max residual partition order (above 4 doesn't usually help much)
// Default: 3, Min: 0, Max: 16
func (e *GstFlacEnc) SetMaxResidualPartitionOrder(value uint) error {
	return e.Element.SetProperty("max-residual-partition-order", value)
}

func (e *GstFlacEnc) GetMaxResidualPartitionOrder() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-residual-partition-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Min residual partition order (above 4 doesn't usually help much)
// Default: 3, Min: 0, Max: 16
func (e *GstFlacEnc) SetMinResidualPartitionOrder(value uint) error {
	return e.Element.SetProperty("min-residual-partition-order", value)
}

func (e *GstFlacEnc) GetMinResidualPartitionOrder() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-residual-partition-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Precision in bits of quantized linear-predictor coefficients; 0 = automatic
// Default: 0, Min: 0, Max: 32
func (e *GstFlacEnc) SetQlpCoeffPrecision(value uint) error {
	return e.Element.SetProperty("qlp-coeff-precision", value)
}

func (e *GstFlacEnc) GetQlpCoeffPrecision() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qlp-coeff-precision")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Loose mid side stereo
// Default: false
func (e *GstFlacEnc) SetLooseMidSideStereo(value bool) error {
	return e.Element.SetProperty("loose-mid-side-stereo", value)
}

func (e *GstFlacEnc) GetLooseMidSideStereo() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loose-mid-side-stereo")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// 0 = try only calc'd parameter k; else try all [k-dist..k+dist] parameters, use best
// Default: 0, Min: 0, Max: 15
func (e *GstFlacEnc) SetRiceParameterSearchDist(value uint) error {
	return e.Element.SetProperty("rice-parameter-search-dist", value)
}

func (e *GstFlacEnc) GetRiceParameterSearchDist() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rice-parameter-search-dist")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// search for escape codes in the entropy coding stage for slightly better compression
// Default: false
func (e *GstFlacEnc) SetEscapeCoding(value bool) error {
	return e.Element.SetProperty("escape-coding", value)
}

func (e *GstFlacEnc) GetEscapeCoding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("escape-coding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// do exhaustive search of LP coefficient quantization (expensive!)
// Default: false
func (e *GstFlacEnc) SetExhaustiveModelSearch(value bool) error {
	return e.Element.SetProperty("exhaustive-model-search", value)
}

func (e *GstFlacEnc) GetExhaustiveModelSearch() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("exhaustive-model-search")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Do mid side stereo (only for stereo input)
// Default: true
func (e *GstFlacEnc) SetMidSideStereo(value bool) error {
	return e.Element.SetProperty("mid-side-stereo", value)
}

func (e *GstFlacEnc) GetMidSideStereo() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mid-side-stereo")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Write a PADDING block with this length in bytes
// Default: 0, Min: 0, Max: -1
func (e *GstFlacEnc) SetPadding(value uint) error {
	return e.Element.SetProperty("padding", value)
}

func (e *GstFlacEnc) GetPadding() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// false = use qlp_coeff_precision, true = search around qlp_coeff_precision, take best
// Default: false
func (e *GstFlacEnc) SetQlpCoeffPrecSearch(value bool) error {
	return e.Element.SetProperty("qlp-coeff-prec-search", value)
}

func (e *GstFlacEnc) GetQlpCoeffPrecSearch() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qlp-coeff-prec-search")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Speed versus compression tradeoff
// Default: 5 (5)
func (e *GstFlacEnc) SetQuality(value GstFlacEncQuality) error {
	e.Element.SetArg("quality", string(value))
	return nil
}

func (e *GstFlacEnc) GetQuality() (GstFlacEncQuality, error) {
	var value GstFlacEncQuality
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFlacEncQuality)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFlacEncQuality")
	}
	return value, nil
}

// true to limit encoder to generating a Subset stream, else false
// Default: true
func (e *GstFlacEnc) SetStreamableSubset(value bool) error {
	return e.Element.SetProperty("streamable-subset", value)
}

func (e *GstFlacEnc) GetStreamableSubset() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable-subset")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Rewrite tags in a FLAC file
type GstFlacTag struct {
	*gst.Element
}

func NewFlacTag() (*GstFlacTag, error) {
	e, err := gst.NewElement("flactag")
	if err != nil {
		return nil, err
	}
	return &GstFlacTag{Element: e}, nil
}

func NewFlacTagWithName(name string) (*GstFlacTag, error) {
	e, err := gst.NewElementWithName("flactag", name)
	if err != nil {
		return nil, err
	}
	return &GstFlacTag{Element: e}, nil
}

type GstFlacEncQuality string

const (
	GstFlacEncQuality0 GstFlacEncQuality = "0" // 0 - Fastest compression
	GstFlacEncQuality1 GstFlacEncQuality = "1" // 1
	GstFlacEncQuality2 GstFlacEncQuality = "2" // 2
	GstFlacEncQuality3 GstFlacEncQuality = "3" // 3
	GstFlacEncQuality4 GstFlacEncQuality = "4" // 4
	GstFlacEncQuality5 GstFlacEncQuality = "5" // 5 - Default
	GstFlacEncQuality6 GstFlacEncQuality = "6" // 6
	GstFlacEncQuality7 GstFlacEncQuality = "7" // 7
	GstFlacEncQuality8 GstFlacEncQuality = "8" // 8 - Highest compression
	GstFlacEncQuality9 GstFlacEncQuality = "9" // 9 - Insane
)
