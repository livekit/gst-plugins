// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A Vulkan image copier
type GstVulkanImageIdentity struct {
	*GstVulkanVideoFilter
}

func NewVulkanImageIdentity() (*GstVulkanImageIdentity, error) {
	e, err := gst.NewElement("vulkanimageidentity")
	if err != nil {
		return nil, err
	}
	return &GstVulkanImageIdentity{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVulkanImageIdentityWithName(name string) (*GstVulkanImageIdentity, error) {
	e, err := gst.NewElementWithName("vulkanimageidentity", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanImageIdentity{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Vulkan Overlay Composition element
type GstVulkanOverlayCompositor struct {
	*GstVulkanVideoFilter
}

func NewVulkanOverlayCompositor() (*GstVulkanOverlayCompositor, error) {
	e, err := gst.NewElement("vulkanoverlaycompositor")
	if err != nil {
		return nil, err
	}
	return &GstVulkanOverlayCompositor{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVulkanOverlayCompositorWithName(name string) (*GstVulkanOverlayCompositor, error) {
	e, err := gst.NewElementWithName("vulkanoverlaycompositor", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanOverlayCompositor{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Performs operations with SPIRV shaders in Vulkan
type GstVulkanShaderSpv struct {
	*GstVulkanVideoFilter
}

func NewVulkanShaderSpv() (*GstVulkanShaderSpv, error) {
	e, err := gst.NewElement("vulkanshaderspv")
	if err != nil {
		return nil, err
	}
	return &GstVulkanShaderSpv{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVulkanShaderSpvWithName(name string) (*GstVulkanShaderSpv, error) {
	e, err := gst.NewElementWithName("vulkanshaderspv", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanShaderSpv{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// SPIRV fragment source
// Default: NULL
func (e *GstVulkanShaderSpv) SetFragmentLocation(value string) error {
	return e.Element.SetProperty("fragment-location", value)
}

func (e *GstVulkanShaderSpv) GetFragmentLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("fragment-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// SPIRV vertex binary

func (e *GstVulkanShaderSpv) SetVertex(value interface{}) error {
	return e.Element.SetProperty("vertex", value)
}

func (e *GstVulkanShaderSpv) GetVertex() (interface{}, error) {
	return e.Element.GetProperty("vertex")
}

// SPIRV vertex source
// Default: NULL
func (e *GstVulkanShaderSpv) SetVertexLocation(value string) error {
	return e.Element.SetProperty("vertex-location", value)
}

func (e *GstVulkanShaderSpv) GetVertexLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("vertex-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// SPIRV fragment binary

func (e *GstVulkanShaderSpv) SetFragment(value interface{}) error {
	return e.Element.SetProperty("fragment", value)
}

func (e *GstVulkanShaderSpv) GetFragment() (interface{}, error) {
	return e.Element.GetProperty("fragment")
}

// A videosink based on Vulkan
type GstVulkanSink struct {
	*GstVideoSink
}

func NewVulkanSink() (*GstVulkanSink, error) {
	e, err := gst.NewElement("vulkansink")
	if err != nil {
		return nil, err
	}
	return &GstVulkanSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewVulkanSinkWithName(name string) (*GstVulkanSink, error) {
	e, err := gst.NewElementWithName("vulkansink", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Vulkan device

func (e *GstVulkanSink) GetDevice() (interface{}, error) {
	return e.Element.GetProperty("device")
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstVulkanSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstVulkanSink) GetForceAspectRatio() (bool, error) {
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

// The pixel aspect ratio of the device
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstVulkanSink) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstVulkanSink) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

// A Vulkan data uploader
type GstVulkanUpload struct {
	*base.GstBaseTransform
}

func NewVulkanUpload() (*GstVulkanUpload, error) {
	e, err := gst.NewElement("vulkanupload")
	if err != nil {
		return nil, err
	}
	return &GstVulkanUpload{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVulkanUploadWithName(name string) (*GstVulkanUpload, error) {
	e, err := gst.NewElementWithName("vulkanupload", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanUpload{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// A Vulkan View Convert
type GstVulkanViewConvert struct {
	*GstVulkanVideoFilter
}

func NewVulkanViewConvert() (*GstVulkanViewConvert, error) {
	e, err := gst.NewElement("vulkanviewconvert")
	if err != nil {
		return nil, err
	}
	return &GstVulkanViewConvert{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVulkanViewConvertWithName(name string) (*GstVulkanViewConvert, error) {
	e, err := gst.NewElementWithName("vulkanviewconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanViewConvert{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Output anaglyph type to generate when downmixing to mono
// Default: green-magenta-dubois (0)
func (e *GstVulkanViewConvert) SetDownmixMode(value GstVulkanStereoDownmix) error {
	e.Element.SetArg("downmix-mode", string(value))
	return nil
}

func (e *GstVulkanViewConvert) GetDownmixMode() (GstVulkanStereoDownmix, error) {
	var value GstVulkanStereoDownmix
	var ok bool
	v, err := e.Element.GetProperty("downmix-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVulkanStereoDownmix)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVulkanStereoDownmix")
	}
	return value, nil
}

// Override any input information about multiview layout flags
// Default: none
func (e *GstVulkanViewConvert) SetInputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("input-flags-override", value)
}

func (e *GstVulkanViewConvert) GetInputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("input-flags-override")
}

// Override any input information about multiview layout
// Default: none (-1)
func (e *GstVulkanViewConvert) SetInputModeOverride(value interface{}) error {
	return e.Element.SetProperty("input-mode-override", value)
}

func (e *GstVulkanViewConvert) GetInputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("input-mode-override")
}

// Override automatic negotiation for output multiview layout flags
// Default: none
func (e *GstVulkanViewConvert) SetOutputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("output-flags-override", value)
}

func (e *GstVulkanViewConvert) GetOutputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("output-flags-override")
}

// Override automatic output mode selection for multiview layout
// Default: none (-1)
func (e *GstVulkanViewConvert) SetOutputModeOverride(value interface{}) error {
	return e.Element.SetProperty("output-mode-override", value)
}

func (e *GstVulkanViewConvert) GetOutputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("output-mode-override")
}

// A Vulkan Color Convert
type GstVulkanColorConvert struct {
	*GstVulkanVideoFilter
}

func NewVulkanColorConvert() (*GstVulkanColorConvert, error) {
	e, err := gst.NewElement("vulkancolorconvert")
	if err != nil {
		return nil, err
	}
	return &GstVulkanColorConvert{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVulkanColorConvertWithName(name string) (*GstVulkanColorConvert, error) {
	e, err := gst.NewElementWithName("vulkancolorconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanColorConvert{GstVulkanVideoFilter: &GstVulkanVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// A Vulkan data downloader
type GstVulkanDownload struct {
	*base.GstBaseTransform
}

func NewVulkanDownload() (*GstVulkanDownload, error) {
	e, err := gst.NewElement("vulkandownload")
	if err != nil {
		return nil, err
	}
	return &GstVulkanDownload{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVulkanDownloadWithName(name string) (*GstVulkanDownload, error) {
	e, err := gst.NewElementWithName("vulkandownload", name)
	if err != nil {
		return nil, err
	}
	return &GstVulkanDownload{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

type GstVulkanStereoDownmix string

const (
	GstVulkanStereoDownmixGreenMagentaDubois GstVulkanStereoDownmix = "green-magenta-dubois" // GST_VULKAN_STEREO_DOWNMIX_ANAGLYPH_GREEN_MAGENTA_DUBOIS
	GstVulkanStereoDownmixRedCyanDubois      GstVulkanStereoDownmix = "red-cyan-dubois"      // GST_VULKAN_STEREO_DOWNMIX_ANAGLYPH_RED_CYAN_DUBOIS
	GstVulkanStereoDownmixAmberBlueDubois    GstVulkanStereoDownmix = "amber-blue-dubois"    // GST_VULKAN_STEREO_DOWNMIX_ANAGLYPH_AMBER_BLUE_DUBOIS
)
