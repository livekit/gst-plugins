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

package base

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type GstDataQueue struct {
	Element *gst.Element
}

// ----- Properties -----

// current-level-bytes (uint)

func (e *GstDataQueue) GetCurrentLevelBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// current-level-time (uint64)

func (e *GstDataQueue) GetCurrentLevelTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// current-level-visible (uint)

func (e *GstDataQueue) GetCurrentLevelVisible() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-visible")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

