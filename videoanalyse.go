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

type Videoanalyse struct {
	*base.GstBaseTransform
}

func NewVideoanalyse() (*Videoanalyse, error) {
	e, err := gst.NewElement("videoanalyse")
	if err != nil {
		return nil, err
	}
	return &Videoanalyse{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideoanalyseWithName(name string) (*Videoanalyse, error) {
	e, err := gst.NewElementWithName("videoanalyse", name)
	if err != nil {
		return nil, err
	}
	return &Videoanalyse{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// message (bool)
//
// Post statics messages

func (e *Videoanalyse) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videoanalyse) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

