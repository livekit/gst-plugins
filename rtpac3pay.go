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

type Rtpac3Pay struct {
	Element *gst.Element
}

func NewRtpac3Pay() (*Rtpac3Pay, error) {
	e, err := gst.NewElement("rtpac3pay")
	if err != nil {
		return nil, err
	}
	return &Rtpac3Pay{Element: e}, nil
}

func NewRtpac3PayWithName(name string) (*Rtpac3Pay, error) {
	e, err := gst.NewElementWithName("rtpac3pay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpac3Pay{Element: e}, nil
}

