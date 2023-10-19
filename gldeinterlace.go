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

	"github.com/livekit/gstplugins/base"
)

type Gldeinterlace struct {
	*base.GstBaseTransform
}

func NewGldeinterlace() (*Gldeinterlace, error) {
	e, err := gst.NewElement("gldeinterlace")
	if err != nil {
		return nil, err
	}
	return &Gldeinterlace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGldeinterlaceWithName(name string) (*Gldeinterlace, error) {
	e, err := gst.NewElementWithName("gldeinterlace", name)
	if err != nil {
		return nil, err
	}
	return &Gldeinterlace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// method (GstGldeinterlaceMethod)
//
// Select which deinterlace method apply to GL video texture

func (e *Gldeinterlace) GetMethod() (GstGldeinterlaceMethod, error) {
	var value GstGldeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGldeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGldeinterlaceMethod")
	}
	return value, nil
}

func (e *Gldeinterlace) SetMethod(value GstGldeinterlaceMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// ----- Constants -----

type GstGldeinterlaceMethod string

const (
	GstGldeinterlaceMethodVfir GstGldeinterlaceMethod = "vfir" // vfir (0) – Blur Vertical
	GstGldeinterlaceMethodGreedyh GstGldeinterlaceMethod = "greedyh" // greedyh (1) – Motion Adaptive: Advanced Detection
)

