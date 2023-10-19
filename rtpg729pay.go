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

type Rtpg729Pay struct {
	Element *gst.Element
}

func NewRtpg729Pay() (*Rtpg729Pay, error) {
	e, err := gst.NewElement("rtpg729pay")
	if err != nil {
		return nil, err
	}
	return &Rtpg729Pay{Element: e}, nil
}

func NewRtpg729PayWithName(name string) (*Rtpg729Pay, error) {
	e, err := gst.NewElementWithName("rtpg729pay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpg729Pay{Element: e}, nil
}
