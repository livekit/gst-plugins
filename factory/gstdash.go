// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Dynamic Adaptive Streaming over HTTP demuxer
type GstDashDemux struct {
	*GstAdaptiveDemux
}

func NewDashDemux() (*GstDashDemux, error) {
	e, err := gst.NewElement("dashdemux")
	if err != nil {
		return nil, err
	}
	return &GstDashDemux{GstAdaptiveDemux: &GstAdaptiveDemux{Bin: &gst.Bin{Element: e}}}, nil
}

func NewDashDemuxWithName(name string) (*GstDashDemux, error) {
	e, err := gst.NewElementWithName("dashdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstDashDemux{GstAdaptiveDemux: &GstAdaptiveDemux{Bin: &gst.Bin{Element: e}}}, nil
}

// Default presentation delay (in seconds, milliseconds or fragments) (e.g. 12s, 2500ms, 3f)
// Default: 10s
func (e *GstDashDemux) SetPresentationDelay(value string) error {
	return e.Element.SetProperty("presentation-delay", value)
}

func (e *GstDashDemux) GetPresentationDelay() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("presentation-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Percentage of the available bandwidth to use when selecting representations (deprecated)
// Default: 0.8, Min: 0, Max: 1
func (e *GstDashDemux) SetBandwidthUsage(value float32) error {
	return e.Element.SetProperty("bandwidth-usage", value)
}

func (e *GstDashDemux) GetBandwidthUsage() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("bandwidth-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Max of bitrate supported by target video decoder (0 = no maximum)
// Default: 0, Min: 0, Max: -1
func (e *GstDashDemux) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstDashDemux) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum number of seconds of buffer accumulated during playback(deprecated)
// Default: 30, Min: 2, Max: -1
func (e *GstDashDemux) SetMaxBufferingTime(value uint) error {
	return e.Element.SetProperty("max-buffering-time", value)
}

func (e *GstDashDemux) GetMaxBufferingTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-buffering-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max video framerate to select (0/1 = no maximum)
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstDashDemux) SetMaxVideoFramerate(value interface{}) error {
	return e.Element.SetProperty("max-video-framerate", value)
}

func (e *GstDashDemux) GetMaxVideoFramerate() (interface{}, error) {
	return e.Element.GetProperty("max-video-framerate")
}

// Max video height to select (0 = no maximum)
// Default: 0, Min: 0, Max: -1
func (e *GstDashDemux) SetMaxVideoHeight(value uint) error {
	return e.Element.SetProperty("max-video-height", value)
}

func (e *GstDashDemux) GetMaxVideoHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-video-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max video width to select (0 = no maximum)
// Default: 0, Min: 0, Max: -1
func (e *GstDashDemux) SetMaxVideoWidth(value uint) error {
	return e.Element.SetProperty("max-video-width", value)
}

func (e *GstDashDemux) GetMaxVideoWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-video-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Dynamic Adaptive Streaming over HTTP sink
type GstDashSink struct {
	*gst.Bin
}

func NewDashSink() (*GstDashSink, error) {
	e, err := gst.NewElement("dashsink")
	if err != nil {
		return nil, err
	}
	return &GstDashSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewDashSinkWithName(name string) (*GstDashSink, error) {
	e, err := gst.NewElementWithName("dashsink", name)
	if err != nil {
		return nil, err
	}
	return &GstDashSink{Bin: &gst.Bin{Element: e}}, nil
}

// Provides the explicit duration of a period in milliseconds
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstDashSink) SetPeriodDuration(value uint64) error {
	return e.Element.SetProperty("period-duration", value)
}

func (e *GstDashSink) GetPeriodDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("period-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Send keyframe requests to ensure correct fragmentation. If this is disabled then the input must have keyframes in regular intervals
// Default: true
func (e *GstDashSink) SetSendKeyframeRequests(value bool) error {
	return e.Element.SetProperty("send-keyframe-requests", value)
}

func (e *GstDashSink) GetSendKeyframeRequests() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-keyframe-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Provides a dynamic mpd
// Default: false
func (e *GstDashSink) SetDynamic(value bool) error {
	return e.Element.SetProperty("dynamic", value)
}

func (e *GstDashSink) GetDynamic() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dynamic")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Provides to the manifest a minimum update period in milliseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstDashSink) SetMinimumUpdatePeriod(value uint64) error {
	return e.Element.SetProperty("minimum-update-period", value)
}

func (e *GstDashSink) GetMinimumUpdatePeriod() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("minimum-update-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Path where the MPD and its fragents will be written
// Default: NULL
func (e *GstDashSink) SetMpdRootPath(value string) error {
	return e.Element.SetProperty("mpd-root-path", value)
}

func (e *GstDashSink) GetMpdRootPath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mpd-root-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Muxer type to be used by dashsink to generate the fragment
// Default: ts (0)
func (e *GstDashSink) SetMuxer(value GstDashSinkMuxerType) error {
	e.Element.SetArg("muxer", string(value))
	return nil
}

func (e *GstDashSink) GetMuxer() (GstDashSinkMuxerType, error) {
	var value GstDashSinkMuxerType
	var ok bool
	v, err := e.Element.GetProperty("muxer")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDashSinkMuxerType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDashSinkMuxerType")
	}
	return value, nil
}

// The target duration in seconds of a segment/file. (0 - disabled, useful for management of segment duration by the streaming server)
// Default: 15, Min: 0, Max: -1
func (e *GstDashSink) SetTargetDuration(value uint) error {
	return e.Element.SetProperty("target-duration", value)
}

func (e *GstDashSink) GetTargetDuration() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Use segment list instead of segment template to create the segments
// Default: false
func (e *GstDashSink) SetUseSegmentList(value bool) error {
	return e.Element.SetProperty("use-segment-list", value)
}

func (e *GstDashSink) GetUseSegmentList() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-segment-list")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Provides to the manifest a minimum buffer time in milliseconds
// Default: 2000, Min: 0, Max: 18446744073709551615
func (e *GstDashSink) SetMinBufferTime(value uint64) error {
	return e.Element.SetProperty("min-buffer-time", value)
}

func (e *GstDashSink) GetMinBufferTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// BaseURL to set in the MPD
// Default: NULL
func (e *GstDashSink) SetMpdBaseurl(value string) error {
	return e.Element.SetProperty("mpd-baseurl", value)
}

func (e *GstDashSink) GetMpdBaseurl() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mpd-baseurl")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// filename of the mpd to write
// Default: dash.mpd
func (e *GstDashSink) SetMpdFilename(value string) error {
	return e.Element.SetProperty("mpd-filename", value)
}

func (e *GstDashSink) GetMpdFilename() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mpd-filename")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Provides to the manifest a suggested presentation delay in milliseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstDashSink) SetSuggestedPresentationDelay(value uint64) error {
	return e.Element.SetProperty("suggested-presentation-delay", value)
}

func (e *GstDashSink) GetSuggestedPresentationDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("suggested-presentation-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

type GstDashSinkMuxerType string

const (
	GstDashSinkMuxerTypeTs      GstDashSinkMuxerType = "ts"      // Use mpegtsmux
	GstDashSinkMuxerTypeMp4     GstDashSinkMuxerType = "mp4"     // Use mp4mux (deprecated, non-functional)
	GstDashSinkMuxerTypeDashmp4 GstDashSinkMuxerType = "dashmp4" // Use dashmp4mux
)
