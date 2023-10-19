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

type Openjpegdec struct {
	Element *gst.Element
}

func NewOpenjpegdec() (*Openjpegdec, error) {
	e, err := gst.NewElement("openjpegdec")
	if err != nil {
		return nil, err
	}
	return &Openjpegdec{Element: e}, nil
}

func NewOpenjpegdecWithName(name string) (*Openjpegdec, error) {
	e, err := gst.NewElementWithName("openjpegdec", name)
	if err != nil {
		return nil, err
	}
	return &Openjpegdec{Element: e}, nil
}

// ----- Properties -----

// max-slice-threads (int)
//
// Maximum number of worker threads to spawn. (0 = auto)

func (e *Openjpegdec) GetMaxSliceThreads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-slice-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegdec) SetMaxSliceThreads(value int) error {
	return e.Element.SetProperty("max-slice-threads", value)
}

// max-threads (int)
//
// Maximum number of worker threads to spawn used by openjpeg internally. (0 = no thread)

func (e *Openjpegdec) GetMaxThreads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegdec) SetMaxThreads(value int) error {
	return e.Element.SetProperty("max-threads", value)
}

