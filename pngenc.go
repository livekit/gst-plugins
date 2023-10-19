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

type Pngenc struct {
	Element *gst.Element
}

func NewPngenc() (*Pngenc, error) {
	e, err := gst.NewElement("pngenc")
	if err != nil {
		return nil, err
	}
	return &Pngenc{Element: e}, nil
}

func NewPngencWithName(name string) (*Pngenc, error) {
	e, err := gst.NewElementWithName("pngenc", name)
	if err != nil {
		return nil, err
	}
	return &Pngenc{Element: e}, nil
}

// ----- Properties -----

// compression-level (uint)
//
// PNG compression level

func (e *Pngenc) GetCompressionLevel() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("compression-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Pngenc) SetCompressionLevel(value uint) error {
	return e.Element.SetProperty("compression-level", value)
}

// snapshot (bool)
//
// Send EOS after encoding a frame, useful for snapshots

func (e *Pngenc) GetSnapshot() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("snapshot")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Pngenc) SetSnapshot(value bool) error {
	return e.Element.SetProperty("snapshot", value)
}

