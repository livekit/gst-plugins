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

type Capssetter struct {
	*base.GstBaseTransform
}

func NewCapssetter() (*Capssetter, error) {
	e, err := gst.NewElement("capssetter")
	if err != nil {
		return nil, err
	}
	return &Capssetter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCapssetterWithName(name string) (*Capssetter, error) {
	e, err := gst.NewElementWithName("capssetter", name)
	if err != nil {
		return nil, err
	}
	return &Capssetter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// Merge these caps (thereby overwriting) in the stream

func (e *Capssetter) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Capssetter) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// join (bool)
//
// Match incoming caps' mime-type to mime-type of provided caps

func (e *Capssetter) GetJoin() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("join")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Capssetter) SetJoin(value bool) error {
	return e.Element.SetProperty("join", value)
}

// replace (bool)
//
// Drop fields of incoming caps

func (e *Capssetter) GetReplace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("replace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Capssetter) SetReplace(value bool) error {
	return e.Element.SetProperty("replace", value)
}

