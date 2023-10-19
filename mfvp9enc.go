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

type Mfvp9Enc struct {
	Element *gst.Element
}

func NewMfvp9Enc() (*Mfvp9Enc, error) {
	e, err := gst.NewElement("mfvp9enc")
	if err != nil {
		return nil, err
	}
	return &Mfvp9Enc{Element: e}, nil
}

func NewMfvp9EncWithName(name string) (*Mfvp9Enc, error) {
	e, err := gst.NewElementWithName("mfvp9enc", name)
	if err != nil {
		return nil, err
	}
	return &Mfvp9Enc{Element: e}, nil
}

// ----- Properties -----

// adapter-luid (int64)
//
// DXGI Adapter LUID for this elemenet

func (e *Mfvp9Enc) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// bitrate (uint)
//
// Bitrate in kbit/sec

func (e *Mfvp9Enc) GetBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfvp9Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// d3d11-aware (bool)
//
// Whether element supports Direct3D11 texture as an input or not

func (e *Mfvp9Enc) GetD3D11Aware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("d3d11-aware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// gop-size (int)

func (e *Mfvp9Enc) GetGopSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gop-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mfvp9Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// low-latency (bool)

func (e *Mfvp9Enc) GetLowLatency() (bool, error) {
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

func (e *Mfvp9Enc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// max-bitrate (uint)

func (e *Mfvp9Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfvp9Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// rc-mode (GstMfvp9EncRcmode)

func (e *Mfvp9Enc) GetRcMode() (GstMfvp9EncRcmode, error) {
	var value GstMfvp9EncRcmode
	var ok bool
	v, err := e.Element.GetProperty("rc-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMfvp9EncRcmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMfvp9EncRcmode")
	}
	return value, nil
}

func (e *Mfvp9Enc) SetRcMode(value GstMfvp9EncRcmode) error {
	e.Element.SetArg("rc-mode", string(value))
	return nil
}

// ----- Constants -----

type GstMfvp9EncRcmode string

const (
	GstMfvp9EncRcmodeCbr GstMfvp9EncRcmode = "cbr" // cbr (0) – Constant bitrate
	GstMfvp9EncRcmodeQvbr GstMfvp9EncRcmode = "qvbr" // qvbr (1) – Quality-based variable bitrate
)

