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

type Libde265Dec struct {
	Element *gst.Element
}

func NewLibde265Dec() (*Libde265Dec, error) {
	e, err := gst.NewElement("libde265dec")
	if err != nil {
		return nil, err
	}
	return &Libde265Dec{Element: e}, nil
}

func NewLibde265DecWithName(name string) (*Libde265Dec, error) {
	e, err := gst.NewElementWithName("libde265dec", name)
	if err != nil {
		return nil, err
	}
	return &Libde265Dec{Element: e}, nil
}

// ----- Properties -----

// max-threads (int)
//
// Maximum number of worker threads to spawn. (0 = auto)

func (e *Libde265Dec) GetMaxThreads() (int, error) {
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

func (e *Libde265Dec) SetMaxThreads(value int) error {
	return e.Element.SetProperty("max-threads", value)
}

