// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Libvisual corona plugin
type GstVisualcorona struct {
	*GstVisual
}

func NewVisualcorona() (*GstVisualcorona, error) {
	e, err := gst.NewElement("libvisual_corona")
	if err != nil {
		return nil, err
	}
	return &GstVisualcorona{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisualcoronaWithName(name string) (*GstVisualcorona, error) {
	e, err := gst.NewElementWithName("libvisual_corona", name)
	if err != nil {
		return nil, err
	}
	return &GstVisualcorona{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisualcorona) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisualcorona) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisualcorona) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisualcorona) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Infinite visual plugin
type GstVisualinfinite struct {
	*GstVisual
}

func NewVisualinfinite() (*GstVisualinfinite, error) {
	e, err := gst.NewElement("libvisual_infinite")
	if err != nil {
		return nil, err
	}
	return &GstVisualinfinite{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisualinfiniteWithName(name string) (*GstVisualinfinite, error) {
	e, err := gst.NewElementWithName("libvisual_infinite", name)
	if err != nil {
		return nil, err
	}
	return &GstVisualinfinite{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisualinfinite) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisualinfinite) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisualinfinite) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisualinfinite) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// jakdaw visual plugin
type GstVisualjakdaw struct {
	*GstVisual
}

func NewVisualjakdaw() (*GstVisualjakdaw, error) {
	e, err := gst.NewElement("libvisual_jakdaw")
	if err != nil {
		return nil, err
	}
	return &GstVisualjakdaw{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisualjakdawWithName(name string) (*GstVisualjakdaw, error) {
	e, err := gst.NewElementWithName("libvisual_jakdaw", name)
	if err != nil {
		return nil, err
	}
	return &GstVisualjakdaw{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisualjakdaw) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisualjakdaw) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisualjakdaw) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisualjakdaw) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// Jess visual plugin
type GstVisualjess struct {
	*GstVisual
}

func NewVisualjess() (*GstVisualjess, error) {
	e, err := gst.NewElement("libvisual_jess")
	if err != nil {
		return nil, err
	}
	return &GstVisualjess{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisualjessWithName(name string) (*GstVisualjess, error) {
	e, err := gst.NewElementWithName("libvisual_jess", name)
	if err != nil {
		return nil, err
	}
	return &GstVisualjess{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisualjess) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisualjess) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisualjess) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisualjess) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// Libvisual analyzer plugin
type GstVisuallvAnalyzer struct {
	*GstVisual
}

func NewVisuallvAnalyzer() (*GstVisuallvAnalyzer, error) {
	e, err := gst.NewElement("libvisual_lv_analyzer")
	if err != nil {
		return nil, err
	}
	return &GstVisuallvAnalyzer{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisuallvAnalyzerWithName(name string) (*GstVisuallvAnalyzer, error) {
	e, err := gst.NewElementWithName("libvisual_lv_analyzer", name)
	if err != nil {
		return nil, err
	}
	return &GstVisuallvAnalyzer{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisuallvAnalyzer) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisuallvAnalyzer) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisuallvAnalyzer) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisuallvAnalyzer) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// Libvisual scope plugin
type GstVisuallvScope struct {
	*GstVisual
}

func NewVisuallvScope() (*GstVisuallvScope, error) {
	e, err := gst.NewElement("libvisual_lv_scope")
	if err != nil {
		return nil, err
	}
	return &GstVisuallvScope{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisuallvScopeWithName(name string) (*GstVisuallvScope, error) {
	e, err := gst.NewElementWithName("libvisual_lv_scope", name)
	if err != nil {
		return nil, err
	}
	return &GstVisuallvScope{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisuallvScope) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisuallvScope) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisuallvScope) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisuallvScope) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// Libvisual Oinksie visual plugin
type GstVisualoinksie struct {
	*GstVisual
}

func NewVisualoinksie() (*GstVisualoinksie, error) {
	e, err := gst.NewElement("libvisual_oinksie")
	if err != nil {
		return nil, err
	}
	return &GstVisualoinksie{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisualoinksieWithName(name string) (*GstVisualoinksie, error) {
	e, err := gst.NewElementWithName("libvisual_oinksie", name)
	if err != nil {
		return nil, err
	}
	return &GstVisualoinksie{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisualoinksie) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisualoinksie) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisualoinksie) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisualoinksie) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// Bumpscope visual plugin
type GstVisualbumpscope struct {
	*GstVisual
}

func NewVisualbumpscope() (*GstVisualbumpscope, error) {
	e, err := gst.NewElement("libvisual_bumpscope")
	if err != nil {
		return nil, err
	}
	return &GstVisualbumpscope{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

func NewVisualbumpscopeWithName(name string) (*GstVisualbumpscope, error) {
	e, err := gst.NewElementWithName("libvisual_bumpscope", name)
	if err != nil {
		return nil, err
	}
	return &GstVisualbumpscope{GstVisual: &GstVisual{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}}, nil
}

// Shading color to use (big-endian ARGB)
// Default: 657930, Min: 0, Max: -1
func (e *GstVisualbumpscope) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

func (e *GstVisualbumpscope) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Shader function to apply on each frame
// Default: fade (1)
func (e *GstVisualbumpscope) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstVisualbumpscope) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

type GstVisual struct {
	*GstAudioVisualizer
}
