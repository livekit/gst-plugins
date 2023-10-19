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

	"github.com/livekit/gstplugins/base"
)

type Vapostproc struct {
	*base.GstBaseTransform
}

func NewVapostproc() (*Vapostproc, error) {
	e, err := gst.NewElement("vapostproc")
	if err != nil {
		return nil, err
	}
	return &Vapostproc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVapostprocWithName(name string) (*Vapostproc, error) {
	e, err := gst.NewElementWithName("vapostproc", name)
	if err != nil {
		return nil, err
	}
	return &Vapostproc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// add-borders (bool)
//
// If set to TRUE the filter will add black borders if necessary to
// keep the display aspect ratio.

func (e *Vapostproc) GetAddBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vapostproc) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

// disable-passthrough (bool)
//
// If set to TRUE the filter will not enable passthrough mode, thus
// each frame will be processed. It's useful for cropping, for
// example.

func (e *Vapostproc) GetDisablePassthrough() (bool, error) {
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

func (e *Vapostproc) SetDisablePassthrough(value bool) error {
	return e.Element.SetProperty("disable-passthrough", value)
}

// scale-method (GstVaScaleMethod)
//
// Sets the scale method algorithm to use when resizing.

func (e *Vapostproc) GetScaleMethod() (GstVaScaleMethod, error) {
	var value GstVaScaleMethod
	var ok bool
	v, err := e.Element.GetProperty("scale-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVaScaleMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVaScaleMethod")
	}
	return value, nil
}

func (e *Vapostproc) SetScaleMethod(value GstVaScaleMethod) error {
	e.Element.SetArg("scale-method", string(value))
	return nil
}

// ----- Constants -----

type GstVaScaleMethod string

const (
	GstVaScaleMethodDefault GstVaScaleMethod = "default" // default (0) – Default scaling method
	GstVaScaleMethodFast GstVaScaleMethod = "fast" // fast (256) – Fast scaling method
	GstVaScaleMethodHq GstVaScaleMethod = "hq" // hq (512) – High quality scaling method
)

