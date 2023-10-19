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

type Voamrwbenc struct {
	Element *gst.Element
}

func NewVoamrwbenc() (*Voamrwbenc, error) {
	e, err := gst.NewElement("voamrwbenc")
	if err != nil {
		return nil, err
	}
	return &Voamrwbenc{Element: e}, nil
}

func NewVoamrwbencWithName(name string) (*Voamrwbenc, error) {
	e, err := gst.NewElementWithName("voamrwbenc", name)
	if err != nil {
		return nil, err
	}
	return &Voamrwbenc{Element: e}, nil
}

// ----- Properties -----

// band-mode (GstVoAmrWbEncBandMode)
//
// Encoding Band Mode (Kbps)

func (e *Voamrwbenc) GetBandMode() (GstVoAmrWbEncBandMode, error) {
	var value GstVoAmrWbEncBandMode
	var ok bool
	v, err := e.Element.GetProperty("band-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVoAmrWbEncBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVoAmrWbEncBandMode")
	}
	return value, nil
}

func (e *Voamrwbenc) SetBandMode(value GstVoAmrWbEncBandMode) error {
	e.Element.SetArg("band-mode", string(value))
	return nil
}

// ----- Constants -----

type GstVoAmrWbEncBandMode string

const (
	GstVoAmrWbEncBandModeMr660 GstVoAmrWbEncBandMode = "MR660" // MR660 (0) – MR660
	GstVoAmrWbEncBandModeMr885 GstVoAmrWbEncBandMode = "MR885" // MR885 (1) – MR885
	GstVoAmrWbEncBandModeMr1265 GstVoAmrWbEncBandMode = "MR1265" // MR1265 (2) – MR1265
	GstVoAmrWbEncBandModeMr1425 GstVoAmrWbEncBandMode = "MR1425" // MR1425 (2) – MR1425
	GstVoAmrWbEncBandModeMr1585 GstVoAmrWbEncBandMode = "MR1585" // MR1585 (3) – MR1585
	GstVoAmrWbEncBandModeMr1825 GstVoAmrWbEncBandMode = "MR1825" // MR1825 (4) – MR1825
	GstVoAmrWbEncBandModeMr1985 GstVoAmrWbEncBandMode = "MR1985" // MR1985 (5) – MR1985
	GstVoAmrWbEncBandModeMr2305 GstVoAmrWbEncBandMode = "MR2305" // MR2305 (6) – MR2305
	GstVoAmrWbEncBandModeMr2385 GstVoAmrWbEncBandMode = "MR2385" // MR2385 (7) – MR2385
	GstVoAmrWbEncBandModeMrdtx GstVoAmrWbEncBandMode = "MRDTX" // MRDTX (8) – MRDTX
)

