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

type SwitchbinexampleLaunchLine struct {
	Element *gst.Element
}

func NewSwitchbinexampleLaunchLine() (*SwitchbinexampleLaunchLine, error) {
	e, err := gst.NewElement("switchbinExample launch line")
	if err != nil {
		return nil, err
	}
	return &SwitchbinexampleLaunchLine{Element: e}, nil
}

func NewSwitchbinexampleLaunchLineWithName(name string) (*SwitchbinexampleLaunchLine, error) {
	e, err := gst.NewElementWithName("switchbinExample launch line", name)
	if err != nil {
		return nil, err
	}
	return &SwitchbinexampleLaunchLine{Element: e}, nil
}

// ----- Properties -----

// current-path (uint)
//
// Returns the currently selected path number. If there is no current
// path (due to no caps, or unsupported caps), the value is G_MAXUINT. Read-only.

func (e *SwitchbinexampleLaunchLine) GetCurrentPath() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// num-paths (uint)
//
// Configure how many paths the switchbin will be choosing between. Attempting
// to configure a path outside the range 0..(n-1) will fail. Reducing the
// number of paths will release any paths outside the new range, which might
// trigger activation of a new path by re-assessing the current caps.

func (e *SwitchbinexampleLaunchLine) GetNumPaths() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-paths")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *SwitchbinexampleLaunchLine) SetNumPaths(value uint) error {
	return e.Element.SetProperty("num-paths", value)
}

