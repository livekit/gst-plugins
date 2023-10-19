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

type Line21Encoder struct {
	*base.GstBaseTransform
}

func NewLine21Encoder() (*Line21Encoder, error) {
	e, err := gst.NewElement("line21encoder")
	if err != nil {
		return nil, err
	}
	return &Line21Encoder{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLine21EncoderWithName(name string) (*Line21Encoder, error) {
	e, err := gst.NewElementWithName("line21encoder", name)
	if err != nil {
		return nil, err
	}
	return &Line21Encoder{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// remove-caption-meta (bool)
//
// Selects whether the encoded GstVideoCaptionMeta should be removed from
// the outgoing video buffers or whether it should be kept.

func (e *Line21Encoder) GetRemoveCaptionMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-caption-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Line21Encoder) SetRemoveCaptionMeta(value bool) error {
	return e.Element.SetProperty("remove-caption-meta", value)
}

