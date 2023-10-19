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

type Oggmux struct {
	Element *gst.Element
}

func NewOggmux() (*Oggmux, error) {
	e, err := gst.NewElement("oggmux")
	if err != nil {
		return nil, err
	}
	return &Oggmux{Element: e}, nil
}

func NewOggmuxWithName(name string) (*Oggmux, error) {
	e, err := gst.NewElementWithName("oggmux", name)
	if err != nil {
		return nil, err
	}
	return &Oggmux{Element: e}, nil
}

// ----- Properties -----

// max-delay (uint64)
//
// Maximum delay in multiplexing streams

func (e *Oggmux) GetMaxDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Oggmux) SetMaxDelay(value uint64) error {
	return e.Element.SetProperty("max-delay", value)
}

// max-page-delay (uint64)
//
// Maximum delay for sending out a page

func (e *Oggmux) GetMaxPageDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-page-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Oggmux) SetMaxPageDelay(value uint64) error {
	return e.Element.SetProperty("max-page-delay", value)
}

// max-tolerance (uint64)
//
// Maximum timestamp difference for maintaining perfect granules

func (e *Oggmux) GetMaxTolerance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Oggmux) SetMaxTolerance(value uint64) error {
	return e.Element.SetProperty("max-tolerance", value)
}

// skeleton (bool)
//
// Whether to include a Skeleton track

func (e *Oggmux) GetSkeleton() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skeleton")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Oggmux) SetSkeleton(value bool) error {
	return e.Element.SetProperty("skeleton", value)
}

