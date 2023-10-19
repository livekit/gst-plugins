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

type Speexdec struct {
	Element *gst.Element
}

func NewSpeexdec() (*Speexdec, error) {
	e, err := gst.NewElement("speexdec")
	if err != nil {
		return nil, err
	}
	return &Speexdec{Element: e}, nil
}

func NewSpeexdecWithName(name string) (*Speexdec, error) {
	e, err := gst.NewElementWithName("speexdec", name)
	if err != nil {
		return nil, err
	}
	return &Speexdec{Element: e}, nil
}

// ----- Properties -----

// enh (bool)
//
// Enable perceptual enhancement

func (e *Speexdec) GetEnh() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enh")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Speexdec) SetEnh(value bool) error {
	return e.Element.SetProperty("enh", value)
}

