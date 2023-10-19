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

type DshowvdecMpeg4 struct {
	Element *gst.Element
}

func NewDshowvdecMpeg4() (*DshowvdecMpeg4, error) {
	e, err := gst.NewElement("dshowvdec_mpeg4")
	if err != nil {
		return nil, err
	}
	return &DshowvdecMpeg4{Element: e}, nil
}

func NewDshowvdecMpeg4WithName(name string) (*DshowvdecMpeg4, error) {
	e, err := gst.NewElementWithName("dshowvdec_mpeg4", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecMpeg4{Element: e}, nil
}

