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

type Vp9Enc struct {
	Element *gst.Element
}

func NewVp9Enc() (*Vp9Enc, error) {
	e, err := gst.NewElement("vp9enc")
	if err != nil {
		return nil, err
	}
	return &Vp9Enc{Element: e}, nil
}

func NewVp9EncWithName(name string) (*Vp9Enc, error) {
	e, err := gst.NewElementWithName("vp9enc", name)
	if err != nil {
		return nil, err
	}
	return &Vp9Enc{Element: e}, nil
}

// ----- Properties -----

// aq-mode (GstVpxaq)
//
// Adaptive Quantization Mode

func (e *Vp9Enc) GetAqMode() (GstVpxaq, error) {
	var value GstVpxaq
	var ok bool
	v, err := e.Element.GetProperty("aq-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVpxaq)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVpxaq")
	}
	return value, nil
}

func (e *Vp9Enc) SetAqMode(value GstVpxaq) error {
	e.Element.SetArg("aq-mode", string(value))
	return nil
}

// frame-parallel-decoding (bool)
//
// Whether encoded bitstream should allow parallel processing of video frames in the decoder

func (e *Vp9Enc) GetFrameParallelDecoding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("frame-parallel-decoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vp9Enc) SetFrameParallelDecoding(value bool) error {
	return e.Element.SetProperty("frame-parallel-decoding", value)
}

// row-mt (bool)
//
// Whether each row should be encoded using multiple threads

func (e *Vp9Enc) GetRowMt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("row-mt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vp9Enc) SetRowMt(value bool) error {
	return e.Element.SetProperty("row-mt", value)
}

// tile-columns (int)
//
// Number of tile columns, log2

func (e *Vp9Enc) GetTileColumns() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Vp9Enc) SetTileColumns(value int) error {
	return e.Element.SetProperty("tile-columns", value)
}

// tile-rows (int)
//
// Number of tile rows, log2

func (e *Vp9Enc) GetTileRows() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Vp9Enc) SetTileRows(value int) error {
	return e.Element.SetProperty("tile-rows", value)
}

// ----- Constants -----

type GstVpxaq string

const (
	GstVpxaqOff GstVpxaq = "off" // off (0) – GST_VPX_AQ_OFF
	GstVpxaqVariance GstVpxaq = "variance" // variance (1) – GST_VPX_AQ_VARIANCE
	GstVpxaqComplexity GstVpxaq = "complexity" // complexity (2) – GST_VPX_AQ_COMPLEXITY
	GstVpxaqCyclicRefresh GstVpxaq = "cyclic-refresh" // cyclic-refresh (3) – GST_VPX_AQ_CYCLIC_REFRESH
	GstVpxaqEquator360 GstVpxaq = "equator360" // equator360 (4) – GST_VPX_AQ_EQUATOR360
	GstVpxaqPerceptual GstVpxaq = "perceptual" // perceptual (5) – GST_VPX_AQ_PERCEPTUAL
	GstVpxaqPsnr GstVpxaq = "psnr" // psnr (6) – GST_VPX_AQ_PSNR
	GstVpxaqLookahead GstVpxaq = "lookahead" // lookahead (7) – GST_VPX_AQ_LOOKAHEAD
)

