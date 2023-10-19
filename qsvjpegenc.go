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

type Qsvjpegenc struct {
	Element *gst.Element
}

func NewQsvjpegenc() (*Qsvjpegenc, error) {
	e, err := gst.NewElement("qsvjpegenc")
	if err != nil {
		return nil, err
	}
	return &Qsvjpegenc{Element: e}, nil
}

func NewQsvjpegencWithName(name string) (*Qsvjpegenc, error) {
	e, err := gst.NewElementWithName("qsvjpegenc", name)
	if err != nil {
		return nil, err
	}
	return &Qsvjpegenc{Element: e}, nil
}

// ----- Properties -----

// quality (uint)
//
// Encoding quality, 100 for best quality

func (e *Qsvjpegenc) GetQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvjpegenc) SetQuality(value uint) error {
	return e.Element.SetProperty("quality", value)
}

