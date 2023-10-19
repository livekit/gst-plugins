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

type Rtpg726Depay struct {
	Element *gst.Element
}

func NewRtpg726Depay() (*Rtpg726Depay, error) {
	e, err := gst.NewElement("rtpg726depay")
	if err != nil {
		return nil, err
	}
	return &Rtpg726Depay{Element: e}, nil
}

func NewRtpg726DepayWithName(name string) (*Rtpg726Depay, error) {
	e, err := gst.NewElementWithName("rtpg726depay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpg726Depay{Element: e}, nil
}

// ----- Properties -----

// force-aal2 (bool)
//
// Force AAL2 decoding for compatibility with bad payloaders

func (e *Rtpg726Depay) GetForceAal2() (bool, error) {
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

func (e *Rtpg726Depay) SetForceAal2(value bool) error {
	return e.Element.SetProperty("force-aal2", value)
}

