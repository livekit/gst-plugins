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

type Audiodynamic struct {
	*base.GstBaseTransform
}

func NewAudiodynamic() (*Audiodynamic, error) {
	e, err := gst.NewElement("audiodynamic")
	if err != nil {
		return nil, err
	}
	return &Audiodynamic{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiodynamicWithName(name string) (*Audiodynamic, error) {
	e, err := gst.NewElementWithName("audiodynamic", name)
	if err != nil {
		return nil, err
	}
	return &Audiodynamic{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// characteristics (GstAudioDynamicCharacteristics)
//
// Selects whether the ratio should be applied smooth (soft-knee) or hard (hard-knee).

func (e *Audiodynamic) GetCharacteristics() (GstAudioDynamicCharacteristics, error) {
	var value GstAudioDynamicCharacteristics
	var ok bool
	v, err := e.Element.GetProperty("characteristics")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioDynamicCharacteristics)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioDynamicCharacteristics")
	}
	return value, nil
}

func (e *Audiodynamic) SetCharacteristics(value GstAudioDynamicCharacteristics) error {
	e.Element.SetArg("characteristics", string(value))
	return nil
}

// mode (GstAudioDynamicMode)
//
// Selects whether the filter should work on loud samples (compressor) orquiet samples (expander).

func (e *Audiodynamic) GetMode() (GstAudioDynamicMode, error) {
	var value GstAudioDynamicMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioDynamicMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioDynamicMode")
	}
	return value, nil
}

func (e *Audiodynamic) SetMode(value GstAudioDynamicMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ratio (float32)
//
// Ratio that should be applied

func (e *Audiodynamic) GetRatio() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiodynamic) SetRatio(value float32) error {
	return e.Element.SetProperty("ratio", value)
}

// threshold (float32)
//
// Threshold until the filter is activated

func (e *Audiodynamic) GetThreshold() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiodynamic) SetThreshold(value float32) error {
	return e.Element.SetProperty("threshold", value)
}

// ----- Constants -----

type GstAudioDynamicCharacteristics string

const (
	GstAudioDynamicCharacteristicsHardKnee GstAudioDynamicCharacteristics = "hard-knee" // hard-knee (0) – Hard Knee (default)
	GstAudioDynamicCharacteristicsSoftKnee GstAudioDynamicCharacteristics = "soft-knee" // soft-knee (1) – Soft Knee (smooth)
)

type GstAudioDynamicMode string

const (
	GstAudioDynamicModeCompressor GstAudioDynamicMode = "compressor" // compressor (0) – Compressor (default)
	GstAudioDynamicModeExpander GstAudioDynamicMode = "expander" // expander (1) – Expander
)

