// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Take image snapshots and record movies from camera
type GstCameraBin struct {
	*gst.Pipeline
}

func NewCameraBin() (*GstCameraBin, error) {
	e, err := gst.NewElement("camerabin")
	if err != nil {
		return nil, err
	}
	return &GstCameraBin{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

func NewCameraBinWithName(name string) (*GstCameraBin, error) {
	e, err := gst.NewElementWithName("camerabin", name)
	if err != nil {
		return nil, err
	}
	return &GstCameraBin{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

// Location to save the captured files. A %%d might be used on thefilename as a placeholder for a numeric index of the capture.Default is cap_%%d
// Default: cap_%%d
func (e *GstCameraBin) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstCameraBin) GetLocation() (string, error) {
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

// If capture preview images should be posted to the bus
// Default: false
func (e *GstCameraBin) SetPostPreviews(value bool) error {
	return e.Element.SetProperty("post-previews", value)
}

func (e *GstCameraBin) GetPostPreviews() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-previews")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Caps for video capture
// Default: ANY
func (e *GstCameraBin) SetVideoCaptureCaps(value *gst.Caps) error {
	return e.Element.SetProperty("video-capture-caps", value)
}

func (e *GstCameraBin) GetVideoCaptureCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("video-capture-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Formats supported for capturing videos represented as GstCaps

func (e *GstCameraBin) GetVideoCaptureSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("video-capture-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// The element that will process captured video frames. (Should be set on NULL state)

func (e *GstCameraBin) SetVideoFilter(value *gst.Element) error {
	return e.Element.SetProperty("video-filter", value)
}

func (e *GstCameraBin) GetVideoFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Formats supported for capturing audio represented as GstCaps

func (e *GstCameraBin) GetAudioCaptureSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("audio-capture-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// The capture mode (still image capture or video recording)
// Default: mode-image (1)
func (e *GstCameraBin) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

func (e *GstCameraBin) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

// The caps of the preview image to be posted

func (e *GstCameraBin) SetPreviewCaps(value *gst.Caps) error {
	return e.Element.SetProperty("preview-caps", value)
}

func (e *GstCameraBin) GetPreviewCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("preview-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// The element that will process preview buffers. (Should be set on NULL state)

func (e *GstCameraBin) SetPreviewFilter(value *gst.Element) error {
	return e.Element.SetProperty("preview-filter", value)
}

func (e *GstCameraBin) GetPreviewFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("preview-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Digital zoom factor (e.g. 1.5 means 1.5x)
// Default: 1, Min: 1, Max: 10
func (e *GstCameraBin) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

func (e *GstCameraBin) GetZoom() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The GstEncodingProfile to use for video recording. Audio is enabled when this profile supports audio.

func (e *GstCameraBin) SetVideoProfile(value interface{}) error {
	return e.Element.SetProperty("video-profile", value)
}

func (e *GstCameraBin) GetVideoProfile() (interface{}, error) {
	return e.Element.GetProperty("video-profile")
}

// The video sink of the viewfinder. It is only taken into use on the next null to ready transition

func (e *GstCameraBin) SetViewfinderSink(value *gst.Element) error {
	return e.Element.SetProperty("viewfinder-sink", value)
}

func (e *GstCameraBin) GetViewfinderSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("viewfinder-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Flags to control behaviour
// Default: (none)
func (e *GstCameraBin) SetFlags(value GstCamFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstCameraBin) GetFlags() (GstCamFlags, error) {
	var value GstCamFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCamFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCamFlags")
	}
	return value, nil
}

// If camerabin2 is idle (not doing captures).
// Default: true
func (e *GstCameraBin) GetIdle() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("idle")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The element that will process frames going to the viewfinder. (Should be set on NULL state)

func (e *GstCameraBin) SetViewfinderFilter(value *gst.Element) error {
	return e.Element.SetProperty("viewfinder-filter", value)
}

func (e *GstCameraBin) GetViewfinderFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("viewfinder-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Caps for image capture
// Default: ANY
func (e *GstCameraBin) SetImageCaptureCaps(value *gst.Caps) error {
	return e.Element.SetProperty("image-capture-caps", value)
}

func (e *GstCameraBin) GetImageCaptureCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("image-capture-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Formats supported for capturing images represented as GstCaps

func (e *GstCameraBin) GetImageCaptureSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("image-capture-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// The element that will process captured audio buffers when recording. (Should be set on NULL state)

func (e *GstCameraBin) SetAudioFilter(value *gst.Element) error {
	return e.Element.SetProperty("audio-filter", value)
}

func (e *GstCameraBin) GetAudioFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The element that will process captured image frames. (Should be set on NULL state)

func (e *GstCameraBin) SetImageFilter(value *gst.Element) error {
	return e.Element.SetProperty("image-filter", value)
}

func (e *GstCameraBin) GetImageFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("image-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Restricts the caps that can be used on the viewfinder
// Default: ANY
func (e *GstCameraBin) SetViewfinderCaps(value *gst.Caps) error {
	return e.Element.SetProperty("viewfinder-caps", value)
}

func (e *GstCameraBin) GetViewfinderCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("viewfinder-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// The caps that the camera source can produce on the viewfinder pad

func (e *GstCameraBin) GetViewfinderSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("viewfinder-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Format to capture audio for video recording represented as GstCaps
// Default: ANY
func (e *GstCameraBin) SetAudioCaptureCaps(value *gst.Caps) error {
	return e.Element.SetProperty("audio-capture-caps", value)
}

func (e *GstCameraBin) GetAudioCaptureCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("audio-capture-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// The audio source element to be used on video recordings. It is only taken into use on the next null to ready transition

func (e *GstCameraBin) SetAudioSource(value *gst.Element) error {
	return e.Element.SetProperty("audio-source", value)
}

func (e *GstCameraBin) GetAudioSource() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The camera source element to be used. It is only taken into use on the next null to ready transition

func (e *GstCameraBin) SetCameraSource(value *gst.Element) error {
	return e.Element.SetProperty("camera-source", value)
}

func (e *GstCameraBin) GetCameraSource() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("camera-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The GstEncodingProfile to use for image captures.

func (e *GstCameraBin) SetImageProfile(value interface{}) error {
	return e.Element.SetProperty("image-profile", value)
}

func (e *GstCameraBin) GetImageProfile() (interface{}, error) {
	return e.Element.GetProperty("image-profile")
}

// Digital zoom factor (e.g. 1.5 means 1.5x)
// Default: 10, Min: 1, Max: 3.40282e+38
func (e *GstCameraBin) GetMaxZoom() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("max-zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// If the audio recording should be muted. Note that this still saves audio data to the resulting file, but they are silent. Use a video-profile without audio to disable audio completely
// Default: false
func (e *GstCameraBin) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstCameraBin) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Viewfinder Bin used in camerabin2
type GstViewfinderBin struct {
	*gst.Bin
}

func NewViewfinderBin() (*GstViewfinderBin, error) {
	e, err := gst.NewElement("viewfinderbin")
	if err != nil {
		return nil, err
	}
	return &GstViewfinderBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewViewfinderBinWithName(name string) (*GstViewfinderBin, error) {
	e, err := gst.NewElementWithName("viewfinderbin", name)
	if err != nil {
		return nil, err
	}
	return &GstViewfinderBin{Bin: &gst.Bin{Element: e}}, nil
}

// If video converters should be disabled (must be set on NULL)
// Default: false
func (e *GstViewfinderBin) SetDisableConverters(value bool) error {
	return e.Element.SetProperty("disable-converters", value)
}

func (e *GstViewfinderBin) GetDisableConverters() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-converters")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// the video output element to use (NULL = default)

func (e *GstViewfinderBin) SetVideoSink(value *gst.Element) error {
	return e.Element.SetProperty("video-sink", value)
}

func (e *GstViewfinderBin) GetVideoSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Wrapper camera src element for camerabin2
type GstWrapperCameraBinSrc struct {
	*GstBaseCameraSrc
}

func NewWrapperCameraBinSrc() (*GstWrapperCameraBinSrc, error) {
	e, err := gst.NewElement("wrappercamerabinsrc")
	if err != nil {
		return nil, err
	}
	return &GstWrapperCameraBinSrc{GstBaseCameraSrc: &GstBaseCameraSrc{Bin: &gst.Bin{Element: e}}}, nil
}

func NewWrapperCameraBinSrcWithName(name string) (*GstWrapperCameraBinSrc, error) {
	e, err := gst.NewElementWithName("wrappercamerabinsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstWrapperCameraBinSrc{GstBaseCameraSrc: &GstBaseCameraSrc{Bin: &gst.Bin{Element: e}}}, nil
}

// Optional video source filter element

func (e *GstWrapperCameraBinSrc) SetVideoSourceFilter(value *gst.Element) error {
	return e.Element.SetProperty("video-source-filter", value)
}

func (e *GstWrapperCameraBinSrc) GetVideoSourceFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-source-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The video source element to be used

func (e *GstWrapperCameraBinSrc) SetVideoSource(value *gst.Element) error {
	return e.Element.SetProperty("video-source", value)
}

func (e *GstWrapperCameraBinSrc) GetVideoSource() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

type GstCamFlags int

const (
	GstCamFlagsNoAudioConversion      GstCamFlags = 0x00000001 // Do not use audio conversion elements
	GstCamFlagsNoVideoConversion      GstCamFlags = 0x00000002 // Do not use video conversion elements
	GstCamFlagsNoViewfinderConversion GstCamFlags = 0x00000004 // Do not use viewfinder conversion elements
	GstCamFlagsNoImageConversion      GstCamFlags = 0x00000008 // Do not use image conversion elements
)
