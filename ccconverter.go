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

type Ccconverter struct {
	*base.GstBaseTransform
}

func NewCcconverter() (*Ccconverter, error) {
	e, err := gst.NewElement("ccconverter")
	if err != nil {
		return nil, err
	}
	return &Ccconverter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCcconverterWithName(name string) (*Ccconverter, error) {
	e, err := gst.NewElementWithName("ccconverter", name)
	if err != nil {
		return nil, err
	}
	return &Ccconverter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// cdp-mode (GstCcconverterCdpmode)
//
// Only insert the selection sections into CEA 708 CDP packets.

// ----- Constants -----

type GstCcconverterCdpmode string

const (
	GstCcconverterCdpmodeTimeCode GstCcconverterCdpmode = "time-code" // time-code (0x00000001) – Store time code information in CDP packets
	GstCcconverterCdpmodeCcData GstCcconverterCdpmode = "cc-data" // cc-data (0x00000002) – Store CC data in CDP packets
	GstCcconverterCdpmodeCcSvcInfo GstCcconverterCdpmode = "cc-svc-info" // cc-svc-info (0x00000004) – Store CC service information in CDP packets
)

