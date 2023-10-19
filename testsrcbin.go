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

type Testsrcbin struct {
	Element *gst.Element
}

func NewTestsrcbin() (*Testsrcbin, error) {
	e, err := gst.NewElement("testsrcbin")
	if err != nil {
		return nil, err
	}
	return &Testsrcbin{Element: e}, nil
}

func NewTestsrcbinWithName(name string) (*Testsrcbin, error) {
	e, err := gst.NewElementWithName("testsrcbin", name)
	if err != nil {
		return nil, err
	}
	return &Testsrcbin{Element: e}, nil
}

// ----- Properties -----

// expose-sources-async (bool)
//
// Whether to expose sources at random time to simulate a source that is
// reading a file and exposing the srcpads later.

func (e *Testsrcbin) GetExposeSourcesAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("expose-sources-async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Testsrcbin) SetExposeSourcesAsync(value bool) error {
	return e.Element.SetProperty("expose-sources-async", value)
}

// stream-types (string)
//
// String describing the stream types to expose, eg. "video+audio".

func (e *Testsrcbin) GetStreamTypes() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("stream-types")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Testsrcbin) SetStreamTypes(value string) error {
	return e.Element.SetProperty("stream-types", value)
}

