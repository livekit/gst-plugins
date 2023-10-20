// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Dynamic Adaptive Streaming over HTTP demuxer
type GstDashDemux2 struct {
	*GstAdaptiveDemux2
}

func NewDashDemux2() (*GstDashDemux2, error) {
	e, err := gst.NewElement("dashdemux2")
	if err != nil {
		return nil, err
	}
	return &GstDashDemux2{GstAdaptiveDemux2: &GstAdaptiveDemux2{Bin: &gst.Bin{Element: e}}}, nil
}

func NewDashDemux2WithName(name string) (*GstDashDemux2, error) {
	e, err := gst.NewElementWithName("dashdemux2", name)
	if err != nil {
		return nil, err
	}
	return &GstDashDemux2{GstAdaptiveDemux2: &GstAdaptiveDemux2{Bin: &gst.Bin{Element: e}}}, nil
}

// Max video framerate to select (0/1 = no maximum)
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstDashDemux2) SetMaxVideoFramerate(value interface{}) error {
	return e.Element.SetProperty("max-video-framerate", value)
}

func (e *GstDashDemux2) GetMaxVideoFramerate() (interface{}, error) {
	return e.Element.GetProperty("max-video-framerate")
}

// Max video height to select (0 = no maximum)
// Default: 0, Min: 0, Max: -1
func (e *GstDashDemux2) SetMaxVideoHeight(value uint) error {
	return e.Element.SetProperty("max-video-height", value)
}

func (e *GstDashDemux2) GetMaxVideoHeight() (uint, error) {
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
func (e *GstDashDemux2) SetMaxVideoWidth(value uint) error {
	return e.Element.SetProperty("max-video-width", value)
}

func (e *GstDashDemux2) GetMaxVideoWidth() (uint, error) {
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

// Default presentation delay (in seconds, milliseconds or fragments) (e.g. 12s, 2500ms, 3f)
// Default: 10s
func (e *GstDashDemux2) SetPresentationDelay(value string) error {
	return e.Element.SetProperty("presentation-delay", value)
}

func (e *GstDashDemux2) GetPresentationDelay() (string, error) {
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

// Initial bitrate to use to choose first alternate (0 = automatic) (bits/s)
// Default: 0, Min: 0, Max: -1
func (e *GstDashDemux2) SetStartBitrate(value uint) error {
	return e.Element.SetProperty("start-bitrate", value)
}

func (e *GstDashDemux2) GetStartBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("start-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max of bitrate supported by target video decoder (0 = no maximum)
// Default: 0, Min: 0, Max: -1
func (e *GstDashDemux2) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstDashDemux2) GetMaxBitrate() (uint, error) {
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

// HTTP Live Streaming demuxer
type GstHLSDemux2 struct {
	*GstAdaptiveDemux2
}

func NewHLSDemux2() (*GstHLSDemux2, error) {
	e, err := gst.NewElement("hlsdemux2")
	if err != nil {
		return nil, err
	}
	return &GstHLSDemux2{GstAdaptiveDemux2: &GstAdaptiveDemux2{Bin: &gst.Bin{Element: e}}}, nil
}

func NewHLSDemux2WithName(name string) (*GstHLSDemux2, error) {
	e, err := gst.NewElementWithName("hlsdemux2", name)
	if err != nil {
		return nil, err
	}
	return &GstHLSDemux2{GstAdaptiveDemux2: &GstAdaptiveDemux2{Bin: &gst.Bin{Element: e}}}, nil
}

// Initial bitrate to use to choose first alternate (0 = automatic) (bits/s)
// Default: 0, Min: 0, Max: -1
func (e *GstHLSDemux2) SetStartBitrate(value uint) error {
	return e.Element.SetProperty("start-bitrate", value)
}

func (e *GstHLSDemux2) GetStartBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("start-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Parse and demultiplex a Smooth Streaming manifest into audio and video streams
type GstMssDemux2 struct {
	*GstAdaptiveDemux2
}

func NewMssDemux2() (*GstMssDemux2, error) {
	e, err := gst.NewElement("mssdemux2")
	if err != nil {
		return nil, err
	}
	return &GstMssDemux2{GstAdaptiveDemux2: &GstAdaptiveDemux2{Bin: &gst.Bin{Element: e}}}, nil
}

func NewMssDemux2WithName(name string) (*GstMssDemux2, error) {
	e, err := gst.NewElementWithName("mssdemux2", name)
	if err != nil {
		return nil, err
	}
	return &GstMssDemux2{GstAdaptiveDemux2: &GstAdaptiveDemux2{Bin: &gst.Bin{Element: e}}}, nil
}

type GstAdaptiveDemux2 struct {
	*gst.Bin
}

// Network connection speed to use (0 = automatic) (bits/s)
// Default: 0, Min: 0, Max: -1
func (e *GstAdaptiveDemux2) SetConnectionBitrate(value uint) error {
	return e.Element.SetProperty("connection-bitrate", value)
}

func (e *GstAdaptiveDemux2) GetConnectionBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("connection-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Currently buffered level of video track(s) (ns)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAdaptiveDemux2) GetCurrentLevelTimeVideo() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-time-video")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Low watermark for parsed data below which downloads are resumed (in fragments, 0=disable)
// Default: 0, Min: 0, Max: 3.40282e+38
func (e *GstAdaptiveDemux2) SetLowWatermarkFragments(value float64) error {
	return e.Element.SetProperty("low-watermark-fragments", value)
}

func (e *GstAdaptiveDemux2) GetLowWatermarkFragments() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("low-watermark-fragments")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Low watermark for parsed data below which downloads are resumed (in ns, 0=automatic)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAdaptiveDemux2) SetLowWatermarkTime(value uint64) error {
	return e.Element.SetProperty("low-watermark-time", value)
}

func (e *GstAdaptiveDemux2) GetLowWatermarkTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("low-watermark-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum bitrate to use when switching to alternates (bits/s)
// Default: 0, Min: 0, Max: -1
func (e *GstAdaptiveDemux2) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstAdaptiveDemux2) GetMaxBitrate() (uint, error) {
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

// Minimum bitrate to use when switching to alternates (bits/s)
// Default: 0, Min: 0, Max: -1
func (e *GstAdaptiveDemux2) SetMinBitrate(value uint) error {
	return e.Element.SetProperty("min-bitrate", value)
}

func (e *GstAdaptiveDemux2) GetMinBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Limit of the available bitrate to use when switching to alternates
// Default: 0.8, Min: 0, Max: 1
func (e *GstAdaptiveDemux2) SetBandwidthTargetRatio(value float32) error {
	return e.Element.SetProperty("bandwidth-target-ratio", value)
}

func (e *GstAdaptiveDemux2) GetBandwidthTargetRatio() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("bandwidth-target-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Network connection speed to use in kbps (0 = calculate from downloaded fragments)
// Default: 0, Min: 0, Max: 4294967
func (e *GstAdaptiveDemux2) SetConnectionSpeed(value uint) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstAdaptiveDemux2) GetConnectionSpeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Report of current download bandwidth (based on arriving data) (bits/s)
// Default: 0, Min: 0, Max: -1
func (e *GstAdaptiveDemux2) GetCurrentBandwidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Currently buffered level of audio track(s) (ns)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstAdaptiveDemux2) GetCurrentLevelTimeAudio() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-time-audio")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// High watermark for parsed data above which downloads are paused (in fragments, 0=disable)
// Default: 0, Min: 0, Max: 3.40282e+38
func (e *GstAdaptiveDemux2) SetHighWatermarkFragments(value float64) error {
	return e.Element.SetProperty("high-watermark-fragments", value)
}

func (e *GstAdaptiveDemux2) GetHighWatermarkFragments() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("high-watermark-fragments")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// High watermark for parsed data above which downloads are paused (in ns, 0=disable)
// Default: 30000000000, Min: 0, Max: 18446744073709551615
func (e *GstAdaptiveDemux2) SetHighWatermarkTime(value uint64) error {
	return e.Element.SetProperty("high-watermark-time", value)
}

func (e *GstAdaptiveDemux2) GetHighWatermarkTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("high-watermark-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Upper limit on the high watermark for parsed data, above which downloads are paused (in ns, 0=disable)
// Default: 30000000000, Min: 0, Max: 18446744073709551615
func (e *GstAdaptiveDemux2) SetMaxBufferingTime(value uint64) error {
	return e.Element.SetProperty("max-buffering-time", value)
}

func (e *GstAdaptiveDemux2) GetMaxBufferingTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-buffering-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}
