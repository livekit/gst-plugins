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

type Jp2Kdecimator struct {
	Element *gst.Element
}

func NewJp2Kdecimator() (*Jp2Kdecimator, error) {
	e, err := gst.NewElement("jp2kdecimator")
	if err != nil {
		return nil, err
	}
	return &Jp2Kdecimator{Element: e}, nil
}

func NewJp2KdecimatorWithName(name string) (*Jp2Kdecimator, error) {
	e, err := gst.NewElementWithName("jp2kdecimator", name)
	if err != nil {
		return nil, err
	}
	return &Jp2Kdecimator{Element: e}, nil
}

// ----- Properties -----

// max-decomposition-levels (int)
//
// Maximum number of decomposition levels to keep (-1 == all)

func (e *Jp2Kdecimator) GetMaxDecompositionLevels() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-decomposition-levels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Jp2Kdecimator) SetMaxDecompositionLevels(value int) error {
	return e.Element.SetProperty("max-decomposition-levels", value)
}

// max-layers (int)
//
// Maximum number of layers to keep (0 == all)

func (e *Jp2Kdecimator) GetMaxLayers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-layers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Jp2Kdecimator) SetMaxLayers(value int) error {
	return e.Element.SetProperty("max-layers", value)
}

