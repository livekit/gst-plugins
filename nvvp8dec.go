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

type Nvvp8Dec struct {
	Element *gst.Element
}

func NewNvvp8Dec() (*Nvvp8Dec, error) {
	e, err := gst.NewElement("nvvp8dec")
	if err != nil {
		return nil, err
	}
	return &Nvvp8Dec{Element: e}, nil
}

func NewNvvp8DecWithName(name string) (*Nvvp8Dec, error) {
	e, err := gst.NewElementWithName("nvvp8dec", name)
	if err != nil {
		return nil, err
	}
	return &Nvvp8Dec{Element: e}, nil
}

