// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// HTTP Live Streaming demuxer
type GstHLSDemux struct {
	*GstAdaptiveDemux
}

func NewHLSDemux() (*GstHLSDemux, error) {
	e, err := gst.NewElement("hlsdemux")
	if err != nil {
		return nil, err
	}
	return &GstHLSDemux{GstAdaptiveDemux: &GstAdaptiveDemux{Bin: &gst.Bin{Element: e}}}, nil
}

func NewHLSDemuxWithName(name string) (*GstHLSDemux, error) {
	e, err := gst.NewElementWithName("hlsdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstHLSDemux{GstAdaptiveDemux: &GstAdaptiveDemux{Bin: &gst.Bin{Element: e}}}, nil
}

// HTTP Live Streaming sink
type GstHlsSink struct {
	*gst.Bin
}

func NewHlsSink() (*GstHlsSink, error) {
	e, err := gst.NewElement("hlssink")
	if err != nil {
		return nil, err
	}
	return &GstHlsSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewHlsSinkWithName(name string) (*GstHlsSink, error) {
	e, err := gst.NewElementWithName("hlssink", name)
	if err != nil {
		return nil, err
	}
	return &GstHlsSink{Bin: &gst.Bin{Element: e}}, nil
}

// The target duration in seconds of a segment/file. (0 - disabled, useful for management of segment duration by the streaming server)
// Default: 15, Min: 0, Max: -1
func (e *GstHlsSink) SetTargetDuration(value uint) error {
	return e.Element.SetProperty("target-duration", value)
}

func (e *GstHlsSink) GetTargetDuration() (uint, error) {
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

// Location of the file to write
// Default: segment%%05d.ts
func (e *GstHlsSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstHlsSink) GetLocation() (string, error) {
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

// Maximum number of files to keep on disk. Once the maximum is reached,old files start to be deleted to make room for new ones.
// Default: 10, Min: 0, Max: -1
func (e *GstHlsSink) SetMaxFiles(value uint) error {
	return e.Element.SetProperty("max-files", value)
}

func (e *GstHlsSink) GetMaxFiles() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-files")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Length of HLS playlist. To allow players to conform to section 6.3.3 of the HLS specification, this should be at least 3. If set to 0, the playlist will be infinite.
// Default: 5, Min: 0, Max: -1
func (e *GstHlsSink) SetPlaylistLength(value uint) error {
	return e.Element.SetProperty("playlist-length", value)
}

func (e *GstHlsSink) GetPlaylistLength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("playlist-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Location of the playlist to write
// Default: playlist.m3u8
func (e *GstHlsSink) SetPlaylistLocation(value string) error {
	return e.Element.SetProperty("playlist-location", value)
}

func (e *GstHlsSink) GetPlaylistLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("playlist-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Location of the playlist to write
// Default: NULL
func (e *GstHlsSink) SetPlaylistRoot(value string) error {
	return e.Element.SetProperty("playlist-root", value)
}

func (e *GstHlsSink) GetPlaylistRoot() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("playlist-root")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// HTTP Live Streaming sink
type GstHlsSink2 struct {
	*gst.Bin
}

func NewHlsSink2() (*GstHlsSink2, error) {
	e, err := gst.NewElement("hlssink2")
	if err != nil {
		return nil, err
	}
	return &GstHlsSink2{Bin: &gst.Bin{Element: e}}, nil
}

func NewHlsSink2WithName(name string) (*GstHlsSink2, error) {
	e, err := gst.NewElementWithName("hlssink2", name)
	if err != nil {
		return nil, err
	}
	return &GstHlsSink2{Bin: &gst.Bin{Element: e}}, nil
}

// Location of the playlist to write
// Default: playlist.m3u8
func (e *GstHlsSink2) SetPlaylistLocation(value string) error {
	return e.Element.SetProperty("playlist-location", value)
}

func (e *GstHlsSink2) GetPlaylistLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("playlist-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Location of the playlist to write
// Default: NULL
func (e *GstHlsSink2) SetPlaylistRoot(value string) error {
	return e.Element.SetProperty("playlist-root", value)
}

func (e *GstHlsSink2) GetPlaylistRoot() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("playlist-root")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Send keyframe requests to ensure correct fragmentation. If this is disabled then the input must have keyframes in regular intervals
// Default: true
func (e *GstHlsSink2) SetSendKeyframeRequests(value bool) error {
	return e.Element.SetProperty("send-keyframe-requests", value)
}

func (e *GstHlsSink2) GetSendKeyframeRequests() (bool, error) {
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

// The target duration in seconds of a segment/file. (0 - disabled, useful for management of segment duration by the streaming server)
// Default: 15, Min: 0, Max: -1
func (e *GstHlsSink2) SetTargetDuration(value uint) error {
	return e.Element.SetProperty("target-duration", value)
}

func (e *GstHlsSink2) GetTargetDuration() (uint, error) {
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

// Location of the file to write
// Default: segment%%05d.ts
func (e *GstHlsSink2) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstHlsSink2) GetLocation() (string, error) {
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

// Maximum number of files to keep on disk. Once the maximum is reached,old files start to be deleted to make room for new ones.
// Default: 10, Min: 0, Max: -1
func (e *GstHlsSink2) SetMaxFiles(value uint) error {
	return e.Element.SetProperty("max-files", value)
}

func (e *GstHlsSink2) GetMaxFiles() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-files")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Length of HLS playlist. To allow players to conform to section 6.3.3 of the HLS specification, this should be at least 3. If set to 0, the playlist will be infinite.
// Default: 5, Min: 0, Max: -1
func (e *GstHlsSink2) SetPlaylistLength(value uint) error {
	return e.Element.SetProperty("playlist-length", value)
}

func (e *GstHlsSink2) GetPlaylistLength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("playlist-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
