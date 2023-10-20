// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Download and buffer a URI as needed
type GstURISourceBin struct {
	*gst.Bin
}

func NewURISourceBin() (*GstURISourceBin, error) {
	e, err := gst.NewElement("urisourcebin")
	if err != nil {
		return nil, err
	}
	return &GstURISourceBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewURISourceBinWithName(name string) (*GstURISourceBin, error) {
	e, err := gst.NewElementWithName("urisourcebin", name)
	if err != nil {
		return nil, err
	}
	return &GstURISourceBin{Bin: &gst.Bin{Element: e}}, nil
}

// Source object used

func (e *GstURISourceBin) GetSource() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("source")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Perform buffering on demuxed/parsed media
// Default: true
func (e *GstURISourceBin) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

func (e *GstURISourceBin) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The directory where buffers are downloaded to, if 'download' is enabled. If not set (default), the XDG cache directory is used.
// Default: NULL
func (e *GstURISourceBin) SetDownloadDir(value string) error {
	return e.Element.SetProperty("download-dir", value)
}

func (e *GstURISourceBin) GetDownloadDir() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("download-dir")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Low threshold for buffering to start. Only used if use-buffering is True
// Default: 0.01, Min: 0, Max: 1
func (e *GstURISourceBin) SetLowWatermark(value float64) error {
	return e.Element.SetProperty("low-watermark", value)
}

func (e *GstURISourceBin) GetLowWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("low-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Extract the elementary streams of non-raw sources
// Default: false
func (e *GstURISourceBin) SetParseStreams(value bool) error {
	return e.Element.SetProperty("parse-streams", value)
}

func (e *GstURISourceBin) GetParseStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("parse-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Attempt download buffering when buffering network streams
// Default: false
func (e *GstURISourceBin) SetDownload(value bool) error {
	return e.Element.SetProperty("download", value)
}

func (e *GstURISourceBin) GetDownload() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("download")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// High threshold for buffering to finish. Only used if use-buffering is True
// Default: 0.6, Min: 0, Max: 1
func (e *GstURISourceBin) SetHighWatermark(value float64) error {
	return e.Element.SetProperty("high-watermark", value)
}

func (e *GstURISourceBin) GetHighWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("high-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Max. amount of data in the ring buffer (bytes, 0 = ring buffer disabled)
// Default: 0, Min: 0, Max: 4294967295
func (e *GstURISourceBin) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

func (e *GstURISourceBin) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// A set of statistics over all the queue-like elements contained in this element
// Default: application/x-urisourcebin-stats, minimum-byte-level=(uint)0, maximum-byte-level=(uint)0, average-byte-level=(uint)0, minimum-time-level=(guint64)0, maximum-time-level=(guint64)0, average-time-level=(guint64)0;
func (e *GstURISourceBin) GetStatistics() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("statistics")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// URI to decode
// Default: NULL
func (e *GstURISourceBin) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstURISourceBin) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Buffer duration when buffering streams (-1 default value)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstURISourceBin) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

func (e *GstURISourceBin) GetBufferDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Buffer size when buffering streams (-1 default value)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstURISourceBin) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstURISourceBin) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstURISourceBin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstURISourceBin) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Autoplug and decode to raw media
type GstDecodebin3 struct {
	*gst.Bin
}

func NewDecodebin3() (*GstDecodebin3, error) {
	e, err := gst.NewElement("decodebin3")
	if err != nil {
		return nil, err
	}
	return &GstDecodebin3{Bin: &gst.Bin{Element: e}}, nil
}

func NewDecodebin3WithName(name string) (*GstDecodebin3, error) {
	e, err := gst.NewElementWithName("decodebin3", name)
	if err != nil {
		return nil, err
	}
	return &GstDecodebin3{Bin: &gst.Bin{Element: e}}, nil
}

// The caps on which to stop decoding. (NULL = default)
// Default: video/x-raw(ANY); audio/x-raw(ANY); text/x-raw(ANY); subpicture/x-dvd; subpicture/x-dvb; subpicture/x-xsub; subpicture/x-pgs; closedcaption/x-cea-608; closedcaption/x-cea-708; application/x-onvif-metadata
func (e *GstDecodebin3) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstDecodebin3) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Autoplug and play media from an uri
type GstPlayBin struct {
	*gst.Pipeline
}

func NewPlayBin() (*GstPlayBin, error) {
	e, err := gst.NewElement("playbin")
	if err != nil {
		return nil, err
	}
	return &GstPlayBin{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

func NewPlayBinWithName(name string) (*GstPlayBin, error) {
	e, err := gst.NewElementWithName("playbin", name)
	if err != nil {
		return nil, err
	}
	return &GstPlayBin{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

// Currently playing text stream (-1 = auto)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstPlayBin) SetCurrentText(value int) error {
	return e.Element.SetProperty("current-text", value)
}

func (e *GstPlayBin) GetCurrentText() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("current-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Mute the audio channel without changing the volume
// Default: false
func (e *GstPlayBin) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstPlayBin) GetMute() (bool, error) {
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

// the text output element to use (NULL = default subtitleoverlay)

func (e *GstPlayBin) SetTextSink(value *gst.Element) error {
	return e.Element.SetProperty("text-sink", value)
}

func (e *GstPlayBin) GetTextSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("text-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Current audio stream combiner (NULL = input-selector)

func (e *GstPlayBin) SetAudioStreamCombiner(value *gst.Element) error {
	return e.Element.SetProperty("audio-stream-combiner", value)
}

func (e *GstPlayBin) GetAudioStreamCombiner() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-stream-combiner")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstPlayBin) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstPlayBin) GetForceAspectRatio() (bool, error) {
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

// Total number of audio streams
// Default: 0, Min: 0, Max: 2147483647
func (e *GstPlayBin) GetNAudio() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("n-audio")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pango font description of font to be used for subtitle rendering
// Default: NULL
func (e *GstPlayBin) SetSubtitleFontDesc(value string) error {
	return e.Element.SetProperty("subtitle-font-desc", value)
}

// URI of the media to play
// Default: NULL
func (e *GstPlayBin) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstPlayBin) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// the audio output element to use (NULL = default sink)

func (e *GstPlayBin) SetAudioSink(value *gst.Element) error {
	return e.Element.SetProperty("audio-sink", value)
}

func (e *GstPlayBin) GetAudioSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The synchronisation offset between audio and video in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstPlayBin) SetAvOffset(value int64) error {
	return e.Element.SetProperty("av-offset", value)
}

func (e *GstPlayBin) GetAvOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("av-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The currently playing URI of a subtitle
// Default: NULL
func (e *GstPlayBin) GetCurrentSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Total number of text streams
// Default: 0, Min: 0, Max: 2147483647
func (e *GstPlayBin) GetNText() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("n-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// the video filter(s) to apply, if possible

func (e *GstPlayBin) SetVideoFilter(value *gst.Element) error {
	return e.Element.SetProperty("video-filter", value)
}

func (e *GstPlayBin) GetVideoFilter() (*gst.Element, error) {
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

// The synchronisation offset between text and video in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstPlayBin) SetTextOffset(value int64) error {
	return e.Element.SetProperty("text-offset", value)
}

func (e *GstPlayBin) GetTextOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("text-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// the visualization element to use (NULL = default)

func (e *GstPlayBin) SetVisPlugin(value *gst.Element) error {
	return e.Element.SetProperty("vis-plugin", value)
}

func (e *GstPlayBin) GetVisPlugin() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("vis-plugin")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// the audio filter(s) to apply, if possible

func (e *GstPlayBin) SetAudioFilter(value *gst.Element) error {
	return e.Element.SetProperty("audio-filter", value)
}

func (e *GstPlayBin) GetAudioFilter() (*gst.Element, error) {
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

// Buffer duration when buffering network streams
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstPlayBin) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

func (e *GstPlayBin) GetBufferDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Buffer size when buffering network streams
// Default: -1, Min: -1, Max: 2147483647
func (e *GstPlayBin) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstPlayBin) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstPlayBin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstPlayBin) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Source element

func (e *GstPlayBin) GetSource() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("source")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Override details of the multiview frame layout
// Default: none
func (e *GstPlayBin) SetVideoMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("video-multiview-flags", value)
}

func (e *GstPlayBin) GetVideoMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-flags")
}

// Current video stream combiner (NULL = input-selector)

func (e *GstPlayBin) SetVideoStreamCombiner(value *gst.Element) error {
	return e.Element.SetProperty("video-stream-combiner", value)
}

func (e *GstPlayBin) GetVideoStreamCombiner() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-stream-combiner")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// the video output element to use (NULL = default sink)

func (e *GstPlayBin) SetVideoSink(value *gst.Element) error {
	return e.Element.SetProperty("video-sink", value)
}

func (e *GstPlayBin) GetVideoSink() (*gst.Element, error) {
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

// The currently playing URI
// Default: NULL
func (e *GstPlayBin) GetCurrentUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Currently playing video stream (-1 = auto)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstPlayBin) SetCurrentVideo(value int) error {
	return e.Element.SetProperty("current-video", value)
}

func (e *GstPlayBin) GetCurrentVideo() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("current-video")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Flags to control behaviour
// Default: soft-colorbalance+deinterlace+soft-volume+text+audio+video
func (e *GstPlayBin) SetFlags(value GstPlayFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstPlayBin) GetFlags() (GstPlayFlags, error) {
	var value GstPlayFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstPlayFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstPlayFlags")
	}
	return value, nil
}

// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstPlayBin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstPlayBin) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Re-interpret a video stream as one of several frame-packed stereoscopic modes.
// Default: none (-1)
func (e *GstPlayBin) SetVideoMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("video-multiview-mode", value)
}

func (e *GstPlayBin) GetVideoMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-mode")
}

// Currently playing audio stream (-1 = auto)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstPlayBin) SetCurrentAudio(value int) error {
	return e.Element.SetProperty("current-audio", value)
}

func (e *GstPlayBin) GetCurrentAudio() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("current-audio")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Max. amount of data in the ring buffer (bytes, 0 = ring buffer disabled)
// Default: 0, Min: 0, Max: 4294967295
func (e *GstPlayBin) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

func (e *GstPlayBin) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The last sample (NULL = no video available)

func (e *GstPlayBin) GetSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// Current text stream combiner (NULL = input-selector)

func (e *GstPlayBin) SetTextStreamCombiner(value *gst.Element) error {
	return e.Element.SetProperty("text-stream-combiner", value)
}

func (e *GstPlayBin) GetTextStreamCombiner() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("text-stream-combiner")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The audio volume, 1.0=100%%
// Default: 1, Min: 0, Max: 10
func (e *GstPlayBin) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstPlayBin) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Total number of video streams
// Default: 0, Min: 0, Max: 2147483647
func (e *GstPlayBin) GetNVideo() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("n-video")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Optional URI of a subtitle
// Default: NULL
func (e *GstPlayBin) SetSuburi(value string) error {
	return e.Element.SetProperty("suburi", value)
}

func (e *GstPlayBin) GetSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Convenience sink for multiple streams
type GstPlaySink struct {
	*gst.Bin
}

func NewPlaySink() (*GstPlaySink, error) {
	e, err := gst.NewElement("playsink")
	if err != nil {
		return nil, err
	}
	return &GstPlaySink{Bin: &gst.Bin{Element: e}}, nil
}

func NewPlaySinkWithName(name string) (*GstPlaySink, error) {
	e, err := gst.NewElementWithName("playsink", name)
	if err != nil {
		return nil, err
	}
	return &GstPlaySink{Bin: &gst.Bin{Element: e}}, nil
}

// the audio filter(s) to apply, if possible

func (e *GstPlaySink) SetAudioFilter(value *gst.Element) error {
	return e.Element.SetProperty("audio-filter", value)
}

func (e *GstPlaySink) GetAudioFilter() (*gst.Element, error) {
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

// Flags to control behaviour
// Default: soft-colorbalance+soft-volume+text+audio+video
func (e *GstPlaySink) SetFlags(value GstPlayFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstPlaySink) GetFlags() (GstPlayFlags, error) {
	var value GstPlayFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstPlayFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstPlayFlags")
	}
	return value, nil
}

// How to send events received in send_event function
// Default: default (0)
func (e *GstPlaySink) SetSendEventMode(value GstPlaySinkSendEventMode) error {
	e.Element.SetArg("send-event-mode", string(value))
	return nil
}

func (e *GstPlaySink) GetSendEventMode() (GstPlaySinkSendEventMode, error) {
	var value GstPlaySinkSendEventMode
	var ok bool
	v, err := e.Element.GetProperty("send-event-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstPlaySinkSendEventMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstPlaySinkSendEventMode")
	}
	return value, nil
}

// the audio output element to use (NULL = default sink)

func (e *GstPlaySink) SetAudioSink(value *gst.Element) error {
	return e.Element.SetProperty("audio-sink", value)
}

func (e *GstPlaySink) GetAudioSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The synchronisation offset between audio and video in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstPlaySink) SetAvOffset(value int64) error {
	return e.Element.SetProperty("av-offset", value)
}

func (e *GstPlaySink) GetAvOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("av-offset")
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
func (e *GstPlaySink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstPlaySink) GetForceAspectRatio() (bool, error) {
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

// The last sample (NULL = no video available)

func (e *GstPlaySink) GetSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// the text output element to use (NULL = default subtitleoverlay)

func (e *GstPlaySink) SetTextSink(value *gst.Element) error {
	return e.Element.SetProperty("text-sink", value)
}

func (e *GstPlaySink) GetTextSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("text-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// the video output element to use (NULL = default sink)

func (e *GstPlaySink) SetVideoSink(value *gst.Element) error {
	return e.Element.SetProperty("video-sink", value)
}

func (e *GstPlaySink) GetVideoSink() (*gst.Element, error) {
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

// The audio volume, 1.0=100%%
// Default: 1, Min: 0, Max: 10
func (e *GstPlaySink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstPlaySink) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Pango font description of font to be used for subtitle rendering
// Default: NULL
func (e *GstPlaySink) SetSubtitleFontDesc(value string) error {
	return e.Element.SetProperty("subtitle-font-desc", value)
}

// The synchronisation offset between text and video in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstPlaySink) SetTextOffset(value int64) error {
	return e.Element.SetProperty("text-offset", value)
}

func (e *GstPlaySink) GetTextOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("text-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Mute the audio channel without changing the volume
// Default: false
func (e *GstPlaySink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstPlaySink) GetMute() (bool, error) {
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

// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstPlaySink) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstPlaySink) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// the video filter(s) to apply, if possible

func (e *GstPlaySink) SetVideoFilter(value *gst.Element) error {
	return e.Element.SetProperty("video-filter", value)
}

func (e *GstPlaySink) GetVideoFilter() (*gst.Element, error) {
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

// the visualization element to use (NULL = default)

func (e *GstPlaySink) SetVisPlugin(value *gst.Element) error {
	return e.Element.SetProperty("vis-plugin", value)
}

func (e *GstPlaySink) GetVisPlugin() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("vis-plugin")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Synchronizes a group of streams to have equal durations and starting points
type GstStreamSynchronizer struct {
	*gst.Element
}

func NewStreamSynchronizer() (*GstStreamSynchronizer, error) {
	e, err := gst.NewElement("streamsynchronizer")
	if err != nil {
		return nil, err
	}
	return &GstStreamSynchronizer{Element: e}, nil
}

func NewStreamSynchronizerWithName(name string) (*GstStreamSynchronizer, error) {
	e, err := gst.NewElementWithName("streamsynchronizer", name)
	if err != nil {
		return nil, err
	}
	return &GstStreamSynchronizer{Element: e}, nil
}

// Overlays a video stream with subtitles
type GstSubtitleOverlay struct {
	*gst.Bin
}

func NewSubtitleOverlay() (*GstSubtitleOverlay, error) {
	e, err := gst.NewElement("subtitleoverlay")
	if err != nil {
		return nil, err
	}
	return &GstSubtitleOverlay{Bin: &gst.Bin{Element: e}}, nil
}

func NewSubtitleOverlayWithName(name string) (*GstSubtitleOverlay, error) {
	e, err := gst.NewElementWithName("subtitleoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstSubtitleOverlay{Bin: &gst.Bin{Element: e}}, nil
}

// Pango font description of font to be used for subtitle rendering
// Default: NULL
func (e *GstSubtitleOverlay) SetFontDesc(value string) error {
	return e.Element.SetProperty("font-desc", value)
}

func (e *GstSubtitleOverlay) GetFontDesc() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-desc")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Whether to show subtitles
// Default: false
func (e *GstSubtitleOverlay) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstSubtitleOverlay) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstSubtitleOverlay) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstSubtitleOverlay) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The synchronisation offset between text and video in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstSubtitleOverlay) SetSubtitleTsOffset(value int64) error {
	return e.Element.SetProperty("subtitle-ts-offset", value)
}

func (e *GstSubtitleOverlay) GetSubtitleTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("subtitle-ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Autoplug and decode an URI to raw media
type GstURIDecodeBin struct {
	*gst.Bin
}

func NewURIDecodeBin() (*GstURIDecodeBin, error) {
	e, err := gst.NewElement("uridecodebin")
	if err != nil {
		return nil, err
	}
	return &GstURIDecodeBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewURIDecodeBinWithName(name string) (*GstURIDecodeBin, error) {
	e, err := gst.NewElementWithName("uridecodebin", name)
	if err != nil {
		return nil, err
	}
	return &GstURIDecodeBin{Bin: &gst.Bin{Element: e}}, nil
}

// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstURIDecodeBin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstURIDecodeBin) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Perform buffering on demuxed/parsed media
// Default: false
func (e *GstURIDecodeBin) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

func (e *GstURIDecodeBin) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Buffer duration when buffering streams (-1 default value)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstURIDecodeBin) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

func (e *GstURIDecodeBin) GetBufferDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Buffer size when buffering streams (-1 default value)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstURIDecodeBin) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstURIDecodeBin) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Post stream-topology messages
// Default: false
func (e *GstURIDecodeBin) SetPostStreamTopology(value bool) error {
	return e.Element.SetProperty("post-stream-topology", value)
}

func (e *GstURIDecodeBin) GetPostStreamTopology() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-stream-topology")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Expose all streams, including those of unknown type or that don't match the 'caps' property
// Default: true
func (e *GstURIDecodeBin) SetExposeAllStreams(value bool) error {
	return e.Element.SetProperty("expose-all-streams", value)
}

func (e *GstURIDecodeBin) GetExposeAllStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("expose-all-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use only sofware decoders to process streams
// Default: false
func (e *GstURIDecodeBin) SetForceSwDecoders(value bool) error {
	return e.Element.SetProperty("force-sw-decoders", value)
}

func (e *GstURIDecodeBin) GetForceSwDecoders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-sw-decoders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Max. amount of data in the ring buffer (bytes, 0 = ring buffer disabled)
// Default: 0, Min: 0, Max: 4294967295
func (e *GstURIDecodeBin) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

func (e *GstURIDecodeBin) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Source object used

func (e *GstURIDecodeBin) GetSource() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("source")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// URI to decode
// Default: NULL
func (e *GstURIDecodeBin) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstURIDecodeBin) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The caps on which to stop decoding. (NULL = default)
// Default: video/x-raw(ANY); audio/x-raw(ANY); text/x-raw(ANY); subpicture/x-dvd; subpicture/x-dvb; subpicture/x-xsub; subpicture/x-pgs; closedcaption/x-cea-608; closedcaption/x-cea-708; application/x-onvif-metadata
func (e *GstURIDecodeBin) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstURIDecodeBin) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstURIDecodeBin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstURIDecodeBin) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Attempt download buffering when buffering network streams
// Default: false
func (e *GstURIDecodeBin) SetDownload(value bool) error {
	return e.Element.SetProperty("download", value)
}

func (e *GstURIDecodeBin) GetDownload() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("download")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Autoplug and decode an URI to raw media
type GstURIDecodeBin3 struct {
	*gst.Bin
}

func NewURIDecodeBin3() (*GstURIDecodeBin3, error) {
	e, err := gst.NewElement("uridecodebin3")
	if err != nil {
		return nil, err
	}
	return &GstURIDecodeBin3{Bin: &gst.Bin{Element: e}}, nil
}

func NewURIDecodeBin3WithName(name string) (*GstURIDecodeBin3, error) {
	e, err := gst.NewElementWithName("uridecodebin3", name)
	if err != nil {
		return nil, err
	}
	return &GstURIDecodeBin3{Bin: &gst.Bin{Element: e}}, nil
}

// Perform buffering on demuxed/parsed media
// Default: false
func (e *GstURIDecodeBin3) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

func (e *GstURIDecodeBin3) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The currently playing URI of a subtitle
// Default: NULL
func (e *GstURIDecodeBin3) GetCurrentSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The currently playing URI
// Default: NULL
func (e *GstURIDecodeBin3) GetCurrentUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The directory where buffers are downloaded to, if 'download' is enabled. If not set (default), the XDG cache directory is used.
// Default: NULL
func (e *GstURIDecodeBin3) SetDownloadDir(value string) error {
	return e.Element.SetProperty("download-dir", value)
}

func (e *GstURIDecodeBin3) GetDownloadDir() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("download-dir")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Max. amount of data in the ring buffer (bytes, 0 = ring buffer disabled)
// Default: 0, Min: 0, Max: 4294967295
func (e *GstURIDecodeBin3) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

func (e *GstURIDecodeBin3) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Optional URI of a subtitle
// Default: NULL
func (e *GstURIDecodeBin3) SetSuburi(value string) error {
	return e.Element.SetProperty("suburi", value)
}

func (e *GstURIDecodeBin3) GetSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Buffer duration when buffering streams (-1 default value)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstURIDecodeBin3) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

func (e *GstURIDecodeBin3) GetBufferDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Buffer size when buffering streams (-1 default value)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstURIDecodeBin3) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstURIDecodeBin3) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The caps on which to stop decoding. (NULL = default)
// Default: video/x-raw(ANY); audio/x-raw(ANY); text/x-raw(ANY); subpicture/x-dvd; subpicture/x-dvb; subpicture/x-xsub; subpicture/x-pgs; closedcaption/x-cea-608; closedcaption/x-cea-708; application/x-onvif-metadata
func (e *GstURIDecodeBin3) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstURIDecodeBin3) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstURIDecodeBin3) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstURIDecodeBin3) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Attempt download buffering when buffering network streams
// Default: false
func (e *GstURIDecodeBin3) SetDownload(value bool) error {
	return e.Element.SetProperty("download", value)
}

func (e *GstURIDecodeBin3) GetDownload() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("download")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, URI changes are applied immediately
// Default: false
func (e *GstURIDecodeBin3) SetInstantUri(value bool) error {
	return e.Element.SetProperty("instant-uri", value)
}

func (e *GstURIDecodeBin3) GetInstantUri() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("instant-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Source object used

func (e *GstURIDecodeBin3) GetSource() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("source")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// URI to decode
// Default: NULL
func (e *GstURIDecodeBin3) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstURIDecodeBin3) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Autoplug and decode to raw media
type GstDecodeBin struct {
	*gst.Bin
}

func NewDecodeBin() (*GstDecodeBin, error) {
	e, err := gst.NewElement("decodebin")
	if err != nil {
		return nil, err
	}
	return &GstDecodeBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewDecodeBinWithName(name string) (*GstDecodeBin, error) {
	e, err := gst.NewElementWithName("decodebin", name)
	if err != nil {
		return nil, err
	}
	return &GstDecodeBin{Bin: &gst.Bin{Element: e}}, nil
}

// Use only sofware decoders to process streams
// Default: false
func (e *GstDecodeBin) SetForceSwDecoders(value bool) error {
	return e.Element.SetProperty("force-sw-decoders", value)
}

func (e *GstDecodeBin) GetForceSwDecoders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-sw-decoders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Low threshold for buffering to start
// Default: 10, Min: 0, Max: 100
func (e *GstDecodeBin) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

func (e *GstDecodeBin) GetLowPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("low-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Max. number of buffers in the queue (0=automatic)
// Default: 0, Min: 0, Max: -1
func (e *GstDecodeBin) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

func (e *GstDecodeBin) GetMaxSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (in ns, 0=automatic)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstDecodeBin) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstDecodeBin) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The caps of the input data. (NULL = use typefind element)

func (e *GstDecodeBin) SetSinkCaps(value *gst.Caps) error {
	return e.Element.SetProperty("sink-caps", value)
}

func (e *GstDecodeBin) GetSinkCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("sink-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Post stream-topology messages
// Default: false
func (e *GstDecodeBin) SetPostStreamTopology(value bool) error {
	return e.Element.SetProperty("post-stream-topology", value)
}

func (e *GstDecodeBin) GetPostStreamTopology() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-stream-topology")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstDecodeBin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstDecodeBin) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Emit GST_MESSAGE_BUFFERING based on low-/high-percent thresholds
// Default: false
func (e *GstDecodeBin) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

func (e *GstDecodeBin) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The caps on which to stop decoding.
// Default: video/x-raw(ANY); audio/x-raw(ANY); text/x-raw(ANY); subpicture/x-dvd; subpicture/x-dvb; subpicture/x-xsub; subpicture/x-pgs; closedcaption/x-cea-608; closedcaption/x-cea-708; application/x-onvif-metadata
func (e *GstDecodeBin) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstDecodeBin) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstDecodeBin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstDecodeBin) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Expose all streams, including those of unknown type or that don't match the 'caps' property
// Default: true
func (e *GstDecodeBin) SetExposeAllStreams(value bool) error {
	return e.Element.SetProperty("expose-all-streams", value)
}

func (e *GstDecodeBin) GetExposeAllStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("expose-all-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// High threshold for buffering to finish
// Default: 99, Min: 0, Max: 100
func (e *GstDecodeBin) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

func (e *GstDecodeBin) GetHighPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("high-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Max. amount of bytes in the queue (0=automatic)
// Default: 0, Min: 0, Max: -1
func (e *GstDecodeBin) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

func (e *GstDecodeBin) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Parse and de-multiplex to elementary stream
type GstParseBin struct {
	*gst.Bin
}

func NewParseBin() (*GstParseBin, error) {
	e, err := gst.NewElement("parsebin")
	if err != nil {
		return nil, err
	}
	return &GstParseBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewParseBinWithName(name string) (*GstParseBin, error) {
	e, err := gst.NewElementWithName("parsebin", name)
	if err != nil {
		return nil, err
	}
	return &GstParseBin{Bin: &gst.Bin{Element: e}}, nil
}

// The caps of the input data. (NULL = use typefind element)

func (e *GstParseBin) SetSinkCaps(value *gst.Caps) error {
	return e.Element.SetProperty("sink-caps", value)
}

func (e *GstParseBin) GetSinkCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("sink-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstParseBin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstParseBin) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstParseBin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstParseBin) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Expose all streams, including those of unknown type or that don't match the 'caps' property
// Default: true
func (e *GstParseBin) SetExposeAllStreams(value bool) error {
	return e.Element.SetProperty("expose-all-streams", value)
}

func (e *GstParseBin) GetExposeAllStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("expose-all-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Autoplug and play media from an uri
type GstPlayBin3 struct {
	*gst.Pipeline
}

func NewPlayBin3() (*GstPlayBin3, error) {
	e, err := gst.NewElement("playbin3")
	if err != nil {
		return nil, err
	}
	return &GstPlayBin3{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

func NewPlayBin3WithName(name string) (*GstPlayBin3, error) {
	e, err := gst.NewElementWithName("playbin3", name)
	if err != nil {
		return nil, err
	}
	return &GstPlayBin3{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

// the audio output element to use (NULL = default sink)

func (e *GstPlayBin3) SetAudioSink(value *gst.Element) error {
	return e.Element.SetProperty("audio-sink", value)
}

func (e *GstPlayBin3) GetAudioSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Optional URI of a subtitle
// Default: NULL
func (e *GstPlayBin3) SetSuburi(value string) error {
	return e.Element.SetProperty("suburi", value)
}

func (e *GstPlayBin3) GetSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// URI of the media to play
// Default: NULL
func (e *GstPlayBin3) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstPlayBin3) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// the video filter(s) to apply, if possible

func (e *GstPlayBin3) SetVideoFilter(value *gst.Element) error {
	return e.Element.SetProperty("video-filter", value)
}

func (e *GstPlayBin3) GetVideoFilter() (*gst.Element, error) {
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

// Pango font description of font to be used for subtitle rendering
// Default: NULL
func (e *GstPlayBin3) SetSubtitleFontDesc(value string) error {
	return e.Element.SetProperty("subtitle-font-desc", value)
}

// The synchronisation offset between text and video in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstPlayBin3) SetTextOffset(value int64) error {
	return e.Element.SetProperty("text-offset", value)
}

func (e *GstPlayBin3) GetTextOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("text-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Buffer duration when buffering network streams
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstPlayBin3) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

func (e *GstPlayBin3) GetBufferDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The last sample (NULL = no video available)

func (e *GstPlayBin3) GetSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// Current text stream combiner (default: none)

func (e *GstPlayBin3) SetTextStreamCombiner(value *gst.Element) error {
	return e.Element.SetProperty("text-stream-combiner", value)
}

func (e *GstPlayBin3) GetTextStreamCombiner() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("text-stream-combiner")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// the visualization element to use (NULL = default)

func (e *GstPlayBin3) SetVisPlugin(value *gst.Element) error {
	return e.Element.SetProperty("vis-plugin", value)
}

func (e *GstPlayBin3) GetVisPlugin() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("vis-plugin")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The synchronisation offset between audio and video in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstPlayBin3) SetAvOffset(value int64) error {
	return e.Element.SetProperty("av-offset", value)
}

func (e *GstPlayBin3) GetAvOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("av-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The currently playing URI of a subtitle
// Default: NULL
func (e *GstPlayBin3) GetCurrentSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The currently playing URI
// Default: NULL
func (e *GstPlayBin3) GetCurrentUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstPlayBin3) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstPlayBin3) GetForceAspectRatio() (bool, error) {
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

// When enabled, URI changes are applied immediately
// Default: false
func (e *GstPlayBin3) SetInstantUri(value bool) error {
	return e.Element.SetProperty("instant-uri", value)
}

func (e *GstPlayBin3) GetInstantUri() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("instant-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Override details of the multiview frame layout
// Default: none
func (e *GstPlayBin3) SetVideoMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("video-multiview-flags", value)
}

func (e *GstPlayBin3) GetVideoMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-flags")
}

// The audio volume, 1.0=100%%
// Default: 1, Min: 0, Max: 10
func (e *GstPlayBin3) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstPlayBin3) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// the audio filter(s) to apply, if possible

func (e *GstPlayBin3) SetAudioFilter(value *gst.Element) error {
	return e.Element.SetProperty("audio-filter", value)
}

func (e *GstPlayBin3) GetAudioFilter() (*gst.Element, error) {
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

// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.
// Default: NULL
func (e *GstPlayBin3) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

func (e *GstPlayBin3) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// the text output element to use (NULL = default subtitleoverlay)

func (e *GstPlayBin3) SetTextSink(value *gst.Element) error {
	return e.Element.SetProperty("text-sink", value)
}

func (e *GstPlayBin3) GetTextSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("text-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Current audio stream combiner (default: none))

func (e *GstPlayBin3) SetAudioStreamCombiner(value *gst.Element) error {
	return e.Element.SetProperty("audio-stream-combiner", value)
}

func (e *GstPlayBin3) GetAudioStreamCombiner() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-stream-combiner")
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
// Default: soft-colorbalance+deinterlace+buffering+soft-volume+text+audio+video
func (e *GstPlayBin3) SetFlags(value GstPlayFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstPlayBin3) GetFlags() (GstPlayFlags, error) {
	var value GstPlayFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstPlayFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstPlayFlags")
	}
	return value, nil
}

// Re-interpret a video stream as one of several frame-packed stereoscopic modes.
// Default: none (-1)
func (e *GstPlayBin3) SetVideoMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("video-multiview-mode", value)
}

func (e *GstPlayBin3) GetVideoMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-mode")
}

// the video output element to use (NULL = default sink)

func (e *GstPlayBin3) SetVideoSink(value *gst.Element) error {
	return e.Element.SetProperty("video-sink", value)
}

func (e *GstPlayBin3) GetVideoSink() (*gst.Element, error) {
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

// Max. amount of data in the ring buffer (bytes, 0 = ring buffer disabled)
// Default: 0, Min: 0, Max: 4294967295
func (e *GstPlayBin3) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

func (e *GstPlayBin3) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Current video stream combiner (default: none)

func (e *GstPlayBin3) SetVideoStreamCombiner(value *gst.Element) error {
	return e.Element.SetProperty("video-stream-combiner", value)
}

func (e *GstPlayBin3) GetVideoStreamCombiner() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-stream-combiner")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Buffer size when buffering network streams
// Default: -1, Min: -1, Max: 2147483647
func (e *GstPlayBin3) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstPlayBin3) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstPlayBin3) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstPlayBin3) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Mute the audio channel without changing the volume
// Default: false
func (e *GstPlayBin3) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstPlayBin3) GetMute() (bool, error) {
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

type GstAutoplugSelectResult string

const (
	GstAutoplugSelectResultTry    GstAutoplugSelectResult = "try"    // GST_AUTOPLUG_SELECT_TRY
	GstAutoplugSelectResultExpose GstAutoplugSelectResult = "expose" // GST_AUTOPLUG_SELECT_EXPOSE
	GstAutoplugSelectResultSkip   GstAutoplugSelectResult = "skip"   // GST_AUTOPLUG_SELECT_SKIP
)

type GstPlayFlags int

const (
	GstPlayFlagsVideo            GstPlayFlags = 0x00000001 // Render the video stream
	GstPlayFlagsAudio            GstPlayFlags = 0x00000002 // Render the audio stream
	GstPlayFlagsText             GstPlayFlags = 0x00000004 // Render subtitles
	GstPlayFlagsVis              GstPlayFlags = 0x00000008 // Render visualisation when no video is present
	GstPlayFlagsSoftVolume       GstPlayFlags = 0x00000010 // Use software volume
	GstPlayFlagsNativeAudio      GstPlayFlags = 0x00000020 // Only use native audio formats
	GstPlayFlagsNativeVideo      GstPlayFlags = 0x00000040 // Only use native video formats
	GstPlayFlagsDownload         GstPlayFlags = 0x00000080 // Attempt progressive download buffering
	GstPlayFlagsBuffering        GstPlayFlags = 0x00000100 // Buffer demuxed/parsed data
	GstPlayFlagsDeinterlace      GstPlayFlags = 0x00000200 // Deinterlace video if necessary
	GstPlayFlagsSoftColorbalance GstPlayFlags = 0x00000400 // Use software color balance
	GstPlayFlagsForceFilters     GstPlayFlags = 0x00000800 // Force audio/video filter(s) to be applied
	GstPlayFlagsForceSwDecoders  GstPlayFlags = 0x00001000 // Force only software-based decoders (no effect for playbin3)
)

type GstPlaySinkSendEventMode string

const (
	GstPlaySinkSendEventModeDefault GstPlaySinkSendEventMode = "default" // Default GstBin's send_event handling (default)
	GstPlaySinkSendEventModeFirst   GstPlaySinkSendEventMode = "first"   // Sends the event to sinks until the first one handles it
)
