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

type Speed struct {
	Element *gst.Element
}

func NewSpeed() (*Speed, error) {
	e, err := gst.NewElement("speed")
	if err != nil {
		return nil, err
	}
	return &Speed{Element: e}, nil
}

func NewSpeedWithName(name string) (*Speed, error) {
	e, err := gst.NewElementWithName("speed", name)
	if err != nil {
		return nil, err
	}
	return &Speed{Element: e}, nil
}

// ----- Properties -----

// speed (float32)
//
// speed

func (e *Speed) GetSpeed() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Speed) SetSpeed(value float32) error {
	return e.Element.SetProperty("speed", value)
}

