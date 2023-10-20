// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Deinterlace Methods ported from DScaler/TvTime
type GstDeinterlace struct {
	*gst.Element
}

func NewDeinterlace() (*GstDeinterlace, error) {
	e, err := gst.NewElement("deinterlace")
	if err != nil {
		return nil, err
	}
	return &GstDeinterlace{Element: e}, nil
}

func NewDeinterlaceWithName(name string) (*GstDeinterlace, error) {
	e, err := gst.NewElementWithName("deinterlace", name)
	if err != nil {
		return nil, err
	}
	return &GstDeinterlace{Element: e}, nil
}

// Fields to use for deinterlacing
// Default: all (0)
func (e *GstDeinterlace) SetFields(value GstDeinterlaceFields) error {
	e.Element.SetArg("fields", string(value))
	return nil
}

func (e *GstDeinterlace) GetFields() (GstDeinterlaceFields, error) {
	var value GstDeinterlaceFields
	var ok bool
	v, err := e.Element.GetProperty("fields")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceFields)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceFields")
	}
	return value, nil
}

// Ignore obscure telecine patterns (only consider P, I and 2:3 variants).
// Default: true
func (e *GstDeinterlace) SetIgnoreObscure(value bool) error {
	return e.Element.SetProperty("ignore-obscure", value)
}

func (e *GstDeinterlace) GetIgnoreObscure() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-obscure")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Pattern locking mode
// Default: none (0)
func (e *GstDeinterlace) SetLocking(value GstDeinterlaceLocking) error {
	e.Element.SetArg("locking", string(value))
	return nil
}

func (e *GstDeinterlace) GetLocking() (GstDeinterlaceLocking, error) {
	var value GstDeinterlaceLocking
	var ok bool
	v, err := e.Element.GetProperty("locking")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceLocking)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceLocking")
	}
	return value, nil
}

// Deinterlace Method
// Default: linear (4)
func (e *GstDeinterlace) SetMethod(value GstDeinterlaceMethods) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstDeinterlace) GetMethod() (GstDeinterlaceMethods, error) {
	var value GstDeinterlaceMethods
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceMethods)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceMethods")
	}
	return value, nil
}

// Deinterlace Mode
// Default: auto (0)
func (e *GstDeinterlace) SetMode(value GstDeinterlaceModes) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstDeinterlace) GetMode() (GstDeinterlaceModes, error) {
	var value GstDeinterlaceModes
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceModes")
	}
	return value, nil
}

// Deinterlace top field first
// Default: auto (0)
func (e *GstDeinterlace) SetTff(value GstDeinterlaceFieldLayout) error {
	e.Element.SetArg("tff", string(value))
	return nil
}

func (e *GstDeinterlace) GetTff() (GstDeinterlaceFieldLayout, error) {
	var value GstDeinterlaceFieldLayout
	var ok bool
	v, err := e.Element.GetProperty("tff")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceFieldLayout)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceFieldLayout")
	}
	return value, nil
}

// Drop orphan fields at the beginning of telecine patterns in active locking mode.
// Default: true
func (e *GstDeinterlace) SetDropOrphans(value bool) error {
	return e.Element.SetProperty("drop-orphans", value)
}

func (e *GstDeinterlace) GetDropOrphans() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-orphans")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstDeinterlaceFieldLayout string

const (
	GstDeinterlaceFieldLayoutAuto GstDeinterlaceFieldLayout = "auto" // Auto detection
	GstDeinterlaceFieldLayoutTff  GstDeinterlaceFieldLayout = "tff"  // Top field first
	GstDeinterlaceFieldLayoutBff  GstDeinterlaceFieldLayout = "bff"  // Bottom field first
)

type GstDeinterlaceFields string

const (
	GstDeinterlaceFieldsAll    GstDeinterlaceFields = "all"    // All fields
	GstDeinterlaceFieldsTop    GstDeinterlaceFields = "top"    // Top fields only
	GstDeinterlaceFieldsBottom GstDeinterlaceFields = "bottom" // Bottom fields only
	GstDeinterlaceFieldsAuto   GstDeinterlaceFields = "auto"   // Automatically detect
)

type GstDeinterlaceLocking string

const (
	GstDeinterlaceLockingNone    GstDeinterlaceLocking = "none"    // No pattern locking
	GstDeinterlaceLockingAuto    GstDeinterlaceLocking = "auto"    // Choose passive/active locking depending on whether upstream is live
	GstDeinterlaceLockingActive  GstDeinterlaceLocking = "active"  // Block until pattern-locked. Use accurate timestamp interpolation within a pattern repeat.
	GstDeinterlaceLockingPassive GstDeinterlaceLocking = "passive" // Do not block. Use naïve timestamp adjustment until pattern-locked based on state history.
)

type GstDeinterlaceMethods string

const (
	GstDeinterlaceMethodsTomsmocomp  GstDeinterlaceMethods = "tomsmocomp"  // Motion Adaptive: Motion Search
	GstDeinterlaceMethodsGreedyh     GstDeinterlaceMethods = "greedyh"     // Motion Adaptive: Advanced Detection
	GstDeinterlaceMethodsGreedyl     GstDeinterlaceMethods = "greedyl"     // Motion Adaptive: Simple Detection
	GstDeinterlaceMethodsVfir        GstDeinterlaceMethods = "vfir"        // Blur Vertical
	GstDeinterlaceMethodsLinear      GstDeinterlaceMethods = "linear"      // Linear
	GstDeinterlaceMethodsLinearblend GstDeinterlaceMethods = "linearblend" // Blur: Temporal (Do Not Use)
	GstDeinterlaceMethodsScalerbob   GstDeinterlaceMethods = "scalerbob"   // Double lines
	GstDeinterlaceMethodsWeave       GstDeinterlaceMethods = "weave"       // Weave (Do Not Use)
	GstDeinterlaceMethodsWeavetff    GstDeinterlaceMethods = "weavetff"    // Progressive: Top Field First (Do Not Use)
	GstDeinterlaceMethodsWeavebff    GstDeinterlaceMethods = "weavebff"    // Progressive: Bottom Field First (Do Not Use)
	GstDeinterlaceMethodsYadif       GstDeinterlaceMethods = "yadif"       // YADIF Adaptive Deinterlacer
)

type GstDeinterlaceModes string

const (
	GstDeinterlaceModesAuto       GstDeinterlaceModes = "auto"        // Auto detection (best effort)
	GstDeinterlaceModesInterlaced GstDeinterlaceModes = "interlaced"  // Force deinterlacing
	GstDeinterlaceModesDisabled   GstDeinterlaceModes = "disabled"    // Run in passthrough mode
	GstDeinterlaceModesAutoStrict GstDeinterlaceModes = "auto-strict" // Auto detection (strict)
)
