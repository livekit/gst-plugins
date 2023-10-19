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

	"github.com/livekit/gstplugins/base"
)

type AvtpcrfcheckexamplePipeline struct {
	*base.GstBaseTransform
}

func NewAvtpcrfcheckexamplePipeline() (*AvtpcrfcheckexamplePipeline, error) {
	e, err := gst.NewElement("avtpcrfcheckExample pipeline")
	if err != nil {
		return nil, err
	}
	return &AvtpcrfcheckexamplePipeline{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAvtpcrfcheckexamplePipelineWithName(name string) (*AvtpcrfcheckexamplePipeline, error) {
	e, err := gst.NewElementWithName("avtpcrfcheckExample pipeline", name)
	if err != nil {
		return nil, err
	}
	return &AvtpcrfcheckexamplePipeline{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// drop-invalid (bool)
//
// Drop the packets which are not within 25%%%% of the sample period of the CRF timestamps

func (e *AvtpcrfcheckexamplePipeline) GetDropInvalid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-invalid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvtpcrfcheckexamplePipeline) SetDropInvalid(value bool) error {
	return e.Element.SetProperty("drop-invalid", value)
}

