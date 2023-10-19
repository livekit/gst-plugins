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

type Amrnbdec struct {
	Element *gst.Element
}

func NewAmrnbdec() (*Amrnbdec, error) {
	e, err := gst.NewElement("amrnbdec")
	if err != nil {
		return nil, err
	}
	return &Amrnbdec{Element: e}, nil
}

func NewAmrnbdecWithName(name string) (*Amrnbdec, error) {
	e, err := gst.NewElementWithName("amrnbdec", name)
	if err != nil {
		return nil, err
	}
	return &Amrnbdec{Element: e}, nil
}

// ----- Properties -----

// variant (GstAmrnbVariant)
//
// The decoder variant

func (e *Amrnbdec) GetVariant() (GstAmrnbVariant, error) {
	var value GstAmrnbVariant
	var ok bool
	v, err := e.Element.GetProperty("variant")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmrnbVariant)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmrnbVariant")
	}
	return value, nil
}

func (e *Amrnbdec) SetVariant(value GstAmrnbVariant) error {
	e.Element.SetArg("variant", string(value))
	return nil
}

// ----- Constants -----

type GstAmrnbVariant string

const (
	GstAmrnbVariantIf1 GstAmrnbVariant = "IF1" // IF1 (0) – IF1
	GstAmrnbVariantIf2 GstAmrnbVariant = "IF2" // IF2 (1) – IF2
)

