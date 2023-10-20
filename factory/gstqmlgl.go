// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A filter that renders a QML scene onto a video stream
type GstQtOverlay struct {
	*GstGLFilter
}

func NewQtOverlay() (*GstQtOverlay, error) {
	e, err := gst.NewElement("qmlgloverlay")
	if err != nil {
		return nil, err
	}
	return &GstQtOverlay{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewQtOverlayWithName(name string) (*GstQtOverlay, error) {
	e, err := gst.NewElementWithName("qmlgloverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstQtOverlay{GstGLFilter: &GstGLFilter{GstGLBaseFilter: &GstGLBaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Handle Quality-of-Service events
// Default: true
func (e *GstQtOverlay) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstQtOverlay) GetQos() (bool, error) {
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

// The root QQuickItem from the qml-scene used to render

func (e *GstQtOverlay) GetRootItem() (interface{}, error) {
	return e.Element.GetProperty("root-item")
}

// The QQuickItem to place the input video in the object hierarchy

func (e *GstQtOverlay) SetWidget(value interface{}) error {
	return e.Element.SetProperty("widget", value)
}

func (e *GstQtOverlay) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}

// Get OpenGL context

func (e *GstQtOverlay) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// The contents of the QML scene
// Default: NULL
func (e *GstQtOverlay) SetQmlScene(value string) error {
	return e.Element.SetProperty("qml-scene", value)
}

func (e *GstQtOverlay) GetQmlScene() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("qml-scene")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// A video sink that renders to a QQuickItem
type GstQtSink struct {
	*GstVideoSink
}

func NewQtSink() (*GstQtSink, error) {
	e, err := gst.NewElement("qmlglsink")
	if err != nil {
		return nil, err
	}
	return &GstQtSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewQtSinkWithName(name string) (*GstQtSink, error) {
	e, err := gst.NewElementWithName("qmlglsink", name)
	if err != nil {
		return nil, err
	}
	return &GstQtSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstQtSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstQtSink) GetTsOffset() (int64, error) {
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

// The pixel aspect ratio of the device
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstQtSink) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstQtSink) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

// Maximum processing time for a buffer in nanoseconds
// Default: 15000000, Min: 0, Max: 18446744073709551615
func (e *GstQtSink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstQtSink) GetProcessingDeadline() (uint64, error) {
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

// Additional render delay of the sink in nanoseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQtSink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

func (e *GstQtSink) GetRenderDelay() (uint64, error) {
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

// The maximum bits per second to render (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQtSink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstQtSink) GetMaxBitrate() (uint64, error) {
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
func (e *GstQtSink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

func (e *GstQtSink) GetMaxLateness() (int64, error) {
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

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstQtSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstQtSink) GetForceAspectRatio() (bool, error) {
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

// Whether to render video frames during preroll
// Default: true
func (e *GstQtSink) SetShowPrerollFrame(value bool) error {
	return e.Element.SetProperty("show-preroll-frame", value)
}

func (e *GstQtSink) GetShowPrerollFrame() (bool, error) {
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

// Sink Statistics
// Default: application/x-gst-base-sink-stats, average-rate=(double)0, dropped=(guint64)0, rendered=(guint64)0;
func (e *GstQtSink) GetStats() (*gst.Structure, error) {
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

// Sync on the clock
// Default: true
func (e *GstQtSink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstQtSink) GetSync() (bool, error) {
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

// The time to keep between rendered buffers (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQtSink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

func (e *GstQtSink) GetThrottleTime() (uint64, error) {
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

// Size in bytes to pull per buffer (0 = default)
// Default: 4096, Min: 0, Max: -1
func (e *GstQtSink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstQtSink) GetBlocksize() (uint, error) {
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
// Default: true
func (e *GstQtSink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

func (e *GstQtSink) GetEnableLastSample() (bool, error) {
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

// Generate Quality-of-Service events upstream
// Default: true
func (e *GstQtSink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstQtSink) GetQos() (bool, error) {
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

// The QQuickItem to place in the object hierarchy

func (e *GstQtSink) SetWidget(value interface{}) error {
	return e.Element.SetProperty("widget", value)
}

func (e *GstQtSink) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}

// Go asynchronously to PAUSED
// Default: true
func (e *GstQtSink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

func (e *GstQtSink) GetAsync() (bool, error) {
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

// The last sample received in the sink

func (e *GstQtSink) GetLastSample() (*gst.Sample, error) {
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

// A video src that captures a window from a QML view
type GstQtSrc struct {
	*base.GstPushSrc
}

func NewQtSrc() (*GstQtSrc, error) {
	e, err := gst.NewElement("qmlglsrc")
	if err != nil {
		return nil, err
	}
	return &GstQtSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewQtSrcWithName(name string) (*GstQtSrc, error) {
	e, err := gst.NewElementWithName("qmlglsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstQtSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Size in bytes to read per buffer (-1 = default)
// Default: 4096, Min: 0, Max: -1
func (e *GstQtSrc) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstQtSrc) GetBlocksize() (uint, error) {
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

// Apply current stream time to buffers
// Default: false
func (e *GstQtSrc) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

func (e *GstQtSrc) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of buffers to output before sending EOS (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstQtSrc) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstQtSrc) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Run typefind before negotiating (deprecated, non-functional)
// Default: false
func (e *GstQtSrc) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

func (e *GstQtSrc) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When set it will not create a new FBO for the QML render thread
// Default: false
func (e *GstQtSrc) SetUseDefaultFbo(value bool) error {
	return e.Element.SetProperty("use-default-fbo", value)
}

func (e *GstQtSrc) GetUseDefaultFbo() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-default-fbo")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The QQuickWindow to place in the object hierarchy

func (e *GstQtSrc) SetWindow(value interface{}) error {
	return e.Element.SetProperty("window", value)
}

func (e *GstQtSrc) GetWindow() (interface{}, error) {
	return e.Element.GetProperty("window")
}
