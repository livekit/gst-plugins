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

type Ristrtpdeext struct {
	Element *gst.Element
}

func NewRistrtpdeext() (*Ristrtpdeext, error) {
	e, err := gst.NewElement("ristrtpdeext")
	if err != nil {
		return nil, err
	}
	return &Ristrtpdeext{Element: e}, nil
}

func NewRistrtpdeextWithName(name string) (*Ristrtpdeext, error) {
	e, err := gst.NewElementWithName("ristrtpdeext", name)
	if err != nil {
		return nil, err
	}
	return &Ristrtpdeext{Element: e}, nil
}

// ----- Properties -----

// have-ext-seqnum (bool)
//
// Has an extended sequence number extension been seen

func (e *Ristrtpdeext) GetHaveExtSeqnum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("have-ext-seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// max-ext-seqnum (uint)
//
// Largest extended sequence number received

func (e *Ristrtpdeext) GetMaxExtSeqnum() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-ext-seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

