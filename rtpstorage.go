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

type Rtpstorage struct {
	Element *gst.Element
}

func NewRtpstorage() (*Rtpstorage, error) {
	e, err := gst.NewElement("rtpstorage")
	if err != nil {
		return nil, err
	}
	return &Rtpstorage{Element: e}, nil
}

func NewRtpstorageWithName(name string) (*Rtpstorage, error) {
	e, err := gst.NewElementWithName("rtpstorage", name)
	if err != nil {
		return nil, err
	}
	return &Rtpstorage{Element: e}, nil
}

// ----- Properties -----

// internal-storage (GstGobject)
//
// Internal RtpStorage object

func (e *Rtpstorage) GetInternalStorage() (interface{}, error) {
	return e.Element.GetProperty("internal-storage")
}

// size-time (uint64)
//
// The amount of data to keep in the storage (in ns, 0-disable)

func (e *Rtpstorage) GetSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Rtpstorage) SetSizeTime(value uint64) error {
	return e.Element.SetProperty("size-time", value)
}

