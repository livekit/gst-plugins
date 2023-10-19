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

type Navigationtest struct {
	*base.GstBaseTransform
}

func NewNavigationtest() (*Navigationtest, error) {
	e, err := gst.NewElement("navigationtest")
	if err != nil {
		return nil, err
	}
	return &Navigationtest{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewNavigationtestWithName(name string) (*Navigationtest, error) {
	e, err := gst.NewElementWithName("navigationtest", name)
	if err != nil {
		return nil, err
	}
	return &Navigationtest{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// display-mouse (bool)
//
// Toggles display of mouse events.

func (e *Navigationtest) GetDisplayMouse() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-mouse")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Navigationtest) SetDisplayMouse(value bool) error {
	return e.Element.SetProperty("display-mouse", value)
}

// display-touch (bool)
//
// Toggles display of touch events.

func (e *Navigationtest) GetDisplayTouch() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-touch")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Navigationtest) SetDisplayTouch(value bool) error {
	return e.Element.SetProperty("display-touch", value)
}

