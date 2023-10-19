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

type Segmentation struct {
	*base.GstBaseTransform
}

func NewSegmentation() (*Segmentation, error) {
	e, err := gst.NewElement("segmentation")
	if err != nil {
		return nil, err
	}
	return &Segmentation{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSegmentationWithName(name string) (*Segmentation, error) {
	e, err := gst.NewElementWithName("segmentation", name)
	if err != nil {
		return nil, err
	}
	return &Segmentation{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// learning-rate (float32)
//
// Speed with which a motionless foreground pixel would become background (inverse of number of frames)

func (e *Segmentation) GetLearningRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("learning-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Segmentation) SetLearningRate(value float32) error {
	return e.Element.SetProperty("learning-rate", value)
}

// method (GstSegmentationMethod)
//
// Segmentation method to use

func (e *Segmentation) GetMethod() (GstSegmentationMethod, error) {
	var value GstSegmentationMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSegmentationMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSegmentationMethod")
	}
	return value, nil
}

func (e *Segmentation) SetMethod(value GstSegmentationMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// test-mode (bool)
//
// If true, the output RGB is overwritten with the calculated foreground (white color)

func (e *Segmentation) GetTestMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("test-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Segmentation) SetTestMode(value bool) error {
	return e.Element.SetProperty("test-mode", value)
}

// ----- Constants -----

type GstSegmentationMethod string

const (
	GstSegmentationMethodCodebook GstSegmentationMethod = "codebook" // codebook (0) – Codebook-based segmentation (Bradski2008)
	GstSegmentationMethodMog GstSegmentationMethod = "mog" // mog (1) – Mixture-of-Gaussians segmentation (Bowden2001)
	GstSegmentationMethodMog2 GstSegmentationMethod = "mog2" // mog2 (2) – Mixture-of-Gaussians segmentation (Zivkovic2004)
)

