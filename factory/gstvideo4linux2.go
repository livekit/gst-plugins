// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Controls a Video4Linux2 radio device
type GstV4l2Radio struct {
	*gst.Element
}

func NewV4l2Radio() (*GstV4l2Radio, error) {
	e, err := gst.NewElement("v4l2radio")
	if err != nil {
		return nil, err
	}
	return &GstV4l2Radio{Element: e}, nil
}

func NewV4l2RadioWithName(name string) (*GstV4l2Radio, error) {
	e, err := gst.NewElementWithName("v4l2radio", name)
	if err != nil {
		return nil, err
	}
	return &GstV4l2Radio{Element: e}, nil
}

// Video4Linux2 radio device location
// Default: /dev/radio0
func (e *GstV4l2Radio) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstV4l2Radio) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Station frequency in Hz
// Default: 0, Min: 87500000, Max: 108000000
func (e *GstV4l2Radio) SetFrequency(value int) error {
	return e.Element.SetProperty("frequency", value)
}

func (e *GstV4l2Radio) GetFrequency() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Displays frames on a video4linux2 device
type GstV4l2Sink struct {
	*GstVideoSink
}

func NewV4l2Sink() (*GstV4l2Sink, error) {
	e, err := gst.NewElement("v4l2sink")
	if err != nil {
		return nil, err
	}
	return &GstV4l2Sink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewV4l2SinkWithName(name string) (*GstV4l2Sink, error) {
	e, err := gst.NewElementWithName("v4l2sink", name)
	if err != nil {
		return nil, err
	}
	return &GstV4l2Sink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// The leftmost (x) coordinate of the video crop; top left corner of image is 0,0
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetCropLeft(value int) error {
	return e.Element.SetProperty("crop-left", value)
}

func (e *GstV4l2Sink) GetCropLeft() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("crop-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The topmost (y) coordinate of the video crop; top left corner of image is 0,0
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetCropTop(value int) error {
	return e.Element.SetProperty("crop-top", value)
}

func (e *GstV4l2Sink) GetCropTop() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("crop-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// File descriptor of the device
// Default: -1, Min: -1, Max: 2147483647
func (e *GstV4l2Sink) GetDeviceFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Extra v4l2 controls (CIDs) for the device

func (e *GstV4l2Sink) SetExtraControls(value *gst.Structure) error {
	return e.Element.SetProperty("extra-controls", value)
}

func (e *GstV4l2Sink) GetExtraControls() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("extra-controls")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// When enabled, the pixel aspect ratio will be enforced
// Default: true
func (e *GstV4l2Sink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstV4l2Sink) GetForceAspectRatio() (bool, error) {
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

// Hue or color balance
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstV4l2Sink) GetHue() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The width of the video overlay; default is equal to negotiated image width
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Sink) SetOverlayWidth(value uint) error {
	return e.Element.SetProperty("overlay-width", value)
}

func (e *GstV4l2Sink) GetOverlayWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("overlay-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The height of the video crop; default is equal to negotiated image height
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Sink) SetCropHeight(value uint) error {
	return e.Element.SetProperty("crop-height", value)
}

func (e *GstV4l2Sink) GetCropHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// video standard
// Default: none (0)
func (e *GstV4l2Sink) SetNorm(value V4L2TVNorms) error {
	return e.Element.SetProperty("norm", value)
}

func (e *GstV4l2Sink) GetNorm() (V4L2TVNorms, error) {
	var value V4L2TVNorms
	var ok bool
	v, err := e.Element.GetProperty("norm")
	if err != nil {
		return value, err
	}
	value, ok = v.(V4L2TVNorms)
	if !ok {
		return value, fmt.Errorf("could not cast value to V4L2TVNorms")
	}
	return value, nil
}

// Overwrite the pixel aspect ratio of the device
// Default: NULL
func (e *GstV4l2Sink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstV4l2Sink) GetPixelAspectRatio() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pixel-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Picture color saturation or chroma gain
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstV4l2Sink) GetSaturation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Name of the device
// Default: NULL
func (e *GstV4l2Sink) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Device location
// Default: /dev/video1
func (e *GstV4l2Sink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstV4l2Sink) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Device type flags
// Default: (none)
func (e *GstV4l2Sink) GetFlags() (GstV4l2DeviceTypeFlags, error) {
	var value GstV4l2DeviceTypeFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstV4l2DeviceTypeFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstV4l2DeviceTypeFlags")
	}
	return value, nil
}

// Picture contrast or luma gain
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstV4l2Sink) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The width of the video crop; default is equal to negotiated image width
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Sink) SetCropWidth(value uint) error {
	return e.Element.SetProperty("crop-width", value)
}

func (e *GstV4l2Sink) GetCropWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// I/O mode
// Default: auto (0)
func (e *GstV4l2Sink) SetIoMode(value GstV4l2IOMode) error {
	e.Element.SetArg("io-mode", string(value))
	return nil
}

func (e *GstV4l2Sink) GetIoMode() (GstV4l2IOMode, error) {
	var value GstV4l2IOMode
	var ok bool
	v, err := e.Element.GetProperty("io-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstV4l2IOMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstV4l2IOMode")
	}
	return value, nil
}

// The height of the video overlay; default is equal to negotiated image height
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Sink) SetOverlayHeight(value uint) error {
	return e.Element.SetProperty("overlay-height", value)
}

func (e *GstV4l2Sink) GetOverlayHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("overlay-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The leftmost (x) coordinate of the video overlay; top left corner of screen is 0,0
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetOverlayLeft(value int) error {
	return e.Element.SetProperty("overlay-left", value)
}

func (e *GstV4l2Sink) GetOverlayLeft() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The topmost (y) coordinate of the video overlay; top left corner of screen is 0,0
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetOverlayTop(value int) error {
	return e.Element.SetProperty("overlay-top", value)
}

func (e *GstV4l2Sink) GetOverlayTop() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Picture brightness, or more precisely, the black level
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Sink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstV4l2Sink) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Reads frames from a Video4Linux2 device
type GstV4l2Src struct {
	*base.GstPushSrc
}

func NewV4l2Src() (*GstV4l2Src, error) {
	e, err := gst.NewElement("v4l2src")
	if err != nil {
		return nil, err
	}
	return &GstV4l2Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewV4l2SrcWithName(name string) (*GstV4l2Src, error) {
	e, err := gst.NewElementWithName("v4l2src", name)
	if err != nil {
		return nil, err
	}
	return &GstV4l2Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Picture brightness, or more precisely, the black level
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Src) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstV4l2Src) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The bounding region for crop rectangles ('<x, y, width, height>').

func (e *GstV4l2Src) GetCropBounds() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("crop-bounds")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Pixels to crop at top of video capture input
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Src) SetCropTop(value uint) error {
	return e.Element.SetProperty("crop-top", value)
}

func (e *GstV4l2Src) GetCropTop() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// File descriptor of the device
// Default: -1, Min: -1, Max: 2147483647
func (e *GstV4l2Src) GetDeviceFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Name of the device
// Default: NULL
func (e *GstV4l2Src) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Extra v4l2 controls (CIDs) for the device

func (e *GstV4l2Src) SetExtraControls(value *gst.Structure) error {
	return e.Element.SetProperty("extra-controls", value)
}

func (e *GstV4l2Src) GetExtraControls() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("extra-controls")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Device type flags
// Default: (none)
func (e *GstV4l2Src) GetFlags() (GstV4l2DeviceTypeFlags, error) {
	var value GstV4l2DeviceTypeFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstV4l2DeviceTypeFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstV4l2DeviceTypeFlags")
	}
	return value, nil
}

// video standard
// Default: none (0)
func (e *GstV4l2Src) SetNorm(value V4L2TVNorms) error {
	return e.Element.SetProperty("norm", value)
}

func (e *GstV4l2Src) GetNorm() (V4L2TVNorms, error) {
	var value V4L2TVNorms
	var ok bool
	v, err := e.Element.GetProperty("norm")
	if err != nil {
		return value, err
	}
	value, ok = v.(V4L2TVNorms)
	if !ok {
		return value, fmt.Errorf("could not cast value to V4L2TVNorms")
	}
	return value, nil
}

// Pixels to crop at left of video capture input
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Src) SetCropLeft(value uint) error {
	return e.Element.SetProperty("crop-left", value)
}

func (e *GstV4l2Src) GetCropLeft() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Hue or color balance
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Src) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstV4l2Src) GetHue() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Overwrite the pixel aspect ratio of the device
// Default: NULL
func (e *GstV4l2Src) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstV4l2Src) GetPixelAspectRatio() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pixel-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Picture contrast or luma gain
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Src) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstV4l2Src) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pixels to crop at right of video capture input
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Src) SetCropRight(value uint) error {
	return e.Element.SetProperty("crop-right", value)
}

func (e *GstV4l2Src) GetCropRight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Picture color saturation or chroma gain
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstV4l2Src) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstV4l2Src) GetSaturation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pixels to crop at bottom of video capture input
// Default: 0, Min: 0, Max: -1
func (e *GstV4l2Src) SetCropBottom(value uint) error {
	return e.Element.SetProperty("crop-bottom", value)
}

func (e *GstV4l2Src) GetCropBottom() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Device location
// Default: /dev/video0
func (e *GstV4l2Src) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstV4l2Src) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// When enabled, the pixel aspect ratio will be enforced
// Default: true
func (e *GstV4l2Src) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstV4l2Src) GetForceAspectRatio() (bool, error) {
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

// I/O mode
// Default: auto (0)
func (e *GstV4l2Src) SetIoMode(value GstV4l2IOMode) error {
	e.Element.SetArg("io-mode", string(value))
	return nil
}

func (e *GstV4l2Src) GetIoMode() (GstV4l2IOMode, error) {
	var value GstV4l2IOMode
	var ok bool
	v, err := e.Element.GetProperty("io-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstV4l2IOMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstV4l2IOMode")
	}
	return value, nil
}

type GstV4l2DeviceTypeFlags int

const (
	GstV4l2DeviceTypeFlagsCapture    GstV4l2DeviceTypeFlags = 0x00000001 // Device supports video capture
	GstV4l2DeviceTypeFlagsOutput     GstV4l2DeviceTypeFlags = 0x00000002 // Device supports video playback
	GstV4l2DeviceTypeFlagsOverlay    GstV4l2DeviceTypeFlags = 0x00000004 // Device supports video overlay
	GstV4l2DeviceTypeFlagsVbiCapture GstV4l2DeviceTypeFlags = 0x00000010 // Device supports the VBI capture
	GstV4l2DeviceTypeFlagsVbiOutput  GstV4l2DeviceTypeFlags = 0x00000020 // Device supports the VBI output
	GstV4l2DeviceTypeFlagsTuner      GstV4l2DeviceTypeFlags = 0x00010000 // Device has a tuner or modulator
	GstV4l2DeviceTypeFlagsAudio      GstV4l2DeviceTypeFlags = 0x00020000 // Device has audio inputs or outputs
)

type GstV4l2IOMode string

const (
	GstV4l2IOModeAuto         GstV4l2IOMode = "auto"          // GST_V4L2_IO_AUTO
	GstV4l2IOModeRw           GstV4l2IOMode = "rw"            // GST_V4L2_IO_RW
	GstV4l2IOModeMmap         GstV4l2IOMode = "mmap"          // GST_V4L2_IO_MMAP
	GstV4l2IOModeUserptr      GstV4l2IOMode = "userptr"       // GST_V4L2_IO_USERPTR
	GstV4l2IOModeDmabuf       GstV4l2IOMode = "dmabuf"        // GST_V4L2_IO_DMABUF
	GstV4l2IOModeDmabufImport GstV4l2IOMode = "dmabuf-import" // GST_V4L2_IO_DMABUF_IMPORT
)

type V4L2TVNorms string

const (
	V4L2TVNormsNone    V4L2TVNorms = "none"      // none
	V4L2TVNormsNTSC    V4L2TVNorms = "NTSC"      // NTSC
	V4L2TVNormsNTSCM   V4L2TVNorms = "NTSC-M"    // NTSC-M
	V4L2TVNormsNTSCMJP V4L2TVNorms = "NTSC-M-JP" // NTSC-M-JP
	V4L2TVNormsNTSCMKR V4L2TVNorms = "NTSC-M-KR" // NTSC-M-KR
	V4L2TVNormsNTSC443 V4L2TVNorms = "NTSC-443"  // NTSC-443
	V4L2TVNormsPAL     V4L2TVNorms = "PAL"       // PAL
	V4L2TVNormsPALBG   V4L2TVNorms = "PAL-BG"    // PAL-BG
	V4L2TVNormsPALB    V4L2TVNorms = "PAL-B"     // PAL-B
	V4L2TVNormsPALB1   V4L2TVNorms = "PAL-B1"    // PAL-B1
	V4L2TVNormsPALG    V4L2TVNorms = "PAL-G"     // PAL-G
	V4L2TVNormsPALH    V4L2TVNorms = "PAL-H"     // PAL-H
	V4L2TVNormsPALI    V4L2TVNorms = "PAL-I"     // PAL-I
	V4L2TVNormsPALDK   V4L2TVNorms = "PAL-DK"    // PAL-DK
	V4L2TVNormsPALD    V4L2TVNorms = "PAL-D"     // PAL-D
	V4L2TVNormsPALD1   V4L2TVNorms = "PAL-D1"    // PAL-D1
	V4L2TVNormsPALK    V4L2TVNorms = "PAL-K"     // PAL-K
	V4L2TVNormsPALM    V4L2TVNorms = "PAL-M"     // PAL-M
	V4L2TVNormsPALN    V4L2TVNorms = "PAL-N"     // PAL-N
	V4L2TVNormsPALNc   V4L2TVNorms = "PAL-Nc"    // PAL-Nc
	V4L2TVNormsPAL60   V4L2TVNorms = "PAL-60"    // PAL-60
	V4L2TVNormsSECAM   V4L2TVNorms = "SECAM"     // SECAM
	V4L2TVNormsSECAMB  V4L2TVNorms = "SECAM-B"   // SECAM-B
	V4L2TVNormsSECAMG  V4L2TVNorms = "SECAM-G"   // SECAM-G
	V4L2TVNormsSECAMH  V4L2TVNorms = "SECAM-H"   // SECAM-H
	V4L2TVNormsSECAMDK V4L2TVNorms = "SECAM-DK"  // SECAM-DK
	V4L2TVNormsSECAMD  V4L2TVNorms = "SECAM-D"   // SECAM-D
	V4L2TVNormsSECAMK  V4L2TVNorms = "SECAM-K"   // SECAM-K
	V4L2TVNormsSECAMK1 V4L2TVNorms = "SECAM-K1"  // SECAM-K1
	V4L2TVNormsSECAML  V4L2TVNorms = "SECAM-L"   // SECAM-L
	V4L2TVNormsSECAMLc V4L2TVNorms = "SECAM-Lc"  // SECAM-Lc
)
