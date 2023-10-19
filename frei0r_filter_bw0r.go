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

type Frei0RFilterBw0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterBw0R() (*Frei0RFilterBw0R, error) {
	e, err := gst.NewElement("frei0r-filter-bw0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterBw0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterBw0RWithName(name string) (*Frei0RFilterBw0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-bw0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterBw0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

