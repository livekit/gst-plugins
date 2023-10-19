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

type Accurip struct {
	*base.GstBaseTransform
}

func NewAccurip() (*Accurip, error) {
	e, err := gst.NewElement("accurip")
	if err != nil {
		return nil, err
	}
	return &Accurip{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAccuripWithName(name string) (*Accurip, error) {
	e, err := gst.NewElementWithName("accurip", name)
	if err != nil {
		return nil, err
	}
	return &Accurip{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// first-track (bool)
//
// Indicate to the CRC calculation algorithm that this is the first track of a set

func (e *Accurip) GetFirstTrack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("first-track")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Accurip) SetFirstTrack(value bool) error {
	return e.Element.SetProperty("first-track", value)
}

// last-track (bool)
//
// Indicate to the CRC calculation algorithm that this is the last track of a set

func (e *Accurip) GetLastTrack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("last-track")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Accurip) SetLastTrack(value bool) error {
	return e.Element.SetProperty("last-track", value)
}

