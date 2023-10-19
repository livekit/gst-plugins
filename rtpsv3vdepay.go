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

type Rtpsv3Vdepay struct {
	Element *gst.Element
}

func NewRtpsv3Vdepay() (*Rtpsv3Vdepay, error) {
	e, err := gst.NewElement("rtpsv3vdepay")
	if err != nil {
		return nil, err
	}
	return &Rtpsv3Vdepay{Element: e}, nil
}

func NewRtpsv3VdepayWithName(name string) (*Rtpsv3Vdepay, error) {
	e, err := gst.NewElementWithName("rtpsv3vdepay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpsv3Vdepay{Element: e}, nil
}

