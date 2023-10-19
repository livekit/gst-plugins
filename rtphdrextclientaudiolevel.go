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

type Rtphdrextclientaudiolevel struct {
	Element *gst.Element
}

func NewRtphdrextclientaudiolevel() (*Rtphdrextclientaudiolevel, error) {
	e, err := gst.NewElement("rtphdrextclientaudiolevel")
	if err != nil {
		return nil, err
	}
	return &Rtphdrextclientaudiolevel{Element: e}, nil
}

func NewRtphdrextclientaudiolevelWithName(name string) (*Rtphdrextclientaudiolevel, error) {
	e, err := gst.NewElementWithName("rtphdrextclientaudiolevel", name)
	if err != nil {
		return nil, err
	}
	return &Rtphdrextclientaudiolevel{Element: e}, nil
}

// ----- Properties -----

// vad (bool)
//
// If the vad extension attribute is enabled or not, default to FALSE.

func (e *Rtphdrextclientaudiolevel) GetVad() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vad")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

