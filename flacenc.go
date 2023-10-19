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
)

type Flacenc struct {
	Element *gst.Element
}

func NewFlacenc() (*Flacenc, error) {
	e, err := gst.NewElement("flacenc")
	if err != nil {
		return nil, err
	}
	return &Flacenc{Element: e}, nil
}

func NewFlacencWithName(name string) (*Flacenc, error) {
	e, err := gst.NewElementWithName("flacenc", name)
	if err != nil {
		return nil, err
	}
	return &Flacenc{Element: e}, nil
}

// ----- Properties -----

// blocksize (uint)
//
// Blocksize in samples

func (e *Flacenc) GetBlocksize() (uint, error) {
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

func (e *Flacenc) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// escape-coding (bool)
//
// search for escape codes in the entropy coding stage for slightly better compression

func (e *Flacenc) GetEscapeCoding() (bool, error) {
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

func (e *Flacenc) SetEscapeCoding(value bool) error {
	return e.Element.SetProperty("escape-coding", value)
}

// exhaustive-model-search (bool)
//
// do exhaustive search of LP coefficient quantization (expensive!)

func (e *Flacenc) GetExhaustiveModelSearch() (bool, error) {
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

func (e *Flacenc) SetExhaustiveModelSearch(value bool) error {
	return e.Element.SetProperty("exhaustive-model-search", value)
}

// loose-mid-side-stereo (bool)
//
// Loose mid side stereo

func (e *Flacenc) GetLooseMidSideStereo() (bool, error) {
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

func (e *Flacenc) SetLooseMidSideStereo(value bool) error {
	return e.Element.SetProperty("loose-mid-side-stereo", value)
}

// max-lpc-order (uint)
//
// Max LPC order; 0 => use only fixed predictors

func (e *Flacenc) GetMaxLpcOrder() (uint, error) {
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

func (e *Flacenc) SetMaxLpcOrder(value uint) error {
	return e.Element.SetProperty("max-lpc-order", value)
}

// max-residual-partition-order (uint)
//
// Max residual partition order (above 4 doesn't usually help much)

func (e *Flacenc) GetMaxResidualPartitionOrder() (uint, error) {
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

func (e *Flacenc) SetMaxResidualPartitionOrder(value uint) error {
	return e.Element.SetProperty("max-residual-partition-order", value)
}

// mid-side-stereo (bool)
//
// Do mid side stereo (only for stereo input)

func (e *Flacenc) GetMidSideStereo() (bool, error) {
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

func (e *Flacenc) SetMidSideStereo(value bool) error {
	return e.Element.SetProperty("mid-side-stereo", value)
}

// min-residual-partition-order (uint)
//
// Min residual partition order (above 4 doesn't usually help much)

func (e *Flacenc) GetMinResidualPartitionOrder() (uint, error) {
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

func (e *Flacenc) SetMinResidualPartitionOrder(value uint) error {
	return e.Element.SetProperty("min-residual-partition-order", value)
}

// padding (uint)
//
// Write a PADDING block with this length in bytes

func (e *Flacenc) GetPadding() (uint, error) {
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

func (e *Flacenc) SetPadding(value uint) error {
	return e.Element.SetProperty("padding", value)
}

// qlp-coeff-prec-search (bool)
//
// false = use qlp_coeff_precision, true = search around qlp_coeff_precision, take best

func (e *Flacenc) GetQlpCoeffPrecSearch() (bool, error) {
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

func (e *Flacenc) SetQlpCoeffPrecSearch(value bool) error {
	return e.Element.SetProperty("qlp-coeff-prec-search", value)
}

// qlp-coeff-precision (uint)
//
// Precision in bits of quantized linear-predictor coefficients; 0 = automatic

func (e *Flacenc) GetQlpCoeffPrecision() (uint, error) {
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

func (e *Flacenc) SetQlpCoeffPrecision(value uint) error {
	return e.Element.SetProperty("qlp-coeff-precision", value)
}

// quality (GstFlacEncQuality)
//
// Speed versus compression tradeoff

func (e *Flacenc) GetQuality() (GstFlacEncQuality, error) {
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

func (e *Flacenc) SetQuality(value GstFlacEncQuality) error {
	e.Element.SetArg("quality", string(value))
	return nil
}

// rice-parameter-search-dist (uint)
//
// 0 = try only calc'd parameter k; else try all [k-dist..k+dist] parameters, use best

func (e *Flacenc) GetRiceParameterSearchDist() (uint, error) {
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

func (e *Flacenc) SetRiceParameterSearchDist(value uint) error {
	return e.Element.SetProperty("rice-parameter-search-dist", value)
}

// seekpoints (int)
//
// Add SEEKTABLE metadata (if > 0, number of entries, if < 0, interval in sec)

func (e *Flacenc) GetSeekpoints() (int, error) {
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

func (e *Flacenc) SetSeekpoints(value int) error {
	return e.Element.SetProperty("seekpoints", value)
}

// streamable-subset (bool)
//
// true to limit encoder to generating a Subset stream, else false

func (e *Flacenc) GetStreamableSubset() (bool, error) {
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

func (e *Flacenc) SetStreamableSubset(value bool) error {
	return e.Element.SetProperty("streamable-subset", value)
}

// ----- Constants -----

type GstFlacEncQuality string

const (
	GstFlacEncQuality0 GstFlacEncQuality = "0" // 0 (0) – 0 - Fastest compression
	GstFlacEncQuality1 GstFlacEncQuality = "1" // 1 (1) – 1
	GstFlacEncQuality2 GstFlacEncQuality = "2" // 2 (2) – 2
	GstFlacEncQuality3 GstFlacEncQuality = "3" // 3 (3) – 3
	GstFlacEncQuality4 GstFlacEncQuality = "4" // 4 (4) – 4
	GstFlacEncQuality5 GstFlacEncQuality = "5" // 5 (5) – 5 - Default
	GstFlacEncQuality6 GstFlacEncQuality = "6" // 6 (6) – 6
	GstFlacEncQuality7 GstFlacEncQuality = "7" // 7 (7) – 7
	GstFlacEncQuality8 GstFlacEncQuality = "8" // 8 (8) – 8 - Highest compression
	GstFlacEncQuality9 GstFlacEncQuality = "9" // 9 (9) – 9 - Insane
)

