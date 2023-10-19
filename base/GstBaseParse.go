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

type GstBaseParse struct {
	Element *gst.Element
}

// ----- Properties -----

// disable-passthrough (bool)
//
// If set to TRUE, baseparse will unconditionally force parsing of the
// incoming data. This can be required in the rare cases where the incoming
// side-data (caps, pts, dts, ...) is not trusted by the user and wants to
// force validation and parsing of the incoming data.
// If set to FALSE, decision of whether to parse the data or not is up to
// the implementation (standard behaviour).

func (e *GstBaseParse) GetDisablePassthrough() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-passthrough")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *GstBaseParse) SetDisablePassthrough(value bool) error {
	return e.Element.SetProperty("disable-passthrough", value)
}

