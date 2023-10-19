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

type Audiorate struct {
	Element *gst.Element
}

func NewAudiorate() (*Audiorate, error) {
	e, err := gst.NewElement("audiorate")
	if err != nil {
		return nil, err
	}
	return &Audiorate{Element: e}, nil
}

func NewAudiorateWithName(name string) (*Audiorate, error) {
	e, err := gst.NewElementWithName("audiorate", name)
	if err != nil {
		return nil, err
	}
	return &Audiorate{Element: e}, nil
}

// ----- Properties -----

// add (uint64)
//
// Number of added samples

func (e *Audiorate) GetAdd() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("add")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// drop (uint64)
//
// Number of dropped samples

func (e *Audiorate) GetDrop() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// in (uint64)
//
// Number of input samples

func (e *Audiorate) GetIn() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("in")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// out (uint64)
//
// Number of output samples

func (e *Audiorate) GetOut() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("out")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// silent (bool)
//
// Don't emit notify for dropped and duplicated frames

func (e *Audiorate) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audiorate) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// skip-to-first (bool)
//
// Don't produce buffers before the first one we receive.

func (e *Audiorate) GetSkipToFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-to-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audiorate) SetSkipToFirst(value bool) error {
	return e.Element.SetProperty("skip-to-first", value)
}

// tolerance (uint64)
//
// The difference between incoming timestamp and next timestamp must exceed
// the given value for audiorate to add or drop samples.

func (e *Audiorate) GetTolerance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audiorate) SetTolerance(value uint64) error {
	return e.Element.SetProperty("tolerance", value)
}

