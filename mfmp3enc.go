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

type Mfmp3Enc struct {
	Element *gst.Element
}

func NewMfmp3Enc() (*Mfmp3Enc, error) {
	e, err := gst.NewElement("mfmp3enc")
	if err != nil {
		return nil, err
	}
	return &Mfmp3Enc{Element: e}, nil
}

func NewMfmp3EncWithName(name string) (*Mfmp3Enc, error) {
	e, err := gst.NewElementWithName("mfmp3enc", name)
	if err != nil {
		return nil, err
	}
	return &Mfmp3Enc{Element: e}, nil
}

// ----- Properties -----

// bitrate (uint)
//
// Bitrate in bit/sec, (0 = auto), valid values are { 0, 8000, 16000, 18000, 20000, 24000, 32000, 40000, 48000, 56000, 64000, 80000, 96000, 112000, 128000, 160000, 192000, 224000, 256000, 320000 }

func (e *Mfmp3Enc) GetBitrate() (uint, error) {
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

func (e *Mfmp3Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

