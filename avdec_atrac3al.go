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

type AvdecAtrac3Al struct {
	Element *gst.Element
}

func NewAvdecAtrac3Al() (*AvdecAtrac3Al, error) {
	e, err := gst.NewElement("avdec_atrac3al")
	if err != nil {
		return nil, err
	}
	return &AvdecAtrac3Al{Element: e}, nil
}

func NewAvdecAtrac3AlWithName(name string) (*AvdecAtrac3Al, error) {
	e, err := gst.NewElementWithName("avdec_atrac3al", name)
	if err != nil {
		return nil, err
	}
	return &AvdecAtrac3Al{Element: e}, nil
}

