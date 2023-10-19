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

	"github.com/livekit/gstplugins/base"
)

type Vadeinterlace struct {
	*base.GstBaseTransform
}

func NewVadeinterlace() (*Vadeinterlace, error) {
	e, err := gst.NewElement("vadeinterlace")
	if err != nil {
		return nil, err
	}
	return &Vadeinterlace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVadeinterlaceWithName(name string) (*Vadeinterlace, error) {
	e, err := gst.NewElementWithName("vadeinterlace", name)
	if err != nil {
		return nil, err
	}
	return &Vadeinterlace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// method (GstVaDeinterlaceMethods)
//
// Selects the different deinterlacing algorithms that can be used.

// ----- Constants -----

type GstVaDeinterlaceMethods string

const (
	GstVaDeinterlaceMethodsBob GstVaDeinterlaceMethods = "bob" // bob (1) – Bob: Interpolating missing lines by using the adjacent lines.
	GstVaDeinterlaceMethodsAdaptive GstVaDeinterlaceMethods = "adaptive" // adaptive (3) – Adaptive: Interpolating missing lines by using spatial/temporal references.
	GstVaDeinterlaceMethodsCompensated GstVaDeinterlaceMethods = "compensated" // compensated (4) – Compensation: Recreating missing lines by using motion vector.
)

