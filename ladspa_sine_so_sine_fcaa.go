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

type LadspaSineSoSineFcaa struct {
	*base.GstBaseTransform
}

func NewLadspaSineSoSineFcaa() (*LadspaSineSoSineFcaa, error) {
	e, err := gst.NewElement("ladspa-sine-so-sine-fcaa")
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFcaa{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLadspaSineSoSineFcaaWithName(name string) (*LadspaSineSoSineFcaa, error) {
	e, err := gst.NewElementWithName("ladspa-sine-so-sine-fcaa", name)
	if err != nil {
		return nil, err
	}
	return &LadspaSineSoSineFcaa{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// frequency (float32)
//
// Frequency (Hz)

func (e *LadspaSineSoSineFcaa) GetFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LadspaSineSoSineFcaa) SetFrequency(value float32) error {
	return e.Element.SetProperty("frequency", value)
}

