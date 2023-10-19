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

type Fieldanalysis struct {
	Element *gst.Element
}

func NewFieldanalysis() (*Fieldanalysis, error) {
	e, err := gst.NewElement("fieldanalysis")
	if err != nil {
		return nil, err
	}
	return &Fieldanalysis{Element: e}, nil
}

func NewFieldanalysisWithName(name string) (*Fieldanalysis, error) {
	e, err := gst.NewElementWithName("fieldanalysis", name)
	if err != nil {
		return nil, err
	}
	return &Fieldanalysis{Element: e}, nil
}

// ----- Properties -----

// block-height (uint64)
//
// Block height for windowed comb detection

func (e *Fieldanalysis) GetBlockHeight() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("block-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fieldanalysis) SetBlockHeight(value uint64) error {
	return e.Element.SetProperty("block-height", value)
}

// block-threshold (uint64)
//
// Block threshold for windowed comb detection

func (e *Fieldanalysis) GetBlockThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("block-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fieldanalysis) SetBlockThreshold(value uint64) error {
	return e.Element.SetProperty("block-threshold", value)
}

// block-width (uint64)
//
// Block width for windowed comb detection

func (e *Fieldanalysis) GetBlockWidth() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("block-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fieldanalysis) SetBlockWidth(value uint64) error {
	return e.Element.SetProperty("block-width", value)
}

// comb-method (GstFieldAnalysisCombMethod)
//
// Metric to be used for identifying comb artifacts if using windowed comb detection

func (e *Fieldanalysis) GetCombMethod() (interface{}, error) {
	return e.Element.GetProperty("comb-method")
}

func (e *Fieldanalysis) SetCombMethod(value interface{}) error {
	return e.Element.SetProperty("comb-method", value)
}

// field-metric (GstFieldAnalysisFieldMetric)
//
// Metric to be used for comparing same parity fields to decide if they are a repeated field for telecine

func (e *Fieldanalysis) GetFieldMetric() (GstFieldAnalysisFieldMetric, error) {
	var value GstFieldAnalysisFieldMetric
	var ok bool
	v, err := e.Element.GetProperty("field-metric")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFieldAnalysisFieldMetric)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFieldAnalysisFieldMetric")
	}
	return value, nil
}

func (e *Fieldanalysis) SetFieldMetric(value GstFieldAnalysisFieldMetric) error {
	e.Element.SetArg("field-metric", string(value))
	return nil
}

// field-threshold (float32)
//
// Threshold for field metric decisions

func (e *Fieldanalysis) GetFieldThreshold() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("field-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Fieldanalysis) SetFieldThreshold(value float32) error {
	return e.Element.SetProperty("field-threshold", value)
}

// frame-metric (GstFieldAnalysisFrameMetric)
//
// Metric to be used for comparing opposite parity fields to decide if they are a progressive frame

func (e *Fieldanalysis) GetFrameMetric() (GstFieldAnalysisFrameMetric, error) {
	var value GstFieldAnalysisFrameMetric
	var ok bool
	v, err := e.Element.GetProperty("frame-metric")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFieldAnalysisFrameMetric)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFieldAnalysisFrameMetric")
	}
	return value, nil
}

func (e *Fieldanalysis) SetFrameMetric(value GstFieldAnalysisFrameMetric) error {
	e.Element.SetArg("frame-metric", string(value))
	return nil
}

// frame-threshold (float32)
//
// Threshold for frame metric decisions

func (e *Fieldanalysis) GetFrameThreshold() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frame-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Fieldanalysis) SetFrameThreshold(value float32) error {
	return e.Element.SetProperty("frame-threshold", value)
}

// ignored-lines (uint64)
//
// Ignore this many lines from the top and bottom for windowed comb detection

func (e *Fieldanalysis) GetIgnoredLines() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ignored-lines")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fieldanalysis) SetIgnoredLines(value uint64) error {
	return e.Element.SetProperty("ignored-lines", value)
}

// noise-floor (uint)
//
// Noise floor for appropriate metrics (per-pixel metric values with a score less than this will be ignored)

func (e *Fieldanalysis) GetNoiseFloor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("noise-floor")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Fieldanalysis) SetNoiseFloor(value uint) error {
	return e.Element.SetProperty("noise-floor", value)
}

// spatial-threshold (int64)
//
// Threshold for combing metric decisions

func (e *Fieldanalysis) GetSpatialThreshold() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("spatial-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Fieldanalysis) SetSpatialThreshold(value int64) error {
	return e.Element.SetProperty("spatial-threshold", value)
}

// ----- Constants -----

type GstFieldAnalysisFieldMetric string

const (
	GstFieldAnalysisFieldMetricSad GstFieldAnalysisFieldMetric = "sad" // sad (0) – Sum of Absolute Differences
	GstFieldAnalysisFieldMetricSsd GstFieldAnalysisFieldMetric = "ssd" // ssd (1) – Sum of Squared Differences
	GstFieldAnalysisFieldMetric3Tap GstFieldAnalysisFieldMetric = "3-tap" // 3-tap (2) – Difference of 3-tap [1,4,1] Horizontal Filter
)

type GstFieldAnalysisFrameMetric string

const (
	GstFieldAnalysisFrameMetric5Tap GstFieldAnalysisFrameMetric = "5-tap" // 5-tap (0) – 5-tap [1,-3,4,-3,1] Vertical Filter
	GstFieldAnalysisFrameMetricWindowedComb GstFieldAnalysisFrameMetric = "windowed-comb" // windowed-comb (1) – Windowed Comb Detection (not optimised)
)

type FieldAnalysisCombMethod string

const (
	FieldAnalysisCombMethod32Detect FieldAnalysisCombMethod = "32-detect" // 32-detect (0) – Difference to above sample in same field small and difference to sample in other field large
	FieldAnalysisCombMethodIsCombed FieldAnalysisCombMethod = "isCombed" // isCombed (1) – Differences between current sample and the above/below samples in other field multiplied together, larger than squared spatial threshold (from Tritical's isCombed)
	FieldAnalysisCombMethod5Tap FieldAnalysisCombMethod = "5-tap" // 5-tap (2) – 5-tap [1,-3,4,-3,1] vertical filter result is larger than spatial threshold*6
)

