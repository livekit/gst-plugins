// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Creates an interlaced video from progressive frames
type GstInterlace struct {
	*gst.Element
}

func NewInterlace() (*GstInterlace, error) {
	e, err := gst.NewElement("interlace")
	if err != nil {
		return nil, err
	}
	return &GstInterlace{Element: e}, nil
}

func NewInterlaceWithName(name string) (*GstInterlace, error) {
	e, err := gst.NewElementWithName("interlace", name)
	if err != nil {
		return nil, err
	}
	return &GstInterlace{Element: e}, nil
}

// Allow generation of buffers with RFF flag set, i.e., duration of 3 fields
// Default: false
func (e *GstInterlace) SetAllowRff(value bool) error {
	return e.Element.SetProperty("allow-rff", value)
}

func (e *GstInterlace) GetAllowRff() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-rff")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The output field pattern
// Default: 2:3 (2)
func (e *GstInterlace) SetFieldPattern(value GstInterlacePattern) error {
	e.Element.SetArg("field-pattern", string(value))
	return nil
}

func (e *GstInterlace) GetFieldPattern() (GstInterlacePattern, error) {
	var value GstInterlacePattern
	var ok bool
	v, err := e.Element.GetProperty("field-pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstInterlacePattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstInterlacePattern")
	}
	return value, nil
}

// The initial field pattern offset. Counts from 0.
// Default: 0, Min: 0, Max: 12
func (e *GstInterlace) SetPatternOffset(value uint) error {
	return e.Element.SetProperty("pattern-offset", value)
}

func (e *GstInterlace) GetPatternOffset() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pattern-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Interlaced stream should be top field first
// Default: false
func (e *GstInterlace) SetTopFieldFirst(value bool) error {
	return e.Element.SetProperty("top-field-first", value)
}

func (e *GstInterlace) GetTopFieldFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("top-field-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstInterlacePattern string

const (
	GstInterlacePattern11   GstInterlacePattern = "1:1"     // 1:1 (e.g. 60p -> 60i)
	GstInterlacePattern22   GstInterlacePattern = "2:2"     // 2:2 (e.g. 30p -> 60i)
	GstInterlacePattern23   GstInterlacePattern = "2:3"     // 2:3 (e.g. 24p -> 60i telecine)
	GstInterlacePattern2332 GstInterlacePattern = "2:3:3:2" // 2:3:3:2 (e.g. 24p -> 60i telecine)
	GstInterlacePattern2113 GstInterlacePattern = "2-11:3"  // Euro 2-11:3 (e.g. 24p -> 50i telecine)
	GstInterlacePattern343  GstInterlacePattern = "3:4-3"   // 3:4-3 (e.g. 16p -> 60i telecine)
	GstInterlacePattern374  GstInterlacePattern = "3-7:4"   // 3-7:4 (e.g. 16p -> 50i telecine)
	GstInterlacePattern334  GstInterlacePattern = "3:3:4"   // 3:3:4 (e.g. 18p -> 60i telecine)
	GstInterlacePattern33   GstInterlacePattern = "3:3"     // 3:3 (e.g. 20p -> 60i telecine)
	GstInterlacePattern324  GstInterlacePattern = "3:2-4"   // 3:2-4 (e.g. 27.5p -> 60i telecine)
	GstInterlacePattern124  GstInterlacePattern = "1:2-4"   // 1:2-4 (e.g. 27.5p -> 50i telecine)
)
