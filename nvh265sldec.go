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

type Nvh265Sldec struct {
	Element *gst.Element
}

func NewNvh265Sldec() (*Nvh265Sldec, error) {
	e, err := gst.NewElement("nvh265sldec")
	if err != nil {
		return nil, err
	}
	return &Nvh265Sldec{Element: e}, nil
}

func NewNvh265SldecWithName(name string) (*Nvh265Sldec, error) {
	e, err := gst.NewElementWithName("nvh265sldec", name)
	if err != nil {
		return nil, err
	}
	return &Nvh265Sldec{Element: e}, nil
}

// ----- Properties -----

// cuda-device-id (uint)
//
// Assigned CUDA device id

func (e *Nvh265Sldec) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

