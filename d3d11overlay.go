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

type D3D11Overlay struct {
	*base.GstBaseTransform
}

func NewD3D11Overlay() (*D3D11Overlay, error) {
	e, err := gst.NewElement("d3d11overlay")
	if err != nil {
		return nil, err
	}
	return &D3D11Overlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewD3D11OverlayWithName(name string) (*D3D11Overlay, error) {
	e, err := gst.NewElementWithName("d3d11overlay", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Overlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}
