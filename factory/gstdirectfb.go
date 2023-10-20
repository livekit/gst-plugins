// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A DirectFB based videosink
type GstDfbVideoSink struct {
	*GstVideoSink
}

func NewDfbVideoSink() (*GstDfbVideoSink, error) {
	e, err := gst.NewElement("dfbvideosink")
	if err != nil {
		return nil, err
	}
	return &GstDfbVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewDfbVideoSinkWithName(name string) (*GstDfbVideoSink, error) {
	e, err := gst.NewElementWithName("dfbvideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstDfbVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// The saturation of the video
// Default: -1, Min: 0, Max: 65535
func (e *GstDfbVideoSink) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstDfbVideoSink) GetSaturation() (int, error) {
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

// The brightness of the video
// Default: -1, Min: 0, Max: 65535
func (e *GstDfbVideoSink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstDfbVideoSink) GetBrightness() (int, error) {
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

// Enable the last-sample property
// Default: true
func (e *GstDfbVideoSink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

func (e *GstDfbVideoSink) GetEnableLastSample() (bool, error) {
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

// The hue of the video
// Default: -1, Min: 0, Max: 65535
func (e *GstDfbVideoSink) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstDfbVideoSink) GetHue() (int, error) {
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

// Additional render delay of the sink in nanoseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstDfbVideoSink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

func (e *GstDfbVideoSink) GetRenderDelay() (uint64, error) {
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

// The time to keep between rendered buffers (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstDfbVideoSink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

func (e *GstDfbVideoSink) GetThrottleTime() (uint64, error) {
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
func (e *GstDfbVideoSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstDfbVideoSink) GetTsOffset() (int64, error) {
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

// The contrast of the video
// Default: -1, Min: 0, Max: 65535
func (e *GstDfbVideoSink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstDfbVideoSink) GetContrast() (int, error) {
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

// The cooperative level handling the access permission (set this to 'administrative' when the cursor is required)
// Default: exclusive (1)
func (e *GstDfbVideoSink) SetLayerMode(value GstDfbVideoSinkLayerMode) error {
	e.Element.SetArg("layer-mode", string(value))
	return nil
}

func (e *GstDfbVideoSink) GetLayerMode() (GstDfbVideoSinkLayerMode, error) {
	var value GstDfbVideoSinkLayerMode
	var ok bool
	v, err := e.Element.GetProperty("layer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDfbVideoSinkLayerMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDfbVideoSinkLayerMode")
	}
	return value, nil
}

// Generate Quality-of-Service events upstream
// Default: true
func (e *GstDfbVideoSink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstDfbVideoSink) GetQos() (bool, error) {
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

// Sink Statistics
// Default: application/x-gst-base-sink-stats, average-rate=(double)0, dropped=(guint64)0, rendered=(guint64)0;
func (e *GstDfbVideoSink) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Wait for next vertical sync to draw frames
// Default: true
func (e *GstDfbVideoSink) SetVsync(value bool) error {
	return e.Element.SetProperty("vsync", value)
}

func (e *GstDfbVideoSink) GetVsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vsync")
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
// Default: 4096, Min: 0, Max: -1
func (e *GstDfbVideoSink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstDfbVideoSink) GetBlocksize() (uint, error) {
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

// The last sample received in the sink

func (e *GstDfbVideoSink) GetLastSample() (*gst.Sample, error) {
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

// Whether to render video frames during preroll
// Default: true
func (e *GstDfbVideoSink) SetShowPrerollFrame(value bool) error {
	return e.Element.SetProperty("show-preroll-frame", value)
}

func (e *GstDfbVideoSink) GetShowPrerollFrame() (bool, error) {
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

// The target surface for video

func (e *GstDfbVideoSink) SetSurface(value interface{}) error {
	return e.Element.SetProperty("surface", value)
}

// Maximum processing time for a buffer in nanoseconds
// Default: 15000000, Min: 0, Max: 18446744073709551615
func (e *GstDfbVideoSink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstDfbVideoSink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
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
// Default: true
func (e *GstDfbVideoSink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstDfbVideoSink) GetSync() (bool, error) {
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

// Go asynchronously to PAUSED
// Default: true
func (e *GstDfbVideoSink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

func (e *GstDfbVideoSink) GetAsync() (bool, error) {
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
func (e *GstDfbVideoSink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstDfbVideoSink) GetMaxBitrate() (uint64, error) {
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

// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)
// Default: 5000000, Min: -1, Max: 9223372036854775807
func (e *GstDfbVideoSink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

func (e *GstDfbVideoSink) GetMaxLateness() (int64, error) {
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

// The pixel aspect ratio of the device
// Default: NULL
func (e *GstDfbVideoSink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstDfbVideoSink) GetPixelAspectRatio() (string, error) {
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

type GstDfbVideoSinkLayerMode string

const (
	GstDfbVideoSinkLayerModeNone           GstDfbVideoSinkLayerMode = "none"           // NONE
	GstDfbVideoSinkLayerModeExclusive      GstDfbVideoSinkLayerMode = "exclusive"      // DLSCL_EXCLUSIVE
	GstDfbVideoSinkLayerModeAdministrative GstDfbVideoSinkLayerMode = "administrative" // DLSCL_ADMINISTRATIVE
)
