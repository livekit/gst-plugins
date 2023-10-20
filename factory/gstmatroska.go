// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Demuxes Matroska/WebM streams into video/audio/subtitles
type GstMatroskaDemux struct {
	*gst.Element
}

func NewMatroskaDemux() (*GstMatroskaDemux, error) {
	e, err := gst.NewElement("matroskademux")
	if err != nil {
		return nil, err
	}
	return &GstMatroskaDemux{Element: e}, nil
}

func NewMatroskaDemuxWithName(name string) (*GstMatroskaDemux, error) {
	e, err := gst.NewElementWithName("matroskademux", name)
	if err != nil {
		return nil, err
	}
	return &GstMatroskaDemux{Element: e}, nil
}

// Maximum backtrack distance in seconds when seeking without and index in pull mode and search for a keyframe (0 = disable backtracking).
// Default: 30, Min: 0, Max: -1
func (e *GstMatroskaDemux) SetMaxBacktrackDistance(value uint) error {
	return e.Element.SetProperty("max-backtrack-distance", value)
}

func (e *GstMatroskaDemux) GetMaxBacktrackDistance() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-backtrack-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The demuxer sends out segment events for skipping gaps longer than this (0 = disabled).
// Default: 2000000000, Min: 0, Max: 18446744073709551615
func (e *GstMatroskaDemux) SetMaxGapTime(value uint64) error {
	return e.Element.SetProperty("max-gap-time", value)
}

func (e *GstMatroskaDemux) GetMaxGapTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-gap-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Muxes video/audio/subtitle streams into a matroska stream
type GstMatroskaMux struct {
	*gst.Element
}

func NewMatroskaMux() (*GstMatroskaMux, error) {
	e, err := gst.NewElement("matroskamux")
	if err != nil {
		return nil, err
	}
	return &GstMatroskaMux{Element: e}, nil
}

func NewMatroskaMuxWithName(name string) (*GstMatroskaMux, error) {
	e, err := gst.NewElementWithName("matroskamux", name)
	if err != nil {
		return nil, err
	}
	return &GstMatroskaMux{Element: e}, nil
}

// The name the application that creates the matroska file.
// Default: GStreamer Matroska muxer
func (e *GstMatroskaMux) SetWritingApp(value string) error {
	return e.Element.SetProperty("writing-app", value)
}

func (e *GstMatroskaMux) GetWritingApp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("writing-app")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Date and time of creation. This will be used for the DateUTC field. NULL means that the current time will be used.

func (e *GstMatroskaMux) SetCreationTime(value interface{}) error {
	return e.Element.SetProperty("creation-time", value)
}

func (e *GstMatroskaMux) GetCreationTime() (interface{}, error) {
	return e.Element.GetProperty("creation-time")
}

// A new cluster will be created if its duration exceeds this value. 0 means no maximum duration.
// Default: 65535000000, Min: 0, Max: 9223372036854775807
func (e *GstMatroskaMux) SetMaxClusterDuration(value int64) error {
	return e.Element.SetProperty("max-cluster-duration", value)
}

func (e *GstMatroskaMux) GetMaxClusterDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-cluster-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// An index entry is created every so many nanoseconds.
// Default: 0, Min: 0, Max: 9223372036854775807
func (e *GstMatroskaMux) SetMinIndexInterval(value int64) error {
	return e.Element.SetProperty("min-index-interval", value)
}

func (e *GstMatroskaMux) GetMinIndexInterval() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("min-index-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Offsets all streams so that the earliest stream starts at 0.
// Default: false
func (e *GstMatroskaMux) SetOffsetToZero(value bool) error {
	return e.Element.SetProperty("offset-to-zero", value)
}

func (e *GstMatroskaMux) GetOffsetToZero() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("offset-to-zero")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// This parameter determines what Matroska features can be used.
// Default: 2, Min: 1, Max: 2
func (e *GstMatroskaMux) SetVersion(value int) error {
	return e.Element.SetProperty("version", value)
}

func (e *GstMatroskaMux) GetVersion() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("version")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// An offset to add to all clusters/blocks (in nanoseconds)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstMatroskaMux) SetClusterTimestampOffset(value uint64) error {
	return e.Element.SetProperty("cluster-timestamp-offset", value)
}

func (e *GstMatroskaMux) GetClusterTimestampOffset() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("cluster-timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Desired cluster duration as nanoseconds. A new cluster will be created irrespective of this property if a force key unit event is received. 0 means create a new cluster for each video keyframe or for each audio buffer in audio only streams.
// Default: 500000000, Min: 0, Max: 9223372036854775807
func (e *GstMatroskaMux) SetMinClusterDuration(value int64) error {
	return e.Element.SetProperty("min-cluster-duration", value)
}

func (e *GstMatroskaMux) GetMinClusterDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("min-cluster-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written.
// Default: false
func (e *GstMatroskaMux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *GstMatroskaMux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// TimecodeScale used to calculate the Raw Timecode of a Block
// Default: 1000000, Min: 1, Max: 1000000000
func (e *GstMatroskaMux) SetTimecodescale(value int64) error {
	return e.Element.SetProperty("timecodescale", value)
}

func (e *GstMatroskaMux) GetTimecodescale() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timecodescale")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Parses Matroska/WebM streams into video/audio/subtitles
type GstMatroskaParse struct {
	*gst.Element
}

func NewMatroskaParse() (*GstMatroskaParse, error) {
	e, err := gst.NewElement("matroskaparse")
	if err != nil {
		return nil, err
	}
	return &GstMatroskaParse{Element: e}, nil
}

func NewMatroskaParseWithName(name string) (*GstMatroskaParse, error) {
	e, err := gst.NewElementWithName("matroskaparse", name)
	if err != nil {
		return nil, err
	}
	return &GstMatroskaParse{Element: e}, nil
}

// Muxes video and audio streams into a WebM stream
type GstWebMMux struct {
	*GstMatroskaMux
}

func NewWebMMux() (*GstWebMMux, error) {
	e, err := gst.NewElement("webmmux")
	if err != nil {
		return nil, err
	}
	return &GstWebMMux{GstMatroskaMux: &GstMatroskaMux{Element: e}}, nil
}

func NewWebMMuxWithName(name string) (*GstWebMMux, error) {
	e, err := gst.NewElementWithName("webmmux", name)
	if err != nil {
		return nil, err
	}
	return &GstWebMMux{GstMatroskaMux: &GstMatroskaMux{Element: e}}, nil
}
