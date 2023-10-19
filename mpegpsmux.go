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

type Mpegpsmux struct {
	Element *gst.Element
}

func NewMpegpsmux() (*Mpegpsmux, error) {
	e, err := gst.NewElement("mpegpsmux")
	if err != nil {
		return nil, err
	}
	return &Mpegpsmux{Element: e}, nil
}

func NewMpegpsmuxWithName(name string) (*Mpegpsmux, error) {
	e, err := gst.NewElementWithName("mpegpsmux", name)
	if err != nil {
		return nil, err
	}
	return &Mpegpsmux{Element: e}, nil
}

// ----- Properties -----

// aggregate-gops (bool)
//
// Whether to aggregate GOPs and push them out as buffer lists

func (e *Mpegpsmux) GetAggregateGops() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("aggregate-gops")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpegpsmux) SetAggregateGops(value bool) error {
	return e.Element.SetProperty("aggregate-gops", value)
}

