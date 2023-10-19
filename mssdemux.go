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

type Mssdemux struct {
	Element *gst.Element
}

func NewMssdemux() (*Mssdemux, error) {
	e, err := gst.NewElement("mssdemux")
	if err != nil {
		return nil, err
	}
	return &Mssdemux{Element: e}, nil
}

func NewMssdemuxWithName(name string) (*Mssdemux, error) {
	e, err := gst.NewElementWithName("mssdemux", name)
	if err != nil {
		return nil, err
	}
	return &Mssdemux{Element: e}, nil
}

// ----- Properties -----

// max-queue-size-buffers (uint)
//
// Maximum buffers that can be stored in each internal stream queue (0 = infinite) (deprecated)

func (e *Mssdemux) GetMaxQueueSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-queue-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mssdemux) SetMaxQueueSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-queue-size-buffers", value)
}

