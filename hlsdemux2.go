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

type Hlsdemux2 struct {
	Element *gst.Element
}

func NewHlsdemux2() (*Hlsdemux2, error) {
	e, err := gst.NewElement("hlsdemux2")
	if err != nil {
		return nil, err
	}
	return &Hlsdemux2{Element: e}, nil
}

func NewHlsdemux2WithName(name string) (*Hlsdemux2, error) {
	e, err := gst.NewElementWithName("hlsdemux2", name)
	if err != nil {
		return nil, err
	}
	return &Hlsdemux2{Element: e}, nil
}

// ----- Properties -----

// start-bitrate (uint)
//
// Initial bitrate to use to choose first alternate (0 = automatic) (bits/s)

func (e *Hlsdemux2) GetStartBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("start-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Hlsdemux2) SetStartBitrate(value uint) error {
	return e.Element.SetProperty("start-bitrate", value)
}

