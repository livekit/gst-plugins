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

type Rtpg726Pay struct {
	Element *gst.Element
}

func NewRtpg726Pay() (*Rtpg726Pay, error) {
	e, err := gst.NewElement("rtpg726pay")
	if err != nil {
		return nil, err
	}
	return &Rtpg726Pay{Element: e}, nil
}

func NewRtpg726PayWithName(name string) (*Rtpg726Pay, error) {
	e, err := gst.NewElementWithName("rtpg726pay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpg726Pay{Element: e}, nil
}

// ----- Properties -----

// force-aal2 (bool)
//
// Force AAL2 encoding for compatibility with bad depayloaders

func (e *Rtpg726Pay) GetForceAal2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aal2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpg726Pay) SetForceAal2(value bool) error {
	return e.Element.SetProperty("force-aal2", value)
}

