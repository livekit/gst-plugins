// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type Playbin struct {
	Element *gst.Element
}

func NewPlaybin() (*Playbin, error) {
	e, err := gst.NewElement("playbin")
	if err != nil {
		return nil, err
	}
	return &Playbin{Element: e}, nil
}

func NewPlaybinWithName(name string) (*Playbin, error) {
	e, err := gst.NewElementWithName("playbin", name)
	if err != nil {
		return nil, err
	}
	return &Playbin{Element: e}, nil
}

// ----- Properties -----

// audio-filter (GstElement)
//
// the audio filter(s) to apply, if possible

func (e *Playbin) GetAudioFilter() (interface{}, error) {
	return e.Element.GetProperty("audio-filter")
}

func (e *Playbin) SetAudioFilter(value interface{}) error {
	return e.Element.SetProperty("audio-filter", value)
}

// audio-sink (GstElement)
//
// the audio output element to use (NULL = default sink)

func (e *Playbin) GetAudioSink() (interface{}, error) {
	return e.Element.GetProperty("audio-sink")
}

func (e *Playbin) SetAudioSink(value interface{}) error {
	return e.Element.SetProperty("audio-sink", value)
}

// audio-stream-combiner (GstElement)
//
// Get or set the current audio stream combiner. By default, an input-selector
// is created and deleted as-needed.

func (e *Playbin) GetAudioStreamCombiner() (interface{}, error) {
	return e.Element.GetProperty("audio-stream-combiner")
}

func (e *Playbin) SetAudioStreamCombiner(value interface{}) error {
	return e.Element.SetProperty("audio-stream-combiner", value)
}

// av-offset (int64)
//
// Control the synchronisation offset between the audio and video streams.
// Positive values make the audio ahead of the video and negative values make
// the audio go behind the video.

func (e *Playbin) GetAvOffset() (int64, error) {
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

func (e *Playbin) SetAvOffset(value int64) error {
	return e.Element.SetProperty("av-offset", value)
}

// buffer-duration (int64)
//
// Buffer duration when buffering network streams

func (e *Playbin) GetBufferDuration() (int64, error) {
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

func (e *Playbin) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

// buffer-size (int)
//
// Buffer size when buffering network streams

func (e *Playbin) GetBufferSize() (int, error) {
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

func (e *Playbin) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Playbin) GetConnectionSpeed() (uint64, error) {
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

func (e *Playbin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// current-audio (int)
//
// Get or set the currently playing audio stream. By default the first audio
// stream with data is played.

func (e *Playbin) GetCurrentAudio() (int, error) {
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

func (e *Playbin) SetCurrentAudio(value int) error {
	return e.Element.SetProperty("current-audio", value)
}

// current-suburi (string)
//
// The currently playing subtitle uri.

func (e *Playbin) GetCurrentSuburi() (string, error) {
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

// current-text (int)
//
// Get or set the currently playing subtitle stream. By default the first
// subtitle stream with data is played.

func (e *Playbin) GetCurrentText() (int, error) {
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

func (e *Playbin) SetCurrentText(value int) error {
	return e.Element.SetProperty("current-text", value)
}

// current-uri (string)
//
// The currently playing uri.

func (e *Playbin) GetCurrentUri() (string, error) {
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

// current-video (int)
//
// Get or set the currently playing video stream. By default the first video
// stream with data is played.

func (e *Playbin) GetCurrentVideo() (int, error) {
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

func (e *Playbin) SetCurrentVideo(value int) error {
	return e.Element.SetProperty("current-video", value)
}

// flags (GstPlayFlags)
//
// Control the behaviour of playbin.

func (e *Playbin) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *Playbin) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Playbin) GetForceAspectRatio() (bool, error) {
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

func (e *Playbin) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// mute (bool)
//
// Mute the audio channel without changing the volume

func (e *Playbin) GetMute() (bool, error) {
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

func (e *Playbin) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// n-audio (int)
//
// Get the total number of available audio streams.

func (e *Playbin) GetNAudio() (int, error) {
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

// n-text (int)
//
// Get the total number of available subtitle streams.

func (e *Playbin) GetNText() (int, error) {
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

// n-video (int)
//
// Get the total number of available video streams.

func (e *Playbin) GetNVideo() (int, error) {
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

// ring-buffer-max-size (uint64)
//
// The maximum size of the ring buffer in bytes. If set to 0, the ring
// buffer is disabled. Default 0.

func (e *Playbin) GetRingBufferMaxSize() (uint64, error) {
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

func (e *Playbin) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

// sample (GstSample)
//
// Get the currently rendered or prerolled sample in the video sink.
// The GstCaps in the sample will describe the format of the buffer.

func (e *Playbin) GetSample() (*gst.Sample, error) {
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

// source (GstElement)
//
// Source element

func (e *Playbin) GetSource() (interface{}, error) {
	return e.Element.GetProperty("source")
}

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Playbin) GetSubtitleEncoding() (string, error) {
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

func (e *Playbin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

// subtitle-font-desc (string)
//
// Pango font description of font to be used for subtitle rendering

func (e *Playbin) GetSubtitleFontDesc() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-font-desc")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Playbin) SetSubtitleFontDesc(value string) error {
	return e.Element.SetProperty("subtitle-font-desc", value)
}

// suburi (string)
//
// Set the next subtitle URI that playbin will play. This property can be
// set from the about-to-finish signal to queue the next subtitle media file.

func (e *Playbin) GetSuburi() (string, error) {
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

func (e *Playbin) SetSuburi(value string) error {
	return e.Element.SetProperty("suburi", value)
}

// text-offset (int64)
//
// Control the synchronisation offset between the text and video streams.
// Positive values make the text ahead of the video and negative values make
// the text go behind the video.

func (e *Playbin) GetTextOffset() (int64, error) {
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

func (e *Playbin) SetTextOffset(value int64) error {
	return e.Element.SetProperty("text-offset", value)
}

// text-sink (GstElement)
//
// the text output element to use (NULL = default subtitleoverlay)

func (e *Playbin) GetTextSink() (interface{}, error) {
	return e.Element.GetProperty("text-sink")
}

func (e *Playbin) SetTextSink(value interface{}) error {
	return e.Element.SetProperty("text-sink", value)
}

// text-stream-combiner (GstElement)
//
// Get or set the current text stream combiner. By default, an input-selector
// is created and deleted as-needed.

func (e *Playbin) GetTextStreamCombiner() (interface{}, error) {
	return e.Element.GetProperty("text-stream-combiner")
}

func (e *Playbin) SetTextStreamCombiner(value interface{}) error {
	return e.Element.SetProperty("text-stream-combiner", value)
}

// uri (string)
//
// Set the next URI that playbin will play. This property can be set from the
// about-to-finish signal to queue the next media file.

func (e *Playbin) GetUri() (string, error) {
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

func (e *Playbin) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// video-filter (GstElement)
//
// the video filter(s) to apply, if possible

func (e *Playbin) GetVideoFilter() (interface{}, error) {
	return e.Element.GetProperty("video-filter")
}

func (e *Playbin) SetVideoFilter(value interface{}) error {
	return e.Element.SetProperty("video-filter", value)
}

// video-multiview-flags (GstVideoMultiviewFlags)
//
// Override details of the multiview frame layout

func (e *Playbin) GetVideoMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-flags")
}

func (e *Playbin) SetVideoMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("video-multiview-flags", value)
}

// video-multiview-mode (GstVideoMultiviewFramePacking)
//
// Re-interpret a video stream as one of several frame-packed stereoscopic modes.

func (e *Playbin) GetVideoMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-mode")
}

func (e *Playbin) SetVideoMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("video-multiview-mode", value)
}

// video-sink (GstElement)
//
// the video output element to use (NULL = default sink)

func (e *Playbin) GetVideoSink() (interface{}, error) {
	return e.Element.GetProperty("video-sink")
}

func (e *Playbin) SetVideoSink(value interface{}) error {
	return e.Element.SetProperty("video-sink", value)
}

// video-stream-combiner (GstElement)
//
// Get or set the current video stream combiner. By default, an input-selector
// is created and deleted as-needed.

func (e *Playbin) GetVideoStreamCombiner() (interface{}, error) {
	return e.Element.GetProperty("video-stream-combiner")
}

func (e *Playbin) SetVideoStreamCombiner(value interface{}) error {
	return e.Element.SetProperty("video-stream-combiner", value)
}

// vis-plugin (GstElement)
//
// the visualization element to use (NULL = default)

func (e *Playbin) GetVisPlugin() (interface{}, error) {
	return e.Element.GetProperty("vis-plugin")
}

func (e *Playbin) SetVisPlugin(value interface{}) error {
	return e.Element.SetProperty("vis-plugin", value)
}

// volume (float64)
//
// Get or set the current audio stream volume. 1.0 means 100%,
// 0.0 means mute. This uses a linear volume scale.

func (e *Playbin) GetVolume() (float64, error) {
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

func (e *Playbin) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

