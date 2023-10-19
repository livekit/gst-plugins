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

type Vah264Enc struct {
	Element *gst.Element
}

func NewVah264Enc() (*Vah264Enc, error) {
	e, err := gst.NewElement("vah264enc")
	if err != nil {
		return nil, err
	}
	return &Vah264Enc{Element: e}, nil
}

func NewVah264EncWithName(name string) (*Vah264Enc, error) {
	e, err := gst.NewElementWithName("vah264enc", name)
	if err != nil {
		return nil, err
	}
	return &Vah264Enc{Element: e}, nil
}

// ----- Properties -----

// aud (bool)
//
// Insert the AU (Access Unit) delimeter for each frame.

func (e *Vah264Enc) GetAud() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("aud")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vah264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

// b-frames (uint)
//
// Number of B-frames between two reference frames.

func (e *Vah264Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

// b-pyramid (bool)
//
// Enable the b-pyramid reference structure in GOP.

func (e *Vah264Enc) GetBPyramid() (bool, error) {
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

func (e *Vah264Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

// bitrate (uint)
//
// The desired target bitrate, expressed in kbps. Not available in CQP mode.

// cabac (bool)
//
// It enables CABAC entropy coding mode to improve compression ratio,
// but requires main profile at least.

func (e *Vah264Enc) GetCabac() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cabac")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vah264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

// cc-insert (bool)
//
// Closed Caption Insert mode. Only CEA-708 RAW format is supported for now.

func (e *Vah264Enc) GetCcInsert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cc-insert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vah264Enc) SetCcInsert(value bool) error {
	return e.Element.SetProperty("cc-insert", value)
}

// cpb-size (uint)
//
// The desired max CPB size in Kb (0: auto-calculate).

func (e *Vah264Enc) GetCpbSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cpb-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetCpbSize(value uint) error {
	return e.Element.SetProperty("cpb-size", value)
}

// dct8x8 (bool)
//
// Enable adaptive use of 8x8 transforms in I-frames. This improves
// the compression ratio but requires high profile at least.

func (e *Vah264Enc) GetDct8X8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dct8x8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vah264Enc) SetDct8X8(value bool) error {
	return e.Element.SetProperty("dct8x8", value)
}

// i-frames (uint)
//
// Force the number of i-frames insertion within one GOP.

func (e *Vah264Enc) GetIFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("i-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetIFrames(value uint) error {
	return e.Element.SetProperty("i-frames", value)
}

// key-int-max (uint)
//
// The maximal distance between two keyframes.

func (e *Vah264Enc) GetKeyIntMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("key-int-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetKeyIntMax(value uint) error {
	return e.Element.SetProperty("key-int-max", value)
}

// max-qp (uint)
//
// The maximum quantizer value.

func (e *Vah264Enc) GetMaxQp() (uint, error) {
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

func (e *Vah264Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

// mbbrc (GstVaFeature)
//
// Macroblock level bitrate control. Not available in CQP mode.

func (e *Vah264Enc) GetMbbrc() (interface{}, error) {
	return e.Element.GetProperty("mbbrc")
}

func (e *Vah264Enc) SetMbbrc(value interface{}) error {
	return e.Element.SetProperty("mbbrc", value)
}

// min-qp (uint)
//
// The minimum quantizer value.

func (e *Vah264Enc) GetMinQp() (uint, error) {
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

func (e *Vah264Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

// num-slices (uint)
//
// The number of slices per frame.

func (e *Vah264Enc) GetNumSlices() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slices")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetNumSlices(value uint) error {
	return e.Element.SetProperty("num-slices", value)
}

// qpb (uint)
//
// The quantizer value for B frame. Available only in CQP mode.

func (e *Vah264Enc) GetQpb() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qpb")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetQpb(value uint) error {
	return e.Element.SetProperty("qpb", value)
}

// qpi (uint)
//
// The quantizer value for I frame.

// qpp (uint)
//
// The quantizer value for P frame. Available only in CQP mode.

func (e *Vah264Enc) GetQpp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qpp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetQpp(value uint) error {
	return e.Element.SetProperty("qpp", value)
}

// ref-frames (uint)
//
// The number of reference frames.

func (e *Vah264Enc) GetRefFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ref-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Vah264Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

// target-percentage (uint)
//
// The target percentage of the max bitrate, and expressed in uint, equal to
// "target percentage" * 100. Available only when rate-control is VBR.

// target-usage (uint)
//
// The target usage of the encoder.

// trellis (bool)
//
// It enable the trellis quantization method. Trellis is an improved
// quantization algorithm.

func (e *Vah264Enc) GetTrellis() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vah264Enc) SetTrellis(value bool) error {
	return e.Element.SetProperty("trellis", value)
}

