// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Marks a video signal with a pattern
type GstSimpleVideoMark struct {
	*GstVideoFilter
}

func NewSimpleVideoMark() (*GstSimpleVideoMark, error) {
	e, err := gst.NewElement("simplevideomark")
	if err != nil {
		return nil, err
	}
	return &GstSimpleVideoMark{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewSimpleVideoMarkWithName(name string) (*GstSimpleVideoMark, error) {
	e, err := gst.NewElementWithName("simplevideomark", name)
	if err != nil {
		return nil, err
	}
	return &GstSimpleVideoMark{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The extra data pattern markers
// Default: 10, Min: 0, Max: 18446744073709551615
func (e *GstSimpleVideoMark) SetPatternData(value uint64) error {
	return e.Element.SetProperty("pattern-data", value)
}

func (e *GstSimpleVideoMark) GetPatternData() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("pattern-data")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The number of extra data pattern markers
// Default: 5, Min: 0, Max: 64
func (e *GstSimpleVideoMark) SetPatternDataCount(value int) error {
	return e.Element.SetProperty("pattern-data-count", value)
}

func (e *GstSimpleVideoMark) GetPatternDataCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-data-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The height of the pattern markers
// Default: 16, Min: 1, Max: 2147483647
func (e *GstSimpleVideoMark) SetPatternHeight(value int) error {
	return e.Element.SetProperty("pattern-height", value)
}

func (e *GstSimpleVideoMark) GetPatternHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The width of the pattern markers
// Default: 4, Min: 1, Max: 2147483647
func (e *GstSimpleVideoMark) SetPatternWidth(value int) error {
	return e.Element.SetProperty("pattern-width", value)
}

func (e *GstSimpleVideoMark) GetPatternWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The offset from the bottom border where the pattern starts
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSimpleVideoMark) SetBottomOffset(value int) error {
	return e.Element.SetProperty("bottom-offset", value)
}

func (e *GstSimpleVideoMark) GetBottomOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bottom-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Enable or disable the filter
// Default: true
func (e *GstSimpleVideoMark) SetEnabled(value bool) error {
	return e.Element.SetProperty("enabled", value)
}

func (e *GstSimpleVideoMark) GetEnabled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enabled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The offset from the left border where the pattern starts
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSimpleVideoMark) SetLeftOffset(value int) error {
	return e.Element.SetProperty("left-offset", value)
}

func (e *GstSimpleVideoMark) GetLeftOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("left-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The number of pattern markers
// Default: 4, Min: 0, Max: 2147483647
func (e *GstSimpleVideoMark) SetPatternCount(value int) error {
	return e.Element.SetProperty("pattern-count", value)
}

func (e *GstSimpleVideoMark) GetPatternCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Detect patterns in a video signal
type GstSimpleVideoMarkDetect struct {
	*GstVideoFilter
}

func NewSimpleVideoMarkDetect() (*GstSimpleVideoMarkDetect, error) {
	e, err := gst.NewElement("simplevideomarkdetect")
	if err != nil {
		return nil, err
	}
	return &GstSimpleVideoMarkDetect{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewSimpleVideoMarkDetectWithName(name string) (*GstSimpleVideoMarkDetect, error) {
	e, err := gst.NewElementWithName("simplevideomarkdetect", name)
	if err != nil {
		return nil, err
	}
	return &GstSimpleVideoMarkDetect{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The number of pattern markers
// Default: 4, Min: 0, Max: 2147483647
func (e *GstSimpleVideoMarkDetect) SetPatternCount(value int) error {
	return e.Element.SetProperty("pattern-count", value)
}

func (e *GstSimpleVideoMarkDetect) GetPatternCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The number of extra data pattern markers
// Default: 5, Min: 0, Max: 2147483647
func (e *GstSimpleVideoMarkDetect) SetPatternDataCount(value int) error {
	return e.Element.SetProperty("pattern-data-count", value)
}

func (e *GstSimpleVideoMarkDetect) GetPatternDataCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-data-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The height of the pattern markers
// Default: 16, Min: 1, Max: 2147483647
func (e *GstSimpleVideoMarkDetect) SetPatternHeight(value int) error {
	return e.Element.SetProperty("pattern-height", value)
}

func (e *GstSimpleVideoMarkDetect) GetPatternHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The sensitivity around the center for detecting the markers (0.0 = lowest, 1.0 highest)
// Default: 0.3, Min: 0, Max: 1
func (e *GstSimpleVideoMarkDetect) SetPatternSensitivity(value float64) error {
	return e.Element.SetProperty("pattern-sensitivity", value)
}

func (e *GstSimpleVideoMarkDetect) GetPatternSensitivity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("pattern-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The offset from the bottom border where the pattern starts
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSimpleVideoMarkDetect) SetBottomOffset(value int) error {
	return e.Element.SetProperty("bottom-offset", value)
}

func (e *GstSimpleVideoMarkDetect) GetBottomOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bottom-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Post detected data as bus messages
// Default: true
func (e *GstSimpleVideoMarkDetect) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

func (e *GstSimpleVideoMarkDetect) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The center of the black/white separation (0.0 = lowest, 1.0 highest)
// Default: 0.5, Min: 0, Max: 1
func (e *GstSimpleVideoMarkDetect) SetPatternCenter(value float64) error {
	return e.Element.SetProperty("pattern-center", value)
}

func (e *GstSimpleVideoMarkDetect) GetPatternCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("pattern-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The offset from the left border where the pattern starts
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSimpleVideoMarkDetect) SetLeftOffset(value int) error {
	return e.Element.SetProperty("left-offset", value)
}

func (e *GstSimpleVideoMarkDetect) GetLeftOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("left-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The width of the pattern markers
// Default: 4, Min: 1, Max: 2147483647
func (e *GstSimpleVideoMarkDetect) SetPatternWidth(value int) error {
	return e.Element.SetProperty("pattern-width", value)
}

func (e *GstSimpleVideoMarkDetect) GetPatternWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pattern-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Analyse video signal
type GstVideoAnalyse struct {
	*GstVideoFilter
}

func NewVideoAnalyse() (*GstVideoAnalyse, error) {
	e, err := gst.NewElement("videoanalyse")
	if err != nil {
		return nil, err
	}
	return &GstVideoAnalyse{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoAnalyseWithName(name string) (*GstVideoAnalyse, error) {
	e, err := gst.NewElementWithName("videoanalyse", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoAnalyse{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Post statics messages
// Default: true
func (e *GstVideoAnalyse) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

func (e *GstVideoAnalyse) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
