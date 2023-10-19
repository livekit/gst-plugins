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

type Ccextractor struct {
	Element *gst.Element
}

func NewCcextractor() (*Ccextractor, error) {
	e, err := gst.NewElement("ccextractor")
	if err != nil {
		return nil, err
	}
	return &Ccextractor{Element: e}, nil
}

func NewCcextractorWithName(name string) (*Ccextractor, error) {
	e, err := gst.NewElementWithName("ccextractor", name)
	if err != nil {
		return nil, err
	}
	return &Ccextractor{Element: e}, nil
}

// ----- Properties -----

// remove-caption-meta (bool)
//
// Selects whether the GstVideoCaptionMeta should be removed from the
// outgoing video buffers or whether it should be kept.

func (e *Ccextractor) GetRemoveCaptionMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-caption-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ccextractor) SetRemoveCaptionMeta(value bool) error {
	return e.Element.SetProperty("remove-caption-meta", value)
}

