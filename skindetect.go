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

type Skindetect struct {
	*base.GstBaseTransform
}

func NewSkindetect() (*Skindetect, error) {
	e, err := gst.NewElement("skindetect")
	if err != nil {
		return nil, err
	}
	return &Skindetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSkindetectWithName(name string) (*Skindetect, error) {
	e, err := gst.NewElementWithName("skindetect", name)
	if err != nil {
		return nil, err
	}
	return &Skindetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// method (GstSkindetectMethod)
//
// Method to use

func (e *Skindetect) GetMethod() (GstSkindetectMethod, error) {
	var value GstSkindetectMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSkindetectMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSkindetectMethod")
	}
	return value, nil
}

func (e *Skindetect) SetMethod(value GstSkindetectMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// postprocess (bool)
//
// Apply opening-closing to skin detection to extract large, significant blobs

func (e *Skindetect) GetPostprocess() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("postprocess")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Skindetect) SetPostprocess(value bool) error {
	return e.Element.SetProperty("postprocess", value)
}

// ----- Constants -----

type GstSkindetectMethod string

const (
	GstSkindetectMethodHsv GstSkindetectMethod = "hsv" // hsv (0) – Classic HSV thresholding
	GstSkindetectMethodRgb GstSkindetectMethod = "rgb" // rgb (1) – Normalised-RGB colorspace thresholding
)

