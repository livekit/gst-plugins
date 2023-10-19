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
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type AvtpaafpayexamplePipeline struct {
	Element *gst.Element
}

func NewAvtpaafpayexamplePipeline() (*AvtpaafpayexamplePipeline, error) {
	e, err := gst.NewElement("avtpaafpayExample pipeline")
	if err != nil {
		return nil, err
	}
	return &AvtpaafpayexamplePipeline{Element: e}, nil
}

func NewAvtpaafpayexamplePipelineWithName(name string) (*AvtpaafpayexamplePipeline, error) {
	e, err := gst.NewElementWithName("avtpaafpayExample pipeline", name)
	if err != nil {
		return nil, err
	}
	return &AvtpaafpayexamplePipeline{Element: e}, nil
}

// ----- Properties -----

// timestamp-mode (GstAvtpAafTimestampMode)
//
// AAF timestamping mode

func (e *AvtpaafpayexamplePipeline) GetTimestampMode() (GstAvtpAafTimestampMode, error) {
	var value GstAvtpAafTimestampMode
	var ok bool
	v, err := e.Element.GetProperty("timestamp-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAvtpAafTimestampMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAvtpAafTimestampMode")
	}
	return value, nil
}

func (e *AvtpaafpayexamplePipeline) SetTimestampMode(value GstAvtpAafTimestampMode) error {
	e.Element.SetArg("timestamp-mode", string(value))
	return nil
}

// ----- Constants -----

type GstAvtpAafTimestampMode string

const (
	GstAvtpAafTimestampModeNormal GstAvtpAafTimestampMode = "normal" // normal (0) – Normal timestamping mode
	GstAvtpAafTimestampModeSparse GstAvtpAafTimestampMode = "sparse" // sparse (1) – Sparse timestamping mode
)

