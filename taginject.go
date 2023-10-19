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

type Taginject struct {
	*base.GstBaseTransform
}

func NewTaginject() (*Taginject, error) {
	e, err := gst.NewElement("taginject")
	if err != nil {
		return nil, err
	}
	return &Taginject{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewTaginjectWithName(name string) (*Taginject, error) {
	e, err := gst.NewElementWithName("taginject", name)
	if err != nil {
		return nil, err
	}
	return &Taginject{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// scope (GstTagScope)
//
// Scope of tags to inject (stream | global).

func (e *Taginject) GetScope() (interface{}, error) {
	return e.Element.GetProperty("scope")
}

func (e *Taginject) SetScope(value interface{}) error {
	return e.Element.SetProperty("scope", value)
}

// tags (string)
//
// List of tags to inject into the target file

func (e *Taginject) GetTags() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("tags")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Taginject) SetTags(value string) error {
	return e.Element.SetProperty("tags", value)
}

