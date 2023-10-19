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
	"github.com/tinyzimmer/go-gst/gst"
)

type AvdecAdpcmYamaha struct {
	Element *gst.Element
}

func NewAvdecAdpcmYamaha() (*AvdecAdpcmYamaha, error) {
	e, err := gst.NewElement("avdec_adpcm_yamaha")
	if err != nil {
		return nil, err
	}
	return &AvdecAdpcmYamaha{Element: e}, nil
}

func NewAvdecAdpcmYamahaWithName(name string) (*AvdecAdpcmYamaha, error) {
	e, err := gst.NewElementWithName("avdec_adpcm_yamaha", name)
	if err != nil {
		return nil, err
	}
	return &AvdecAdpcmYamaha{Element: e}, nil
}

