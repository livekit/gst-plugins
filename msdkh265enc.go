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

type Msdkh265Enc struct {
	Element *gst.Element
}

func NewMsdkh265Enc() (*Msdkh265Enc, error) {
	e, err := gst.NewElement("msdkh265enc")
	if err != nil {
		return nil, err
	}
	return &Msdkh265Enc{Element: e}, nil
}

func NewMsdkh265EncWithName(name string) (*Msdkh265Enc, error) {
	e, err := gst.NewElementWithName("msdkh265enc", name)
	if err != nil {
		return nil, err
	}
	return &Msdkh265Enc{Element: e}, nil
}

// ----- Properties -----

// b-pyramid (bool)
//
// Enable B-Pyramid Reference structure

func (e *Msdkh265Enc) GetBPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

// dblk-idc (uint)
//
// Option of disable deblocking idc

func (e *Msdkh265Enc) GetDblkIdc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dblk-idc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetDblkIdc(value uint) error {
	return e.Element.SetProperty("dblk-idc", value)
}

// intra-refresh-type (GstMsdkEncIntraRefreshType)
//
// Set intra refresh type

func (e *Msdkh265Enc) GetIntraRefreshType() (GstMsdkEncIntraRefreshType, error) {
	var value GstMsdkEncIntraRefreshType
	var ok bool
	v, err := e.Element.GetProperty("intra-refresh-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncIntraRefreshType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncIntraRefreshType")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetIntraRefreshType(value GstMsdkEncIntraRefreshType) error {
	e.Element.SetArg("intra-refresh-type", string(value))
	return nil
}

// low-power (bool)
//
// Enable low power mode (DEPRECATED, use tune instead)

func (e *Msdkh265Enc) GetLowPower() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-power")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetLowPower(value bool) error {
	return e.Element.SetProperty("low-power", value)
}

// max-qp (uint)
//
// Maximum quantizer scale for I/P/B frames

func (e *Msdkh265Enc) GetMaxQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

// max-qp-b (uint)

func (e *Msdkh265Enc) GetMaxQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

// max-qp-i (uint)

func (e *Msdkh265Enc) GetMaxQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

// max-qp-p (uint)

func (e *Msdkh265Enc) GetMaxQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

// max-slice-size (uint)
//
// Maximum slice size in bytes (if enabled MSDK will ignore the control over num-slices)

func (e *Msdkh265Enc) GetMaxSliceSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-slice-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

// min-qp (uint)
//
// Minimal quantizer scale for I/P/B frames

func (e *Msdkh265Enc) GetMinQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

// min-qp-b (uint)

func (e *Msdkh265Enc) GetMinQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

// min-qp-i (uint)

func (e *Msdkh265Enc) GetMinQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

// min-qp-p (uint)

func (e *Msdkh265Enc) GetMinQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

// num-tile-cols (uint)
//
// number of columns for tiled encoding

func (e *Msdkh265Enc) GetNumTileCols() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-cols")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetNumTileCols(value uint) error {
	return e.Element.SetProperty("num-tile-cols", value)
}

// num-tile-rows (uint)
//
// number of rows for tiled encoding

func (e *Msdkh265Enc) GetNumTileRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetNumTileRows(value uint) error {
	return e.Element.SetProperty("num-tile-rows", value)
}

// p-pyramid (bool)
//
// Enable P-Pyramid Reference structure

func (e *Msdkh265Enc) GetPPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("p-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetPPyramid(value bool) error {
	return e.Element.SetProperty("p-pyramid", value)
}

// pic-timing-sei (bool)

func (e *Msdkh265Enc) GetPicTimingSei() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pic-timing-sei")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetPicTimingSei(value bool) error {
	return e.Element.SetProperty("pic-timing-sei", value)
}

// transform-skip (GstMsdkEncTransformSkip)
//
// Transform Skip option

func (e *Msdkh265Enc) GetTransformSkip() (GstMsdkEncTransformSkip, error) {
	var value GstMsdkEncTransformSkip
	var ok bool
	v, err := e.Element.GetProperty("transform-skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncTransformSkip)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncTransformSkip")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetTransformSkip(value GstMsdkEncTransformSkip) error {
	e.Element.SetArg("transform-skip", string(value))
	return nil
}

// tune (GstMsdkEncTuneMode)
//
// Encoder tuning option

func (e *Msdkh265Enc) GetTune() (GstMsdkEncTuneMode, error) {
	var value GstMsdkEncTuneMode
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncTuneMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncTuneMode")
	}
	return value, nil
}

func (e *Msdkh265Enc) SetTune(value GstMsdkEncTuneMode) error {
	e.Element.SetArg("tune", string(value))
	return nil
}

// ----- Constants -----

type GstMsdkEncTransformSkip string

const (
	GstMsdkEncTransformSkipAuto GstMsdkEncTransformSkip = "auto" // auto (0) – SDK desides what to do
	GstMsdkEncTransformSkipOff GstMsdkEncTransformSkip = "off" // off (32) – transform_skip_enabled_flag will be set to 0 in PPS
	GstMsdkEncTransformSkipOn GstMsdkEncTransformSkip = "on" // on (16) – transform_skip_enabled_flag will be set to 1 in PPS
)

type GstMsdkEncTuneMode string

const (
	GstMsdkEncTuneModeAuto GstMsdkEncTuneMode = "auto" // auto (0) – Auto
	GstMsdkEncTuneModeNone GstMsdkEncTuneMode = "none" // none (32) – None
	GstMsdkEncTuneModeLowPower GstMsdkEncTuneMode = "low-power" // low-power (16) – Low power mode
)

type GstMsdkEncIntraRefreshType string

const (
	GstMsdkEncIntraRefreshTypeNo GstMsdkEncIntraRefreshType = "no" // no (0) – No (default)
	GstMsdkEncIntraRefreshTypeVertical GstMsdkEncIntraRefreshType = "vertical" // vertical (1) – Vertical
	GstMsdkEncIntraRefreshTypeHorizontal GstMsdkEncIntraRefreshType = "horizontal" // horizontal (2) – Horizontal
	GstMsdkEncIntraRefreshTypeSlice GstMsdkEncIntraRefreshType = "slice" // slice (3) – Slice
)

