// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Removes all color information except for one color
type GstChromaHold struct {
	*GstVideoFilter
}

func NewChromaHold() (*GstChromaHold, error) {
	e, err := gst.NewElement("chromahold")
	if err != nil {
		return nil, err
	}
	return &GstChromaHold{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewChromaHoldWithName(name string) (*GstChromaHold, error) {
	e, err := gst.NewElementWithName("chromahold", name)
	if err != nil {
		return nil, err
	}
	return &GstChromaHold{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The Blue target
// Default: 0, Min: 0, Max: 255
func (e *GstChromaHold) SetTargetB(value uint) error {
	return e.Element.SetProperty("target-b", value)
}

func (e *GstChromaHold) GetTargetB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The Green target
// Default: 0, Min: 0, Max: 255
func (e *GstChromaHold) SetTargetG(value uint) error {
	return e.Element.SetProperty("target-g", value)
}

func (e *GstChromaHold) GetTargetG() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The Red target
// Default: 255, Min: 0, Max: 255
func (e *GstChromaHold) SetTargetR(value uint) error {
	return e.Element.SetProperty("target-r", value)
}

func (e *GstChromaHold) GetTargetR() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Tolerance for the target color
// Default: 30, Min: 0, Max: 180
func (e *GstChromaHold) SetTolerance(value uint) error {
	return e.Element.SetProperty("tolerance", value)
}

func (e *GstChromaHold) GetTolerance() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Color Look-up Table filter
type GstColorEffects struct {
	*GstVideoFilter
}

func NewColorEffects() (*GstColorEffects, error) {
	e, err := gst.NewElement("coloreffects")
	if err != nil {
		return nil, err
	}
	return &GstColorEffects{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewColorEffectsWithName(name string) (*GstColorEffects, error) {
	e, err := gst.NewElementWithName("coloreffects", name)
	if err != nil {
		return nil, err
	}
	return &GstColorEffects{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Color effect preset to use
// Default: none (0)
func (e *GstColorEffects) SetPreset(value GstColorEffectsPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstColorEffects) GetPreset() (GstColorEffectsPreset, error) {
	var value GstColorEffectsPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstColorEffectsPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstColorEffectsPreset")
	}
	return value, nil
}

type GstColorEffectsPreset string

const (
	GstColorEffectsPresetNone       GstColorEffectsPreset = "none"       // Do nothing preset
	GstColorEffectsPresetHeat       GstColorEffectsPreset = "heat"       // Fake heat camera toning
	GstColorEffectsPresetSepia      GstColorEffectsPreset = "sepia"      // Sepia toning
	GstColorEffectsPresetXray       GstColorEffectsPreset = "xray"       // Invert and slightly shade to blue
	GstColorEffectsPresetXpro       GstColorEffectsPreset = "xpro"       // Cross processing toning
	GstColorEffectsPresetYellowblue GstColorEffectsPreset = "yellowblue" // Yellow foreground Blue background color filter
)
