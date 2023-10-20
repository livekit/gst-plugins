// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Analyse fields from video frames to identify if they are progressive/telecined/interlaced
type GstFieldAnalysis struct {
	*gst.Element
}

func NewFieldAnalysis() (*GstFieldAnalysis, error) {
	e, err := gst.NewElement("fieldanalysis")
	if err != nil {
		return nil, err
	}
	return &GstFieldAnalysis{Element: e}, nil
}

func NewFieldAnalysisWithName(name string) (*GstFieldAnalysis, error) {
	e, err := gst.NewElementWithName("fieldanalysis", name)
	if err != nil {
		return nil, err
	}
	return &GstFieldAnalysis{Element: e}, nil
}

// Metric to be used for comparing same parity fields to decide if they are a repeated field for telecine
// Default: ssd (1)
func (e *GstFieldAnalysis) SetFieldMetric(value GstFieldAnalysisFieldMetric) error {
	e.Element.SetArg("field-metric", string(value))
	return nil
}

func (e *GstFieldAnalysis) GetFieldMetric() (GstFieldAnalysisFieldMetric, error) {
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

// Threshold for field metric decisions
// Default: 0.08, Min: 0, Max: 3.40282e+38
func (e *GstFieldAnalysis) SetFieldThreshold(value float32) error {
	return e.Element.SetProperty("field-threshold", value)
}

func (e *GstFieldAnalysis) GetFieldThreshold() (float32, error) {
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

// Threshold for frame metric decisions
// Default: 0.002, Min: 0, Max: 3.40282e+38
func (e *GstFieldAnalysis) SetFrameThreshold(value float32) error {
	return e.Element.SetProperty("frame-threshold", value)
}

func (e *GstFieldAnalysis) GetFrameThreshold() (float32, error) {
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

// Ignore this many lines from the top and bottom for windowed comb detection
// Default: 2, Min: 2, Max: 18446744073709551615
func (e *GstFieldAnalysis) SetIgnoredLines(value uint64) error {
	return e.Element.SetProperty("ignored-lines", value)
}

func (e *GstFieldAnalysis) GetIgnoredLines() (uint64, error) {
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

// Noise floor for appropriate metrics (per-pixel metric values with a score less than this will be ignored)
// Default: 16, Min: 0, Max: -1
func (e *GstFieldAnalysis) SetNoiseFloor(value uint) error {
	return e.Element.SetProperty("noise-floor", value)
}

func (e *GstFieldAnalysis) GetNoiseFloor() (uint, error) {
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

// Block height for windowed comb detection
// Default: 16, Min: 0, Max: 18446744073709551615
func (e *GstFieldAnalysis) SetBlockHeight(value uint64) error {
	return e.Element.SetProperty("block-height", value)
}

func (e *GstFieldAnalysis) GetBlockHeight() (uint64, error) {
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

// Block width for windowed comb detection
// Default: 16, Min: 1, Max: 18446744073709551615
func (e *GstFieldAnalysis) SetBlockWidth(value uint64) error {
	return e.Element.SetProperty("block-width", value)
}

func (e *GstFieldAnalysis) GetBlockWidth() (uint64, error) {
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

// Metric to be used for identifying comb artifacts if using windowed comb detection
// Default: 5-tap (2)
func (e *GstFieldAnalysis) SetCombMethod(value FieldAnalysisCombMethod) error {
	e.Element.SetArg("comb-method", string(value))
	return nil
}

func (e *GstFieldAnalysis) GetCombMethod() (FieldAnalysisCombMethod, error) {
	var value FieldAnalysisCombMethod
	var ok bool
	v, err := e.Element.GetProperty("comb-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(FieldAnalysisCombMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to FieldAnalysisCombMethod")
	}
	return value, nil
}

// Threshold for combing metric decisions
// Default: 9, Min: 0, Max: 9223372036854775807
func (e *GstFieldAnalysis) SetSpatialThreshold(value int64) error {
	return e.Element.SetProperty("spatial-threshold", value)
}

func (e *GstFieldAnalysis) GetSpatialThreshold() (int64, error) {
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

// Block threshold for windowed comb detection
// Default: 80, Min: 0, Max: 18446744073709551615
func (e *GstFieldAnalysis) SetBlockThreshold(value uint64) error {
	return e.Element.SetProperty("block-threshold", value)
}

func (e *GstFieldAnalysis) GetBlockThreshold() (uint64, error) {
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

// Metric to be used for comparing opposite parity fields to decide if they are a progressive frame
// Default: 5-tap (0)
func (e *GstFieldAnalysis) SetFrameMetric(value GstFieldAnalysisFrameMetric) error {
	e.Element.SetArg("frame-metric", string(value))
	return nil
}

func (e *GstFieldAnalysis) GetFrameMetric() (GstFieldAnalysisFrameMetric, error) {
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

type GstFieldAnalysisFieldMetric string

const (
	GstFieldAnalysisFieldMetricSad  GstFieldAnalysisFieldMetric = "sad"   // Sum of Absolute Differences
	GstFieldAnalysisFieldMetricSsd  GstFieldAnalysisFieldMetric = "ssd"   // Sum of Squared Differences
	GstFieldAnalysisFieldMetric3Tap GstFieldAnalysisFieldMetric = "3-tap" // Difference of 3-tap [1,4,1] Horizontal Filter
)

type GstFieldAnalysisFrameMetric string

const (
	GstFieldAnalysisFrameMetric5Tap         GstFieldAnalysisFrameMetric = "5-tap"         // 5-tap [1,-3,4,-3,1] Vertical Filter
	GstFieldAnalysisFrameMetricWindowedComb GstFieldAnalysisFrameMetric = "windowed-comb" // Windowed Comb Detection (not optimised)
)

type FieldAnalysisCombMethod string

const (
	FieldAnalysisCombMethod32Detect FieldAnalysisCombMethod = "32-detect" // Difference to above sample in same field small and difference to sample in other field large
	FieldAnalysisCombMethodIsCombed FieldAnalysisCombMethod = "isCombed"  // Differences between current sample and the above/below samples in other field multiplied together, larger than squared spatial threshold (from Tritical's isCombed)
	FieldAnalysisCombMethod5Tap     FieldAnalysisCombMethod = "5-tap"     // 5-tap [1,-3,4,-3,1] vertical filter result is larger than spatial threshold*6
)
