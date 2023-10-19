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

type Nvh265Enc struct {
	Element *gst.Element
}

func NewNvh265Enc() (*Nvh265Enc, error) {
	e, err := gst.NewElement("nvh265enc")
	if err != nil {
		return nil, err
	}
	return &Nvh265Enc{Element: e}, nil
}

func NewNvh265EncWithName(name string) (*Nvh265Enc, error) {
	e, err := gst.NewElementWithName("nvh265enc", name)
	if err != nil {
		return nil, err
	}
	return &Nvh265Enc{Element: e}, nil
}

// ----- Properties -----

// aud (bool)
//
// Use AU (Access Unit) delimiter

func (e *Nvh265Enc) GetAud() (bool, error) {
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

func (e *Nvh265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

// b-adapt (bool)
//
// Enable adaptive B-frame insert when lookahead is enabled

func (e *Nvh265Enc) GetBAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Nvh265Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

// bframes (uint)
//
// Number of B-frames between I and P

func (e *Nvh265Enc) GetBframes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bframes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Nvh265Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

// rc-lookahead (uint)
//
// Number of frames for frame type lookahead

func (e *Nvh265Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Nvh265Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

// temporal-aq (bool)
//
// Temporal Adaptive Quantization

func (e *Nvh265Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Nvh265Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

// vbv-buffer-size (uint)
//
// VBV(HRD) Buffer Size in kbits (0 = NVENC default)

func (e *Nvh265Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Nvh265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

// weighted-pred (bool)
//
// Weighted Prediction

func (e *Nvh265Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Nvh265Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

