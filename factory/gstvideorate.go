// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Drops/duplicates/adjusts timestamps on video frames to make a perfect stream
type GstVideoRate struct {
	*base.GstBaseTransform
}

func NewVideoRate() (*GstVideoRate, error) {
	e, err := gst.NewElement("videorate")
	if err != nil {
		return nil, err
	}
	return &GstVideoRate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideoRateWithName(name string) (*GstVideoRate, error) {
	e, err := gst.NewElementWithName("videorate", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoRate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Number of input frames
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstVideoRate) GetIn() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("in")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum duration of duplicated buffers to close current segment
// Default: 1000000000, Min: 0, Max: 18446744073709551615
func (e *GstVideoRate) SetMaxClosingSegmentDuplicationDuration(value uint64) error {
	return e.Element.SetProperty("max-closing-segment-duplication-duration", value)
}

func (e *GstVideoRate) GetMaxClosingSegmentDuplicationDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-closing-segment-duplication-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Do not duplicate frames if the gap exceeds this period (in ns) (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstVideoRate) SetMaxDuplicationTime(value uint64) error {
	return e.Element.SetProperty("max-duplication-time", value)
}

func (e *GstVideoRate) GetMaxDuplicationTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-duplication-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Don't emit notify for dropped and duplicated frames
// Default: true
func (e *GstVideoRate) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstVideoRate) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Don't produce buffers before the first one we receive
// Default: false
func (e *GstVideoRate) SetSkipToFirst(value bool) error {
	return e.Element.SetProperty("skip-to-first", value)
}

func (e *GstVideoRate) GetSkipToFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-to-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Period over which to average the framerate (in ns) (0 = disabled)
// Default: 0, Min: 0, Max: 9223372036854775807
func (e *GstVideoRate) SetAveragePeriod(value uint64) error {
	return e.Element.SetProperty("average-period", value)
}

func (e *GstVideoRate) GetAveragePeriod() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("average-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Number of duplicated frames
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstVideoRate) GetDuplicate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("duplicate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum framerate allowed to pass through (in frames per second, implies drop-only)
// Default: 2147483647, Min: 1, Max: 2147483647
func (e *GstVideoRate) SetMaxRate(value int) error {
	return e.Element.SetProperty("max-rate", value)
}

func (e *GstVideoRate) GetMaxRate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Value indicating how much to prefer new frames (unused)
// Default: 1, Min: 0, Max: 1
func (e *GstVideoRate) SetNewPref(value float64) error {
	return e.Element.SetProperty("new-pref", value)
}

func (e *GstVideoRate) GetNewPref() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("new-pref")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Number of output frames
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstVideoRate) GetOut() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("out")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Factor of speed for frame displaying
// Default: 1, Min: 0, Max: 1.79769e+308
func (e *GstVideoRate) SetRate(value float64) error {
	return e.Element.SetProperty("rate", value)
}

func (e *GstVideoRate) GetRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Number of dropped frames
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstVideoRate) GetDrop() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Only drop frames, no duplicates are produced
// Default: false
func (e *GstVideoRate) SetDropOnly(value bool) error {
	return e.Element.SetProperty("drop-only", value)
}

func (e *GstVideoRate) GetDropOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
