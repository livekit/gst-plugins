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

type AvdecDsdLsbfPlanar struct {
	Element *gst.Element
}

func NewAvdecDsdLsbfPlanar() (*AvdecDsdLsbfPlanar, error) {
	e, err := gst.NewElement("avdec_dsd_lsbf_planar")
	if err != nil {
		return nil, err
	}
	return &AvdecDsdLsbfPlanar{Element: e}, nil
}

func NewAvdecDsdLsbfPlanarWithName(name string) (*AvdecDsdLsbfPlanar, error) {
	e, err := gst.NewElementWithName("avdec_dsd_lsbf_planar", name)
	if err != nil {
		return nil, err
	}
	return &AvdecDsdLsbfPlanar{Element: e}, nil
}

