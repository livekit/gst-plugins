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

type GstBaseSrc struct {
	Element *gst.Element
}

// ----- Properties -----

// blocksize (uint)

func (e *GstBaseSrc) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *GstBaseSrc) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// do-timestamp (bool)

func (e *GstBaseSrc) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *GstBaseSrc) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

// num-buffers (int)

func (e *GstBaseSrc) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *GstBaseSrc) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// typefind (bool)

func (e *GstBaseSrc) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *GstBaseSrc) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

