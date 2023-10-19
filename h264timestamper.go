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

type H264Timestamper struct {
	Element *gst.Element
}

func NewH264Timestamper() (*H264Timestamper, error) {
	e, err := gst.NewElement("h264timestamper")
	if err != nil {
		return nil, err
	}
	return &H264Timestamper{Element: e}, nil
}

func NewH264TimestamperWithName(name string) (*H264Timestamper, error) {
	e, err := gst.NewElementWithName("h264timestamper", name)
	if err != nil {
		return nil, err
	}
	return &H264Timestamper{Element: e}, nil
}

