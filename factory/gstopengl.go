// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Uploads data into OpenGL
type GstGLUploadElement struct {
	*GstGLBaseFilter
}

func NewGLUploadElement() (*GstGLUploadElement, error) {
	e, err := gst.NewElement("glupload")
	if err != nil {
		return nil, err
	}
	return &GstGLUploadElement{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewGLUploadElementWithName(name string) (*GstGLUploadElement, error) {
	e, err := gst.NewElementWithName("glupload", name)
	if err != nil {
		return nil, err
	}
	return &GstGLUploadElement{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// GL Shading Language effects - Sepia Toning Effect
type GleffectsSepia struct {
	*GstGLEffects
}

func NewGleffectsSepia() (*GleffectsSepia, error) {
	e, err := gst.NewElement("gleffects_sepia")
	if err != nil {
		return nil, err
	}
	return &GleffectsSepia{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsSepiaWithName(name string) (*GleffectsSepia, error) {
	e, err := gst.NewElementWithName("gleffects_sepia", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsSepia{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsSepia) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsSepia) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Splits a stereoscopic stream into separate left/right streams
type GstGLStereoSplit struct {
	*gst.Element
}

func NewGLStereoSplit() (*GstGLStereoSplit, error) {
	e, err := gst.NewElement("glstereosplit")
	if err != nil {
		return nil, err
	}
	return &GstGLStereoSplit{Element: e}, nil
}

func NewGLStereoSplitWithName(name string) (*GstGLStereoSplit, error) {
	e, err := gst.NewElementWithName("glstereosplit", name)
	if err != nil {
		return nil, err
	}
	return &GstGLStereoSplit{Element: e}, nil
}

// GL Shading Language effects - Square Effect
type GleffectsSquare struct {
	*GstGLEffects
}

func NewGleffectsSquare() (*GleffectsSquare, error) {
	e, err := gst.NewElement("gleffects_square")
	if err != nil {
		return nil, err
	}
	return &GleffectsSquare{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsSquareWithName(name string) (*GleffectsSquare, error) {
	e, err := gst.NewElementWithName("gleffects_square", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsSquare{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsSquare) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsSquare) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Squeeze Effect
type GleffectsSqueeze struct {
	*GstGLEffects
}

func NewGleffectsSqueeze() (*GleffectsSqueeze, error) {
	e, err := gst.NewElement("gleffects_squeeze")
	if err != nil {
		return nil, err
	}
	return &GleffectsSqueeze{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsSqueezeWithName(name string) (*GleffectsSqueeze, error) {
	e, err := gst.NewElementWithName("gleffects_squeeze", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsSqueeze{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsSqueeze) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsSqueeze) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Glowing negative effect
type GleffectsXray struct {
	*GstGLEffects
}

func NewGleffectsXray() (*GleffectsXray, error) {
	e, err := gst.NewElement("gleffects_xray")
	if err != nil {
		return nil, err
	}
	return &GleffectsXray{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsXrayWithName(name string) (*GleffectsXray, error) {
	e, err := gst.NewElementWithName("gleffects_xray", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsXray{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsXray) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsXray) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Perform operations with a GLSL shader
type GstGLFilterShader struct {
	*GstGLFilter
}

func NewGLFilterShader() (*GstGLFilterShader, error) {
	e, err := gst.NewElement("glshader")
	if err != nil {
		return nil, err
	}
	return &GstGLFilterShader{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLFilterShaderWithName(name string) (*GstGLFilterShader, error) {
	e, err := gst.NewElementWithName("glshader", name)
	if err != nil {
		return nil, err
	}
	return &GstGLFilterShader{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// GLSL fragment source
// Default: NULL
func (e *GstGLFilterShader) SetFragment(value string) error {
	return e.Element.SetProperty("fragment", value)
}

func (e *GstGLFilterShader) GetFragment() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("fragment")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// GstGLShader to use

func (e *GstGLFilterShader) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

func (e *GstGLFilterShader) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

// GLSL Uniforms

func (e *GstGLFilterShader) SetUniforms(value *gst.Structure) error {
	return e.Element.SetProperty("uniforms", value)
}

func (e *GstGLFilterShader) GetUniforms() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("uniforms")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Emit the 'create-shader' signal for the next frame
// Default: false
func (e *GstGLFilterShader) SetUpdateShader(value bool) error {
	return e.Element.SetProperty("update-shader", value)
}

// GLSL vertex source
// Default: NULL
func (e *GstGLFilterShader) SetVertex(value string) error {
	return e.Element.SetProperty("vertex", value)
}

func (e *GstGLFilterShader) GetVertex() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("vertex")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Infrastructure to process GL textures
type GstGLSinkBin struct {
	*gst.Bin
}

func NewGLSinkBin() (*GstGLSinkBin, error) {
	e, err := gst.NewElement("glsinkbin")
	if err != nil {
		return nil, err
	}
	return &GstGLSinkBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewGLSinkBinWithName(name string) (*GstGLSinkBin, error) {
	e, err := gst.NewElementWithName("glsinkbin", name)
	if err != nil {
		return nil, err
	}
	return &GstGLSinkBin{Bin: &gst.Bin{Element: e}}, nil
}

// Generate Quality-of-Service events upstream
// Default: false
func (e *GstGLSinkBin) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstGLSinkBin) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Additional render delay of the sink in nanoseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstGLSinkBin) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

func (e *GstGLSinkBin) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Sync on the clock
// Default: false
func (e *GstGLSinkBin) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstGLSinkBin) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Size in bytes to pull per buffer (0 = default)
// Default: 0, Min: 0, Max: -1
func (e *GstGLSinkBin) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstGLSinkBin) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable the last-sample property
// Default: false
func (e *GstGLSinkBin) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

func (e *GstGLSinkBin) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// hue
// Default: 0, Min: -1, Max: 1
func (e *GstGLSinkBin) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstGLSinkBin) GetHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)
// Default: 0, Min: -1, Max: 9223372036854775807
func (e *GstGLSinkBin) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

func (e *GstGLSinkBin) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The GL sink chain to use

func (e *GstGLSinkBin) SetSink(value *gst.Element) error {
	return e.Element.SetProperty("sink", value)
}

func (e *GstGLSinkBin) GetSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Go asynchronously to PAUSED
// Default: false
func (e *GstGLSinkBin) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

func (e *GstGLSinkBin) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The maximum bits per second to render (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstGLSinkBin) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstGLSinkBin) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: false
func (e *GstGLSinkBin) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstGLSinkBin) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The last sample received in the sink

func (e *GstGLSinkBin) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// saturation
// Default: 1, Min: 0, Max: 2
func (e *GstGLSinkBin) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstGLSinkBin) GetSaturation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The time to keep between rendered buffers (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstGLSinkBin) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

func (e *GstGLSinkBin) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstGLSinkBin) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstGLSinkBin) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// brightness
// Default: 0, Min: -1, Max: 1
func (e *GstGLSinkBin) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstGLSinkBin) GetBrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// contrast
// Default: 1, Min: 0, Max: 2
func (e *GstGLSinkBin) SetContrast(value float64) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstGLSinkBin) GetContrast() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Convert stereoscopic/multiview video formats
type GstGLViewConvertElement struct {
	*GstGLFilter
}

func NewGLViewConvertElement() (*GstGLViewConvertElement, error) {
	e, err := gst.NewElement("glviewconvert")
	if err != nil {
		return nil, err
	}
	return &GstGLViewConvertElement{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLViewConvertElementWithName(name string) (*GstGLViewConvertElement, error) {
	e, err := gst.NewElementWithName("glviewconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstGLViewConvertElement{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Override automatic output mode selection for multiview layout
// Default: none (-1)
func (e *GstGLViewConvertElement) SetOutputModeOverride(value interface{}) error {
	return e.Element.SetProperty("output-mode-override", value)
}

func (e *GstGLViewConvertElement) GetOutputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("output-mode-override")
}

// Output anaglyph type to generate when downmixing to mono
// Default: green-magenta-dubois (0)
func (e *GstGLViewConvertElement) SetDownmixMode(value GstGLStereoDownmix) error {
	e.Element.SetArg("downmix-mode", string(value))
	return nil
}

func (e *GstGLViewConvertElement) GetDownmixMode() (GstGLStereoDownmix, error) {
	var value GstGLStereoDownmix
	var ok bool
	v, err := e.Element.GetProperty("downmix-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLStereoDownmix)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLStereoDownmix")
	}
	return value, nil
}

// Override any input information about multiview layout flags
// Default: none
func (e *GstGLViewConvertElement) SetInputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("input-flags-override", value)
}

func (e *GstGLViewConvertElement) GetInputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("input-flags-override")
}

// Override any input information about multiview layout
// Default: none (-1)
func (e *GstGLViewConvertElement) SetInputModeOverride(value interface{}) error {
	return e.Element.SetProperty("input-mode-override", value)
}

func (e *GstGLViewConvertElement) GetInputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("input-mode-override")
}

// Override automatic negotiation for output multiview layout flags
// Default: none
func (e *GstGLViewConvertElement) SetOutputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("output-flags-override", value)
}

func (e *GstGLViewConvertElement) GetOutputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("output-flags-override")
}

// Downloads data from OpenGL
type GstGLDownloadElement struct {
	*GstGLBaseFilter
}

func NewGLDownloadElement() (*GstGLDownloadElement, error) {
	e, err := gst.NewElement("gldownload")
	if err != nil {
		return nil, err
	}
	return &GstGLDownloadElement{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewGLDownloadElementWithName(name string) (*GstGLDownloadElement, error) {
	e, err := gst.NewElementWithName("gldownload", name)
	if err != nil {
		return nil, err
	}
	return &GstGLDownloadElement{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// GL Shading Language effects - Sobel edge detection Effect
type GleffectsSobel struct {
	*GstGLEffects
}

func NewGleffectsSobel() (*GleffectsSobel, error) {
	e, err := gst.NewElement("gleffects_sobel")
	if err != nil {
		return nil, err
	}
	return &GleffectsSobel{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsSobelWithName(name string) (*GleffectsSobel, error) {
	e, err := gst.NewElementWithName("gleffects_sobel", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsSobel{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsSobel) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsSobel) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Invert colors to get dark edges on bright background when using sobel effect
// Default: false
func (e *GleffectsSobel) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

func (e *GleffectsSobel) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Stretch Effect
type GleffectsStretch struct {
	*GstGLEffects
}

func NewGleffectsStretch() (*GleffectsStretch, error) {
	e, err := gst.NewElement("gleffects_stretch")
	if err != nil {
		return nil, err
	}
	return &GleffectsStretch{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsStretchWithName(name string) (*GleffectsStretch, error) {
	e, err := gst.NewElementWithName("gleffects_stretch", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsStretch{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsStretch) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsStretch) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Infrastructure to process GL textures
type GstGLImageSinkBin struct {
	*GstGLSinkBin
}

func NewGLImageSinkBin() (*GstGLImageSinkBin, error) {
	e, err := gst.NewElement("glimagesink")
	if err != nil {
		return nil, err
	}
	return &GstGLImageSinkBin{GstGLSinkBin: &GstGLSinkBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewGLImageSinkBinWithName(name string) (*GstGLImageSinkBin, error) {
	e, err := gst.NewElementWithName("glimagesink", name)
	if err != nil {
		return nil, err
	}
	return &GstGLImageSinkBin{GstGLSinkBin: &GstGLSinkBin{Bin: &gst.Bin{Element: e}}}, nil
}

// When enabled, alpha will be ignored and converted to black
// Default: true
func (e *GstGLImageSinkBin) SetIgnoreAlpha(value bool) error {
	return e.Element.SetProperty("ignore-alpha", value)
}

func (e *GstGLImageSinkBin) GetIgnoreAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Output multiview layout modifier flags
// Default: none
func (e *GstGLImageSinkBin) SetOutputMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("output-multiview-flags", value)
}

func (e *GstGLImageSinkBin) GetOutputMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-flags")
}

// Choose output mode for multiview/3D video
// Default: mono (0)
func (e *GstGLImageSinkBin) SetOutputMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("output-multiview-mode", value)
}

func (e *GstGLImageSinkBin) GetOutputMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-mode")
}

// The render rectangle ('<x, y, width, height>')

func (e *GstGLImageSinkBin) SetRenderRectangle(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// rotate method
// Default: none (0)
func (e *GstGLImageSinkBin) SetRotateMethod(value GstGLRotateMethod) error {
	e.Element.SetArg("rotate-method", string(value))
	return nil
}

func (e *GstGLImageSinkBin) GetRotateMethod() (GstGLRotateMethod, error) {
	var value GstGLRotateMethod
	var ok bool
	v, err := e.Element.GetProperty("rotate-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLRotateMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLRotateMethod")
	}
	return value, nil
}

// Get OpenGL context

func (e *GstGLImageSinkBin) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstGLImageSinkBin) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstGLImageSinkBin) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, XEvents will be selected and handled
// Default: true
func (e *GstGLImageSinkBin) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

func (e *GstGLImageSinkBin) GetHandleEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Output anaglyph type to generate when downmixing to mono
// Default: green-magenta-dubois (0)
func (e *GstGLImageSinkBin) SetOutputMultiviewDownmixMode(value GstGLStereoDownmix) error {
	e.Element.SetArg("output-multiview-downmix-mode", string(value))
	return nil
}

func (e *GstGLImageSinkBin) GetOutputMultiviewDownmixMode() (GstGLStereoDownmix, error) {
	var value GstGLStereoDownmix
	var ok bool
	v, err := e.Element.GetProperty("output-multiview-downmix-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLStereoDownmix)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLStereoDownmix")
	}
	return value, nil
}

// The pixel aspect ratio of the device
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstGLImageSinkBin) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstGLImageSinkBin) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

// Whether to render video frames during preroll
// Default: true
func (e *GstGLImageSinkBin) SetShowPrerollFrame(value bool) error {
	return e.Element.SetProperty("show-preroll-frame", value)
}

func (e *GstGLImageSinkBin) GetShowPrerollFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-preroll-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Flatten a stream containing GstVideoOverlayCompositionMeta
type GstGLOverlayCompositorElement struct {
	*GstGLFilter
}

func NewGLOverlayCompositorElement() (*GstGLOverlayCompositorElement, error) {
	e, err := gst.NewElement("gloverlaycompositor")
	if err != nil {
		return nil, err
	}
	return &GstGLOverlayCompositorElement{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLOverlayCompositorElementWithName(name string) (*GstGLOverlayCompositorElement, error) {
	e, err := gst.NewElementWithName("gloverlaycompositor", name)
	if err != nil {
		return nil, err
	}
	return &GstGLOverlayCompositorElement{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// OpenGL stereo video combiner
type GstGLStereoMix struct {
	*GstGLMixer
}

func NewGLStereoMix() (*GstGLStereoMix, error) {
	e, err := gst.NewElement("glstereomix")
	if err != nil {
		return nil, err
	}
	return &GstGLStereoMix{GstGLMixer: &GstGLMixer{GstGLBaseMixer: &GstGLBaseMixer{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}}}, nil
}

func NewGLStereoMixWithName(name string) (*GstGLStereoMix, error) {
	e, err := gst.NewElementWithName("glstereomix", name)
	if err != nil {
		return nil, err
	}
	return &GstGLStereoMix{GstGLMixer: &GstGLMixer{GstGLBaseMixer: &GstGLBaseMixer{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}}}, nil
}

// Output anaglyph type to generate when downmixing to mono
// Default: green-magenta-dubois (0)
func (e *GstGLStereoMix) SetDownmixMode(value GstGLStereoDownmix) error {
	e.Element.SetArg("downmix-mode", string(value))
	return nil
}

func (e *GstGLStereoMix) GetDownmixMode() (GstGLStereoDownmix, error) {
	var value GstGLStereoDownmix
	var ok bool
	v, err := e.Element.GetProperty("downmix-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLStereoDownmix)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLStereoDownmix")
	}
	return value, nil
}

// GL Shading Language effects - FishEye Effect
type GleffectsFisheye struct {
	*GstGLEffects
}

func NewGleffectsFisheye() (*GleffectsFisheye, error) {
	e, err := gst.NewElement("gleffects_fisheye")
	if err != nil {
		return nil, err
	}
	return &GleffectsFisheye{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsFisheyeWithName(name string) (*GleffectsFisheye, error) {
	e, err := gst.NewElementWithName("gleffects_fisheye", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsFisheye{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsFisheye) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsFisheye) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Glow Lighting Effect
type GleffectsGlow struct {
	*GstGLEffects
}

func NewGleffectsGlow() (*GleffectsGlow, error) {
	e, err := gst.NewElement("gleffects_glow")
	if err != nil {
		return nil, err
	}
	return &GleffectsGlow{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsGlowWithName(name string) (*GleffectsGlow, error) {
	e, err := gst.NewElementWithName("gleffects_glow", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsGlow{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsGlow) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsGlow) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Colorspace converter and video scaler
type GstGLColorscale struct {
	*GstGLFilter
}

func NewGLColorscale() (*GstGLColorscale, error) {
	e, err := gst.NewElement("glcolorscale")
	if err != nil {
		return nil, err
	}
	return &GstGLColorscale{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLColorscaleWithName(name string) (*GstGLColorscale, error) {
	e, err := gst.NewElementWithName("glcolorscale", name)
	if err != nil {
		return nil, err
	}
	return &GstGLColorscale{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// GL Shading Language effects - Bulge Effect
type GleffectsBulge struct {
	*GstGLEffects
}

func NewGleffectsBulge() (*GleffectsBulge, error) {
	e, err := gst.NewElement("gleffects_bulge")
	if err != nil {
		return nil, err
	}
	return &GleffectsBulge{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsBulgeWithName(name string) (*GleffectsBulge, error) {
	e, err := gst.NewElementWithName("gleffects_bulge", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsBulge{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsBulge) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsBulge) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Do nothing Effect
type GleffectsIdentity struct {
	*GstGLEffects
}

func NewGleffectsIdentity() (*GleffectsIdentity, error) {
	e, err := gst.NewElement("gleffects_identity")
	if err != nil {
		return nil, err
	}
	return &GleffectsIdentity{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsIdentityWithName(name string) (*GleffectsIdentity, error) {
	e, err := gst.NewElementWithName("gleffects_identity", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsIdentity{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsIdentity) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsIdentity) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Overlay GL video texture with a JPEG/PNG image
type GstGLOverlay struct {
	*GstGLFilter
}

func NewGLOverlay() (*GstGLOverlay, error) {
	e, err := gst.NewElement("gloverlay")
	if err != nil {
		return nil, err
	}
	return &GstGLOverlay{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLOverlayWithName(name string) (*GstGLOverlay, error) {
	e, err := gst.NewElementWithName("gloverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstGLOverlay{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Vertical offset of overlay image in fractions of video image height, from top-left corner of video image
// Default: 0, Min: 0, Max: 1
func (e *GstGLOverlay) SetRelativeY(value float64) error {
	return e.Element.SetProperty("relative-y", value)
}

func (e *GstGLOverlay) GetRelativeY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("relative-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Global alpha of overlay image
// Default: 1, Min: 0, Max: 1
func (e *GstGLOverlay) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *GstGLOverlay) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Location of image file to overlay
// Default: NULL
func (e *GstGLOverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstGLOverlay) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// For positive value, horizontal offset of overlay image in pixels from left of video image. For negative value, horizontal offset of overlay image in pixels from right of video image
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstGLOverlay) SetOffsetX(value int) error {
	return e.Element.SetProperty("offset-x", value)
}

func (e *GstGLOverlay) GetOffsetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// For positive value, vertical offset of overlay image in pixels from top of video image. For negative value, vertical offset of overlay image in pixels from bottom of video image
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstGLOverlay) SetOffsetY(value int) error {
	return e.Element.SetProperty("offset-y", value)
}

func (e *GstGLOverlay) GetOffsetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Height of overlay image in pixels (0 = same as overlay image)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGLOverlay) SetOverlayHeight(value int) error {
	return e.Element.SetProperty("overlay-height", value)
}

func (e *GstGLOverlay) GetOverlayHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Width of overlay image in pixels (0 = same as overlay image)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGLOverlay) SetOverlayWidth(value int) error {
	return e.Element.SetProperty("overlay-width", value)
}

func (e *GstGLOverlay) GetOverlayWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Horizontal offset of overlay image in fractions of video image width, from top-left corner of video image
// Default: 0, Min: 0, Max: 1
func (e *GstGLOverlay) SetRelativeX(value float64) error {
	return e.Element.SetProperty("relative-x", value)
}

func (e *GstGLOverlay) GetRelativeX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("relative-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Creates a test video stream
type GstGLTestSrc struct {
	*GstGLBaseSrc
}

func NewGLTestSrc() (*GstGLTestSrc, error) {
	e, err := gst.NewElement("gltestsrc")
	if err != nil {
		return nil, err
	}
	return &GstGLTestSrc{GstGLBaseSrc: &GstGLBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewGLTestSrcWithName(name string) (*GstGLTestSrc, error) {
	e, err := gst.NewElementWithName("gltestsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstGLTestSrc{GstGLBaseSrc: &GstGLBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// Whether to act as a live source
// Default: false
func (e *GstGLTestSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstGLTestSrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Type of test pattern to generate
// Default: smpte (0)
func (e *GstGLTestSrc) SetPattern(value GstGLTestSrcPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

func (e *GstGLTestSrc) GetPattern() (GstGLTestSrcPattern, error) {
	var value GstGLTestSrcPattern
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLTestSrcPattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLTestSrcPattern")
	}
	return value, nil
}

// OpenGL video_mixer bin
type GstGLVideoMixerBin struct {
	*GstGLMixerBin
}

func NewGLVideoMixerBin() (*GstGLVideoMixerBin, error) {
	e, err := gst.NewElement("glvideomixer")
	if err != nil {
		return nil, err
	}
	return &GstGLVideoMixerBin{GstGLMixerBin: &GstGLMixerBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewGLVideoMixerBinWithName(name string) (*GstGLVideoMixerBin, error) {
	e, err := gst.NewElementWithName("glvideomixer", name)
	if err != nil {
		return nil, err
	}
	return &GstGLVideoMixerBin{GstGLMixerBin: &GstGLMixerBin{Bin: &gst.Bin{Element: e}}}, nil
}

// Background type
// Default: checker (0)
func (e *GstGLVideoMixerBin) SetBackground(value GstGLVideoMixerBackground) error {
	e.Element.SetArg("background", string(value))
	return nil
}

func (e *GstGLVideoMixerBin) GetBackground() (GstGLVideoMixerBackground, error) {
	var value GstGLVideoMixerBackground
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLVideoMixerBackground)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLVideoMixerBackground")
	}
	return value, nil
}

// Adjusts brightness, contrast, hue, saturation on a video stream
type GstGLColorBalance struct {
	*GstGLFilter
}

func NewGLColorBalance() (*GstGLColorBalance, error) {
	e, err := gst.NewElement("glcolorbalance")
	if err != nil {
		return nil, err
	}
	return &GstGLColorBalance{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLColorBalanceWithName(name string) (*GstGLColorBalance, error) {
	e, err := gst.NewElementWithName("glcolorbalance", name)
	if err != nil {
		return nil, err
	}
	return &GstGLColorBalance{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// brightness
// Default: 0, Min: -1, Max: 1
func (e *GstGLColorBalance) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstGLColorBalance) GetBrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// contrast
// Default: 1, Min: 0, Max: 2
func (e *GstGLColorBalance) SetContrast(value float64) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstGLColorBalance) GetContrast() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// hue
// Default: 0, Min: -1, Max: 1
func (e *GstGLColorBalance) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstGLColorBalance) GetHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// saturation
// Default: 1, Min: 0, Max: 2
func (e *GstGLColorBalance) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstGLColorBalance) GetSaturation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Converts between color spaces using OpenGL shaders
type GstGLColorConvertElement struct {
	*GstGLBaseFilter
}

func NewGLColorConvertElement() (*GstGLColorConvertElement, error) {
	e, err := gst.NewElement("glcolorconvert")
	if err != nil {
		return nil, err
	}
	return &GstGLColorConvertElement{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewGLColorConvertElementWithName(name string) (*GstGLColorConvertElement, error) {
	e, err := gst.NewElementWithName("glcolorconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstGLColorConvertElement{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// GL Shading Language effects - Mirror Effect
type GleffectsMirror struct {
	*GstGLEffects
}

func NewGleffectsMirror() (*GleffectsMirror, error) {
	e, err := gst.NewElement("gleffects_mirror")
	if err != nil {
		return nil, err
	}
	return &GleffectsMirror{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsMirrorWithName(name string) (*GleffectsMirror, error) {
	e, err := gst.NewElementWithName("gleffects_mirror", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsMirror{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsMirror) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsMirror) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Infrastructure to process GL textures
type GstGLFilterBin struct {
	*gst.Bin
}

func NewGLFilterBin() (*GstGLFilterBin, error) {
	e, err := gst.NewElement("glfilterbin")
	if err != nil {
		return nil, err
	}
	return &GstGLFilterBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewGLFilterBinWithName(name string) (*GstGLFilterBin, error) {
	e, err := gst.NewElementWithName("glfilterbin", name)
	if err != nil {
		return nil, err
	}
	return &GstGLFilterBin{Bin: &gst.Bin{Element: e}}, nil
}

// The GL filter chain to use

func (e *GstGLFilterBin) SetFilter(value *gst.Element) error {
	return e.Element.SetProperty("filter", value)
}

func (e *GstGLFilterBin) GetFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// OpenGL video_mixer empty bin
type GstGLMixerBin struct {
	*gst.Bin
}

func NewGLMixerBinWithProperties(properties map[string]interface{}) (*GstGLMixerBin, error) {
	e, err := gst.NewElementWithProperties("glmixerbin", properties)
	if err != nil {
		return nil, err
	}
	return &GstGLMixerBin{Bin: &gst.Bin{Element: e}}, nil
}

// Start time to use if start-time-selection=set
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstGLMixerBin) SetStartTime(value uint64) error {
	return e.Element.SetProperty("start-time", value)
}

func (e *GstGLMixerBin) GetStartTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("start-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Decides which start time is output
// Default: zero (0)
func (e *GstGLMixerBin) SetStartTimeSelection(value GstGLMixerBinStartTimeSelection) error {
	e.Element.SetArg("start-time-selection", string(value))
	return nil
}

func (e *GstGLMixerBin) GetStartTimeSelection() (GstGLMixerBinStartTimeSelection, error) {
	var value GstGLMixerBinStartTimeSelection
	var ok bool
	v, err := e.Element.GetProperty("start-time-selection")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLMixerBinStartTimeSelection)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLMixerBinStartTimeSelection")
	}
	return value, nil
}

// Get OpenGL context

func (e *GstGLMixerBin) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// Always operate in live mode and aggregate on timeout regardless of whether any live sources are linked upstream
// Default: false
func (e *GstGLMixerBin) SetForceLive(value bool) error {
	return e.Element.SetProperty("force-live", value)
}

func (e *GstGLMixerBin) GetForceLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Additional latency in live mode to allow upstream to take longer to produce buffers for the current position (in nanoseconds)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstGLMixerBin) SetLatency(value uint64) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstGLMixerBin) GetLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// When sources with a higher latency are expected to be plugged in dynamically after the aggregator has started playing, this allows overriding the minimum latency reported by the initial source(s). This is only taken into account when larger than the actually reported minimum latency. (nanoseconds)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstGLMixerBin) SetMinUpstreamLatency(value uint64) error {
	return e.Element.SetProperty("min-upstream-latency", value)
}

func (e *GstGLMixerBin) GetMinUpstreamLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-upstream-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The GL mixer chain to use

func (e *GstGLMixerBin) SetMixer(value *gst.Element) error {
	return e.Element.SetProperty("mixer", value)
}

func (e *GstGLMixerBin) GetMixer() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("mixer")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// GL Shading Language effects - Heat Signature Effect
type GleffectsHeat struct {
	*GstGLEffects
}

func NewGleffectsHeat() (*GleffectsHeat, error) {
	e, err := gst.NewElement("gleffects_heat")
	if err != nil {
		return nil, err
	}
	return &GleffectsHeat{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsHeatWithName(name string) (*GleffectsHeat, error) {
	e, err := gst.NewElementWithName("gleffects_heat", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsHeat{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsHeat) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsHeat) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Luma Cross Processing Effect
type GleffectsLumaxpro struct {
	*GstGLEffects
}

func NewGleffectsLumaxpro() (*GleffectsLumaxpro, error) {
	e, err := gst.NewElement("gleffects_lumaxpro")
	if err != nil {
		return nil, err
	}
	return &GleffectsLumaxpro{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsLumaxproWithName(name string) (*GleffectsLumaxpro, error) {
	e, err := gst.NewElementWithName("gleffects_lumaxpro", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsLumaxpro{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsLumaxpro) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsLumaxpro) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Laplacian Convolution Demo Effect
type GleffectsLaplacian struct {
	*GstGLEffects
}

func NewGleffectsLaplacian() (*GleffectsLaplacian, error) {
	e, err := gst.NewElement("gleffects_laplacian")
	if err != nil {
		return nil, err
	}
	return &GleffectsLaplacian{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsLaplacianWithName(name string) (*GleffectsLaplacian, error) {
	e, err := gst.NewElementWithName("gleffects_laplacian", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsLaplacian{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsLaplacian) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsLaplacian) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Invert colors to get dark edges on bright background when using sobel effect
// Default: false
func (e *GleffectsLaplacian) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

func (e *GleffectsLaplacian) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - All Grey but Red Effect
type GleffectsSin struct {
	*GstGLEffects
}

func NewGleffectsSin() (*GleffectsSin, error) {
	e, err := gst.NewElement("gleffects_sin")
	if err != nil {
		return nil, err
	}
	return &GleffectsSin{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsSinWithName(name string) (*GleffectsSin, error) {
	e, err := gst.NewElementWithName("gleffects_sin", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsSin{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsSin) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsSin) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Light Tunnel Effect
type GleffectsTunnel struct {
	*GstGLEffects
}

func NewGleffectsTunnel() (*GleffectsTunnel, error) {
	e, err := gst.NewElement("gleffects_tunnel")
	if err != nil {
		return nil, err
	}
	return &GleffectsTunnel{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsTunnelWithName(name string) (*GleffectsTunnel, error) {
	e, err := gst.NewElementWithName("gleffects_tunnel", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsTunnel{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsTunnel) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsTunnel) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// A videosink based on OpenGL
type GstGLImageSink struct {
	*GstVideoSink
}

func NewGLImageSink() (*GstGLImageSink, error) {
	e, err := gst.NewElement("glimagesinkelement")
	if err != nil {
		return nil, err
	}
	return &GstGLImageSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewGLImageSinkWithName(name string) (*GstGLImageSink, error) {
	e, err := gst.NewElementWithName("glimagesinkelement", name)
	if err != nil {
		return nil, err
	}
	return &GstGLImageSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// The pixel aspect ratio of the device
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstGLImageSink) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstGLImageSink) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

// rotate method
// Default: none (0)
func (e *GstGLImageSink) SetRotateMethod(value GstGLRotateMethod) error {
	e.Element.SetArg("rotate-method", string(value))
	return nil
}

func (e *GstGLImageSink) GetRotateMethod() (GstGLRotateMethod, error) {
	var value GstGLRotateMethod
	var ok bool
	v, err := e.Element.GetProperty("rotate-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLRotateMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLRotateMethod")
	}
	return value, nil
}

// Get OpenGL context

func (e *GstGLImageSink) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstGLImageSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstGLImageSink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Output anaglyph type to generate when downmixing to mono
// Default: green-magenta-dubois (0)
func (e *GstGLImageSink) SetOutputMultiviewDownmixMode(value GstGLStereoDownmix) error {
	e.Element.SetArg("output-multiview-downmix-mode", string(value))
	return nil
}

func (e *GstGLImageSink) GetOutputMultiviewDownmixMode() (GstGLStereoDownmix, error) {
	var value GstGLStereoDownmix
	var ok bool
	v, err := e.Element.GetProperty("output-multiview-downmix-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLStereoDownmix)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLStereoDownmix")
	}
	return value, nil
}

// Choose output mode for multiview/3D video
// Default: mono (0)
func (e *GstGLImageSink) SetOutputMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("output-multiview-mode", value)
}

func (e *GstGLImageSink) GetOutputMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-mode")
}

// The render rectangle ('<x, y, width, height>')

func (e *GstGLImageSink) SetRenderRectangle(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// When enabled, XEvents will be selected and handled
// Default: true
func (e *GstGLImageSink) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

func (e *GstGLImageSink) GetHandleEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, alpha will be ignored and converted to black
// Default: true
func (e *GstGLImageSink) SetIgnoreAlpha(value bool) error {
	return e.Element.SetProperty("ignore-alpha", value)
}

func (e *GstGLImageSink) GetIgnoreAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Output multiview layout modifier flags
// Default: none
func (e *GstGLImageSink) SetOutputMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("output-multiview-flags", value)
}

func (e *GstGLImageSink) GetOutputMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("output-multiview-flags")
}

// Infrastructure to process GL textures
type GstGLSrcBin struct {
	*gst.Bin
}

func NewGLSrcBin() (*GstGLSrcBin, error) {
	e, err := gst.NewElement("glsrcbin")
	if err != nil {
		return nil, err
	}
	return &GstGLSrcBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewGLSrcBinWithName(name string) (*GstGLSrcBin, error) {
	e, err := gst.NewElementWithName("glsrcbin", name)
	if err != nil {
		return nil, err
	}
	return &GstGLSrcBin{Bin: &gst.Bin{Element: e}}, nil
}

// The GL src chain to use

func (e *GstGLSrcBin) SetSrc(value *gst.Element) error {
	return e.Element.SetProperty("src", value)
}

func (e *GstGLSrcBin) GetSrc() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("src")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Flip video on the GPU
type GstGLVideoFlip struct {
	*gst.Bin
}

func NewGLVideoFlip() (*GstGLVideoFlip, error) {
	e, err := gst.NewElement("glvideoflip")
	if err != nil {
		return nil, err
	}
	return &GstGLVideoFlip{Bin: &gst.Bin{Element: e}}, nil
}

func NewGLVideoFlipWithName(name string) (*GstGLVideoFlip, error) {
	e, err := gst.NewElementWithName("glvideoflip", name)
	if err != nil {
		return nil, err
	}
	return &GstGLVideoFlip{Bin: &gst.Bin{Element: e}}, nil
}

// method (deprecated, use video-direction instead)
// Default: none (0)
func (e *GstGLVideoFlip) SetMethod(value GstGLVideoFlipMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstGLVideoFlip) GetMethod() (GstGLVideoFlipMethod, error) {
	var value GstGLVideoFlipMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLVideoFlipMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLVideoFlipMethod")
	}
	return value, nil
}

// Deinterlacing based on fragment shaders
type GstGLDeinterlace struct {
	*GstGLFilter
}

func NewGLDeinterlace() (*GstGLDeinterlace, error) {
	e, err := gst.NewElement("gldeinterlace")
	if err != nil {
		return nil, err
	}
	return &GstGLDeinterlace{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLDeinterlaceWithName(name string) (*GstGLDeinterlace, error) {
	e, err := gst.NewElementWithName("gldeinterlace", name)
	if err != nil {
		return nil, err
	}
	return &GstGLDeinterlace{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Select which deinterlace method apply to GL video texture
// Default: vfir (0)
func (e *GstGLDeinterlace) SetMethod(value GstGLDeinterlaceMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstGLDeinterlace) GetMethod() (GstGLDeinterlaceMethod, error) {
	var value GstGLDeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLDeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLDeinterlaceMethod")
	}
	return value, nil
}

// Saves a background frame and replace it with a pixbuf
type GstGLDifferenceMatte struct {
	*GstGLFilter
}

func NewGLDifferenceMatte() (*GstGLDifferenceMatte, error) {
	e, err := gst.NewElement("gldifferencematte")
	if err != nil {
		return nil, err
	}
	return &GstGLDifferenceMatte{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLDifferenceMatteWithName(name string) (*GstGLDifferenceMatte, error) {
	e, err := gst.NewElementWithName("gldifferencematte", name)
	if err != nil {
		return nil, err
	}
	return &GstGLDifferenceMatte{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Background image location
// Default: NULL
func (e *GstGLDifferenceMatte) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstGLDifferenceMatte) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// OpenGL mosaic
type GstGLMosaic struct {
	*GstGLMixer
}

func NewGLMosaic() (*GstGLMosaic, error) {
	e, err := gst.NewElement("glmosaic")
	if err != nil {
		return nil, err
	}
	return &GstGLMosaic{GstGLMixer: &GstGLMixer{GstGLBaseMixer: &GstGLBaseMixer{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}}}, nil
}

func NewGLMosaicWithName(name string) (*GstGLMosaic, error) {
	e, err := gst.NewElementWithName("glmosaic", name)
	if err != nil {
		return nil, err
	}
	return &GstGLMosaic{GstGLMixer: &GstGLMixer{GstGLBaseMixer: &GstGLBaseMixer{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}}}, nil
}

// Use client callbacks to define the scene
type GstGLFilterApp struct {
	*GstGLFilter
}

func NewGLFilterApp() (*GstGLFilterApp, error) {
	e, err := gst.NewElement("glfilterapp")
	if err != nil {
		return nil, err
	}
	return &GstGLFilterApp{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLFilterAppWithName(name string) (*GstGLFilterApp, error) {
	e, err := gst.NewElementWithName("glfilterapp", name)
	if err != nil {
		return nil, err
	}
	return &GstGLFilterApp{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Map input texture on the 6 cube faces
type GstGLFilterCube struct {
	*GstGLFilter
}

func NewGLFilterCube() (*GstGLFilterCube, error) {
	e, err := gst.NewElement("glfiltercube")
	if err != nil {
		return nil, err
	}
	return &GstGLFilterCube{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLFilterCubeWithName(name string) (*GstGLFilterCube, error) {
	e, err := gst.NewElementWithName("glfiltercube", name)
	if err != nil {
		return nil, err
	}
	return &GstGLFilterCube{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Specifies the distance from the viewer to the far clipping plane
// Default: 100, Min: 0, Max: 1000
func (e *GstGLFilterCube) SetZfar(value float64) error {
	return e.Element.SetProperty("zfar", value)
}

func (e *GstGLFilterCube) GetZfar() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zfar")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Specifies the distance from the viewer to the near clipping plane
// Default: 0.1, Min: 0, Max: 100
func (e *GstGLFilterCube) SetZnear(value float64) error {
	return e.Element.SetProperty("znear", value)
}

func (e *GstGLFilterCube) GetZnear() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("znear")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Field of view in the x direction
// Default: 0, Min: 0, Max: 100
func (e *GstGLFilterCube) SetAspect(value float64) error {
	return e.Element.SetProperty("aspect", value)
}

func (e *GstGLFilterCube) GetAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Background blue color
// Default: 0, Min: 0, Max: 1
func (e *GstGLFilterCube) SetBlue(value float32) error {
	return e.Element.SetProperty("blue", value)
}

func (e *GstGLFilterCube) GetBlue() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Field of view angle in degrees
// Default: 45, Min: 0, Max: 180
func (e *GstGLFilterCube) SetFovy(value float64) error {
	return e.Element.SetProperty("fovy", value)
}

func (e *GstGLFilterCube) GetFovy() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("fovy")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Background green color
// Default: 0, Min: 0, Max: 1
func (e *GstGLFilterCube) SetGreen(value float32) error {
	return e.Element.SetProperty("green", value)
}

func (e *GstGLFilterCube) GetGreen() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("green")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Background red color
// Default: 0, Min: 0, Max: 1
func (e *GstGLFilterCube) SetRed(value float32) error {
	return e.Element.SetProperty("red", value)
}

func (e *GstGLFilterCube) GetRed() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("red")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// GL Shading Language effects - Blur with 9x9 separable convolution Effect
type GleffectsBlur struct {
	*GstGLEffects
}

func NewGleffectsBlur() (*GleffectsBlur, error) {
	e, err := gst.NewElement("gleffects_blur")
	if err != nil {
		return nil, err
	}
	return &GleffectsBlur{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsBlurWithName(name string) (*GleffectsBlur, error) {
	e, err := gst.NewElementWithName("gleffects_blur", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsBlur{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsBlur) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsBlur) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Twirl Effect
type GleffectsTwirl struct {
	*GstGLEffects
}

func NewGleffectsTwirl() (*GleffectsTwirl, error) {
	e, err := gst.NewElement("gleffects_twirl")
	if err != nil {
		return nil, err
	}
	return &GleffectsTwirl{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsTwirlWithName(name string) (*GleffectsTwirl, error) {
	e, err := gst.NewElementWithName("gleffects_twirl", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsTwirl{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsTwirl) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsTwirl) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GL Shading Language effects - Cross Processing Effect
type GleffectsXpro struct {
	*GstGLEffects
}

func NewGleffectsXpro() (*GleffectsXpro, error) {
	e, err := gst.NewElement("gleffects_xpro")
	if err != nil {
		return nil, err
	}
	return &GleffectsXpro{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGleffectsXproWithName(name string) (*GleffectsXpro, error) {
	e, err := gst.NewElementWithName("gleffects_xpro", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsXpro{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GleffectsXpro) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GleffectsXpro) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Glass Filter
type GstGLFilterGlass struct {
	*GstGLFilter
}

func NewGLFilterGlass() (*GstGLFilterGlass, error) {
	e, err := gst.NewElement("glfilterglass")
	if err != nil {
		return nil, err
	}
	return &GstGLFilterGlass{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLFilterGlassWithName(name string) (*GstGLFilterGlass, error) {
	e, err := gst.NewElementWithName("glfilterglass", name)
	if err != nil {
		return nil, err
	}
	return &GstGLFilterGlass{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Transform video on the GPU
type GstGLTransformation struct {
	*GstGLFilter
}

func NewGLTransformation() (*GstGLTransformation, error) {
	e, err := gst.NewElement("gltransformation")
	if err != nil {
		return nil, err
	}
	return &GstGLTransformation{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLTransformationWithName(name string) (*GstGLTransformation, error) {
	e, err := gst.NewElementWithName("gltransformation", name)
	if err != nil {
		return nil, err
	}
	return &GstGLTransformation{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Field of view angle in degrees
// Default: 90, Min: 0, Max: 3.40282e+38
func (e *GstGLTransformation) SetFov(value float32) error {
	return e.Element.SetProperty("fov", value)
}

func (e *GstGLTransformation) GetFov() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fov")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Rotates the video around the X-Axis in degrees.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetRotationX(value float32) error {
	return e.Element.SetProperty("rotation-x", value)
}

func (e *GstGLTransformation) GetRotationX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rotation-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Rotates the video around the Z-Axis in degrees.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetRotationZ(value float32) error {
	return e.Element.SetProperty("rotation-z", value)
}

func (e *GstGLTransformation) GetRotationZ() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rotation-z")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Scale multiplier for the X-Axis.
// Default: 1, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetScaleX(value float32) error {
	return e.Element.SetProperty("scale-x", value)
}

func (e *GstGLTransformation) GetScaleX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scale-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Translates the video at the Y-Axis, in universal [0-1] coordinate.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetTranslationY(value float32) error {
	return e.Element.SetProperty("translation-y", value)
}

func (e *GstGLTransformation) GetTranslationY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("translation-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Rotates the video around the Y-Axis in degrees.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetRotationY(value float32) error {
	return e.Element.SetProperty("rotation-y", value)
}

func (e *GstGLTransformation) GetRotationY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rotation-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Scale multiplier for the Y-Axis.
// Default: 1, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetScaleY(value float32) error {
	return e.Element.SetProperty("scale-y", value)
}

func (e *GstGLTransformation) GetScaleY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scale-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Translates the video at the X-Axis, in universal [0-1] coordinate.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetTranslationX(value float32) error {
	return e.Element.SetProperty("translation-x", value)
}

func (e *GstGLTransformation) GetTranslationX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("translation-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The final Graphene 4x4 Matrix for transformation

func (e *GstGLTransformation) SetMvpMatrix(value interface{}) error {
	return e.Element.SetProperty("mvp-matrix", value)
}

func (e *GstGLTransformation) GetMvpMatrix() (interface{}, error) {
	return e.Element.GetProperty("mvp-matrix")
}

// Use orthographic projection
// Default: false
func (e *GstGLTransformation) SetOrtho(value bool) error {
	return e.Element.SetProperty("ortho", value)
}

func (e *GstGLTransformation) GetOrtho() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ortho")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Rotation pivot point X coordinate, where 0 is the center, -1 the left border, +1 the right border and <-1, >1 outside.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetPivotX(value float32) error {
	return e.Element.SetProperty("pivot-x", value)
}

func (e *GstGLTransformation) GetPivotX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pivot-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Rotation pivot point X coordinate, where 0 is the center, -1 the left border, +1 the right border and <-1, >1 outside.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetPivotY(value float32) error {
	return e.Element.SetProperty("pivot-y", value)
}

func (e *GstGLTransformation) GetPivotY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pivot-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Relevant for rotation in 3D space. You look into the negative Z axis direction
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetPivotZ(value float32) error {
	return e.Element.SetProperty("pivot-z", value)
}

func (e *GstGLTransformation) GetPivotZ() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pivot-z")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Translates the video at the Z-Axis, in universal [0-1] coordinate.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstGLTransformation) SetTranslationZ(value float32) error {
	return e.Element.SetProperty("translation-z", value)
}

func (e *GstGLTransformation) GetTranslationZ() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("translation-z")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// OpenGL video_mixer
type GstGLVideoMixer struct {
	*GstGLMixer
}

func NewGLVideoMixer() (*GstGLVideoMixer, error) {
	e, err := gst.NewElement("glvideomixerelement")
	if err != nil {
		return nil, err
	}
	return &GstGLVideoMixer{GstGLMixer: &GstGLMixer{GstGLBaseMixer: &GstGLBaseMixer{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}}}, nil
}

func NewGLVideoMixerWithName(name string) (*GstGLVideoMixer, error) {
	e, err := gst.NewElementWithName("glvideomixerelement", name)
	if err != nil {
		return nil, err
	}
	return &GstGLVideoMixer{GstGLMixer: &GstGLMixer{GstGLBaseMixer: &GstGLBaseMixer{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}}}, nil
}

// Background type
// Default: checker (0)
func (e *GstGLVideoMixer) SetBackground(value GstGLVideoMixerBackground) error {
	e.Element.SetArg("background", string(value))
	return nil
}

func (e *GstGLVideoMixer) GetBackground() (GstGLVideoMixerBackground, error) {
	var value GstGLVideoMixerBackground
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLVideoMixerBackground)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLVideoMixerBackground")
	}
	return value, nil
}

// Adds an alpha channel to video using OpenGL - uniform or chroma-keying
type GstGLAlpha struct {
	*GstGLFilter
}

func NewGLAlpha() (*GstGLAlpha, error) {
	e, err := gst.NewElement("glalpha")
	if err != nil {
		return nil, err
	}
	return &GstGLAlpha{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGLAlphaWithName(name string) (*GstGLAlpha, error) {
	e, err := gst.NewElementWithName("glalpha", name)
	if err != nil {
		return nil, err
	}
	return &GstGLAlpha{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Size of the colorcube to change
// Default: 20, Min: 0, Max: 90
func (e *GstGLAlpha) SetAngle(value float32) error {
	return e.Element.SetProperty("angle", value)
}

func (e *GstGLAlpha) GetAngle() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// How the alpha channels should be created
// Default: set (0)
func (e *GstGLAlpha) SetMethod(value GstGLAlphaMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstGLAlpha) GetMethod() (GstGLAlphaMethod, error) {
	var value GstGLAlphaMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLAlphaMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLAlphaMethod")
	}
	return value, nil
}

// The blue color value for custom RGB chroma keying
// Default: 0, Min: 0, Max: 255
func (e *GstGLAlpha) SetTargetB(value uint) error {
	return e.Element.SetProperty("target-b", value)
}

func (e *GstGLAlpha) GetTargetB() (uint, error) {
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

// The green color value for custom RGB chroma keying
// Default: 255, Min: 0, Max: 255
func (e *GstGLAlpha) SetTargetG(value uint) error {
	return e.Element.SetProperty("target-g", value)
}

func (e *GstGLAlpha) GetTargetG() (uint, error) {
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

// Sensitivity to bright colors
// Default: 100, Min: 0, Max: 128
func (e *GstGLAlpha) SetWhiteSensitivity(value uint) error {
	return e.Element.SetProperty("white-sensitivity", value)
}

func (e *GstGLAlpha) GetWhiteSensitivity() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("white-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The value for the alpha channel
// Default: 1, Min: 0, Max: 1
func (e *GstGLAlpha) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *GstGLAlpha) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Sensitivity to dark colors
// Default: 100, Min: 0, Max: 128
func (e *GstGLAlpha) SetBlackSensitivity(value uint) error {
	return e.Element.SetProperty("black-sensitivity", value)
}

func (e *GstGLAlpha) GetBlackSensitivity() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("black-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Size of noise radius
// Default: 2, Min: 0, Max: 64
func (e *GstGLAlpha) SetNoiseLevel(value float32) error {
	return e.Element.SetProperty("noise-level", value)
}

func (e *GstGLAlpha) GetNoiseLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("noise-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The red color value for custom RGB chroma keying
// Default: 0, Min: 0, Max: 255
func (e *GstGLAlpha) SetTargetR(value uint) error {
	return e.Element.SetProperty("target-r", value)
}

func (e *GstGLAlpha) GetTargetR() (uint, error) {
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

// GL Shading Language effects
type GstGLEffectsGeneric struct {
	*GstGLEffects
}

func NewGLEffectsGeneric() (*GstGLEffectsGeneric, error) {
	e, err := gst.NewElement("gleffects")
	if err != nil {
		return nil, err
	}
	return &GstGLEffectsGeneric{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewGLEffectsGenericWithName(name string) (*GstGLEffectsGeneric, error) {
	e, err := gst.NewElementWithName("gleffects", name)
	if err != nil {
		return nil, err
	}
	return &GstGLEffectsGeneric{GstGLEffects: &GstGLEffects{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Select which effect apply to GL video texture
// Default: identity (0)
func (e *GstGLEffectsGeneric) SetEffect(value GstGLEffectsEffect) error {
	e.Element.SetArg("effect", string(value))
	return nil
}

func (e *GstGLEffectsGeneric) GetEffect() (GstGLEffectsEffect, error) {
	var value GstGLEffectsEffect
	var ok bool
	v, err := e.Element.GetProperty("effect")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGLEffectsEffect)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGLEffectsEffect")
	}
	return value, nil
}

// Switch video texture left to right, useful with webcams
// Default: false
func (e *GstGLEffectsGeneric) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

func (e *GstGLEffectsGeneric) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Invert colors to get dark edges on bright background when using sobel effect
// Default: false
func (e *GstGLEffectsGeneric) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

func (e *GstGLEffectsGeneric) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstGLAlphaMethod string

const (
	GstGLAlphaMethodSet    GstGLAlphaMethod = "set"    // Set/adjust alpha channel
	GstGLAlphaMethodGreen  GstGLAlphaMethod = "green"  // Chroma Key on pure green
	GstGLAlphaMethodBlue   GstGLAlphaMethod = "blue"   // Chroma Key on pure blue
	GstGLAlphaMethodCustom GstGLAlphaMethod = "custom" // Chroma Key on custom RGB values
)

type GstGLRotateMethod string

const (
	GstGLRotateMethodNone               GstGLRotateMethod = "none"                 // Identity (no rotation)
	GstGLRotateMethodClockwise          GstGLRotateMethod = "clockwise"            // Rotate clockwise 90 degrees
	GstGLRotateMethodRotate180          GstGLRotateMethod = "rotate-180"           // Rotate 180 degrees
	GstGLRotateMethodCounterclockwise   GstGLRotateMethod = "counterclockwise"     // Rotate counter-clockwise 90 degrees
	GstGLRotateMethodHorizontalFlip     GstGLRotateMethod = "horizontal-flip"      // Flip horizontally
	GstGLRotateMethodVerticalFlip       GstGLRotateMethod = "vertical-flip"        // Flip vertically
	GstGLRotateMethodUpperLeftDiagonal  GstGLRotateMethod = "upper-left-diagonal"  // Flip across upper left/lower right diagonal
	GstGLRotateMethodUpperRightDiagonal GstGLRotateMethod = "upper-right-diagonal" // Flip across upper right/lower left diagonal
	GstGLRotateMethodAutomatic          GstGLRotateMethod = "automatic"            // Select rotate method based on image-orientation tag
)

type GstGLVideoMixerSizingPolicy string

const (
	GstGLVideoMixerSizingPolicyNone            GstGLVideoMixerSizingPolicy = "none"              // None: Image is scaled to fill configured destination rectangle without padding or keeping the aspect ratio
	GstGLVideoMixerSizingPolicyKeepAspectRatio GstGLVideoMixerSizingPolicy = "keep-aspect-ratio" // Keep Aspect Ratio: Image is scaled to fit destination rectangle specified by GstGLVideoMixerPad:{xpos, ypos, width, height} with preserved aspect ratio. The empty space of the resulting image will be distributed in the destination rectangle according to the GstGLVideoMixerPad:{xalign, yalign} values
)

type GstGLEffects struct {
	*GstGLFilter
}

type GstGLStereoDownmix string

const (
	GstGLStereoDownmixGreenMagentaDubois GstGLStereoDownmix = "green-magenta-dubois" // GST_GL_STEREO_DOWNMIX_ANAGLYPH_GREEN_MAGENTA_DUBOIS
	GstGLStereoDownmixRedCyanDubois      GstGLStereoDownmix = "red-cyan-dubois"      // GST_GL_STEREO_DOWNMIX_ANAGLYPH_RED_CYAN_DUBOIS
	GstGLStereoDownmixAmberBlueDubois    GstGLStereoDownmix = "amber-blue-dubois"    // GST_GL_STEREO_DOWNMIX_ANAGLYPH_AMBER_BLUE_DUBOIS
)

type GstGLTestSrcPattern string

const (
	GstGLTestSrcPatternSmpte      GstGLTestSrcPattern = "smpte"      // SMPTE 100%% color bars
	GstGLTestSrcPatternSnow       GstGLTestSrcPattern = "snow"       // Random (television snow)
	GstGLTestSrcPatternBlack      GstGLTestSrcPattern = "black"      // 100%% Black
	GstGLTestSrcPatternWhite      GstGLTestSrcPattern = "white"      // 100%% White
	GstGLTestSrcPatternRed        GstGLTestSrcPattern = "red"        // Red
	GstGLTestSrcPatternGreen      GstGLTestSrcPattern = "green"      // Green
	GstGLTestSrcPatternBlue       GstGLTestSrcPattern = "blue"       // Blue
	GstGLTestSrcPatternCheckers1  GstGLTestSrcPattern = "checkers-1" // Checkers 1px
	GstGLTestSrcPatternCheckers2  GstGLTestSrcPattern = "checkers-2" // Checkers 2px
	GstGLTestSrcPatternCheckers4  GstGLTestSrcPattern = "checkers-4" // Checkers 4px
	GstGLTestSrcPatternCheckers8  GstGLTestSrcPattern = "checkers-8" // Checkers 8px
	GstGLTestSrcPatternCircular   GstGLTestSrcPattern = "circular"   // Circular
	GstGLTestSrcPatternBlink      GstGLTestSrcPattern = "blink"      // Blink
	GstGLTestSrcPatternMandelbrot GstGLTestSrcPattern = "mandelbrot" // Mandelbrot Fractal
)

type GstGLVideoFlipMethod string

const (
	GstGLVideoFlipMethodNone               GstGLVideoFlipMethod = "none"                 // Identity (no rotation)
	GstGLVideoFlipMethodClockwise          GstGLVideoFlipMethod = "clockwise"            // Rotate clockwise 90 degrees
	GstGLVideoFlipMethodRotate180          GstGLVideoFlipMethod = "rotate-180"           // Rotate 180 degrees
	GstGLVideoFlipMethodCounterclockwise   GstGLVideoFlipMethod = "counterclockwise"     // Rotate counter-clockwise 90 degrees
	GstGLVideoFlipMethodHorizontalFlip     GstGLVideoFlipMethod = "horizontal-flip"      // Flip horizontally
	GstGLVideoFlipMethodVerticalFlip       GstGLVideoFlipMethod = "vertical-flip"        // Flip vertically
	GstGLVideoFlipMethodUpperLeftDiagonal  GstGLVideoFlipMethod = "upper-left-diagonal"  // Flip across upper left/lower right diagonal
	GstGLVideoFlipMethodUpperRightDiagonal GstGLVideoFlipMethod = "upper-right-diagonal" // Flip across upper right/lower left diagonal
	GstGLVideoFlipMethodAutomatic          GstGLVideoFlipMethod = "automatic"            // Select flip method based on image-orientation tag
)

type GstGLVideoMixerBlendEquation string

const (
	GstGLVideoMixerBlendEquationAdd             GstGLVideoMixerBlendEquation = "add"              // Add
	GstGLVideoMixerBlendEquationSubtract        GstGLVideoMixerBlendEquation = "subtract"         // Subtract
	GstGLVideoMixerBlendEquationReverseSubtract GstGLVideoMixerBlendEquation = "reverse-subtract" // Reverse Subtract
)

type GstGLDeinterlaceMethod string

const (
	GstGLDeinterlaceMethodVfir    GstGLDeinterlaceMethod = "vfir"    // Blur Vertical
	GstGLDeinterlaceMethodGreedyh GstGLDeinterlaceMethod = "greedyh" // Motion Adaptive: Advanced Detection
)

type GstGLEffectsEffect string

const (
	GstGLEffectsEffectIdentity  GstGLEffectsEffect = "identity"  // Do nothing Effect
	GstGLEffectsEffectMirror    GstGLEffectsEffect = "mirror"    // Mirror Effect
	GstGLEffectsEffectSqueeze   GstGLEffectsEffect = "squeeze"   // Squeeze Effect
	GstGLEffectsEffectStretch   GstGLEffectsEffect = "stretch"   // Stretch Effect
	GstGLEffectsEffectTunnel    GstGLEffectsEffect = "tunnel"    // Light Tunnel Effect
	GstGLEffectsEffectFisheye   GstGLEffectsEffect = "fisheye"   // FishEye Effect
	GstGLEffectsEffectTwirl     GstGLEffectsEffect = "twirl"     // Twirl Effect
	GstGLEffectsEffectBulge     GstGLEffectsEffect = "bulge"     // Bulge Effect
	GstGLEffectsEffectSquare    GstGLEffectsEffect = "square"    // Square Effect
	GstGLEffectsEffectHeat      GstGLEffectsEffect = "heat"      // Heat Signature Effect
	GstGLEffectsEffectSepia     GstGLEffectsEffect = "sepia"     // Sepia Toning Effect
	GstGLEffectsEffectXpro      GstGLEffectsEffect = "xpro"      // Cross Processing Effect
	GstGLEffectsEffectLumaxpro  GstGLEffectsEffect = "lumaxpro"  // Luma Cross Processing Effect
	GstGLEffectsEffectXray      GstGLEffectsEffect = "xray"      // Glowing negative effect
	GstGLEffectsEffectSin       GstGLEffectsEffect = "sin"       // All Grey but Red Effect
	GstGLEffectsEffectGlow      GstGLEffectsEffect = "glow"      // Glow Lighting Effect
	GstGLEffectsEffectSobel     GstGLEffectsEffect = "sobel"     // Sobel edge detection Effect
	GstGLEffectsEffectBlur      GstGLEffectsEffect = "blur"      // Blur with 9x9 separable convolution Effect
	GstGLEffectsEffectLaplacian GstGLEffectsEffect = "laplacian" // Laplacian Convolution Demo Effect
)

type GstGLVideoMixerBackground string

const (
	GstGLVideoMixerBackgroundChecker     GstGLVideoMixerBackground = "checker"     // Checker pattern
	GstGLVideoMixerBackgroundBlack       GstGLVideoMixerBackground = "black"       // Black
	GstGLVideoMixerBackgroundWhite       GstGLVideoMixerBackground = "white"       // White
	GstGLVideoMixerBackgroundTransparent GstGLVideoMixerBackground = "transparent" // Transparent Background to enable further compositing
)

type GstGLVideoMixerBlendFunction string

const (
	GstGLVideoMixerBlendFunctionZero                 GstGLVideoMixerBlendFunction = "zero"                    // Zero
	GstGLVideoMixerBlendFunctionOne                  GstGLVideoMixerBlendFunction = "one"                     // One
	GstGLVideoMixerBlendFunctionSrcColor             GstGLVideoMixerBlendFunction = "src-color"               // Source Color
	GstGLVideoMixerBlendFunctionOneMinusSrcColor     GstGLVideoMixerBlendFunction = "one-minus-src-color"     // One Minus Source Color
	GstGLVideoMixerBlendFunctionDstColor             GstGLVideoMixerBlendFunction = "dst-color"               // Destination Color
	GstGLVideoMixerBlendFunctionOneMinusDstColor     GstGLVideoMixerBlendFunction = "one-minus-dst-color"     // One Minus Destination Color
	GstGLVideoMixerBlendFunctionSrcAlpha             GstGLVideoMixerBlendFunction = "src-alpha"               // Source Alpha
	GstGLVideoMixerBlendFunctionOneMinusSrcAlpha     GstGLVideoMixerBlendFunction = "one-minus-src-alpha"     // One Minus Source Alpha
	GstGLVideoMixerBlendFunctionDstAlpha             GstGLVideoMixerBlendFunction = "dst-alpha"               // Destination Alpha
	GstGLVideoMixerBlendFunctionOneMinusDstAlpha     GstGLVideoMixerBlendFunction = "one-minus-dst-alpha"     // One Minus Destination Alpha
	GstGLVideoMixerBlendFunctionConstantColor        GstGLVideoMixerBlendFunction = "constant-color"          // Constant Color
	GstGLVideoMixerBlendFunctionOneMinusContantColor GstGLVideoMixerBlendFunction = "one-minus-contant-color" // One Minus Constant Color
	GstGLVideoMixerBlendFunctionConstantAlpha        GstGLVideoMixerBlendFunction = "constant-alpha"          // Constant Alpha
	GstGLVideoMixerBlendFunctionOneMinusContantAlpha GstGLVideoMixerBlendFunction = "one-minus-contant-alpha" // One Minus Constant Alpha
	GstGLVideoMixerBlendFunctionSrcAlphaSaturate     GstGLVideoMixerBlendFunction = "src-alpha-saturate"      // Source Alpha Saturate
)

type GstGLMixerBinStartTimeSelection string

const (
	GstGLMixerBinStartTimeSelectionZero  GstGLMixerBinStartTimeSelection = "zero"  // Start at 0 running time (default)
	GstGLMixerBinStartTimeSelectionFirst GstGLMixerBinStartTimeSelection = "first" // Start at first observed input running time
	GstGLMixerBinStartTimeSelectionSet   GstGLMixerBinStartTimeSelection = "set"   // Set start time with start-time property
)
